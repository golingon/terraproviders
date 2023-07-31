// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	healthcareconsentstore "github.com/golingon/terraproviders/googlebeta/4.75.1/healthcareconsentstore"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareConsentStore creates a new instance of [HealthcareConsentStore].
func NewHealthcareConsentStore(name string, args HealthcareConsentStoreArgs) *HealthcareConsentStore {
	return &HealthcareConsentStore{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareConsentStore)(nil)

// HealthcareConsentStore represents the Terraform resource google_healthcare_consent_store.
type HealthcareConsentStore struct {
	Name      string
	Args      HealthcareConsentStoreArgs
	state     *healthcareConsentStoreState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareConsentStore].
func (hcs *HealthcareConsentStore) Type() string {
	return "google_healthcare_consent_store"
}

// LocalName returns the local name for [HealthcareConsentStore].
func (hcs *HealthcareConsentStore) LocalName() string {
	return hcs.Name
}

// Configuration returns the configuration (args) for [HealthcareConsentStore].
func (hcs *HealthcareConsentStore) Configuration() interface{} {
	return hcs.Args
}

// DependOn is used for other resources to depend on [HealthcareConsentStore].
func (hcs *HealthcareConsentStore) DependOn() terra.Reference {
	return terra.ReferenceResource(hcs)
}

// Dependencies returns the list of resources [HealthcareConsentStore] depends_on.
func (hcs *HealthcareConsentStore) Dependencies() terra.Dependencies {
	return hcs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareConsentStore].
func (hcs *HealthcareConsentStore) LifecycleManagement() *terra.Lifecycle {
	return hcs.Lifecycle
}

// Attributes returns the attributes for [HealthcareConsentStore].
func (hcs *HealthcareConsentStore) Attributes() healthcareConsentStoreAttributes {
	return healthcareConsentStoreAttributes{ref: terra.ReferenceResource(hcs)}
}

// ImportState imports the given attribute values into [HealthcareConsentStore]'s state.
func (hcs *HealthcareConsentStore) ImportState(av io.Reader) error {
	hcs.state = &healthcareConsentStoreState{}
	if err := json.NewDecoder(av).Decode(hcs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hcs.Type(), hcs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareConsentStore] has state.
func (hcs *HealthcareConsentStore) State() (*healthcareConsentStoreState, bool) {
	return hcs.state, hcs.state != nil
}

// StateMust returns the state for [HealthcareConsentStore]. Panics if the state is nil.
func (hcs *HealthcareConsentStore) StateMust() *healthcareConsentStoreState {
	if hcs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hcs.Type(), hcs.LocalName()))
	}
	return hcs.state
}

// HealthcareConsentStoreArgs contains the configurations for google_healthcare_consent_store.
type HealthcareConsentStoreArgs struct {
	// Dataset: string, required
	Dataset terra.StringValue `hcl:"dataset,attr" validate:"required"`
	// DefaultConsentTtl: string, optional
	DefaultConsentTtl terra.StringValue `hcl:"default_consent_ttl,attr"`
	// EnableConsentCreateOnUpdate: bool, optional
	EnableConsentCreateOnUpdate terra.BoolValue `hcl:"enable_consent_create_on_update,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *healthcareconsentstore.Timeouts `hcl:"timeouts,block"`
}
type healthcareConsentStoreAttributes struct {
	ref terra.Reference
}

// Dataset returns a reference to field dataset of google_healthcare_consent_store.
func (hcs healthcareConsentStoreAttributes) Dataset() terra.StringValue {
	return terra.ReferenceAsString(hcs.ref.Append("dataset"))
}

// DefaultConsentTtl returns a reference to field default_consent_ttl of google_healthcare_consent_store.
func (hcs healthcareConsentStoreAttributes) DefaultConsentTtl() terra.StringValue {
	return terra.ReferenceAsString(hcs.ref.Append("default_consent_ttl"))
}

// EnableConsentCreateOnUpdate returns a reference to field enable_consent_create_on_update of google_healthcare_consent_store.
func (hcs healthcareConsentStoreAttributes) EnableConsentCreateOnUpdate() terra.BoolValue {
	return terra.ReferenceAsBool(hcs.ref.Append("enable_consent_create_on_update"))
}

// Id returns a reference to field id of google_healthcare_consent_store.
func (hcs healthcareConsentStoreAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hcs.ref.Append("id"))
}

// Labels returns a reference to field labels of google_healthcare_consent_store.
func (hcs healthcareConsentStoreAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hcs.ref.Append("labels"))
}

// Name returns a reference to field name of google_healthcare_consent_store.
func (hcs healthcareConsentStoreAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hcs.ref.Append("name"))
}

func (hcs healthcareConsentStoreAttributes) Timeouts() healthcareconsentstore.TimeoutsAttributes {
	return terra.ReferenceAsSingle[healthcareconsentstore.TimeoutsAttributes](hcs.ref.Append("timeouts"))
}

type healthcareConsentStoreState struct {
	Dataset                     string                                `json:"dataset"`
	DefaultConsentTtl           string                                `json:"default_consent_ttl"`
	EnableConsentCreateOnUpdate bool                                  `json:"enable_consent_create_on_update"`
	Id                          string                                `json:"id"`
	Labels                      map[string]string                     `json:"labels"`
	Name                        string                                `json:"name"`
	Timeouts                    *healthcareconsentstore.TimeoutsState `json:"timeouts"`
}
