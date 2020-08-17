package simplivity

import (
	"fmt"
	"testing"

	"github.com/HewlettPackard/simplivity-go/ovc"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSimplivityVirtualMachine_base(t *testing.T) {
	var vm ovc.VirtualMachine
	rn := "simplivity_vm.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSimplivityVMDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSimplivityVirtualMachine,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSimplivityVirtualMachineExists(rn, &vm),
					resource.TestCheckResourceAttr(rn, "name", "test_vm"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckSimplivityVirtualMachineExists(n string, vm *ovc.VirtualMachine) resource.TestCheckFunc {
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

func testAccCheckSimplivityVMDestroy(s *terraform.State) error {
	testAccProvider.Meta().(*Config).Client
	return nil
}

var testAccSimplivityVirtualMachine = `resource "simplivity_vm" "test" {
  name = "test_vm"
}`
