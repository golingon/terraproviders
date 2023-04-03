// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataappserviceplan "github.com/golingon/terraproviders/azurerm/3.49.0/dataappserviceplan"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAppServicePlan creates a new instance of [DataAppServicePlan].
func NewDataAppServicePlan(name string, args DataAppServicePlanArgs) *DataAppServicePlan {
	return &DataAppServicePlan{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppServicePlan)(nil)

// DataAppServicePlan represents the Terraform data resource azurerm_app_service_plan.
type DataAppServicePlan struct {
	Name string
	Args DataAppServicePlanArgs
}

// DataSource returns the Terraform object type for [DataAppServicePlan].
func (asp *DataAppServicePlan) DataSource() string {
	return "azurerm_app_service_plan"
}

// LocalName returns the local name for [DataAppServicePlan].
func (asp *DataAppServicePlan) LocalName() string {
	return asp.Name
}

// Configuration returns the configuration (args) for [DataAppServicePlan].
func (asp *DataAppServicePlan) Configuration() interface{} {
	return asp.Args
}

// Attributes returns the attributes for [DataAppServicePlan].
func (asp *DataAppServicePlan) Attributes() dataAppServicePlanAttributes {
	return dataAppServicePlanAttributes{ref: terra.ReferenceDataResource(asp)}
}

// DataAppServicePlanArgs contains the configurations for azurerm_app_service_plan.
type DataAppServicePlanArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: min=0
	Sku []dataappserviceplan.Sku `hcl:"sku,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataappserviceplan.Timeouts `hcl:"timeouts,block"`
}
type dataAppServicePlanAttributes struct {
	ref terra.Reference
}

// AppServiceEnvironmentId returns a reference to field app_service_environment_id of azurerm_app_service_plan.
func (asp dataAppServicePlanAttributes) AppServiceEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("app_service_environment_id"))
}

// Id returns a reference to field id of azurerm_app_service_plan.
func (asp dataAppServicePlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("id"))
}

// IsXenon returns a reference to field is_xenon of azurerm_app_service_plan.
func (asp dataAppServicePlanAttributes) IsXenon() terra.BoolValue {
	return terra.ReferenceAsBool(asp.ref.Append("is_xenon"))
}

// Kind returns a reference to field kind of azurerm_app_service_plan.
func (asp dataAppServicePlanAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_app_service_plan.
func (asp dataAppServicePlanAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("location"))
}

// MaximumElasticWorkerCount returns a reference to field maximum_elastic_worker_count of azurerm_app_service_plan.
func (asp dataAppServicePlanAttributes) MaximumElasticWorkerCount() terra.NumberValue {
	return terra.ReferenceAsNumber(asp.ref.Append("maximum_elastic_worker_count"))
}

// MaximumNumberOfWorkers returns a reference to field maximum_number_of_workers of azurerm_app_service_plan.
func (asp dataAppServicePlanAttributes) MaximumNumberOfWorkers() terra.NumberValue {
	return terra.ReferenceAsNumber(asp.ref.Append("maximum_number_of_workers"))
}

// Name returns a reference to field name of azurerm_app_service_plan.
func (asp dataAppServicePlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("name"))
}

// PerSiteScaling returns a reference to field per_site_scaling of azurerm_app_service_plan.
func (asp dataAppServicePlanAttributes) PerSiteScaling() terra.BoolValue {
	return terra.ReferenceAsBool(asp.ref.Append("per_site_scaling"))
}

// Reserved returns a reference to field reserved of azurerm_app_service_plan.
func (asp dataAppServicePlanAttributes) Reserved() terra.BoolValue {
	return terra.ReferenceAsBool(asp.ref.Append("reserved"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_app_service_plan.
func (asp dataAppServicePlanAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_app_service_plan.
func (asp dataAppServicePlanAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asp.ref.Append("tags"))
}

// ZoneRedundant returns a reference to field zone_redundant of azurerm_app_service_plan.
func (asp dataAppServicePlanAttributes) ZoneRedundant() terra.BoolValue {
	return terra.ReferenceAsBool(asp.ref.Append("zone_redundant"))
}

func (asp dataAppServicePlanAttributes) Sku() terra.ListValue[dataappserviceplan.SkuAttributes] {
	return terra.ReferenceAsList[dataappserviceplan.SkuAttributes](asp.ref.Append("sku"))
}

func (asp dataAppServicePlanAttributes) Timeouts() dataappserviceplan.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataappserviceplan.TimeoutsAttributes](asp.ref.Append("timeouts"))
}
