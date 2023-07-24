// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	resourcegroupcostmanagementview "github.com/golingon/terraproviders/azurerm/3.66.0/resourcegroupcostmanagementview"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewResourceGroupCostManagementView creates a new instance of [ResourceGroupCostManagementView].
func NewResourceGroupCostManagementView(name string, args ResourceGroupCostManagementViewArgs) *ResourceGroupCostManagementView {
	return &ResourceGroupCostManagementView{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ResourceGroupCostManagementView)(nil)

// ResourceGroupCostManagementView represents the Terraform resource azurerm_resource_group_cost_management_view.
type ResourceGroupCostManagementView struct {
	Name      string
	Args      ResourceGroupCostManagementViewArgs
	state     *resourceGroupCostManagementViewState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ResourceGroupCostManagementView].
func (rgcmv *ResourceGroupCostManagementView) Type() string {
	return "azurerm_resource_group_cost_management_view"
}

// LocalName returns the local name for [ResourceGroupCostManagementView].
func (rgcmv *ResourceGroupCostManagementView) LocalName() string {
	return rgcmv.Name
}

// Configuration returns the configuration (args) for [ResourceGroupCostManagementView].
func (rgcmv *ResourceGroupCostManagementView) Configuration() interface{} {
	return rgcmv.Args
}

// DependOn is used for other resources to depend on [ResourceGroupCostManagementView].
func (rgcmv *ResourceGroupCostManagementView) DependOn() terra.Reference {
	return terra.ReferenceResource(rgcmv)
}

// Dependencies returns the list of resources [ResourceGroupCostManagementView] depends_on.
func (rgcmv *ResourceGroupCostManagementView) Dependencies() terra.Dependencies {
	return rgcmv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ResourceGroupCostManagementView].
func (rgcmv *ResourceGroupCostManagementView) LifecycleManagement() *terra.Lifecycle {
	return rgcmv.Lifecycle
}

// Attributes returns the attributes for [ResourceGroupCostManagementView].
func (rgcmv *ResourceGroupCostManagementView) Attributes() resourceGroupCostManagementViewAttributes {
	return resourceGroupCostManagementViewAttributes{ref: terra.ReferenceResource(rgcmv)}
}

// ImportState imports the given attribute values into [ResourceGroupCostManagementView]'s state.
func (rgcmv *ResourceGroupCostManagementView) ImportState(av io.Reader) error {
	rgcmv.state = &resourceGroupCostManagementViewState{}
	if err := json.NewDecoder(av).Decode(rgcmv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rgcmv.Type(), rgcmv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ResourceGroupCostManagementView] has state.
func (rgcmv *ResourceGroupCostManagementView) State() (*resourceGroupCostManagementViewState, bool) {
	return rgcmv.state, rgcmv.state != nil
}

// StateMust returns the state for [ResourceGroupCostManagementView]. Panics if the state is nil.
func (rgcmv *ResourceGroupCostManagementView) StateMust() *resourceGroupCostManagementViewState {
	if rgcmv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rgcmv.Type(), rgcmv.LocalName()))
	}
	return rgcmv.state
}

// ResourceGroupCostManagementViewArgs contains the configurations for azurerm_resource_group_cost_management_view.
type ResourceGroupCostManagementViewArgs struct {
	// Accumulated: bool, required
	Accumulated terra.BoolValue `hcl:"accumulated,attr" validate:"required"`
	// ChartType: string, required
	ChartType terra.StringValue `hcl:"chart_type,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ReportType: string, required
	ReportType terra.StringValue `hcl:"report_type,attr" validate:"required"`
	// ResourceGroupId: string, required
	ResourceGroupId terra.StringValue `hcl:"resource_group_id,attr" validate:"required"`
	// Timeframe: string, required
	Timeframe terra.StringValue `hcl:"timeframe,attr" validate:"required"`
	// Dataset: required
	Dataset *resourcegroupcostmanagementview.Dataset `hcl:"dataset,block" validate:"required"`
	// Kpi: min=0
	Kpi []resourcegroupcostmanagementview.Kpi `hcl:"kpi,block" validate:"min=0"`
	// Pivot: min=0
	Pivot []resourcegroupcostmanagementview.Pivot `hcl:"pivot,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *resourcegroupcostmanagementview.Timeouts `hcl:"timeouts,block"`
}
type resourceGroupCostManagementViewAttributes struct {
	ref terra.Reference
}

