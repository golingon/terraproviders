// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudapp "github.com/golingon/terraproviders/azurerm/3.69.0/springcloudapp"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpringCloudApp creates a new instance of [SpringCloudApp].
func NewSpringCloudApp(name string, args SpringCloudAppArgs) *SpringCloudApp {
	return &SpringCloudApp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudApp)(nil)

// SpringCloudApp represents the Terraform resource azurerm_spring_cloud_app.
type SpringCloudApp struct {
	Name      string
	Args      SpringCloudAppArgs
	state     *springCloudAppState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudApp].
func (sca *SpringCloudApp) Type() string {
	return "azurerm_spring_cloud_app"
}

// LocalName returns the local name for [SpringCloudApp].
func (sca *SpringCloudApp) LocalName() string {
	return sca.Name
}

// Configuration returns the configuration (args) for [SpringCloudApp].
func (sca *SpringCloudApp) Configuration() interface{} {
	return sca.Args
}

// DependOn is used for other resources to depend on [SpringCloudApp].
func (sca *SpringCloudApp) DependOn() terra.Reference {
	return terra.ReferenceResource(sca)
}

// Dependencies returns the list of resources [SpringCloudApp] depends_on.
func (sca *SpringCloudApp) Dependencies() terra.Dependencies {
	return sca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudApp].
func (sca *SpringCloudApp) LifecycleManagement() *terra.Lifecycle {
	return sca.Lifecycle
}

// Attributes returns the attributes for [SpringCloudApp].
func (sca *SpringCloudApp) Attributes() springCloudAppAttributes {
	return springCloudAppAttributes{ref: terra.ReferenceResource(sca)}
}

// ImportState imports the given attribute values into [SpringCloudApp]'s state.
func (sca *SpringCloudApp) ImportState(av io.Reader) error {
	sca.state = &springCloudAppState{}
	if err := json.NewDecoder(av).Decode(sca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sca.Type(), sca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudApp] has state.
func (sca *SpringCloudApp) State() (*springCloudAppState, bool) {
	return sca.state, sca.state != nil
}

// StateMust returns the state for [SpringCloudApp]. Panics if the state is nil.
func (sca *SpringCloudApp) StateMust() *springCloudAppState {
	if sca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sca.Type(), sca.LocalName()))
	}
	return sca.state
}

