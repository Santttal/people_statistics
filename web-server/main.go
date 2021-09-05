package main

import (
	"encoding/json"
	"fmt"
	request_reader "github.com/Santttal/people-statistics/web-server/request-reader"
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

	csvReader := request_reader.CsvReaderWrapper{}
	records, err := csvReader.Read(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, err := json.Marshal(records)

	w.Write(b)

	w.WriteHeader(http.StatusOK)
}
