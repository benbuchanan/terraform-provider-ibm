---
layout: "ibm"
page_title: "IBM : ibm_cm_object"
description: |-
  Manages cm_object.
subcategory: "Catalog Management API"
---

# ibm_cm_object

Provides a resource for cm_object. This allows cm_object to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cm_object" "cm_object" {
  catalog_identifier = ibm_cm_catalog.cm_catalog.id
  publish {
		permit_ibm_public_publish = true
		ibm_approved = true
		public_approved = true
		portal_approval_record = "portal_approval_record"
		portal_url = "portal_url"
  }
  state {
		current = "current"
		current_entered = "2021-01-31T09:44:12Z"
		pending = "pending"
		pending_requested = "2021-01-31T09:44:12Z"
		previous = "previous"
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `catalog_id` - (Optional, String) The id of the catalog containing this offering.
* `catalog_identifier` - (Required, Forces new resource, String) Catalog identifier.
* `catalog_name` - (Optional, String) The name of the catalog.
* `created` - (Optional, String) The date and time this catalog was created.
* `crn` - (Optional, String) The crn for this specific object.
* `data` - (Optional, Map) Map of data values for this object.
* `kind` - (Optional, String) Kind of object.
* `label` - (Optional, String) Display name in the requested language.
* `label_i18n` - (Optional, Map) A map of translated strings, by language code.
* `name` - (Optional, String) The programmatic name of this object.
* `parent_id` - (Optional, String) The parent for this specific object.
* `publish` - (Optional, List) Publish information.
Nested scheme for **publish**:
	* `ibm_approved` - (Optional, Boolean) Indicates if this offering has been approved for use by all IBMers.
	* `permit_ibm_public_publish` - (Optional, Boolean) Is it permitted to request publishing to IBM or Public.
	* `portal_approval_record` - (Optional, String) The portal's approval record ID.
	* `portal_url` - (Optional, String) The portal UI URL.
	* `public_approved` - (Optional, Boolean) Indicates if this offering has been approved for use by all IBM Cloud users.
* `short_description` - (Optional, String) Short description in the requested language.
* `short_description_i18n` - (Optional, Map) A map of translated strings, by language code.
* `state` - (Optional, List) Offering state.
Nested scheme for **state**:
	* `current` - (Optional, String) one of: new, validated, account-published, ibm-published, public-published.
	* `current_entered` - (Optional, String) Date and time of current request.
	* `pending` - (Optional, String) one of: new, validated, account-published, ibm-published, public-published.
	* `pending_requested` - (Optional, String) Date and time of pending request.
	* `previous` - (Optional, String) one of: new, validated, account-published, ibm-published, public-published.
* `tags` - (Optional, List) List of tags associated with this catalog.
* `updated` - (Optional, String) The date and time this catalog was last updated.
* `url` - (Optional, String) The url for this specific object.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cm_object.
* `catalog_object_id` - (String) unique id.
* `rev` - (String) Cloudant revision.

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:
```
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` and `iaas_classic_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`
  - Select `Classic Infrastructure API Keys` option from view dropdown for `iaas_classic_api_key`
2. For iaas_classic_username
  - Go to [Users](https://cloud.ibm.com/iam/users)
  - Click on user.
  - Find user name in the `VPN password` section under `User Details` tab

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_cm_object` resource by using `id`.
The `id` property can be formed from `catalog_identifier`, and `object_identifier` in the following format:

```
<catalog_identifier>/<object_identifier>
```
* `catalog_identifier`: A string. Catalog identifier.
* `object_identifier`: A string. Object identifier.

# Syntax
```
$ terraform import ibm_cm_object.cm_object <catalog_identifier>/<object_identifier>
```
