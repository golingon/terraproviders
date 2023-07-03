// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	portaltenantconfiguration "github.com/golingon/terraproviders/azurerm/3.63.0/portaltenantconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPortalTenantConfiguration creates a new instance of [PortalTenantConfiguration].
func NewPortalTenantConfiguration(name string, args PortalTenantConfigurationArgs) *PortalTenantConfiguration {
	return &PortalTenantConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PortalTenantConfiguration)(nil)

// PortalTenantConfiguration represents the Terraform resource azurerm_portal_tenant_configuration.
type PortalTenantConfiguration struct {
	Name      string
	Args      PortalTenantConfigurationArgs
	state     *portalTenantConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PortalTenantConfiguration].
func (ptc *PortalTenantConfiguration) Type() string {
	return "azurerm_portal_tenant_configuration"
}

// LocalName returns the local name for [PortalTenantConfiguration].
func (ptc *PortalTenantConfiguration) LocalName() string {
	return ptc.Name
}

// Configuration returns the configuration (args) for [PortalTenantConfiguration].
func (ptc *PortalTenantConfiguration) Configuration() interface{} {
	return ptc.Args
}

// DependOn is used for other resources to depend on [PortalTenantConfiguration].
func (ptc *PortalTenantConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(ptc)
}

// Dependencies returns the list of resources [PortalTenantConfiguration] depends_on.
func (ptc *PortalTenantConfiguration) Dependencies() terra.Dependencies {
	return ptc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PortalTenantConfiguration].
func (ptc *PortalTenantConfiguration) LifecycleManagement() *terra.Lifecycle {
	return ptc.Lifecycle
}

// Attributes returns the attributes for [PortalTenantConfiguration].
func (ptc *PortalTenantConfiguration) Attributes() portalTenantConfigurationAttributes {
	return portalTenantConfigurationAttributes{ref: terra.ReferenceResource(ptc)}
}

// ImportState imports the given attribute values into [PortalTenantConfiguration]'s state.
func (ptc *PortalTenantConfiguration) ImportState(av io.Reader) error {
	ptc.state = &portalTenantConfigurationState{}
	if err := json.NewDecoder(av).Decode(ptc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ptc.Type(), ptc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PortalTenantConfiguration] has state.
func (ptc *PortalTenantConfiguration) State() (*portalTenantConfigurationState, bool) {
	return ptc.state, ptc.state != nil
}

// StateMust returns the state for [PortalTenantConfiguration]. Panics if the state is nil.
func (ptc *PortalTenantConfiguration) StateMust() *portalTenantConfigurationState {
	if ptc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ptc.Type(), ptc.LocalName()))
	}
	return ptc.state
}

// PortalTenantConfigurationArgs contains the configurations for azurerm_portal_tenant_configuration.
type PortalTenantConfigurationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PrivateMarkdownStorageEnforced: bool, required
	PrivateMarkdownStorageEnforced terra.BoolValue `hcl:"private_markdown_storage_enforced,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *portaltenantconfiguration.Timeouts `hcl:"timeouts,block"`
}
type portalTenantConfigurationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_portal_tenant_configuration.
func (ptc portalTenantConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ptc.ref.Append("id"))
}

// PrivateMarkdownStorageEnforced returns a reference to field private_markdown_storage_enforced of azurerm_portal_tenant_configuration.
func (ptc portalTenantConfigurationAttributes) PrivateMarkdownStorageEnforced() terra.BoolValue {
	return terra.ReferenceAsBool(ptc.ref.Append("private_markdown_storage_enforced"))
}

func (ptc portalTenantConfigurationAttributes) Timeouts() portaltenantconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[portaltenantconfiguration.TimeoutsAttributes](ptc.ref.Append("timeouts"))
}

type portalTenantConfigurationState struct {
	Id                             string                                   `json:"id"`
	PrivateMarkdownStorageEnforced bool                                     `json:"private_markdown_storage_enforced"`
	Timeouts                       *portaltenantconfiguration.TimeoutsState `json:"timeouts"`
}