// SpringCloudAppArgs contains the configurations for azurerm_spring_cloud_app.
type SpringCloudAppArgs struct {
	// AddonJson: string, optional
	AddonJson terra.StringValue `hcl:"addon_json,attr"`
	// HttpsOnly: bool, optional
	HttpsOnly terra.BoolValue `hcl:"https_only,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsPublic: bool, optional
	IsPublic terra.BoolValue `hcl:"is_public,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicEndpointEnabled: bool, optional
	PublicEndpointEnabled terra.BoolValue `hcl:"public_endpoint_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// TlsEnabled: bool, optional
	TlsEnabled terra.BoolValue `hcl:"tls_enabled,attr"`
	// CustomPersistentDisk: min=0
	CustomPersistentDisk []springcloudapp.CustomPersistentDisk `hcl:"custom_persistent_disk,block" validate:"min=0"`
	// Identity: optional
	Identity *springcloudapp.Identity `hcl:"identity,block"`
	// IngressSettings: optional
	IngressSettings *springcloudapp.IngressSettings `hcl:"ingress_settings,block"`
	// PersistentDisk: optional
	PersistentDisk *springcloudapp.PersistentDisk `hcl:"persistent_disk,block"`
	// Timeouts: optional
	Timeouts *springcloudapp.Timeouts `hcl:"timeouts,block"`
}
type springCloudAppAttributes struct {
	ref terra.Reference
}

// AddonJson returns a reference to field addon_json of azurerm_spring_cloud_app.
func (sca springCloudAppAttributes) AddonJson() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("addon_json"))
}

// Fqdn returns a reference to field fqdn of azurerm_spring_cloud_app.
func (sca springCloudAppAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("fqdn"))
}

// HttpsOnly returns a reference to field https_only of azurerm_spring_cloud_app.
func (sca springCloudAppAttributes) HttpsOnly() terra.BoolValue {
	return terra.ReferenceAsBool(sca.ref.Append("https_only"))
}

// Id returns a reference to field id of azurerm_spring_cloud_app.
func (sca springCloudAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("id"))
}

// IsPublic returns a reference to field is_public of azurerm_spring_cloud_app.
func (sca springCloudAppAttributes) IsPublic() terra.BoolValue {
	return terra.ReferenceAsBool(sca.ref.Append("is_public"))
}

// Name returns a reference to field name of azurerm_spring_cloud_app.
func (sca springCloudAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("name"))
}

// PublicEndpointEnabled returns a reference to field public_endpoint_enabled of azurerm_spring_cloud_app.
func (sca springCloudAppAttributes) PublicEndpointEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sca.ref.Append("public_endpoint_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_spring_cloud_app.
func (sca springCloudAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("resource_group_name"))
}

// ServiceName returns a reference to field service_name of azurerm_spring_cloud_app.
func (sca springCloudAppAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("service_name"))
}

// TlsEnabled returns a reference to field tls_enabled of azurerm_spring_cloud_app.
func (sca springCloudAppAttributes) TlsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sca.ref.Append("tls_enabled"))
}

// Url returns a reference to field url of azurerm_spring_cloud_app.
func (sca springCloudAppAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(sca.ref.Append("url"))
}

func (sca springCloudAppAttributes) CustomPersistentDisk() terra.ListValue[springcloudapp.CustomPersistentDiskAttributes] {
	return terra.ReferenceAsList[springcloudapp.CustomPersistentDiskAttributes](sca.ref.Append("custom_persistent_disk"))
}

func (sca springCloudAppAttributes) Identity() terra.ListValue[springcloudapp.IdentityAttributes] {
	return terra.ReferenceAsList[springcloudapp.IdentityAttributes](sca.ref.Append("identity"))
}

func (sca springCloudAppAttributes) IngressSettings() terra.ListValue[springcloudapp.IngressSettingsAttributes] {
	return terra.ReferenceAsList[springcloudapp.IngressSettingsAttributes](sca.ref.Append("ingress_settings"))
}

func (sca springCloudAppAttributes) PersistentDisk() terra.ListValue[springcloudapp.PersistentDiskAttributes] {
	return terra.ReferenceAsList[springcloudapp.PersistentDiskAttributes](sca.ref.Append("persistent_disk"))
}

func (sca springCloudAppAttributes) Timeouts() springcloudapp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudapp.TimeoutsAttributes](sca.ref.Append("timeouts"))
}

type springCloudAppState struct {
	AddonJson             string                                     `json:"addon_json"`
	Fqdn                  string                                     `json:"fqdn"`
	HttpsOnly             bool                                       `json:"https_only"`
	Id                    string                                     `json:"id"`
	IsPublic              bool                                       `json:"is_public"`
	Name                  string                                     `json:"name"`
	PublicEndpointEnabled bool                                       `json:"public_endpoint_enabled"`
	ResourceGroupName     string                                     `json:"resource_group_name"`
	ServiceName           string                                     `json:"service_name"`
	TlsEnabled            bool                                       `json:"tls_enabled"`
	Url                   string                                     `json:"url"`
	CustomPersistentDisk  []springcloudapp.CustomPersistentDiskState `json:"custom_persistent_disk"`
	Identity              []springcloudapp.IdentityState             `json:"identity"`
	IngressSettings       []springcloudapp.IngressSettingsState      `json:"ingress_settings"`
	PersistentDisk        []springcloudapp.PersistentDiskState       `json:"persistent_disk"`
	Timeouts              *springcloudapp.TimeoutsState              `json:"timeouts"`
}
