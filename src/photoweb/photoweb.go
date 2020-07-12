package  main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
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
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		fileName := h.Filename
		t, err := os.Create("./uploads/" + fileName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _,err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id=" + fileName, http.StatusFound)
	}


}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := "./uploads/" + imageId
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}


func main() {
	http.HandleFunc("/upload", uploaderHandler)
	http.HandleFunc("/view", viewHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}