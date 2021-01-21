# ApiUser

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | The userid which is required for login. This value cannot be changed. | [optional] [default to null]
**FirstName** | **string** | The first name of the user. | [optional] [default to null]
**LastName** | **string** | The last name of the user. | [optional] [default to null]
**EmailAddress** | **string** | The email address associated with the user. | [optional] [default to null]
**Source** | **string** | The user source which is the origin of this user. This value cannot be changed. | [optional] [default to null]
**Status** | **string** | The user&#x27;s status, e.g. active or disabled. | [default to null]
**ReadOnly** | **bool** | Indicates whether the user&#x27;s properties could be modified by the Nexus Repository Manager. When false only roles are considered during update. | [optional] [default to null]
**Roles** | **[]string** | The roles which the user has been assigned within Nexus. | [optional] [default to null]
**ExternalRoles** | **[]string** | The roles which the user has been assigned in an external source, e.g. LDAP group. These cannot be changed within the Nexus Repository Manager. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

