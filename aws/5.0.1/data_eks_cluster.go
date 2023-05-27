// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataekscluster "github.com/golingon/terraproviders/aws/5.0.1/dataekscluster"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEksCluster creates a new instance of [DataEksCluster].
func NewDataEksCluster(name string, args DataEksClusterArgs) *DataEksCluster {
	return &DataEksCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEksCluster)(nil)

// DataEksCluster represents the Terraform data resource aws_eks_cluster.
type DataEksCluster struct {
	Name string
	Args DataEksClusterArgs
}

// DataSource returns the Terraform object type for [DataEksCluster].
func (ec *DataEksCluster) DataSource() string {
	return "aws_eks_cluster"
}

// LocalName returns the local name for [DataEksCluster].
func (ec *DataEksCluster) LocalName() string {
	return ec.Name
}

// Configuration returns the configuration (args) for [DataEksCluster].
func (ec *DataEksCluster) Configuration() interface{} {
	return ec.Args
}

// Attributes returns the attributes for [DataEksCluster].
func (ec *DataEksCluster) Attributes() dataEksClusterAttributes {
	return dataEksClusterAttributes{ref: terra.ReferenceDataResource(ec)}
}

// DataEksClusterArgs contains the configurations for aws_eks_cluster.
type DataEksClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// CertificateAuthority: min=0
	CertificateAuthority []dataekscluster.CertificateAuthority `hcl:"certificate_authority,block" validate:"min=0"`
	// Identity: min=0
	Identity []dataekscluster.Identity `hcl:"identity,block" validate:"min=0"`
	// KubernetesNetworkConfig: min=0
	KubernetesNetworkConfig []dataekscluster.KubernetesNetworkConfig `hcl:"kubernetes_network_config,block" validate:"min=0"`
	// OutpostConfig: min=0
	OutpostConfig []dataekscluster.OutpostConfig `hcl:"outpost_config,block" validate:"min=0"`
	// VpcConfig: min=0
	VpcConfig []dataekscluster.VpcConfig `hcl:"vpc_config,block" validate:"min=0"`
}
type dataEksClusterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_eks_cluster.
func (ec dataEksClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("arn"))
}

// ClusterId returns a reference to field cluster_id of aws_eks_cluster.
func (ec dataEksClusterAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("cluster_id"))
}

// CreatedAt returns a reference to field created_at of aws_eks_cluster.
func (ec dataEksClusterAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("created_at"))
}

// EnabledClusterLogTypes returns a reference to field enabled_cluster_log_types of aws_eks_cluster.
func (ec dataEksClusterAttributes) EnabledClusterLogTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ec.ref.Append("enabled_cluster_log_types"))
}

// Endpoint returns a reference to field endpoint of aws_eks_cluster.
func (ec dataEksClusterAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("endpoint"))
}

// Id returns a reference to field id of aws_eks_cluster.
func (ec dataEksClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("id"))
}

// Name returns a reference to field name of aws_eks_cluster.
func (ec dataEksClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("name"))
}

// PlatformVersion returns a reference to field platform_version of aws_eks_cluster.
func (ec dataEksClusterAttributes) PlatformVersion() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("platform_version"))
}

// RoleArn returns a reference to field role_arn of aws_eks_cluster.
func (ec dataEksClusterAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("role_arn"))
}

// Status returns a reference to field status of aws_eks_cluster.
func (ec dataEksClusterAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_eks_cluster.
func (ec dataEksClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ec.ref.Append("tags"))
}

// Version returns a reference to field version of aws_eks_cluster.
func (ec dataEksClusterAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("version"))
}

func (ec dataEksClusterAttributes) CertificateAuthority() terra.ListValue[dataekscluster.CertificateAuthorityAttributes] {
	return terra.ReferenceAsList[dataekscluster.CertificateAuthorityAttributes](ec.ref.Append("certificate_authority"))
}

func (ec dataEksClusterAttributes) Identity() terra.ListValue[dataekscluster.IdentityAttributes] {
	return terra.ReferenceAsList[dataekscluster.IdentityAttributes](ec.ref.Append("identity"))
}

func (ec dataEksClusterAttributes) KubernetesNetworkConfig() terra.ListValue[dataekscluster.KubernetesNetworkConfigAttributes] {
	return terra.ReferenceAsList[dataekscluster.KubernetesNetworkConfigAttributes](ec.ref.Append("kubernetes_network_config"))
}

func (ec dataEksClusterAttributes) OutpostConfig() terra.ListValue[dataekscluster.OutpostConfigAttributes] {
	return terra.ReferenceAsList[dataekscluster.OutpostConfigAttributes](ec.ref.Append("outpost_config"))
}

func (ec dataEksClusterAttributes) VpcConfig() terra.ListValue[dataekscluster.VpcConfigAttributes] {
	return terra.ReferenceAsList[dataekscluster.VpcConfigAttributes](ec.ref.Append("vpc_config"))
}
