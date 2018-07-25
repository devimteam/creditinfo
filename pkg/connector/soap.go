package connector

import (
	"encoding/xml"
	"github.com/fiorix/wsdl2go/soap"
	"time"
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
	BeginQuery(parameters *BeginQuery) (*BeginQueryResponse, error)

	// EndQuery was auto-generated from WSDL.
	EndQuery(parameters *EndQuery) (*EndQueryResponse, error)

	// ListConnectors was auto-generated from WSDL.
	ListConnectors(parameters *ListConnectors) (*ListConnectorsResponse, error)

	// Query was auto-generated from WSDL.
	Query(parameters *Query) (*QueryResponse, error)

	// Schema was auto-generated from WSDL.
	Schema(parameters *Schema) (*SchemaResponse, error)
}

// DateTime in WSDL format.
type DateTime string

// Char was auto-generated from WSDL.
type Char int

// Duration was auto-generated from WSDL.
type Duration time.Duration

// Guid was auto-generated from WSDL.
type Guid string

// ArrayOfConnectorInfo was auto-generated from WSDL.
type ArrayOfConnectorInfo struct {
	ConnectorInfo []*ConnectorInfo `xml:"ConnectorInfo,omitempty" json:"ConnectorInfo,omitempty" yaml:"ConnectorInfo,omitempty"`
}

// ArrayOfConnectorSchemas was auto-generated from WSDL.
type ArrayOfConnectorSchemas struct {
	ConnectorSchemas []*ConnectorSchemas `xml:"ConnectorSchemas,omitempty" json:"ConnectorSchemas,omitempty" yaml:"ConnectorSchemas,omitempty"`
}

// ArrayOfguid was auto-generated from WSDL.
type ArrayOfguid struct {
	Guid []*Guid `xml:"guid,omitempty" json:"guid,omitempty" yaml:"guid,omitempty"`
}

// BeginQuery was auto-generated from WSDL.
type BeginQuery struct {
	Request *MultiConnectorRequest `xml:"request,omitempty" json:"request,omitempty" yaml:"request,omitempty"`
}

// BeginQueryResponse was auto-generated from WSDL.
type BeginQueryResponse struct {
	BeginQueryResult *MultiConnectorTicket `xml:"BeginQueryResult,omitempty" json:"BeginQueryResult,omitempty" yaml:"BeginQueryResult,omitempty"`
}

// ConnectorInfo was auto-generated from WSDL.
type ConnectorInfo struct {
	Description *string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
	Id          *Guid   `xml:"Id,omitempty" json:"Id,omitempty" yaml:"Id,omitempty"`
	QueryStep   *Guid   `xml:"QueryStep,omitempty" json:"QueryStep,omitempty" yaml:"QueryStep,omitempty"`
}

// ConnectorSchemaRequest was auto-generated from WSDL.
type ConnectorSchemaRequest struct {
	ConnectorIds *ArrayOfguid `xml:"ConnectorIds,omitempty" json:"ConnectorIds,omitempty" yaml:"ConnectorIds,omitempty"`
	MessageId    *Guid        `xml:"MessageId,omitempty" json:"MessageId,omitempty" yaml:"MessageId,omitempty"`
}

// ConnectorSchemaResponse was auto-generated from WSDL.
type ConnectorSchemaResponse struct {
	MessageId *Guid                    `xml:"MessageId,omitempty" json:"MessageId,omitempty" yaml:"MessageId,omitempty"`
	Schemas   *ArrayOfConnectorSchemas `xml:"Schemas,omitempty" json:"Schemas,omitempty" yaml:"Schemas,omitempty"`
}

// ConnectorSchemas was auto-generated from WSDL.
type ConnectorSchemas struct {
	Id           *Guid        `xml:"Id,omitempty" json:"Id,omitempty" yaml:"Id,omitempty"`
	InputSchema  *interface{} `xml:"InputSchema,omitempty" json:"InputSchema,omitempty" yaml:"InputSchema,omitempty"`
	OutputSchema *interface{} `xml:"OutputSchema,omitempty" json:"OutputSchema,omitempty" yaml:"OutputSchema,omitempty"`
}

