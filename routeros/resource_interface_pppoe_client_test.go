package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testInterfacePPPoEClientAddress = "routeros_interface_pppoe_client.test"

func TestAccInterfacePPPoEClientTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/interface/pppoe-client", "routeros_interface_pppoe_client"),
				Steps: []resource.TestStep{
					{
						Config: testAccInterfacePPPoEClientConfig(),
						Check: resource.ComposeTestCheckFunc(
							testAccCheckInterfacePPPoEClientExists(testInterfacePPPoEClientAddress),
							resource.TestCheckResourceAttr(testInterfacePPPoEClientAddress, "name", "PPPoE-Out"),
						),
					},
				},
			})
		})
	}
}

func testAccCheckInterfacePPPoEClientExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no id is set")
		}

		return nil
	}
}

func testAccInterfacePPPoEClientConfig() string {
	return providerConfig + `

resource "routeros_interface_pppoe_client" "test" {
  interface    = "ether1" 
  password     = "StrongPass"
  service_name = "pppoeservice"
  name         = "PPPoE-Out"
  disabled     = false
  user         = "MT-User"
}
`
}