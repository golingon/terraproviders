// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataEksClusterAuth creates a new instance of [DataEksClusterAuth].
func NewDataEksClusterAuth(name string, args DataEksClusterAuthArgs) *DataEksClusterAuth {
	return &DataEksClusterAuth{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEksClusterAuth)(nil)

// DataEksClusterAuth represents the Terraform data resource aws_eks_cluster_auth.
type DataEksClusterAuth struct {
	Name string
	Args DataEksClusterAuthArgs
}

// DataSource returns the Terraform object type for [DataEksClusterAuth].
func (eca *DataEksClusterAuth) DataSource() string {
	return "aws_eks_cluster_auth"
}

// LocalName returns the local name for [DataEksClusterAuth].
func (eca *DataEksClusterAuth) LocalName() string {
	return eca.Name
}

// Configuration returns the configuration (args) for [DataEksClusterAuth].
func (eca *DataEksClusterAuth) Configuration() interface{} {
	return eca.Args
}

// Attributes returns the attributes for [DataEksClusterAuth].
func (eca *DataEksClusterAuth) Attributes() dataEksClusterAuthAttributes {
	return dataEksClusterAuthAttributes{ref: terra.ReferenceDataResource(eca)}
}

// DataEksClusterAuthArgs contains the configurations for aws_eks_cluster_auth.
type DataEksClusterAuthArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type dataEksClusterAuthAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_eks_cluster_auth.
func (eca dataEksClusterAuthAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eca.ref.Append("id"))
}

// Name returns a reference to field name of aws_eks_cluster_auth.
func (eca dataEksClusterAuthAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(eca.ref.Append("name"))
}

// Token returns a reference to field token of aws_eks_cluster_auth.
func (eca dataEksClusterAuthAttributes) Token() terra.StringValue {
	return terra.ReferenceAsString(eca.ref.Append("token"))
}
