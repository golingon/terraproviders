// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	serviceusageconsumerquotaoverride "github.com/golingon/terraproviders/googlebeta/4.72.1/serviceusageconsumerquotaoverride"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServiceUsageConsumerQuotaOverride creates a new instance of [ServiceUsageConsumerQuotaOverride].
func NewServiceUsageConsumerQuotaOverride(name string, args ServiceUsageConsumerQuotaOverrideArgs) *ServiceUsageConsumerQuotaOverride {
	return &ServiceUsageConsumerQuotaOverride{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceUsageConsumerQuotaOverride)(nil)

// ServiceUsageConsumerQuotaOverride represents the Terraform resource google_service_usage_consumer_quota_override.
type ServiceUsageConsumerQuotaOverride struct {
	Name      string
	Args      ServiceUsageConsumerQuotaOverrideArgs
	state     *serviceUsageConsumerQuotaOverrideState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServiceUsageConsumerQuotaOverride].
func (sucqo *ServiceUsageConsumerQuotaOverride) Type() string {
	return "google_service_usage_consumer_quota_override"
}

// LocalName returns the local name for [ServiceUsageConsumerQuotaOverride].
func (sucqo *ServiceUsageConsumerQuotaOverride) LocalName() string {
	return sucqo.Name
}

// Configuration returns the configuration (args) for [ServiceUsageConsumerQuotaOverride].
func (sucqo *ServiceUsageConsumerQuotaOverride) Configuration() interface{} {
	return sucqo.Args
}

// DependOn is used for other resources to depend on [ServiceUsageConsumerQuotaOverride].
func (sucqo *ServiceUsageConsumerQuotaOverride) DependOn() terra.Reference {
	return terra.ReferenceResource(sucqo)
}

// Dependencies returns the list of resources [ServiceUsageConsumerQuotaOverride] depends_on.
func (sucqo *ServiceUsageConsumerQuotaOverride) Dependencies() terra.Dependencies {
	return sucqo.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServiceUsageConsumerQuotaOverride].
func (sucqo *ServiceUsageConsumerQuotaOverride) LifecycleManagement() *terra.Lifecycle {
	return sucqo.Lifecycle
}

// Attributes returns the attributes for [ServiceUsageConsumerQuotaOverride].
func (sucqo *ServiceUsageConsumerQuotaOverride) Attributes() serviceUsageConsumerQuotaOverrideAttributes {
	return serviceUsageConsumerQuotaOverrideAttributes{ref: terra.ReferenceResource(sucqo)}
}

// ImportState imports the given attribute values into [ServiceUsageConsumerQuotaOverride]'s state.
func (sucqo *ServiceUsageConsumerQuotaOverride) ImportState(av io.Reader) error {
	sucqo.state = &serviceUsageConsumerQuotaOverrideState{}
	if err := json.NewDecoder(av).Decode(sucqo.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sucqo.Type(), sucqo.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServiceUsageConsumerQuotaOverride] has state.
func (sucqo *ServiceUsageConsumerQuotaOverride) State() (*serviceUsageConsumerQuotaOverrideState, bool) {
	return sucqo.state, sucqo.state != nil
}

// StateMust returns the state for [ServiceUsageConsumerQuotaOverride]. Panics if the state is nil.
func (sucqo *ServiceUsageConsumerQuotaOverride) StateMust() *serviceUsageConsumerQuotaOverrideState {
	if sucqo.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sucqo.Type(), sucqo.LocalName()))
	}
	return sucqo.state
}

// ServiceUsageConsumerQuotaOverrideArgs contains the configurations for google_service_usage_consumer_quota_override.
type ServiceUsageConsumerQuotaOverrideArgs struct {
	// Dimensions: map of string, optional
	Dimensions terra.MapValue[terra.StringValue] `hcl:"dimensions,attr"`
	// Force: bool, optional
	Force terra.BoolValue `hcl:"force,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Limit: string, required
	Limit terra.StringValue `hcl:"limit,attr" validate:"required"`
	// Metric: string, required
	Metric terra.StringValue `hcl:"metric,attr" validate:"required"`
	// OverrideValue: string, required
	OverrideValue terra.StringValue `hcl:"override_value,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *serviceusageconsumerquotaoverride.Timeouts `hcl:"timeouts,block"`
}
type serviceUsageConsumerQuotaOverrideAttributes struct {
	ref terra.Reference
}

// Dimensions returns a reference to field dimensions of google_service_usage_consumer_quota_override.
func (sucqo serviceUsageConsumerQuotaOverrideAttributes) Dimensions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sucqo.ref.Append("dimensions"))
}

// Force returns a reference to field force of google_service_usage_consumer_quota_override.
func (sucqo serviceUsageConsumerQuotaOverrideAttributes) Force() terra.BoolValue {
	return terra.ReferenceAsBool(sucqo.ref.Append("force"))
}

// Id returns a reference to field id of google_service_usage_consumer_quota_override.
func (sucqo serviceUsageConsumerQuotaOverrideAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sucqo.ref.Append("id"))
}

// Limit returns a reference to field limit of google_service_usage_consumer_quota_override.
func (sucqo serviceUsageConsumerQuotaOverrideAttributes) Limit() terra.StringValue {
	return terra.ReferenceAsString(sucqo.ref.Append("limit"))
}

// Metric returns a reference to field metric of google_service_usage_consumer_quota_override.
func (sucqo serviceUsageConsumerQuotaOverrideAttributes) Metric() terra.StringValue {
	return terra.ReferenceAsString(sucqo.ref.Append("metric"))
}

// Name returns a reference to field name of google_service_usage_consumer_quota_override.
func (sucqo serviceUsageConsumerQuotaOverrideAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sucqo.ref.Append("name"))
}

// OverrideValue returns a reference to field override_value of google_service_usage_consumer_quota_override.
func (sucqo serviceUsageConsumerQuotaOverrideAttributes) OverrideValue() terra.StringValue {
	return terra.ReferenceAsString(sucqo.ref.Append("override_value"))
}

// Project returns a reference to field project of google_service_usage_consumer_quota_override.
func (sucqo serviceUsageConsumerQuotaOverrideAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(sucqo.ref.Append("project"))
}

// Service returns a reference to field service of google_service_usage_consumer_quota_override.
func (sucqo serviceUsageConsumerQuotaOverrideAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(sucqo.ref.Append("service"))
}

func (sucqo serviceUsageConsumerQuotaOverrideAttributes) Timeouts() serviceusageconsumerquotaoverride.TimeoutsAttributes {
	return terra.ReferenceAsSingle[serviceusageconsumerquotaoverride.TimeoutsAttributes](sucqo.ref.Append("timeouts"))
}

type serviceUsageConsumerQuotaOverrideState struct {
	Dimensions    map[string]string                                `json:"dimensions"`
	Force         bool                                             `json:"force"`
	Id            string                                           `json:"id"`
	Limit         string                                           `json:"limit"`
	Metric        string                                           `json:"metric"`
	Name          string                                           `json:"name"`
	OverrideValue string                                           `json:"override_value"`
	Project       string                                           `json:"project"`
	Service       string                                           `json:"service"`
	Timeouts      *serviceusageconsumerquotaoverride.TimeoutsState `json:"timeouts"`
}
