package main

// @title Publicis Project API
// @version 1.0
// @description This is a sample service for managing Publicis project
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email visvishwa@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5000
// @BasePath /

import (
	_ "awesomeProject/docs" // docs is generated by Swag CLI, you have to import it.
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	_ "strconv"

	"net/http"
)

type PublicisProject struct {
	Project    string `json:"project"`
	Technology string `json:"technology"`
	Location   string `json:"location"`
}

var projectArray []PublicisProject

// UpdatefUNCTION godoc
// @Summary Update project identified by the given project_name
// @Description Update the project corresponding to the input project_name
// @Tags project
// @Accept  json
// @Produce  json
// @Param project_name path string true "project_name of the project to be updated"
// @Success 200 {object} PublicisProject
// @Router /updateTechnology/{project_name} [post]
func updateFunction(writer http.ResponseWriter, request *http.Request) {
	param := mux.Vars(request)

	for index, vals := range projectArray {
		if vals.Project == param["project_name"] {
			projectArray = append(projectArray[:index], projectArray[index+1:]...)

			var psproj PublicisProject
			json.NewDecoder(request.Body).Decode(&psproj)
			psproj.Project = param["project_name"]
			projectArray = append(projectArray, psproj)
			json.NewEncoder(writer).Encode(projectArray)
			return
		}
	}
	json.NewEncoder(writer).Encode("No project found with specified project name")

}

// DeleteFunction godoc
// @Summary Delete order identified by the given project_name
// @Description Delete the project corresponding to the input project_name
// @Tags project
// @Accept  json
// @Produce  json
// @Param project_name path string true "name of the project to be deleted"
// @Success 204 "No Content"
// @Router /deleteTechnology/{project_name} [delete]
func deleteFunction(writer http.ResponseWriter, request *http.Request) {

	var param = mux.Vars(request)
	for index, values := range projectArray {
		if values.Project == param["project_name"] {
			projectArray = append(projectArray[:index], projectArray[index+1:]...)
			return
		}
	}
	json.NewEncoder(writer).Encode("No Project found with given project name")
}

// CreateFunction godoc
// @Summary Create a new Project
// @Description Create a new Project with the input paylod
// @Tags project
// @Accept  json
// @Produce  json
// @Param project body PublicisProject true "Create Project"
// @Success 200 {object} PublicisProject
// @Router /createTechnology [post]
func createFunction(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("Authorization") == "VishwaRatna" {
		writer.Header().Set("Content-Type", "Application/Json")
		var newProject PublicisProject
		json.NewDecoder(request.Body).Decode(&newProject)

		projectArray = append(projectArray, newProject)
		json.NewEncoder(writer).Encode(projectArray)
		return
	}
	json.NewEncoder(writer).Encode("Not Authenticated")

}

// GetFunction godoc
// @Summary Get details of all Projects in Publicis Sapient
// @Description Get details of all Projects in Publicis Sapient
// @Tags project
// @Accept  json
// @Produce  json
// @Success 200 {array} PublicisProject
// @Router /getTechnology [get]
func getFunction(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(projectArray)

	fmt.Println("some")

}

func main() {

	router := mux.NewRouter()

	projectArray = append(projectArray, PublicisProject{
		Project:    "Banking",
		Technology: "Golang",
		Location:   "India",
	}, PublicisProject{
		Project:    "Telecommunication",
		Technology: "Java-8",
		Location:   "Australia",
	})

	fmt.Println("Listening on 5000 port.....")

	// Read-all
	router.HandleFunc("/getTechnology", getFunction).Methods("GET")

	// Create
	router.HandleFunc("/createTechnology", createFunction).Methods("POST")

	// Delete
	router.HandleFunc("/deleteTechnology/{project_name}", deleteFunction).Methods("DELETE")

	// Update
	router.HandleFunc("/updateTechnology/{project_name}", updateFunction).Methods("PUT")

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":5000", router))

}
