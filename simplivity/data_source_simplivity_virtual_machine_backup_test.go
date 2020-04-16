package simplivity

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccSimplivityVirtualMachineBackupDataSource_name_existing(t *testing.T) {
	name := "vm_new_back"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSimplivityVirtualMachineBackupDataSourceConfig(name),
				Check:  testBackupCheck(),
			},
		},
	})
}

func testBackupCheck() resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("data.simplivity_vm_backup.data_test", "id", "vm_new_back"),
		resource.TestCheckResourceAttr("data.simplivity_vm_backup.data_test", "name", "vm_new_back"),
	)
}

func testAccCheckSimplivityVirtualMachineBackupDataSourceConfig(name string) string {
	return fmt.Sprintf(`
data "simplivity_vm_backup" "data_test" {
  name = "%s"
}
`, name)
}
