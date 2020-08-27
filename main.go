package main

import(
	"./routing"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"time"
)
func main(){
	routing.ApiMapper()

	srv := &http.Server{
		//Handler: handlers.LoggingHandler(logger.GeneralLogger.Writer(), router.Router),
		Handler:      handlers.LoggingHandler(os.Stdout, routing.Router),
		Addr:         "localhost:" + "8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()

}