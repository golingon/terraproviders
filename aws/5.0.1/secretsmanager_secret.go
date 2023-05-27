// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	secretsmanagersecret "github.com/golingon/terraproviders/aws/5.0.1/secretsmanagersecret"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSecretsmanagerSecret creates a new instance of [SecretsmanagerSecret].
func NewSecretsmanagerSecret(name string, args SecretsmanagerSecretArgs) *SecretsmanagerSecret {
	return &SecretsmanagerSecret{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecretsmanagerSecret)(nil)

// SecretsmanagerSecret represents the Terraform resource aws_secretsmanager_secret.
type SecretsmanagerSecret struct {
	Name      string
	Args      SecretsmanagerSecretArgs
	state     *secretsmanagerSecretState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecretsmanagerSecret].
func (ss *SecretsmanagerSecret) Type() string {
	return "aws_secretsmanager_secret"
}

// LocalName returns the local name for [SecretsmanagerSecret].
func (ss *SecretsmanagerSecret) LocalName() string {
	return ss.Name
}

// Configuration returns the configuration (args) for [SecretsmanagerSecret].
func (ss *SecretsmanagerSecret) Configuration() interface{} {
	return ss.Args
}

// DependOn is used for other resources to depend on [SecretsmanagerSecret].
func (ss *SecretsmanagerSecret) DependOn() terra.Reference {
	return terra.ReferenceResource(ss)
}

// Dependencies returns the list of resources [SecretsmanagerSecret] depends_on.
func (ss *SecretsmanagerSecret) Dependencies() terra.Dependencies {
	return ss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecretsmanagerSecret].
func (ss *SecretsmanagerSecret) LifecycleManagement() *terra.Lifecycle {
	return ss.Lifecycle
}

// Attributes returns the attributes for [SecretsmanagerSecret].
func (ss *SecretsmanagerSecret) Attributes() secretsmanagerSecretAttributes {
	return secretsmanagerSecretAttributes{ref: terra.ReferenceResource(ss)}
}

// ImportState imports the given attribute values into [SecretsmanagerSecret]'s state.
func (ss *SecretsmanagerSecret) ImportState(av io.Reader) error {
	ss.state = &secretsmanagerSecretState{}
	if err := json.NewDecoder(av).Decode(ss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ss.Type(), ss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecretsmanagerSecret] has state.
func (ss *SecretsmanagerSecret) State() (*secretsmanagerSecretState, bool) {
	return ss.state, ss.state != nil
}

// StateMust returns the state for [SecretsmanagerSecret]. Panics if the state is nil.
func (ss *SecretsmanagerSecret) StateMust() *secretsmanagerSecretState {
	if ss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ss.Type(), ss.LocalName()))
	}
	return ss.state
}

// SecretsmanagerSecretArgs contains the configurations for aws_secretsmanager_secret.
type SecretsmanagerSecretArgs struct {
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
	Replica []secretsmanagersecret.Replica `hcl:"replica,block" validate:"min=0"`
}
type secretsmanagerSecretAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_secretsmanager_secret.
func (ss secretsmanagerSecretAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("arn"))
}

// Description returns a reference to field description of aws_secretsmanager_secret.
func (ss secretsmanagerSecretAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("description"))
}

// ForceOverwriteReplicaSecret returns a reference to field force_overwrite_replica_secret of aws_secretsmanager_secret.
func (ss secretsmanagerSecretAttributes) ForceOverwriteReplicaSecret() terra.BoolValue {
	return terra.ReferenceAsBool(ss.ref.Append("force_overwrite_replica_secret"))
}

// Id returns a reference to field id of aws_secretsmanager_secret.
func (ss secretsmanagerSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_secretsmanager_secret.
func (ss secretsmanagerSecretAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("kms_key_id"))
}

// Name returns a reference to field name of aws_secretsmanager_secret.
func (ss secretsmanagerSecretAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_secretsmanager_secret.
func (ss secretsmanagerSecretAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name_prefix"))
}

// Policy returns a reference to field policy of aws_secretsmanager_secret.
func (ss secretsmanagerSecretAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("policy"))
}

// RecoveryWindowInDays returns a reference to field recovery_window_in_days of aws_secretsmanager_secret.
func (ss secretsmanagerSecretAttributes) RecoveryWindowInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(ss.ref.Append("recovery_window_in_days"))
}

// Tags returns a reference to field tags of aws_secretsmanager_secret.
func (ss secretsmanagerSecretAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ss.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_secretsmanager_secret.
func (ss secretsmanagerSecretAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ss.ref.Append("tags_all"))
}

func (ss secretsmanagerSecretAttributes) Replica() terra.SetValue[secretsmanagersecret.ReplicaAttributes] {
	return terra.ReferenceAsSet[secretsmanagersecret.ReplicaAttributes](ss.ref.Append("replica"))
}

type secretsmanagerSecretState struct {
	Arn                         string                              `json:"arn"`
	Description                 string                              `json:"description"`
	ForceOverwriteReplicaSecret bool                                `json:"force_overwrite_replica_secret"`
	Id                          string                              `json:"id"`
	KmsKeyId                    string                              `json:"kms_key_id"`
	Name                        string                              `json:"name"`
	NamePrefix                  string                              `json:"name_prefix"`
	Policy                      string                              `json:"policy"`
	RecoveryWindowInDays        float64                             `json:"recovery_window_in_days"`
	Tags                        map[string]string                   `json:"tags"`
	TagsAll                     map[string]string                   `json:"tags_all"`
	Replica                     []secretsmanagersecret.ReplicaState `json:"replica"`
}
