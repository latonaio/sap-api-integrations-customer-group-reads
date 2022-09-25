package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-customer-group-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToCustomerGroup(raw []byte, l *logger.Logger) ([]CustomerGroup, error) {
	pm := &responses.CustomerGroup{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CustomerGroup. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	customerGroup := make([]CustomerGroup, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		customerGroup = append(customerGroup, CustomerGroup{
			CustomerGroup: data.CustomerGroup,
			ToText:        data.ToText.Deferred.URI,
		})
	}

	return customerGroup, nil
}

func ConvertToText(raw []byte, l *logger.Logger) ([]Text, error) {
	pm := &responses.Text{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Text. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	text := make([]Text, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		text = append(text, Text{
			CustomerGroup:     data.CustomerGroup,
			Language:          data.Language,
			CustomerGroupName: data.CustomerGroupName,
		})
	}

	return text, nil
}
