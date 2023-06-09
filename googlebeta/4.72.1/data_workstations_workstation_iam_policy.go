// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataWorkstationsWorkstationIamPolicy creates a new instance of [DataWorkstationsWorkstationIamPolicy].
func NewDataWorkstationsWorkstationIamPolicy(name string, args DataWorkstationsWorkstationIamPolicyArgs) *DataWorkstationsWorkstationIamPolicy {
	return &DataWorkstationsWorkstationIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataWorkstationsWorkstationIamPolicy)(nil)

// DataWorkstationsWorkstationIamPolicy represents the Terraform data resource google_workstations_workstation_iam_policy.
type DataWorkstationsWorkstationIamPolicy struct {
	Name string
	Args DataWorkstationsWorkstationIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataWorkstationsWorkstationIamPolicy].
func (wwip *DataWorkstationsWorkstationIamPolicy) DataSource() string {
	return "google_workstations_workstation_iam_policy"
}

// LocalName returns the local name for [DataWorkstationsWorkstationIamPolicy].
func (wwip *DataWorkstationsWorkstationIamPolicy) LocalName() string {
	return wwip.Name
}

// Configuration returns the configuration (args) for [DataWorkstationsWorkstationIamPolicy].
func (wwip *DataWorkstationsWorkstationIamPolicy) Configuration() interface{} {
	return wwip.Args
}

// Attributes returns the attributes for [DataWorkstationsWorkstationIamPolicy].
func (wwip *DataWorkstationsWorkstationIamPolicy) Attributes() dataWorkstationsWorkstationIamPolicyAttributes {
	return dataWorkstationsWorkstationIamPolicyAttributes{ref: terra.ReferenceDataResource(wwip)}
}

// DataWorkstationsWorkstationIamPolicyArgs contains the configurations for google_workstations_workstation_iam_policy.
type DataWorkstationsWorkstationIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// WorkstationClusterId: string, required
	WorkstationClusterId terra.StringValue `hcl:"workstation_cluster_id,attr" validate:"required"`
	// WorkstationConfigId: string, required
	WorkstationConfigId terra.StringValue `hcl:"workstation_config_id,attr" validate:"required"`
	// WorkstationId: string, required
	WorkstationId terra.StringValue `hcl:"workstation_id,attr" validate:"required"`
}
type dataWorkstationsWorkstationIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_workstations_workstation_iam_policy.
func (wwip dataWorkstationsWorkstationIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(wwip.ref.Append("etag"))
}

// Id returns a reference to field id of google_workstations_workstation_iam_policy.
func (wwip dataWorkstationsWorkstationIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wwip.ref.Append("id"))
}

// Location returns a reference to field location of google_workstations_workstation_iam_policy.
func (wwip dataWorkstationsWorkstationIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(wwip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_workstations_workstation_iam_policy.
func (wwip dataWorkstationsWorkstationIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(wwip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_workstations_workstation_iam_policy.
func (wwip dataWorkstationsWorkstationIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(wwip.ref.Append("project"))
}

// WorkstationClusterId returns a reference to field workstation_cluster_id of google_workstations_workstation_iam_policy.
func (wwip dataWorkstationsWorkstationIamPolicyAttributes) WorkstationClusterId() terra.StringValue {
	return terra.ReferenceAsString(wwip.ref.Append("workstation_cluster_id"))
}

// WorkstationConfigId returns a reference to field workstation_config_id of google_workstations_workstation_iam_policy.
func (wwip dataWorkstationsWorkstationIamPolicyAttributes) WorkstationConfigId() terra.StringValue {
	return terra.ReferenceAsString(wwip.ref.Append("workstation_config_id"))
}

// WorkstationId returns a reference to field workstation_id of google_workstations_workstation_iam_policy.
func (wwip dataWorkstationsWorkstationIamPolicyAttributes) WorkstationId() terra.StringValue {
	return terra.ReferenceAsString(wwip.ref.Append("workstation_id"))
}
