---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_vpn_servers"
description: |-
  Manages IBM Cloud VPN Servers.
---

# ibm_is_vpn_servers

Provides a read-only data source for VPNServerCollection. For more information, about VPN Server, see [Creating a VPN server](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-create-server&interface=ui).

## Example Usage

```terraform
data "ibm_is_vpn_servers" "example" {
}
```

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the VPNServerCollection.
* `vpn_servers` - (Required, List) Collection of VPN servers.
	
	Nested scheme for **vpn_servers**:
	* `certificate` - (Required, List) The certificate instance for this VPN server.
		
		Nested scheme for **certificate**:
		* `crn` - (Required, String) The CRN for this certificate instance.
	* `client_authentication` - (Required, List) The methods used to authenticate VPN clients to this VPN server. VPN clients must authenticate against all provided methods.
	  * Constraints: The minimum length is `1` item.
		
		Nested scheme for **client_authentication**:
		* `method` - (Required, String) The type of authentication.
		  * Constraints: Allowable values are: certificate, username
		* `identity_provider` - (Optional, List) The type of identity provider to be used by VPN client.
			
			Nested scheme for **identity_provider**:
			* `provider_type` - (Optional, String) The type of identity provider to be used by the VPN client.- `iam`: IBM identity and access managementThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the route on which the unexpected property value was encountered.
			  * Constraints: Allowable values are: iam
		* `client_ca` - (Optional, List) The certificate instance used for the VPN client certificate authority (CA).
			
			Nested scheme for **client_ca**:
			* `crn` - (Required, String) The CRN for this certificate instance.
		* `crl` - (Optional, String) The certificate revocation list contents, encoded in PEM format.
	* `client_auto_delete` - (Required, Boolean) If set to `true`, disconnected VPN clients will be automatically deleted after the `client_auto_delete_timeout` time has passed.
	* `client_auto_delete_timeout` - (Required, Integer) Hours after which disconnected VPN clients will be automatically deleted. If `0`, disconnected VPN clients will be deleted immediately.
	  * Constraints: The maximum value is `24`. The minimum value is `0`.
	* `client_dns_server_ips` - (Required, List) The DNS server addresses that will be provided to VPN clients that are connected to this VPN server.
		
		Nested scheme for **client_dns_server_ips**:
		* `address` - (Required, String) The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
		  * Constraints: The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`
	* `client_idle_timeout` - (Required, Integer) The seconds a VPN client can be idle before this VPN server will disconnect it.  If `0`, the server will not disconnect idle clients.
	  * Constraints: The maximum value is `28800`. The minimum value is `0`.
	* `client_ip_pool` - (Required, String) The VPN client IPv4 address pool, expressed in CIDR format.
	  * Constraints: The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$/`
	* `created_at` - (Required, String) The date and time that the VPN server was created.
	* `crn` - (Required, String) The CRN for this VPN server.
	* `enable_split_tunneling` - (Required, Boolean) Indicates whether the split tunneling is enabled on this VPN server.
	* `health_state` - (Required, String) The health of this resource.- `ok`: Healthy- `degraded`: Suffering from compromised performance, capacity, or connectivity- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle state. A resource with a lifecycle state of `failed` or `deleting` will have a health state of `inapplicable`. A `pending` resource may also have this state.
	  * Constraints: Allowable values are: ok, degraded, faulted, inapplicable
	* `hostname` - (Required, String) Fully qualified domain name assigned to this VPN server.
	* `href` - (Required, String) The URL for this VPN server.
	  * Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`
	* `id` - (Required, String) The unique identifier for this VPN server.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`
	* `lifecycle_state` - (Required, String) The lifecycle state of the VPN server.
	  * Constraints: Allowable values are: deleting, failed, pending, stable, updating, waiting, suspended
	* `name` - (Required, String) The unique user-defined name for this VPN server.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`
	* `port` - (Required, Integer) The port number used by this VPN server.
	  * Constraints: The maximum value is `65535`. The minimum value is `1`.
	* `private_ips` - (Required, List) The reserved IPs bound to this VPN server.
		
		Nested scheme for **private_ips**:
		* `address` - (Required, String) The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
		  * Constraints: The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`
		* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
			
			Nested scheme for **deleted**:
			* `more_info` - (Required, String) Link to documentation about deleted resources.
			  * Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`
		* `href` - (Required, String) The URL for this reserved IP.
		  * Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`
		* `id` - (Required, String) The unique identifier for this reserved IP.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`
		* `name` - (Required, String) The user-defined or system-provided name for this reserved IP.
		  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`
		* `resource_type` - (Required, String) The resource type.
		  * Constraints: Allowable values are: subnet_reserved_ip
	* `protocol` - (Required, String) The transport protocol used by this VPN server.
	  * Constraints: Allowable values are: udp, tcp
	* `resource_group` - (Required, List) The resource group for this VPN server.
		
		Nested scheme for **resource_group**:
		* `href` - (Required, String) The URL for this resource group.
		  * Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`
		* `id` - (Required, String) The unique identifier for this resource group.
		  * Constraints: The value must match regular expression `/^[0-9a-f]{32}$/`
		* `name` - (Required, String) The user-defined name for this resource group.
		  * Constraints: The maximum length is `40` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-_ ]+$/`
	* `resource_type` - (Required, String) The type of resource referenced.
	  * Constraints: Allowable values are: vpn_server
	* `security_groups` - (Required, List) The security groups targeting this VPN server.
		
		Nested scheme for **security_groups**:
		* `crn` - (Required, String) The security group's CRN.
		* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
			
			Nested scheme for **deleted**:
			* `more_info` - (Required, String) Link to documentation about deleted resources.
			  * Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`
		* `href` - (Required, String) The security group's canonical URL.
		  * Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`
		* `id` - (Required, String) The unique identifier for this security group.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`
		* `name` - (Required, String) The user-defined name for this security group. Names must be unique within the VPC the security group resides in.
		  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`
	* `subnets` - (Required, List) The subnets this VPN server is part of.
		
		Nested scheme for **subnets**:
		* `crn` - (Required, String) The CRN for this subnet.
		* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
			
			Nested scheme for **deleted**:
			* `more_info` - (Required, String) Link to documentation about deleted resources.
			  * Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`
		* `href` - (Required, String) The URL for this subnet.
		  * Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`
		* `id` - (Required, String) The unique identifier for this subnet.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`
		* `name` - (Required, String) The user-defined name for this subnet.
		  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`
		* `resource_type` - (Required, String) The resource type.
		  * Constraints: Allowable values are: subnet
