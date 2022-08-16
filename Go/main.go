package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	Handlers "github.com/mcorengia1/go-backend-template/src/handlers"
)

func handleRequest(r *mux.Router) {

	getR := r.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/", Handlers.ExampleHandler)

	// See the documentation of the cors library there are many configurations according to your needs
	// Origins from all the combinations of the domains that are going to make the requests must be included
	// ie: http:xxxx.com / https:xxxx.com / http:www.xxxx.com / https:www.xxxx.com
	c := cors.New(cors.Options{
		//AllowCredentials: true,
		//AllowedOrigins: []string{"http://yourdomain.com", "https://yourdomain.com",
		//"http://www.yourdomain.com", "https://www.yourdomain.com"},
	})

	handler := c.Handler(r)
	err := http.ListenAndServe(":8000", handler)
	fmt.Println(err)
}

func main() {

	router := mux.NewRouter()
	handleRequest(router)

}
