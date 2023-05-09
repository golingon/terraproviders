// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	dataiampolicy "github.com/golingon/terraproviders/google/4.63.1/dataiampolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataIamPolicy creates a new instance of [DataIamPolicy].
func NewDataIamPolicy(name string, args DataIamPolicyArgs) *DataIamPolicy {
	return &DataIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamPolicy)(nil)

// DataIamPolicy represents the Terraform data resource google_iam_policy.
type DataIamPolicy struct {
	Name string
	Args DataIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataIamPolicy].
func (ip *DataIamPolicy) DataSource() string {
	return "google_iam_policy"
}

// LocalName returns the local name for [DataIamPolicy].
func (ip *DataIamPolicy) LocalName() string {
	return ip.Name
}

// Configuration returns the configuration (args) for [DataIamPolicy].
func (ip *DataIamPolicy) Configuration() interface{} {
	return ip.Args
}

// Attributes returns the attributes for [DataIamPolicy].
func (ip *DataIamPolicy) Attributes() dataIamPolicyAttributes {
	return dataIamPolicyAttributes{ref: terra.ReferenceDataResource(ip)}
}

// DataIamPolicyArgs contains the configurations for google_iam_policy.
type DataIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// AuditConfig: min=0
	AuditConfig []dataiampolicy.AuditConfig `hcl:"audit_config,block" validate:"min=0"`
	// Binding: min=0
	Binding []dataiampolicy.Binding `hcl:"binding,block" validate:"min=0"`
}
type dataIamPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_iam_policy.
func (ip dataIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_iam_policy.
func (ip dataIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("policy_data"))
}

func (ip dataIamPolicyAttributes) AuditConfig() terra.SetValue[dataiampolicy.AuditConfigAttributes] {
	return terra.ReferenceAsSet[dataiampolicy.AuditConfigAttributes](ip.ref.Append("audit_config"))
}

func (ip dataIamPolicyAttributes) Binding() terra.SetValue[dataiampolicy.BindingAttributes] {
	return terra.ReferenceAsSet[dataiampolicy.BindingAttributes](ip.ref.Append("binding"))
}
