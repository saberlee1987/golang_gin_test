package main

import (
	"fmt"
	"github.com.saber/golang_gin_test/db"
	docs "github.com.saber/golang_gin_test/docs"
	"github.com.saber/golang_gin_test/dto"
	"github.com.saber/golang_gin_test/hi"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"runtime"
)

// @title saber golang gin
// @version 1.0.0-1400/11/26
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email saberazizi66@yahoo.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5000
// @BasePath /
// @schemes http
func main() {
	fmt.Println("Hello World !!!!!!!")
	fmt.Println(hi.SayHello("Saber", "Azizi"))
	router := gin.Default()
	docs.SwaggerInfo_swagger.BasePath = "/"
	docs.SwaggerInfo_swagger.Title = "golang gin swagger"
	url := ginSwagger.URL("http://localhost:5000/swagger/v3/api-docs/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/hello", hello)
	router.GET("/os", os)
	router.GET("/findAll", findAllPerson)
	router.POST("/add", addPerson)
	router.Run(":5000")
}

// HealthCheck godoc
// @Summary hello
// @Description get the status of server.
// @Tags hello
// @Accept */*
// @Param firstName query string true "firstName param"
// @Param lastName query string true "lastName param"
// @Produce json
// @Success 200 {object} dto.HelloDto
// @Router /hello [get]
func hello(context *gin.Context) {
	firstName := context.Query("firstName")
	lastName := context.Query("lastName")
	message := fmt.Sprintf("Hello %s %s", firstName, lastName)
	helloDto := dto.HelloDto{
		Message: message,
	}
	context.JSON(200, helloDto)
}

// HealthCheck godoc
// @Summary os
// @Description get the status of server.
// @Tags os
// @Accept */*
// @Produce json
// @Success 200 {string} string
// @Router /os [get]
func os(context *gin.Context) {
	context.String(200, runtime.GOOS)
}

// HealthCheck godoc
// @Summary find All person
// @Description get the status of server.
// @Tags findAllPerson
// @Accept */*
// @Produce json
// @Success 200 {object}  dto.FindAllPersonResponse
// @Router /findAll [get]
func findAllPerson(context *gin.Context) {
	persons, err := db.FindAllPersons()
	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{
			"error": err,
		})
	}
	response := dto.FindAllPersonResponse{
		Persons: persons,
	}
	context.JSON(200, response)
}

// HealthCheck godoc
// @Summary add person
// @Description post the status of server.
// @Tags addPerson
// @Accept application/json
//@Param personDto body dto.Person true "person body"
// @Produce json
// @Success 200 {object}  dto.AddPersonsResponseDto
// @Failure 400,404,406,500,504 {object} dto.ErrorResponseDto
// @Router /add [post]
func addPerson(context *gin.Context) {
	var person dto.Person
	err := context.ShouldBindJSON(&person)
	var errorResponseDto dto.ErrorResponseDto
	if err != nil {
		errorResponseDto.Code = -1
		errorResponseDto.Text = err.Error()
		context.JSON(http.StatusBadRequest, err)
	}
	fmt.Printf("Request for add person with body ===> %s\n", person)
	addPersonResponseDto, errorResponse := db.AddPerson(person)
	if errorResponse != nil {
		fmt.Printf("Error for add person with body ===> %s\n", errorResponse)
		context.JSON(http.StatusNotAcceptable, errorResponse)
	} else {
		fmt.Printf("Response for add person with body ===> %s\n", addPersonResponseDto)
		context.JSON(200, addPersonResponseDto)
	}
}
