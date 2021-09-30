package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	//"log"
	"net/http"
	"os"
	csv "path/handler"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/csv" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Asset not Found\n"))
		return
	}

	csv.Conv()

	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	w.Write([]byte(data))
}

func main() {

	http.HandleFunc("/", Template)

	http.HandleFunc("/csv", rootHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func Template(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("template/index.html"))
	/* if err != nil {
		log.Fatal(err)
	} */
	template.Execute(w, "wow")
}
