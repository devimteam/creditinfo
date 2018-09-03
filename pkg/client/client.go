package client

import (
	"time"

	wsse "github.com/casualcode/soap"
	"github.com/devimteam/creditinfo/internal/connector"
	"github.com/devimteam/creditinfo/pkg/request"
	"github.com/devimteam/creditinfo/pkg/response"
	"github.com/fiorix/wsdl2go/soap"
	"github.com/satori/go.uuid"
)

//Soap Client provides an interface for getting data out creditinfo service
type Client interface {
	GetIndividualReport(nationalId *string, phone *string, birthDate *time.Time) (*response.ResultResponse, error)
	GetIndividualReportBeginQuery(nationalId *string, phone *string, birthDate *time.Time) (ticket *connector.MultiConnectorTicket, err error)
	EndQuery(ticket *connector.MultiConnectorTicket) (*response.ResultResponse, error)
}
type CreditInfoParams struct {
	Username    string
	Password    string
	ConnectorId string
	StrategyId  string
	Endpoint    string
}

type creditInfo struct {
	params CreditInfoParams
	svc    connector.MultiConnectorService
}

// NewClient returns a Client interface for given soap api creditinfo
func NewCreditInfoClient(params CreditInfoParams) Client {

	svc := connector.NewMultiConnectorService(&soap.Client{
		URL:                    params.Endpoint,
		Header:                 getWsseHeader(params.Username, params.Password),
		Namespace:              connector.Namespace,
		ExcludeActionNamespace: true,
	})

	return &creditInfo{
		params: params,
		svc:    svc,
	}
}

func (client *creditInfo) GetIndividualReport(nationalId *string, phone *string, birthDate *time.Time) (response *response.ResultResponse, err error) {
	messageId := uuid.NewV4().String()
	dataId := uuid.NewV4().String()

	var birthDdateFormat string
	if birthDate != nil {
		birthDdateFormat = birthDate.Format("2006-01-02")
	}

	return client.svc.Query(&connector.Query{
		Request: client.getMultiConnectorRequest(messageId, dataId, *nationalId, &request.CustomFields{
			MobilePhone: phone,
			DateOfBirth: &birthDdateFormat,
		}),
	})
}

func (client *creditInfo) getMultiConnectorRequest(messageId, dataId, nationalId string, customFields *request.CustomFields) *request.MultiConnectorRequest {
	return &request.MultiConnectorRequest{
		MessageId: &messageId,
		RequestXml: &request.RequestXml{
			RequestXmlItem: &request.RequestXmlItem{
				Connector: &request.ConnectorRequest{
					Id: &client.params.ConnectorId,
					Data: &request.ConnectorDataRequest{
						Id: &dataId,
						Request: &request.Request{
							Strategy: &request.Strategy{
								Id: client.params.StrategyId,
							},
							Cb5SearchParameters: request.Cb5SearchParameters{
								NationalId: &nationalId,
							},
							CustomFields: customFields,
							Consent:      true,
						},
					},
				},
			},
		},
	}
}
func (client *creditInfo) GetIndividualReportBeginQuery(nationalId *string, phone *string, birthDate *time.Time) (ticket *connector.MultiConnectorTicket, err error) {
	messageId := uuid.NewV4().String()
	dataId := uuid.NewV4().String()

	var birthDdateFormat string
	if birthDate != nil {
		birthDdateFormat = birthDate.Format("2006-01-02")
	}

	return client.svc.BeginQuery(&connector.BeginQuery{
		Request: client.getMultiConnectorRequest(messageId, dataId, *nationalId, &request.CustomFields{
			MobilePhone: phone,
			DateOfBirth: &birthDdateFormat,
		}),
	})
}

func (client *creditInfo) EndQuery(ticket *connector.MultiConnectorTicket) (*response.ResultResponse, error) {
	return client.svc.EndQuery(&connector.EndQuery{
		Ticket: ticket,
	})
}

func getWsseHeader(username string, password string) *wsse.Header {
	env := &wsse.Envelope{
		XmlnsSoapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsUniv:    "http://www.example.pl/ws/test/universal",
		Header: &wsse.Header{
			WsseSecurity: &wsse.WsseSecurity{
				MustUnderstand: "1",
				XmlnsWsse:      "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
				XmlnsWsu:       "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
				UsernameToken: &wsse.UsernameToken{
					WsuId:    "UsernameToken-1",
					Username: &wsse.Username{},
					Password: &wsse.Password{
						Type: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText",
					},
				},
			},
		},
	}

	env.Header.WsseSecurity.UsernameToken.Username.Value = username
	env.Header.WsseSecurity.UsernameToken.Password.Value = password

	return env.Header
}
