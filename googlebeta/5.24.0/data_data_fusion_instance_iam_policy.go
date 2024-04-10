// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataDataFusionInstanceIamPolicy creates a new instance of [DataDataFusionInstanceIamPolicy].
func NewDataDataFusionInstanceIamPolicy(name string, args DataDataFusionInstanceIamPolicyArgs) *DataDataFusionInstanceIamPolicy {
	return &DataDataFusionInstanceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDataFusionInstanceIamPolicy)(nil)

// DataDataFusionInstanceIamPolicy represents the Terraform data resource google_data_fusion_instance_iam_policy.
type DataDataFusionInstanceIamPolicy struct {
	Name string
	Args DataDataFusionInstanceIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataDataFusionInstanceIamPolicy].
func (dfiip *DataDataFusionInstanceIamPolicy) DataSource() string {
	return "google_data_fusion_instance_iam_policy"
}

// LocalName returns the local name for [DataDataFusionInstanceIamPolicy].
func (dfiip *DataDataFusionInstanceIamPolicy) LocalName() string {
	return dfiip.Name
}

// Configuration returns the configuration (args) for [DataDataFusionInstanceIamPolicy].
func (dfiip *DataDataFusionInstanceIamPolicy) Configuration() interface{} {
	return dfiip.Args
}

// Attributes returns the attributes for [DataDataFusionInstanceIamPolicy].
func (dfiip *DataDataFusionInstanceIamPolicy) Attributes() dataDataFusionInstanceIamPolicyAttributes {
	return dataDataFusionInstanceIamPolicyAttributes{ref: terra.ReferenceDataResource(dfiip)}
}

// DataDataFusionInstanceIamPolicyArgs contains the configurations for google_data_fusion_instance_iam_policy.
type DataDataFusionInstanceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}
type dataDataFusionInstanceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_data_fusion_instance_iam_policy.
func (dfiip dataDataFusionInstanceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dfiip.ref.Append("etag"))
}

// Id returns a reference to field id of google_data_fusion_instance_iam_policy.
func (dfiip dataDataFusionInstanceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfiip.ref.Append("id"))
}

// Name returns a reference to field name of google_data_fusion_instance_iam_policy.
func (dfiip dataDataFusionInstanceIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfiip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_data_fusion_instance_iam_policy.
func (dfiip dataDataFusionInstanceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dfiip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_data_fusion_instance_iam_policy.
func (dfiip dataDataFusionInstanceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dfiip.ref.Append("project"))
}

// Region returns a reference to field region of google_data_fusion_instance_iam_policy.
func (dfiip dataDataFusionInstanceIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dfiip.ref.Append("region"))
}
