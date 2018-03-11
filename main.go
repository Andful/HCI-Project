package main

import(
	_ "fmt"
	"net/http"
	"html/template"
)

func statistics(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/pages/statistics.tmpl")
	t, err = t.ParseGlob("./template/components/*.tmpl")
	if err != nil {
		panic(err)
		return
	}

	err = t.Execute(w,nil)

	if err != nil {
		panic(err)
		return
	}
}

func app(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/main.tmpl")
	t, err = t.ParseGlob("./template/components/*.tmpl")
	if err != nil {
		panic(err)
		return
	}

	err = t.Execute(w,nil)

	if err != nil {
		panic(err)
		return
	}
}

func main() {
	http.HandleFunc("/statistics", statistics)
		http.HandleFunc("/app", app)
        http.Handle("/", http.FileServer(http.Dir("./")))
        panic(http.ListenAndServe(":8080", nil))
}
