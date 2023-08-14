// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	staticsite "github.com/golingon/terraproviders/azurerm/3.69.0/staticsite"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStaticSite creates a new instance of [StaticSite].
func NewStaticSite(name string, args StaticSiteArgs) *StaticSite {
	return &StaticSite{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StaticSite)(nil)

// StaticSite represents the Terraform resource azurerm_static_site.
type StaticSite struct {
	Name      string
	Args      StaticSiteArgs
	state     *staticSiteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StaticSite].
func (ss *StaticSite) Type() string {
	return "azurerm_static_site"
}

// LocalName returns the local name for [StaticSite].
func (ss *StaticSite) LocalName() string {
	return ss.Name
}

// Configuration returns the configuration (args) for [StaticSite].
func (ss *StaticSite) Configuration() interface{} {
	return ss.Args
}

// DependOn is used for other resources to depend on [StaticSite].
func (ss *StaticSite) DependOn() terra.Reference {
	return terra.ReferenceResource(ss)
}

// Dependencies returns the list of resources [StaticSite] depends_on.
func (ss *StaticSite) Dependencies() terra.Dependencies {
	return ss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StaticSite].
func (ss *StaticSite) LifecycleManagement() *terra.Lifecycle {
	return ss.Lifecycle
}

// Attributes returns the attributes for [StaticSite].
func (ss *StaticSite) Attributes() staticSiteAttributes {
	return staticSiteAttributes{ref: terra.ReferenceResource(ss)}
}

// ImportState imports the given attribute values into [StaticSite]'s state.
func (ss *StaticSite) ImportState(av io.Reader) error {
	ss.state = &staticSiteState{}
	if err := json.NewDecoder(av).Decode(ss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ss.Type(), ss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StaticSite] has state.
func (ss *StaticSite) State() (*staticSiteState, bool) {
	return ss.state, ss.state != nil
}

// StateMust returns the state for [StaticSite]. Panics if the state is nil.
func (ss *StaticSite) StateMust() *staticSiteState {
	if ss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ss.Type(), ss.LocalName()))
	}
	return ss.state
}

// StaticSiteArgs contains the configurations for azurerm_static_site.
type StaticSiteArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuSize: string, optional
	SkuSize terra.StringValue `hcl:"sku_size,attr"`
	// SkuTier: string, optional
	SkuTier terra.StringValue `hcl:"sku_tier,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Identity: optional
	Identity *staticsite.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *staticsite.Timeouts `hcl:"timeouts,block"`
}
type staticSiteAttributes struct {
	ref terra.Reference
}

// ApiKey returns a reference to field api_key of azurerm_static_site.
func (ss staticSiteAttributes) ApiKey() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("api_key"))
}

// DefaultHostName returns a reference to field default_host_name of azurerm_static_site.
func (ss staticSiteAttributes) DefaultHostName() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("default_host_name"))
}

// Id returns a reference to field id of azurerm_static_site.
func (ss staticSiteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_static_site.
func (ss staticSiteAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_static_site.
func (ss staticSiteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_static_site.
func (ss staticSiteAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("resource_group_name"))
}

// SkuSize returns a reference to field sku_size of azurerm_static_site.
func (ss staticSiteAttributes) SkuSize() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("sku_size"))
}

// SkuTier returns a reference to field sku_tier of azurerm_static_site.
func (ss staticSiteAttributes) SkuTier() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("sku_tier"))
}

// Tags returns a reference to field tags of azurerm_static_site.
func (ss staticSiteAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ss.ref.Append("tags"))
}

func (ss staticSiteAttributes) Identity() terra.ListValue[staticsite.IdentityAttributes] {
	return terra.ReferenceAsList[staticsite.IdentityAttributes](ss.ref.Append("identity"))
}

func (ss staticSiteAttributes) Timeouts() staticsite.TimeoutsAttributes {
	return terra.ReferenceAsSingle[staticsite.TimeoutsAttributes](ss.ref.Append("timeouts"))
}

type staticSiteState struct {
	ApiKey            string                     `json:"api_key"`
	DefaultHostName   string                     `json:"default_host_name"`
	Id                string                     `json:"id"`
	Location          string                     `json:"location"`
	Name              string                     `json:"name"`
	ResourceGroupName string                     `json:"resource_group_name"`
	SkuSize           string                     `json:"sku_size"`
	SkuTier           string                     `json:"sku_tier"`
	Tags              map[string]string          `json:"tags"`
	Identity          []staticsite.IdentityState `json:"identity"`
	Timeouts          *staticsite.TimeoutsState  `json:"timeouts"`
}
