// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	orbitalcontactprofile "github.com/golingon/terraproviders/azurerm/3.49.0/orbitalcontactprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewOrbitalContactProfile(name string, args OrbitalContactProfileArgs) *OrbitalContactProfile {
	return &OrbitalContactProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OrbitalContactProfile)(nil)

type OrbitalContactProfile struct {
	Name  string
	Args  OrbitalContactProfileArgs
	state *orbitalContactProfileState
}

func (ocp *OrbitalContactProfile) Type() string {
	return "azurerm_orbital_contact_profile"
}

func (ocp *OrbitalContactProfile) LocalName() string {
	return ocp.Name
}

func (ocp *OrbitalContactProfile) Configuration() interface{} {
	return ocp.Args
}

func (ocp *OrbitalContactProfile) Attributes() orbitalContactProfileAttributes {
	return orbitalContactProfileAttributes{ref: terra.ReferenceResource(ocp)}
}

func (ocp *OrbitalContactProfile) ImportState(av io.Reader) error {
	ocp.state = &orbitalContactProfileState{}
	if err := json.NewDecoder(av).Decode(ocp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ocp.Type(), ocp.LocalName(), err)
	}
	return nil
}

func (ocp *OrbitalContactProfile) State() (*orbitalContactProfileState, bool) {
	return ocp.state, ocp.state != nil
}

func (ocp *OrbitalContactProfile) StateMust() *orbitalContactProfileState {
	if ocp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ocp.Type(), ocp.LocalName()))
	}
	return ocp.state
}

func (ocp *OrbitalContactProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(ocp)
}

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
	// DependsOn contains resources that OrbitalContactProfile depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type orbitalContactProfileAttributes struct {
	ref terra.Reference
}

func (ocp orbitalContactProfileAttributes) AutoTracking() terra.StringValue {
	return terra.ReferenceString(ocp.ref.Append("auto_tracking"))
}

func (ocp orbitalContactProfileAttributes) EventHubUri() terra.StringValue {
	return terra.ReferenceString(ocp.ref.Append("event_hub_uri"))
}

func (ocp orbitalContactProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ocp.ref.Append("id"))
}

func (ocp orbitalContactProfileAttributes) Location() terra.StringValue {
	return terra.ReferenceString(ocp.ref.Append("location"))
}

func (ocp orbitalContactProfileAttributes) MinimumElevationDegrees() terra.NumberValue {
	return terra.ReferenceNumber(ocp.ref.Append("minimum_elevation_degrees"))
}

func (ocp orbitalContactProfileAttributes) MinimumVariableContactDuration() terra.StringValue {
	return terra.ReferenceString(ocp.ref.Append("minimum_variable_contact_duration"))
}

func (ocp orbitalContactProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ocp.ref.Append("name"))
}

func (ocp orbitalContactProfileAttributes) NetworkConfigurationSubnetId() terra.StringValue {
	return terra.ReferenceString(ocp.ref.Append("network_configuration_subnet_id"))
}

func (ocp orbitalContactProfileAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(ocp.ref.Append("resource_group_name"))
}

func (ocp orbitalContactProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ocp.ref.Append("tags"))
}

func (ocp orbitalContactProfileAttributes) Links() terra.ListValue[orbitalcontactprofile.LinksAttributes] {
	return terra.ReferenceList[orbitalcontactprofile.LinksAttributes](ocp.ref.Append("links"))
}

func (ocp orbitalContactProfileAttributes) Timeouts() orbitalcontactprofile.TimeoutsAttributes {
	return terra.ReferenceSingle[orbitalcontactprofile.TimeoutsAttributes](ocp.ref.Append("timeouts"))
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
