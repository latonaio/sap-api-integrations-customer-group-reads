package responses

type CustomerGroup struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			CustomerGroup string `json:"CustomerGroup"`
			ToText        struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Text"`
		} `json:"results"`
	} `json:"d"`
}
