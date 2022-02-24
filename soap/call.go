package soap

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func soapCall(ws string, action string, payloadInterface interface{}) ([]byte, error) {
	v := &Envelope{
		XmlnsSoapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsUniv:    "http://www.example.pl/ws/test/universal",
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
		XMLName: xml.Name{
			Local: "http://com.saber.spring_camel_cxf_soap_provider.soap.services/",
		},
	}

	payload, err := xml.MarshalIndent(v, "", "  ")

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

	rawbody, err := ioutil.ReadAll(res.Body)
	if len(rawbody) == 0 {
		return nil, errors.New("Empty response")
	}

	soapResponse, err := SoapFomMTOM(rawbody)
	if err != nil {
		return nil, err
	}

	// test for fault
	err = checkFault(soapResponse)
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
		sFault := fault.Code + " | " + fault.String + " | " + fault.Actor + " | " + fault.Detail
		return errors.New(sFault)
	}

	return nil
}
func SoapCallHandleResponse(ws string, action string, payloadInterface interface{}, result interface{}) error {
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
