// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	kmssecretciphertext "github.com/golingon/terraproviders/googlebeta/4.73.1/kmssecretciphertext"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKmsSecretCiphertext creates a new instance of [KmsSecretCiphertext].
func NewKmsSecretCiphertext(name string, args KmsSecretCiphertextArgs) *KmsSecretCiphertext {
	return &KmsSecretCiphertext{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KmsSecretCiphertext)(nil)

// KmsSecretCiphertext represents the Terraform resource google_kms_secret_ciphertext.
type KmsSecretCiphertext struct {
	Name      string
	Args      KmsSecretCiphertextArgs
	state     *kmsSecretCiphertextState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KmsSecretCiphertext].
func (ksc *KmsSecretCiphertext) Type() string {
	return "google_kms_secret_ciphertext"
}

// LocalName returns the local name for [KmsSecretCiphertext].
func (ksc *KmsSecretCiphertext) LocalName() string {
	return ksc.Name
}

// Configuration returns the configuration (args) for [KmsSecretCiphertext].
func (ksc *KmsSecretCiphertext) Configuration() interface{} {
	return ksc.Args
}

// DependOn is used for other resources to depend on [KmsSecretCiphertext].
func (ksc *KmsSecretCiphertext) DependOn() terra.Reference {
	return terra.ReferenceResource(ksc)
}

// Dependencies returns the list of resources [KmsSecretCiphertext] depends_on.
func (ksc *KmsSecretCiphertext) Dependencies() terra.Dependencies {
	return ksc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KmsSecretCiphertext].
func (ksc *KmsSecretCiphertext) LifecycleManagement() *terra.Lifecycle {
	return ksc.Lifecycle
}

// Attributes returns the attributes for [KmsSecretCiphertext].
func (ksc *KmsSecretCiphertext) Attributes() kmsSecretCiphertextAttributes {
	return kmsSecretCiphertextAttributes{ref: terra.ReferenceResource(ksc)}
}

// ImportState imports the given attribute values into [KmsSecretCiphertext]'s state.
func (ksc *KmsSecretCiphertext) ImportState(av io.Reader) error {
	ksc.state = &kmsSecretCiphertextState{}
	if err := json.NewDecoder(av).Decode(ksc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ksc.Type(), ksc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KmsSecretCiphertext] has state.
func (ksc *KmsSecretCiphertext) State() (*kmsSecretCiphertextState, bool) {
	return ksc.state, ksc.state != nil
}

// StateMust returns the state for [KmsSecretCiphertext]. Panics if the state is nil.
func (ksc *KmsSecretCiphertext) StateMust() *kmsSecretCiphertextState {
	if ksc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ksc.Type(), ksc.LocalName()))
	}
	return ksc.state
}

// KmsSecretCiphertextArgs contains the configurations for google_kms_secret_ciphertext.
type KmsSecretCiphertextArgs struct {
	// AdditionalAuthenticatedData: string, optional
	AdditionalAuthenticatedData terra.StringValue `hcl:"additional_authenticated_data,attr"`
	// CryptoKey: string, required
	CryptoKey terra.StringValue `hcl:"crypto_key,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Plaintext: string, required
	Plaintext terra.StringValue `hcl:"plaintext,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *kmssecretciphertext.Timeouts `hcl:"timeouts,block"`
}
type kmsSecretCiphertextAttributes struct {
	ref terra.Reference
}

// AdditionalAuthenticatedData returns a reference to field additional_authenticated_data of google_kms_secret_ciphertext.
func (ksc kmsSecretCiphertextAttributes) AdditionalAuthenticatedData() terra.StringValue {
	return terra.ReferenceAsString(ksc.ref.Append("additional_authenticated_data"))
}

// Ciphertext returns a reference to field ciphertext of google_kms_secret_ciphertext.
func (ksc kmsSecretCiphertextAttributes) Ciphertext() terra.StringValue {
	return terra.ReferenceAsString(ksc.ref.Append("ciphertext"))
}

// CryptoKey returns a reference to field crypto_key of google_kms_secret_ciphertext.
func (ksc kmsSecretCiphertextAttributes) CryptoKey() terra.StringValue {
	return terra.ReferenceAsString(ksc.ref.Append("crypto_key"))
}

// Id returns a reference to field id of google_kms_secret_ciphertext.
func (ksc kmsSecretCiphertextAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ksc.ref.Append("id"))
}

// Plaintext returns a reference to field plaintext of google_kms_secret_ciphertext.
func (ksc kmsSecretCiphertextAttributes) Plaintext() terra.StringValue {
	return terra.ReferenceAsString(ksc.ref.Append("plaintext"))
}

func (ksc kmsSecretCiphertextAttributes) Timeouts() kmssecretciphertext.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kmssecretciphertext.TimeoutsAttributes](ksc.ref.Append("timeouts"))
}

type kmsSecretCiphertextState struct {
	AdditionalAuthenticatedData string                             `json:"additional_authenticated_data"`
	Ciphertext                  string                             `json:"ciphertext"`
	CryptoKey                   string                             `json:"crypto_key"`
	Id                          string                             `json:"id"`
	Plaintext                   string                             `json:"plaintext"`
	Timeouts                    *kmssecretciphertext.TimeoutsState `json:"timeouts"`
}