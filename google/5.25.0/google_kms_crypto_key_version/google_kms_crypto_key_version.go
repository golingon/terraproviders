// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_kms_crypto_key_version

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

// Resource represents the Terraform resource google_kms_crypto_key_version.
type Resource struct {
	Name      string
	Args      Args
	state     *googleKmsCryptoKeyVersionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gkckv *Resource) Type() string {
	return "google_kms_crypto_key_version"
}

// LocalName returns the local name for [Resource].
func (gkckv *Resource) LocalName() string {
	return gkckv.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gkckv *Resource) Configuration() interface{} {
	return gkckv.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gkckv *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gkckv)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gkckv *Resource) Dependencies() terra.Dependencies {
	return gkckv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gkckv *Resource) LifecycleManagement() *terra.Lifecycle {
	return gkckv.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gkckv *Resource) Attributes() googleKmsCryptoKeyVersionAttributes {
	return googleKmsCryptoKeyVersionAttributes{ref: terra.ReferenceResource(gkckv)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gkckv *Resource) ImportState(state io.Reader) error {
	gkckv.state = &googleKmsCryptoKeyVersionState{}
	if err := json.NewDecoder(state).Decode(gkckv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gkckv.Type(), gkckv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gkckv *Resource) State() (*googleKmsCryptoKeyVersionState, bool) {
	return gkckv.state, gkckv.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gkckv *Resource) StateMust() *googleKmsCryptoKeyVersionState {
	if gkckv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gkckv.Type(), gkckv.LocalName()))
	}
	return gkckv.state
}

// Args contains the configurations for google_kms_crypto_key_version.
type Args struct {
	// CryptoKey: string, required
	CryptoKey terra.StringValue `hcl:"crypto_key,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// ExternalProtectionLevelOptions: optional
	ExternalProtectionLevelOptions *ExternalProtectionLevelOptions `hcl:"external_protection_level_options,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleKmsCryptoKeyVersionAttributes struct {
	ref terra.Reference
}

// Algorithm returns a reference to field algorithm of google_kms_crypto_key_version.
func (gkckv googleKmsCryptoKeyVersionAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceAsString(gkckv.ref.Append("algorithm"))
}

// CryptoKey returns a reference to field crypto_key of google_kms_crypto_key_version.
func (gkckv googleKmsCryptoKeyVersionAttributes) CryptoKey() terra.StringValue {
	return terra.ReferenceAsString(gkckv.ref.Append("crypto_key"))
}

// GenerateTime returns a reference to field generate_time of google_kms_crypto_key_version.
func (gkckv googleKmsCryptoKeyVersionAttributes) GenerateTime() terra.StringValue {
	return terra.ReferenceAsString(gkckv.ref.Append("generate_time"))
}

// Id returns a reference to field id of google_kms_crypto_key_version.
func (gkckv googleKmsCryptoKeyVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gkckv.ref.Append("id"))
}

// Name returns a reference to field name of google_kms_crypto_key_version.
func (gkckv googleKmsCryptoKeyVersionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gkckv.ref.Append("name"))
}

// ProtectionLevel returns a reference to field protection_level of google_kms_crypto_key_version.
func (gkckv googleKmsCryptoKeyVersionAttributes) ProtectionLevel() terra.StringValue {
	return terra.ReferenceAsString(gkckv.ref.Append("protection_level"))
}

// State returns a reference to field state of google_kms_crypto_key_version.
func (gkckv googleKmsCryptoKeyVersionAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(gkckv.ref.Append("state"))
}

func (gkckv googleKmsCryptoKeyVersionAttributes) Attestation() terra.ListValue[AttestationAttributes] {
	return terra.ReferenceAsList[AttestationAttributes](gkckv.ref.Append("attestation"))
}

func (gkckv googleKmsCryptoKeyVersionAttributes) ExternalProtectionLevelOptions() terra.ListValue[ExternalProtectionLevelOptionsAttributes] {
	return terra.ReferenceAsList[ExternalProtectionLevelOptionsAttributes](gkckv.ref.Append("external_protection_level_options"))
}

func (gkckv googleKmsCryptoKeyVersionAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gkckv.ref.Append("timeouts"))
}

type googleKmsCryptoKeyVersionState struct {
	Algorithm                      string                                `json:"algorithm"`
	CryptoKey                      string                                `json:"crypto_key"`
	GenerateTime                   string                                `json:"generate_time"`
	Id                             string                                `json:"id"`
	Name                           string                                `json:"name"`
	ProtectionLevel                string                                `json:"protection_level"`
	State                          string                                `json:"state"`
	Attestation                    []AttestationState                    `json:"attestation"`
	ExternalProtectionLevelOptions []ExternalProtectionLevelOptionsState `json:"external_protection_level_options"`
	Timeouts                       *TimeoutsState                        `json:"timeouts"`
}
