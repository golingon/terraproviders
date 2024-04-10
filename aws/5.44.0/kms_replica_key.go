// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewKmsReplicaKey creates a new instance of [KmsReplicaKey].
func NewKmsReplicaKey(name string, args KmsReplicaKeyArgs) *KmsReplicaKey {
	return &KmsReplicaKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KmsReplicaKey)(nil)

// KmsReplicaKey represents the Terraform resource aws_kms_replica_key.
type KmsReplicaKey struct {
	Name      string
	Args      KmsReplicaKeyArgs
	state     *kmsReplicaKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KmsReplicaKey].
func (krk *KmsReplicaKey) Type() string {
	return "aws_kms_replica_key"
}

// LocalName returns the local name for [KmsReplicaKey].
func (krk *KmsReplicaKey) LocalName() string {
	return krk.Name
}

// Configuration returns the configuration (args) for [KmsReplicaKey].
func (krk *KmsReplicaKey) Configuration() interface{} {
	return krk.Args
}

// DependOn is used for other resources to depend on [KmsReplicaKey].
func (krk *KmsReplicaKey) DependOn() terra.Reference {
	return terra.ReferenceResource(krk)
}

// Dependencies returns the list of resources [KmsReplicaKey] depends_on.
func (krk *KmsReplicaKey) Dependencies() terra.Dependencies {
	return krk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KmsReplicaKey].
func (krk *KmsReplicaKey) LifecycleManagement() *terra.Lifecycle {
	return krk.Lifecycle
}

// Attributes returns the attributes for [KmsReplicaKey].
func (krk *KmsReplicaKey) Attributes() kmsReplicaKeyAttributes {
	return kmsReplicaKeyAttributes{ref: terra.ReferenceResource(krk)}
}

// ImportState imports the given attribute values into [KmsReplicaKey]'s state.
func (krk *KmsReplicaKey) ImportState(av io.Reader) error {
	krk.state = &kmsReplicaKeyState{}
	if err := json.NewDecoder(av).Decode(krk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", krk.Type(), krk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KmsReplicaKey] has state.
func (krk *KmsReplicaKey) State() (*kmsReplicaKeyState, bool) {
	return krk.state, krk.state != nil
}

// StateMust returns the state for [KmsReplicaKey]. Panics if the state is nil.
func (krk *KmsReplicaKey) StateMust() *kmsReplicaKeyState {
	if krk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", krk.Type(), krk.LocalName()))
	}
	return krk.state
}

// KmsReplicaKeyArgs contains the configurations for aws_kms_replica_key.
type KmsReplicaKeyArgs struct {
	// BypassPolicyLockoutSafetyCheck: bool, optional
	BypassPolicyLockoutSafetyCheck terra.BoolValue `hcl:"bypass_policy_lockout_safety_check,attr"`
	// DeletionWindowInDays: number, optional
	DeletionWindowInDays terra.NumberValue `hcl:"deletion_window_in_days,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Policy: string, optional
	Policy terra.StringValue `hcl:"policy,attr"`
	// PrimaryKeyArn: string, required
	PrimaryKeyArn terra.StringValue `hcl:"primary_key_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type kmsReplicaKeyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kms_replica_key.
func (krk kmsReplicaKeyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(krk.ref.Append("arn"))
}

// BypassPolicyLockoutSafetyCheck returns a reference to field bypass_policy_lockout_safety_check of aws_kms_replica_key.
func (krk kmsReplicaKeyAttributes) BypassPolicyLockoutSafetyCheck() terra.BoolValue {
	return terra.ReferenceAsBool(krk.ref.Append("bypass_policy_lockout_safety_check"))
}

// DeletionWindowInDays returns a reference to field deletion_window_in_days of aws_kms_replica_key.
func (krk kmsReplicaKeyAttributes) DeletionWindowInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(krk.ref.Append("deletion_window_in_days"))
}

// Description returns a reference to field description of aws_kms_replica_key.
func (krk kmsReplicaKeyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(krk.ref.Append("description"))
}

// Enabled returns a reference to field enabled of aws_kms_replica_key.
func (krk kmsReplicaKeyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(krk.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_kms_replica_key.
func (krk kmsReplicaKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(krk.ref.Append("id"))
}

// KeyId returns a reference to field key_id of aws_kms_replica_key.
func (krk kmsReplicaKeyAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(krk.ref.Append("key_id"))
}

// KeyRotationEnabled returns a reference to field key_rotation_enabled of aws_kms_replica_key.
func (krk kmsReplicaKeyAttributes) KeyRotationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(krk.ref.Append("key_rotation_enabled"))
}

// KeySpec returns a reference to field key_spec of aws_kms_replica_key.
func (krk kmsReplicaKeyAttributes) KeySpec() terra.StringValue {
	return terra.ReferenceAsString(krk.ref.Append("key_spec"))
}

// KeyUsage returns a reference to field key_usage of aws_kms_replica_key.
func (krk kmsReplicaKeyAttributes) KeyUsage() terra.StringValue {
	return terra.ReferenceAsString(krk.ref.Append("key_usage"))
}

// Policy returns a reference to field policy of aws_kms_replica_key.
func (krk kmsReplicaKeyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(krk.ref.Append("policy"))
}

// PrimaryKeyArn returns a reference to field primary_key_arn of aws_kms_replica_key.
func (krk kmsReplicaKeyAttributes) PrimaryKeyArn() terra.StringValue {
	return terra.ReferenceAsString(krk.ref.Append("primary_key_arn"))
}

// Tags returns a reference to field tags of aws_kms_replica_key.
func (krk kmsReplicaKeyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](krk.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_kms_replica_key.
func (krk kmsReplicaKeyAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](krk.ref.Append("tags_all"))
}

type kmsReplicaKeyState struct {
	Arn                            string            `json:"arn"`
	BypassPolicyLockoutSafetyCheck bool              `json:"bypass_policy_lockout_safety_check"`
	DeletionWindowInDays           float64           `json:"deletion_window_in_days"`
	Description                    string            `json:"description"`
	Enabled                        bool              `json:"enabled"`
	Id                             string            `json:"id"`
	KeyId                          string            `json:"key_id"`
	KeyRotationEnabled             bool              `json:"key_rotation_enabled"`
	KeySpec                        string            `json:"key_spec"`
	KeyUsage                       string            `json:"key_usage"`
	Policy                         string            `json:"policy"`
	PrimaryKeyArn                  string            `json:"primary_key_arn"`
	Tags                           map[string]string `json:"tags"`
	TagsAll                        map[string]string `json:"tags_all"`
}
