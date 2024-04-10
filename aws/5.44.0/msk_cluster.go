// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	mskcluster "github.com/golingon/terraproviders/aws/5.44.0/mskcluster"
	"io"
)

// NewMskCluster creates a new instance of [MskCluster].
func NewMskCluster(name string, args MskClusterArgs) *MskCluster {
	return &MskCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MskCluster)(nil)

// MskCluster represents the Terraform resource aws_msk_cluster.
type MskCluster struct {
	Name      string
	Args      MskClusterArgs
	state     *mskClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MskCluster].
func (mc *MskCluster) Type() string {
	return "aws_msk_cluster"
}

// LocalName returns the local name for [MskCluster].
func (mc *MskCluster) LocalName() string {
	return mc.Name
}

// Configuration returns the configuration (args) for [MskCluster].
func (mc *MskCluster) Configuration() interface{} {
	return mc.Args
}

// DependOn is used for other resources to depend on [MskCluster].
func (mc *MskCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(mc)
}

// Dependencies returns the list of resources [MskCluster] depends_on.
func (mc *MskCluster) Dependencies() terra.Dependencies {
	return mc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MskCluster].
func (mc *MskCluster) LifecycleManagement() *terra.Lifecycle {
	return mc.Lifecycle
}

// Attributes returns the attributes for [MskCluster].
func (mc *MskCluster) Attributes() mskClusterAttributes {
	return mskClusterAttributes{ref: terra.ReferenceResource(mc)}
}

// ImportState imports the given attribute values into [MskCluster]'s state.
func (mc *MskCluster) ImportState(av io.Reader) error {
	mc.state = &mskClusterState{}
	if err := json.NewDecoder(av).Decode(mc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mc.Type(), mc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MskCluster] has state.
func (mc *MskCluster) State() (*mskClusterState, bool) {
	return mc.state, mc.state != nil
}

// StateMust returns the state for [MskCluster]. Panics if the state is nil.
func (mc *MskCluster) StateMust() *mskClusterState {
	if mc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mc.Type(), mc.LocalName()))
	}
	return mc.state
}

// MskClusterArgs contains the configurations for aws_msk_cluster.
type MskClusterArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// EnhancedMonitoring: string, optional
	EnhancedMonitoring terra.StringValue `hcl:"enhanced_monitoring,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KafkaVersion: string, required
	KafkaVersion terra.StringValue `hcl:"kafka_version,attr" validate:"required"`
	// NumberOfBrokerNodes: number, required
	NumberOfBrokerNodes terra.NumberValue `hcl:"number_of_broker_nodes,attr" validate:"required"`
	// StorageMode: string, optional
	StorageMode terra.StringValue `hcl:"storage_mode,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// BrokerNodeGroupInfo: required
	BrokerNodeGroupInfo *mskcluster.BrokerNodeGroupInfo `hcl:"broker_node_group_info,block" validate:"required"`
	// ClientAuthentication: optional
	ClientAuthentication *mskcluster.ClientAuthentication `hcl:"client_authentication,block"`
	// ConfigurationInfo: optional
	ConfigurationInfo *mskcluster.ConfigurationInfo `hcl:"configuration_info,block"`
	// EncryptionInfo: optional
	EncryptionInfo *mskcluster.EncryptionInfo `hcl:"encryption_info,block"`
	// LoggingInfo: optional
	LoggingInfo *mskcluster.LoggingInfo `hcl:"logging_info,block"`
	// OpenMonitoring: optional
	OpenMonitoring *mskcluster.OpenMonitoring `hcl:"open_monitoring,block"`
	// Timeouts: optional
	Timeouts *mskcluster.Timeouts `hcl:"timeouts,block"`
}
type mskClusterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_msk_cluster.
func (mc mskClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("arn"))
}

// BootstrapBrokers returns a reference to field bootstrap_brokers of aws_msk_cluster.
func (mc mskClusterAttributes) BootstrapBrokers() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers"))
}

// BootstrapBrokersPublicSaslIam returns a reference to field bootstrap_brokers_public_sasl_iam of aws_msk_cluster.
func (mc mskClusterAttributes) BootstrapBrokersPublicSaslIam() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers_public_sasl_iam"))
}