// ConnectorsListRequest was auto-generated from WSDL.
type ConnectorsListRequest struct {
	MessageId *Guid `xml:"MessageId,omitempty" json:"MessageId,omitempty" yaml:"MessageId,omitempty"`
}

// ConnectorsListResponse was auto-generated from WSDL.
type ConnectorsListResponse struct {
	Connectors *ArrayOfConnectorInfo `xml:"Connectors,omitempty" json:"Connectors,omitempty" yaml:"Connectors,omitempty"`
	MessageId  *Guid                 `xml:"MessageId,omitempty" json:"MessageId,omitempty" yaml:"MessageId,omitempty"`
}

// EndQuery was auto-generated from WSDL.
type EndQuery struct {
	Ticket *MultiConnectorTicket `xml:"ticket,omitempty" json:"ticket,omitempty" yaml:"ticket,omitempty"`
}

// EndQueryResponse was auto-generated from WSDL.
type EndQueryResponse struct {
	EndQueryResult *MultiConnectorResponse `xml:"EndQueryResult,omitempty" json:"EndQueryResult,omitempty" yaml:"EndQueryResult,omitempty"`
}

// Error was auto-generated from WSDL.
type Error struct {
	Message *string `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`
}

// FailedAuthentication was auto-generated from WSDL.
type FailedAuthentication struct {
}

// InProgress was auto-generated from WSDL.
type InProgress struct {
}

// ListConnectors was auto-generated from WSDL.
type ListConnectors struct {
	Request *ConnectorsListRequest `xml:"request,omitempty" json:"request,omitempty" yaml:"request,omitempty"`
}

// ListConnectorsResponse was auto-generated from WSDL.
type ListConnectorsResponse struct {
	ListConnectorsResult *ConnectorsListResponse `xml:"ListConnectorsResult,omitempty" json:"ListConnectorsResult,omitempty" yaml:"ListConnectorsResult,omitempty"`
}

// MultiConnectorRequest was auto-generated from WSDL.
type MultiConnectorRequest struct {
	MessageId     *Guid       `xml:"MessageId,omitempty" json:"MessageId,omitempty" yaml:"MessageId,omitempty"`
	OperationCode *string     `xml:"OperationCode,omitempty" json:"OperationCode,omitempty" yaml:"OperationCode,omitempty"`
	RequestXml    *RequestXml `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector RequestXml,omitempty" json:"RequestXml,omitempty" yaml:"RequestXml,omitempty"`
	ScheduledTime *DateTime   `xml:"ScheduledTime,omitempty" json:"ScheduledTime,omitempty" yaml:"ScheduledTime,omitempty"`
	Timeout       *Duration   `xml:"Timeout,omitempty" json:"Timeout,omitempty" yaml:"Timeout,omitempty"`
}

// MultiConnectorResponse was auto-generated from WSDL.
type MultiConnectorResponse struct {
	MessageId     *Guid        `xml:"MessageId,omitempty" json:"MessageId,omitempty" yaml:"MessageId,omitempty"`
	OperationCode *string      `xml:"OperationCode,omitempty" json:"OperationCode,omitempty" yaml:"OperationCode,omitempty"`
	ResponseXml   *interface{} `xml:"ResponseXml,omitempty" json:"ResponseXml,omitempty" yaml:"ResponseXml,omitempty"`
	Timestamp     *DateTime    `xml:"Timestamp,omitempty" json:"Timestamp,omitempty" yaml:"Timestamp,omitempty"`
}

// MultiConnectorTicket was auto-generated from WSDL.
type MultiConnectorTicket struct {
	MessageId *Guid `xml:"MessageId,omitempty" json:"MessageId,omitempty" yaml:"MessageId,omitempty"`
}

// Query was auto-generated from WSDL.
type Query struct {
	Request *MultiConnectorRequest `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector request,omitempty" json:"request,omitempty" yaml:"request,omitempty"`
}

// QueryResponse was auto-generated from WSDL.
type QueryResponse struct {
	QueryResult *MultiConnectorResponse `xml:"QueryResult,omitempty" json:"QueryResult,omitempty" yaml:"QueryResult,omitempty"`
}

