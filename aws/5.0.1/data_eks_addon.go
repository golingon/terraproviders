// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataEksAddon creates a new instance of [DataEksAddon].
func NewDataEksAddon(name string, args DataEksAddonArgs) *DataEksAddon {
	return &DataEksAddon{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEksAddon)(nil)

// DataEksAddon represents the Terraform data resource aws_eks_addon.
type DataEksAddon struct {
	Name string
	Args DataEksAddonArgs
}

// DataSource returns the Terraform object type for [DataEksAddon].
func (ea *DataEksAddon) DataSource() string {
	return "aws_eks_addon"
}

// LocalName returns the local name for [DataEksAddon].
func (ea *DataEksAddon) LocalName() string {
	return ea.Name
}

// Configuration returns the configuration (args) for [DataEksAddon].
func (ea *DataEksAddon) Configuration() interface{} {
	return ea.Args
}

// Attributes returns the attributes for [DataEksAddon].
func (ea *DataEksAddon) Attributes() dataEksAddonAttributes {
	return dataEksAddonAttributes{ref: terra.ReferenceDataResource(ea)}
}

// DataEksAddonArgs contains the configurations for aws_eks_addon.
type DataEksAddonArgs struct {
	// AddonName: string, required
	AddonName terra.StringValue `hcl:"addon_name,attr" validate:"required"`
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataEksAddonAttributes struct {
	ref terra.Reference
}

// AddonName returns a reference to field addon_name of aws_eks_addon.
func (ea dataEksAddonAttributes) AddonName() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("addon_name"))
}

// AddonVersion returns a reference to field addon_version of aws_eks_addon.
func (ea dataEksAddonAttributes) AddonVersion() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("addon_version"))
}

// Arn returns a reference to field arn of aws_eks_addon.
func (ea dataEksAddonAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("arn"))
}

// ClusterName returns a reference to field cluster_name of aws_eks_addon.
func (ea dataEksAddonAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("cluster_name"))
}

// ConfigurationValues returns a reference to field configuration_values of aws_eks_addon.
func (ea dataEksAddonAttributes) ConfigurationValues() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("configuration_values"))
}

// CreatedAt returns a reference to field created_at of aws_eks_addon.
func (ea dataEksAddonAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("created_at"))
}

// Id returns a reference to field id of aws_eks_addon.
func (ea dataEksAddonAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("id"))
}

// ModifiedAt returns a reference to field modified_at of aws_eks_addon.
func (ea dataEksAddonAttributes) ModifiedAt() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("modified_at"))
}

// ServiceAccountRoleArn returns a reference to field service_account_role_arn of aws_eks_addon.
func (ea dataEksAddonAttributes) ServiceAccountRoleArn() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("service_account_role_arn"))
}

// Tags returns a reference to field tags of aws_eks_addon.
func (ea dataEksAddonAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ea.ref.Append("tags"))
}