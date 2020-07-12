package  main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)
const (
	TEMPLATE_DIR = "./views"
)

var templates = make(map[string] *template.Template)

func init() {
	fileInfoArr, err := ioutil.ReadDir("./views")
	fmt.Println("init run")
	if err != nil {
		panic(err)
		return
	}
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		fmt.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templatePath] = t
	}
}

func uploaderHandler( w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		if err := renderHtml(w, "./views/upload.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
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

func logError(w http.ResponseWriter, err error)  {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}

func listViewHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	if (err != nil) {
		logError(w, err)
		return
	}
	locals := make(map[string] interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	renderHtml(w, "./views/list.html", locals)

}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) error {
	err := templates[tmpl].Execute(w, locals)
	return err
}

func main() {
	fmt.Println("main start")
	http.HandleFunc("/upload", uploaderHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/listview", listViewHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}