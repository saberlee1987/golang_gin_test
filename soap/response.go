package soap

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type ResponseEnvelope struct {
	XMLName          xml.Name     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	ResponseBodyBody ResponseBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

type ResponseBody struct {
	XMLName         xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Fault           Fault    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	FindAllResponse string   `xml:"http://services.soap.spring_camel_cxf_soap_provider.saber.com/ FindAllPersonResponse"`
}

type Fault struct {
	XMLName     xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault" json:"xmlname"`
	Code        string   `xml:"faultcode,omitempty" json:"code"`
	FaultString string   `xml:"faultstring,omitempty" json:"faultString"`
	Actor       string   `xml:"faultactor,omitempty" json:"actor"`
	Detail      string   `xml:"detail,omitempty" json:"detail"`
}

//type FindAllResponse struct {
//	XMLName     xml.Name `xml:"http://services.soap.spring_camel_cxf_soap_provider.saber.com/" json:"xmlname"`
//	Response []PersonSoap `xml:"response" json:"response"`
//}
//
//type PersonSoap struct {
//	Firstname string `xml:"firstname" json:"firstname"`
//	Lastname string `xml:"lastname" json:"lastname"`
//	NationalCode string `xml:"nationalCode" json:"nationalCode"`
//	Email string `xml:"email" json:"email"`
//	Age int `xml:"age" json:"age"`
//	Mobile int `xml:"mobile" json:"mobile"`
//}

func (f Fault) String() string {
	body, err := json.Marshal(f)
	if err != nil {
		return fmt.Sprintf("{\"code\":\"%s\",\"faultString\":\"%s\",\"actor\":\"%s\",\"detail\":\"%s\"}", f.Code,
			f.FaultString, f.Actor, f.Detail)
	}
	return string(body)
}
