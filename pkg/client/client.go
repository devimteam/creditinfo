package client

import "github.com/satori/go.uuid"

import (
	"fmt"

	"github.com/devimteam/creditinfo/pkg/connector"
	"github.com/fiorix/wsdl2go/soap"
	wsse "github.com/casualcode/soap"
)

const (
	TestEndpoint string = "https://idmtest.creditinfo.co.ke/MultiConnector.svc"
	LiveEndpoint string = "https://idmtest.creditinfo.co.ke"
)

//Soap Client provides an interface for getting data out creditinfo service
type Client interface {
	GetIndividualReport(nationalId string) (*QueryResponse, error)
}

type creditInfo struct {
	username    string
	password    string
	connectorId string
	strategyId  string
	isLive      bool
}

// NewClient returns a Client interface for given consul address
func NewCreditInfoService(username string,
	password string,
	connectorId string,
	strategyId string,
	isLive bool,
) Client {
	return &creditInfo{
		username:    username,
		password:    password,
		connectorId: connectorId,
		strategyId:  strategyId,
		isLive:      isLive,
	}
}

func (client creditInfo) GetIndividualReport(nationalId string) (*QueryResponse, error) {
	cli := client.getSoapClient()

	messageId := generateUUID()
	dataId := generateUUID()
	connectorGuuid := connector.Guid(client.connectorId)

	multiConnectorService := connector.NewMultiConnectorService(cli)
	response, err := multiConnectorService.Query(&connector.Query{
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
									Id: client.strategyId,
								},
								Cb5SearchParameters: connector.Cb5SearchParameters{
									NationalId: &nationalId,
								},
								CustomFields: connector.CustomFields{},
								Consent:      true,
							},
						},
					},
				},
			},
		},
	})

	fmt.Println(*response.QueryResult.ResponseXml.Response.Connector.Data.Response.Status)

	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	return &QueryResponse{}, nil
}

func (client creditInfo) getSoapClient() *soap.Client {
	return &soap.Client{
		URL:                    client.getEndpoint(),
		Header:                 client.getWsseHeader(),
		Namespace:              connector.Namespace,
		UserAgent:              "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36",
		ExcludeActionNamespace: true,
	}
}

func (client creditInfo) getEndpoint() string {
	if client.isLive {
		return LiveEndpoint
	}

	return TestEndpoint
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

	env.Header.WsseSecurity.UsernameToken.Username.Value = client.username
	env.Header.WsseSecurity.UsernameToken.Password.Value = client.password

	return env.Header
}

func generateUUID() connector.Guid {
	return connector.Guid(uuid.NewV4().String())
}
