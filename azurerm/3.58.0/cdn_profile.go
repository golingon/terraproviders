// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cdnprofile "github.com/golingon/terraproviders/azurerm/3.58.0/cdnprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCdnProfile creates a new instance of [CdnProfile].
func NewCdnProfile(name string, args CdnProfileArgs) *CdnProfile {
	return &CdnProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CdnProfile)(nil)

// CdnProfile represents the Terraform resource azurerm_cdn_profile.
type CdnProfile struct {
	Name      string
	Args      CdnProfileArgs
	state     *cdnProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CdnProfile].
func (cp *CdnProfile) Type() string {
	return "azurerm_cdn_profile"
}

// LocalName returns the local name for [CdnProfile].
func (cp *CdnProfile) LocalName() string {
	return cp.Name
}

// Configuration returns the configuration (args) for [CdnProfile].
func (cp *CdnProfile) Configuration() interface{} {
	return cp.Args
}

// DependOn is used for other resources to depend on [CdnProfile].
func (cp *CdnProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(cp)
}

// Dependencies returns the list of resources [CdnProfile] depends_on.
func (cp *CdnProfile) Dependencies() terra.Dependencies {
	return cp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CdnProfile].
func (cp *CdnProfile) LifecycleManagement() *terra.Lifecycle {
	return cp.Lifecycle
}

// Attributes returns the attributes for [CdnProfile].
func (cp *CdnProfile) Attributes() cdnProfileAttributes {
	return cdnProfileAttributes{ref: terra.ReferenceResource(cp)}
}

// ImportState imports the given attribute values into [CdnProfile]'s state.
func (cp *CdnProfile) ImportState(av io.Reader) error {
	cp.state = &cdnProfileState{}
	if err := json.NewDecoder(av).Decode(cp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cp.Type(), cp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CdnProfile] has state.
func (cp *CdnProfile) State() (*cdnProfileState, bool) {
	return cp.state, cp.state != nil
}

// StateMust returns the state for [CdnProfile]. Panics if the state is nil.
func (cp *CdnProfile) StateMust() *cdnProfileState {
	if cp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cp.Type(), cp.LocalName()))
	}
	return cp.state
}

// CdnProfileArgs contains the configurations for azurerm_cdn_profile.
type CdnProfileArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *cdnprofile.Timeouts `hcl:"timeouts,block"`
}
type cdnProfileAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_cdn_profile.
func (cp cdnProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_cdn_profile.
func (cp cdnProfileAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_cdn_profile.
func (cp cdnProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cdn_profile.
func (cp cdnProfileAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_cdn_profile.
func (cp cdnProfileAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_cdn_profile.
func (cp cdnProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cp.ref.Append("tags"))
}

func (cp cdnProfileAttributes) Timeouts() cdnprofile.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cdnprofile.TimeoutsAttributes](cp.ref.Append("timeouts"))
}

type cdnProfileState struct {
	Id                string                    `json:"id"`
	Location          string                    `json:"location"`
	Name              string                    `json:"name"`
	ResourceGroupName string                    `json:"resource_group_name"`
	Sku               string                    `json:"sku"`
	Tags              map[string]string         `json:"tags"`
	Timeouts          *cdnprofile.TimeoutsState `json:"timeouts"`
}
