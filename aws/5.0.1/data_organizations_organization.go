// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataorganizationsorganization "github.com/golingon/terraproviders/aws/5.0.1/dataorganizationsorganization"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataOrganizationsOrganization creates a new instance of [DataOrganizationsOrganization].
func NewDataOrganizationsOrganization(name string, args DataOrganizationsOrganizationArgs) *DataOrganizationsOrganization {
	return &DataOrganizationsOrganization{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataOrganizationsOrganization)(nil)

// DataOrganizationsOrganization represents the Terraform data resource aws_organizations_organization.
type DataOrganizationsOrganization struct {
	Name string
	Args DataOrganizationsOrganizationArgs
}

// DataSource returns the Terraform object type for [DataOrganizationsOrganization].
func (oo *DataOrganizationsOrganization) DataSource() string {
	return "aws_organizations_organization"
}

// LocalName returns the local name for [DataOrganizationsOrganization].
func (oo *DataOrganizationsOrganization) LocalName() string {
	return oo.Name
}

// Configuration returns the configuration (args) for [DataOrganizationsOrganization].
func (oo *DataOrganizationsOrganization) Configuration() interface{} {
	return oo.Args
}

// Attributes returns the attributes for [DataOrganizationsOrganization].
func (oo *DataOrganizationsOrganization) Attributes() dataOrganizationsOrganizationAttributes {
	return dataOrganizationsOrganizationAttributes{ref: terra.ReferenceDataResource(oo)}
}

// DataOrganizationsOrganizationArgs contains the configurations for aws_organizations_organization.
type DataOrganizationsOrganizationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Accounts: min=0
	Accounts []dataorganizationsorganization.Accounts `hcl:"accounts,block" validate:"min=0"`
	// NonMasterAccounts: min=0
	NonMasterAccounts []dataorganizationsorganization.NonMasterAccounts `hcl:"non_master_accounts,block" validate:"min=0"`
	// Roots: min=0
	Roots []dataorganizationsorganization.Roots `hcl:"roots,block" validate:"min=0"`
}
type dataOrganizationsOrganizationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_organizations_organization.
func (oo dataOrganizationsOrganizationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("arn"))
}

// AwsServiceAccessPrincipals returns a reference to field aws_service_access_principals of aws_organizations_organization.
func (oo dataOrganizationsOrganizationAttributes) AwsServiceAccessPrincipals() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](oo.ref.Append("aws_service_access_principals"))
}

// EnabledPolicyTypes returns a reference to field enabled_policy_types of aws_organizations_organization.
func (oo dataOrganizationsOrganizationAttributes) EnabledPolicyTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](oo.ref.Append("enabled_policy_types"))
}

// FeatureSet returns a reference to field feature_set of aws_organizations_organization.
func (oo dataOrganizationsOrganizationAttributes) FeatureSet() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("feature_set"))
}

// Id returns a reference to field id of aws_organizations_organization.
func (oo dataOrganizationsOrganizationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("id"))
}

// MasterAccountArn returns a reference to field master_account_arn of aws_organizations_organization.
func (oo dataOrganizationsOrganizationAttributes) MasterAccountArn() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("master_account_arn"))
}

// MasterAccountEmail returns a reference to field master_account_email of aws_organizations_organization.
func (oo dataOrganizationsOrganizationAttributes) MasterAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("master_account_email"))
}

// MasterAccountId returns a reference to field master_account_id of aws_organizations_organization.
func (oo dataOrganizationsOrganizationAttributes) MasterAccountId() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("master_account_id"))
}

func (oo dataOrganizationsOrganizationAttributes) Accounts() terra.ListValue[dataorganizationsorganization.AccountsAttributes] {
	return terra.ReferenceAsList[dataorganizationsorganization.AccountsAttributes](oo.ref.Append("accounts"))
}

func (oo dataOrganizationsOrganizationAttributes) NonMasterAccounts() terra.ListValue[dataorganizationsorganization.NonMasterAccountsAttributes] {
	return terra.ReferenceAsList[dataorganizationsorganization.NonMasterAccountsAttributes](oo.ref.Append("non_master_accounts"))
}

func (oo dataOrganizationsOrganizationAttributes) Roots() terra.ListValue[dataorganizationsorganization.RootsAttributes] {
	return terra.ReferenceAsList[dataorganizationsorganization.RootsAttributes](oo.ref.Append("roots"))
}