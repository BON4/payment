package http

import (
	"net/http"

	"github.com/BON4/payment/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TransctionHandler struct {
	tUc    domain.TxUsecase
	logger *logrus.Entry
}

func NewTransactionHandler(group *gin.RouterGroup, tUc domain.TxUsecase, logger *logrus.Entry) {
	t := &TransctionHandler{
		logger: logger,
		tUc:    tUc,
	}

	group.POST("/upload", t.Upload)
	group.GET("/csv", t.Download)
	group.GET("/json", t.List)
}

type uploadResponse struct {
	UploadedObjectCount int64 `json:"upload_object_count"`
}

// @Summary      Upload
// @Description  Uploads csv file and saves it in DB.
// @Tags         payments
// @Accept       mpfd
// @Produce      json
// @Param        file        formData  file    true  "provide csv file"
// @Success      200     {object}  uploadResponse
// @Failure      400     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /payments/upload [post]
func (t *TransctionHandler) Upload(ctx *gin.Context) {
	h, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f, err := h.Open()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer f.Close()

	count, err := t.tUc.CSVInsert(ctx, f)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, uploadResponse{UploadedObjectCount: count})
}

// @Summary      Download
// @Description  Downloads csv file.
// @Tags         payments
// @Produce      mpfd
// @Param        page_size         query     int              false "page size"
// @Param        page_number       query     int              false "page number"
// @Param        transaction_id    query     int              false "search by transaction_id"
// @Param        terminal_id       query     []int            false "search by terminal id"
// @Param        status            query     string           false "search by status"     Enums(accepted, declined)
// @Param        payment_type      query     string           false "search by payment_type  Enums(cash, card)"
// @Param        post_date_from    query     string           false "search objects starting from specified date"     Format(dateTime)
// @Param        post_date_to      query     string           false "search objects ending with specified date"       Format(dateTime)
// @Param        payment_narrative       query     string              false  "search by the partially specified payment_narrative"
// @Success      200
// @Failure      400     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /payments/csv [get]
func (t *TransctionHandler) Download(ctx *gin.Context) {
	form := domain.FindTxRequest{}
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileType := "text/csv" // set the type

	//TODO: get file size
	// fileSize := 512

	//Transmit the headers
	ctx.Header("Expires", "0")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Control", "private, no-transform, no-store, must-revalidate")
	ctx.Header("Content-Disposition", "attachment; filename="+"out.csv")
	ctx.Header("Content-Type", fileType)
	// ctx.Header("Content-Length", strconv.FormatInt(int64(fileSize), 10))

	if err := t.tUc.CSVRetrive(ctx.Request.Context(), form, ctx.Writer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// @Summary      List
// @Description  Retrives lis of json formaated objects
// @Tags         payments
// @Produce      json
// @Param        page_size         query     int              false "page size"
// @Param        page_number       query     int              false "page number"
// @Param        transaction_id    query     int              false "search by transaction_id"
// @Param        terminal_id       query     []int            false "search by terminal id"
// @Param        status            query     string           false "search by status"     Enums(accepted, declined)
// @Param        payment_type      query     string           false "search by payment_type  Enums(cash, card)"
// @Param        post_date_from    query     string           false "search objects starting from specified date"     Format(dateTime)
// @Param        post_date_to      query     string           false "search objects ending with specified date"       Format(dateTime)
// @Param        payment_narrative       query     string              false  "search by the partially specified payment_narrative"
// @Success      200
// @Failure      400     {object}  error
// @Failure      401     {object}  error
// @Failure      500     {object}  error
// @Router       /payments/json [get]
func (t *TransctionHandler) List(ctx *gin.Context) {
	form := domain.FindTxRequest{}
	if err := ctx.Bind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txs, err := t.tUc.List(ctx.Request.Context(), form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, txs)
}
