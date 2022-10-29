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

	group.POST("/upload/:file_name", t.Upload)
	group.GET("/csv", t.Download)
	group.GET("/json", t.List)
}

func (t *TransctionHandler) Upload(ctx *gin.Context) {
	n := ctx.Param("file_name")
	if len(n) < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no filename provided"})
		return
	}

	h, err := ctx.FormFile(n)
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

	ctx.JSON(http.StatusOK, gin.H{"objectCount": count})
}

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
