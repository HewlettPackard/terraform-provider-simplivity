package simplivity

import (
	"fmt"
	"testing"

	"github.com/HewlettPackard/simplivity-go/ovc"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSimplivityVirtualMachineBackup_base(t *testing.T) {
	var backup ovc.Backup
	rn := "simplivity_vm_backup.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSimplivityVirtualMachineBackup,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSimplivityVirtualMachineBackupExists(rn, &backup),
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

func testAccCheckSimplivityVirtualMachineBackupExists(n string, backup *ovc.Backup) resource.TestCheckFunc {
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

		testBackup, err := config.Client.Backups.GetByName(rs.Primary.ID)
		if err != nil {
			return err
		}

		if testBackup.Name != rs.Primary.ID {
			return fmt.Errorf("Instance not found")
		}

		*backup = *testBackup
		return nil
	}
}

var testAccSimplivityVirtualMachineBackup = `resource "simplivity_vm_backup" "test" {
  name = "vm_test_backup"
  vm_name = "test_1"
}`
