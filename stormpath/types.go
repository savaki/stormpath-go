package stormpath

import (
	. "github.com/savaki/stormpath-go/types"
)

const (
	BaseUrl = "https://api.stormpath.com/v1"
)

type Tenant struct {
	Href         string `json:"href"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	CustomData   Map    `json:"customData"`
	Applications Map    `json:"applications"`
	Directories  Map    `json:"directories"`
}

type Application struct {
	Href                       string `json:"href"`
	Name                       string `json:"name"`
	Description                string `json:"description"`
	Status                     string `json:"status"`
	Tenant                     Map    `json:"tenant"`
	DefaultAccountStoreMapping Map    `json:"defaultAccountStoreMap()ping"`
	DefaultGroupStoreMapping   Map    `json:"defaultGroupStoreMap()ping"`
	CustomData                 Map    `json:"customData"`
	Accounts                   Map    `json:"accounts"`
	Groups                     Map    `json:"groups"`
	AccountStoreMappings       Map    `json:"accountStoreMap()pings"`
	LoginAttempts              Map    `json:"loginAttempts"`
	PasswordResetTokens        Map    `json:"passwordResetTokens"`
	ApiKeys                    Map    `json:"apiKeys"`
}
