// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_iam_policy

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_iam_policy.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gip *DataSource) DataSource() string {
	return "google_iam_policy"
}

// LocalName returns the local name for [DataSource].
func (gip *DataSource) LocalName() string {
	return gip.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gip *DataSource) Configuration() interface{} {
	return gip.Args
}

// Attributes returns the attributes for [DataSource].
func (gip *DataSource) Attributes() dataGoogleIamPolicyAttributes {
	return dataGoogleIamPolicyAttributes{ref: terra.ReferenceDataSource(gip)}
}

// DataArgs contains the configurations for google_iam_policy.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// AuditConfig: min=0
	AuditConfig []DataAuditConfig `hcl:"audit_config,block" validate:"min=0"`
	// Binding: min=0
	Binding []DataBinding `hcl:"binding,block" validate:"min=0"`
}

type dataGoogleIamPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_iam_policy.
func (gip dataGoogleIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_iam_policy.
func (gip dataGoogleIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gip.ref.Append("policy_data"))
}

func (gip dataGoogleIamPolicyAttributes) AuditConfig() terra.SetValue[DataAuditConfigAttributes] {
	return terra.ReferenceAsSet[DataAuditConfigAttributes](gip.ref.Append("audit_config"))
}

func (gip dataGoogleIamPolicyAttributes) Binding() terra.SetValue[DataBindingAttributes] {
	return terra.ReferenceAsSet[DataBindingAttributes](gip.ref.Append("binding"))
}
