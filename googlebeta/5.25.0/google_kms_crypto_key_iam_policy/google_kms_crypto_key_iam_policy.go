// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_kms_crypto_key_iam_policy

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

// Resource represents the Terraform resource google_kms_crypto_key_iam_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleKmsCryptoKeyIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gkckip *Resource) Type() string {
	return "google_kms_crypto_key_iam_policy"
}

// LocalName returns the local name for [Resource].
func (gkckip *Resource) LocalName() string {
	return gkckip.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gkckip *Resource) Configuration() interface{} {
	return gkckip.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gkckip *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gkckip)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gkckip *Resource) Dependencies() terra.Dependencies {
	return gkckip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gkckip *Resource) LifecycleManagement() *terra.Lifecycle {
	return gkckip.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gkckip *Resource) Attributes() googleKmsCryptoKeyIamPolicyAttributes {
	return googleKmsCryptoKeyIamPolicyAttributes{ref: terra.ReferenceResource(gkckip)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gkckip *Resource) ImportState(state io.Reader) error {
	gkckip.state = &googleKmsCryptoKeyIamPolicyState{}
	if err := json.NewDecoder(state).Decode(gkckip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gkckip.Type(), gkckip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gkckip *Resource) State() (*googleKmsCryptoKeyIamPolicyState, bool) {
	return gkckip.state, gkckip.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gkckip *Resource) StateMust() *googleKmsCryptoKeyIamPolicyState {
	if gkckip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gkckip.Type(), gkckip.LocalName()))
	}
	return gkckip.state
}

// Args contains the configurations for google_kms_crypto_key_iam_policy.
type Args struct {
	// CryptoKeyId: string, required
	CryptoKeyId terra.StringValue `hcl:"crypto_key_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
}

type googleKmsCryptoKeyIamPolicyAttributes struct {
	ref terra.Reference
}

// CryptoKeyId returns a reference to field crypto_key_id of google_kms_crypto_key_iam_policy.
func (gkckip googleKmsCryptoKeyIamPolicyAttributes) CryptoKeyId() terra.StringValue {
	return terra.ReferenceAsString(gkckip.ref.Append("crypto_key_id"))
}

// Etag returns a reference to field etag of google_kms_crypto_key_iam_policy.
func (gkckip googleKmsCryptoKeyIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gkckip.ref.Append("etag"))
}

// Id returns a reference to field id of google_kms_crypto_key_iam_policy.
func (gkckip googleKmsCryptoKeyIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gkckip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_kms_crypto_key_iam_policy.
func (gkckip googleKmsCryptoKeyIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gkckip.ref.Append("policy_data"))
}

type googleKmsCryptoKeyIamPolicyState struct {
	CryptoKeyId string `json:"crypto_key_id"`
	Etag        string `json:"etag"`
	Id          string `json:"id"`
	PolicyData  string `json:"policy_data"`
}
