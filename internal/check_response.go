package internal

import (
	"encoding/json"

	resty "github.com/go-resty/resty/v2"
)

func checkResponse(resp *resty.Response) (*resty.Response, error) {
	if resp.StatusCode()/100 > 3 {
		var kce KeycloakError
		err := json.Unmarshal(resp.Body(), &kce)
		if err != nil {
			return nil, err
		}
		return nil, kce
	}
	return resp, nil
}
