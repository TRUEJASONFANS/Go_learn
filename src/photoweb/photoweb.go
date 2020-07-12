package  main

import (
	"html/template"
	"log"
	"net/http"
)

func uploaderHandler( w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		t, err := template.ParseFiles("upload.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
		return
	}

}


func main() {
	http.HandleFunc("/upload", uploaderHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}