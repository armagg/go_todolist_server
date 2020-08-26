package routing
import(
	"github.com/gorilla/mux"
	"../view/auth"
	"../view/events"
)
var Router = mux.NewRouter()


func ApiMapper() {
	authHadler := Router.PathPrefix("/auth").Subrouter()	
	eventHandler := Router.PathPrefix("/event").Subrouter()
	authHadler.HandleFunc("/sign_up", auth.SignUpHandler).Methods("POST")
	eventHandler.HandleFunc("/get_event", events.GetEventHandler).Methods("GET")


}