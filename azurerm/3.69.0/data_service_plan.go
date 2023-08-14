// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataserviceplan "github.com/golingon/terraproviders/azurerm/3.69.0/dataserviceplan"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataServicePlan creates a new instance of [DataServicePlan].
func NewDataServicePlan(name string, args DataServicePlanArgs) *DataServicePlan {
	return &DataServicePlan{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServicePlan)(nil)

// DataServicePlan represents the Terraform data resource azurerm_service_plan.
type DataServicePlan struct {
	Name string
	Args DataServicePlanArgs
}

// DataSource returns the Terraform object type for [DataServicePlan].
func (sp *DataServicePlan) DataSource() string {
	return "azurerm_service_plan"
}

// LocalName returns the local name for [DataServicePlan].
func (sp *DataServicePlan) LocalName() string {
	return sp.Name
}

// Configuration returns the configuration (args) for [DataServicePlan].
func (sp *DataServicePlan) Configuration() interface{} {
	return sp.Args
}

// Attributes returns the attributes for [DataServicePlan].
func (sp *DataServicePlan) Attributes() dataServicePlanAttributes {
	return dataServicePlanAttributes{ref: terra.ReferenceDataResource(sp)}
}

// DataServicePlanArgs contains the configurations for azurerm_service_plan.
type DataServicePlanArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataserviceplan.Timeouts `hcl:"timeouts,block"`
}
type dataServicePlanAttributes struct {
	ref terra.Reference
}

// AppServiceEnvironmentId returns a reference to field app_service_environment_id of azurerm_service_plan.
func (sp dataServicePlanAttributes) AppServiceEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("app_service_environment_id"))
}

// Id returns a reference to field id of azurerm_service_plan.
func (sp dataServicePlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_service_plan.
func (sp dataServicePlanAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_service_plan.
func (sp dataServicePlanAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("location"))
}

// MaximumElasticWorkerCount returns a reference to field maximum_elastic_worker_count of azurerm_service_plan.
func (sp dataServicePlanAttributes) MaximumElasticWorkerCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sp.ref.Append("maximum_elastic_worker_count"))
}

// Name returns a reference to field name of azurerm_service_plan.
func (sp dataServicePlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("name"))
}

// OsType returns a reference to field os_type of azurerm_service_plan.
func (sp dataServicePlanAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("os_type"))
}

// PerSiteScalingEnabled returns a reference to field per_site_scaling_enabled of azurerm_service_plan.
func (sp dataServicePlanAttributes) PerSiteScalingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("per_site_scaling_enabled"))
}

// Reserved returns a reference to field reserved of azurerm_service_plan.
func (sp dataServicePlanAttributes) Reserved() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("reserved"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_service_plan.
func (sp dataServicePlanAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_service_plan.
func (sp dataServicePlanAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_service_plan.
func (sp dataServicePlanAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sp.ref.Append("tags"))
}

// WorkerCount returns a reference to field worker_count of azurerm_service_plan.
func (sp dataServicePlanAttributes) WorkerCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sp.ref.Append("worker_count"))
}

// ZoneBalancingEnabled returns a reference to field zone_balancing_enabled of azurerm_service_plan.
func (sp dataServicePlanAttributes) ZoneBalancingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("zone_balancing_enabled"))
}

func (sp dataServicePlanAttributes) Timeouts() dataserviceplan.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataserviceplan.TimeoutsAttributes](sp.ref.Append("timeouts"))
}
