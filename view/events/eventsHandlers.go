package events

import (
	"net/http"
	"encoding/json"
)


func GetEventHandler(w http.ResponseWriter, r *http.Request){
	response := getEventResponse{}
	// if user request is okay we responde the task details 
	// from data base ... to be complete 
	response.name = "complete this project"
	response.date = "soooon :)))) "
	response.priority = 127
	 output, err := json.Marshal(response)
	 if err != nil {
		 http.Error(w, err.Error(), 500)
		 return
	 }
	 w.Write(output)
	 w.WriteHeader(200)
}
