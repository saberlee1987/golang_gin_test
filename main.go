package main

import (
	"fmt"
	"github.com.saber/golang_gin_test/db"
	docs "github.com.saber/golang_gin_test/docs"
	"github.com.saber/golang_gin_test/dto"
	"github.com.saber/golang_gin_test/hi"
	"github.com.saber/golang_gin_test/soap"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"runtime"
)

// @title saber golang gin
// @version 1.0.0-1400/11/30
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
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	fmt.Println("Hello World !!!!!!!")
	fmt.Println(hi.SayHello("Saber", "Azizi"))

	result := soap.ResponseEnvelope{}
	err := soap.CallHandleResponse("https://localhost:9091/soapServices/personSoap", "FindAll", &soap.FindAll{}, &result)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)

	//
	//	findByNationalCode:=soap.FindPersonByNationalCode{
	//		Payload: &soap.NationalCode{
	//			Payload: "0079028748",
	//		},
	//	}
	//
	//err = soap.CallHandleResponse("https://localhost:9091/soapServices/personSoap", "FindByNationalCode", findByNationalCode, &result)
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//fmt.Println(result)

	router := gin.Default()

	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(router)

	docs.SwaggerInfo_swagger.BasePath = "/"
	docs.SwaggerInfo_swagger.Title = "golang gin swagger"
	url := ginSwagger.URL("http://localhost:5000/swagger/v3/api-docs/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/hello", hello)
	router.GET("/os", os)
	personRoute := router.Group("/person")
	{
		personRoute.GET("/findAll", findAllPerson)
		personRoute.GET("/find/:nationalCode", findPersonByNationalCode)
		personRoute.PUT("/update/:nationalCode", updatePersonByNationalCode)
		personRoute.DELETE("/delete/:nationalCode", deletePersonByNationalCode)
		personRoute.POST("/add", addPerson)
	}

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
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
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
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func os(context *gin.Context) {
	context.String(200, runtime.GOOS)
}

// HealthCheck godoc
// @Summary findAllPerson
// @Description get the status of server.
// @Tags person api
// @Accept */*
// @Produce json
// @Success 200 {object}  dto.FindAllPersonResponse
// @Router /person/findAll [get]
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func findAllPerson(context *gin.Context) {
	var persons []dto.Person
	err := db.FindAllPersonsOrm(&persons)
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
// @Summary findPersonByNationalCode
// @Description get the status of server.
// @Tags person api
// @Accept */*
// @Param nationalCode path string true "nationalCode param"
// @Produce json
// @Success 200 {object}  dto.PersonDto
// @Failure 400,404,406,500,504 {object} dto.ErrorResponseDto
// @Router /person/find/{nationalCode} [get]
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func findPersonByNationalCode(context *gin.Context) {
	nationalCode := context.Param("nationalCode")
	personDto, err := db.FindPersonByNationalCodeOrm(nationalCode)
	if err != nil {
		fmt.Println(err)
		context.JSON(406, err)
		return
	}
	context.JSON(200, personDto)
}

// HealthCheck godoc
// @Summary deletePersonByNationalCode
// @Description get the status of server.
// @Tags person api
// @Accept */*
// @Param nationalCode path string true "nationalCode param"
// @Produce json
// @Success 200 {object}  dto.DeletePersonResponseDto
// @Failure 400,404,406,500,504 {object} dto.ErrorResponseDto
// @Router /person/delete/{nationalCode} [delete]
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func deletePersonByNationalCode(context *gin.Context) {
	nationalCode := context.Param("nationalCode")
	response, err := db.DeletePersonByNationalCodeOrm(nationalCode)
	if err != nil {
		fmt.Println(err)
		context.JSON(406, err)
		return
	}
	context.JSON(200, response)
}

// HealthCheck godoc
// @Summary update person by nationalCode
// @Description put the status of server.
// @Tags person api
// @Accept */*
// @Param nationalCode path string true "nationalCode param"
//@Param personDto body dto.PersonDto true "person body"
// @Produce json
// @Success 200 {object}  dto.UpdatePersonsResponseDto
// @Failure 400,404,406,500,504 {object} dto.ErrorResponseDto
// @Router /person/update/{nationalCode} [put]
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func updatePersonByNationalCode(context *gin.Context) {
	nationalCode := context.Param("nationalCode")

	var person dto.PersonDto
	err := context.ShouldBindJSON(&person)
	if err != nil {
		var errorResponseDto dto.ErrorResponseDto
		errorResponseDto.Code = -1
		errorResponseDto.Text = "BadRequest"
		var validations []dto.ValidationDto
		for _, fieldErr := range err.(validator.ValidationErrors) {
			validation := dto.ValidationDto{}
			validation.FieldName = fieldErr.Field()
			validation.DetailMessage = fmt.Sprintf("Error for %s actual value %s is %s your input %v", fieldErr.StructField(), fieldErr.ActualTag(), fieldErr.Param(), fieldErr.Value())
			validations = append(validations, validation)
		}
		errorResponseDto.Validations = validations
		fmt.Printf("Error for binding json to person with error %s\n", errorResponseDto)
		context.JSON(http.StatusBadRequest, errorResponseDto)
		return
	}
	fmt.Printf("Request for update person with nationalCode %s with  body ===> %s\n", nationalCode, person)

	response, errorResponseDto := db.UpdatePersonOrm(person, nationalCode)
	if errorResponseDto != nil {
		fmt.Println(errorResponseDto)
		context.JSON(406, errorResponseDto)
		return
	}
	fmt.Printf("Request for update person with nationalCode %s with respone body ===> %s\n", nationalCode, response)
	context.JSON(200, response)
}

// HealthCheck godoc
// @Summary add person
// @Description post the status of server.
// @Tags person api
// @Accept application/json
//@Param personDto body dto.PersonDto true "person body"
// @Produce json
// @Success 200 {object}  dto.AddPersonsResponseDto
// @Failure 400,404,406,500,504 {object} dto.ErrorResponseDto
// @Router /person/add [post]
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func addPerson(context *gin.Context) {
	var person dto.PersonDto

	err := context.ShouldBindJSON(&person)
	var errorResponseDto dto.ErrorResponseDto
	if err != nil {
		errorResponseDto.Code = -1
		errorResponseDto.Text = "BadRequest"
		var validations []dto.ValidationDto
		for _, fieldErr := range err.(validator.ValidationErrors) {
			validation := dto.ValidationDto{}
			validation.FieldName = fieldErr.Field()
			validation.DetailMessage = fmt.Sprintf("Error for %s actual value %s is %s your input %v", fieldErr.StructField(), fieldErr.ActualTag(), fieldErr.Param(), fieldErr.Value())
			validations = append(validations, validation)
		}
		errorResponseDto.Validations = validations
		fmt.Printf("Error for binding json to person with error %s\n", errorResponseDto)
		context.JSON(http.StatusBadRequest, errorResponseDto)
		return
	}
	fmt.Printf("Request for add person with body ===> %s\n", person)
	addPersonResponseDto, errorResponse := db.AddPersonOrm(person)
	if errorResponse != nil {
		fmt.Printf("Error for add person with body ===> %s\n", errorResponse)
		context.JSON(http.StatusNotAcceptable, errorResponse)
	} else {
		fmt.Printf("Response for add person with body ===> %s\n", addPersonResponseDto)
		context.JSON(200, addPersonResponseDto)
	}
}
