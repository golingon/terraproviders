// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appserviceslotcustomhostnamebinding "github.com/golingon/terraproviders/azurerm/3.63.0/appserviceslotcustomhostnamebinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppServiceSlotCustomHostnameBinding creates a new instance of [AppServiceSlotCustomHostnameBinding].
func NewAppServiceSlotCustomHostnameBinding(name string, args AppServiceSlotCustomHostnameBindingArgs) *AppServiceSlotCustomHostnameBinding {
	return &AppServiceSlotCustomHostnameBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServiceSlotCustomHostnameBinding)(nil)

// AppServiceSlotCustomHostnameBinding represents the Terraform resource azurerm_app_service_slot_custom_hostname_binding.
type AppServiceSlotCustomHostnameBinding struct {
	Name      string
	Args      AppServiceSlotCustomHostnameBindingArgs
	state     *appServiceSlotCustomHostnameBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppServiceSlotCustomHostnameBinding].
func (asschb *AppServiceSlotCustomHostnameBinding) Type() string {
	return "azurerm_app_service_slot_custom_hostname_binding"
}

// LocalName returns the local name for [AppServiceSlotCustomHostnameBinding].
func (asschb *AppServiceSlotCustomHostnameBinding) LocalName() string {
	return asschb.Name
}

// Configuration returns the configuration (args) for [AppServiceSlotCustomHostnameBinding].
func (asschb *AppServiceSlotCustomHostnameBinding) Configuration() interface{} {
	return asschb.Args
}

// DependOn is used for other resources to depend on [AppServiceSlotCustomHostnameBinding].
func (asschb *AppServiceSlotCustomHostnameBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(asschb)
}

// Dependencies returns the list of resources [AppServiceSlotCustomHostnameBinding] depends_on.
func (asschb *AppServiceSlotCustomHostnameBinding) Dependencies() terra.Dependencies {
	return asschb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppServiceSlotCustomHostnameBinding].
func (asschb *AppServiceSlotCustomHostnameBinding) LifecycleManagement() *terra.Lifecycle {
	return asschb.Lifecycle
}

// Attributes returns the attributes for [AppServiceSlotCustomHostnameBinding].
func (asschb *AppServiceSlotCustomHostnameBinding) Attributes() appServiceSlotCustomHostnameBindingAttributes {
	return appServiceSlotCustomHostnameBindingAttributes{ref: terra.ReferenceResource(asschb)}
}

// ImportState imports the given attribute values into [AppServiceSlotCustomHostnameBinding]'s state.
func (asschb *AppServiceSlotCustomHostnameBinding) ImportState(av io.Reader) error {
	asschb.state = &appServiceSlotCustomHostnameBindingState{}
	if err := json.NewDecoder(av).Decode(asschb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asschb.Type(), asschb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppServiceSlotCustomHostnameBinding] has state.
func (asschb *AppServiceSlotCustomHostnameBinding) State() (*appServiceSlotCustomHostnameBindingState, bool) {
	return asschb.state, asschb.state != nil
}

// StateMust returns the state for [AppServiceSlotCustomHostnameBinding]. Panics if the state is nil.
func (asschb *AppServiceSlotCustomHostnameBinding) StateMust() *appServiceSlotCustomHostnameBindingState {
	if asschb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asschb.Type(), asschb.LocalName()))
	}
	return asschb.state
}

// AppServiceSlotCustomHostnameBindingArgs contains the configurations for azurerm_app_service_slot_custom_hostname_binding.
type AppServiceSlotCustomHostnameBindingArgs struct {
	// AppServiceSlotId: string, required
	AppServiceSlotId terra.StringValue `hcl:"app_service_slot_id,attr" validate:"required"`
	// Hostname: string, required
	Hostname terra.StringValue `hcl:"hostname,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SslState: string, optional
	SslState terra.StringValue `hcl:"ssl_state,attr"`
	// Thumbprint: string, optional
	Thumbprint terra.StringValue `hcl:"thumbprint,attr"`
	// Timeouts: optional
	Timeouts *appserviceslotcustomhostnamebinding.Timeouts `hcl:"timeouts,block"`
}
type appServiceSlotCustomHostnameBindingAttributes struct {
	ref terra.Reference
}

// AppServiceSlotId returns a reference to field app_service_slot_id of azurerm_app_service_slot_custom_hostname_binding.
func (asschb appServiceSlotCustomHostnameBindingAttributes) AppServiceSlotId() terra.StringValue {
	return terra.ReferenceAsString(asschb.ref.Append("app_service_slot_id"))
}

// Hostname returns a reference to field hostname of azurerm_app_service_slot_custom_hostname_binding.
func (asschb appServiceSlotCustomHostnameBindingAttributes) Hostname() terra.StringValue {
	return terra.ReferenceAsString(asschb.ref.Append("hostname"))
}

// Id returns a reference to field id of azurerm_app_service_slot_custom_hostname_binding.
func (asschb appServiceSlotCustomHostnameBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asschb.ref.Append("id"))
}

// SslState returns a reference to field ssl_state of azurerm_app_service_slot_custom_hostname_binding.
func (asschb appServiceSlotCustomHostnameBindingAttributes) SslState() terra.StringValue {
	return terra.ReferenceAsString(asschb.ref.Append("ssl_state"))
}

// Thumbprint returns a reference to field thumbprint of azurerm_app_service_slot_custom_hostname_binding.
func (asschb appServiceSlotCustomHostnameBindingAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(asschb.ref.Append("thumbprint"))
}

// VirtualIp returns a reference to field virtual_ip of azurerm_app_service_slot_custom_hostname_binding.
func (asschb appServiceSlotCustomHostnameBindingAttributes) VirtualIp() terra.StringValue {
	return terra.ReferenceAsString(asschb.ref.Append("virtual_ip"))
}

func (asschb appServiceSlotCustomHostnameBindingAttributes) Timeouts() appserviceslotcustomhostnamebinding.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appserviceslotcustomhostnamebinding.TimeoutsAttributes](asschb.ref.Append("timeouts"))
}

type appServiceSlotCustomHostnameBindingState struct {
	AppServiceSlotId string                                             `json:"app_service_slot_id"`
	Hostname         string                                             `json:"hostname"`
	Id               string                                             `json:"id"`
	SslState         string                                             `json:"ssl_state"`
	Thumbprint       string                                             `json:"thumbprint"`
	VirtualIp        string                                             `json:"virtual_ip"`
	Timeouts         *appserviceslotcustomhostnamebinding.TimeoutsState `json:"timeouts"`
}