// BootstrapBrokersPublicSaslScram returns a reference to field bootstrap_brokers_public_sasl_scram of aws_msk_cluster.
func (mc mskClusterAttributes) BootstrapBrokersPublicSaslScram() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers_public_sasl_scram"))
}

// BootstrapBrokersPublicTls returns a reference to field bootstrap_brokers_public_tls of aws_msk_cluster.
func (mc mskClusterAttributes) BootstrapBrokersPublicTls() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers_public_tls"))
}

// BootstrapBrokersSaslIam returns a reference to field bootstrap_brokers_sasl_iam of aws_msk_cluster.
func (mc mskClusterAttributes) BootstrapBrokersSaslIam() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers_sasl_iam"))
}

// BootstrapBrokersSaslScram returns a reference to field bootstrap_brokers_sasl_scram of aws_msk_cluster.
func (mc mskClusterAttributes) BootstrapBrokersSaslScram() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers_sasl_scram"))
}

// BootstrapBrokersTls returns a reference to field bootstrap_brokers_tls of aws_msk_cluster.
func (mc mskClusterAttributes) BootstrapBrokersTls() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers_tls"))
}

// BootstrapBrokersVpcConnectivitySaslIam returns a reference to field bootstrap_brokers_vpc_connectivity_sasl_iam of aws_msk_cluster.
func (mc mskClusterAttributes) BootstrapBrokersVpcConnectivitySaslIam() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers_vpc_connectivity_sasl_iam"))
}

// BootstrapBrokersVpcConnectivitySaslScram returns a reference to field bootstrap_brokers_vpc_connectivity_sasl_scram of aws_msk_cluster.
func (mc mskClusterAttributes) BootstrapBrokersVpcConnectivitySaslScram() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers_vpc_connectivity_sasl_scram"))
}

// BootstrapBrokersVpcConnectivityTls returns a reference to field bootstrap_brokers_vpc_connectivity_tls of aws_msk_cluster.
func (mc mskClusterAttributes) BootstrapBrokersVpcConnectivityTls() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("bootstrap_brokers_vpc_connectivity_tls"))
}

// ClusterName returns a reference to field cluster_name of aws_msk_cluster.
func (mc mskClusterAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("cluster_name"))
}

// ClusterUuid returns a reference to field cluster_uuid of aws_msk_cluster.
func (mc mskClusterAttributes) ClusterUuid() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("cluster_uuid"))
}

// CurrentVersion returns a reference to field current_version of aws_msk_cluster.
func (mc mskClusterAttributes) CurrentVersion() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("current_version"))
}

// EnhancedMonitoring returns a reference to field enhanced_monitoring of aws_msk_cluster.
func (mc mskClusterAttributes) EnhancedMonitoring() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("enhanced_monitoring"))
}

// Id returns a reference to field id of aws_msk_cluster.
func (mc mskClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("id"))
}

// KafkaVersion returns a reference to field kafka_version of aws_msk_cluster.
func (mc mskClusterAttributes) KafkaVersion() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("kafka_version"))
}

// NumberOfBrokerNodes returns a reference to field number_of_broker_nodes of aws_msk_cluster.
func (mc mskClusterAttributes) NumberOfBrokerNodes() terra.NumberValue {
	return terra.ReferenceAsNumber(mc.ref.Append("number_of_broker_nodes"))
}

// StorageMode returns a reference to field storage_mode of aws_msk_cluster.
func (mc mskClusterAttributes) StorageMode() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("storage_mode"))
}

// Tags returns a reference to field tags of aws_msk_cluster.
func (mc mskClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_msk_cluster.
func (mc mskClusterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mc.ref.Append("tags_all"))
}

// ZookeeperConnectString returns a reference to field zookeeper_connect_string of aws_msk_cluster.
func (mc mskClusterAttributes) ZookeeperConnectString() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("zookeeper_connect_string"))
}

// ZookeeperConnectStringTls returns a reference to field zookeeper_connect_string_tls of aws_msk_cluster.
func (mc mskClusterAttributes) ZookeeperConnectStringTls() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("zookeeper_connect_string_tls"))
}

func (mc mskClusterAttributes) BrokerNodeGroupInfo() terra.ListValue[mskcluster.BrokerNodeGroupInfoAttributes] {
	return terra.ReferenceAsList[mskcluster.BrokerNodeGroupInfoAttributes](mc.ref.Append("broker_node_group_info"))
}

