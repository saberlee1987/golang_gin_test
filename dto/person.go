package dto

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Firstname    string `json:"firstname"  binding:"required"`
	LastName     string `json:"lastname"  binding:"required"`
	NationalCode string `json:"nationalCode"  binding:"required,min=10,max=10"`
	Age          int    `json:"age"  binding:"required"`
	Email        string `json:"email"  binding:"required,email"`
	Mobile       string `json:"mobile"  binding:"required"`
}

func (p Person) String() string {
	marshal, err := json.Marshal(p)
	if err != nil {
		return fmt.Sprintf("{\"firstname\":\"%s\",\"lastname\":\"%s\",\"nationalCode\":\"%s\",\"age\":%d,\"email\":\"%s\",\"mobile\":\"%s\"}",
			p.Firstname, p.LastName, p.NationalCode, p.Age, p.Email, p.Mobile)
	}
	return string(marshal)
}
