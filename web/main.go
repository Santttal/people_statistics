package main

import (
	"fmt"
	"github.com/Santttal/people-statistics/web/request"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server..")
	router := createRouter()
	handler := request.ReportHandler{CsvReader: createReader()}
	router.Handle("/add-report", handler).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8084", router))
}

func createRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}

func createReader() request.PersonRecordsReader {
	return request.CsvReaderWrapper{}
}
