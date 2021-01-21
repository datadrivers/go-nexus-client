# CreateLdapServerXo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | LDAP server name | [default to null]
**Protocol** | **string** | LDAP server connection Protocol to use | [default to null]
**UseTrustStore** | **bool** | Whether to use certificates stored in Nexus Repository Manager&#x27;s truststore | [optional] [default to null]
**Host** | **string** | LDAP server connection hostname | [default to null]
**Port** | **int32** | LDAP server connection port to use | [default to null]
**SearchBase** | **string** | LDAP location to be added to the connection URL | [default to null]
**AuthScheme** | **string** | Authentication scheme used for connecting to LDAP server | [default to null]
**AuthRealm** | **string** | The SASL realm to bind to. Required if authScheme is CRAM_MD5 or DIGEST_MD5 | [optional] [default to null]
**AuthUsername** | **string** | This must be a fully qualified username if simple authentication is used. Required if authScheme other than none. | [optional] [default to null]
**ConnectionTimeoutSeconds** | **int32** | How long to wait before timeout | [default to null]
**ConnectionRetryDelaySeconds** | **int32** | How long to wait before retrying | [default to null]
**MaxIncidentsCount** | **int32** | How many retry attempts | [default to null]
**UserBaseDn** | **string** | The relative DN where user objects are found (e.g. ou&#x3D;people). This value will have the Search base DN value appended to form the full User search base DN. | [optional] [default to null]
**UserSubtree** | **bool** | Are users located in structures below the user base DN? | [optional] [default to null]
**UserObjectClass** | **string** | LDAP class for user objects | [optional] [default to null]
**UserLdapFilter** | **string** | LDAP search filter to limit user search | [optional] [default to null]
**UserIdAttribute** | **string** | This is used to find a user given its user ID | [optional] [default to null]
**UserRealNameAttribute** | **string** | This is used to find a real name given the user ID | [optional] [default to null]
**UserEmailAddressAttribute** | **string** | This is used to find an email address given the user ID | [optional] [default to null]
**UserPasswordAttribute** | **string** | If this field is blank the user will be authenticated against a bind with the LDAP server | [optional] [default to null]
**LdapGroupsAsRoles** | **bool** | Denotes whether LDAP assigned roles are used as Nexus Repository Manager roles | [optional] [default to null]
**GroupType** | **string** | Defines a type of groups used: static (a group contains a list of users) or dynamic (a user contains a list of groups). Required if ldapGroupsAsRoles is true. | [default to null]
**GroupBaseDn** | **string** | The relative DN where group objects are found (e.g. ou&#x3D;Group). This value will have the Search base DN value appended to form the full Group search base DN. | [optional] [default to null]
**GroupSubtree** | **bool** | Are groups located in structures below the group base DN | [optional] [default to null]
**GroupObjectClass** | **string** | LDAP class for group objects. Required if groupType is static | [optional] [default to null]
**GroupIdAttribute** | **string** | This field specifies the attribute of the Object class that defines the Group ID. Required if groupType is static | [optional] [default to null]
**GroupMemberAttribute** | **string** | LDAP attribute containing the usernames for the group. Required if groupType is static | [optional] [default to null]
**GroupMemberFormat** | **string** | The format of user ID stored in the group member attribute. Required if groupType is static | [optional] [default to null]
**UserMemberOfAttribute** | **string** | Set this to the attribute used to store the attribute which holds groups DN in the user object. Required if groupType is dynamic | [optional] [default to null]
**AuthPassword** | **string** | The password to bind with. Required if authScheme other than none. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

