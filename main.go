package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

//Models
//todo add type of task
type SingleTask struct {
	ID string `json:"id"`
	Description string `json:"description"`
	IsCompleted bool `json:"isCompleted"`
	Effort int `json:"effort"`
}

type TaskList struct {
	UID string `json:"uid"`
	AllTasks []SingleTask `json:"allTasks"`
}

//temp database
var tasks = []SingleTask{
	{ID: "1", Description: "cs2030s lab6", IsCompleted: false, Effort: 50},
	{ID: "2", Description: "Wish Bryan Happy Brithday", IsCompleted: false, Effort: 5},
	{ID: "3", Description: "Run 5km", IsCompleted: true, Effort: 20},
	{ID: "4", Description: "Buy a new pen", IsCompleted: false, Effort: 10},

}

var IDCounter = 5

func main() {
	fmt.Println("code is running")

	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")


	router := mux.NewRouter()

	router.HandleFunc("/test", testHandler).Methods("GET")
	router.HandleFunc("/singleTask/{id}", getSingleTask).Methods("GET")
	router.HandleFunc("/singleTask", createSingleTask).Methods("POST")
	router.HandleFunc("/singleTask/{id}", updateSingleTask).Methods("PUT")
	router.HandleFunc("/singleTask/{id}", deleteSingleTask).Methods("DELETE")
	router.HandleFunc("/taskList", getTaskList).Methods("GET")
	
	http.ListenAndServe(":8000", router)
}

//handlers
func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode( struct {
		Message string
	}{"test successful"})
}

func getTaskList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getSingleTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for _, task := range tasks {
		if task.ID == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}	
	//if reach here means id invalid
	w.WriteHeader(400)
	w.Write([]byte("Invalid task ID"))
}

func createSingleTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var newTask SingleTask
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if (err != nil) {
		w.WriteHeader(400)
		w.Write([]byte("Unable to decode post body"))
	}
	newTask.ID = strconv.Itoa(IDCounter) //temporary fix
	IDCounter++
	tasks = append(tasks, newTask)
	json.NewEncoder(w).Encode(newTask)
}

func updateSingleTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var idParam string = mux.Vars(r)["id"]
	var updatedTask SingleTask
	_ = json.NewDecoder(r.Body).Decode(&updatedTask)
	for i, t := range tasks {
		if t.ID == idParam {
			updatedTask.ID = idParam
			tasks[i] = updatedTask
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}
	//if reach here means id invalid
	w.WriteHeader(400)
	w.Write([]byte("Invalid task ID"))
}

func deleteSingleTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	//todo in the future need proper error handling and
	//this is inefficient code
	for index, task := range tasks {
		if task.ID == params["id"] {
			//this is slicing in go
			tasks = append(tasks[:index], tasks[index + 1:]...)
			json.NewEncoder(w).Encode(tasks)
			return
		}
	}
	//if reach here means id invalid
	w.WriteHeader(400)
	w.Write([]byte("Invalid task ID"))
	
}