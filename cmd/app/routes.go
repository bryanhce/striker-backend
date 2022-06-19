package app

func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/test", TestHandler).Methods("GET")
	a.Router.HandleFunc("/task-list/single-task/{id:[a-z0-9-]+}", a.GetSingleTask).Methods("GET")
	a.Router.HandleFunc("/task-list/single-task/{id:[a-z0-9-]+}", a.UpdateSingleTask).Methods("PUT")
	a.Router.HandleFunc("/task-list/single-task/{id:[a-z0-9-]+}", a.DeleteSingleTask).Methods("DELETE")
	//has query param date
	a.Router.HandleFunc("/task-list/single-task", a.CreateSingleTask).Methods("POST") 
	//has query param date
	a.Router.HandleFunc("/task-list/{userId:[A-Za-z0-9]+}", a.GetTaskList).Methods("GET")
	//has query param year-month 
	a.Router.HandleFunc("/calendar/{userId:[A-Za-z0-9]+}", a.GetMonthlyTasks).Methods("GET") 
	//has query param start-date & end-date
	a.Router.HandleFunc("/analytics/{userId:[A-Za-z0-9]+}", a.GetAnalytics).Methods("GET") 
}

/*
* notes:
* id refers to task id and in regex expression no capital alphabets and need to include -
* for userId regex expression there is caps letters but no -
*/
