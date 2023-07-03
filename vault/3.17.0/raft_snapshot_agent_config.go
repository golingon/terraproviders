// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRaftSnapshotAgentConfig creates a new instance of [RaftSnapshotAgentConfig].
func NewRaftSnapshotAgentConfig(name string, args RaftSnapshotAgentConfigArgs) *RaftSnapshotAgentConfig {
	return &RaftSnapshotAgentConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RaftSnapshotAgentConfig)(nil)

// RaftSnapshotAgentConfig represents the Terraform resource vault_raft_snapshot_agent_config.
type RaftSnapshotAgentConfig struct {
	Name      string
	Args      RaftSnapshotAgentConfigArgs
	state     *raftSnapshotAgentConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RaftSnapshotAgentConfig].
func (rsac *RaftSnapshotAgentConfig) Type() string {
	return "vault_raft_snapshot_agent_config"
}

// LocalName returns the local name for [RaftSnapshotAgentConfig].
func (rsac *RaftSnapshotAgentConfig) LocalName() string {
	return rsac.Name
}

// Configuration returns the configuration (args) for [RaftSnapshotAgentConfig].
func (rsac *RaftSnapshotAgentConfig) Configuration() interface{} {
	return rsac.Args
}

// DependOn is used for other resources to depend on [RaftSnapshotAgentConfig].
func (rsac *RaftSnapshotAgentConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(rsac)
}

// Dependencies returns the list of resources [RaftSnapshotAgentConfig] depends_on.
func (rsac *RaftSnapshotAgentConfig) Dependencies() terra.Dependencies {
	return rsac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RaftSnapshotAgentConfig].
func (rsac *RaftSnapshotAgentConfig) LifecycleManagement() *terra.Lifecycle {
	return rsac.Lifecycle
}

// Attributes returns the attributes for [RaftSnapshotAgentConfig].
func (rsac *RaftSnapshotAgentConfig) Attributes() raftSnapshotAgentConfigAttributes {
	return raftSnapshotAgentConfigAttributes{ref: terra.ReferenceResource(rsac)}
}

// ImportState imports the given attribute values into [RaftSnapshotAgentConfig]'s state.
func (rsac *RaftSnapshotAgentConfig) ImportState(av io.Reader) error {
	rsac.state = &raftSnapshotAgentConfigState{}
	if err := json.NewDecoder(av).Decode(rsac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rsac.Type(), rsac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RaftSnapshotAgentConfig] has state.
func (rsac *RaftSnapshotAgentConfig) State() (*raftSnapshotAgentConfigState, bool) {
	return rsac.state, rsac.state != nil
}

// StateMust returns the state for [RaftSnapshotAgentConfig]. Panics if the state is nil.
func (rsac *RaftSnapshotAgentConfig) StateMust() *raftSnapshotAgentConfigState {
	if rsac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rsac.Type(), rsac.LocalName()))
	}
	return rsac.state
}

