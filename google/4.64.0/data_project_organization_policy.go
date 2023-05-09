// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	dataprojectorganizationpolicy "github.com/golingon/terraproviders/google/4.64.0/dataprojectorganizationpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataProjectOrganizationPolicy creates a new instance of [DataProjectOrganizationPolicy].
func NewDataProjectOrganizationPolicy(name string, args DataProjectOrganizationPolicyArgs) *DataProjectOrganizationPolicy {
	return &DataProjectOrganizationPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataProjectOrganizationPolicy)(nil)

// DataProjectOrganizationPolicy represents the Terraform data resource google_project_organization_policy.
type DataProjectOrganizationPolicy struct {
	Name string
	Args DataProjectOrganizationPolicyArgs
}

// DataSource returns the Terraform object type for [DataProjectOrganizationPolicy].
func (pop *DataProjectOrganizationPolicy) DataSource() string {
	return "google_project_organization_policy"
}

// LocalName returns the local name for [DataProjectOrganizationPolicy].
func (pop *DataProjectOrganizationPolicy) LocalName() string {
	return pop.Name
}

// Configuration returns the configuration (args) for [DataProjectOrganizationPolicy].
func (pop *DataProjectOrganizationPolicy) Configuration() interface{} {
	return pop.Args
}

// Attributes returns the attributes for [DataProjectOrganizationPolicy].
func (pop *DataProjectOrganizationPolicy) Attributes() dataProjectOrganizationPolicyAttributes {
	return dataProjectOrganizationPolicyAttributes{ref: terra.ReferenceDataResource(pop)}
}

// DataProjectOrganizationPolicyArgs contains the configurations for google_project_organization_policy.
type DataProjectOrganizationPolicyArgs struct {
	// Constraint: string, required
	Constraint terra.StringValue `hcl:"constraint,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
	// BooleanPolicy: min=0
	BooleanPolicy []dataprojectorganizationpolicy.BooleanPolicy `hcl:"boolean_policy,block" validate:"min=0"`
	// ListPolicy: min=0
	ListPolicy []dataprojectorganizationpolicy.ListPolicy `hcl:"list_policy,block" validate:"min=0"`
	// RestorePolicy: min=0
	RestorePolicy []dataprojectorganizationpolicy.RestorePolicy `hcl:"restore_policy,block" validate:"min=0"`
}
type dataProjectOrganizationPolicyAttributes struct {
	ref terra.Reference
}

// Constraint returns a reference to field constraint of google_project_organization_policy.
func (pop dataProjectOrganizationPolicyAttributes) Constraint() terra.StringValue {
	return terra.ReferenceAsString(pop.ref.Append("constraint"))
}

// Etag returns a reference to field etag of google_project_organization_policy.
func (pop dataProjectOrganizationPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(pop.ref.Append("etag"))
}

// Id returns a reference to field id of google_project_organization_policy.
func (pop dataProjectOrganizationPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pop.ref.Append("id"))
}

// Project returns a reference to field project of google_project_organization_policy.
func (pop dataProjectOrganizationPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pop.ref.Append("project"))
}

// UpdateTime returns a reference to field update_time of google_project_organization_policy.
func (pop dataProjectOrganizationPolicyAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(pop.ref.Append("update_time"))
}

// Version returns a reference to field version of google_project_organization_policy.
func (pop dataProjectOrganizationPolicyAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(pop.ref.Append("version"))
}

func (pop dataProjectOrganizationPolicyAttributes) BooleanPolicy() terra.ListValue[dataprojectorganizationpolicy.BooleanPolicyAttributes] {
	return terra.ReferenceAsList[dataprojectorganizationpolicy.BooleanPolicyAttributes](pop.ref.Append("boolean_policy"))
}

func (pop dataProjectOrganizationPolicyAttributes) ListPolicy() terra.ListValue[dataprojectorganizationpolicy.ListPolicyAttributes] {
	return terra.ReferenceAsList[dataprojectorganizationpolicy.ListPolicyAttributes](pop.ref.Append("list_policy"))
}

func (pop dataProjectOrganizationPolicyAttributes) RestorePolicy() terra.ListValue[dataprojectorganizationpolicy.RestorePolicyAttributes] {
	return terra.ReferenceAsList[dataprojectorganizationpolicy.RestorePolicyAttributes](pop.ref.Append("restore_policy"))
}
