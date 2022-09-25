package sap_api_output_formatter

type CustomerGroupReads struct {
	ConnectionKey     string `json:"connection_key"`
	Result            bool   `json:"result"`
	RedisKey          string `json:"redis_key"`
	Filepath          string `json:"filepath"`
	Product           string `json:"Product"`
	APISchema         string `json:"api_schema"`
	CustomerGroupCode string `json:"customer_group_code"`
	Deleted           string `json:"deleted"`
}

type CustomerGroup struct {
	CustomerGroup string `json:"CustomerGroup"`
	ToText        string `json:"to_Text"`
}

type Text struct {
	CustomerGroup     string `json:"CustomerGroup"`
	Language          string `json:"Language"`
	CustomerGroupName string `json:"CustomerGroupName"`
}
