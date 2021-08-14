package handlers

import(
	"encoding/json"
"io/ioutil"
"net/http"
"github.com/gorilla/mux"
"github.com/sample-fb/models"
)
//UpdateAccount is used to update the account details
func UpdateAccount(w http.ResponseWriter, req *http.Request){
	// read url params of the request
	urlParams := mux.Vars(req)
	var account models.Account
	var isUpdate bool
	//read the body of the request
	body, err:=ioutil.ReadAll(req.Body)
		if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
		}

		if err= json.Unmarshal(body, &account); err!=nil{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		for i:= range accounts{
			if urlParams["id"] == accounts[i].ID{
				if account.Firstname != ""{
					accounts[i].Firstname = account.Firstname
					isUpdate = true
				}
				if account.Middlename != ""{
					accounts[i].Middlename = account.Middlename
					isUpdate = true
				}
				if account.Lastname != ""{
					accounts[i].Lastname = account.Lastname
					isUpdate = true
				}
				if account.PhoneNumber != ""{
					accounts[i].PhoneNumber = account.PhoneNumber
					isUpdate = true
				}
				if account.Email!= ""{
					accounts[i].Email= account.Email
					isUpdate = true
				}
				if account.Password != ""{
					accounts[i].Password = account.Password
					isUpdate = true
				}
				if account.Gender != ""{
					accounts[i].Gender = account.Gender
					isUpdate = true
				}
				if account.Address != nil {
					if account.Address.HouseNumber != ""{
						accounts[i].Address.HouseNumber = account.Address.HouseNumber
						isUpdate = true
					}
					if account.Address.Street != ""{
						accounts[i].Address.Street = account.Address.Street
						isUpdate = true
					}
					if account.Address.City != ""{
						accounts[i].Address.City = account.Address.City
						isUpdate = true
					}
					if account.Address.Zipcode != ""{
						accounts[i].Address.Zipcode = account.Address.Zipcode
						isUpdate = true
					}
				}


			}
		}
		if isUpdate == false{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte ("No changed detected"))
		return

		}
		w.WriteHeader(http.StatusNoContent)
		return
}