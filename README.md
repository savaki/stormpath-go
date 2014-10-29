stormpath-go
============

[![GoDoc](https://godoc.org/github.com/savaki/stormpath-go?status.svg)](https://godoc.org/github.com/savaki/stormpath-go)

## Tenant API

| Method | Path | Description | Status |
|---|---|---|
| GET | /v1/tenants/:tenantId | retrieve a specific tenant | TBD |
| GET | /v1/tenants/:tenantId/applications | retrieve applications associated with tenant | TBD |
| GET | /v1/tenants/:tenantId/directories | retrieve directories associated with tenant | TBD |
| GET | /v1/tenants/:tenantId/groups | retrieve groups associated with tenant | TBD |

## Application API

| Method | Path | Description | Status |
|---|---|---|
| GET | /v1/applications/:applicationId | retrieve application info | TBD |
| DELETE | /v1/applications/:applicationId | delete an application | TBD |
| GET | /v1/applications/:applicationId/accounts | all accounts accessible to the application | TBD |
| POST | /v1/applications/:applicationId/accounts | register a new account | TBD |
| GET | /v1/applications/:applicationId/groups | groups associated with an application | TBD |
| POST | /v1/applications/:applicationId/groups | create a new application group | TBD |
| POST | /v1/applications/:applicationId/loginAttempts | login user | TBD |
| POST | /v1/applications/:applicationId/passwordResetTokens | reset password token | TBD | 
| POST | /v1/applications/:applicationId/verificationEmails | resend verification email | TBD |
| GET | /v1/applications/:applicationId/accountStoreMappings | all groups and directories that are assigned to that application for the purpose of providing accounts that may login to the application | TBD |

## Account Store Mappings API

| Method | Path | Description | Status |
|---|---|---|
| POST | /v1/accountStoreMappings | create a new account storm mapping | TBD |
| GET | /v1/accountStoreMappings/:accountStoreMappingId | retrieve an individual account storm mapping | TBD |
| POST | /v1/accountStoreMappings/:accountStoreMappingId | update an individual account storm mapping | TBD |
| DELETE | /v1/accountStoreMappings/:accountStoreMappingId | deletes an individual account storm mapping | TBD |

## Directory Resource API

| Method | Path | Description | Status |
|---|---|---|
| GET | /v1/directories/:directoryId | retrieve a directory resource | TBD | 
| POST | /v1/directories/:directoryId | update a directory resource | TBD | 
| DELETE | /v1/directories/:directoryId | delete a directory resource | TBD | 
| GET | /v1/directories/:directoryId/groups | list directory groups | TBD |
| GET | /v1/directories/:directoryId/accounts | list directory accounts | TBD | 
| POST | /v1/directories/:directoryId/groups | create a new group | TBD |

## Group Resource API

| Method | Path | Description | Status |
|---|---|---|
| GET | /v1/groups/:groupId | retrieve a group | TBD |
| POST | /v1/groups/:groupId | update a group | TBD |
| DELETE | /v1/groups/:groupId | delete a group | TBD |
| GET | /v1/applications/:applicationUid/groups | all application accessible groups | TBD |
| GET | /v1/accounts/:accountUid/groups | all account accessiblle groups | TBD | 
| GET | /v1/groups/:groupId/accounts | search group search | TBD |
| GET | /v1/groups/:groupId/accountMemberships | all memberships for a group | TBD |

## Group Membership API

| Method | Path | Description | Status |
|---|---|---|
| POST | /v1/groupMemberships | create a group membership | TBD |
| GET | /v1/groupMemberships/:groupMembershipId | retrieve a specific group membership | TBD |
| POST| /v1/groupMemberships/:groupMembershipId | update a specific group membership | TBD |
| DELETE | /v1/groupMemberships/:groupMembershipId | delete a specific group membership | TBD |

## Account API

| Method | Path | Description | Status |
|---|---|---|
| POST | /v1/directories/:directoryId/accounts | create an account | TBD |
| GET | /v1/accounts/:accountId | retrieve a specific account | TBD |
| POST | /v1/accounts/:accountId | update a specific account | TBD |
| DELETE | /v1/accounts/:accountId | delete a specific account | TBD |
| POST | /v1/groupMemberships | assign an account to a group | TBD | 
| POST | /v1/accounts/emailVerificationsToken/:verificationToken | verify an email address | TBD |
| GET | /v1/accounts/:accountId/groups | all groups where account is a member | TBD |
| GET | /v1/accounts/:accountId/groupMemberships | all groups where account is a member | TBD |

## Custom Data API

| Method | Path | Description | Status |
|---|---|---|
| GET | /v1/accounts/:accountId/customData | retrieve custom account data | TBD |
| POST | /v1/accounts/:accountId/customData | create custom account data | TBD |
| GET | /v1/groups/:groupId/customData | retrieve custom group data | TBD |
| POST | /v1/groups/:groupId/customData | create custom group data | TBD |
| GET | /v1/applications/:applicationId/customData | retrieve custom application data | TBD |
| POST | /v1/applications/:applicationId/customData | create custom application data | TBD |
| GET | /v1/directories/:directoryId/customData | retrieve custom directory data | TBD |
| POST | /v1/directories/:directoryId/customData | create custom directory data | TBD |
| GET | /v1/tenants/:tenantId/customData | retrieve custom tenant data | TBD |
| POST | /v1/tenants/:tenantId/customData | create custom tenant data | TBD |



