// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceDelegatedAdministrators,
			TypeName: "aws_organizations_delegated_administrators",
			Name:     "Delegated Administrators",
		},
		{
			Factory:  dataSourceDelegatedServices,
			TypeName: "aws_organizations_delegated_services",
			Name:     "Delegated Services",
		},
		{
			Factory:  dataSourceOrganization,
			TypeName: "aws_organizations_organization",
			Name:     "Organization",
		},
		{
			Factory:  dataSourceOrganizationalUnit,
			TypeName: "aws_organizations_organizational_unit",
			Name:     "Organizational Unit",
		},
		{
			Factory:  dataSourceOrganizationalUnitChildAccounts,
			TypeName: "aws_organizations_organizational_unit_child_accounts",
			Name:     "Organizational Unit Child Accounts",
		},
		{
			Factory:  dataSourceOrganizationalUnitDescendantAccounts,
			TypeName: "aws_organizations_organizational_unit_descendant_accounts",
			Name:     "Organizational Unit Descendant Accounts",
		},
		{
			Factory:  dataSourceOrganizationalUnitDescendantOrganizationalUnits,
			TypeName: "aws_organizations_organizational_unit_descendant_organizational_units",
			Name:     "Organizational Unit Descendant Organization Units",
		},
		{
			Factory:  dataSourceOrganizationalUnits,
			TypeName: "aws_organizations_organizational_units",
			Name:     "Organizational Unit",
		},
		{
			Factory:  dataSourcePolicies,
			TypeName: "aws_organizations_policies",
			Name:     "Policies",
		},
		{
			Factory:  dataSourcePoliciesForTarget,
			TypeName: "aws_organizations_policies_for_target",
			Name:     "Policies For Target",
		},
		{
			Factory:  dataSourcePolicy,
			TypeName: "aws_organizations_policy",
			Name:     "Policy",
		},
		{
			Factory:  dataSourceResourceTags,
			TypeName: "aws_organizations_resource_tags",
			Name:     "Resource Tags",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAccount,
			TypeName: "aws_organizations_account",
			Name:     "Account",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceDelegatedAdministrator,
			TypeName: "aws_organizations_delegated_administrator",
			Name:     "Delegated Administrator",
		},
		{
			Factory:  resourceOrganization,
			TypeName: "aws_organizations_organization",
			Name:     "Organization",
		},
		{
			Factory:  resourceOrganizationalUnit,
			TypeName: "aws_organizations_organizational_unit",
			Name:     "Organizational Unit",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourcePolicy,
			TypeName: "aws_organizations_policy",
			Name:     "Policy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourcePolicyAttachment,
			TypeName: "aws_organizations_policy_attachment",
			Name:     "Policy Attachment",
		},
		{
			Factory:  resourceResourcePolicy,
			TypeName: "aws_organizations_resource_policy",
			Name:     "Resource Policy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTag,
			TypeName: "aws_organizations_tag",
			Name:     "Organizations Resource Tag",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Organizations
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
