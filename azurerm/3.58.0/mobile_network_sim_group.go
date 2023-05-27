// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mobilenetworksimgroup "github.com/golingon/terraproviders/azurerm/3.58.0/mobilenetworksimgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMobileNetworkSimGroup creates a new instance of [MobileNetworkSimGroup].
func NewMobileNetworkSimGroup(name string, args MobileNetworkSimGroupArgs) *MobileNetworkSimGroup {
	return &MobileNetworkSimGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MobileNetworkSimGroup)(nil)

// MobileNetworkSimGroup represents the Terraform resource azurerm_mobile_network_sim_group.
type MobileNetworkSimGroup struct {
	Name      string
	Args      MobileNetworkSimGroupArgs
	state     *mobileNetworkSimGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MobileNetworkSimGroup].
func (mnsg *MobileNetworkSimGroup) Type() string {
	return "azurerm_mobile_network_sim_group"
}

// LocalName returns the local name for [MobileNetworkSimGroup].
func (mnsg *MobileNetworkSimGroup) LocalName() string {
	return mnsg.Name
}

// Configuration returns the configuration (args) for [MobileNetworkSimGroup].
func (mnsg *MobileNetworkSimGroup) Configuration() interface{} {
	return mnsg.Args
}

// DependOn is used for other resources to depend on [MobileNetworkSimGroup].
func (mnsg *MobileNetworkSimGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(mnsg)
}

// Dependencies returns the list of resources [MobileNetworkSimGroup] depends_on.
func (mnsg *MobileNetworkSimGroup) Dependencies() terra.Dependencies {
	return mnsg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MobileNetworkSimGroup].
func (mnsg *MobileNetworkSimGroup) LifecycleManagement() *terra.Lifecycle {
	return mnsg.Lifecycle
}

// Attributes returns the attributes for [MobileNetworkSimGroup].
func (mnsg *MobileNetworkSimGroup) Attributes() mobileNetworkSimGroupAttributes {
	return mobileNetworkSimGroupAttributes{ref: terra.ReferenceResource(mnsg)}
}

// ImportState imports the given attribute values into [MobileNetworkSimGroup]'s state.
func (mnsg *MobileNetworkSimGroup) ImportState(av io.Reader) error {
	mnsg.state = &mobileNetworkSimGroupState{}
	if err := json.NewDecoder(av).Decode(mnsg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mnsg.Type(), mnsg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MobileNetworkSimGroup] has state.
func (mnsg *MobileNetworkSimGroup) State() (*mobileNetworkSimGroupState, bool) {
	return mnsg.state, mnsg.state != nil
}

// StateMust returns the state for [MobileNetworkSimGroup]. Panics if the state is nil.
func (mnsg *MobileNetworkSimGroup) StateMust() *mobileNetworkSimGroupState {
	if mnsg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mnsg.Type(), mnsg.LocalName()))
	}
	return mnsg.state
}

// MobileNetworkSimGroupArgs contains the configurations for azurerm_mobile_network_sim_group.
type MobileNetworkSimGroupArgs struct {
	// EncryptionKeyUrl: string, optional
	EncryptionKeyUrl terra.StringValue `hcl:"encryption_key_url,attr"`
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
	// Identity: optional
	Identity *mobilenetworksimgroup.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *mobilenetworksimgroup.Timeouts `hcl:"timeouts,block"`
}
type mobileNetworkSimGroupAttributes struct {
	ref terra.Reference
}

// EncryptionKeyUrl returns a reference to field encryption_key_url of azurerm_mobile_network_sim_group.
func (mnsg mobileNetworkSimGroupAttributes) EncryptionKeyUrl() terra.StringValue {
	return terra.ReferenceAsString(mnsg.ref.Append("encryption_key_url"))
}

// Id returns a reference to field id of azurerm_mobile_network_sim_group.
func (mnsg mobileNetworkSimGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mnsg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mobile_network_sim_group.
func (mnsg mobileNetworkSimGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mnsg.ref.Append("location"))
}

// MobileNetworkId returns a reference to field mobile_network_id of azurerm_mobile_network_sim_group.
func (mnsg mobileNetworkSimGroupAttributes) MobileNetworkId() terra.StringValue {
	return terra.ReferenceAsString(mnsg.ref.Append("mobile_network_id"))
}

// Name returns a reference to field name of azurerm_mobile_network_sim_group.
func (mnsg mobileNetworkSimGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mnsg.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_mobile_network_sim_group.
func (mnsg mobileNetworkSimGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mnsg.ref.Append("tags"))
}

func (mnsg mobileNetworkSimGroupAttributes) Identity() terra.ListValue[mobilenetworksimgroup.IdentityAttributes] {
	return terra.ReferenceAsList[mobilenetworksimgroup.IdentityAttributes](mnsg.ref.Append("identity"))
}

func (mnsg mobileNetworkSimGroupAttributes) Timeouts() mobilenetworksimgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mobilenetworksimgroup.TimeoutsAttributes](mnsg.ref.Append("timeouts"))
}

type mobileNetworkSimGroupState struct {
	EncryptionKeyUrl string                                `json:"encryption_key_url"`
	Id               string                                `json:"id"`
	Location         string                                `json:"location"`
	MobileNetworkId  string                                `json:"mobile_network_id"`
	Name             string                                `json:"name"`
	Tags             map[string]string                     `json:"tags"`
	Identity         []mobilenetworksimgroup.IdentityState `json:"identity"`
	Timeouts         *mobilenetworksimgroup.TimeoutsState  `json:"timeouts"`
}
