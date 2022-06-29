package app

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
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

    a.Router = mux.NewRouter()  

	a.InitializeRoutes()
}

func (a *App) Run(addr string) { 
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedHeaders([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	credentials := handlers.AllowCredentials()
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(headers, methods, origins, credentials)(a.Router)))
}