// RaftSnapshotAgentConfigArgs contains the configurations for vault_raft_snapshot_agent_config.
type RaftSnapshotAgentConfigArgs struct {
	// AwsAccessKeyId: string, optional
	AwsAccessKeyId terra.StringValue `hcl:"aws_access_key_id,attr"`
	// AwsS3Bucket: string, optional
	AwsS3Bucket terra.StringValue `hcl:"aws_s3_bucket,attr"`
	// AwsS3DisableTls: bool, optional
	AwsS3DisableTls terra.BoolValue `hcl:"aws_s3_disable_tls,attr"`
	// AwsS3EnableKms: bool, optional
	AwsS3EnableKms terra.BoolValue `hcl:"aws_s3_enable_kms,attr"`
	// AwsS3Endpoint: string, optional
	AwsS3Endpoint terra.StringValue `hcl:"aws_s3_endpoint,attr"`
	// AwsS3ForcePathStyle: bool, optional
	AwsS3ForcePathStyle terra.BoolValue `hcl:"aws_s3_force_path_style,attr"`
	// AwsS3KmsKey: string, optional
	AwsS3KmsKey terra.StringValue `hcl:"aws_s3_kms_key,attr"`
	// AwsS3Region: string, optional
	AwsS3Region terra.StringValue `hcl:"aws_s3_region,attr"`
	// AwsS3ServerSideEncryption: bool, optional
	AwsS3ServerSideEncryption terra.BoolValue `hcl:"aws_s3_server_side_encryption,attr"`
	// AwsSecretAccessKey: string, optional
	AwsSecretAccessKey terra.StringValue `hcl:"aws_secret_access_key,attr"`
	// AwsSessionToken: string, optional
	AwsSessionToken terra.StringValue `hcl:"aws_session_token,attr"`
	// AzureAccountKey: string, optional
	AzureAccountKey terra.StringValue `hcl:"azure_account_key,attr"`
	// AzureAccountName: string, optional
	AzureAccountName terra.StringValue `hcl:"azure_account_name,attr"`
	// AzureBlobEnvironment: string, optional
	AzureBlobEnvironment terra.StringValue `hcl:"azure_blob_environment,attr"`
	// AzureContainerName: string, optional
	AzureContainerName terra.StringValue `hcl:"azure_container_name,attr"`
	// AzureEndpoint: string, optional
	AzureEndpoint terra.StringValue `hcl:"azure_endpoint,attr"`
	// FilePrefix: string, optional
	FilePrefix terra.StringValue `hcl:"file_prefix,attr"`
	// GoogleDisableTls: bool, optional
	GoogleDisableTls terra.BoolValue `hcl:"google_disable_tls,attr"`
	// GoogleEndpoint: string, optional
	GoogleEndpoint terra.StringValue `hcl:"google_endpoint,attr"`
	// GoogleGcsBucket: string, optional
	GoogleGcsBucket terra.StringValue `hcl:"google_gcs_bucket,attr"`
	// GoogleServiceAccountKey: string, optional
	GoogleServiceAccountKey terra.StringValue `hcl:"google_service_account_key,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntervalSeconds: number, required
	IntervalSeconds terra.NumberValue `hcl:"interval_seconds,attr" validate:"required"`
	// LocalMaxSpace: number, optional
	LocalMaxSpace terra.NumberValue `hcl:"local_max_space,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// PathPrefix: string, required
	PathPrefix terra.StringValue `hcl:"path_prefix,attr" validate:"required"`
	// Retain: number, optional
	Retain terra.NumberValue `hcl:"retain,attr"`
	// StorageType: string, required
	StorageType terra.StringValue `hcl:"storage_type,attr" validate:"required"`
}
type raftSnapshotAgentConfigAttributes struct {
	ref terra.Reference
}

// AwsAccessKeyId returns a reference to field aws_access_key_id of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) AwsAccessKeyId() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("aws_access_key_id"))
}

// AwsS3Bucket returns a reference to field aws_s3_bucket of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) AwsS3Bucket() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("aws_s3_bucket"))
}

// AwsS3DisableTls returns a reference to field aws_s3_disable_tls of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) AwsS3DisableTls() terra.BoolValue {
	return terra.ReferenceAsBool(rsac.ref.Append("aws_s3_disable_tls"))
}

// AwsS3EnableKms returns a reference to field aws_s3_enable_kms of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) AwsS3EnableKms() terra.BoolValue {
	return terra.ReferenceAsBool(rsac.ref.Append("aws_s3_enable_kms"))
}

// AwsS3Endpoint returns a reference to field aws_s3_endpoint of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) AwsS3Endpoint() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("aws_s3_endpoint"))
}

// AwsS3ForcePathStyle returns a reference to field aws_s3_force_path_style of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) AwsS3ForcePathStyle() terra.BoolValue {
	return terra.ReferenceAsBool(rsac.ref.Append("aws_s3_force_path_style"))
}

// AwsS3KmsKey returns a reference to field aws_s3_kms_key of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) AwsS3KmsKey() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("aws_s3_kms_key"))
}

// AwsS3Region returns a reference to field aws_s3_region of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) AwsS3Region() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("aws_s3_region"))
}

// AwsS3ServerSideEncryption returns a reference to field aws_s3_server_side_encryption of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) AwsS3ServerSideEncryption() terra.BoolValue {
	return terra.ReferenceAsBool(rsac.ref.Append("aws_s3_server_side_encryption"))
}

// AwsSecretAccessKey returns a reference to field aws_secret_access_key of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) AwsSecretAccessKey() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("aws_secret_access_key"))
}