type Strategy struct {
	// Strategy ID
	Id string `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Request Id,omitempty"`
}

type Cb5SearchParameters struct {
	NationalId *string `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Request NationalId,omitempty" json:"NationalId,omitempty" yaml:"NationalId,omitempty"`
}

type CustomFields struct {
}

// Request was auto-generated from WSDL.
type Request struct {
	XMLName             xml.Name            `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Request request"`
	Strategy            *Strategy           `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Request Strategy,omitempty"`
	Cb5SearchParameters Cb5SearchParameters `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Request Cb5SearchParameters,omitempty"`
	CustomFields        CustomFields        `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Request CustomFields,omitempty"`
	// Subject has given consent to run this query
	Consent bool `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Request Consent,omitempty"`
}

type ConnectorData struct {
	// Unique ID of request to particular connector.
	Id      *Guid    `xml:"id,attr,omitempty"`
	Request *Request `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Request request,omitempty"`
}

type Connector struct {
	Data *ConnectorData `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Messages/Request data,omitempty"`
	// Connector ID
	Id *Guid `xml:"id,attr,omitempty"`
	// If set to true it will try to retrieve data from cache.
	UseCache bool `xml:"useCache,attr,omitempty"`
	// Max age of cached data.
	MaxAge *Duration `xml:"maxAge,attr,omitempty"`
	// How many times should connector try to get data if the response if faulted.
	Retry byte `xml:"retry,attr,omitempty"`
	// If set to true connector nor cache won't be requested.
	Ignore bool `xml:"ignore,attr,omitempty"`
	// If set to true response contains more information about processing of a query
	IncludeMetadata bool `xml:"includeMetadata,attr,omitempty"`
}

type RequestXmlItem struct {
	Connector *Connector `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Messages/Request connector,omitempty"`
}

// RequestXml was auto-generated from WSDL.
type RequestXml struct {
	RequestXmlItem *RequestXmlItem `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector RequestXml,omitempty" json:"connector,omitempty" yaml:"connector,omitempty"`
}

// Schema was auto-generated from WSDL.
type Schema struct {
	Request *ConnectorSchemaRequest `xml:"request,omitempty" json:"request,omitempty" yaml:"request,omitempty"`
}

// SchemaResponse was auto-generated from WSDL.
type SchemaResponse struct {
	SchemaResult *ConnectorSchemaResponse `xml:"SchemaResult,omitempty" json:"SchemaResult,omitempty" yaml:"SchemaResult,omitempty"`
}

