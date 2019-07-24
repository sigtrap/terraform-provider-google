// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccTpuNode_tpuNodeBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTpuNodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTpuNode_tpuNodeBasicExample(context),
			},
			{
				ResourceName:            "google_tpu_node.tpu",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"zone"},
			},
		},
	})
}

func testAccTpuNode_tpuNodeBasicExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_tpu_tensorflow_versions" "available" { }

resource "google_tpu_node" "tpu" {
	name           = "test-tpu%{random_suffix}"
	zone           = "us-central1-b"

	accelerator_type   = "v3-8"
	tensorflow_version = "${data.google_tpu_tensorflow_versions.available.versions[0]}"
	cidr_block         = "10.2.0.0/29"
}
`, context)
}

func TestAccTpuNode_tpuNodeFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTpuNodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTpuNode_tpuNodeFullExample(context),
			},
			{
				ResourceName:            "google_tpu_node.tpu",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"zone"},
			},
		},
	})
}

func testAccTpuNode_tpuNodeFullExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_tpu_tensorflow_versions" "available" { }

resource "google_tpu_node" "tpu" {
	name               = "test-tpu%{random_suffix}"
	zone               = "us-central1-b"

	accelerator_type   = "v3-8"

	cidr_block         = "10.3.0.0/29"
	tensorflow_version = "${data.google_tpu_tensorflow_versions.available.versions[0]}"

	description = "Terraform Google Provider test TPU"
	network = "default"

	labels = {
		foo = "bar"
	}

	scheduling_config {
		preemptible = true
	}
}
`, context)
}

func testAccCheckTpuNodeDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_tpu_node" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{TpuBasePath}}projects/{{project}}/locations/{{zone}}/nodes/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("TpuNode still exists at %s", url)
		}
	}

	return nil
}
