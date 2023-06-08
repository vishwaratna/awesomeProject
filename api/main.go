package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type PublicisProject struct {
	Project    string `json:"project"`
	Technology string `json:"technology"`
	Location   string `json:"location"`
}

var projectArray []PublicisProject

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

	fmt.Println("Listening on 8085 port.....")
	router.HandleFunc("/getTechnology", getFunction).Methods("GET")
	router.HandleFunc("/createTechnology", createFunction).Methods("POST")
	router.HandleFunc("/deleteTechnology/{project_name}", deleteFunction).Methods("DELETE")
	router.HandleFunc("/updateTechnology/{project_name}", updateFunction).Methods("PUT")

	http.ListenAndServe(":5000", router)

}

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

func getFunction(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(projectArray)

	fmt.Println("some")

}
