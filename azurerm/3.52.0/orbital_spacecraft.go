// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	orbitalspacecraft "github.com/golingon/terraproviders/azurerm/3.52.0/orbitalspacecraft"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOrbitalSpacecraft creates a new instance of [OrbitalSpacecraft].
func NewOrbitalSpacecraft(name string, args OrbitalSpacecraftArgs) *OrbitalSpacecraft {
	return &OrbitalSpacecraft{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OrbitalSpacecraft)(nil)

// OrbitalSpacecraft represents the Terraform resource azurerm_orbital_spacecraft.
type OrbitalSpacecraft struct {
	Name      string
	Args      OrbitalSpacecraftArgs
	state     *orbitalSpacecraftState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OrbitalSpacecraft].
func (os *OrbitalSpacecraft) Type() string {
	return "azurerm_orbital_spacecraft"
}

// LocalName returns the local name for [OrbitalSpacecraft].
func (os *OrbitalSpacecraft) LocalName() string {
	return os.Name
}

// Configuration returns the configuration (args) for [OrbitalSpacecraft].
func (os *OrbitalSpacecraft) Configuration() interface{} {
	return os.Args
}

// DependOn is used for other resources to depend on [OrbitalSpacecraft].
func (os *OrbitalSpacecraft) DependOn() terra.Reference {
	return terra.ReferenceResource(os)
}

// Dependencies returns the list of resources [OrbitalSpacecraft] depends_on.
func (os *OrbitalSpacecraft) Dependencies() terra.Dependencies {
	return os.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OrbitalSpacecraft].
func (os *OrbitalSpacecraft) LifecycleManagement() *terra.Lifecycle {
	return os.Lifecycle
}

// Attributes returns the attributes for [OrbitalSpacecraft].
func (os *OrbitalSpacecraft) Attributes() orbitalSpacecraftAttributes {
	return orbitalSpacecraftAttributes{ref: terra.ReferenceResource(os)}
}

// ImportState imports the given attribute values into [OrbitalSpacecraft]'s state.
func (os *OrbitalSpacecraft) ImportState(av io.Reader) error {
	os.state = &orbitalSpacecraftState{}
	if err := json.NewDecoder(av).Decode(os.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", os.Type(), os.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OrbitalSpacecraft] has state.
func (os *OrbitalSpacecraft) State() (*orbitalSpacecraftState, bool) {
	return os.state, os.state != nil
}

// StateMust returns the state for [OrbitalSpacecraft]. Panics if the state is nil.
func (os *OrbitalSpacecraft) StateMust() *orbitalSpacecraftState {
	if os.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", os.Type(), os.LocalName()))
	}
	return os.state
}

// OrbitalSpacecraftArgs contains the configurations for azurerm_orbital_spacecraft.
type OrbitalSpacecraftArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NoradId: string, required
	NoradId terra.StringValue `hcl:"norad_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TitleLine: string, required
	TitleLine terra.StringValue `hcl:"title_line,attr" validate:"required"`
	// TwoLineElements: list of string, required
	TwoLineElements terra.ListValue[terra.StringValue] `hcl:"two_line_elements,attr" validate:"required"`
	// Links: min=1
	Links []orbitalspacecraft.Links `hcl:"links,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *orbitalspacecraft.Timeouts `hcl:"timeouts,block"`
}
type orbitalSpacecraftAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_orbital_spacecraft.
func (os orbitalSpacecraftAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_orbital_spacecraft.
func (os orbitalSpacecraftAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_orbital_spacecraft.
func (os orbitalSpacecraftAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("name"))
}

// NoradId returns a reference to field norad_id of azurerm_orbital_spacecraft.
func (os orbitalSpacecraftAttributes) NoradId() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("norad_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_orbital_spacecraft.
func (os orbitalSpacecraftAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_orbital_spacecraft.
func (os orbitalSpacecraftAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](os.ref.Append("tags"))
}

// TitleLine returns a reference to field title_line of azurerm_orbital_spacecraft.
func (os orbitalSpacecraftAttributes) TitleLine() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("title_line"))
}

// TwoLineElements returns a reference to field two_line_elements of azurerm_orbital_spacecraft.
func (os orbitalSpacecraftAttributes) TwoLineElements() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](os.ref.Append("two_line_elements"))
}

func (os orbitalSpacecraftAttributes) Links() terra.ListValue[orbitalspacecraft.LinksAttributes] {
	return terra.ReferenceAsList[orbitalspacecraft.LinksAttributes](os.ref.Append("links"))
}

func (os orbitalSpacecraftAttributes) Timeouts() orbitalspacecraft.TimeoutsAttributes {
	return terra.ReferenceAsSingle[orbitalspacecraft.TimeoutsAttributes](os.ref.Append("timeouts"))
}

type orbitalSpacecraftState struct {
	Id                string                           `json:"id"`
	Location          string                           `json:"location"`
	Name              string                           `json:"name"`
	NoradId           string                           `json:"norad_id"`
	ResourceGroupName string                           `json:"resource_group_name"`
	Tags              map[string]string                `json:"tags"`
	TitleLine         string                           `json:"title_line"`
	TwoLineElements   []string                         `json:"two_line_elements"`
	Links             []orbitalspacecraft.LinksState   `json:"links"`
	Timeouts          *orbitalspacecraft.TimeoutsState `json:"timeouts"`
}
