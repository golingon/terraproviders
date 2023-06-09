// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datamskbrokernodes "github.com/golingon/terraproviders/aws/4.66.1/datamskbrokernodes"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMskBrokerNodes creates a new instance of [DataMskBrokerNodes].
func NewDataMskBrokerNodes(name string, args DataMskBrokerNodesArgs) *DataMskBrokerNodes {
	return &DataMskBrokerNodes{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMskBrokerNodes)(nil)

// DataMskBrokerNodes represents the Terraform data resource aws_msk_broker_nodes.
type DataMskBrokerNodes struct {
	Name string
	Args DataMskBrokerNodesArgs
}

// DataSource returns the Terraform object type for [DataMskBrokerNodes].
func (mbn *DataMskBrokerNodes) DataSource() string {
	return "aws_msk_broker_nodes"
}

// LocalName returns the local name for [DataMskBrokerNodes].
func (mbn *DataMskBrokerNodes) LocalName() string {
	return mbn.Name
}

// Configuration returns the configuration (args) for [DataMskBrokerNodes].
func (mbn *DataMskBrokerNodes) Configuration() interface{} {
	return mbn.Args
}

// Attributes returns the attributes for [DataMskBrokerNodes].
func (mbn *DataMskBrokerNodes) Attributes() dataMskBrokerNodesAttributes {
	return dataMskBrokerNodesAttributes{ref: terra.ReferenceDataResource(mbn)}
}

// DataMskBrokerNodesArgs contains the configurations for aws_msk_broker_nodes.
type DataMskBrokerNodesArgs struct {
	// ClusterArn: string, required
	ClusterArn terra.StringValue `hcl:"cluster_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NodeInfoList: min=0
	NodeInfoList []datamskbrokernodes.NodeInfoList `hcl:"node_info_list,block" validate:"min=0"`
}
type dataMskBrokerNodesAttributes struct {
	ref terra.Reference
}

// ClusterArn returns a reference to field cluster_arn of aws_msk_broker_nodes.
func (mbn dataMskBrokerNodesAttributes) ClusterArn() terra.StringValue {
	return terra.ReferenceAsString(mbn.ref.Append("cluster_arn"))
}

// Id returns a reference to field id of aws_msk_broker_nodes.
func (mbn dataMskBrokerNodesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mbn.ref.Append("id"))
}

func (mbn dataMskBrokerNodesAttributes) NodeInfoList() terra.ListValue[datamskbrokernodes.NodeInfoListAttributes] {
	return terra.ReferenceAsList[datamskbrokernodes.NodeInfoListAttributes](mbn.ref.Append("node_info_list"))
}
