// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appserviceplan "github.com/golingon/terraproviders/azurerm/3.68.0/appserviceplan"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppServicePlan creates a new instance of [AppServicePlan].
func NewAppServicePlan(name string, args AppServicePlanArgs) *AppServicePlan {
	return &AppServicePlan{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServicePlan)(nil)

// AppServicePlan represents the Terraform resource azurerm_app_service_plan.
type AppServicePlan struct {
	Name      string
	Args      AppServicePlanArgs
	state     *appServicePlanState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppServicePlan].
func (asp *AppServicePlan) Type() string {
	return "azurerm_app_service_plan"
}

// LocalName returns the local name for [AppServicePlan].
func (asp *AppServicePlan) LocalName() string {
	return asp.Name
}

// Configuration returns the configuration (args) for [AppServicePlan].
func (asp *AppServicePlan) Configuration() interface{} {
	return asp.Args
}

// DependOn is used for other resources to depend on [AppServicePlan].
func (asp *AppServicePlan) DependOn() terra.Reference {
	return terra.ReferenceResource(asp)
}

// Dependencies returns the list of resources [AppServicePlan] depends_on.
func (asp *AppServicePlan) Dependencies() terra.Dependencies {
	return asp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppServicePlan].
func (asp *AppServicePlan) LifecycleManagement() *terra.Lifecycle {
	return asp.Lifecycle
}

// Attributes returns the attributes for [AppServicePlan].
func (asp *AppServicePlan) Attributes() appServicePlanAttributes {
	return appServicePlanAttributes{ref: terra.ReferenceResource(asp)}
}

// ImportState imports the given attribute values into [AppServicePlan]'s state.
func (asp *AppServicePlan) ImportState(av io.Reader) error {
	asp.state = &appServicePlanState{}
	if err := json.NewDecoder(av).Decode(asp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asp.Type(), asp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppServicePlan] has state.
func (asp *AppServicePlan) State() (*appServicePlanState, bool) {
	return asp.state, asp.state != nil
}

// StateMust returns the state for [AppServicePlan]. Panics if the state is nil.
func (asp *AppServicePlan) StateMust() *appServicePlanState {
	if asp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asp.Type(), asp.LocalName()))
	}
	return asp.state
}

// AppServicePlanArgs contains the configurations for azurerm_app_service_plan.
type AppServicePlanArgs struct {
	// AppServiceEnvironmentId: string, optional
	AppServiceEnvironmentId terra.StringValue `hcl:"app_service_environment_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsXenon: bool, optional
	IsXenon terra.BoolValue `hcl:"is_xenon,attr"`
	// Kind: string, optional
	Kind terra.StringValue `hcl:"kind,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MaximumElasticWorkerCount: number, optional
	MaximumElasticWorkerCount terra.NumberValue `hcl:"maximum_elastic_worker_count,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PerSiteScaling: bool, optional
	PerSiteScaling terra.BoolValue `hcl:"per_site_scaling,attr"`
	// Reserved: bool, optional
	Reserved terra.BoolValue `hcl:"reserved,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ZoneRedundant: bool, optional
	ZoneRedundant terra.BoolValue `hcl:"zone_redundant,attr"`
	// Sku: required
	Sku *appserviceplan.Sku `hcl:"sku,block" validate:"required"`
	// Timeouts: optional
	Timeouts *appserviceplan.Timeouts `hcl:"timeouts,block"`
}
type appServicePlanAttributes struct {
	ref terra.Reference
}

// AppServiceEnvironmentId returns a reference to field app_service_environment_id of azurerm_app_service_plan.
func (asp appServicePlanAttributes) AppServiceEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("app_service_environment_id"))
}

// Id returns a reference to field id of azurerm_app_service_plan.
func (asp appServicePlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("id"))
}

// IsXenon returns a reference to field is_xenon of azurerm_app_service_plan.
func (asp appServicePlanAttributes) IsXenon() terra.BoolValue {
	return terra.ReferenceAsBool(asp.ref.Append("is_xenon"))
}

// Kind returns a reference to field kind of azurerm_app_service_plan.
func (asp appServicePlanAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_app_service_plan.
func (asp appServicePlanAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("location"))
}

// MaximumElasticWorkerCount returns a reference to field maximum_elastic_worker_count of azurerm_app_service_plan.
func (asp appServicePlanAttributes) MaximumElasticWorkerCount() terra.NumberValue {
	return terra.ReferenceAsNumber(asp.ref.Append("maximum_elastic_worker_count"))
}

// MaximumNumberOfWorkers returns a reference to field maximum_number_of_workers of azurerm_app_service_plan.
func (asp appServicePlanAttributes) MaximumNumberOfWorkers() terra.NumberValue {
	return terra.ReferenceAsNumber(asp.ref.Append("maximum_number_of_workers"))
}

// Name returns a reference to field name of azurerm_app_service_plan.
func (asp appServicePlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("name"))
}

// PerSiteScaling returns a reference to field per_site_scaling of azurerm_app_service_plan.
func (asp appServicePlanAttributes) PerSiteScaling() terra.BoolValue {
	return terra.ReferenceAsBool(asp.ref.Append("per_site_scaling"))
}

// Reserved returns a reference to field reserved of azurerm_app_service_plan.
func (asp appServicePlanAttributes) Reserved() terra.BoolValue {
	return terra.ReferenceAsBool(asp.ref.Append("reserved"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_app_service_plan.
func (asp appServicePlanAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_app_service_plan.
func (asp appServicePlanAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asp.ref.Append("tags"))
}

// ZoneRedundant returns a reference to field zone_redundant of azurerm_app_service_plan.
func (asp appServicePlanAttributes) ZoneRedundant() terra.BoolValue {
	return terra.ReferenceAsBool(asp.ref.Append("zone_redundant"))
}

func (asp appServicePlanAttributes) Sku() terra.ListValue[appserviceplan.SkuAttributes] {
	return terra.ReferenceAsList[appserviceplan.SkuAttributes](asp.ref.Append("sku"))
}

func (asp appServicePlanAttributes) Timeouts() appserviceplan.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appserviceplan.TimeoutsAttributes](asp.ref.Append("timeouts"))
}

type appServicePlanState struct {
	AppServiceEnvironmentId   string                        `json:"app_service_environment_id"`
	Id                        string                        `json:"id"`
	IsXenon                   bool                          `json:"is_xenon"`
	Kind                      string                        `json:"kind"`
	Location                  string                        `json:"location"`
	MaximumElasticWorkerCount float64                       `json:"maximum_elastic_worker_count"`
	MaximumNumberOfWorkers    float64                       `json:"maximum_number_of_workers"`
	Name                      string                        `json:"name"`
	PerSiteScaling            bool                          `json:"per_site_scaling"`
	Reserved                  bool                          `json:"reserved"`
	ResourceGroupName         string                        `json:"resource_group_name"`
	Tags                      map[string]string             `json:"tags"`
	ZoneRedundant             bool                          `json:"zone_redundant"`
	Sku                       []appserviceplan.SkuState     `json:"sku"`
	Timeouts                  *appserviceplan.TimeoutsState `json:"timeouts"`
}
