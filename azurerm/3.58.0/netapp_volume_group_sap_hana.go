// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	netappvolumegroupsaphana "github.com/golingon/terraproviders/azurerm/3.58.0/netappvolumegroupsaphana"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetappVolumeGroupSapHana creates a new instance of [NetappVolumeGroupSapHana].
func NewNetappVolumeGroupSapHana(name string, args NetappVolumeGroupSapHanaArgs) *NetappVolumeGroupSapHana {
	return &NetappVolumeGroupSapHana{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetappVolumeGroupSapHana)(nil)

// NetappVolumeGroupSapHana represents the Terraform resource azurerm_netapp_volume_group_sap_hana.
type NetappVolumeGroupSapHana struct {
	Name      string
	Args      NetappVolumeGroupSapHanaArgs
	state     *netappVolumeGroupSapHanaState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetappVolumeGroupSapHana].
func (nvgsh *NetappVolumeGroupSapHana) Type() string {
	return "azurerm_netapp_volume_group_sap_hana"
}

// LocalName returns the local name for [NetappVolumeGroupSapHana].
func (nvgsh *NetappVolumeGroupSapHana) LocalName() string {
	return nvgsh.Name
}

// Configuration returns the configuration (args) for [NetappVolumeGroupSapHana].
func (nvgsh *NetappVolumeGroupSapHana) Configuration() interface{} {
	return nvgsh.Args
}

// DependOn is used for other resources to depend on [NetappVolumeGroupSapHana].
func (nvgsh *NetappVolumeGroupSapHana) DependOn() terra.Reference {
	return terra.ReferenceResource(nvgsh)
}

// Dependencies returns the list of resources [NetappVolumeGroupSapHana] depends_on.
func (nvgsh *NetappVolumeGroupSapHana) Dependencies() terra.Dependencies {
	return nvgsh.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetappVolumeGroupSapHana].
func (nvgsh *NetappVolumeGroupSapHana) LifecycleManagement() *terra.Lifecycle {
	return nvgsh.Lifecycle
}

// Attributes returns the attributes for [NetappVolumeGroupSapHana].
func (nvgsh *NetappVolumeGroupSapHana) Attributes() netappVolumeGroupSapHanaAttributes {
	return netappVolumeGroupSapHanaAttributes{ref: terra.ReferenceResource(nvgsh)}
}

// ImportState imports the given attribute values into [NetappVolumeGroupSapHana]'s state.
func (nvgsh *NetappVolumeGroupSapHana) ImportState(av io.Reader) error {
	nvgsh.state = &netappVolumeGroupSapHanaState{}
	if err := json.NewDecoder(av).Decode(nvgsh.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nvgsh.Type(), nvgsh.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetappVolumeGroupSapHana] has state.
func (nvgsh *NetappVolumeGroupSapHana) State() (*netappVolumeGroupSapHanaState, bool) {
	return nvgsh.state, nvgsh.state != nil
}

// StateMust returns the state for [NetappVolumeGroupSapHana]. Panics if the state is nil.
func (nvgsh *NetappVolumeGroupSapHana) StateMust() *netappVolumeGroupSapHanaState {
	if nvgsh.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nvgsh.Type(), nvgsh.LocalName()))
	}
	return nvgsh.state
}

// NetappVolumeGroupSapHanaArgs contains the configurations for azurerm_netapp_volume_group_sap_hana.
type NetappVolumeGroupSapHanaArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// ApplicationIdentifier: string, required
	ApplicationIdentifier terra.StringValue `hcl:"application_identifier,attr" validate:"required"`
	// GroupDescription: string, required
	GroupDescription terra.StringValue `hcl:"group_description,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *netappvolumegroupsaphana.Timeouts `hcl:"timeouts,block"`
	// Volume: min=2,max=5
	Volume []netappvolumegroupsaphana.Volume `hcl:"volume,block" validate:"min=2,max=5"`
}
type netappVolumeGroupSapHanaAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_netapp_volume_group_sap_hana.
func (nvgsh netappVolumeGroupSapHanaAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(nvgsh.ref.Append("account_name"))
}

// ApplicationIdentifier returns a reference to field application_identifier of azurerm_netapp_volume_group_sap_hana.
func (nvgsh netappVolumeGroupSapHanaAttributes) ApplicationIdentifier() terra.StringValue {
	return terra.ReferenceAsString(nvgsh.ref.Append("application_identifier"))
}

// GroupDescription returns a reference to field group_description of azurerm_netapp_volume_group_sap_hana.
func (nvgsh netappVolumeGroupSapHanaAttributes) GroupDescription() terra.StringValue {
	return terra.ReferenceAsString(nvgsh.ref.Append("group_description"))
}

// Id returns a reference to field id of azurerm_netapp_volume_group_sap_hana.
func (nvgsh netappVolumeGroupSapHanaAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nvgsh.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_netapp_volume_group_sap_hana.
func (nvgsh netappVolumeGroupSapHanaAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nvgsh.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_netapp_volume_group_sap_hana.
func (nvgsh netappVolumeGroupSapHanaAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nvgsh.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_netapp_volume_group_sap_hana.
func (nvgsh netappVolumeGroupSapHanaAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nvgsh.ref.Append("resource_group_name"))
}

func (nvgsh netappVolumeGroupSapHanaAttributes) Timeouts() netappvolumegroupsaphana.TimeoutsAttributes {
	return terra.ReferenceAsSingle[netappvolumegroupsaphana.TimeoutsAttributes](nvgsh.ref.Append("timeouts"))
}

func (nvgsh netappVolumeGroupSapHanaAttributes) Volume() terra.ListValue[netappvolumegroupsaphana.VolumeAttributes] {
	return terra.ReferenceAsList[netappvolumegroupsaphana.VolumeAttributes](nvgsh.ref.Append("volume"))
}

type netappVolumeGroupSapHanaState struct {
	AccountName           string                                  `json:"account_name"`
	ApplicationIdentifier string                                  `json:"application_identifier"`
	GroupDescription      string                                  `json:"group_description"`
	Id                    string                                  `json:"id"`
	Location              string                                  `json:"location"`
	Name                  string                                  `json:"name"`
	ResourceGroupName     string                                  `json:"resource_group_name"`
	Timeouts              *netappvolumegroupsaphana.TimeoutsState `json:"timeouts"`
	Volume                []netappvolumegroupsaphana.VolumeState  `json:"volume"`
}
