package main

import (
	"os"
    "net/http"
    "html/template"
    "github.com/nlittlepoole/horcrux"
)

var AppTemplatePath string

func runTemplatedApp(w http.ResponseWriter, req *http.Request) {
    parsedTemplate, _ := template.ParseFiles(AppTemplatePath)
    if err := parsedTemplate.Execute(w, horcrux.GetEnvironmentHorcrux()); err != nil {
    	panic(err)
    }
}

func main() {
	AppTemplatePath = os.Args[1]
    http.HandleFunc("/app", runTemplatedApp)
    http.ListenAndServe(":8090", nil)
}