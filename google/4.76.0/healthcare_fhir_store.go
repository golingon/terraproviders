// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	healthcarefhirstore "github.com/golingon/terraproviders/google/4.76.0/healthcarefhirstore"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareFhirStore creates a new instance of [HealthcareFhirStore].
func NewHealthcareFhirStore(name string, args HealthcareFhirStoreArgs) *HealthcareFhirStore {
	return &HealthcareFhirStore{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareFhirStore)(nil)

// HealthcareFhirStore represents the Terraform resource google_healthcare_fhir_store.
type HealthcareFhirStore struct {
	Name      string
	Args      HealthcareFhirStoreArgs
	state     *healthcareFhirStoreState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareFhirStore].
func (hfs *HealthcareFhirStore) Type() string {
	return "google_healthcare_fhir_store"
}

// LocalName returns the local name for [HealthcareFhirStore].
func (hfs *HealthcareFhirStore) LocalName() string {
	return hfs.Name
}

// Configuration returns the configuration (args) for [HealthcareFhirStore].
func (hfs *HealthcareFhirStore) Configuration() interface{} {
	return hfs.Args
}

// DependOn is used for other resources to depend on [HealthcareFhirStore].
func (hfs *HealthcareFhirStore) DependOn() terra.Reference {
	return terra.ReferenceResource(hfs)
}

// Dependencies returns the list of resources [HealthcareFhirStore] depends_on.
func (hfs *HealthcareFhirStore) Dependencies() terra.Dependencies {
	return hfs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareFhirStore].
func (hfs *HealthcareFhirStore) LifecycleManagement() *terra.Lifecycle {
	return hfs.Lifecycle
}

// Attributes returns the attributes for [HealthcareFhirStore].
func (hfs *HealthcareFhirStore) Attributes() healthcareFhirStoreAttributes {
	return healthcareFhirStoreAttributes{ref: terra.ReferenceResource(hfs)}
}

// ImportState imports the given attribute values into [HealthcareFhirStore]'s state.
func (hfs *HealthcareFhirStore) ImportState(av io.Reader) error {
	hfs.state = &healthcareFhirStoreState{}
	if err := json.NewDecoder(av).Decode(hfs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hfs.Type(), hfs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareFhirStore] has state.
func (hfs *HealthcareFhirStore) State() (*healthcareFhirStoreState, bool) {
	return hfs.state, hfs.state != nil
}

// StateMust returns the state for [HealthcareFhirStore]. Panics if the state is nil.
func (hfs *HealthcareFhirStore) StateMust() *healthcareFhirStoreState {
	if hfs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hfs.Type(), hfs.LocalName()))
	}
	return hfs.state
}

// HealthcareFhirStoreArgs contains the configurations for google_healthcare_fhir_store.
type HealthcareFhirStoreArgs struct {
	// ComplexDataTypeReferenceParsing: string, optional
	ComplexDataTypeReferenceParsing terra.StringValue `hcl:"complex_data_type_reference_parsing,attr"`
	// Dataset: string, required
	Dataset terra.StringValue `hcl:"dataset,attr" validate:"required"`
	// DisableReferentialIntegrity: bool, optional
	DisableReferentialIntegrity terra.BoolValue `hcl:"disable_referential_integrity,attr"`
	// DisableResourceVersioning: bool, optional
	DisableResourceVersioning terra.BoolValue `hcl:"disable_resource_versioning,attr"`
	// EnableHistoryImport: bool, optional
	EnableHistoryImport terra.BoolValue `hcl:"enable_history_import,attr"`
	// EnableUpdateCreate: bool, optional
	EnableUpdateCreate terra.BoolValue `hcl:"enable_update_create,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// NotificationConfig: optional
	NotificationConfig *healthcarefhirstore.NotificationConfig `hcl:"notification_config,block"`
	// StreamConfigs: min=0
	StreamConfigs []healthcarefhirstore.StreamConfigs `hcl:"stream_configs,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *healthcarefhirstore.Timeouts `hcl:"timeouts,block"`
}
type healthcareFhirStoreAttributes struct {
	ref terra.Reference
}

// ComplexDataTypeReferenceParsing returns a reference to field complex_data_type_reference_parsing of google_healthcare_fhir_store.
func (hfs healthcareFhirStoreAttributes) ComplexDataTypeReferenceParsing() terra.StringValue {
	return terra.ReferenceAsString(hfs.ref.Append("complex_data_type_reference_parsing"))
}

