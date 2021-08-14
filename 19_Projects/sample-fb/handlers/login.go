package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/sample-fb/models"
)

//LoginHandler is used to login to the account
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var acc models.Account

	// read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err = json.Unmarshal(body, &acc); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if acc.Email == "" || acc.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email and password are required to log in and can't be nill, "))
		return
	}
	for i := range accounts {
		// First check if email passed by user and that of present in the account are same
		if acc.Email == accounts[i].Email && acc.Password == accounts[i].Password {
			loginResponse := struct {
				LoginToken string `json: "logintoken"`
			}{
				//generate new id for the account
				LoginToken: uuid.New().String(),
			}

			resp, err := json.Marshal(&loginResponse)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(resp)
			return
		}

	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Invalid credentials"))
	return
}
