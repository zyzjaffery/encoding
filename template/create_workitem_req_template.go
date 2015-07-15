package template

import (
	"bytes"
	"github.com/yaozong/encoding/types"
	"text/template"
)

var createTemplate *template.Template

func init() {
	createTemplate = template.Must(template.New("createWorkitem").ParseFiles(
		"create_workitem_req_json.template", "create_workitem_req_json_min.template", "create_workitem_req_soap.template"))
}

func MarshalWithSoapTemplate(soapEnv *types.SoapEnvelope) ([]byte, error) {
	return marshalWithTemplate(soapEnv, "create_workitem_req_soap.template")
}

func MarshalWithJsonTemplate(soapEnv *types.SoapEnvelope) ([]byte, error) {
	return marshalWithTemplate(soapEnv, "create_workitem_req_json.template")
}

func MarshalWithMinimisedJsonTemplate(soapEnv *types.SoapEnvelope) ([]byte, error) {
	return marshalWithTemplate(soapEnv, "create_workitem_req_json_min.template")
}

func marshalWithTemplate(soapEnv *types.SoapEnvelope, templateName string) ([]byte, error) {
	bts := new(bytes.Buffer)
	err := createTemplate.ExecuteTemplate(bts, templateName, soapEnv)
	if err != nil {
		return nil, err
	} else {
		return bts.Bytes(), nil
	}
}
