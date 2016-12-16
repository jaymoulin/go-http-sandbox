package main
import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jaymoulin/http/app/handler"
)
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HomeHandler)
	r.HandleFunc("/products", handler.ProductsHandler)
	r.HandleFunc("/articles", handler.ArticlesHandler)
	r.HandleFunc("/choices", handler.ChoicesHandler)
	r.HandleFunc("/validate", handler.ValidateHandler)
	r.HandleFunc("/result", handler.ResultHandler)
	r.HandleFunc("/login", handler.LoginHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	log.Fatal(http.ListenAndServe(":8000", r))
}
