package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-customer-group-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetCustomerGroup(customerGroup, language, customerGroupName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "CustomerGroup":
			func() {
				c.CustomerGroup(customerGroup)
				wg.Done()
			}()
		case "Text":
			func() {
				c.Text(language, customerGroupName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) CustomerGroup(customerGroup string) {
	customerGroupData, err := c.callCustomerGroupSrvAPIRequirementCustomerGroup("A_CustomerGroup", customerGroup)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(customerGroupData)

	textData, err := c.callText(customerGroupData[0].ToText)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(textData)
}

func (c *SAPAPICaller) callCustomerGroupSrvAPIRequirementCustomerGroup(api, customerGroup string) ([]sap_api_output_formatter.CustomerGroup, error) {
	url := strings.Join([]string{c.baseURL, "API_CUSTOMERGROUP_SRV", api}, "/")
	param := c.getQueryWithCustomerGroup(map[string]string{}, customerGroup)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCustomerGroup(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callText(url string) ([]sap_api_output_formatter.Text, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Text(language, customerGroupName string) {
	data, err := c.callSalesDistrictSrvAPIRequirementText("A_CustomerGroupText", language, customerGroupName)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callSalesDistrictSrvAPIRequirementText(api, language, customerGroupName string) ([]sap_api_output_formatter.Text, error) {
	url := strings.Join([]string{c.baseURL, "API_CUSTOMERGROUP_SRV", api}, "/")

	param := c.getQueryWithText(map[string]string{}, language, customerGroupName)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithCustomerGroup(params map[string]string, customerGroup string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("CustomerGroup eq '%s'", customerGroup)
	return params
}

func (c *SAPAPICaller) getQueryWithText(params map[string]string, language, customerGroupName string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Language eq '%s' and substringof('%s', CustomerGroupName)", language, customerGroupName)
	return params
}
