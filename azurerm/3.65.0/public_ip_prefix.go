// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	publicipprefix "github.com/golingon/terraproviders/azurerm/3.65.0/publicipprefix"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPublicIpPrefix creates a new instance of [PublicIpPrefix].
func NewPublicIpPrefix(name string, args PublicIpPrefixArgs) *PublicIpPrefix {
	return &PublicIpPrefix{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PublicIpPrefix)(nil)

// PublicIpPrefix represents the Terraform resource azurerm_public_ip_prefix.
type PublicIpPrefix struct {
	Name      string
	Args      PublicIpPrefixArgs
	state     *publicIpPrefixState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PublicIpPrefix].
func (pip *PublicIpPrefix) Type() string {
	return "azurerm_public_ip_prefix"
}

// LocalName returns the local name for [PublicIpPrefix].
func (pip *PublicIpPrefix) LocalName() string {
	return pip.Name
}

// Configuration returns the configuration (args) for [PublicIpPrefix].
func (pip *PublicIpPrefix) Configuration() interface{} {
	return pip.Args
}

// DependOn is used for other resources to depend on [PublicIpPrefix].
func (pip *PublicIpPrefix) DependOn() terra.Reference {
	return terra.ReferenceResource(pip)
}

// Dependencies returns the list of resources [PublicIpPrefix] depends_on.
func (pip *PublicIpPrefix) Dependencies() terra.Dependencies {
	return pip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PublicIpPrefix].
func (pip *PublicIpPrefix) LifecycleManagement() *terra.Lifecycle {
	return pip.Lifecycle
}

// Attributes returns the attributes for [PublicIpPrefix].
func (pip *PublicIpPrefix) Attributes() publicIpPrefixAttributes {
	return publicIpPrefixAttributes{ref: terra.ReferenceResource(pip)}
}

// ImportState imports the given attribute values into [PublicIpPrefix]'s state.
func (pip *PublicIpPrefix) ImportState(av io.Reader) error {
	pip.state = &publicIpPrefixState{}
	if err := json.NewDecoder(av).Decode(pip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pip.Type(), pip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PublicIpPrefix] has state.
func (pip *PublicIpPrefix) State() (*publicIpPrefixState, bool) {
	return pip.state, pip.state != nil
}

// StateMust returns the state for [PublicIpPrefix]. Panics if the state is nil.
func (pip *PublicIpPrefix) StateMust() *publicIpPrefixState {
	if pip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pip.Type(), pip.LocalName()))
	}
	return pip.state
}

// PublicIpPrefixArgs contains the configurations for azurerm_public_ip_prefix.
type PublicIpPrefixArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpVersion: string, optional
	IpVersion terra.StringValue `hcl:"ip_version,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrefixLength: number, optional
	PrefixLength terra.NumberValue `hcl:"prefix_length,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, optional
	Sku terra.StringValue `hcl:"sku,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Zones: set of string, optional
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr"`
	// Timeouts: optional
	Timeouts *publicipprefix.Timeouts `hcl:"timeouts,block"`
}
type publicIpPrefixAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_public_ip_prefix.
func (pip publicIpPrefixAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("id"))
}

// IpPrefix returns a reference to field ip_prefix of azurerm_public_ip_prefix.
func (pip publicIpPrefixAttributes) IpPrefix() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("ip_prefix"))
}

// IpVersion returns a reference to field ip_version of azurerm_public_ip_prefix.
func (pip publicIpPrefixAttributes) IpVersion() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("ip_version"))
}

// Location returns a reference to field location of azurerm_public_ip_prefix.
func (pip publicIpPrefixAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_public_ip_prefix.
func (pip publicIpPrefixAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("name"))
}

// PrefixLength returns a reference to field prefix_length of azurerm_public_ip_prefix.
func (pip publicIpPrefixAttributes) PrefixLength() terra.NumberValue {
	return terra.ReferenceAsNumber(pip.ref.Append("prefix_length"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_public_ip_prefix.
func (pip publicIpPrefixAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_public_ip_prefix.
func (pip publicIpPrefixAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_public_ip_prefix.
func (pip publicIpPrefixAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pip.ref.Append("tags"))
}

// Zones returns a reference to field zones of azurerm_public_ip_prefix.
func (pip publicIpPrefixAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pip.ref.Append("zones"))
}

func (pip publicIpPrefixAttributes) Timeouts() publicipprefix.TimeoutsAttributes {
	return terra.ReferenceAsSingle[publicipprefix.TimeoutsAttributes](pip.ref.Append("timeouts"))
}

type publicIpPrefixState struct {
	Id                string                        `json:"id"`
	IpPrefix          string                        `json:"ip_prefix"`
	IpVersion         string                        `json:"ip_version"`
	Location          string                        `json:"location"`
	Name              string                        `json:"name"`
	PrefixLength      float64                       `json:"prefix_length"`
	ResourceGroupName string                        `json:"resource_group_name"`
	Sku               string                        `json:"sku"`
	Tags              map[string]string             `json:"tags"`
	Zones             []string                      `json:"zones"`
	Timeouts          *publicipprefix.TimeoutsState `json:"timeouts"`
}
