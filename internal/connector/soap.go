package connector

import (
	"fmt"
	"time"

	"github.com/devimteam/creditinfo/pkg/request"
	"github.com/devimteam/creditinfo/pkg/response"
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
	// BeginQuery was auto-generated from WSDL.
	BeginQuery(parameters *BeginQuery) (*MultiConnectorTicket, error)
	// EndQuery was auto-generated from WSDL.
	EndQuery(parameters *EndQuery) (*response.ResultResponse, error)
	// Query was auto-generated from WSDL.
	Query(parameters *Query) (*response.ResultResponse, error)
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

// BeginQuery was auto-generated from WSDL.
type BeginQuery struct {
	Request *request.MultiConnectorRequest `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector request,omitempty" json:"request,omitempty" yaml:"request,omitempty"`
}

// BeginQueryResponse was auto-generated from WSDL.
type BeginQueryResponse struct {
	BeginQueryResult *MultiConnectorTicket `xml:"BeginQueryResult,omitempty" json:"BeginQueryResult,omitempty" yaml:"BeginQueryResult,omitempty"`
}

// EndQuery was auto-generated from WSDL.
type EndQuery struct {
	Ticket *MultiConnectorTicket `xml:"ticket,omitempty" json:"ticket,omitempty" yaml:"ticket,omitempty"`
}

// EndQueryResponse was auto-generated from WSDL.
type EndQueryResponse struct {
	EndQueryResult *response.MultiConnectorResponse `xml:"EndQueryResult,omitempty" json:"EndQueryResult,omitempty" yaml:"EndQueryResult,omitempty"`
}

// MultiConnectorTicket was auto-generated from WSDL.
type MultiConnectorTicket struct {
	MessageId *Guid `xml:"MessageId,omitempty" json:"MessageId,omitempty" yaml:"MessageId,omitempty"`
}

// Operation wrapper for BeginQuery.
// OperationMultiConnectorService_BeginQuery_InputMessage was auto-generated
// from WSDL.
type OperationMultiConnectorService_BeginQuery_InputMessage struct {
	Parameters *BeginQuery `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector BeginQuery,omitempty" json:"BeginQuery,omitempty" yaml:"BeginQuery,omitempty"`
}

// Operation wrapper for BeginQuery.
// OperationMultiConnectorService_BeginQuery_OutputMessage was
// auto-generated from WSDL.
type OperationMultiConnectorService_BeginQuery_OutputMessage struct {
	Parameters *BeginQueryResponse `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector BeginQueryResponse,omitempty" json:"BeginQueryResponse,omitempty" yaml:"BeginQueryResponse,omitempty"`
}

// Operation wrapper for EndQuery.
// OperationMultiConnectorService_EndQuery_InputMessage was auto-generated
// from WSDL.
type OperationMultiConnectorService_EndQuery_InputMessage struct {
	Parameters *EndQuery `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector EndQuery,omitempty" json:"EndQuery,omitempty" yaml:"EndQuery,omitempty"`
}

// Operation wrapper for EndQuery.
// OperationMultiConnectorService_EndQuery_OutputMessage was auto-generated
// from WSDL.
type OperationMultiConnectorService_EndQuery_OutputMessage struct {
	Parameters *EndQueryResponse `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector EndQueryResponse,omitempty" json:"EndQueryResponse,omitempty" yaml:"EndQueryResponse,omitempty"`
}

// Query was auto-generated from WSDL.
type Query struct {
	Request *request.MultiConnectorRequest `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector request,omitempty" json:"request,omitempty" yaml:"request,omitempty"`
}

// QueryResponse was auto-generated from WSDL.
type QueryResponse struct {
	QueryResult *response.MultiConnectorResponse `xml:"QueryResult,omitempty" json:"QueryResult,omitempty" yaml:"QueryResult,omitempty"`
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
func (p *multiConnectorService) Query(parameters *Query) (*response.ResultResponse, error) {
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
		return nil, fmt.Errorf("status: %s; message: %s", status, out.Parameters.QueryResult.ResponseXml.Response.Connector.Data.Response.Infomsg)
	}

	return out.Parameters.QueryResult.ResponseXml.Response.Connector.Data.Response, nil
}

// BeginQuery was auto-generated from WSDL.
func (p *multiConnectorService) BeginQuery(parameters *BeginQuery) (*MultiConnectorTicket, error) {
	input := struct {
		OperationMultiConnectorService_BeginQuery_InputMessage `xml:"tns:BeginQuery"`
	}{
		OperationMultiConnectorService_BeginQuery_InputMessage{
			parameters,
		},
	}

	out := struct {
		OperationMultiConnectorService_BeginQuery_OutputMessage `xml:"tns:BeginQueryResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://creditinfo.com/schemas/2012/09/MultiConnector/MultiConnectorService/BeginQuery", input, &out); err != nil {
		return nil, err
	}
	return out.Parameters.BeginQueryResult, nil
}

// EndQuery was auto-generated from WSDL.
func (p *multiConnectorService) EndQuery(parameters *EndQuery) (*response.ResultResponse, error) {
	input := struct {
		OperationMultiConnectorService_EndQuery_InputMessage `xml:"tns:EndQuery"`
	}{
		OperationMultiConnectorService_EndQuery_InputMessage{
			parameters,
		},
	}

	out := struct {
		OperationMultiConnectorService_EndQuery_OutputMessage `xml:"tns:EndQueryResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://creditinfo.com/schemas/2012/09/MultiConnector/MultiConnectorService/EndQuery", input, &out); err != nil {
		return nil, err
	}

	if err := out.Parameters.EndQueryResult.ResponseXml.Response.Connector.Error; err != nil {
		return nil, errors.New(err.Message)
	}

	if status := out.Parameters.EndQueryResult.ResponseXml.Response.Connector.Data.Response.Status; status != "ok" {
		return nil, errors.New(status)
	}

	return out.Parameters.EndQueryResult.ResponseXml.Response.Connector.Data.Response, nil
}
