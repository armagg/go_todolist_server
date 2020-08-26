package auth

import (
	"net/http"
	"encoding/json"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request){
	response := signUpResponse{}
	// if user request is okay we can make a new user 
	// in our data base ... to be complete 
	response.Message = "its okay"
	response.Token = "adsfklhnavjkg"// its just a silly example
	 output, err := json.Marshal(response)
	 if err != nil {
		 http.Error(w, err.Error(), 500)
		 return
	 }
	 w.Write(output)
	 w.WriteHeader(200)
}

func SignInHandler(w http.ResponseWriter, r *http.Request){
	//to be complete
}