// Timeout was auto-generated from WSDL.
type Timeout struct {
	Value *string `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

// UnsupportedSecurityToken was auto-generated from WSDL.
type UnsupportedSecurityToken struct {
}

// Response was auto-generated from WSDL.
type Response []interface{}

// Operation wrapper for BeginQuery.
// OperationMultiConnectorService_BeginQuery_InputMessage was auto-generated
// from WSDL.
type OperationMultiConnectorService_BeginQuery_InputMessage struct {
	Parameters *BeginQuery `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for BeginQuery.
// OperationMultiConnectorService_BeginQuery_OutputMessage was
// auto-generated from WSDL.
type OperationMultiConnectorService_BeginQuery_OutputMessage struct {
	Parameters *BeginQueryResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for EndQuery.
// OperationMultiConnectorService_EndQuery_InputMessage was auto-generated
// from WSDL.
type OperationMultiConnectorService_EndQuery_InputMessage struct {
	Parameters *EndQuery `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for EndQuery.
// OperationMultiConnectorService_EndQuery_OutputMessage was auto-generated
// from WSDL.
type OperationMultiConnectorService_EndQuery_OutputMessage struct {
	Parameters *EndQueryResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for ListConnectors.
// OperationMultiConnectorService_ListConnectors_InputMessage was
// auto-generated from WSDL.
type OperationMultiConnectorService_ListConnectors_InputMessage struct {
	Parameters *ListConnectors `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for ListConnectors.
// OperationMultiConnectorService_ListConnectors_OutputMessage
// was auto-generated from WSDL.
type OperationMultiConnectorService_ListConnectors_OutputMessage struct {
	Parameters *ListConnectorsResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for Query.
// OperationMultiConnectorService_Query_InputMessage was auto-generated
// from WSDL.
type OperationMultiConnectorService_Query_InputMessage struct {
	Parameters *Query `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector Query,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for Query.
// OperationMultiConnectorService_Query_OutputMessage was auto-generated
// from WSDL.
type OperationMultiConnectorService_Query_OutputMessage struct {
	Parameters *QueryResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for Schema.
// OperationMultiConnectorService_Schema_InputMessage was auto-generated
// from WSDL.
type OperationMultiConnectorService_Schema_InputMessage struct {
	Parameters *Schema `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for Schema.
// OperationMultiConnectorService_Schema_OutputMessage was auto-generated
// from WSDL.
type OperationMultiConnectorService_Schema_OutputMessage struct {
	Parameters *SchemaResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// multiConnectorService implements the MultiConnectorService interface.
type multiConnectorService struct {
	cli *soap.Client
}

// BeginQuery was auto-generated from WSDL.
func (p *multiConnectorService) BeginQuery(parameters *BeginQuery) (*BeginQueryResponse, error) {
	α := struct {
		OperationMultiConnectorService_BeginQuery_InputMessage `xml:"tns:BeginQuery"`
	}{
		OperationMultiConnectorService_BeginQuery_InputMessage{
			parameters,
		},
	}

	γ := struct {
		OperationMultiConnectorService_BeginQuery_OutputMessage `xml:"BeginQueryResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://creditinfo.com/schemas/2012/09/MultiConnector/MultiConnectorService/BeginQuery", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// EndQuery was auto-generated from WSDL.
func (p *multiConnectorService) EndQuery(parameters *EndQuery) (*EndQueryResponse, error) {
	α := struct {
		OperationMultiConnectorService_EndQuery_InputMessage `xml:"tns:EndQuery"`
	}{
		OperationMultiConnectorService_EndQuery_InputMessage{
			parameters,
		},
	}

	γ := struct {
		OperationMultiConnectorService_EndQuery_OutputMessage `xml:"EndQueryResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://creditinfo.com/schemas/2012/09/MultiConnector/MultiConnectorService/EndQuery", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// ListConnectors was auto-generated from WSDL.
func (p *multiConnectorService) ListConnectors(parameters *ListConnectors) (*ListConnectorsResponse, error) {
	α := struct {
		OperationMultiConnectorService_ListConnectors_InputMessage `xml:"tns:ListConnectors"`
	}{
		OperationMultiConnectorService_ListConnectors_InputMessage{
			parameters,
		},
	}

	γ := struct {
		OperationMultiConnectorService_ListConnectors_OutputMessage `xml:"ListConnectorsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://creditinfo.com/schemas/2012/09/MultiConnector/MultiConnectorService/ListConnectors", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// Query was auto-generated from WSDL.
func (p *multiConnectorService) Query(parameters *Query) (*QueryResponse, error) {
	α := struct {
		OperationMultiConnectorService_Query_InputMessage `xml:"tns:Query"`
	}{
		OperationMultiConnectorService_Query_InputMessage{
			parameters,
		},
	}

	γ := struct {
		OperationMultiConnectorService_Query_OutputMessage `xml:"QueryResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://creditinfo.com/schemas/2012/09/MultiConnector/MultiConnectorService/Query", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// Schema was auto-generated from WSDL.
func (p *multiConnectorService) Schema(parameters *Schema) (*SchemaResponse, error) {
	α := struct {
		OperationMultiConnectorService_Schema_InputMessage `xml:"tns:Schema"`
	}{
		OperationMultiConnectorService_Schema_InputMessage{
			parameters,
		},
	}

	γ := struct {
		OperationMultiConnectorService_Schema_OutputMessage `xml:"SchemaResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://creditinfo.com/schemas/2012/09/MultiConnector/MultiConnectorService/Schema", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}
