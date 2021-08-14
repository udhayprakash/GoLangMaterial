package handlers

import (
	"encoding/json"
	"net/http"
)

//ListAccounts is used to get the account details
func ListAccounts(w http.ResponseWriter, r *http.Request) {
	JSONResp, err := json.Marshal(&accounts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(JSONResp)
}
