// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_msk_cluster

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_msk_cluster.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (amc *DataSource) DataSource() string {
	return "aws_msk_cluster"
}

// LocalName returns the local name for [DataSource].
func (amc *DataSource) LocalName() string {
	return amc.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (amc *DataSource) Configuration() interface{} {
	return amc.Args
}

// Attributes returns the attributes for [DataSource].
func (amc *DataSource) Attributes() dataAwsMskClusterAttributes {
	return dataAwsMskClusterAttributes{ref: terra.ReferenceDataSource(amc)}
}

// DataArgs contains the configurations for aws_msk_cluster.
type DataArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}

type dataAwsMskClusterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_msk_cluster.
func (amc dataAwsMskClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("arn"))
}

// BootstrapBrokers returns a reference to field bootstrap_brokers of aws_msk_cluster.
func (amc dataAwsMskClusterAttributes) BootstrapBrokers() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("bootstrap_brokers"))
}

// BootstrapBrokersPublicSaslIam returns a reference to field bootstrap_brokers_public_sasl_iam of aws_msk_cluster.
func (amc dataAwsMskClusterAttributes) BootstrapBrokersPublicSaslIam() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("bootstrap_brokers_public_sasl_iam"))
}

// BootstrapBrokersPublicSaslScram returns a reference to field bootstrap_brokers_public_sasl_scram of aws_msk_cluster.
func (amc dataAwsMskClusterAttributes) BootstrapBrokersPublicSaslScram() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("bootstrap_brokers_public_sasl_scram"))
}

// BootstrapBrokersPublicTls returns a reference to field bootstrap_brokers_public_tls of aws_msk_cluster.
func (amc dataAwsMskClusterAttributes) BootstrapBrokersPublicTls() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("bootstrap_brokers_public_tls"))
}

// BootstrapBrokersSaslIam returns a reference to field bootstrap_brokers_sasl_iam of aws_msk_cluster.
func (amc dataAwsMskClusterAttributes) BootstrapBrokersSaslIam() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("bootstrap_brokers_sasl_iam"))
}

// BootstrapBrokersSaslScram returns a reference to field bootstrap_brokers_sasl_scram of aws_msk_cluster.
func (amc dataAwsMskClusterAttributes) BootstrapBrokersSaslScram() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("bootstrap_brokers_sasl_scram"))
}

// BootstrapBrokersTls returns a reference to field bootstrap_brokers_tls of aws_msk_cluster.
func (amc dataAwsMskClusterAttributes) BootstrapBrokersTls() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("bootstrap_brokers_tls"))
}

// ClusterName returns a reference to field cluster_name of aws_msk_cluster.
func (amc dataAwsMskClusterAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("cluster_name"))
}

// ClusterUuid returns a reference to field cluster_uuid of aws_msk_cluster.
func (amc dataAwsMskClusterAttributes) ClusterUuid() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("cluster_uuid"))
}

// Id returns a reference to field id of aws_msk_cluster.
func (amc dataAwsMskClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("id"))
}

// KafkaVersion returns a reference to field kafka_version of aws_msk_cluster.
func (amc dataAwsMskClusterAttributes) KafkaVersion() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("kafka_version"))
}

// NumberOfBrokerNodes returns a reference to field number_of_broker_nodes of aws_msk_cluster.
func (amc dataAwsMskClusterAttributes) NumberOfBrokerNodes() terra.NumberValue {
	return terra.ReferenceAsNumber(amc.ref.Append("number_of_broker_nodes"))
}

// Tags returns a reference to field tags of aws_msk_cluster.
func (amc dataAwsMskClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](amc.ref.Append("tags"))
}

// ZookeeperConnectString returns a reference to field zookeeper_connect_string of aws_msk_cluster.
func (amc dataAwsMskClusterAttributes) ZookeeperConnectString() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("zookeeper_connect_string"))
}

// ZookeeperConnectStringTls returns a reference to field zookeeper_connect_string_tls of aws_msk_cluster.
func (amc dataAwsMskClusterAttributes) ZookeeperConnectStringTls() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("zookeeper_connect_string_tls"))
}
