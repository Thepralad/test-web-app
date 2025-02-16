package main

import(
	"html/template"
	"net/http"
)



func main(){
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/notfound", notFoundHandler)
	http.ListenAndServe(":8080", nil)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "pagenotfound.html")
}
func homeHandler(res http.ResponseWriter, req *http.Request){
	if req.URL.Path != "/"{
		http.Redirect(res, req, "/notfound", http.StatusSeeOther)
	}
	page := struct{
		Title string
		Body string
	}{
		Title: "Welcome",
		Body: "This is a just a test for deploying my first Go app",
	}
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(res, page)
}