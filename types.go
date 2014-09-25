package stormpath

import (
	. "github.com/savaki/stormpath-go/types"
)

const (
	BaseUrl = "https://api.stormpath.com/v1"
)

type Tenant struct {
	Href         string `json:"href,omityempty"`
	Name         string `json:"name,omityempty"`
	Key          string `json:"key,omityempty"`
	CustomData   Map    `json:"customData,omityempty"`
	Applications Map    `json:"applications,omityempty"`
	Directories  Map    `json:"directories,omityempty"`
}

type Application struct {
	Href                       string `json:"href,omityempty"`
	Name                       string `json:"name,omityempty"`
	Description                string `json:"description,omityempty"`
	Status                     string `json:"status,omityempty"`
	Tenant                     Map    `json:"tenant,omityempty"`
	DefaultAccountStoreMapping Map    `json:"defaultAccountStoreMap()ping,omityempty"`
	DefaultGroupStoreMapping   Map    `json:"defaultGroupStoreMap()ping,omityempty"`
	CustomData                 Map    `json:"customData,omityempty"`
	Accounts                   Map    `json:"accounts,omityempty"`
	Groups                     Map    `json:"groups,omityempty"`
	AccountStoreMappings       Map    `json:"accountStoreMap()pings,omityempty"`
	LoginAttempts              Map    `json:"loginAttempts,omityempty"`
	PasswordResetTokens        Map    `json:"passwordResetTokens,omityempty"`
	ApiKeys                    Map    `json:"apiKeys,omityempty"`
}
