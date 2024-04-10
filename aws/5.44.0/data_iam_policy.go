// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataIamPolicy creates a new instance of [DataIamPolicy].
func NewDataIamPolicy(name string, args DataIamPolicyArgs) *DataIamPolicy {
	return &DataIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamPolicy)(nil)

// DataIamPolicy represents the Terraform data resource aws_iam_policy.
type DataIamPolicy struct {
	Name string
	Args DataIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataIamPolicy].
func (ip *DataIamPolicy) DataSource() string {
	return "aws_iam_policy"
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

// DataIamPolicyArgs contains the configurations for aws_iam_policy.
type DataIamPolicyArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// PathPrefix: string, optional
	PathPrefix terra.StringValue `hcl:"path_prefix,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataIamPolicyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iam_policy.
func (ip dataIamPolicyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("arn"))
}

// Description returns a reference to field description of aws_iam_policy.
func (ip dataIamPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("description"))
}

// Id returns a reference to field id of aws_iam_policy.
func (ip dataIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("id"))
}

// Name returns a reference to field name of aws_iam_policy.
func (ip dataIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("name"))
}

// Path returns a reference to field path of aws_iam_policy.
func (ip dataIamPolicyAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("path"))
}

// PathPrefix returns a reference to field path_prefix of aws_iam_policy.
func (ip dataIamPolicyAttributes) PathPrefix() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("path_prefix"))
}

// Policy returns a reference to field policy of aws_iam_policy.
func (ip dataIamPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("policy"))
}

// PolicyId returns a reference to field policy_id of aws_iam_policy.
func (ip dataIamPolicyAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("policy_id"))
}

// Tags returns a reference to field tags of aws_iam_policy.
func (ip dataIamPolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ip.ref.Append("tags"))
}