func (mc mskClusterAttributes) ClientAuthentication() terra.ListValue[mskcluster.ClientAuthenticationAttributes] {
	return terra.ReferenceAsList[mskcluster.ClientAuthenticationAttributes](mc.ref.Append("client_authentication"))
}

func (mc mskClusterAttributes) ConfigurationInfo() terra.ListValue[mskcluster.ConfigurationInfoAttributes] {
	return terra.ReferenceAsList[mskcluster.ConfigurationInfoAttributes](mc.ref.Append("configuration_info"))
}

func (mc mskClusterAttributes) EncryptionInfo() terra.ListValue[mskcluster.EncryptionInfoAttributes] {
	return terra.ReferenceAsList[mskcluster.EncryptionInfoAttributes](mc.ref.Append("encryption_info"))
}

func (mc mskClusterAttributes) LoggingInfo() terra.ListValue[mskcluster.LoggingInfoAttributes] {
	return terra.ReferenceAsList[mskcluster.LoggingInfoAttributes](mc.ref.Append("logging_info"))
}

func (mc mskClusterAttributes) OpenMonitoring() terra.ListValue[mskcluster.OpenMonitoringAttributes] {
	return terra.ReferenceAsList[mskcluster.OpenMonitoringAttributes](mc.ref.Append("open_monitoring"))
}

func (mc mskClusterAttributes) Timeouts() mskcluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mskcluster.TimeoutsAttributes](mc.ref.Append("timeouts"))
}

type mskClusterState struct {
	Arn                                      string                                 `json:"arn"`
	BootstrapBrokers                         string                                 `json:"bootstrap_brokers"`
	BootstrapBrokersPublicSaslIam            string                                 `json:"bootstrap_brokers_public_sasl_iam"`
	BootstrapBrokersPublicSaslScram          string                                 `json:"bootstrap_brokers_public_sasl_scram"`
	BootstrapBrokersPublicTls                string                                 `json:"bootstrap_brokers_public_tls"`
	BootstrapBrokersSaslIam                  string                                 `json:"bootstrap_brokers_sasl_iam"`
	BootstrapBrokersSaslScram                string                                 `json:"bootstrap_brokers_sasl_scram"`
	BootstrapBrokersTls                      string                                 `json:"bootstrap_brokers_tls"`
	BootstrapBrokersVpcConnectivitySaslIam   string                                 `json:"bootstrap_brokers_vpc_connectivity_sasl_iam"`
	BootstrapBrokersVpcConnectivitySaslScram string                                 `json:"bootstrap_brokers_vpc_connectivity_sasl_scram"`
	BootstrapBrokersVpcConnectivityTls       string                                 `json:"bootstrap_brokers_vpc_connectivity_tls"`
	ClusterName                              string                                 `json:"cluster_name"`
	ClusterUuid                              string                                 `json:"cluster_uuid"`
	CurrentVersion                           string                                 `json:"current_version"`
	EnhancedMonitoring                       string                                 `json:"enhanced_monitoring"`
	Id                                       string                                 `json:"id"`
	KafkaVersion                             string                                 `json:"kafka_version"`
	NumberOfBrokerNodes                      float64                                `json:"number_of_broker_nodes"`
	StorageMode                              string                                 `json:"storage_mode"`
	Tags                                     map[string]string                      `json:"tags"`
	TagsAll                                  map[string]string                      `json:"tags_all"`
	ZookeeperConnectString                   string                                 `json:"zookeeper_connect_string"`
	ZookeeperConnectStringTls                string                                 `json:"zookeeper_connect_string_tls"`
	BrokerNodeGroupInfo                      []mskcluster.BrokerNodeGroupInfoState  `json:"broker_node_group_info"`
	ClientAuthentication                     []mskcluster.ClientAuthenticationState `json:"client_authentication"`
	ConfigurationInfo                        []mskcluster.ConfigurationInfoState    `json:"configuration_info"`
	EncryptionInfo                           []mskcluster.EncryptionInfoState       `json:"encryption_info"`
	LoggingInfo                              []mskcluster.LoggingInfoState          `json:"logging_info"`
	OpenMonitoring                           []mskcluster.OpenMonitoringState       `json:"open_monitoring"`
	Timeouts                                 *mskcluster.TimeoutsState              `json:"timeouts"`
}
