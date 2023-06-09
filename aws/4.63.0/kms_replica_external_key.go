// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKmsReplicaExternalKey creates a new instance of [KmsReplicaExternalKey].
func NewKmsReplicaExternalKey(name string, args KmsReplicaExternalKeyArgs) *KmsReplicaExternalKey {
	return &KmsReplicaExternalKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KmsReplicaExternalKey)(nil)

// KmsReplicaExternalKey represents the Terraform resource aws_kms_replica_external_key.
type KmsReplicaExternalKey struct {
	Name      string
	Args      KmsReplicaExternalKeyArgs
	state     *kmsReplicaExternalKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KmsReplicaExternalKey].
func (krek *KmsReplicaExternalKey) Type() string {
	return "aws_kms_replica_external_key"
}

// LocalName returns the local name for [KmsReplicaExternalKey].
func (krek *KmsReplicaExternalKey) LocalName() string {
	return krek.Name
}

// Configuration returns the configuration (args) for [KmsReplicaExternalKey].
func (krek *KmsReplicaExternalKey) Configuration() interface{} {
	return krek.Args
}

// DependOn is used for other resources to depend on [KmsReplicaExternalKey].
func (krek *KmsReplicaExternalKey) DependOn() terra.Reference {
	return terra.ReferenceResource(krek)
}

// Dependencies returns the list of resources [KmsReplicaExternalKey] depends_on.
func (krek *KmsReplicaExternalKey) Dependencies() terra.Dependencies {
	return krek.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KmsReplicaExternalKey].
func (krek *KmsReplicaExternalKey) LifecycleManagement() *terra.Lifecycle {
	return krek.Lifecycle
}

// Attributes returns the attributes for [KmsReplicaExternalKey].
func (krek *KmsReplicaExternalKey) Attributes() kmsReplicaExternalKeyAttributes {
	return kmsReplicaExternalKeyAttributes{ref: terra.ReferenceResource(krek)}
}

// ImportState imports the given attribute values into [KmsReplicaExternalKey]'s state.
func (krek *KmsReplicaExternalKey) ImportState(av io.Reader) error {
	krek.state = &kmsReplicaExternalKeyState{}
	if err := json.NewDecoder(av).Decode(krek.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", krek.Type(), krek.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KmsReplicaExternalKey] has state.
func (krek *KmsReplicaExternalKey) State() (*kmsReplicaExternalKeyState, bool) {
	return krek.state, krek.state != nil
}

// StateMust returns the state for [KmsReplicaExternalKey]. Panics if the state is nil.
func (krek *KmsReplicaExternalKey) StateMust() *kmsReplicaExternalKeyState {
	if krek.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", krek.Type(), krek.LocalName()))
	}
	return krek.state
}

// KmsReplicaExternalKeyArgs contains the configurations for aws_kms_replica_external_key.
type KmsReplicaExternalKeyArgs struct {
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
type kmsReplicaExternalKeyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kms_replica_external_key.
func (krek kmsReplicaExternalKeyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(krek.ref.Append("arn"))
}

// BypassPolicyLockoutSafetyCheck returns a reference to field bypass_policy_lockout_safety_check of aws_kms_replica_external_key.
func (krek kmsReplicaExternalKeyAttributes) BypassPolicyLockoutSafetyCheck() terra.BoolValue {
	return terra.ReferenceAsBool(krek.ref.Append("bypass_policy_lockout_safety_check"))
}

// DeletionWindowInDays returns a reference to field deletion_window_in_days of aws_kms_replica_external_key.
func (krek kmsReplicaExternalKeyAttributes) DeletionWindowInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(krek.ref.Append("deletion_window_in_days"))
}

// Description returns a reference to field description of aws_kms_replica_external_key.
func (krek kmsReplicaExternalKeyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(krek.ref.Append("description"))
}

// Enabled returns a reference to field enabled of aws_kms_replica_external_key.
func (krek kmsReplicaExternalKeyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(krek.ref.Append("enabled"))
}

// ExpirationModel returns a reference to field expiration_model of aws_kms_replica_external_key.
func (krek kmsReplicaExternalKeyAttributes) ExpirationModel() terra.StringValue {
	return terra.ReferenceAsString(krek.ref.Append("expiration_model"))
}

// Id returns a reference to field id of aws_kms_replica_external_key.
func (krek kmsReplicaExternalKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(krek.ref.Append("id"))
}

// KeyId returns a reference to field key_id of aws_kms_replica_external_key.
func (krek kmsReplicaExternalKeyAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(krek.ref.Append("key_id"))
}

// KeyMaterialBase64 returns a reference to field key_material_base64 of aws_kms_replica_external_key.
func (krek kmsReplicaExternalKeyAttributes) KeyMaterialBase64() terra.StringValue {
	return terra.ReferenceAsString(krek.ref.Append("key_material_base64"))
}

// KeyState returns a reference to field key_state of aws_kms_replica_external_key.
func (krek kmsReplicaExternalKeyAttributes) KeyState() terra.StringValue {
	return terra.ReferenceAsString(krek.ref.Append("key_state"))
}

// KeyUsage returns a reference to field key_usage of aws_kms_replica_external_key.
func (krek kmsReplicaExternalKeyAttributes) KeyUsage() terra.StringValue {
	return terra.ReferenceAsString(krek.ref.Append("key_usage"))
}

// Policy returns a reference to field policy of aws_kms_replica_external_key.
func (krek kmsReplicaExternalKeyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(krek.ref.Append("policy"))
}

// PrimaryKeyArn returns a reference to field primary_key_arn of aws_kms_replica_external_key.
func (krek kmsReplicaExternalKeyAttributes) PrimaryKeyArn() terra.StringValue {
	return terra.ReferenceAsString(krek.ref.Append("primary_key_arn"))
}

// Tags returns a reference to field tags of aws_kms_replica_external_key.
func (krek kmsReplicaExternalKeyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](krek.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_kms_replica_external_key.
func (krek kmsReplicaExternalKeyAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](krek.ref.Append("tags_all"))
}

// ValidTo returns a reference to field valid_to of aws_kms_replica_external_key.
func (krek kmsReplicaExternalKeyAttributes) ValidTo() terra.StringValue {
	return terra.ReferenceAsString(krek.ref.Append("valid_to"))
}

type kmsReplicaExternalKeyState struct {
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
