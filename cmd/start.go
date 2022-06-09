package cmd

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gomicro1.com/assisgnment/courses/pkg/routing"
)

func Start() {

	s := mux.NewRouter()
	routing.CoursesRouteHandler(s)
	http.Handle("/", s)
	fmt.Println("Listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", s))
}
