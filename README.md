stormpath-go
============

[![GoDoc](https://godoc.org/github.com/savaki/stormpath-go?status.svg)](https://godoc.org/github.com/savaki/stormpath-go)

## Tenant API

| Method | Path | Description | Done? |
| ------ | ---- | ----------- | :------: |
| GET | /v1/tenants/current | retrieve the current tenant | Yes |
| GET | /v1/tenants/:tenantId | retrieve a specific tenant | - |
| GET | /v1/tenants/:tenantId/applications | retrieve applications associated with tenant | Yes |
| GET | /v1/tenants/:tenantId/directories | retrieve directories associated with tenant | Yes |
| GET | /v1/tenants/:tenantId/groups | retrieve groups associated with tenant | - |

## Application API

| Method | Path | Description | Done? |
| ------ | ---- | ----------- | :------: |
| GET | /v1/applications/:applicationId | retrieve application info | - |
| DELETE | /v1/applications/:applicationId | delete an application | - |
| GET | /v1/applications/:applicationId/accounts | all accounts accessible to the application | - |
| POST | /v1/applications/:applicationId/accounts | register a new account | Yes |
| GET | /v1/applications/:applicationId/groups | groups associated with an application | - |
| POST | /v1/applications/:applicationId/groups | create a new application group | - |
| POST | /v1/applications/:applicationId/loginAttempts | login user | Yes |
| POST | /v1/applications/:applicationId/passwordResetTokens | reset password token | - | 
| POST | /v1/applications/:applicationId/verificationEmails | resend verification email | - |
| GET | /v1/applications/:applicationId/accountStoreMappings | all groups and directories that are assigned to that application for the purpose of providing accounts that may login to the application | - |

## Account Store Mappings API

| Method | Path | Description | Done? |
| ------ | ---- | ----------- | :------: |
| POST | /v1/accountStoreMappings | create a new account storm mapping | - |
| GET | /v1/accountStoreMappings/:accountStoreMappingId | retrieve an individual account storm mapping | - |
| POST | /v1/accountStoreMappings/:accountStoreMappingId | update an individual account storm mapping | - |
| DELETE | /v1/accountStoreMappings/:accountStoreMappingId | deletes an individual account storm mapping | - |

## Directory Resource API

| Method | Path | Description | Done? |
| ------ | ---- | ----------- | :------: |
| GET | /v1/directories/:directoryId | retrieve a directory resource | - | 
| POST | /v1/directories/:directoryId | update a directory resource | - | 
| DELETE | /v1/directories/:directoryId | delete a directory resource | - | 
| GET | /v1/directories/:directoryId/groups | list directory groups | - |
| GET | /v1/directories/:directoryId/accounts | list directory accounts | - | 
| POST | /v1/directories/:directoryId/groups | create a new group | - |

## Group Resource API

| Method | Path | Description | Done? |
| ------ | ---- | ----------- | :------: |
| GET | /v1/groups/:groupId | retrieve a group | - |
| POST | /v1/groups/:groupId | update a group | - |
| DELETE | /v1/groups/:groupId | delete a group | - |
| GET | /v1/applications/:applicationUid/groups | all application accessible groups | - |
| GET | /v1/accounts/:accountUid/groups | all account accessiblle groups | - | 
| GET | /v1/groups/:groupId/accounts | search group search | - |
| GET | /v1/groups/:groupId/accountMemberships | all memberships for a group | - |

## Group Membership API

| Method | Path | Description | Done? |
| ------ | ---- | ----------- | :------: |
| POST | /v1/groupMemberships | create a group membership | - |
| GET | /v1/groupMemberships/:groupMembershipId | retrieve a specific group membership | - |
| POST| /v1/groupMemberships/:groupMembershipId | update a specific group membership | - |
| DELETE | /v1/groupMemberships/:groupMembershipId | delete a specific group membership | - |

## Account API

| Method | Path | Description | Done? |
| ------ | ---- | ----------- | :------: |
| POST | /v1/directories/:directoryId/accounts | create an account | - |
| GET | /v1/accounts/:accountId | retrieve a specific account | - |
| POST | /v1/accounts/:accountId | update a specific account | - |
| DELETE | /v1/accounts/:accountId | delete a specific account | Yes |
| POST | /v1/groupMemberships | assign an account to a group | - | 
| POST | /v1/accounts/emailVerificationsToken/:verificationToken | verify an email address | - |
| GET | /v1/accounts/:accountId/groups | all groups where account is a member | - |
| GET | /v1/accounts/:accountId/groupMemberships | all groups where account is a member | - |

## Custom Data API

| Method | Path | Description | Done? |
| ------ | ---- | ----------- | :------: |/customData | retrieve custom account data | - |
| POST | /v1/accounts/:accountId/customData | create custom account data | - |
| GET | /v1/groups/:groupId/customData | retrieve custom group data | - |
| POST | /v1/groups/:groupId/customData | create custom group data | - |
| GET | /v1/applications/:applicationId/customData | retrieve custom application data | - |
| POST | /v1/applications/:applicationId/customData | create custom application data | - |
| GET | /v1/directories/:directoryId/customData | retrieve custom directory data | - |
| POST | /v1/directories/:directoryId/customData | create custom directory data | - |
| GET | /v1/tenants/:tenantId/customData | retrieve custom tenant data | - |
| POST | /v1/tenants/:tenantId/customData | create custom tenant data | - |


