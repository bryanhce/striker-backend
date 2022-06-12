package main

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/test", TestHandler).Methods("GET")
	a.Router.HandleFunc("/singleTask/{id:[A-Za-z0-9]+}", getSingleTask).Methods("GET")
	a.Router.HandleFunc("/singleTask", createSingleTask).Methods("POST")
	a.Router.HandleFunc("/singleTask/{id:[A-Za-z0-9]+}", updateSingleTask).Methods("PUT")
	a.Router.HandleFunc("/singleTask/{id:[A-Za-z0-9]+}", deleteSingleTask).Methods("DELETE")
	a.Router.HandleFunc("/taskList", getTaskList).Methods("GET")
}