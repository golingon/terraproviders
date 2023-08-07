// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	healthcaremedtechservicefhirdestination "github.com/golingon/terraproviders/azurerm/3.68.0/healthcaremedtechservicefhirdestination"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareMedtechServiceFhirDestination creates a new instance of [HealthcareMedtechServiceFhirDestination].
func NewHealthcareMedtechServiceFhirDestination(name string, args HealthcareMedtechServiceFhirDestinationArgs) *HealthcareMedtechServiceFhirDestination {
	return &HealthcareMedtechServiceFhirDestination{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareMedtechServiceFhirDestination)(nil)

// HealthcareMedtechServiceFhirDestination represents the Terraform resource azurerm_healthcare_medtech_service_fhir_destination.
type HealthcareMedtechServiceFhirDestination struct {
	Name      string
	Args      HealthcareMedtechServiceFhirDestinationArgs
	state     *healthcareMedtechServiceFhirDestinationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareMedtechServiceFhirDestination].
func (hmsfd *HealthcareMedtechServiceFhirDestination) Type() string {
	return "azurerm_healthcare_medtech_service_fhir_destination"
}

// LocalName returns the local name for [HealthcareMedtechServiceFhirDestination].
func (hmsfd *HealthcareMedtechServiceFhirDestination) LocalName() string {
	return hmsfd.Name
}

// Configuration returns the configuration (args) for [HealthcareMedtechServiceFhirDestination].
func (hmsfd *HealthcareMedtechServiceFhirDestination) Configuration() interface{} {
	return hmsfd.Args
}

// DependOn is used for other resources to depend on [HealthcareMedtechServiceFhirDestination].
func (hmsfd *HealthcareMedtechServiceFhirDestination) DependOn() terra.Reference {
	return terra.ReferenceResource(hmsfd)
}

// Dependencies returns the list of resources [HealthcareMedtechServiceFhirDestination] depends_on.
func (hmsfd *HealthcareMedtechServiceFhirDestination) Dependencies() terra.Dependencies {
	return hmsfd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareMedtechServiceFhirDestination].
func (hmsfd *HealthcareMedtechServiceFhirDestination) LifecycleManagement() *terra.Lifecycle {
	return hmsfd.Lifecycle
}

// Attributes returns the attributes for [HealthcareMedtechServiceFhirDestination].
func (hmsfd *HealthcareMedtechServiceFhirDestination) Attributes() healthcareMedtechServiceFhirDestinationAttributes {
	return healthcareMedtechServiceFhirDestinationAttributes{ref: terra.ReferenceResource(hmsfd)}
}

// ImportState imports the given attribute values into [HealthcareMedtechServiceFhirDestination]'s state.
func (hmsfd *HealthcareMedtechServiceFhirDestination) ImportState(av io.Reader) error {
	hmsfd.state = &healthcareMedtechServiceFhirDestinationState{}
	if err := json.NewDecoder(av).Decode(hmsfd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hmsfd.Type(), hmsfd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareMedtechServiceFhirDestination] has state.
func (hmsfd *HealthcareMedtechServiceFhirDestination) State() (*healthcareMedtechServiceFhirDestinationState, bool) {
	return hmsfd.state, hmsfd.state != nil
}

// StateMust returns the state for [HealthcareMedtechServiceFhirDestination]. Panics if the state is nil.
func (hmsfd *HealthcareMedtechServiceFhirDestination) StateMust() *healthcareMedtechServiceFhirDestinationState {
	if hmsfd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hmsfd.Type(), hmsfd.LocalName()))
	}
	return hmsfd.state
}

// HealthcareMedtechServiceFhirDestinationArgs contains the configurations for azurerm_healthcare_medtech_service_fhir_destination.
type HealthcareMedtechServiceFhirDestinationArgs struct {
	// DestinationFhirMappingJson: string, required
	DestinationFhirMappingJson terra.StringValue `hcl:"destination_fhir_mapping_json,attr" validate:"required"`
	// DestinationFhirServiceId: string, required
	DestinationFhirServiceId terra.StringValue `hcl:"destination_fhir_service_id,attr" validate:"required"`
	// DestinationIdentityResolutionType: string, required
	DestinationIdentityResolutionType terra.StringValue `hcl:"destination_identity_resolution_type,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MedtechServiceId: string, required
	MedtechServiceId terra.StringValue `hcl:"medtech_service_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *healthcaremedtechservicefhirdestination.Timeouts `hcl:"timeouts,block"`
}
type healthcareMedtechServiceFhirDestinationAttributes struct {
	ref terra.Reference
}

// DestinationFhirMappingJson returns a reference to field destination_fhir_mapping_json of azurerm_healthcare_medtech_service_fhir_destination.
func (hmsfd healthcareMedtechServiceFhirDestinationAttributes) DestinationFhirMappingJson() terra.StringValue {
	return terra.ReferenceAsString(hmsfd.ref.Append("destination_fhir_mapping_json"))
}

// DestinationFhirServiceId returns a reference to field destination_fhir_service_id of azurerm_healthcare_medtech_service_fhir_destination.
func (hmsfd healthcareMedtechServiceFhirDestinationAttributes) DestinationFhirServiceId() terra.StringValue {
	return terra.ReferenceAsString(hmsfd.ref.Append("destination_fhir_service_id"))
}

// DestinationIdentityResolutionType returns a reference to field destination_identity_resolution_type of azurerm_healthcare_medtech_service_fhir_destination.
func (hmsfd healthcareMedtechServiceFhirDestinationAttributes) DestinationIdentityResolutionType() terra.StringValue {
	return terra.ReferenceAsString(hmsfd.ref.Append("destination_identity_resolution_type"))
}

// Id returns a reference to field id of azurerm_healthcare_medtech_service_fhir_destination.
func (hmsfd healthcareMedtechServiceFhirDestinationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hmsfd.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_healthcare_medtech_service_fhir_destination.
func (hmsfd healthcareMedtechServiceFhirDestinationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(hmsfd.ref.Append("location"))
}

// MedtechServiceId returns a reference to field medtech_service_id of azurerm_healthcare_medtech_service_fhir_destination.
func (hmsfd healthcareMedtechServiceFhirDestinationAttributes) MedtechServiceId() terra.StringValue {
	return terra.ReferenceAsString(hmsfd.ref.Append("medtech_service_id"))
}

// Name returns a reference to field name of azurerm_healthcare_medtech_service_fhir_destination.
func (hmsfd healthcareMedtechServiceFhirDestinationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hmsfd.ref.Append("name"))
}

func (hmsfd healthcareMedtechServiceFhirDestinationAttributes) Timeouts() healthcaremedtechservicefhirdestination.TimeoutsAttributes {
	return terra.ReferenceAsSingle[healthcaremedtechservicefhirdestination.TimeoutsAttributes](hmsfd.ref.Append("timeouts"))
}

type healthcareMedtechServiceFhirDestinationState struct {
	DestinationFhirMappingJson        string                                                 `json:"destination_fhir_mapping_json"`
	DestinationFhirServiceId          string                                                 `json:"destination_fhir_service_id"`
	DestinationIdentityResolutionType string                                                 `json:"destination_identity_resolution_type"`
	Id                                string                                                 `json:"id"`
	Location                          string                                                 `json:"location"`
	MedtechServiceId                  string                                                 `json:"medtech_service_id"`
	Name                              string                                                 `json:"name"`
	Timeouts                          *healthcaremedtechservicefhirdestination.TimeoutsState `json:"timeouts"`
}
