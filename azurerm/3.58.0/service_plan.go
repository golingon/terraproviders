// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	serviceplan "github.com/golingon/terraproviders/azurerm/3.58.0/serviceplan"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicePlan creates a new instance of [ServicePlan].
func NewServicePlan(name string, args ServicePlanArgs) *ServicePlan {
	return &ServicePlan{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicePlan)(nil)

// ServicePlan represents the Terraform resource azurerm_service_plan.
type ServicePlan struct {
	Name      string
	Args      ServicePlanArgs
	state     *servicePlanState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicePlan].
func (sp *ServicePlan) Type() string {
	return "azurerm_service_plan"
}

// LocalName returns the local name for [ServicePlan].
func (sp *ServicePlan) LocalName() string {
	return sp.Name
}

// Configuration returns the configuration (args) for [ServicePlan].
func (sp *ServicePlan) Configuration() interface{} {
	return sp.Args
}

// DependOn is used for other resources to depend on [ServicePlan].
func (sp *ServicePlan) DependOn() terra.Reference {
	return terra.ReferenceResource(sp)
}

// Dependencies returns the list of resources [ServicePlan] depends_on.
func (sp *ServicePlan) Dependencies() terra.Dependencies {
	return sp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicePlan].
func (sp *ServicePlan) LifecycleManagement() *terra.Lifecycle {
	return sp.Lifecycle
}

// Attributes returns the attributes for [ServicePlan].
func (sp *ServicePlan) Attributes() servicePlanAttributes {
	return servicePlanAttributes{ref: terra.ReferenceResource(sp)}
}

// ImportState imports the given attribute values into [ServicePlan]'s state.
func (sp *ServicePlan) ImportState(av io.Reader) error {
	sp.state = &servicePlanState{}
	if err := json.NewDecoder(av).Decode(sp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sp.Type(), sp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicePlan] has state.
func (sp *ServicePlan) State() (*servicePlanState, bool) {
	return sp.state, sp.state != nil
}

// StateMust returns the state for [ServicePlan]. Panics if the state is nil.
func (sp *ServicePlan) StateMust() *servicePlanState {
	if sp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sp.Type(), sp.LocalName()))
	}
	return sp.state
}

// ServicePlanArgs contains the configurations for azurerm_service_plan.
type ServicePlanArgs struct {
	// AppServiceEnvironmentId: string, optional
	AppServiceEnvironmentId terra.StringValue `hcl:"app_service_environment_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MaximumElasticWorkerCount: number, optional
	MaximumElasticWorkerCount terra.NumberValue `hcl:"maximum_elastic_worker_count,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OsType: string, required
	OsType terra.StringValue `hcl:"os_type,attr" validate:"required"`
	// PerSiteScalingEnabled: bool, optional
	PerSiteScalingEnabled terra.BoolValue `hcl:"per_site_scaling_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// WorkerCount: number, optional
	WorkerCount terra.NumberValue `hcl:"worker_count,attr"`
	// ZoneBalancingEnabled: bool, optional
	ZoneBalancingEnabled terra.BoolValue `hcl:"zone_balancing_enabled,attr"`
	// Timeouts: optional
	Timeouts *serviceplan.Timeouts `hcl:"timeouts,block"`
}
type servicePlanAttributes struct {
	ref terra.Reference
}

// AppServiceEnvironmentId returns a reference to field app_service_environment_id of azurerm_service_plan.
func (sp servicePlanAttributes) AppServiceEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("app_service_environment_id"))
}

// Id returns a reference to field id of azurerm_service_plan.
func (sp servicePlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_service_plan.
func (sp servicePlanAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_service_plan.
func (sp servicePlanAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("location"))
}

// MaximumElasticWorkerCount returns a reference to field maximum_elastic_worker_count of azurerm_service_plan.
func (sp servicePlanAttributes) MaximumElasticWorkerCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sp.ref.Append("maximum_elastic_worker_count"))
}

// Name returns a reference to field name of azurerm_service_plan.
func (sp servicePlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("name"))
}

// OsType returns a reference to field os_type of azurerm_service_plan.
func (sp servicePlanAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("os_type"))
}

// PerSiteScalingEnabled returns a reference to field per_site_scaling_enabled of azurerm_service_plan.
func (sp servicePlanAttributes) PerSiteScalingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("per_site_scaling_enabled"))
}

// Reserved returns a reference to field reserved of azurerm_service_plan.
func (sp servicePlanAttributes) Reserved() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("reserved"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_service_plan.
func (sp servicePlanAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_service_plan.
func (sp servicePlanAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_service_plan.
func (sp servicePlanAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sp.ref.Append("tags"))
}

// WorkerCount returns a reference to field worker_count of azurerm_service_plan.
func (sp servicePlanAttributes) WorkerCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sp.ref.Append("worker_count"))
}

// ZoneBalancingEnabled returns a reference to field zone_balancing_enabled of azurerm_service_plan.
func (sp servicePlanAttributes) ZoneBalancingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("zone_balancing_enabled"))
}

func (sp servicePlanAttributes) Timeouts() serviceplan.TimeoutsAttributes {
	return terra.ReferenceAsSingle[serviceplan.TimeoutsAttributes](sp.ref.Append("timeouts"))
}

type servicePlanState struct {
	AppServiceEnvironmentId   string                     `json:"app_service_environment_id"`
	Id                        string                     `json:"id"`
	Kind                      string                     `json:"kind"`
	Location                  string                     `json:"location"`
	MaximumElasticWorkerCount float64                    `json:"maximum_elastic_worker_count"`
	Name                      string                     `json:"name"`
	OsType                    string                     `json:"os_type"`
	PerSiteScalingEnabled     bool                       `json:"per_site_scaling_enabled"`
	Reserved                  bool                       `json:"reserved"`
	ResourceGroupName         string                     `json:"resource_group_name"`
	SkuName                   string                     `json:"sku_name"`
	Tags                      map[string]string          `json:"tags"`
	WorkerCount               float64                    `json:"worker_count"`
	ZoneBalancingEnabled      bool                       `json:"zone_balancing_enabled"`
	Timeouts                  *serviceplan.TimeoutsState `json:"timeouts"`
}
