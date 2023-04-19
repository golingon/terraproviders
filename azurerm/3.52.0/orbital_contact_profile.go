// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	orbitalcontactprofile "github.com/golingon/terraproviders/azurerm/3.52.0/orbitalcontactprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOrbitalContactProfile creates a new instance of [OrbitalContactProfile].
func NewOrbitalContactProfile(name string, args OrbitalContactProfileArgs) *OrbitalContactProfile {
	return &OrbitalContactProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OrbitalContactProfile)(nil)

// OrbitalContactProfile represents the Terraform resource azurerm_orbital_contact_profile.
type OrbitalContactProfile struct {
	Name      string
	Args      OrbitalContactProfileArgs
	state     *orbitalContactProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OrbitalContactProfile].
func (ocp *OrbitalContactProfile) Type() string {
	return "azurerm_orbital_contact_profile"
}

// LocalName returns the local name for [OrbitalContactProfile].
func (ocp *OrbitalContactProfile) LocalName() string {
	return ocp.Name
}

// Configuration returns the configuration (args) for [OrbitalContactProfile].
func (ocp *OrbitalContactProfile) Configuration() interface{} {
	return ocp.Args
}

// DependOn is used for other resources to depend on [OrbitalContactProfile].
func (ocp *OrbitalContactProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(ocp)
}

// Dependencies returns the list of resources [OrbitalContactProfile] depends_on.
func (ocp *OrbitalContactProfile) Dependencies() terra.Dependencies {
	return ocp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OrbitalContactProfile].
func (ocp *OrbitalContactProfile) LifecycleManagement() *terra.Lifecycle {
	return ocp.Lifecycle
}

// Attributes returns the attributes for [OrbitalContactProfile].
func (ocp *OrbitalContactProfile) Attributes() orbitalContactProfileAttributes {
	return orbitalContactProfileAttributes{ref: terra.ReferenceResource(ocp)}
}

// ImportState imports the given attribute values into [OrbitalContactProfile]'s state.
func (ocp *OrbitalContactProfile) ImportState(av io.Reader) error {
	ocp.state = &orbitalContactProfileState{}
	if err := json.NewDecoder(av).Decode(ocp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ocp.Type(), ocp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OrbitalContactProfile] has state.
func (ocp *OrbitalContactProfile) State() (*orbitalContactProfileState, bool) {
	return ocp.state, ocp.state != nil
}

// StateMust returns the state for [OrbitalContactProfile]. Panics if the state is nil.
func (ocp *OrbitalContactProfile) StateMust() *orbitalContactProfileState {
	if ocp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ocp.Type(), ocp.LocalName()))
	}
	return ocp.state
}

// OrbitalContactProfileArgs contains the configurations for azurerm_orbital_contact_profile.
type OrbitalContactProfileArgs struct {
	// AutoTracking: string, required
	AutoTracking terra.StringValue `hcl:"auto_tracking,attr" validate:"required"`
	// EventHubUri: string, optional
	EventHubUri terra.StringValue `hcl:"event_hub_uri,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MinimumElevationDegrees: number, optional
	MinimumElevationDegrees terra.NumberValue `hcl:"minimum_elevation_degrees,attr"`
	// MinimumVariableContactDuration: string, required
	MinimumVariableContactDuration terra.StringValue `hcl:"minimum_variable_contact_duration,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkConfigurationSubnetId: string, required
	NetworkConfigurationSubnetId terra.StringValue `hcl:"network_configuration_subnet_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Links: min=1
	Links []orbitalcontactprofile.Links `hcl:"links,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *orbitalcontactprofile.Timeouts `hcl:"timeouts,block"`
}
type orbitalContactProfileAttributes struct {
	ref terra.Reference
}

// AutoTracking returns a reference to field auto_tracking of azurerm_orbital_contact_profile.
func (ocp orbitalContactProfileAttributes) AutoTracking() terra.StringValue {
	return terra.ReferenceAsString(ocp.ref.Append("auto_tracking"))
}

// EventHubUri returns a reference to field event_hub_uri of azurerm_orbital_contact_profile.
func (ocp orbitalContactProfileAttributes) EventHubUri() terra.StringValue {
	return terra.ReferenceAsString(ocp.ref.Append("event_hub_uri"))
}

// Id returns a reference to field id of azurerm_orbital_contact_profile.
func (ocp orbitalContactProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ocp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_orbital_contact_profile.
func (ocp orbitalContactProfileAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ocp.ref.Append("location"))
}

// MinimumElevationDegrees returns a reference to field minimum_elevation_degrees of azurerm_orbital_contact_profile.
func (ocp orbitalContactProfileAttributes) MinimumElevationDegrees() terra.NumberValue {
	return terra.ReferenceAsNumber(ocp.ref.Append("minimum_elevation_degrees"))
}

// MinimumVariableContactDuration returns a reference to field minimum_variable_contact_duration of azurerm_orbital_contact_profile.
func (ocp orbitalContactProfileAttributes) MinimumVariableContactDuration() terra.StringValue {
	return terra.ReferenceAsString(ocp.ref.Append("minimum_variable_contact_duration"))
}

// Name returns a reference to field name of azurerm_orbital_contact_profile.
func (ocp orbitalContactProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ocp.ref.Append("name"))
}

// NetworkConfigurationSubnetId returns a reference to field network_configuration_subnet_id of azurerm_orbital_contact_profile.
func (ocp orbitalContactProfileAttributes) NetworkConfigurationSubnetId() terra.StringValue {
	return terra.ReferenceAsString(ocp.ref.Append("network_configuration_subnet_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_orbital_contact_profile.
func (ocp orbitalContactProfileAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ocp.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_orbital_contact_profile.
func (ocp orbitalContactProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ocp.ref.Append("tags"))
}

func (ocp orbitalContactProfileAttributes) Links() terra.ListValue[orbitalcontactprofile.LinksAttributes] {
	return terra.ReferenceAsList[orbitalcontactprofile.LinksAttributes](ocp.ref.Append("links"))
}

func (ocp orbitalContactProfileAttributes) Timeouts() orbitalcontactprofile.TimeoutsAttributes {
	return terra.ReferenceAsSingle[orbitalcontactprofile.TimeoutsAttributes](ocp.ref.Append("timeouts"))
}

type orbitalContactProfileState struct {
	AutoTracking                   string                               `json:"auto_tracking"`
	EventHubUri                    string                               `json:"event_hub_uri"`
	Id                             string                               `json:"id"`
	Location                       string                               `json:"location"`
	MinimumElevationDegrees        float64                              `json:"minimum_elevation_degrees"`
	MinimumVariableContactDuration string                               `json:"minimum_variable_contact_duration"`
	Name                           string                               `json:"name"`
	NetworkConfigurationSubnetId   string                               `json:"network_configuration_subnet_id"`
	ResourceGroupName              string                               `json:"resource_group_name"`
	Tags                           map[string]string                    `json:"tags"`
	Links                          []orbitalcontactprofile.LinksState   `json:"links"`
	Timeouts                       *orbitalcontactprofile.TimeoutsState `json:"timeouts"`
}
