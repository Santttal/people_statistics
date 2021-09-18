package main

import (
	"encoding/json"
	"fmt"
	"github.com/Santttal/people-statistics/web/request"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server..")
	r := mux.NewRouter()
	r.HandleFunc("/add-report", ReportHandler).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8084", r))
}

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	csvReader := createReader()
	records, err := csvReader.Read(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, err := json.Marshal(records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(b)

	w.WriteHeader(http.StatusOK)
}

func createReader() request.PersonRecordsReader {
	return request.CsvReaderWrapper{}
}
