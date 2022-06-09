package modules

import (
	"gomicro1.com/assisgnment/courses/pkg/config"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"os"

	"github.com/gorilla/mux"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var dB *sql.DB

type Course struct {
	Title string `json:"Title"`
	ID    string `json:"ID"`
}

// Initialize the database
func init() {

	config.ConnectDB()
	dB = config.GetDB()

}

// Function AuthAPIKey validates APIKey for POST,PUT & DELETE methods.
func AuthAPIKey(r *http.Request) bool {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading file")
		os.Exit(1)
	}

	API_KEY := os.Getenv("API_KEY")

	v := r.URL.Query()
	if key, ok := v["key"]; ok {
		if key[0] == API_KEY {
			return true
		} else {
			return false
		}
	} else {
		return false
	}

}

// Fuction CreateCourse creates a new Course
func CreateCourse(w http.ResponseWriter, r *http.Request) {

	if !AuthAPIKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid API key"))
		return
	}

	query, err := dB.Prepare("INSERT INTO course (ID, Title) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	dataMap := make(map[string]string)
	json.Unmarshal(body, &dataMap)
	id := dataMap["ID"]
	title := dataMap["Title"]

	_, err = query.Exec(id, title)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New Course created")
}

// Function GetCoursebyID returns a Course based on id provided.
func GetCourseByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	query, err := dB.Query("SELECT * FROM course WHERE ID = ?", params["ID"])
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	// Course is object of Course struct
	var Course Course

	for query.Next() {
		err := query.Scan(&Course.ID, &Course.Title)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(Course)
}

// Function GetAllCourse returns all Courses stored in database.
func GetAllCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Courses []Course

	query, err := dB.Query("SELECT * FROM db_courses.course")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	for query.Next() {
		var Course Course
		err := query.Scan(&Course.ID, &Course.Title)
		if err != nil {
			panic(err.Error())
		}
		Courses = append(Courses, Course)

	}
	json.NewEncoder(w).Encode(Courses)
}

// Function UpdateCourse updates the Course Title based on id provided.
func UpdateCourse(w http.ResponseWriter, r *http.Request) {

	if !AuthAPIKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid API key"))
		return
	}
	params := mux.Vars(r)
	query, err := dB.Prepare("UPDATE course SET Title = ? WHERE ID = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	dataMap := make(map[string]string)
	json.Unmarshal(body, &dataMap)
	newTitle := dataMap["Title"]

	_, err = query.Exec(newTitle, params["ID"])
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	fmt.Fprintf(w, "Course %s updated", params["ID"])

}

// Fuction DeleteCourse deletes a Course using the course id provided
func DeleteCourse(w http.ResponseWriter, r *http.Request) {

	if !AuthAPIKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid API key"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	query, err := dB.Prepare("DELETE FROM course WHERE ID = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = query.Exec(params["ID"])
	if err != nil {
		panic(err.Error)
	}

	defer query.Close()
	fmt.Fprintf(w, "Course %s deleted", params["ID"])
}
