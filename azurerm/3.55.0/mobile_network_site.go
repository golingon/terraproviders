// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mobilenetworksite "github.com/golingon/terraproviders/azurerm/3.55.0/mobilenetworksite"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMobileNetworkSite creates a new instance of [MobileNetworkSite].
func NewMobileNetworkSite(name string, args MobileNetworkSiteArgs) *MobileNetworkSite {
	return &MobileNetworkSite{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MobileNetworkSite)(nil)

// MobileNetworkSite represents the Terraform resource azurerm_mobile_network_site.
type MobileNetworkSite struct {
	Name      string
	Args      MobileNetworkSiteArgs
	state     *mobileNetworkSiteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MobileNetworkSite].
func (mns *MobileNetworkSite) Type() string {
	return "azurerm_mobile_network_site"
}

// LocalName returns the local name for [MobileNetworkSite].
func (mns *MobileNetworkSite) LocalName() string {
	return mns.Name
}

// Configuration returns the configuration (args) for [MobileNetworkSite].
func (mns *MobileNetworkSite) Configuration() interface{} {
	return mns.Args
}

// DependOn is used for other resources to depend on [MobileNetworkSite].
func (mns *MobileNetworkSite) DependOn() terra.Reference {
	return terra.ReferenceResource(mns)
}

// Dependencies returns the list of resources [MobileNetworkSite] depends_on.
func (mns *MobileNetworkSite) Dependencies() terra.Dependencies {
	return mns.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MobileNetworkSite].
func (mns *MobileNetworkSite) LifecycleManagement() *terra.Lifecycle {
	return mns.Lifecycle
}

// Attributes returns the attributes for [MobileNetworkSite].
func (mns *MobileNetworkSite) Attributes() mobileNetworkSiteAttributes {
	return mobileNetworkSiteAttributes{ref: terra.ReferenceResource(mns)}
}

// ImportState imports the given attribute values into [MobileNetworkSite]'s state.
func (mns *MobileNetworkSite) ImportState(av io.Reader) error {
	mns.state = &mobileNetworkSiteState{}
	if err := json.NewDecoder(av).Decode(mns.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mns.Type(), mns.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MobileNetworkSite] has state.
func (mns *MobileNetworkSite) State() (*mobileNetworkSiteState, bool) {
	return mns.state, mns.state != nil
}

// StateMust returns the state for [MobileNetworkSite]. Panics if the state is nil.
func (mns *MobileNetworkSite) StateMust() *mobileNetworkSiteState {
	if mns.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mns.Type(), mns.LocalName()))
	}
	return mns.state
}

// MobileNetworkSiteArgs contains the configurations for azurerm_mobile_network_site.
type MobileNetworkSiteArgs struct {
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
	// Timeouts: optional
	Timeouts *mobilenetworksite.Timeouts `hcl:"timeouts,block"`
}
type mobileNetworkSiteAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mobile_network_site.
func (mns mobileNetworkSiteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mobile_network_site.
func (mns mobileNetworkSiteAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("location"))
}

// MobileNetworkId returns a reference to field mobile_network_id of azurerm_mobile_network_site.
func (mns mobileNetworkSiteAttributes) MobileNetworkId() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("mobile_network_id"))
}

// Name returns a reference to field name of azurerm_mobile_network_site.
func (mns mobileNetworkSiteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mns.ref.Append("name"))
}

// NetworkFunctionIds returns a reference to field network_function_ids of azurerm_mobile_network_site.
func (mns mobileNetworkSiteAttributes) NetworkFunctionIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mns.ref.Append("network_function_ids"))
}

// Tags returns a reference to field tags of azurerm_mobile_network_site.
func (mns mobileNetworkSiteAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mns.ref.Append("tags"))
}

func (mns mobileNetworkSiteAttributes) Timeouts() mobilenetworksite.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mobilenetworksite.TimeoutsAttributes](mns.ref.Append("timeouts"))
}

type mobileNetworkSiteState struct {
	Id                 string                           `json:"id"`
	Location           string                           `json:"location"`
	MobileNetworkId    string                           `json:"mobile_network_id"`
	Name               string                           `json:"name"`
	NetworkFunctionIds []string                         `json:"network_function_ids"`
	Tags               map[string]string                `json:"tags"`
	Timeouts           *mobilenetworksite.TimeoutsState `json:"timeouts"`
}
