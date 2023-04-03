// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mobilenetworkslice "github.com/golingon/terraproviders/azurerm/3.49.0/mobilenetworkslice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMobileNetworkSlice creates a new instance of [MobileNetworkSlice].
func NewMobileNetworkSlice(name string, args MobileNetworkSliceArgs) *MobileNetworkSlice {
	return &MobileNetworkSlice{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MobileNetworkSlice)(nil)

// MobileNetworkSlice represents the Terraform resource azurerm_mobile_network_slice.
type MobileNetworkSlice struct {
	Name      string
	Args      MobileNetworkSliceArgs
	state     *mobileNetworkSliceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MobileNetworkSlice].
func (mns *MobileNetworkSlice) Type() string {
	return "azurerm_mobile_network_slice"
}

// LocalName returns the local name for [MobileNetworkSlice].
func (mns *MobileNetworkSlice) LocalName() string {
	return mns.Name
}

// Configuration returns the configuration (args) for [MobileNetworkSlice].
func (mns *MobileNetworkSlice) Configuration() interface{} {
	return mns.Args
}

// DependOn is used for other resources to depend on [MobileNetworkSlice].
func (mns *MobileNetworkSlice) DependOn() terra.Reference {
	return terra.ReferenceResource(mns)
}

// Dependencies returns the list of resources [MobileNetworkSlice] depends_on.
func (mns *MobileNetworkSlice) Dependencies() terra.Dependencies {
	return mns.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MobileNetworkSlice].
func (mns *MobileNetworkSlice) LifecycleManagement() *terra.Lifecycle {
	return mns.Lifecycle
}

// Attributes returns the attributes for [MobileNetworkSlice].
func (mns *MobileNetworkSlice) Attributes() mobileNetworkSliceAttributes {
	return mobileNetworkSliceAttributes{ref: terra.ReferenceResource(mns)}
}

// ImportState imports the given attribute values into [MobileNetworkSlice]'s state.
func (mns *MobileNetworkSlice) ImportState(av io.Reader) error {
	mns.state = &mobileNetworkSliceState{}
	if err := json.NewDecoder(av).Decode(mns.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mns.Type(), mns.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MobileNetworkSlice] has state.
func (mns *MobileNetworkSlice) State() (*mobileNetworkSliceState, bool) {
	return mns.state, mns.state != nil
}

// StateMust returns the state for [MobileNetworkSlice]. Panics if the state is nil.
func (mns *MobileNetworkSlice) StateMust() *mobileNetworkSliceState {
	if mns.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mns.Type(), mns.LocalName()))
	}
	return mns.state
}

// MobileNetworkSliceArgs contains the configurations for azurerm_mobile_network_slice.
type MobileNetworkSliceArgs struct {
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
	SingleNetworkSliceSelectionAssistanceInformation *mobilenetworkslice.SingleNetworkSliceSelectionAssistanceInformation `hcl:"single_network_slice_selection_assistance_information,block" validate:"required"`
	// Timeouts: optional
	Timeouts *mobilenetworkslice.Timeouts `hcl:"timeouts,block"`
}
type mobileNetworkSliceAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_mobile_network_slice.
func (mns mobileNetworkSliceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_mobile_network_slice.
func (mns mobileNetworkSliceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mobile_network_slice.
func (mns mobileNetworkSliceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("location"))
}

// MobileNetworkId returns a reference to field mobile_network_id of azurerm_mobile_network_slice.
func (mns mobileNetworkSliceAttributes) MobileNetworkId() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("mobile_network_id"))
}

// Name returns a reference to field name of azurerm_mobile_network_slice.
func (mns mobileNetworkSliceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_mobile_network_slice.
func (mns mobileNetworkSliceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mns.ref.Append("tags"))
}

func (mns mobileNetworkSliceAttributes) SingleNetworkSliceSelectionAssistanceInformation() terra.ListValue[mobilenetworkslice.SingleNetworkSliceSelectionAssistanceInformationAttributes] {
	return terra.ReferenceAsList[mobilenetworkslice.SingleNetworkSliceSelectionAssistanceInformationAttributes](mns.ref.Append("single_network_slice_selection_assistance_information"))
}

func (mns mobileNetworkSliceAttributes) Timeouts() mobilenetworkslice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mobilenetworkslice.TimeoutsAttributes](mns.ref.Append("timeouts"))
}

type mobileNetworkSliceState struct {
	Description                                      string                                                                     `json:"description"`
	Id                                               string                                                                     `json:"id"`
	Location                                         string                                                                     `json:"location"`
	MobileNetworkId                                  string                                                                     `json:"mobile_network_id"`
	Name                                             string                                                                     `json:"name"`
	Tags                                             map[string]string                                                          `json:"tags"`
	SingleNetworkSliceSelectionAssistanceInformation []mobilenetworkslice.SingleNetworkSliceSelectionAssistanceInformationState `json:"single_network_slice_selection_assistance_information"`
	Timeouts                                         *mobilenetworkslice.TimeoutsState                                          `json:"timeouts"`
}
