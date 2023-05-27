// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	healthcaremedtechservice "github.com/golingon/terraproviders/azurerm/3.58.0/healthcaremedtechservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareMedtechService creates a new instance of [HealthcareMedtechService].
func NewHealthcareMedtechService(name string, args HealthcareMedtechServiceArgs) *HealthcareMedtechService {
	return &HealthcareMedtechService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareMedtechService)(nil)

// HealthcareMedtechService represents the Terraform resource azurerm_healthcare_medtech_service.
type HealthcareMedtechService struct {
	Name      string
	Args      HealthcareMedtechServiceArgs
	state     *healthcareMedtechServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareMedtechService].
func (hms *HealthcareMedtechService) Type() string {
	return "azurerm_healthcare_medtech_service"
}

// LocalName returns the local name for [HealthcareMedtechService].
func (hms *HealthcareMedtechService) LocalName() string {
	return hms.Name
}

// Configuration returns the configuration (args) for [HealthcareMedtechService].
func (hms *HealthcareMedtechService) Configuration() interface{} {
	return hms.Args
}

// DependOn is used for other resources to depend on [HealthcareMedtechService].
func (hms *HealthcareMedtechService) DependOn() terra.Reference {
	return terra.ReferenceResource(hms)
}

// Dependencies returns the list of resources [HealthcareMedtechService] depends_on.
func (hms *HealthcareMedtechService) Dependencies() terra.Dependencies {
	return hms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareMedtechService].
func (hms *HealthcareMedtechService) LifecycleManagement() *terra.Lifecycle {
	return hms.Lifecycle
}

// Attributes returns the attributes for [HealthcareMedtechService].
func (hms *HealthcareMedtechService) Attributes() healthcareMedtechServiceAttributes {
	return healthcareMedtechServiceAttributes{ref: terra.ReferenceResource(hms)}
}

// ImportState imports the given attribute values into [HealthcareMedtechService]'s state.
func (hms *HealthcareMedtechService) ImportState(av io.Reader) error {
	hms.state = &healthcareMedtechServiceState{}
	if err := json.NewDecoder(av).Decode(hms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hms.Type(), hms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareMedtechService] has state.
func (hms *HealthcareMedtechService) State() (*healthcareMedtechServiceState, bool) {
	return hms.state, hms.state != nil
}

// StateMust returns the state for [HealthcareMedtechService]. Panics if the state is nil.
func (hms *HealthcareMedtechService) StateMust() *healthcareMedtechServiceState {
	if hms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hms.Type(), hms.LocalName()))
	}
	return hms.state
}

// HealthcareMedtechServiceArgs contains the configurations for azurerm_healthcare_medtech_service.
type HealthcareMedtechServiceArgs struct {
	// DeviceMappingJson: string, required
	DeviceMappingJson terra.StringValue `hcl:"device_mapping_json,attr" validate:"required"`
	// EventhubConsumerGroupName: string, required
	EventhubConsumerGroupName terra.StringValue `hcl:"eventhub_consumer_group_name,attr" validate:"required"`
	// EventhubName: string, required
	EventhubName terra.StringValue `hcl:"eventhub_name,attr" validate:"required"`
	// EventhubNamespaceName: string, required
	EventhubNamespaceName terra.StringValue `hcl:"eventhub_namespace_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Identity: optional
	Identity *healthcaremedtechservice.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *healthcaremedtechservice.Timeouts `hcl:"timeouts,block"`
}
type healthcareMedtechServiceAttributes struct {
	ref terra.Reference
}

// DeviceMappingJson returns a reference to field device_mapping_json of azurerm_healthcare_medtech_service.
func (hms healthcareMedtechServiceAttributes) DeviceMappingJson() terra.StringValue {
	return terra.ReferenceAsString(hms.ref.Append("device_mapping_json"))
}

// EventhubConsumerGroupName returns a reference to field eventhub_consumer_group_name of azurerm_healthcare_medtech_service.
func (hms healthcareMedtechServiceAttributes) EventhubConsumerGroupName() terra.StringValue {
	return terra.ReferenceAsString(hms.ref.Append("eventhub_consumer_group_name"))
}

// EventhubName returns a reference to field eventhub_name of azurerm_healthcare_medtech_service.
func (hms healthcareMedtechServiceAttributes) EventhubName() terra.StringValue {
	return terra.ReferenceAsString(hms.ref.Append("eventhub_name"))
}

// EventhubNamespaceName returns a reference to field eventhub_namespace_name of azurerm_healthcare_medtech_service.
func (hms healthcareMedtechServiceAttributes) EventhubNamespaceName() terra.StringValue {
	return terra.ReferenceAsString(hms.ref.Append("eventhub_namespace_name"))
}

// Id returns a reference to field id of azurerm_healthcare_medtech_service.
func (hms healthcareMedtechServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hms.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_healthcare_medtech_service.
func (hms healthcareMedtechServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(hms.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_healthcare_medtech_service.
func (hms healthcareMedtechServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hms.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_healthcare_medtech_service.
func (hms healthcareMedtechServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hms.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_healthcare_medtech_service.
func (hms healthcareMedtechServiceAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(hms.ref.Append("workspace_id"))
}

func (hms healthcareMedtechServiceAttributes) Identity() terra.ListValue[healthcaremedtechservice.IdentityAttributes] {
	return terra.ReferenceAsList[healthcaremedtechservice.IdentityAttributes](hms.ref.Append("identity"))
}

func (hms healthcareMedtechServiceAttributes) Timeouts() healthcaremedtechservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[healthcaremedtechservice.TimeoutsAttributes](hms.ref.Append("timeouts"))
}

type healthcareMedtechServiceState struct {
	DeviceMappingJson         string                                   `json:"device_mapping_json"`
	EventhubConsumerGroupName string                                   `json:"eventhub_consumer_group_name"`
	EventhubName              string                                   `json:"eventhub_name"`
	EventhubNamespaceName     string                                   `json:"eventhub_namespace_name"`
	Id                        string                                   `json:"id"`
	Location                  string                                   `json:"location"`
	Name                      string                                   `json:"name"`
	Tags                      map[string]string                        `json:"tags"`
	WorkspaceId               string                                   `json:"workspace_id"`
	Identity                  []healthcaremedtechservice.IdentityState `json:"identity"`
	Timeouts                  *healthcaremedtechservice.TimeoutsState  `json:"timeouts"`
}
