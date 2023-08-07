// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	healthcarehl7v2store "github.com/golingon/terraproviders/googlebeta/4.76.0/healthcarehl7v2store"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareHl7V2Store creates a new instance of [HealthcareHl7V2Store].
func NewHealthcareHl7V2Store(name string, args HealthcareHl7V2StoreArgs) *HealthcareHl7V2Store {
	return &HealthcareHl7V2Store{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareHl7V2Store)(nil)

// HealthcareHl7V2Store represents the Terraform resource google_healthcare_hl7_v2_store.
type HealthcareHl7V2Store struct {
	Name      string
	Args      HealthcareHl7V2StoreArgs
	state     *healthcareHl7V2StoreState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareHl7V2Store].
func (hhvs *HealthcareHl7V2Store) Type() string {
	return "google_healthcare_hl7_v2_store"
}

// LocalName returns the local name for [HealthcareHl7V2Store].
func (hhvs *HealthcareHl7V2Store) LocalName() string {
	return hhvs.Name
}

// Configuration returns the configuration (args) for [HealthcareHl7V2Store].
func (hhvs *HealthcareHl7V2Store) Configuration() interface{} {
	return hhvs.Args
}

// DependOn is used for other resources to depend on [HealthcareHl7V2Store].
func (hhvs *HealthcareHl7V2Store) DependOn() terra.Reference {
	return terra.ReferenceResource(hhvs)
}

// Dependencies returns the list of resources [HealthcareHl7V2Store] depends_on.
func (hhvs *HealthcareHl7V2Store) Dependencies() terra.Dependencies {
	return hhvs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareHl7V2Store].
func (hhvs *HealthcareHl7V2Store) LifecycleManagement() *terra.Lifecycle {
	return hhvs.Lifecycle
}

// Attributes returns the attributes for [HealthcareHl7V2Store].
func (hhvs *HealthcareHl7V2Store) Attributes() healthcareHl7V2StoreAttributes {
	return healthcareHl7V2StoreAttributes{ref: terra.ReferenceResource(hhvs)}
}

// ImportState imports the given attribute values into [HealthcareHl7V2Store]'s state.
func (hhvs *HealthcareHl7V2Store) ImportState(av io.Reader) error {
	hhvs.state = &healthcareHl7V2StoreState{}
	if err := json.NewDecoder(av).Decode(hhvs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hhvs.Type(), hhvs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareHl7V2Store] has state.
func (hhvs *HealthcareHl7V2Store) State() (*healthcareHl7V2StoreState, bool) {
	return hhvs.state, hhvs.state != nil
}

// StateMust returns the state for [HealthcareHl7V2Store]. Panics if the state is nil.
func (hhvs *HealthcareHl7V2Store) StateMust() *healthcareHl7V2StoreState {
	if hhvs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hhvs.Type(), hhvs.LocalName()))
	}
	return hhvs.state
}

// HealthcareHl7V2StoreArgs contains the configurations for google_healthcare_hl7_v2_store.
type HealthcareHl7V2StoreArgs struct {
	// Dataset: string, required
	Dataset terra.StringValue `hcl:"dataset,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NotificationConfig: optional
	NotificationConfig *healthcarehl7v2store.NotificationConfig `hcl:"notification_config,block"`
	// NotificationConfigs: min=0
	NotificationConfigs []healthcarehl7v2store.NotificationConfigs `hcl:"notification_configs,block" validate:"min=0"`
	// ParserConfig: optional
	ParserConfig *healthcarehl7v2store.ParserConfig `hcl:"parser_config,block"`
	// Timeouts: optional
	Timeouts *healthcarehl7v2store.Timeouts `hcl:"timeouts,block"`
}
type healthcareHl7V2StoreAttributes struct {
	ref terra.Reference
}

// Dataset returns a reference to field dataset of google_healthcare_hl7_v2_store.
func (hhvs healthcareHl7V2StoreAttributes) Dataset() terra.StringValue {
	return terra.ReferenceAsString(hhvs.ref.Append("dataset"))
}

// Id returns a reference to field id of google_healthcare_hl7_v2_store.
func (hhvs healthcareHl7V2StoreAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hhvs.ref.Append("id"))
}

// Labels returns a reference to field labels of google_healthcare_hl7_v2_store.
func (hhvs healthcareHl7V2StoreAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hhvs.ref.Append("labels"))
}

// Name returns a reference to field name of google_healthcare_hl7_v2_store.
func (hhvs healthcareHl7V2StoreAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hhvs.ref.Append("name"))
}

// SelfLink returns a reference to field self_link of google_healthcare_hl7_v2_store.
func (hhvs healthcareHl7V2StoreAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(hhvs.ref.Append("self_link"))
}

func (hhvs healthcareHl7V2StoreAttributes) NotificationConfig() terra.ListValue[healthcarehl7v2store.NotificationConfigAttributes] {
	return terra.ReferenceAsList[healthcarehl7v2store.NotificationConfigAttributes](hhvs.ref.Append("notification_config"))
}

func (hhvs healthcareHl7V2StoreAttributes) NotificationConfigs() terra.ListValue[healthcarehl7v2store.NotificationConfigsAttributes] {
	return terra.ReferenceAsList[healthcarehl7v2store.NotificationConfigsAttributes](hhvs.ref.Append("notification_configs"))
}

func (hhvs healthcareHl7V2StoreAttributes) ParserConfig() terra.ListValue[healthcarehl7v2store.ParserConfigAttributes] {
	return terra.ReferenceAsList[healthcarehl7v2store.ParserConfigAttributes](hhvs.ref.Append("parser_config"))
}

func (hhvs healthcareHl7V2StoreAttributes) Timeouts() healthcarehl7v2store.TimeoutsAttributes {
	return terra.ReferenceAsSingle[healthcarehl7v2store.TimeoutsAttributes](hhvs.ref.Append("timeouts"))
}

type healthcareHl7V2StoreState struct {
	Dataset             string                                          `json:"dataset"`
	Id                  string                                          `json:"id"`
	Labels              map[string]string                               `json:"labels"`
	Name                string                                          `json:"name"`
	SelfLink            string                                          `json:"self_link"`
	NotificationConfig  []healthcarehl7v2store.NotificationConfigState  `json:"notification_config"`
	NotificationConfigs []healthcarehl7v2store.NotificationConfigsState `json:"notification_configs"`
	ParserConfig        []healthcarehl7v2store.ParserConfigState        `json:"parser_config"`
	Timeouts            *healthcarehl7v2store.TimeoutsState             `json:"timeouts"`
}