// Accumulated returns a reference to field accumulated of azurerm_resource_group_cost_management_view.
func (rgcmv resourceGroupCostManagementViewAttributes) Accumulated() terra.BoolValue {
	return terra.ReferenceAsBool(rgcmv.ref.Append("accumulated"))
}

// ChartType returns a reference to field chart_type of azurerm_resource_group_cost_management_view.
func (rgcmv resourceGroupCostManagementViewAttributes) ChartType() terra.StringValue {
	return terra.ReferenceAsString(rgcmv.ref.Append("chart_type"))
}

// DisplayName returns a reference to field display_name of azurerm_resource_group_cost_management_view.
func (rgcmv resourceGroupCostManagementViewAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(rgcmv.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_resource_group_cost_management_view.
func (rgcmv resourceGroupCostManagementViewAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rgcmv.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_resource_group_cost_management_view.
func (rgcmv resourceGroupCostManagementViewAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rgcmv.ref.Append("name"))
}

// ReportType returns a reference to field report_type of azurerm_resource_group_cost_management_view.
func (rgcmv resourceGroupCostManagementViewAttributes) ReportType() terra.StringValue {
	return terra.ReferenceAsString(rgcmv.ref.Append("report_type"))
}

// ResourceGroupId returns a reference to field resource_group_id of azurerm_resource_group_cost_management_view.
func (rgcmv resourceGroupCostManagementViewAttributes) ResourceGroupId() terra.StringValue {
	return terra.ReferenceAsString(rgcmv.ref.Append("resource_group_id"))
}

// Timeframe returns a reference to field timeframe of azurerm_resource_group_cost_management_view.
func (rgcmv resourceGroupCostManagementViewAttributes) Timeframe() terra.StringValue {
	return terra.ReferenceAsString(rgcmv.ref.Append("timeframe"))
}

func (rgcmv resourceGroupCostManagementViewAttributes) Dataset() terra.ListValue[resourcegroupcostmanagementview.DatasetAttributes] {
	return terra.ReferenceAsList[resourcegroupcostmanagementview.DatasetAttributes](rgcmv.ref.Append("dataset"))
}

func (rgcmv resourceGroupCostManagementViewAttributes) Kpi() terra.ListValue[resourcegroupcostmanagementview.KpiAttributes] {
	return terra.ReferenceAsList[resourcegroupcostmanagementview.KpiAttributes](rgcmv.ref.Append("kpi"))
}

func (rgcmv resourceGroupCostManagementViewAttributes) Pivot() terra.ListValue[resourcegroupcostmanagementview.PivotAttributes] {
	return terra.ReferenceAsList[resourcegroupcostmanagementview.PivotAttributes](rgcmv.ref.Append("pivot"))
}

func (rgcmv resourceGroupCostManagementViewAttributes) Timeouts() resourcegroupcostmanagementview.TimeoutsAttributes {
	return terra.ReferenceAsSingle[resourcegroupcostmanagementview.TimeoutsAttributes](rgcmv.ref.Append("timeouts"))
}

type resourceGroupCostManagementViewState struct {
	Accumulated     bool                                           `json:"accumulated"`
	ChartType       string                                         `json:"chart_type"`
	DisplayName     string                                         `json:"display_name"`
	Id              string                                         `json:"id"`
	Name            string                                         `json:"name"`
	ReportType      string                                         `json:"report_type"`
	ResourceGroupId string                                         `json:"resource_group_id"`
	Timeframe       string                                         `json:"timeframe"`
	Dataset         []resourcegroupcostmanagementview.DatasetState `json:"dataset"`
	Kpi             []resourcegroupcostmanagementview.KpiState     `json:"kpi"`
	Pivot           []resourcegroupcostmanagementview.PivotState   `json:"pivot"`
	Timeouts        *resourcegroupcostmanagementview.TimeoutsState `json:"timeouts"`
}
