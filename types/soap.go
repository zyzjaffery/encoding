package types

import (
	"encoding/xml"
	"errors"
	"reflect"
)

type SoapEnvelope struct {
	XMLName xml.Name   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  SoapHeader `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header" json:"Header"`
	Body    SoapBody   `xml: Body" json:"Body"`
}

type SoapFault struct {
	FaultCode   *string `xml:"faultcode" json:"faultcode"`
	FaultString *string `xml:"faultstring" json:"faultstring"`
	Detail      *Detail `xml:"detail" json:"detail"`
}

type Detail struct {
	Message     *string      `xml:"message" json:"message"`
	ErrorType   *string      `xml:"errortype" json:"errortype"`
	XtracDetail *XtracDetail `xml:"xtracDetail" json:"xtracDetail"`
}

type XtracDetail struct {
	Code *string `xml:"code" json:"code"`
}

type SoapHeader struct {
	B2bCookie *string `xml:"urn:schemas-xtrac-fmr-com:b2b Cookie" json:"Cookie"`
	WsCookie  *string `xml:"http://xmlns.fmr.com/systems/dev/xtrac/2004/06/ Cookie" json:"Cookie"`
}

// always use the same name as type
type SoapBody struct {
	SoapFault      *SoapFault      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault" json:"Fault"`
	CreateWorkItem *CreateWorkItem `xml:"CreateWorkItem" json:"CreateWorkItem"`
}

func isNotEmpty(pl interface{}) bool {

	return !reflect.DeepEqual(pl, reflect.New(reflect.ValueOf(pl).Type()).Elem().Interface())
}

// SoapBody only filled with one element
// GetPayloadType finds the the element that is populated using reflection
func (s SoapBody) GetPayload() (interface{}, error) {

	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		if isNotEmpty(v.Field(i).Interface()) {
			return v.Field(i).Interface(), nil
			break
		}
	}
	return "", errors.New("No payload type found. SOAP body is empty ")
}
