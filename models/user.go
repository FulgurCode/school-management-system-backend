package models

type Admin struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
}

type Teacher struct {
	Id       string `json:"id,omitempty"`
	Fullname string `json:"fullname,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
	Subject  string `json:"subject,omitempty"`
	Address  string `json:"address,omitempty"`
}
