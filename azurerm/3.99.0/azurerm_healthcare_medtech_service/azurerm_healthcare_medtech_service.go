// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_healthcare_medtech_service

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

// Resource represents the Terraform resource azurerm_healthcare_medtech_service.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermHealthcareMedtechServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ahms *Resource) Type() string {
	return "azurerm_healthcare_medtech_service"
}

// LocalName returns the local name for [Resource].
func (ahms *Resource) LocalName() string {
	return ahms.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ahms *Resource) Configuration() interface{} {
	return ahms.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ahms *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ahms)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ahms *Resource) Dependencies() terra.Dependencies {
	return ahms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ahms *Resource) LifecycleManagement() *terra.Lifecycle {
	return ahms.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ahms *Resource) Attributes() azurermHealthcareMedtechServiceAttributes {
	return azurermHealthcareMedtechServiceAttributes{ref: terra.ReferenceResource(ahms)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ahms *Resource) ImportState(state io.Reader) error {
	ahms.state = &azurermHealthcareMedtechServiceState{}
	if err := json.NewDecoder(state).Decode(ahms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ahms.Type(), ahms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ahms *Resource) State() (*azurermHealthcareMedtechServiceState, bool) {
	return ahms.state, ahms.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ahms *Resource) StateMust() *azurermHealthcareMedtechServiceState {
	if ahms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ahms.Type(), ahms.LocalName()))
	}
	return ahms.state
}

// Args contains the configurations for azurerm_healthcare_medtech_service.
type Args struct {
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
	Identity *Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermHealthcareMedtechServiceAttributes struct {
	ref terra.Reference
}

// DeviceMappingJson returns a reference to field device_mapping_json of azurerm_healthcare_medtech_service.
func (ahms azurermHealthcareMedtechServiceAttributes) DeviceMappingJson() terra.StringValue {
	return terra.ReferenceAsString(ahms.ref.Append("device_mapping_json"))
}

// EventhubConsumerGroupName returns a reference to field eventhub_consumer_group_name of azurerm_healthcare_medtech_service.
func (ahms azurermHealthcareMedtechServiceAttributes) EventhubConsumerGroupName() terra.StringValue {
	return terra.ReferenceAsString(ahms.ref.Append("eventhub_consumer_group_name"))
}

// EventhubName returns a reference to field eventhub_name of azurerm_healthcare_medtech_service.
func (ahms azurermHealthcareMedtechServiceAttributes) EventhubName() terra.StringValue {
	return terra.ReferenceAsString(ahms.ref.Append("eventhub_name"))
}

// EventhubNamespaceName returns a reference to field eventhub_namespace_name of azurerm_healthcare_medtech_service.
func (ahms azurermHealthcareMedtechServiceAttributes) EventhubNamespaceName() terra.StringValue {
	return terra.ReferenceAsString(ahms.ref.Append("eventhub_namespace_name"))
}

// Id returns a reference to field id of azurerm_healthcare_medtech_service.
func (ahms azurermHealthcareMedtechServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ahms.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_healthcare_medtech_service.
func (ahms azurermHealthcareMedtechServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ahms.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_healthcare_medtech_service.
func (ahms azurermHealthcareMedtechServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ahms.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_healthcare_medtech_service.
func (ahms azurermHealthcareMedtechServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ahms.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_healthcare_medtech_service.
func (ahms azurermHealthcareMedtechServiceAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(ahms.ref.Append("workspace_id"))
}

func (ahms azurermHealthcareMedtechServiceAttributes) Identity() terra.ListValue[IdentityAttributes] {
	return terra.ReferenceAsList[IdentityAttributes](ahms.ref.Append("identity"))
}

func (ahms azurermHealthcareMedtechServiceAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ahms.ref.Append("timeouts"))
}

type azurermHealthcareMedtechServiceState struct {
	DeviceMappingJson         string            `json:"device_mapping_json"`
	EventhubConsumerGroupName string            `json:"eventhub_consumer_group_name"`
	EventhubName              string            `json:"eventhub_name"`
	EventhubNamespaceName     string            `json:"eventhub_namespace_name"`
	Id                        string            `json:"id"`
	Location                  string            `json:"location"`
	Name                      string            `json:"name"`
	Tags                      map[string]string `json:"tags"`
	WorkspaceId               string            `json:"workspace_id"`
	Identity                  []IdentityState   `json:"identity"`
	Timeouts                  *TimeoutsState    `json:"timeouts"`
}
