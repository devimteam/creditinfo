package unit

import (
	"testing"

	"github.com/devimteam/creditinfo/pkg/client"
	"github.com/l-vitaly/gounit"
)

func TestGetIndividualReportSuccess(t *testing.T) {

	unit := gounit.New(t)

	params := client.CreditInfoParams{
		Username:    "username",
		Password:    "password",
		ConnectorId: "00000000-0000-0000-0000-000000000000",
		StrategyId:  "00000000-0000-0000-0000-000000000000",
		Endpoint:    "https://endpoint.example.com",
	}

	creditInfo := client.NewCreditInfoClient(params)
	nationalId := "12345678"
	phone := "25411111111"
	response, err := creditInfo.GetIndividualReport(&nationalId, &phone, nil)

	t.Log(response.GeneralInformation.RecommendedDecision)

	unit.AssertNil(err, "err not nil")
	unit.AssertEquals(response.Status, "ok", "status not equal 'ok'")
}

func TestGetIndividualReportFail(t *testing.T) {

	unit := gounit.New(t)

	params := client.CreditInfoParams{
		Username:    "username",
		Password:    "password",
		ConnectorId: "00000000-0000-0000-0000-000000000000",
		StrategyId:  "00000000-0000-0000-0000-000000000000",
		Endpoint:    "https://endpoint.example.com",
	}

	creditInfo := client.NewCreditInfoClient(params)
	nationalId := "0"
	phone := "25411111111"
	response, err := creditInfo.GetIndividualReport(&nationalId, &phone, nil)

	unit.AssertNotNil(err, "err nil")
	unit.AssertNil(response, "err nil")

}
