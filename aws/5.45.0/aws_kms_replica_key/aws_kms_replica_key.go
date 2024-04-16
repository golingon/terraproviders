// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_kms_replica_key

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

// Resource represents the Terraform resource aws_kms_replica_key.
type Resource struct {
	Name      string
	Args      Args
	state     *awsKmsReplicaKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (akrk *Resource) Type() string {
	return "aws_kms_replica_key"
}

// LocalName returns the local name for [Resource].
func (akrk *Resource) LocalName() string {
	return akrk.Name
}

// Configuration returns the configuration (args) for [Resource].
func (akrk *Resource) Configuration() interface{} {
	return akrk.Args
}

// DependOn is used for other resources to depend on [Resource].
func (akrk *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(akrk)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (akrk *Resource) Dependencies() terra.Dependencies {
	return akrk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (akrk *Resource) LifecycleManagement() *terra.Lifecycle {
	return akrk.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (akrk *Resource) Attributes() awsKmsReplicaKeyAttributes {
	return awsKmsReplicaKeyAttributes{ref: terra.ReferenceResource(akrk)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (akrk *Resource) ImportState(state io.Reader) error {
	akrk.state = &awsKmsReplicaKeyState{}
	if err := json.NewDecoder(state).Decode(akrk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", akrk.Type(), akrk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (akrk *Resource) State() (*awsKmsReplicaKeyState, bool) {
	return akrk.state, akrk.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (akrk *Resource) StateMust() *awsKmsReplicaKeyState {
	if akrk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", akrk.Type(), akrk.LocalName()))
	}
	return akrk.state
}

// Args contains the configurations for aws_kms_replica_key.
type Args struct {
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

type awsKmsReplicaKeyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kms_replica_key.
func (akrk awsKmsReplicaKeyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(akrk.ref.Append("arn"))
}

// BypassPolicyLockoutSafetyCheck returns a reference to field bypass_policy_lockout_safety_check of aws_kms_replica_key.
func (akrk awsKmsReplicaKeyAttributes) BypassPolicyLockoutSafetyCheck() terra.BoolValue {
	return terra.ReferenceAsBool(akrk.ref.Append("bypass_policy_lockout_safety_check"))
}

// DeletionWindowInDays returns a reference to field deletion_window_in_days of aws_kms_replica_key.
func (akrk awsKmsReplicaKeyAttributes) DeletionWindowInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(akrk.ref.Append("deletion_window_in_days"))
}

// Description returns a reference to field description of aws_kms_replica_key.
func (akrk awsKmsReplicaKeyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(akrk.ref.Append("description"))
}

// Enabled returns a reference to field enabled of aws_kms_replica_key.
func (akrk awsKmsReplicaKeyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(akrk.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_kms_replica_key.
func (akrk awsKmsReplicaKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akrk.ref.Append("id"))
}

// KeyId returns a reference to field key_id of aws_kms_replica_key.
func (akrk awsKmsReplicaKeyAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(akrk.ref.Append("key_id"))
}

// KeyRotationEnabled returns a reference to field key_rotation_enabled of aws_kms_replica_key.
func (akrk awsKmsReplicaKeyAttributes) KeyRotationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(akrk.ref.Append("key_rotation_enabled"))
}

// KeySpec returns a reference to field key_spec of aws_kms_replica_key.
func (akrk awsKmsReplicaKeyAttributes) KeySpec() terra.StringValue {
	return terra.ReferenceAsString(akrk.ref.Append("key_spec"))
}

// KeyUsage returns a reference to field key_usage of aws_kms_replica_key.
func (akrk awsKmsReplicaKeyAttributes) KeyUsage() terra.StringValue {
	return terra.ReferenceAsString(akrk.ref.Append("key_usage"))
}

// Policy returns a reference to field policy of aws_kms_replica_key.
func (akrk awsKmsReplicaKeyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(akrk.ref.Append("policy"))
}

// PrimaryKeyArn returns a reference to field primary_key_arn of aws_kms_replica_key.
func (akrk awsKmsReplicaKeyAttributes) PrimaryKeyArn() terra.StringValue {
	return terra.ReferenceAsString(akrk.ref.Append("primary_key_arn"))
}

// Tags returns a reference to field tags of aws_kms_replica_key.
func (akrk awsKmsReplicaKeyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](akrk.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_kms_replica_key.
func (akrk awsKmsReplicaKeyAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](akrk.ref.Append("tags_all"))
}

type awsKmsReplicaKeyState struct {
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
