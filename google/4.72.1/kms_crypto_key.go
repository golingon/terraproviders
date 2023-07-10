// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	kmscryptokey "github.com/golingon/terraproviders/google/4.72.1/kmscryptokey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKmsCryptoKey creates a new instance of [KmsCryptoKey].
func NewKmsCryptoKey(name string, args KmsCryptoKeyArgs) *KmsCryptoKey {
	return &KmsCryptoKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KmsCryptoKey)(nil)

// KmsCryptoKey represents the Terraform resource google_kms_crypto_key.
type KmsCryptoKey struct {
	Name      string
	Args      KmsCryptoKeyArgs
	state     *kmsCryptoKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KmsCryptoKey].
func (kck *KmsCryptoKey) Type() string {
	return "google_kms_crypto_key"
}

// LocalName returns the local name for [KmsCryptoKey].
func (kck *KmsCryptoKey) LocalName() string {
	return kck.Name
}

// Configuration returns the configuration (args) for [KmsCryptoKey].
func (kck *KmsCryptoKey) Configuration() interface{} {
	return kck.Args
}

// DependOn is used for other resources to depend on [KmsCryptoKey].
func (kck *KmsCryptoKey) DependOn() terra.Reference {
	return terra.ReferenceResource(kck)
}

// Dependencies returns the list of resources [KmsCryptoKey] depends_on.
func (kck *KmsCryptoKey) Dependencies() terra.Dependencies {
	return kck.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KmsCryptoKey].
func (kck *KmsCryptoKey) LifecycleManagement() *terra.Lifecycle {
	return kck.Lifecycle
}

// Attributes returns the attributes for [KmsCryptoKey].
func (kck *KmsCryptoKey) Attributes() kmsCryptoKeyAttributes {
	return kmsCryptoKeyAttributes{ref: terra.ReferenceResource(kck)}
}

// ImportState imports the given attribute values into [KmsCryptoKey]'s state.
func (kck *KmsCryptoKey) ImportState(av io.Reader) error {
	kck.state = &kmsCryptoKeyState{}
	if err := json.NewDecoder(av).Decode(kck.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kck.Type(), kck.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KmsCryptoKey] has state.
func (kck *KmsCryptoKey) State() (*kmsCryptoKeyState, bool) {
	return kck.state, kck.state != nil
}

// StateMust returns the state for [KmsCryptoKey]. Panics if the state is nil.
func (kck *KmsCryptoKey) StateMust() *kmsCryptoKeyState {
	if kck.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kck.Type(), kck.LocalName()))
	}
	return kck.state
}

// KmsCryptoKeyArgs contains the configurations for google_kms_crypto_key.
type KmsCryptoKeyArgs struct {
	// DestroyScheduledDuration: string, optional
	DestroyScheduledDuration terra.StringValue `hcl:"destroy_scheduled_duration,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImportOnly: bool, optional
	ImportOnly terra.BoolValue `hcl:"import_only,attr"`
	// KeyRing: string, required
	KeyRing terra.StringValue `hcl:"key_ring,attr" validate:"required"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Purpose: string, optional
	Purpose terra.StringValue `hcl:"purpose,attr"`
	// RotationPeriod: string, optional
	RotationPeriod terra.StringValue `hcl:"rotation_period,attr"`
	// SkipInitialVersionCreation: bool, optional
	SkipInitialVersionCreation terra.BoolValue `hcl:"skip_initial_version_creation,attr"`
	// Timeouts: optional
	Timeouts *kmscryptokey.Timeouts `hcl:"timeouts,block"`
	// VersionTemplate: optional
	VersionTemplate *kmscryptokey.VersionTemplate `hcl:"version_template,block"`
}
type kmsCryptoKeyAttributes struct {
	ref terra.Reference
}

// DestroyScheduledDuration returns a reference to field destroy_scheduled_duration of google_kms_crypto_key.
func (kck kmsCryptoKeyAttributes) DestroyScheduledDuration() terra.StringValue {
	return terra.ReferenceAsString(kck.ref.Append("destroy_scheduled_duration"))
}

// Id returns a reference to field id of google_kms_crypto_key.
func (kck kmsCryptoKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kck.ref.Append("id"))
}

// ImportOnly returns a reference to field import_only of google_kms_crypto_key.
func (kck kmsCryptoKeyAttributes) ImportOnly() terra.BoolValue {
	return terra.ReferenceAsBool(kck.ref.Append("import_only"))
}

// KeyRing returns a reference to field key_ring of google_kms_crypto_key.
func (kck kmsCryptoKeyAttributes) KeyRing() terra.StringValue {
	return terra.ReferenceAsString(kck.ref.Append("key_ring"))
}

// Labels returns a reference to field labels of google_kms_crypto_key.
func (kck kmsCryptoKeyAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kck.ref.Append("labels"))
}

// Name returns a reference to field name of google_kms_crypto_key.
func (kck kmsCryptoKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kck.ref.Append("name"))
}

// Purpose returns a reference to field purpose of google_kms_crypto_key.
func (kck kmsCryptoKeyAttributes) Purpose() terra.StringValue {
	return terra.ReferenceAsString(kck.ref.Append("purpose"))
}

// RotationPeriod returns a reference to field rotation_period of google_kms_crypto_key.
func (kck kmsCryptoKeyAttributes) RotationPeriod() terra.StringValue {
	return terra.ReferenceAsString(kck.ref.Append("rotation_period"))
}

// SkipInitialVersionCreation returns a reference to field skip_initial_version_creation of google_kms_crypto_key.
func (kck kmsCryptoKeyAttributes) SkipInitialVersionCreation() terra.BoolValue {
	return terra.ReferenceAsBool(kck.ref.Append("skip_initial_version_creation"))
}

func (kck kmsCryptoKeyAttributes) Timeouts() kmscryptokey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kmscryptokey.TimeoutsAttributes](kck.ref.Append("timeouts"))
}

func (kck kmsCryptoKeyAttributes) VersionTemplate() terra.ListValue[kmscryptokey.VersionTemplateAttributes] {
	return terra.ReferenceAsList[kmscryptokey.VersionTemplateAttributes](kck.ref.Append("version_template"))
}

type kmsCryptoKeyState struct {
	DestroyScheduledDuration   string                              `json:"destroy_scheduled_duration"`
	Id                         string                              `json:"id"`
	ImportOnly                 bool                                `json:"import_only"`
	KeyRing                    string                              `json:"key_ring"`
	Labels                     map[string]string                   `json:"labels"`
	Name                       string                              `json:"name"`
	Purpose                    string                              `json:"purpose"`
	RotationPeriod             string                              `json:"rotation_period"`
	SkipInitialVersionCreation bool                                `json:"skip_initial_version_creation"`
	Timeouts                   *kmscryptokey.TimeoutsState         `json:"timeouts"`
	VersionTemplate            []kmscryptokey.VersionTemplateState `json:"version_template"`
}
