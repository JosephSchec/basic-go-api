package main

import (
	"fmt"
	"net/http"

	"github.com/josephschec/basic-go-api/internal/routes"
)

func main() {
	r := routes.NewRouter()
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Listening on port localhost%s\n", addr)

	err := http.ListenAndServe(addr, r)
	if err != nil {
		panic(err)
	}
}
