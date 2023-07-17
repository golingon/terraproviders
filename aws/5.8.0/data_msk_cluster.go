// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataMskCluster creates a new instance of [DataMskCluster].
func NewDataMskCluster(name string, args DataMskClusterArgs) *DataMskCluster {
	return &DataMskCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMskCluster)(nil)

// DataMskCluster represents the Terraform data resource aws_msk_cluster.
type DataMskCluster struct {
	Name string
	Args DataMskClusterArgs
}

// DataSource returns the Terraform object type for [DataMskCluster].
func (mc *DataMskCluster) DataSource() string {
	return "aws_msk_cluster"
}

// LocalName returns the local name for [DataMskCluster].
func (mc *DataMskCluster) LocalName() string {
	return mc.Name
}

// Configuration returns the configuration (args) for [DataMskCluster].
func (mc *DataMskCluster) Configuration() interface{} {
	return mc.Args
}

// Attributes returns the attributes for [DataMskCluster].
func (mc *DataMskCluster) Attributes() dataMskClusterAttributes {
	return dataMskClusterAttributes{ref: terra.ReferenceDataResource(mc)}
}

// DataMskClusterArgs contains the configurations for aws_msk_cluster.
type DataMskClusterArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataMskClusterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_msk_cluster.
func (mc dataMskClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("arn"))
}

// BootstrapBrokers returns a reference to field bootstrap_brokers of aws_msk_cluster.
func (mc dataMskClusterAttributes) BootstrapBrokers() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers"))
}

// BootstrapBrokersPublicSaslIam returns a reference to field bootstrap_brokers_public_sasl_iam of aws_msk_cluster.
func (mc dataMskClusterAttributes) BootstrapBrokersPublicSaslIam() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers_public_sasl_iam"))
}

// BootstrapBrokersPublicSaslScram returns a reference to field bootstrap_brokers_public_sasl_scram of aws_msk_cluster.
func (mc dataMskClusterAttributes) BootstrapBrokersPublicSaslScram() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers_public_sasl_scram"))
}

// BootstrapBrokersPublicTls returns a reference to field bootstrap_brokers_public_tls of aws_msk_cluster.
func (mc dataMskClusterAttributes) BootstrapBrokersPublicTls() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers_public_tls"))
}

// BootstrapBrokersSaslIam returns a reference to field bootstrap_brokers_sasl_iam of aws_msk_cluster.
func (mc dataMskClusterAttributes) BootstrapBrokersSaslIam() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers_sasl_iam"))
}

// BootstrapBrokersSaslScram returns a reference to field bootstrap_brokers_sasl_scram of aws_msk_cluster.
func (mc dataMskClusterAttributes) BootstrapBrokersSaslScram() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers_sasl_scram"))
}

// BootstrapBrokersTls returns a reference to field bootstrap_brokers_tls of aws_msk_cluster.
func (mc dataMskClusterAttributes) BootstrapBrokersTls() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers_tls"))
}

// ClusterName returns a reference to field cluster_name of aws_msk_cluster.
func (mc dataMskClusterAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("cluster_name"))
}

// Id returns a reference to field id of aws_msk_cluster.
func (mc dataMskClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("id"))
}

// KafkaVersion returns a reference to field kafka_version of aws_msk_cluster.
func (mc dataMskClusterAttributes) KafkaVersion() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("kafka_version"))
}

// NumberOfBrokerNodes returns a reference to field number_of_broker_nodes of aws_msk_cluster.
func (mc dataMskClusterAttributes) NumberOfBrokerNodes() terra.NumberValue {
	return terra.ReferenceAsNumber(mc.ref.Append("number_of_broker_nodes"))
}

// Tags returns a reference to field tags of aws_msk_cluster.
func (mc dataMskClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mc.ref.Append("tags"))
}

// ZookeeperConnectString returns a reference to field zookeeper_connect_string of aws_msk_cluster.
func (mc dataMskClusterAttributes) ZookeeperConnectString() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("zookeeper_connect_string"))
}

// ZookeeperConnectStringTls returns a reference to field zookeeper_connect_string_tls of aws_msk_cluster.
func (mc dataMskClusterAttributes) ZookeeperConnectStringTls() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("zookeeper_connect_string_tls"))
}
