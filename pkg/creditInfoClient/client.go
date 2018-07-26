package creditInfoClient

import (
	wsse "github.com/casualcode/soap"
	"github.com/devimteam/creditinfo/internal/connector"
	"github.com/fiorix/wsdl2go/soap"
	"github.com/satori/go.uuid"
)

//Soap Client provides an interface for getting data out creditinfo service
type Client interface {
	GetIndividualReport(nationalId string) (*connector.ResultResponse, error)
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
}

// NewClient returns a Client interface for given soap api creditinfo
func NewCreditInfoClient(params CreditInfoParams) Client {
	return &creditInfo{
		params: params,
	}
}

func (client creditInfo) GetIndividualReport(nationalId string) (*connector.ResultResponse, error) {
	messageId := connector.Guid(uuid.NewV4().String())
	dataId := connector.Guid(uuid.NewV4().String())
	connectorGuuid := connector.Guid(client.params.ConnectorId)

	return connector.NewMultiConnectorService(client.getSoapClient()).Query(&connector.Query{
		Request: &connector.MultiConnectorRequest{
			MessageId: &messageId,
			RequestXml: &connector.RequestXml{
				RequestXmlItem: &connector.RequestXmlItem{
					Connector: &connector.ConnectorRequest{
						Id: &connectorGuuid,
						Data: &connector.ConnectorDataRequest{
							Id: &dataId,
							Request: &connector.Request{
								Strategy: &connector.Strategy{
									Id: client.params.StrategyId,
								},
								Cb5SearchParameters: connector.Cb5SearchParameters{
									NationalId: &nationalId,
								},
								CustomFields: &connector.CustomFields{},
								Consent:      true,
							},
						},
					},
				},
			},
		},
	})
}

func (client creditInfo) getSoapClient() *soap.Client {
	return &soap.Client{
		URL:                    client.params.Endpoint,
		Header:                 client.getWsseHeader(),
		Namespace:              connector.Namespace,
		ExcludeActionNamespace: true,
	}
}

func (client creditInfo) getWsseHeader() *wsse.Header {
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

	env.Header.WsseSecurity.UsernameToken.Username.Value = client.params.Username
	env.Header.WsseSecurity.UsernameToken.Password.Value = client.params.Password

	return env.Header
}
