// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	kmscryptokeyiambinding "github.com/golingon/terraproviders/google/4.73.1/kmscryptokeyiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKmsCryptoKeyIamBinding creates a new instance of [KmsCryptoKeyIamBinding].
func NewKmsCryptoKeyIamBinding(name string, args KmsCryptoKeyIamBindingArgs) *KmsCryptoKeyIamBinding {
	return &KmsCryptoKeyIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KmsCryptoKeyIamBinding)(nil)

// KmsCryptoKeyIamBinding represents the Terraform resource google_kms_crypto_key_iam_binding.
type KmsCryptoKeyIamBinding struct {
	Name      string
	Args      KmsCryptoKeyIamBindingArgs
	state     *kmsCryptoKeyIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KmsCryptoKeyIamBinding].
func (kckib *KmsCryptoKeyIamBinding) Type() string {
	return "google_kms_crypto_key_iam_binding"
}

// LocalName returns the local name for [KmsCryptoKeyIamBinding].
func (kckib *KmsCryptoKeyIamBinding) LocalName() string {
	return kckib.Name
}

// Configuration returns the configuration (args) for [KmsCryptoKeyIamBinding].
func (kckib *KmsCryptoKeyIamBinding) Configuration() interface{} {
	return kckib.Args
}

// DependOn is used for other resources to depend on [KmsCryptoKeyIamBinding].
func (kckib *KmsCryptoKeyIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(kckib)
}

// Dependencies returns the list of resources [KmsCryptoKeyIamBinding] depends_on.
func (kckib *KmsCryptoKeyIamBinding) Dependencies() terra.Dependencies {
	return kckib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KmsCryptoKeyIamBinding].
func (kckib *KmsCryptoKeyIamBinding) LifecycleManagement() *terra.Lifecycle {
	return kckib.Lifecycle
}

// Attributes returns the attributes for [KmsCryptoKeyIamBinding].
func (kckib *KmsCryptoKeyIamBinding) Attributes() kmsCryptoKeyIamBindingAttributes {
	return kmsCryptoKeyIamBindingAttributes{ref: terra.ReferenceResource(kckib)}
}

// ImportState imports the given attribute values into [KmsCryptoKeyIamBinding]'s state.
func (kckib *KmsCryptoKeyIamBinding) ImportState(av io.Reader) error {
	kckib.state = &kmsCryptoKeyIamBindingState{}
	if err := json.NewDecoder(av).Decode(kckib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kckib.Type(), kckib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KmsCryptoKeyIamBinding] has state.
func (kckib *KmsCryptoKeyIamBinding) State() (*kmsCryptoKeyIamBindingState, bool) {
	return kckib.state, kckib.state != nil
}

// StateMust returns the state for [KmsCryptoKeyIamBinding]. Panics if the state is nil.
func (kckib *KmsCryptoKeyIamBinding) StateMust() *kmsCryptoKeyIamBindingState {
	if kckib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kckib.Type(), kckib.LocalName()))
	}
	return kckib.state
}

// KmsCryptoKeyIamBindingArgs contains the configurations for google_kms_crypto_key_iam_binding.
type KmsCryptoKeyIamBindingArgs struct {
	// CryptoKeyId: string, required
	CryptoKeyId terra.StringValue `hcl:"crypto_key_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *kmscryptokeyiambinding.Condition `hcl:"condition,block"`
}
type kmsCryptoKeyIamBindingAttributes struct {
	ref terra.Reference
}

// CryptoKeyId returns a reference to field crypto_key_id of google_kms_crypto_key_iam_binding.
func (kckib kmsCryptoKeyIamBindingAttributes) CryptoKeyId() terra.StringValue {
	return terra.ReferenceAsString(kckib.ref.Append("crypto_key_id"))
}

// Etag returns a reference to field etag of google_kms_crypto_key_iam_binding.
func (kckib kmsCryptoKeyIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(kckib.ref.Append("etag"))
}

// Id returns a reference to field id of google_kms_crypto_key_iam_binding.
func (kckib kmsCryptoKeyIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kckib.ref.Append("id"))
}

// Members returns a reference to field members of google_kms_crypto_key_iam_binding.
func (kckib kmsCryptoKeyIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kckib.ref.Append("members"))
}

// Role returns a reference to field role of google_kms_crypto_key_iam_binding.
func (kckib kmsCryptoKeyIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(kckib.ref.Append("role"))
}

func (kckib kmsCryptoKeyIamBindingAttributes) Condition() terra.ListValue[kmscryptokeyiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[kmscryptokeyiambinding.ConditionAttributes](kckib.ref.Append("condition"))
}

type kmsCryptoKeyIamBindingState struct {
	CryptoKeyId string                                  `json:"crypto_key_id"`
	Etag        string                                  `json:"etag"`
	Id          string                                  `json:"id"`
	Members     []string                                `json:"members"`
	Role        string                                  `json:"role"`
	Condition   []kmscryptokeyiambinding.ConditionState `json:"condition"`
}
