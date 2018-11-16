package request

import (
	"bytes"
	"encoding/xml"
	"strconv"
	"time"
)

// DateTime in WSDL format.
type DateTime string
type Duration time.Duration

func (d Duration) String() string {
	buf := bytes.NewBufferString("PT")
	buf.WriteString(strconv.Itoa(int(time.Duration(d).Seconds())))
	buf.WriteString("S")
	return buf.String()
}

func NilDuration(d *time.Duration) *Duration {
	if d == nil {
		return nil
	}
	x := Duration(*d)
	return &x
}

// MultiConnectorRequest was auto-generated from WSDL.
type MultiConnectorRequest struct {
	MessageId     *string     `xml:"MessageId,omitempty" json:"MessageId,omitempty" yaml:"MessageId,omitempty"`
	OperationCode *string     `xml:"OperationCode,omitempty" json:"OperationCode,omitempty" yaml:"OperationCode,omitempty"`
	RequestXml    *RequestXml `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector RequestXml,omitempty" json:"RequestXml,omitempty" yaml:"RequestXml,omitempty"`
	ScheduledTime *DateTime   `xml:"ScheduledTime,omitempty" json:"ScheduledTime,omitempty" yaml:"ScheduledTime,omitempty"`
	Timeout       *Duration   `xml:"Timeout,omitempty" json:"Timeout,omitempty" yaml:"Timeout,omitempty"`
}

type RequestXmlItem struct {
	Connector *ConnectorRequest `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Messages/Request connector,omitempty"`
}

// RequestXml was auto-generated from WSDL.
type RequestXml struct {
	RequestXmlItem *RequestXmlItem `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector RequestXml,omitempty" json:"connector,omitempty" yaml:"connector,omitempty"`
}

type ConnectorRequest struct {
	Data *ConnectorDataRequest `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Messages/Request data,omitempty"`
	// ConnectorRequest ID
	Id *string `xml:"id,attr,omitempty"`
	// If set to true it will try to retrieve data from cache.
	UseCache bool `xml:"useCache,attr,omitempty"`
	// Max age of cached data.
	MaxAge *time.Duration `xml:"maxAge,attr,omitempty"`
	// How many times should connector try to get data if the response if faulted.
	Retry byte `xml:"retry,attr,omitempty"`
	// If set to true connector nor cache won't be requested.
	Ignore bool `xml:"ignore,attr,omitempty"`
	// If set to true response contains more information about processing of a query
	IncludeMetadata bool `xml:"includeMetadata,attr,omitempty"`
}

type Request struct {
	XMLName             xml.Name            `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Request request"`
	Strategy            *Strategy           `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Request Strategy,omitempty"`
	Cb5SearchParameters Cb5SearchParameters `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Request Cb5SearchParameters,omitempty"`
	CustomFields        *CustomFields       `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Request CustomFields,omitempty"`
	// Subject has given consent to run this query
	Consent bool `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Request Consent,omitempty"`
}

type ConnectorDataRequest struct {
	// Unique ID of request to particular connector.
	Id      *string  `xml:"id,attr,omitempty"`
	Request *Request `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Request request,omitempty"`
}

type Strategy struct {
	// Strategy ID
	Id string `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Request Id,omitempty"`
}

type Cb5SearchParameters struct {
	NationalId *string `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Request NationalId,omitempty" json:"NationalId,omitempty" yaml:"NationalId,omitempty"`
}

type CustomFields struct {
	MobilePhone *string `xml:"MobilePhone,omitempty" json:"Timeout,omitempty" yaml:"Timeout,omitempty"`
	DateOfBirth *Date   `xml:"DateOfBirth,omitempty" json:"DateOfBirth,omitempty" yaml:"DateOfBirth,omitempty"`
}

const xsdDateFormat = "2006-01-02"

type Date time.Time

func (d Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t := time.Time(d).Format(xsdDateFormat)
	return e.EncodeElement(t, start)
}

func (d *Date) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var temp string
	err := dec.DecodeElement(&temp, &start)
	if err != nil {
		return err
	}
	t, err := time.Parse(xsdDateFormat, temp)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}
