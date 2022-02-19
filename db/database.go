package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com.saber/golang_gin_test/dto"
	_ "github.com/go-sql-driver/mysql"
)

func connectToDataBase() (*sql.DB, error) {
	return sql.Open("mysql", "saber66:AdminSaber66@tcp(localhost:3306)/test2")
}

func FindAllPersons() ([]dto.Person, error) {
	db, err := connectToDataBase()
	if err != nil {
		return nil, errors.New("can not connect to database " + err.Error())
	}
	var persons []dto.Person

	rows, err := db.Query("select firstname,lastname,nationalCode,age,email,mobile from persons")

	if err != nil {
		return nil, errors.New("can not fetch data from database " + err.Error())
	}
	person := dto.Person{}
	//var id = 0
	for rows.Next() {
		err := rows.Scan(&person.Firstname, &person.LastName, &person.NationalCode, &person.Age, &person.Email, &person.Mobile)
		if err != nil {
			return nil, errors.New("can not fetch data from database " + err.Error())
		}
		persons = append(persons, person)
	}

	return persons, nil
}

func AddPerson(person dto.Person) (*dto.AddPersonsResponseDto, *dto.ErrorResponseDto) {
	db, err := connectToDataBase()
	var errorResponse dto.ErrorResponseDto
	if err != nil {
		errorResponse.Code = -1
		errorResponse.Text = err.Error()
		return nil, &errorResponse
	}
	response, _ := FindPersonByNationalCode(person.NationalCode)
	if response != nil {
		errorResponse.Code = -1
		errorResponse.Text = fmt.Sprintf("person with nationalCode %s already exist", person.NationalCode)
		return nil, &errorResponse
	}
	statement, err := db.Prepare("insert into persons (age, email, firstname, lastname, nationalCode, mobile) VALUES (?,?,?,?,?,?)")
	if err != nil {
		errorResponse.Code = -1
		errorResponse.Text = err.Error()
		return nil, &errorResponse
	}
	result, err := statement.Exec(person.Age, person.Email, person.Firstname, person.LastName, person.NationalCode, person.Mobile)
	if err != nil {
		errorResponse.Code = -1
		errorResponse.Text = err.Error()
		return nil, &errorResponse
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		errorResponse.Code = -1
		errorResponse.Text = err.Error()
		return nil, &errorResponse
	}
	if rowAffected > 0 {
		addResponse := dto.AddPersonsResponseDto{Code: 0, Text: "your Data is saved successfully"}
		return &addResponse, nil
	} else {
		errorResponse.Code = -1
		errorResponse.Text = "sorry can not insert new person to database"
		return nil, &errorResponse
	}
}

func FindPersonByNationalCode(nationalCode string) (*dto.Person, *dto.ErrorResponseDto) {
	db, err := connectToDataBase()
	var errorResponse dto.ErrorResponseDto
	var person dto.Person
	if err != nil {
		errorResponse.Code = -1
		errorResponse.Text = err.Error()
		return nil, &errorResponse
	}
	statement, err := db.Prepare("select firstname,lastname,nationalCode,age,email,mobile from persons p where p.nationalCode=?")
	if err != nil {
		errorResponse.Code = -1
		errorResponse.Text = err.Error()
		return nil, &errorResponse
	}
	rows, err := statement.Query(nationalCode)
	if err != nil {
		errorResponse.Code = -1
		errorResponse.Text = err.Error()
		return nil, &errorResponse
	}
	if !rows.Next() {
		errorResponse.Code = -1
		errorResponse.Text = fmt.Sprintf("person with nationalCode %s does not exist", nationalCode)
		return nil, &errorResponse
	}
	//var id int
	err = rows.Scan(&person.Firstname, &person.LastName, &person.NationalCode, &person.Age, &person.Email, &person.Mobile)
	if err != nil {
		errorResponse.Code = -1
		errorResponse.Text = err.Error()
		return nil, &errorResponse
	}
	return &person, nil
}