// AwsSessionToken returns a reference to field aws_session_token of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) AwsSessionToken() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("aws_session_token"))
}

// AzureAccountKey returns a reference to field azure_account_key of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) AzureAccountKey() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("azure_account_key"))
}

// AzureAccountName returns a reference to field azure_account_name of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) AzureAccountName() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("azure_account_name"))
}

// AzureBlobEnvironment returns a reference to field azure_blob_environment of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) AzureBlobEnvironment() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("azure_blob_environment"))
}

// AzureContainerName returns a reference to field azure_container_name of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) AzureContainerName() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("azure_container_name"))
}

// AzureEndpoint returns a reference to field azure_endpoint of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) AzureEndpoint() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("azure_endpoint"))
}

// FilePrefix returns a reference to field file_prefix of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) FilePrefix() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("file_prefix"))
}

// GoogleDisableTls returns a reference to field google_disable_tls of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) GoogleDisableTls() terra.BoolValue {
	return terra.ReferenceAsBool(rsac.ref.Append("google_disable_tls"))
}

// GoogleEndpoint returns a reference to field google_endpoint of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) GoogleEndpoint() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("google_endpoint"))
}

// GoogleGcsBucket returns a reference to field google_gcs_bucket of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) GoogleGcsBucket() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("google_gcs_bucket"))
}

// GoogleServiceAccountKey returns a reference to field google_service_account_key of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) GoogleServiceAccountKey() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("google_service_account_key"))
}

// Id returns a reference to field id of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("id"))
}

// IntervalSeconds returns a reference to field interval_seconds of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) IntervalSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(rsac.ref.Append("interval_seconds"))
}

// LocalMaxSpace returns a reference to field local_max_space of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) LocalMaxSpace() terra.NumberValue {
	return terra.ReferenceAsNumber(rsac.ref.Append("local_max_space"))
}

// Name returns a reference to field name of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("namespace"))
}

// PathPrefix returns a reference to field path_prefix of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) PathPrefix() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("path_prefix"))
}

// Retain returns a reference to field retain of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) Retain() terra.NumberValue {
	return terra.ReferenceAsNumber(rsac.ref.Append("retain"))
}

// StorageType returns a reference to field storage_type of vault_raft_snapshot_agent_config.
func (rsac raftSnapshotAgentConfigAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(rsac.ref.Append("storage_type"))
}

type raftSnapshotAgentConfigState struct {
	AwsAccessKeyId            string  `json:"aws_access_key_id"`
	AwsS3Bucket               string  `json:"aws_s3_bucket"`
	AwsS3DisableTls           bool    `json:"aws_s3_disable_tls"`
	AwsS3EnableKms            bool    `json:"aws_s3_enable_kms"`
	AwsS3Endpoint             string  `json:"aws_s3_endpoint"`
	AwsS3ForcePathStyle       bool    `json:"aws_s3_force_path_style"`
	AwsS3KmsKey               string  `json:"aws_s3_kms_key"`
	AwsS3Region               string  `json:"aws_s3_region"`
	AwsS3ServerSideEncryption bool    `json:"aws_s3_server_side_encryption"`
	AwsSecretAccessKey        string  `json:"aws_secret_access_key"`
	AwsSessionToken           string  `json:"aws_session_token"`
	AzureAccountKey           string  `json:"azure_account_key"`
	AzureAccountName          string  `json:"azure_account_name"`
	AzureBlobEnvironment      string  `json:"azure_blob_environment"`
	AzureContainerName        string  `json:"azure_container_name"`
	AzureEndpoint             string  `json:"azure_endpoint"`
	FilePrefix                string  `json:"file_prefix"`
	GoogleDisableTls          bool    `json:"google_disable_tls"`
	GoogleEndpoint            string  `json:"google_endpoint"`
	GoogleGcsBucket           string  `json:"google_gcs_bucket"`
	GoogleServiceAccountKey   string  `json:"google_service_account_key"`
	Id                        string  `json:"id"`
	IntervalSeconds           float64 `json:"interval_seconds"`
	LocalMaxSpace             float64 `json:"local_max_space"`
	Name                      string  `json:"name"`
	Namespace                 string  `json:"namespace"`
	PathPrefix                string  `json:"path_prefix"`
	Retain                    float64 `json:"retain"`
	StorageType               string  `json:"storage_type"`
}
