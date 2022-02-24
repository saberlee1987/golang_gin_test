package soap

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func soapCall(ws string, action string, payloadInterface interface{}) ([]byte, error) {
	v := &Envelope{
		XmlnsSoapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsUniv:    "http://services.soap.spring_camel_cxf_soap_provider.saber.com/",
		Header: &Header{
			WsseSecurity: &WsseSecurity{
				MustUnderstand: "1",
				XmlnsWsse:      "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
				XmlnsWsu:       "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
				UsernameToken: &UsernameToken{
					WsuId:    "UsernameToken-1",
					Username: &Username{},
					Password: &Password{
						Type: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText",
					},
				},
			},
		},
	}

	v.Header.WsseSecurity.UsernameToken.Username.Value = "saber66"
	v.Header.WsseSecurity.UsernameToken.Password.Value = "saber@123"
	v.Body = &Body{
		Payload: payloadInterface,
	}

	payload, err := xml.MarshalIndent(v, "", "  ")

	fmt.Println(string(payload))

	timeout := time.Duration(30 * time.Second)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}

	req, err := http.NewRequest("POST", ws, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "text/xml, multipart/related")
	req.Header.Set("SOAPAction", action)
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	rawBody, err := ioutil.ReadAll(res.Body)
	if len(rawBody) == 0 {
		return nil, errors.New("Empty response")
	}

	soapResponse, err := SoapFomMTOM(rawBody)
	if err != nil {
		return nil, err
	}

	return soapResponse, nil
}
func SoapFomMTOM(soap []byte) ([]byte, error) {
	reg := regexp.MustCompile(`(?ims)<[env:|soap:].+Envelope>`)
	s := reg.FindString(string(soap))
	if s == "" {
		return nil, errors.New("Response without soap envelope")
	}

	return []byte(s), nil
}

func checkFault(soapResponse []byte) error {
	xmlEnvelope := ResponseEnvelope{}

	err := xml.Unmarshal(soapResponse, &xmlEnvelope)
	if err != nil {
		return err
	}

	fault := xmlEnvelope.ResponseBodyBody.Fault
	if fault.XMLName.Local == "Fault" {
		sFault := fault.Code + " | " + fault.FaultString + " | " + fault.Actor + " | " + fault.Detail + "\n"
		return errors.New(sFault)
	}

	return nil
}
func CallHandleResponse(ws string, action string, payloadInterface interface{}, result interface{}) error {
	body, err := soapCall(ws, action, payloadInterface)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	return nil
}
