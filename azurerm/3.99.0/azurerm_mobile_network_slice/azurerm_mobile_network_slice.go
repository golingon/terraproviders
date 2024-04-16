// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_mobile_network_slice

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

// Resource represents the Terraform resource azurerm_mobile_network_slice.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermMobileNetworkSliceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (amns *Resource) Type() string {
	return "azurerm_mobile_network_slice"
}

// LocalName returns the local name for [Resource].
func (amns *Resource) LocalName() string {
	return amns.Name
}

// Configuration returns the configuration (args) for [Resource].
func (amns *Resource) Configuration() interface{} {
	return amns.Args
}

// DependOn is used for other resources to depend on [Resource].
func (amns *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(amns)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (amns *Resource) Dependencies() terra.Dependencies {
	return amns.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (amns *Resource) LifecycleManagement() *terra.Lifecycle {
	return amns.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (amns *Resource) Attributes() azurermMobileNetworkSliceAttributes {
	return azurermMobileNetworkSliceAttributes{ref: terra.ReferenceResource(amns)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (amns *Resource) ImportState(state io.Reader) error {
	amns.state = &azurermMobileNetworkSliceState{}
	if err := json.NewDecoder(state).Decode(amns.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amns.Type(), amns.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (amns *Resource) State() (*azurermMobileNetworkSliceState, bool) {
	return amns.state, amns.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (amns *Resource) StateMust() *azurermMobileNetworkSliceState {
	if amns.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amns.Type(), amns.LocalName()))
	}
	return amns.state
}

// Args contains the configurations for azurerm_mobile_network_slice.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MobileNetworkId: string, required
	MobileNetworkId terra.StringValue `hcl:"mobile_network_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// SingleNetworkSliceSelectionAssistanceInformation: required
	SingleNetworkSliceSelectionAssistanceInformation *SingleNetworkSliceSelectionAssistanceInformation `hcl:"single_network_slice_selection_assistance_information,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermMobileNetworkSliceAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_mobile_network_slice.
func (amns azurermMobileNetworkSliceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amns.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_mobile_network_slice.
func (amns azurermMobileNetworkSliceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amns.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mobile_network_slice.
func (amns azurermMobileNetworkSliceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(amns.ref.Append("location"))
}

// MobileNetworkId returns a reference to field mobile_network_id of azurerm_mobile_network_slice.
func (amns azurermMobileNetworkSliceAttributes) MobileNetworkId() terra.StringValue {
	return terra.ReferenceAsString(amns.ref.Append("mobile_network_id"))
}

// Name returns a reference to field name of azurerm_mobile_network_slice.
func (amns azurermMobileNetworkSliceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amns.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_mobile_network_slice.
func (amns azurermMobileNetworkSliceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](amns.ref.Append("tags"))
}

func (amns azurermMobileNetworkSliceAttributes) SingleNetworkSliceSelectionAssistanceInformation() terra.ListValue[SingleNetworkSliceSelectionAssistanceInformationAttributes] {
	return terra.ReferenceAsList[SingleNetworkSliceSelectionAssistanceInformationAttributes](amns.ref.Append("single_network_slice_selection_assistance_information"))
}

func (amns azurermMobileNetworkSliceAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](amns.ref.Append("timeouts"))
}

type azurermMobileNetworkSliceState struct {
	Description                                      string                                                  `json:"description"`
	Id                                               string                                                  `json:"id"`
	Location                                         string                                                  `json:"location"`
	MobileNetworkId                                  string                                                  `json:"mobile_network_id"`
	Name                                             string                                                  `json:"name"`
	Tags                                             map[string]string                                       `json:"tags"`
	SingleNetworkSliceSelectionAssistanceInformation []SingleNetworkSliceSelectionAssistanceInformationState `json:"single_network_slice_selection_assistance_information"`
	Timeouts                                         *TimeoutsState                                          `json:"timeouts"`
}
