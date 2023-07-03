// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataportaldashboard "github.com/golingon/terraproviders/azurerm/3.63.0/dataportaldashboard"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPortalDashboard creates a new instance of [DataPortalDashboard].
func NewDataPortalDashboard(name string, args DataPortalDashboardArgs) *DataPortalDashboard {
	return &DataPortalDashboard{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPortalDashboard)(nil)

// DataPortalDashboard represents the Terraform data resource azurerm_portal_dashboard.
type DataPortalDashboard struct {
	Name string
	Args DataPortalDashboardArgs
}

// DataSource returns the Terraform object type for [DataPortalDashboard].
func (pd *DataPortalDashboard) DataSource() string {
	return "azurerm_portal_dashboard"
}

// LocalName returns the local name for [DataPortalDashboard].
func (pd *DataPortalDashboard) LocalName() string {
	return pd.Name
}

// Configuration returns the configuration (args) for [DataPortalDashboard].
func (pd *DataPortalDashboard) Configuration() interface{} {
	return pd.Args
}

// Attributes returns the attributes for [DataPortalDashboard].
func (pd *DataPortalDashboard) Attributes() dataPortalDashboardAttributes {
	return dataPortalDashboardAttributes{ref: terra.ReferenceDataResource(pd)}
}

// DataPortalDashboardArgs contains the configurations for azurerm_portal_dashboard.
type DataPortalDashboardArgs struct {
	// DashboardProperties: string, optional
	DashboardProperties terra.StringValue `hcl:"dashboard_properties,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataportaldashboard.Timeouts `hcl:"timeouts,block"`
}
type dataPortalDashboardAttributes struct {
	ref terra.Reference
}

// DashboardProperties returns a reference to field dashboard_properties of azurerm_portal_dashboard.
func (pd dataPortalDashboardAttributes) DashboardProperties() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("dashboard_properties"))
}

// DisplayName returns a reference to field display_name of azurerm_portal_dashboard.
func (pd dataPortalDashboardAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_portal_dashboard.
func (pd dataPortalDashboardAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_portal_dashboard.
func (pd dataPortalDashboardAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_portal_dashboard.
func (pd dataPortalDashboardAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_portal_dashboard.
func (pd dataPortalDashboardAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_portal_dashboard.
func (pd dataPortalDashboardAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pd.ref.Append("tags"))
}

func (pd dataPortalDashboardAttributes) Timeouts() dataportaldashboard.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataportaldashboard.TimeoutsAttributes](pd.ref.Append("timeouts"))
}
