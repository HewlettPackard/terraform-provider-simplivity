package simplivity

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccSimplivityVirtualMachineDataSource_name_existing(t *testing.T) {
	name := "test_vm"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSimplivityVirtualMachineDataSourceConfig(name),
				Check:  testVMCheck(),
			},
		},
	})
}

func testVMCheck() resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("data.simplivity_vm.data_test", "id", "test_vm"),
		resource.TestCheckResourceAttr("data.simplivity_vm.data_test", "name", "test_vm"),
	)
}

func testAccCheckSimplivityVirtualMachineDataSourceConfig(name string) string {
	return fmt.Sprintf(`
data "simplivity_vm" "data_test" {
  name = "%s"
}
`, name)
}
