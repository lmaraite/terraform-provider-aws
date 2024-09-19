package organizations_test

import (
	"fmt"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	tforganizations "github.com/hashicorp/terraform-provider-aws/internal/service/organizations"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccOrganizationsTag_basic(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_organizations_tag.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.OrganizationsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckTagDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccOrganizationsTagConfig(rName, acctest.CtKey1, acctest.CtValue1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTagExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, names.AttrKey, acctest.CtKey1),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, acctest.CtValue1),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccOrganizationsTag_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_organizations_tag.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.OrganizationsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckTagDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccOrganizationsTagConfig(rName, acctest.CtKey1, acctest.CtValue1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTagExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tforganizations.ResourceTag(), resourceName)),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccOrganizationsTag_Value(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_organizations_tag.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.OrganizationsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckTagDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccOrganizationsTagConfig(rName, acctest.CtKey1, acctest.CtValue1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTagExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, names.AttrKey, acctest.CtKey1),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, acctest.CtValue1),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccOrganizationsTagConfig(rName, acctest.CtKey1, acctest.CtValue1Updated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTagExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, names.AttrKey, acctest.CtKey1),
					resource.TestCheckResourceAttr(resourceName, names.AttrValue, acctest.CtValue1Updated),
				),
			},
		},
	})
}

func testAccOrganizationsTagConfig(rName string, key string, value string) string {
	return fmt.Sprintf(`
data "aws_iam_policy_document" "test" {
  statement {
    effect    = "Allow"
    actions   = ["*"]
    resources = ["*"]
  }
}

resource "aws_organizations_policy" "test" {
  name    = %[1]q
  content = data.aws_iam_policy_document.test.json
  lifecycle {
    ignore_changes = [tags]
  }
}

resource "aws_organizations_tag" "test" {
  resource_id = aws_organizations_policy.test.id
  key         = %[2]q
  value       = %[3]q
}
`, rName, key, value)
}
