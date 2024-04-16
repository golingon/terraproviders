// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_kms_replica_external_key

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

// Resource represents the Terraform resource aws_kms_replica_external_key.
type Resource struct {
	Name      string
	Args      Args
	state     *awsKmsReplicaExternalKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (akrek *Resource) Type() string {
	return "aws_kms_replica_external_key"
}

// LocalName returns the local name for [Resource].
func (akrek *Resource) LocalName() string {
	return akrek.Name
}

// Configuration returns the configuration (args) for [Resource].
func (akrek *Resource) Configuration() interface{} {
	return akrek.Args
}

// DependOn is used for other resources to depend on [Resource].
func (akrek *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(akrek)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (akrek *Resource) Dependencies() terra.Dependencies {
	return akrek.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (akrek *Resource) LifecycleManagement() *terra.Lifecycle {
	return akrek.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (akrek *Resource) Attributes() awsKmsReplicaExternalKeyAttributes {
	return awsKmsReplicaExternalKeyAttributes{ref: terra.ReferenceResource(akrek)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (akrek *Resource) ImportState(state io.Reader) error {
	akrek.state = &awsKmsReplicaExternalKeyState{}
	if err := json.NewDecoder(state).Decode(akrek.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", akrek.Type(), akrek.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (akrek *Resource) State() (*awsKmsReplicaExternalKeyState, bool) {
	return akrek.state, akrek.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (akrek *Resource) StateMust() *awsKmsReplicaExternalKeyState {
	if akrek.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", akrek.Type(), akrek.LocalName()))
	}
	return akrek.state
}

// Args contains the configurations for aws_kms_replica_external_key.
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
	// KeyMaterialBase64: string, optional
	KeyMaterialBase64 terra.StringValue `hcl:"key_material_base64,attr"`
	// Policy: string, optional
	Policy terra.StringValue `hcl:"policy,attr"`
	// PrimaryKeyArn: string, required
	PrimaryKeyArn terra.StringValue `hcl:"primary_key_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ValidTo: string, optional
	ValidTo terra.StringValue `hcl:"valid_to,attr"`
}

type awsKmsReplicaExternalKeyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kms_replica_external_key.
func (akrek awsKmsReplicaExternalKeyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(akrek.ref.Append("arn"))
}

// BypassPolicyLockoutSafetyCheck returns a reference to field bypass_policy_lockout_safety_check of aws_kms_replica_external_key.
func (akrek awsKmsReplicaExternalKeyAttributes) BypassPolicyLockoutSafetyCheck() terra.BoolValue {
	return terra.ReferenceAsBool(akrek.ref.Append("bypass_policy_lockout_safety_check"))
}

// DeletionWindowInDays returns a reference to field deletion_window_in_days of aws_kms_replica_external_key.
func (akrek awsKmsReplicaExternalKeyAttributes) DeletionWindowInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(akrek.ref.Append("deletion_window_in_days"))
}

// Description returns a reference to field description of aws_kms_replica_external_key.
func (akrek awsKmsReplicaExternalKeyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(akrek.ref.Append("description"))
}

// Enabled returns a reference to field enabled of aws_kms_replica_external_key.
func (akrek awsKmsReplicaExternalKeyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(akrek.ref.Append("enabled"))
}

// ExpirationModel returns a reference to field expiration_model of aws_kms_replica_external_key.
func (akrek awsKmsReplicaExternalKeyAttributes) ExpirationModel() terra.StringValue {
	return terra.ReferenceAsString(akrek.ref.Append("expiration_model"))
}

// Id returns a reference to field id of aws_kms_replica_external_key.
func (akrek awsKmsReplicaExternalKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akrek.ref.Append("id"))
}

// KeyId returns a reference to field key_id of aws_kms_replica_external_key.
func (akrek awsKmsReplicaExternalKeyAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(akrek.ref.Append("key_id"))
}

// KeyMaterialBase64 returns a reference to field key_material_base64 of aws_kms_replica_external_key.
func (akrek awsKmsReplicaExternalKeyAttributes) KeyMaterialBase64() terra.StringValue {
	return terra.ReferenceAsString(akrek.ref.Append("key_material_base64"))
}

// KeyState returns a reference to field key_state of aws_kms_replica_external_key.
func (akrek awsKmsReplicaExternalKeyAttributes) KeyState() terra.StringValue {
	return terra.ReferenceAsString(akrek.ref.Append("key_state"))
}

// KeyUsage returns a reference to field key_usage of aws_kms_replica_external_key.
func (akrek awsKmsReplicaExternalKeyAttributes) KeyUsage() terra.StringValue {
	return terra.ReferenceAsString(akrek.ref.Append("key_usage"))
}

// Policy returns a reference to field policy of aws_kms_replica_external_key.
func (akrek awsKmsReplicaExternalKeyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(akrek.ref.Append("policy"))
}

// PrimaryKeyArn returns a reference to field primary_key_arn of aws_kms_replica_external_key.
func (akrek awsKmsReplicaExternalKeyAttributes) PrimaryKeyArn() terra.StringValue {
	return terra.ReferenceAsString(akrek.ref.Append("primary_key_arn"))
}

// Tags returns a reference to field tags of aws_kms_replica_external_key.
func (akrek awsKmsReplicaExternalKeyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](akrek.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_kms_replica_external_key.
func (akrek awsKmsReplicaExternalKeyAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](akrek.ref.Append("tags_all"))
}

// ValidTo returns a reference to field valid_to of aws_kms_replica_external_key.
func (akrek awsKmsReplicaExternalKeyAttributes) ValidTo() terra.StringValue {
	return terra.ReferenceAsString(akrek.ref.Append("valid_to"))
}

type awsKmsReplicaExternalKeyState struct {
	Arn                            string            `json:"arn"`
	BypassPolicyLockoutSafetyCheck bool              `json:"bypass_policy_lockout_safety_check"`
	DeletionWindowInDays           float64           `json:"deletion_window_in_days"`
	Description                    string            `json:"description"`
	Enabled                        bool              `json:"enabled"`
	ExpirationModel                string            `json:"expiration_model"`
	Id                             string            `json:"id"`
	KeyId                          string            `json:"key_id"`
	KeyMaterialBase64              string            `json:"key_material_base64"`
	KeyState                       string            `json:"key_state"`
	KeyUsage                       string            `json:"key_usage"`
	Policy                         string            `json:"policy"`
	PrimaryKeyArn                  string            `json:"primary_key_arn"`
	Tags                           map[string]string `json:"tags"`
	TagsAll                        map[string]string `json:"tags_all"`
	ValidTo                        string            `json:"valid_to"`
}
