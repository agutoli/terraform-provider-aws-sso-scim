package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceGroupMember(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceGroupMember,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"aws_sso_scim_group_member.foo", "id"),
				),
			},
		},
	})
}

const testAccResourceGroupMember = `
data "aws_sso_scim_user" "foo" {
  user_name = "terraform-test-permanent-user"
}
data "aws_sso_scim_group" "foo" {
  display_name = "terraform-test-permanent-group"
}
resource "aws_sso_scim_group_member" "foo" {
  group_id = data.aws_sso_scim_group.foo.id
	user_id = data.aws_sso_scim_user.foo.id
}
`
