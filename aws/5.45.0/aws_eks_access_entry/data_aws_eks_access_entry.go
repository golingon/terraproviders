// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_eks_access_entry

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_eks_access_entry.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aeae *DataSource) DataSource() string {
	return "aws_eks_access_entry"
}

// LocalName returns the local name for [DataSource].
func (aeae *DataSource) LocalName() string {
	return aeae.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aeae *DataSource) Configuration() interface{} {
	return aeae.Args
}

// Attributes returns the attributes for [DataSource].
func (aeae *DataSource) Attributes() dataAwsEksAccessEntryAttributes {
	return dataAwsEksAccessEntryAttributes{ref: terra.ReferenceDataSource(aeae)}
}

// DataArgs contains the configurations for aws_eks_access_entry.
type DataArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PrincipalArn: string, required
	PrincipalArn terra.StringValue `hcl:"principal_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}

type dataAwsEksAccessEntryAttributes struct {
	ref terra.Reference
}

// AccessEntryArn returns a reference to field access_entry_arn of aws_eks_access_entry.
func (aeae dataAwsEksAccessEntryAttributes) AccessEntryArn() terra.StringValue {
	return terra.ReferenceAsString(aeae.ref.Append("access_entry_arn"))
}

// ClusterName returns a reference to field cluster_name of aws_eks_access_entry.
func (aeae dataAwsEksAccessEntryAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(aeae.ref.Append("cluster_name"))
}

// CreatedAt returns a reference to field created_at of aws_eks_access_entry.
func (aeae dataAwsEksAccessEntryAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(aeae.ref.Append("created_at"))
}

// Id returns a reference to field id of aws_eks_access_entry.
func (aeae dataAwsEksAccessEntryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aeae.ref.Append("id"))
}

// KubernetesGroups returns a reference to field kubernetes_groups of aws_eks_access_entry.
func (aeae dataAwsEksAccessEntryAttributes) KubernetesGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aeae.ref.Append("kubernetes_groups"))
}

// ModifiedAt returns a reference to field modified_at of aws_eks_access_entry.
func (aeae dataAwsEksAccessEntryAttributes) ModifiedAt() terra.StringValue {
	return terra.ReferenceAsString(aeae.ref.Append("modified_at"))
}

// PrincipalArn returns a reference to field principal_arn of aws_eks_access_entry.
func (aeae dataAwsEksAccessEntryAttributes) PrincipalArn() terra.StringValue {
	return terra.ReferenceAsString(aeae.ref.Append("principal_arn"))
}

// Tags returns a reference to field tags of aws_eks_access_entry.
func (aeae dataAwsEksAccessEntryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aeae.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_eks_access_entry.
func (aeae dataAwsEksAccessEntryAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aeae.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_eks_access_entry.
func (aeae dataAwsEksAccessEntryAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(aeae.ref.Append("type"))
}

// UserName returns a reference to field user_name of aws_eks_access_entry.
func (aeae dataAwsEksAccessEntryAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(aeae.ref.Append("user_name"))
}
