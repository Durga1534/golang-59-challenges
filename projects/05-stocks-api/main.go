package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Durga1534/stocks-api/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
