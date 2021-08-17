package internal

import (
	"encoding/json"

	resty "github.com/go-resty/resty/v2"
)

func checkResponse(resp *resty.Response) (*resty.Response, error) {
	if resp.StatusCode()/100 > 3 {
		var err KeycloakError
		json.Unmarshal(resp.Body(), err)
		return nil, err
	}
	return resp, nil
}
