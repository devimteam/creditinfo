package unit

import (
	"testing"

	"github.com/devimteam/creditinfo/pkg/creditInfoClient"
	"github.com/l-vitaly/gounit"
)

func TestGetIndividualReportSuccess(t *testing.T) {

	unit := gounit.New(t)

	params := creditInfoClient.CreditInfoParams{
		Username:    "username",
		Password:    "password",
		ConnectorId: "00000000-0000-0000-0000-000000000000",
		StrategyId:  "00000000-0000-0000-0000-000000000000",
		Endpoint:    "https://endpoint.example.com",
	}

	creditInfo := creditInfoClient.NewCreditInfoClient(params)

	response, err := creditInfo.GetIndividualReport("0")

	unit.AssertNil(err, "err not nil")

	unit.AssertEquals(response.Status, "ok", "status not equal 'ok'")

}

func TestGetIndividualReportFail(t *testing.T) {

	unit := gounit.New(t)

	params := creditInfoClient.CreditInfoParams{
		Username:    "username",
		Password:    "password",
		ConnectorId: "00000000-0000-0000-0000-000000000000",
		StrategyId:  "00000000-0000-0000-0000-000000000000", //FAIL DATA StrategyId
		Endpoint:    "https://endpoint.example.com",
	}

	creditInfo := creditInfoClient.NewCreditInfoClient(params)

	response, err := creditInfo.GetIndividualReport("0")

	unit.AssertNotNil(err, "err nil")
	unit.AssertNil(response, "err nil")

}
