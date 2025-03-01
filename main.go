package main

import (
	"app/handler"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
)

func main() {
	fmt.Println("Running on port 8080")
	http.ListenAndServe(":8080", handler.GetRouter())
}
