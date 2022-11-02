// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCmObjectDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmObjectDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "catalog_identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "object_identifier"),
				),
			},
		},
	})
}

func TestAccIBMCmObjectDataSourceAllArgs(t *testing.T) {
	catalogObjectName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	catalogObjectCRN := fmt.Sprintf("tf_crn_%d", acctest.RandIntRange(10, 100))
	catalogObjectURL := fmt.Sprintf("tf_url_%d", acctest.RandIntRange(10, 100))
	catalogObjectParentID := fmt.Sprintf("tf_parent_id_%d", acctest.RandIntRange(10, 100))
	catalogObjectLabel := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	catalogObjectShortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	catalogObjectKind := fmt.Sprintf("tf_kind_%d", acctest.RandIntRange(10, 100))
	catalogObjectCatalogName := fmt.Sprintf("tf_catalog_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmObjectDataSourceConfig(catalogObjectName, catalogObjectCRN, catalogObjectURL, catalogObjectParentID, catalogObjectLabel, catalogObjectShortDescription, catalogObjectKind, catalogObjectCatalogName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "catalog_identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "object_identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "catalog_object_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "rev"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "parent_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "label_i18n.%"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "updated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "short_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "short_description_i18n.%"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "kind"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "publish.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "catalog_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_object.cm_object", "data.%"),
				),
			},
		},
	})
}

func testAccCheckIBMCmObjectDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
		}

		resource "ibm_cm_object" "cm_object" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
		}

		data "ibm_cm_object" "cm_object" {
			catalog_identifier = ibm_cm_object.cm_object.catalog_identifier
			object_identifier = ibm_cm_object.cm_object.catalogObject_id
		}
	`)
}

func testAccCheckIBMCmObjectDataSourceConfig(catalogObjectName string, catalogObjectCRN string, catalogObjectURL string, catalogObjectParentID string, catalogObjectLabel string, catalogObjectShortDescription string, catalogObjectKind string, catalogObjectCatalogName string) string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
		}

		resource "ibm_cm_object" "cm_object" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
			name = "%s"
			crn = "%s"
			url = "%s"
			parent_id = "%s"
			label_i18n = "FIXME"
			label = "%s"
			tags = "FIXME"
			created = "2004-10-28T04:39:00.000Z"
			updated = "2004-10-28T04:39:00.000Z"
			short_description = "%s"
			short_description_i18n = "FIXME"
			kind = "%s"
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
			catalog_id = ibm_cm_catalog.cm_catalog.id
			catalog_name = "%s"
			data = "FIXME"
		}

		data "ibm_cm_object" "cm_object" {
			catalog_identifier = ibm_cm_object.cm_object.catalog_identifier
			object_identifier = ibm_cm_object.cm_object.catalogObject_id
		}
	`, catalogObjectName, catalogObjectCRN, catalogObjectURL, catalogObjectParentID, catalogObjectLabel, catalogObjectShortDescription, catalogObjectKind, catalogObjectCatalogName)
}
