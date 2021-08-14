package handlers

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sample-fb/models"

)
//DeleteAccount is used to delete the account
func DeleteAccount(w http.ResponseWriter, r *http.Request){
	urlParams := mux.Vars(r)
	var temp []models.Account

	//get the accounts from List
	for i:= range accounts {
		if urlParams["id"] != accounts[i].ID{
			temp = append(temp, accounts[i])
		} else{
			fmt.Printf("Account %v deleted successfully\n", urlParams["id"])
		}
		accounts = temp
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(fmt.Sprintf("Account %v deleted successfully", urlParams["id"])))

	}
}