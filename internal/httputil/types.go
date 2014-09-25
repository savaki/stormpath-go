package httputil

import "net/http"

type ErrorMessage struct {
	Status           int           `json:"status"`
	Code             int           `json:"code"`
	Message          string        `json:"message"`
	DeveloperMessage string        `json:"developerMessage"`
	MoreInfo         string        `json:"moreInfo"`
	Request          *http.Request `json:"-"`
}

func (e *ErrorMessage) Error() string {
	return e.Message
}
