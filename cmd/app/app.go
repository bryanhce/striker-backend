package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//todo initalize to have params that are values from env
func (a *App) Initialize() { 
	connectionString := 
	"be73b65ac2b869:9e8a2047@tcp(us-cdbr-east-05.cleardb.net)/heroku_10638fdcb554c64"

    var err error
    a.DB, err = sql.Open("mysql", connectionString)
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("connected to database")
	// defer a.DB.Close()

    a.Router = mux.NewRouter()  

	a.InitializeRoutes()
}

//figure out why addr param not used
func (a *App) Run(addr string) { 
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}