// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	organizationsorganization "github.com/golingon/terraproviders/aws/4.60.0/organizationsorganization"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewOrganizationsOrganization(name string, args OrganizationsOrganizationArgs) *OrganizationsOrganization {
	return &OrganizationsOrganization{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OrganizationsOrganization)(nil)

type OrganizationsOrganization struct {
	Name  string
	Args  OrganizationsOrganizationArgs
	state *organizationsOrganizationState
}

func (oo *OrganizationsOrganization) Type() string {
	return "aws_organizations_organization"
}

func (oo *OrganizationsOrganization) LocalName() string {
	return oo.Name
}

func (oo *OrganizationsOrganization) Configuration() interface{} {
	return oo.Args
}

func (oo *OrganizationsOrganization) Attributes() organizationsOrganizationAttributes {
	return organizationsOrganizationAttributes{ref: terra.ReferenceResource(oo)}
}

func (oo *OrganizationsOrganization) ImportState(av io.Reader) error {
	oo.state = &organizationsOrganizationState{}
	if err := json.NewDecoder(av).Decode(oo.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oo.Type(), oo.LocalName(), err)
	}
	return nil
}

func (oo *OrganizationsOrganization) State() (*organizationsOrganizationState, bool) {
	return oo.state, oo.state != nil
}

func (oo *OrganizationsOrganization) StateMust() *organizationsOrganizationState {
	if oo.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oo.Type(), oo.LocalName()))
	}
	return oo.state
}

func (oo *OrganizationsOrganization) DependOn() terra.Reference {
	return terra.ReferenceResource(oo)
}

type OrganizationsOrganizationArgs struct {
	// AwsServiceAccessPrincipals: set of string, optional
	AwsServiceAccessPrincipals terra.SetValue[terra.StringValue] `hcl:"aws_service_access_principals,attr"`
	// EnabledPolicyTypes: set of string, optional
	EnabledPolicyTypes terra.SetValue[terra.StringValue] `hcl:"enabled_policy_types,attr"`
	// FeatureSet: string, optional
	FeatureSet terra.StringValue `hcl:"feature_set,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Accounts: min=0
	Accounts []organizationsorganization.Accounts `hcl:"accounts,block" validate:"min=0"`
	// NonMasterAccounts: min=0
	NonMasterAccounts []organizationsorganization.NonMasterAccounts `hcl:"non_master_accounts,block" validate:"min=0"`
	// Roots: min=0
	Roots []organizationsorganization.Roots `hcl:"roots,block" validate:"min=0"`
	// DependsOn contains resources that OrganizationsOrganization depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type organizationsOrganizationAttributes struct {
	ref terra.Reference
}

func (oo organizationsOrganizationAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(oo.ref.Append("arn"))
}

func (oo organizationsOrganizationAttributes) AwsServiceAccessPrincipals() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](oo.ref.Append("aws_service_access_principals"))
}

func (oo organizationsOrganizationAttributes) EnabledPolicyTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](oo.ref.Append("enabled_policy_types"))
}

func (oo organizationsOrganizationAttributes) FeatureSet() terra.StringValue {
	return terra.ReferenceString(oo.ref.Append("feature_set"))
}

func (oo organizationsOrganizationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(oo.ref.Append("id"))
}

func (oo organizationsOrganizationAttributes) MasterAccountArn() terra.StringValue {
	return terra.ReferenceString(oo.ref.Append("master_account_arn"))
}

func (oo organizationsOrganizationAttributes) MasterAccountEmail() terra.StringValue {
	return terra.ReferenceString(oo.ref.Append("master_account_email"))
}

func (oo organizationsOrganizationAttributes) MasterAccountId() terra.StringValue {
	return terra.ReferenceString(oo.ref.Append("master_account_id"))
}

func (oo organizationsOrganizationAttributes) Accounts() terra.ListValue[organizationsorganization.AccountsAttributes] {
	return terra.ReferenceList[organizationsorganization.AccountsAttributes](oo.ref.Append("accounts"))
}

func (oo organizationsOrganizationAttributes) NonMasterAccounts() terra.ListValue[organizationsorganization.NonMasterAccountsAttributes] {
	return terra.ReferenceList[organizationsorganization.NonMasterAccountsAttributes](oo.ref.Append("non_master_accounts"))
}

func (oo organizationsOrganizationAttributes) Roots() terra.ListValue[organizationsorganization.RootsAttributes] {
	return terra.ReferenceList[organizationsorganization.RootsAttributes](oo.ref.Append("roots"))
}

type organizationsOrganizationState struct {
	Arn                        string                                             `json:"arn"`
	AwsServiceAccessPrincipals []string                                           `json:"aws_service_access_principals"`
	EnabledPolicyTypes         []string                                           `json:"enabled_policy_types"`
	FeatureSet                 string                                             `json:"feature_set"`
	Id                         string                                             `json:"id"`
	MasterAccountArn           string                                             `json:"master_account_arn"`
	MasterAccountEmail         string                                             `json:"master_account_email"`
	MasterAccountId            string                                             `json:"master_account_id"`
	Accounts                   []organizationsorganization.AccountsState          `json:"accounts"`
	NonMasterAccounts          []organizationsorganization.NonMasterAccountsState `json:"non_master_accounts"`
	Roots                      []organizationsorganization.RootsState             `json:"roots"`
}
