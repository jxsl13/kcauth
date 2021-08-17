package internal

import "fmt"

type KeycloakError struct {
	Err              string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (kce KeycloakError) Error() string {
	return fmt.Sprintf("%s: %s", kce.Err, kce.ErrorDescription)
}
