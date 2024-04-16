// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_secretsmanager_secret

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_secretsmanager_secret.
type Resource struct {
	Name      string
	Args      Args
	state     *awsSecretsmanagerSecretState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ass *Resource) Type() string {
	return "aws_secretsmanager_secret"
}

// LocalName returns the local name for [Resource].
func (ass *Resource) LocalName() string {
	return ass.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ass *Resource) Configuration() interface{} {
	return ass.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ass *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ass)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ass *Resource) Dependencies() terra.Dependencies {
	return ass.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ass *Resource) LifecycleManagement() *terra.Lifecycle {
	return ass.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ass *Resource) Attributes() awsSecretsmanagerSecretAttributes {
	return awsSecretsmanagerSecretAttributes{ref: terra.ReferenceResource(ass)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ass *Resource) ImportState(state io.Reader) error {
	ass.state = &awsSecretsmanagerSecretState{}
	if err := json.NewDecoder(state).Decode(ass.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ass.Type(), ass.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ass *Resource) State() (*awsSecretsmanagerSecretState, bool) {
	return ass.state, ass.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ass *Resource) StateMust() *awsSecretsmanagerSecretState {
	if ass.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ass.Type(), ass.LocalName()))
	}
	return ass.state
}

// Args contains the configurations for aws_secretsmanager_secret.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ForceOverwriteReplicaSecret: bool, optional
	ForceOverwriteReplicaSecret terra.BoolValue `hcl:"force_overwrite_replica_secret,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// Policy: string, optional
	Policy terra.StringValue `hcl:"policy,attr"`
	// RecoveryWindowInDays: number, optional
	RecoveryWindowInDays terra.NumberValue `hcl:"recovery_window_in_days,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Replica: min=0
	Replica []Replica `hcl:"replica,block" validate:"min=0"`
}

type awsSecretsmanagerSecretAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_secretsmanager_secret.
func (ass awsSecretsmanagerSecretAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("arn"))
}

// Description returns a reference to field description of aws_secretsmanager_secret.
func (ass awsSecretsmanagerSecretAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("description"))
}

// ForceOverwriteReplicaSecret returns a reference to field force_overwrite_replica_secret of aws_secretsmanager_secret.
func (ass awsSecretsmanagerSecretAttributes) ForceOverwriteReplicaSecret() terra.BoolValue {
	return terra.ReferenceAsBool(ass.ref.Append("force_overwrite_replica_secret"))
}

// Id returns a reference to field id of aws_secretsmanager_secret.
func (ass awsSecretsmanagerSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_secretsmanager_secret.
func (ass awsSecretsmanagerSecretAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("kms_key_id"))
}

// Name returns a reference to field name of aws_secretsmanager_secret.
func (ass awsSecretsmanagerSecretAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_secretsmanager_secret.
func (ass awsSecretsmanagerSecretAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("name_prefix"))
}

// Policy returns a reference to field policy of aws_secretsmanager_secret.
func (ass awsSecretsmanagerSecretAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(ass.ref.Append("policy"))
}

// RecoveryWindowInDays returns a reference to field recovery_window_in_days of aws_secretsmanager_secret.
func (ass awsSecretsmanagerSecretAttributes) RecoveryWindowInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(ass.ref.Append("recovery_window_in_days"))
}

// Tags returns a reference to field tags of aws_secretsmanager_secret.
func (ass awsSecretsmanagerSecretAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ass.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_secretsmanager_secret.
func (ass awsSecretsmanagerSecretAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ass.ref.Append("tags_all"))
}

func (ass awsSecretsmanagerSecretAttributes) Replica() terra.SetValue[ReplicaAttributes] {
	return terra.ReferenceAsSet[ReplicaAttributes](ass.ref.Append("replica"))
}

type awsSecretsmanagerSecretState struct {
	Arn                         string            `json:"arn"`
	Description                 string            `json:"description"`
	ForceOverwriteReplicaSecret bool              `json:"force_overwrite_replica_secret"`
	Id                          string            `json:"id"`
	KmsKeyId                    string            `json:"kms_key_id"`
	Name                        string            `json:"name"`
	NamePrefix                  string            `json:"name_prefix"`
	Policy                      string            `json:"policy"`
	RecoveryWindowInDays        float64           `json:"recovery_window_in_days"`
	Tags                        map[string]string `json:"tags"`
	TagsAll                     map[string]string `json:"tags_all"`
	Replica                     []ReplicaState    `json:"replica"`
}
