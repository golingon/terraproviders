// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_eks_cluster

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_eks_cluster.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aec *DataSource) DataSource() string {
	return "aws_eks_cluster"
}

// LocalName returns the local name for [DataSource].
func (aec *DataSource) LocalName() string {
	return aec.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aec *DataSource) Configuration() interface{} {
	return aec.Args
}

// Attributes returns the attributes for [DataSource].
func (aec *DataSource) Attributes() dataAwsEksClusterAttributes {
	return dataAwsEksClusterAttributes{ref: terra.ReferenceDataSource(aec)}
}

// DataArgs contains the configurations for aws_eks_cluster.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}

type dataAwsEksClusterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_eks_cluster.
func (aec dataAwsEksClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aec.ref.Append("arn"))
}

// ClusterId returns a reference to field cluster_id of aws_eks_cluster.
func (aec dataAwsEksClusterAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(aec.ref.Append("cluster_id"))
}

// CreatedAt returns a reference to field created_at of aws_eks_cluster.
func (aec dataAwsEksClusterAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(aec.ref.Append("created_at"))
}

// EnabledClusterLogTypes returns a reference to field enabled_cluster_log_types of aws_eks_cluster.
func (aec dataAwsEksClusterAttributes) EnabledClusterLogTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aec.ref.Append("enabled_cluster_log_types"))
}

// Endpoint returns a reference to field endpoint of aws_eks_cluster.
func (aec dataAwsEksClusterAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(aec.ref.Append("endpoint"))
}

// Id returns a reference to field id of aws_eks_cluster.
func (aec dataAwsEksClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aec.ref.Append("id"))
}

// Name returns a reference to field name of aws_eks_cluster.
func (aec dataAwsEksClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aec.ref.Append("name"))
}

// PlatformVersion returns a reference to field platform_version of aws_eks_cluster.
func (aec dataAwsEksClusterAttributes) PlatformVersion() terra.StringValue {
	return terra.ReferenceAsString(aec.ref.Append("platform_version"))
}

// RoleArn returns a reference to field role_arn of aws_eks_cluster.
func (aec dataAwsEksClusterAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(aec.ref.Append("role_arn"))
}

// Status returns a reference to field status of aws_eks_cluster.
func (aec dataAwsEksClusterAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(aec.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_eks_cluster.
func (aec dataAwsEksClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aec.ref.Append("tags"))
}

// Version returns a reference to field version of aws_eks_cluster.
func (aec dataAwsEksClusterAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(aec.ref.Append("version"))
}

func (aec dataAwsEksClusterAttributes) AccessConfig() terra.ListValue[DataAccessConfigAttributes] {
	return terra.ReferenceAsList[DataAccessConfigAttributes](aec.ref.Append("access_config"))
}

func (aec dataAwsEksClusterAttributes) CertificateAuthority() terra.ListValue[DataCertificateAuthorityAttributes] {
	return terra.ReferenceAsList[DataCertificateAuthorityAttributes](aec.ref.Append("certificate_authority"))
}

func (aec dataAwsEksClusterAttributes) Identity() terra.ListValue[DataIdentityAttributes] {
	return terra.ReferenceAsList[DataIdentityAttributes](aec.ref.Append("identity"))
}

func (aec dataAwsEksClusterAttributes) KubernetesNetworkConfig() terra.ListValue[DataKubernetesNetworkConfigAttributes] {
	return terra.ReferenceAsList[DataKubernetesNetworkConfigAttributes](aec.ref.Append("kubernetes_network_config"))
}

func (aec dataAwsEksClusterAttributes) OutpostConfig() terra.ListValue[DataOutpostConfigAttributes] {
	return terra.ReferenceAsList[DataOutpostConfigAttributes](aec.ref.Append("outpost_config"))
}

func (aec dataAwsEksClusterAttributes) VpcConfig() terra.ListValue[DataVpcConfigAttributes] {
	return terra.ReferenceAsList[DataVpcConfigAttributes](aec.ref.Append("vpc_config"))
}
