package dto

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id           uint   `json:"id" gorm:"column:id;primaryKey"`
	FirstName    string `json:"firstname" gorm:"column:firstname;varchar(70)"`              // column name is `firstname`
	LastName     string `json:"lastname" gorm:"column:lastname;varchar(80)"`                // column name is `lastname`
	NationalCode string `json:"nationalCode" gorm:"column:nationalCode;unique;varchar(10)"` // column name is `nationalCode`
	Email        string `json:"email" gorm:"column:email;varchar(50)"`
	Age          int    `json:"age" gorm:"column:age;int"`
	Mobile       string `json:"mobile" gorm:"column:mobile;varchar(15)"`
}

func (p *Person) TableName() string {
	return "persons"
}

func (p Person) String() string {
	marshal, err := json.Marshal(p)
	if err != nil {
		return fmt.Sprintf("{\"id\":%d,\"firstName\":\"%s\",\"lastName\":\"%s\",\"nationalCode\":\"%s\",\"age\":%d,\"email\":\"%s\",\"mobile\":\"%s\"}",
			p.Id, p.FirstName, p.LastName, p.NationalCode, p.Age, p.Email, p.Mobile)
	}
	return string(marshal)
}

type PersonDto struct {
	Firstname    string `json:"firstname"  binding:"required" example:"saber"`
	LastName     string `json:"lastname"  binding:"required" example:"azizi"`
	NationalCode string `json:"nationalCode"  binding:"required,min=10,max=10" example:"0079028748"`
	Age          int    `json:"age"  binding:"required" example:"35"`
	Email        string `json:"email"  binding:"required,email" example:"saberazizi66@yahoo.com"`
	Mobile       string `json:"mobile"  binding:"required" example:"09365627895"`
}

func (p PersonDto) String() string {
	marshal, err := json.Marshal(p)
	if err != nil {
		return fmt.Sprintf("{\"firstname\":\"%s\",\"lastname\":\"%s\",\"nationalCode\":\"%s\",\"age\":%d,\"email\":\"%s\",\"mobile\":\"%s\"}",
			p.Firstname, p.LastName, p.NationalCode, p.Age, p.Email, p.Mobile)
	}
	return string(marshal)
}
