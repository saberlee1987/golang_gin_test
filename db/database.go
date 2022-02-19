package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com.saber/golang_gin_test/dto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type dBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func buildDBConfig() *dBConfig {
	dbConfig := dBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "saber66",
		Password: "AdminSaber66",
		DBName:   "test2",
	}
	return &dbConfig
}

func dbURL(dbConfig *dBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func connectToDataBaseOrm() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dbURL(buildDBConfig()))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDataBase() (*sql.DB, error) {
	return sql.Open("mysql", "saber66:AdminSaber66@tcp(localhost:3306)/test2")
}

func FindAllPersonsOrm(persons *[]dto.Person) error {
	db, err := connectToDataBaseOrm()
	if err != nil {
		return err
	}
	err = db.Find(persons).Error
	if err != nil {
		return err
	}
	return nil
}

func FindAllPersons() ([]dto.PersonDto, error) {
	db, err := connectToDataBase()
	if err != nil {
		return nil, errors.New("can not connect to database " + err.Error())
	}
	var persons []dto.PersonDto

	rows, err := db.Query("select firstname,lastname,nationalCode,age,email,mobile from persons")

	if err != nil {
		return nil, errors.New("can not fetch data from database " + err.Error())
	}
	person := dto.PersonDto{}
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

func AddPerson(person dto.PersonDto) (*dto.AddPersonsResponseDto, *dto.ErrorResponseDto) {
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

func FindPersonByNationalCodeOrm(personDto *dto.PersonDto, nationalCode string) *dto.ErrorResponseDto {
	db, err := connectToDataBaseOrm()
	var errorResponse dto.ErrorResponseDto
	if err != nil {
		errorResponse.Code = -1
		errorResponse.Text = err.Error()
		return nil
	}
	err = db.Where("nationalCode =?", nationalCode).First(personDto).Error
	if err != nil {
		errorResponse.Code = -1
		errorResponse.Text = err.Error()
		return &errorResponse
	}
	return nil
}

func FindPersonByNationalCode(nationalCode string) (*dto.PersonDto, *dto.ErrorResponseDto) {
	db, err := connectToDataBase()
	var errorResponse dto.ErrorResponseDto
	var person dto.PersonDto
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
