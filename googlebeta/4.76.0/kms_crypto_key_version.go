// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	kmscryptokeyversion "github.com/golingon/terraproviders/googlebeta/4.76.0/kmscryptokeyversion"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKmsCryptoKeyVersion creates a new instance of [KmsCryptoKeyVersion].
func NewKmsCryptoKeyVersion(name string, args KmsCryptoKeyVersionArgs) *KmsCryptoKeyVersion {
	return &KmsCryptoKeyVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KmsCryptoKeyVersion)(nil)

// KmsCryptoKeyVersion represents the Terraform resource google_kms_crypto_key_version.
type KmsCryptoKeyVersion struct {
	Name      string
	Args      KmsCryptoKeyVersionArgs
	state     *kmsCryptoKeyVersionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KmsCryptoKeyVersion].
func (kckv *KmsCryptoKeyVersion) Type() string {
	return "google_kms_crypto_key_version"
}

// LocalName returns the local name for [KmsCryptoKeyVersion].
func (kckv *KmsCryptoKeyVersion) LocalName() string {
	return kckv.Name
}

// Configuration returns the configuration (args) for [KmsCryptoKeyVersion].
func (kckv *KmsCryptoKeyVersion) Configuration() interface{} {
	return kckv.Args
}

// DependOn is used for other resources to depend on [KmsCryptoKeyVersion].
func (kckv *KmsCryptoKeyVersion) DependOn() terra.Reference {
	return terra.ReferenceResource(kckv)
}

// Dependencies returns the list of resources [KmsCryptoKeyVersion] depends_on.
func (kckv *KmsCryptoKeyVersion) Dependencies() terra.Dependencies {
	return kckv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KmsCryptoKeyVersion].
func (kckv *KmsCryptoKeyVersion) LifecycleManagement() *terra.Lifecycle {
	return kckv.Lifecycle
}

// Attributes returns the attributes for [KmsCryptoKeyVersion].
func (kckv *KmsCryptoKeyVersion) Attributes() kmsCryptoKeyVersionAttributes {
	return kmsCryptoKeyVersionAttributes{ref: terra.ReferenceResource(kckv)}
}

// ImportState imports the given attribute values into [KmsCryptoKeyVersion]'s state.
func (kckv *KmsCryptoKeyVersion) ImportState(av io.Reader) error {
	kckv.state = &kmsCryptoKeyVersionState{}
	if err := json.NewDecoder(av).Decode(kckv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kckv.Type(), kckv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KmsCryptoKeyVersion] has state.
func (kckv *KmsCryptoKeyVersion) State() (*kmsCryptoKeyVersionState, bool) {
	return kckv.state, kckv.state != nil
}

// StateMust returns the state for [KmsCryptoKeyVersion]. Panics if the state is nil.
func (kckv *KmsCryptoKeyVersion) StateMust() *kmsCryptoKeyVersionState {
	if kckv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kckv.Type(), kckv.LocalName()))
	}
	return kckv.state
}

// KmsCryptoKeyVersionArgs contains the configurations for google_kms_crypto_key_version.
type KmsCryptoKeyVersionArgs struct {
	// CryptoKey: string, required
	CryptoKey terra.StringValue `hcl:"crypto_key,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// Attestation: min=0
	Attestation []kmscryptokeyversion.Attestation `hcl:"attestation,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *kmscryptokeyversion.Timeouts `hcl:"timeouts,block"`
}
type kmsCryptoKeyVersionAttributes struct {
	ref terra.Reference
}

// Algorithm returns a reference to field algorithm of google_kms_crypto_key_version.
func (kckv kmsCryptoKeyVersionAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceAsString(kckv.ref.Append("algorithm"))
}

// CryptoKey returns a reference to field crypto_key of google_kms_crypto_key_version.
func (kckv kmsCryptoKeyVersionAttributes) CryptoKey() terra.StringValue {
	return terra.ReferenceAsString(kckv.ref.Append("crypto_key"))
}

// GenerateTime returns a reference to field generate_time of google_kms_crypto_key_version.
func (kckv kmsCryptoKeyVersionAttributes) GenerateTime() terra.StringValue {
	return terra.ReferenceAsString(kckv.ref.Append("generate_time"))
}

// Id returns a reference to field id of google_kms_crypto_key_version.
func (kckv kmsCryptoKeyVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kckv.ref.Append("id"))
}

// Name returns a reference to field name of google_kms_crypto_key_version.
func (kckv kmsCryptoKeyVersionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kckv.ref.Append("name"))
}

// ProtectionLevel returns a reference to field protection_level of google_kms_crypto_key_version.
func (kckv kmsCryptoKeyVersionAttributes) ProtectionLevel() terra.StringValue {
	return terra.ReferenceAsString(kckv.ref.Append("protection_level"))
}

// State returns a reference to field state of google_kms_crypto_key_version.
func (kckv kmsCryptoKeyVersionAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(kckv.ref.Append("state"))
}

func (kckv kmsCryptoKeyVersionAttributes) Attestation() terra.ListValue[kmscryptokeyversion.AttestationAttributes] {
	return terra.ReferenceAsList[kmscryptokeyversion.AttestationAttributes](kckv.ref.Append("attestation"))
}

func (kckv kmsCryptoKeyVersionAttributes) Timeouts() kmscryptokeyversion.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kmscryptokeyversion.TimeoutsAttributes](kckv.ref.Append("timeouts"))
}

type kmsCryptoKeyVersionState struct {
	Algorithm       string                                 `json:"algorithm"`
	CryptoKey       string                                 `json:"crypto_key"`
	GenerateTime    string                                 `json:"generate_time"`
	Id              string                                 `json:"id"`
	Name            string                                 `json:"name"`
	ProtectionLevel string                                 `json:"protection_level"`
	State           string                                 `json:"state"`
	Attestation     []kmscryptokeyversion.AttestationState `json:"attestation"`
	Timeouts        *kmscryptokeyversion.TimeoutsState     `json:"timeouts"`
}
