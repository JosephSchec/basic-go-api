package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/api", apiHandler)
	return router
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "index page found")
}
func apiHandler(w http.ResponseWriter, r *http.Request) {
	apiDat := "some data"
	fmt.Fprintln(w, apiDat)

}
