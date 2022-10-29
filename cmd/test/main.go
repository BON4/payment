package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/BON4/payment/internal/domain"
	"github.com/BON4/payment/internal/pkg/csvdownload"
	"github.com/BON4/payment/internal/pkg/csvupload"

	_ "github.com/lib/pq"
)

var filepath = "example.csv"

func main() {

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/payment_test?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	uploader := csvupload.NewCSVUploader[domain.TransactonHistory](conn, "transacton_history")

	file, err := os.OpenFile(filepath, os.O_RDONLY, 0444)
	if err != nil {
		panic(err)
	}

	par, err := domain.NewTxHitoryUnmarshaller(file)
	if err != nil {
		panic(err)
	}

	n, err := uploader.Upload(context.Background(), "boil", par)
	if err != nil {
		panic(err)
	}

	fmt.Println(n)

	file.Close()

	file, err = os.OpenFile("new_example.csv", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}

	marshaler := domain.NewTxHitoryMarshaller(file)
	downloader := csvdownload.NewCSVDownloader[domain.TransactonHistory](conn, "transacton_history", "tag")

	err = downloader.Download(context.Background(), marshaler)
	if err != nil {
		panic(err)
	}

	file.Close()
}
