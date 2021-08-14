package models

//Account is a data model of the sample_fb service
type Account struct {
	ID          string   `json:"id,omitempty"`
	Firstname   string   `json:"firstname,omitempty"`
	Middlename  string   `json:"middlename,omitempty"`
	Lastname    string   `json:"lastname,omitempty"`
	PhoneNumber string   `json:"phonenumber,omitempty"`
	Email       string   `json:"email,omitempty"`
	Gender      string   `json:"gender,omitempty"`
	Password    string   `json:"password,omitempty"`
	Address     *Address `json:"address,omitempty"`
	OldPassword string   `json: "oldpassword,omitempty"`
	Status      string   `json: "status,omitempty"`
}
