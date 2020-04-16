package simplivity

import (
	"fmt"
	"testing"

	"github.com/HewlettPackard/simplivity-go/ovc"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSimplivityVirtualMachineClone_base(t *testing.T) {
	var vm ovc.VirtualMachine
	rn := "simplivity_vm_clone.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSimplivityVirtualMachineClone,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSimplivityVirtualMachineCloneExists(rn, &vm),
				),
			},
			{
				ResourceName:            rn,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"new_vm_name", "app_consistent"},
			},
		},
	})
}

func testAccCheckSimplivityVirtualMachineCloneExists(n string, vm *ovc.VirtualMachine) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found :%v", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config, err := testProviderConfig()
		if err != nil {
			return err
		}

		testVM, err := config.Client.VirtualMachines.GetByName(rs.Primary.ID)
		if err != nil {
			return err
		}

		if testVM.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}

		*vm = *testVM
		return nil
	}
}

var testAccSimplivityVirtualMachineClone = `resource "simplivity_vm_clone" "test" {
  name = "test_vm"
  new_vm_name = "cloned_test_vm"
}`
