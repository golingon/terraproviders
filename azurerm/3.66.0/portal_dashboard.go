// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	portaldashboard "github.com/golingon/terraproviders/azurerm/3.66.0/portaldashboard"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPortalDashboard creates a new instance of [PortalDashboard].
func NewPortalDashboard(name string, args PortalDashboardArgs) *PortalDashboard {
	return &PortalDashboard{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PortalDashboard)(nil)

// PortalDashboard represents the Terraform resource azurerm_portal_dashboard.
type PortalDashboard struct {
	Name      string
	Args      PortalDashboardArgs
	state     *portalDashboardState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PortalDashboard].
func (pd *PortalDashboard) Type() string {
	return "azurerm_portal_dashboard"
}

// LocalName returns the local name for [PortalDashboard].
func (pd *PortalDashboard) LocalName() string {
	return pd.Name
}

// Configuration returns the configuration (args) for [PortalDashboard].
func (pd *PortalDashboard) Configuration() interface{} {
	return pd.Args
}

// DependOn is used for other resources to depend on [PortalDashboard].
func (pd *PortalDashboard) DependOn() terra.Reference {
	return terra.ReferenceResource(pd)
}

// Dependencies returns the list of resources [PortalDashboard] depends_on.
func (pd *PortalDashboard) Dependencies() terra.Dependencies {
	return pd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PortalDashboard].
func (pd *PortalDashboard) LifecycleManagement() *terra.Lifecycle {
	return pd.Lifecycle
}

// Attributes returns the attributes for [PortalDashboard].
func (pd *PortalDashboard) Attributes() portalDashboardAttributes {
	return portalDashboardAttributes{ref: terra.ReferenceResource(pd)}
}

// ImportState imports the given attribute values into [PortalDashboard]'s state.
func (pd *PortalDashboard) ImportState(av io.Reader) error {
	pd.state = &portalDashboardState{}
	if err := json.NewDecoder(av).Decode(pd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pd.Type(), pd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PortalDashboard] has state.
func (pd *PortalDashboard) State() (*portalDashboardState, bool) {
	return pd.state, pd.state != nil
}

// StateMust returns the state for [PortalDashboard]. Panics if the state is nil.
func (pd *PortalDashboard) StateMust() *portalDashboardState {
	if pd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pd.Type(), pd.LocalName()))
	}
	return pd.state
}

// PortalDashboardArgs contains the configurations for azurerm_portal_dashboard.
type PortalDashboardArgs struct {
	// DashboardProperties: string, required
	DashboardProperties terra.StringValue `hcl:"dashboard_properties,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *portaldashboard.Timeouts `hcl:"timeouts,block"`
}
type portalDashboardAttributes struct {
	ref terra.Reference
}

// DashboardProperties returns a reference to field dashboard_properties of azurerm_portal_dashboard.
func (pd portalDashboardAttributes) DashboardProperties() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("dashboard_properties"))
}

// Id returns a reference to field id of azurerm_portal_dashboard.
func (pd portalDashboardAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_portal_dashboard.
func (pd portalDashboardAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_portal_dashboard.
func (pd portalDashboardAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_portal_dashboard.
func (pd portalDashboardAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_portal_dashboard.
func (pd portalDashboardAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pd.ref.Append("tags"))
}

func (pd portalDashboardAttributes) Timeouts() portaldashboard.TimeoutsAttributes {
	return terra.ReferenceAsSingle[portaldashboard.TimeoutsAttributes](pd.ref.Append("timeouts"))
}

type portalDashboardState struct {
	DashboardProperties string                         `json:"dashboard_properties"`
	Id                  string                         `json:"id"`
	Location            string                         `json:"location"`
	Name                string                         `json:"name"`
	ResourceGroupName   string                         `json:"resource_group_name"`
	Tags                map[string]string              `json:"tags"`
	Timeouts            *portaldashboard.TimeoutsState `json:"timeouts"`
}
