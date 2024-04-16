// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_static_web_app

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_static_web_app.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermStaticWebAppState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aswa *Resource) Type() string {
	return "azurerm_static_web_app"
}

// LocalName returns the local name for [Resource].
func (aswa *Resource) LocalName() string {
	return aswa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aswa *Resource) Configuration() interface{} {
	return aswa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aswa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aswa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aswa *Resource) Dependencies() terra.Dependencies {
	return aswa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aswa *Resource) LifecycleManagement() *terra.Lifecycle {
	return aswa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aswa *Resource) Attributes() azurermStaticWebAppAttributes {
	return azurermStaticWebAppAttributes{ref: terra.ReferenceResource(aswa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aswa *Resource) ImportState(state io.Reader) error {
	aswa.state = &azurermStaticWebAppState{}
	if err := json.NewDecoder(state).Decode(aswa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aswa.Type(), aswa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aswa *Resource) State() (*azurermStaticWebAppState, bool) {
	return aswa.state, aswa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aswa *Resource) StateMust() *azurermStaticWebAppState {
	if aswa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aswa.Type(), aswa.LocalName()))
	}
	return aswa.state
}

// Args contains the configurations for azurerm_static_web_app.
type Args struct {
	// AppSettings: map of string, optional
	AppSettings terra.MapValue[terra.StringValue] `hcl:"app_settings,attr"`
	// ConfigurationFileChangesEnabled: bool, optional
	ConfigurationFileChangesEnabled terra.BoolValue `hcl:"configuration_file_changes_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PreviewEnvironmentsEnabled: bool, optional
	PreviewEnvironmentsEnabled terra.BoolValue `hcl:"preview_environments_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuSize: string, optional
	SkuSize terra.StringValue `hcl:"sku_size,attr"`
	// SkuTier: string, optional
	SkuTier terra.StringValue `hcl:"sku_tier,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// BasicAuth: optional
	BasicAuth *BasicAuth `hcl:"basic_auth,block"`
	// Identity: optional
	Identity *Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermStaticWebAppAttributes struct {
	ref terra.Reference
}

// ApiKey returns a reference to field api_key of azurerm_static_web_app.
func (aswa azurermStaticWebAppAttributes) ApiKey() terra.StringValue {
	return terra.ReferenceAsString(aswa.ref.Append("api_key"))
}

// AppSettings returns a reference to field app_settings of azurerm_static_web_app.
func (aswa azurermStaticWebAppAttributes) AppSettings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aswa.ref.Append("app_settings"))
}

// ConfigurationFileChangesEnabled returns a reference to field configuration_file_changes_enabled of azurerm_static_web_app.
func (aswa azurermStaticWebAppAttributes) ConfigurationFileChangesEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aswa.ref.Append("configuration_file_changes_enabled"))
}

// DefaultHostName returns a reference to field default_host_name of azurerm_static_web_app.
func (aswa azurermStaticWebAppAttributes) DefaultHostName() terra.StringValue {
	return terra.ReferenceAsString(aswa.ref.Append("default_host_name"))
}

// Id returns a reference to field id of azurerm_static_web_app.
func (aswa azurermStaticWebAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aswa.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_static_web_app.
func (aswa azurermStaticWebAppAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aswa.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_static_web_app.
func (aswa azurermStaticWebAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aswa.ref.Append("name"))
}

// PreviewEnvironmentsEnabled returns a reference to field preview_environments_enabled of azurerm_static_web_app.
func (aswa azurermStaticWebAppAttributes) PreviewEnvironmentsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aswa.ref.Append("preview_environments_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_static_web_app.
func (aswa azurermStaticWebAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aswa.ref.Append("resource_group_name"))
}

// SkuSize returns a reference to field sku_size of azurerm_static_web_app.
func (aswa azurermStaticWebAppAttributes) SkuSize() terra.StringValue {
	return terra.ReferenceAsString(aswa.ref.Append("sku_size"))
}

// SkuTier returns a reference to field sku_tier of azurerm_static_web_app.
func (aswa azurermStaticWebAppAttributes) SkuTier() terra.StringValue {
	return terra.ReferenceAsString(aswa.ref.Append("sku_tier"))
}

// Tags returns a reference to field tags of azurerm_static_web_app.
func (aswa azurermStaticWebAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aswa.ref.Append("tags"))
}

func (aswa azurermStaticWebAppAttributes) BasicAuth() terra.ListValue[BasicAuthAttributes] {
	return terra.ReferenceAsList[BasicAuthAttributes](aswa.ref.Append("basic_auth"))
}

func (aswa azurermStaticWebAppAttributes) Identity() terra.ListValue[IdentityAttributes] {
	return terra.ReferenceAsList[IdentityAttributes](aswa.ref.Append("identity"))
}

func (aswa azurermStaticWebAppAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aswa.ref.Append("timeouts"))
}

type azurermStaticWebAppState struct {
	ApiKey                          string            `json:"api_key"`
	AppSettings                     map[string]string `json:"app_settings"`
	ConfigurationFileChangesEnabled bool              `json:"configuration_file_changes_enabled"`
	DefaultHostName                 string            `json:"default_host_name"`
	Id                              string            `json:"id"`
	Location                        string            `json:"location"`
	Name                            string            `json:"name"`
	PreviewEnvironmentsEnabled      bool              `json:"preview_environments_enabled"`
	ResourceGroupName               string            `json:"resource_group_name"`
	SkuSize                         string            `json:"sku_size"`
	SkuTier                         string            `json:"sku_tier"`
	Tags                            map[string]string `json:"tags"`
	BasicAuth                       []BasicAuthState  `json:"basic_auth"`
	Identity                        []IdentityState   `json:"identity"`
	Timeouts                        *TimeoutsState    `json:"timeouts"`
}
