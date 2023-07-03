// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datacloudhsmv2cluster "github.com/golingon/terraproviders/aws/5.6.2/datacloudhsmv2cluster"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCloudhsmV2Cluster creates a new instance of [DataCloudhsmV2Cluster].
func NewDataCloudhsmV2Cluster(name string, args DataCloudhsmV2ClusterArgs) *DataCloudhsmV2Cluster {
	return &DataCloudhsmV2Cluster{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudhsmV2Cluster)(nil)

// DataCloudhsmV2Cluster represents the Terraform data resource aws_cloudhsm_v2_cluster.
type DataCloudhsmV2Cluster struct {
	Name string
	Args DataCloudhsmV2ClusterArgs
}

// DataSource returns the Terraform object type for [DataCloudhsmV2Cluster].
func (cvc *DataCloudhsmV2Cluster) DataSource() string {
	return "aws_cloudhsm_v2_cluster"
}

// LocalName returns the local name for [DataCloudhsmV2Cluster].
func (cvc *DataCloudhsmV2Cluster) LocalName() string {
	return cvc.Name
}

// Configuration returns the configuration (args) for [DataCloudhsmV2Cluster].
func (cvc *DataCloudhsmV2Cluster) Configuration() interface{} {
	return cvc.Args
}

// Attributes returns the attributes for [DataCloudhsmV2Cluster].
func (cvc *DataCloudhsmV2Cluster) Attributes() dataCloudhsmV2ClusterAttributes {
	return dataCloudhsmV2ClusterAttributes{ref: terra.ReferenceDataResource(cvc)}
}

// DataCloudhsmV2ClusterArgs contains the configurations for aws_cloudhsm_v2_cluster.
type DataCloudhsmV2ClusterArgs struct {
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
	// ClusterState: string, optional
	ClusterState terra.StringValue `hcl:"cluster_state,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ClusterCertificates: min=0
	ClusterCertificates []datacloudhsmv2cluster.ClusterCertificates `hcl:"cluster_certificates,block" validate:"min=0"`
}
type dataCloudhsmV2ClusterAttributes struct {
	ref terra.Reference
}

// ClusterId returns a reference to field cluster_id of aws_cloudhsm_v2_cluster.
func (cvc dataCloudhsmV2ClusterAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(cvc.ref.Append("cluster_id"))
}

// ClusterState returns a reference to field cluster_state of aws_cloudhsm_v2_cluster.
func (cvc dataCloudhsmV2ClusterAttributes) ClusterState() terra.StringValue {
	return terra.ReferenceAsString(cvc.ref.Append("cluster_state"))
}

// Id returns a reference to field id of aws_cloudhsm_v2_cluster.
func (cvc dataCloudhsmV2ClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cvc.ref.Append("id"))
}

// SecurityGroupId returns a reference to field security_group_id of aws_cloudhsm_v2_cluster.
func (cvc dataCloudhsmV2ClusterAttributes) SecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(cvc.ref.Append("security_group_id"))
}

// SubnetIds returns a reference to field subnet_ids of aws_cloudhsm_v2_cluster.
func (cvc dataCloudhsmV2ClusterAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cvc.ref.Append("subnet_ids"))
}

// VpcId returns a reference to field vpc_id of aws_cloudhsm_v2_cluster.
func (cvc dataCloudhsmV2ClusterAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(cvc.ref.Append("vpc_id"))
}

func (cvc dataCloudhsmV2ClusterAttributes) ClusterCertificates() terra.ListValue[datacloudhsmv2cluster.ClusterCertificatesAttributes] {
	return terra.ReferenceAsList[datacloudhsmv2cluster.ClusterCertificatesAttributes](cvc.ref.Append("cluster_certificates"))
}
