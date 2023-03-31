// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mobilenetworksimgroup "github.com/golingon/terraproviders/azurerm/3.49.0/mobilenetworksimgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewMobileNetworkSimGroup(name string, args MobileNetworkSimGroupArgs) *MobileNetworkSimGroup {
	return &MobileNetworkSimGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MobileNetworkSimGroup)(nil)

type MobileNetworkSimGroup struct {
	Name  string
	Args  MobileNetworkSimGroupArgs
	state *mobileNetworkSimGroupState
}

func (mnsg *MobileNetworkSimGroup) Type() string {
	return "azurerm_mobile_network_sim_group"
}

func (mnsg *MobileNetworkSimGroup) LocalName() string {
	return mnsg.Name
}

func (mnsg *MobileNetworkSimGroup) Configuration() interface{} {
	return mnsg.Args
}

func (mnsg *MobileNetworkSimGroup) Attributes() mobileNetworkSimGroupAttributes {
	return mobileNetworkSimGroupAttributes{ref: terra.ReferenceResource(mnsg)}
}

func (mnsg *MobileNetworkSimGroup) ImportState(av io.Reader) error {
	mnsg.state = &mobileNetworkSimGroupState{}
	if err := json.NewDecoder(av).Decode(mnsg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mnsg.Type(), mnsg.LocalName(), err)
	}
	return nil
}

func (mnsg *MobileNetworkSimGroup) State() (*mobileNetworkSimGroupState, bool) {
	return mnsg.state, mnsg.state != nil
}

func (mnsg *MobileNetworkSimGroup) StateMust() *mobileNetworkSimGroupState {
	if mnsg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mnsg.Type(), mnsg.LocalName()))
	}
	return mnsg.state
}

func (mnsg *MobileNetworkSimGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(mnsg)
}

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
	// DependsOn contains resources that MobileNetworkSimGroup depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type mobileNetworkSimGroupAttributes struct {
	ref terra.Reference
}

func (mnsg mobileNetworkSimGroupAttributes) EncryptionKeyUrl() terra.StringValue {
	return terra.ReferenceString(mnsg.ref.Append("encryption_key_url"))
}

func (mnsg mobileNetworkSimGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(mnsg.ref.Append("id"))
}

func (mnsg mobileNetworkSimGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceString(mnsg.ref.Append("location"))
}

func (mnsg mobileNetworkSimGroupAttributes) MobileNetworkId() terra.StringValue {
	return terra.ReferenceString(mnsg.ref.Append("mobile_network_id"))
}

func (mnsg mobileNetworkSimGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(mnsg.ref.Append("name"))
}

func (mnsg mobileNetworkSimGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](mnsg.ref.Append("tags"))
}

func (mnsg mobileNetworkSimGroupAttributes) Identity() terra.ListValue[mobilenetworksimgroup.IdentityAttributes] {
	return terra.ReferenceList[mobilenetworksimgroup.IdentityAttributes](mnsg.ref.Append("identity"))
}

func (mnsg mobileNetworkSimGroupAttributes) Timeouts() mobilenetworksimgroup.TimeoutsAttributes {
	return terra.ReferenceSingle[mobilenetworksimgroup.TimeoutsAttributes](mnsg.ref.Append("timeouts"))
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
