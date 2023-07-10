// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appservicecustomhostnamebinding "github.com/golingon/terraproviders/azurerm/3.64.0/appservicecustomhostnamebinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppServiceCustomHostnameBinding creates a new instance of [AppServiceCustomHostnameBinding].
func NewAppServiceCustomHostnameBinding(name string, args AppServiceCustomHostnameBindingArgs) *AppServiceCustomHostnameBinding {
	return &AppServiceCustomHostnameBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServiceCustomHostnameBinding)(nil)

// AppServiceCustomHostnameBinding represents the Terraform resource azurerm_app_service_custom_hostname_binding.
type AppServiceCustomHostnameBinding struct {
	Name      string
	Args      AppServiceCustomHostnameBindingArgs
	state     *appServiceCustomHostnameBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppServiceCustomHostnameBinding].
func (aschb *AppServiceCustomHostnameBinding) Type() string {
	return "azurerm_app_service_custom_hostname_binding"
}

// LocalName returns the local name for [AppServiceCustomHostnameBinding].
func (aschb *AppServiceCustomHostnameBinding) LocalName() string {
	return aschb.Name
}

// Configuration returns the configuration (args) for [AppServiceCustomHostnameBinding].
func (aschb *AppServiceCustomHostnameBinding) Configuration() interface{} {
	return aschb.Args
}

// DependOn is used for other resources to depend on [AppServiceCustomHostnameBinding].
func (aschb *AppServiceCustomHostnameBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(aschb)
}

// Dependencies returns the list of resources [AppServiceCustomHostnameBinding] depends_on.
func (aschb *AppServiceCustomHostnameBinding) Dependencies() terra.Dependencies {
	return aschb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppServiceCustomHostnameBinding].
func (aschb *AppServiceCustomHostnameBinding) LifecycleManagement() *terra.Lifecycle {
	return aschb.Lifecycle
}

// Attributes returns the attributes for [AppServiceCustomHostnameBinding].
func (aschb *AppServiceCustomHostnameBinding) Attributes() appServiceCustomHostnameBindingAttributes {
	return appServiceCustomHostnameBindingAttributes{ref: terra.ReferenceResource(aschb)}
}

// ImportState imports the given attribute values into [AppServiceCustomHostnameBinding]'s state.
func (aschb *AppServiceCustomHostnameBinding) ImportState(av io.Reader) error {
	aschb.state = &appServiceCustomHostnameBindingState{}
	if err := json.NewDecoder(av).Decode(aschb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aschb.Type(), aschb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppServiceCustomHostnameBinding] has state.
func (aschb *AppServiceCustomHostnameBinding) State() (*appServiceCustomHostnameBindingState, bool) {
	return aschb.state, aschb.state != nil
}

// StateMust returns the state for [AppServiceCustomHostnameBinding]. Panics if the state is nil.
func (aschb *AppServiceCustomHostnameBinding) StateMust() *appServiceCustomHostnameBindingState {
	if aschb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aschb.Type(), aschb.LocalName()))
	}
	return aschb.state
}

// AppServiceCustomHostnameBindingArgs contains the configurations for azurerm_app_service_custom_hostname_binding.
type AppServiceCustomHostnameBindingArgs struct {
	// AppServiceName: string, required
	AppServiceName terra.StringValue `hcl:"app_service_name,attr" validate:"required"`
	// Hostname: string, required
	Hostname terra.StringValue `hcl:"hostname,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SslState: string, optional
	SslState terra.StringValue `hcl:"ssl_state,attr"`
	// Thumbprint: string, optional
	Thumbprint terra.StringValue `hcl:"thumbprint,attr"`
	// Timeouts: optional
	Timeouts *appservicecustomhostnamebinding.Timeouts `hcl:"timeouts,block"`
}
type appServiceCustomHostnameBindingAttributes struct {
	ref terra.Reference
}

// AppServiceName returns a reference to field app_service_name of azurerm_app_service_custom_hostname_binding.
func (aschb appServiceCustomHostnameBindingAttributes) AppServiceName() terra.StringValue {
	return terra.ReferenceAsString(aschb.ref.Append("app_service_name"))
}

// Hostname returns a reference to field hostname of azurerm_app_service_custom_hostname_binding.
func (aschb appServiceCustomHostnameBindingAttributes) Hostname() terra.StringValue {
	return terra.ReferenceAsString(aschb.ref.Append("hostname"))
}

// Id returns a reference to field id of azurerm_app_service_custom_hostname_binding.
func (aschb appServiceCustomHostnameBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aschb.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_app_service_custom_hostname_binding.
func (aschb appServiceCustomHostnameBindingAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aschb.ref.Append("resource_group_name"))
}

// SslState returns a reference to field ssl_state of azurerm_app_service_custom_hostname_binding.
func (aschb appServiceCustomHostnameBindingAttributes) SslState() terra.StringValue {
	return terra.ReferenceAsString(aschb.ref.Append("ssl_state"))
}

// Thumbprint returns a reference to field thumbprint of azurerm_app_service_custom_hostname_binding.
func (aschb appServiceCustomHostnameBindingAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(aschb.ref.Append("thumbprint"))
}

// VirtualIp returns a reference to field virtual_ip of azurerm_app_service_custom_hostname_binding.
func (aschb appServiceCustomHostnameBindingAttributes) VirtualIp() terra.StringValue {
	return terra.ReferenceAsString(aschb.ref.Append("virtual_ip"))
}

func (aschb appServiceCustomHostnameBindingAttributes) Timeouts() appservicecustomhostnamebinding.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appservicecustomhostnamebinding.TimeoutsAttributes](aschb.ref.Append("timeouts"))
}

type appServiceCustomHostnameBindingState struct {
	AppServiceName    string                                         `json:"app_service_name"`
	Hostname          string                                         `json:"hostname"`
	Id                string                                         `json:"id"`
	ResourceGroupName string                                         `json:"resource_group_name"`
	SslState          string                                         `json:"ssl_state"`
	Thumbprint        string                                         `json:"thumbprint"`
	VirtualIp         string                                         `json:"virtual_ip"`
	Timeouts          *appservicecustomhostnamebinding.TimeoutsState `json:"timeouts"`
}
