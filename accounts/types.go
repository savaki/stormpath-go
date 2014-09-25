package accounts

import (
	. "github.com/savaki/stormpath-go/types"
)

type Account struct {
	Href                   string `json:"href"`
	Username               string `json:"username"`
	Email                  string `json:"email"`
	FullName               string `json:"fullName"`
	GivenName              string `json:"givenName"`
	MiddleName             string `json:"middleName"`
	Surname                string `json:"surname"`
	Status                 string `json:"status"`
	CustomData             Map    `json:"customData"`
	Groups                 Map    `json:"groups"`
	GroupMemberships       Map    `json:"groupMemberships"`
	Directory              Map    `json:"directory"`
	Tenant                 Map    `json:"tenant"`
	EmailVerificationToken string `json:"emailVerificationToken"`
}
