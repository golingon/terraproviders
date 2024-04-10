// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewKmsKeyRingIamPolicy creates a new instance of [KmsKeyRingIamPolicy].
func NewKmsKeyRingIamPolicy(name string, args KmsKeyRingIamPolicyArgs) *KmsKeyRingIamPolicy {
	return &KmsKeyRingIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KmsKeyRingIamPolicy)(nil)

// KmsKeyRingIamPolicy represents the Terraform resource google_kms_key_ring_iam_policy.
type KmsKeyRingIamPolicy struct {
	Name      string
	Args      KmsKeyRingIamPolicyArgs
	state     *kmsKeyRingIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KmsKeyRingIamPolicy].
func (kkrip *KmsKeyRingIamPolicy) Type() string {
	return "google_kms_key_ring_iam_policy"
}

// LocalName returns the local name for [KmsKeyRingIamPolicy].
func (kkrip *KmsKeyRingIamPolicy) LocalName() string {
	return kkrip.Name
}

// Configuration returns the configuration (args) for [KmsKeyRingIamPolicy].
func (kkrip *KmsKeyRingIamPolicy) Configuration() interface{} {
	return kkrip.Args
}

// DependOn is used for other resources to depend on [KmsKeyRingIamPolicy].
func (kkrip *KmsKeyRingIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(kkrip)
}

// Dependencies returns the list of resources [KmsKeyRingIamPolicy] depends_on.
func (kkrip *KmsKeyRingIamPolicy) Dependencies() terra.Dependencies {
	return kkrip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KmsKeyRingIamPolicy].
func (kkrip *KmsKeyRingIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return kkrip.Lifecycle
}

// Attributes returns the attributes for [KmsKeyRingIamPolicy].
func (kkrip *KmsKeyRingIamPolicy) Attributes() kmsKeyRingIamPolicyAttributes {
	return kmsKeyRingIamPolicyAttributes{ref: terra.ReferenceResource(kkrip)}
}

// ImportState imports the given attribute values into [KmsKeyRingIamPolicy]'s state.
func (kkrip *KmsKeyRingIamPolicy) ImportState(av io.Reader) error {
	kkrip.state = &kmsKeyRingIamPolicyState{}
	if err := json.NewDecoder(av).Decode(kkrip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kkrip.Type(), kkrip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KmsKeyRingIamPolicy] has state.
func (kkrip *KmsKeyRingIamPolicy) State() (*kmsKeyRingIamPolicyState, bool) {
	return kkrip.state, kkrip.state != nil
}

// StateMust returns the state for [KmsKeyRingIamPolicy]. Panics if the state is nil.
func (kkrip *KmsKeyRingIamPolicy) StateMust() *kmsKeyRingIamPolicyState {
	if kkrip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kkrip.Type(), kkrip.LocalName()))
	}
	return kkrip.state
}

// KmsKeyRingIamPolicyArgs contains the configurations for google_kms_key_ring_iam_policy.
type KmsKeyRingIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyRingId: string, required
	KeyRingId terra.StringValue `hcl:"key_ring_id,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
}
type kmsKeyRingIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_kms_key_ring_iam_policy.
func (kkrip kmsKeyRingIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(kkrip.ref.Append("etag"))
}

// Id returns a reference to field id of google_kms_key_ring_iam_policy.
func (kkrip kmsKeyRingIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kkrip.ref.Append("id"))
}

// KeyRingId returns a reference to field key_ring_id of google_kms_key_ring_iam_policy.
func (kkrip kmsKeyRingIamPolicyAttributes) KeyRingId() terra.StringValue {
	return terra.ReferenceAsString(kkrip.ref.Append("key_ring_id"))
}

// PolicyData returns a reference to field policy_data of google_kms_key_ring_iam_policy.
func (kkrip kmsKeyRingIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(kkrip.ref.Append("policy_data"))
}

type kmsKeyRingIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	KeyRingId  string `json:"key_ring_id"`
	PolicyData string `json:"policy_data"`
}
