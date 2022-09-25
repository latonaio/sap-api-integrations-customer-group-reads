package responses

type Text struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			CustomerGroup     string `json:"CustomerGroup"`
			Language          string `json:"Language"`
			CustomerGroupName string `json:"CustomerGroupName"`
		} `json:"results"`
	} `json:"d"`
}
