package client

import (
	"context"
	"net/http"
	"time"

	wsse "github.com/casualcode/soap"
	"github.com/devimteam/creditinfo/pkg/connector"
	"github.com/devimteam/creditinfo/pkg/request"
	"github.com/devimteam/creditinfo/pkg/response"
	"github.com/devimteam/wsdl2go/soap"
	"github.com/gofrs/uuid"
)

//Soap Client provides an interface for getting data out creditinfo service
type Client interface {
	GetIndividualReport(ctx context.Context, nationalId *string, phone *string, birthDate *time.Time) (*response.ResultResponse, error)
	GetIndividualReportBeginQuery(ctx context.Context, nationalId *string, phone *string, birthDate *time.Time) (ticket *connector.MultiConnectorTicket, err error)
	EndQuery(ctx context.Context, ticket *connector.MultiConnectorTicket) (*response.ResultResponse, error)
}

type CreditInfoParams struct {
	Username    string
	Password    string
	ConnectorId string
	StrategyId  string
	Endpoint    string
	Timeout     *time.Duration
}

type creditInfo struct {
	params CreditInfoParams
	svc    connector.MultiConnectorService
}

// NewClient returns a Client interface for given soap api creditinfo
func NewCreditInfoClient(params CreditInfoParams, pre func(context.Context, *http.Request), post func(context.Context, *http.Response)) *creditInfo {

	svc := connector.NewMultiConnectorService(&soap.Client{
		URL:                    params.Endpoint,
		Header:                 getWsseHeader(params.Username, params.Password),
		Namespace:              connector.Namespace,
		ExcludeActionNamespace: true,
		Pre:                    pre,
		Post:                   post,
	})

	return &creditInfo{
		params: params,
		svc:    svc,
	}
}

func (client *creditInfo) GetIndividualReport(ctx context.Context, nationalId *string, phone *string, birthDate *time.Time) (response *response.ResultResponse, err error) {
	messageId, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	dataId, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	var bdate *request.Date
	if birthDate != nil {
		t := request.Date(*birthDate)
		bdate = &t
	}

	return client.svc.Query(ctx, &connector.Query{
		Request: client.getMultiConnectorRequest(
			messageId.String(),
			dataId.String(),
			*nationalId,
			&request.CustomFields{
				MobilePhone: phone,
				DateOfBirth: bdate,
			},
		),
	})
}

func (client *creditInfo) getMultiConnectorRequest(messageId, dataId, nationalId string, customFields *request.CustomFields) *request.MultiConnectorRequest {
	return &request.MultiConnectorRequest{
		MessageId: &messageId,
		Timeout:   request.NilDuration(client.params.Timeout),
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
func (client *creditInfo) GetIndividualReportBeginQuery(ctx context.Context, nationalId *string, phone *string, birthDate *time.Time) (ticket *connector.MultiConnectorTicket, err error) {
	messageId, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	dataId, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	var bdate *request.Date
	if birthDate != nil {
		t := request.Date(*birthDate)
		bdate = &t
	}

	return client.svc.BeginQuery(ctx, &connector.BeginQuery{
		Request: client.getMultiConnectorRequest(messageId.String(), dataId.String(), *nationalId, &request.CustomFields{
			MobilePhone: phone,
			DateOfBirth: bdate,
		}),
	})
}

func (client *creditInfo) EndQuery(ctx context.Context, ticket *connector.MultiConnectorTicket) (*response.ResultResponse, error) {
	return client.svc.EndQuery(ctx, &connector.EndQuery{
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
