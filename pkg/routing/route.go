package routing

import (
	"github.com/gorilla/mux"
	"gomicro1.com/assisgnment/courses/pkg/modules"
)

//Function CoursesRouteHandler instantiate the CRUD operations
func CoursesRouteHandler(s *mux.Router) {

	s.HandleFunc("/course", modules.CreateCourse).Methods("POST")
	s.HandleFunc("/course", modules.GetAllCourse).Methods("GET")
	s.HandleFunc("/course/{ID}", modules.GetCourseByID).Methods("GET")
	s.HandleFunc("/course/{ID}", modules.UpdateCourse).Methods("PUT")
	s.HandleFunc("/course/{ID}", modules.DeleteCourse).Methods("DELETE")
}
