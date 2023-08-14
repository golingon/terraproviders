// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKmsKey creates a new instance of [KmsKey].
func NewKmsKey(name string, args KmsKeyArgs) *KmsKey {
	return &KmsKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KmsKey)(nil)

// KmsKey represents the Terraform resource aws_kms_key.
type KmsKey struct {
	Name      string
	Args      KmsKeyArgs
	state     *kmsKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KmsKey].
func (kk *KmsKey) Type() string {
	return "aws_kms_key"
}

// LocalName returns the local name for [KmsKey].
func (kk *KmsKey) LocalName() string {
	return kk.Name
}

// Configuration returns the configuration (args) for [KmsKey].
func (kk *KmsKey) Configuration() interface{} {
	return kk.Args
}

// DependOn is used for other resources to depend on [KmsKey].
func (kk *KmsKey) DependOn() terra.Reference {
	return terra.ReferenceResource(kk)
}

// Dependencies returns the list of resources [KmsKey] depends_on.
func (kk *KmsKey) Dependencies() terra.Dependencies {
	return kk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KmsKey].
func (kk *KmsKey) LifecycleManagement() *terra.Lifecycle {
	return kk.Lifecycle
}

// Attributes returns the attributes for [KmsKey].
func (kk *KmsKey) Attributes() kmsKeyAttributes {
	return kmsKeyAttributes{ref: terra.ReferenceResource(kk)}
}

// ImportState imports the given attribute values into [KmsKey]'s state.
func (kk *KmsKey) ImportState(av io.Reader) error {
	kk.state = &kmsKeyState{}
	if err := json.NewDecoder(av).Decode(kk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kk.Type(), kk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KmsKey] has state.
func (kk *KmsKey) State() (*kmsKeyState, bool) {
	return kk.state, kk.state != nil
}

// StateMust returns the state for [KmsKey]. Panics if the state is nil.
func (kk *KmsKey) StateMust() *kmsKeyState {
	if kk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kk.Type(), kk.LocalName()))
	}
	return kk.state
}

// KmsKeyArgs contains the configurations for aws_kms_key.
type KmsKeyArgs struct {
	// BypassPolicyLockoutSafetyCheck: bool, optional
	BypassPolicyLockoutSafetyCheck terra.BoolValue `hcl:"bypass_policy_lockout_safety_check,attr"`
	// CustomKeyStoreId: string, optional
	CustomKeyStoreId terra.StringValue `hcl:"custom_key_store_id,attr"`
	// CustomerMasterKeySpec: string, optional
	CustomerMasterKeySpec terra.StringValue `hcl:"customer_master_key_spec,attr"`
	// DeletionWindowInDays: number, optional
	DeletionWindowInDays terra.NumberValue `hcl:"deletion_window_in_days,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EnableKeyRotation: bool, optional
	EnableKeyRotation terra.BoolValue `hcl:"enable_key_rotation,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsEnabled: bool, optional
	IsEnabled terra.BoolValue `hcl:"is_enabled,attr"`
	// KeyUsage: string, optional
	KeyUsage terra.StringValue `hcl:"key_usage,attr"`
	// MultiRegion: bool, optional
	MultiRegion terra.BoolValue `hcl:"multi_region,attr"`
	// Policy: string, optional
	Policy terra.StringValue `hcl:"policy,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type kmsKeyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kms_key.
func (kk kmsKeyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("arn"))
}

// BypassPolicyLockoutSafetyCheck returns a reference to field bypass_policy_lockout_safety_check of aws_kms_key.
func (kk kmsKeyAttributes) BypassPolicyLockoutSafetyCheck() terra.BoolValue {
	return terra.ReferenceAsBool(kk.ref.Append("bypass_policy_lockout_safety_check"))
}

// CustomKeyStoreId returns a reference to field custom_key_store_id of aws_kms_key.
func (kk kmsKeyAttributes) CustomKeyStoreId() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("custom_key_store_id"))
}

// CustomerMasterKeySpec returns a reference to field customer_master_key_spec of aws_kms_key.
func (kk kmsKeyAttributes) CustomerMasterKeySpec() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("customer_master_key_spec"))
}

// DeletionWindowInDays returns a reference to field deletion_window_in_days of aws_kms_key.
func (kk kmsKeyAttributes) DeletionWindowInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(kk.ref.Append("deletion_window_in_days"))
}

// Description returns a reference to field description of aws_kms_key.
func (kk kmsKeyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("description"))
}

// EnableKeyRotation returns a reference to field enable_key_rotation of aws_kms_key.
func (kk kmsKeyAttributes) EnableKeyRotation() terra.BoolValue {
	return terra.ReferenceAsBool(kk.ref.Append("enable_key_rotation"))
}

// Id returns a reference to field id of aws_kms_key.
func (kk kmsKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("id"))
}

// IsEnabled returns a reference to field is_enabled of aws_kms_key.
func (kk kmsKeyAttributes) IsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kk.ref.Append("is_enabled"))
}

// KeyId returns a reference to field key_id of aws_kms_key.
func (kk kmsKeyAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("key_id"))
}

// KeyUsage returns a reference to field key_usage of aws_kms_key.
func (kk kmsKeyAttributes) KeyUsage() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("key_usage"))
}

// MultiRegion returns a reference to field multi_region of aws_kms_key.
func (kk kmsKeyAttributes) MultiRegion() terra.BoolValue {
	return terra.ReferenceAsBool(kk.ref.Append("multi_region"))
}

// Policy returns a reference to field policy of aws_kms_key.
func (kk kmsKeyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(kk.ref.Append("policy"))
}

// Tags returns a reference to field tags of aws_kms_key.
func (kk kmsKeyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kk.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_kms_key.
func (kk kmsKeyAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kk.ref.Append("tags_all"))
}

type kmsKeyState struct {
	Arn                            string            `json:"arn"`
	BypassPolicyLockoutSafetyCheck bool              `json:"bypass_policy_lockout_safety_check"`
	CustomKeyStoreId               string            `json:"custom_key_store_id"`
	CustomerMasterKeySpec          string            `json:"customer_master_key_spec"`
	DeletionWindowInDays           float64           `json:"deletion_window_in_days"`
	Description                    string            `json:"description"`
	EnableKeyRotation              bool              `json:"enable_key_rotation"`
	Id                             string            `json:"id"`
	IsEnabled                      bool              `json:"is_enabled"`
	KeyId                          string            `json:"key_id"`
	KeyUsage                       string            `json:"key_usage"`
	MultiRegion                    bool              `json:"multi_region"`
	Policy                         string            `json:"policy"`
	Tags                           map[string]string `json:"tags"`
	TagsAll                        map[string]string `json:"tags_all"`
}