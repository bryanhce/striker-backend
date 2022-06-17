package app

func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/test/{id:}", TestHandler).Methods("GET")
	a.Router.HandleFunc("/task-list/single-task/{id:[a-z0-9-]+}", a.GetSingleTask).Methods("GET")
	a.Router.HandleFunc("/task-list/single-task/{id:[a-z0-9-]+}", a.UpdateSingleTask).Methods("PUT")
	a.Router.HandleFunc("/task-list/single-task/{id:[a-z0-9-]+}", a.DeleteSingleTask).Methods("DELETE")
	a.Router.HandleFunc("/task-list/single-task", a.CreateSingleTask).Methods("POST") //has query params date as well
	a.Router.HandleFunc("/task-list/{userId:[A-Za-z0-9]+}", a.GetTaskList).Methods("GET") //has query params date as well
}

/*
* notes:
* id refers to task id and in regex expression no capital alphabets and need to include -
* for userId regex expression there is caps letters but no -
*/
