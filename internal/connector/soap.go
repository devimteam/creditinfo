package connector

import (
	"time"

	"github.com/fiorix/wsdl2go/soap"
	"github.com/pkg/errors"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://creditinfo.com/schemas/2012/09/MultiConnector"

// NewMultiConnectorService creates an initializes a MultiConnectorService.
func NewMultiConnectorService(cli *soap.Client) MultiConnectorService {
	return &multiConnectorService{cli}
}

// MultiConnectorService was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type MultiConnectorService interface {
	// Query was auto-generated from WSDL.
	Query(parameters *Query) (*ResultResponse, error)
}

// Char was auto-generated from WSDL.
type Char int

// Duration was auto-generated from WSDL.
type Duration time.Duration

// Guid was auto-generated from WSDL.
type Guid string

// FailedAuthentication was auto-generated from WSDL.
type FailedAuthentication struct {
}

// InProgress was auto-generated from WSDL.
type InProgress struct {
}

// Query was auto-generated from WSDL.
type Query struct {
	Request *MultiConnectorRequest `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector request,omitempty" json:"request,omitempty" yaml:"request,omitempty"`
}

// QueryResponse was auto-generated from WSDL.
type QueryResponse struct {
	QueryResult *MultiConnectorResponse `xml:"QueryResult,omitempty" json:"QueryResult,omitempty" yaml:"QueryResult,omitempty"`
}

// Operation wrapper for Query.
// OperationMultiConnectorService_Query_InputMessage was auto-generated
// from WSDL.
type OperationMultiConnectorService_Query_InputMessage struct {
	Parameters *Query `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector Query,omitempty" json:"Query,omitempty" yaml:"Query,omitempty"`
}

// Operation wrapper for Query.
// OperationMultiConnectorService_Query_OutputMessage was auto-generated
// from WSDL.
type OperationMultiConnectorService_Query_OutputMessage struct {
	Parameters *QueryResponse `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector QueryResponse,omitempty" json:"QueryResponse,omitempty" yaml:"QueryResponse,omitempty"`
}

// multiConnectorService implements the MultiConnectorService interface.
type multiConnectorService struct {
	cli *soap.Client
}

// Query was auto-generated from WSDL.
func (p *multiConnectorService) Query(parameters *Query) (*ResultResponse, error) {
	intput := struct {
		OperationMultiConnectorService_Query_InputMessage `xml:"tns:Query"`
	}{
		OperationMultiConnectorService_Query_InputMessage{
			parameters,
		},
	}

	out := struct {
		OperationMultiConnectorService_Query_OutputMessage `xml:"tns:QueryResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://creditinfo.com/schemas/2012/09/MultiConnector/MultiConnectorService/Query", intput, &out); err != nil {
		return nil, err
	}
	if err := out.Parameters.QueryResult.ResponseXml.Response.Connector.Error; err != nil {
		return nil, errors.New(err.Message)
	}

	if status := out.Parameters.QueryResult.ResponseXml.Response.Connector.Data.Response.Status; status != "ok" {
		return nil, errors.New(status)
	}

	return out.Parameters.QueryResult.ResponseXml.Response.Connector.Data.Response, nil
}
