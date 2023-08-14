// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	kmskeyring "github.com/golingon/terraproviders/google/4.77.0/kmskeyring"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKmsKeyRing creates a new instance of [KmsKeyRing].
func NewKmsKeyRing(name string, args KmsKeyRingArgs) *KmsKeyRing {
	return &KmsKeyRing{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KmsKeyRing)(nil)

// KmsKeyRing represents the Terraform resource google_kms_key_ring.
type KmsKeyRing struct {
	Name      string
	Args      KmsKeyRingArgs
	state     *kmsKeyRingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KmsKeyRing].
func (kkr *KmsKeyRing) Type() string {
	return "google_kms_key_ring"
}

// LocalName returns the local name for [KmsKeyRing].
func (kkr *KmsKeyRing) LocalName() string {
	return kkr.Name
}

// Configuration returns the configuration (args) for [KmsKeyRing].
func (kkr *KmsKeyRing) Configuration() interface{} {
	return kkr.Args
}

// DependOn is used for other resources to depend on [KmsKeyRing].
func (kkr *KmsKeyRing) DependOn() terra.Reference {
	return terra.ReferenceResource(kkr)
}

// Dependencies returns the list of resources [KmsKeyRing] depends_on.
func (kkr *KmsKeyRing) Dependencies() terra.Dependencies {
	return kkr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KmsKeyRing].
func (kkr *KmsKeyRing) LifecycleManagement() *terra.Lifecycle {
	return kkr.Lifecycle
}

// Attributes returns the attributes for [KmsKeyRing].
func (kkr *KmsKeyRing) Attributes() kmsKeyRingAttributes {
	return kmsKeyRingAttributes{ref: terra.ReferenceResource(kkr)}
}

// ImportState imports the given attribute values into [KmsKeyRing]'s state.
func (kkr *KmsKeyRing) ImportState(av io.Reader) error {
	kkr.state = &kmsKeyRingState{}
	if err := json.NewDecoder(av).Decode(kkr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kkr.Type(), kkr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KmsKeyRing] has state.
func (kkr *KmsKeyRing) State() (*kmsKeyRingState, bool) {
	return kkr.state, kkr.state != nil
}

// StateMust returns the state for [KmsKeyRing]. Panics if the state is nil.
func (kkr *KmsKeyRing) StateMust() *kmsKeyRingState {
	if kkr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kkr.Type(), kkr.LocalName()))
	}
	return kkr.state
}

// KmsKeyRingArgs contains the configurations for google_kms_key_ring.
type KmsKeyRingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *kmskeyring.Timeouts `hcl:"timeouts,block"`
}
type kmsKeyRingAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_kms_key_ring.
func (kkr kmsKeyRingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kkr.ref.Append("id"))
}

// Location returns a reference to field location of google_kms_key_ring.
func (kkr kmsKeyRingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(kkr.ref.Append("location"))
}

// Name returns a reference to field name of google_kms_key_ring.
func (kkr kmsKeyRingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kkr.ref.Append("name"))
}

// Project returns a reference to field project of google_kms_key_ring.
func (kkr kmsKeyRingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(kkr.ref.Append("project"))
}

func (kkr kmsKeyRingAttributes) Timeouts() kmskeyring.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kmskeyring.TimeoutsAttributes](kkr.ref.Append("timeouts"))
}

type kmsKeyRingState struct {
	Id       string                    `json:"id"`
	Location string                    `json:"location"`
	Name     string                    `json:"name"`
	Project  string                    `json:"project"`
	Timeouts *kmskeyring.TimeoutsState `json:"timeouts"`
}
