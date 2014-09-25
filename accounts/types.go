package accounts

import (
	. "github.com/savaki/stormpath-go/types"
)

type Account struct {
	Href                   string `json:"href,omitempty,omitempty"`
	Username               string `json:"username,omitempty"`
	Email                  string `json:"email,omitempty"`
	FullName               string `json:"fullName,omitempty"`
	GivenName              string `json:"givenName,omitempty"`
	MiddleName             string `json:"middleName,omitempty"`
	Surname                string `json:"surname,omitempty"`
	Status                 string `json:"status,omitempty"`
	CustomData             Map    `json:"customData,omitempty"`
	Groups                 Map    `json:"groups,omitempty"`
	GroupMemberships       Map    `json:"groupMemberships,omitempty"`
	Directory              Map    `json:"directory,omitempty"`
	Tenant                 Map    `json:"tenant,omitempty"`
	EmailVerificationToken string `json:"emailVerificationToken,omitempty"`
}
