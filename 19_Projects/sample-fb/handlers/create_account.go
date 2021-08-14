package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/sample-fb/models"
)

var (
	accounts []models.Account
)

//CreateAccount is used to create an account
func CreateAccount(w http.ResponseWriter, req *http.Request) {
	//decode json data into our model
	var account models.Account

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err = json.Unmarshal(body, &account); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if account.Firstname == "" || account.Lastname == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Firstname and lastname are required"))
		return
	}
	// generate a new id for the account
	account.ID = uuid.New().String()
	accounts = append(accounts, account)
	w.WriteHeader(http.StatusCreated)
	return
}