// Dataset returns a reference to field dataset of google_healthcare_fhir_store.
func (hfs healthcareFhirStoreAttributes) Dataset() terra.StringValue {
	return terra.ReferenceAsString(hfs.ref.Append("dataset"))
}

// DisableReferentialIntegrity returns a reference to field disable_referential_integrity of google_healthcare_fhir_store.
func (hfs healthcareFhirStoreAttributes) DisableReferentialIntegrity() terra.BoolValue {
	return terra.ReferenceAsBool(hfs.ref.Append("disable_referential_integrity"))
}

// DisableResourceVersioning returns a reference to field disable_resource_versioning of google_healthcare_fhir_store.
func (hfs healthcareFhirStoreAttributes) DisableResourceVersioning() terra.BoolValue {
	return terra.ReferenceAsBool(hfs.ref.Append("disable_resource_versioning"))
}

// EnableHistoryImport returns a reference to field enable_history_import of google_healthcare_fhir_store.
func (hfs healthcareFhirStoreAttributes) EnableHistoryImport() terra.BoolValue {
	return terra.ReferenceAsBool(hfs.ref.Append("enable_history_import"))
}

// EnableUpdateCreate returns a reference to field enable_update_create of google_healthcare_fhir_store.
func (hfs healthcareFhirStoreAttributes) EnableUpdateCreate() terra.BoolValue {
	return terra.ReferenceAsBool(hfs.ref.Append("enable_update_create"))
}

// Id returns a reference to field id of google_healthcare_fhir_store.
func (hfs healthcareFhirStoreAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hfs.ref.Append("id"))
}

// Labels returns a reference to field labels of google_healthcare_fhir_store.
func (hfs healthcareFhirStoreAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hfs.ref.Append("labels"))
}

// Name returns a reference to field name of google_healthcare_fhir_store.
func (hfs healthcareFhirStoreAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hfs.ref.Append("name"))
}

// SelfLink returns a reference to field self_link of google_healthcare_fhir_store.
func (hfs healthcareFhirStoreAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(hfs.ref.Append("self_link"))
}

// Version returns a reference to field version of google_healthcare_fhir_store.
func (hfs healthcareFhirStoreAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(hfs.ref.Append("version"))
}

func (hfs healthcareFhirStoreAttributes) NotificationConfig() terra.ListValue[healthcarefhirstore.NotificationConfigAttributes] {
	return terra.ReferenceAsList[healthcarefhirstore.NotificationConfigAttributes](hfs.ref.Append("notification_config"))
}

func (hfs healthcareFhirStoreAttributes) StreamConfigs() terra.ListValue[healthcarefhirstore.StreamConfigsAttributes] {
	return terra.ReferenceAsList[healthcarefhirstore.StreamConfigsAttributes](hfs.ref.Append("stream_configs"))
}

func (hfs healthcareFhirStoreAttributes) Timeouts() healthcarefhirstore.TimeoutsAttributes {
	return terra.ReferenceAsSingle[healthcarefhirstore.TimeoutsAttributes](hfs.ref.Append("timeouts"))
}

type healthcareFhirStoreState struct {
	ComplexDataTypeReferenceParsing string                                        `json:"complex_data_type_reference_parsing"`
	Dataset                         string                                        `json:"dataset"`
	DisableReferentialIntegrity     bool                                          `json:"disable_referential_integrity"`
	DisableResourceVersioning       bool                                          `json:"disable_resource_versioning"`
	EnableHistoryImport             bool                                          `json:"enable_history_import"`
	EnableUpdateCreate              bool                                          `json:"enable_update_create"`
	Id                              string                                        `json:"id"`
	Labels                          map[string]string                             `json:"labels"`
	Name                            string                                        `json:"name"`
	SelfLink                        string                                        `json:"self_link"`
	Version                         string                                        `json:"version"`
	NotificationConfig              []healthcarefhirstore.NotificationConfigState `json:"notification_config"`
	StreamConfigs                   []healthcarefhirstore.StreamConfigsState      `json:"stream_configs"`
	Timeouts                        *healthcarefhirstore.TimeoutsState            `json:"timeouts"`
}
