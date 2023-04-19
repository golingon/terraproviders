// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iottimeseriesinsightsaccesspolicy "github.com/golingon/terraproviders/azurerm/3.52.0/iottimeseriesinsightsaccesspolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIotTimeSeriesInsightsAccessPolicy creates a new instance of [IotTimeSeriesInsightsAccessPolicy].
func NewIotTimeSeriesInsightsAccessPolicy(name string, args IotTimeSeriesInsightsAccessPolicyArgs) *IotTimeSeriesInsightsAccessPolicy {
	return &IotTimeSeriesInsightsAccessPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotTimeSeriesInsightsAccessPolicy)(nil)

// IotTimeSeriesInsightsAccessPolicy represents the Terraform resource azurerm_iot_time_series_insights_access_policy.
type IotTimeSeriesInsightsAccessPolicy struct {
	Name      string
	Args      IotTimeSeriesInsightsAccessPolicyArgs
	state     *iotTimeSeriesInsightsAccessPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotTimeSeriesInsightsAccessPolicy].
func (itsiap *IotTimeSeriesInsightsAccessPolicy) Type() string {
	return "azurerm_iot_time_series_insights_access_policy"
}

// LocalName returns the local name for [IotTimeSeriesInsightsAccessPolicy].
func (itsiap *IotTimeSeriesInsightsAccessPolicy) LocalName() string {
	return itsiap.Name
}

// Configuration returns the configuration (args) for [IotTimeSeriesInsightsAccessPolicy].
func (itsiap *IotTimeSeriesInsightsAccessPolicy) Configuration() interface{} {
	return itsiap.Args
}

// DependOn is used for other resources to depend on [IotTimeSeriesInsightsAccessPolicy].
func (itsiap *IotTimeSeriesInsightsAccessPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(itsiap)
}

// Dependencies returns the list of resources [IotTimeSeriesInsightsAccessPolicy] depends_on.
func (itsiap *IotTimeSeriesInsightsAccessPolicy) Dependencies() terra.Dependencies {
	return itsiap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotTimeSeriesInsightsAccessPolicy].
func (itsiap *IotTimeSeriesInsightsAccessPolicy) LifecycleManagement() *terra.Lifecycle {
	return itsiap.Lifecycle
}

// Attributes returns the attributes for [IotTimeSeriesInsightsAccessPolicy].
func (itsiap *IotTimeSeriesInsightsAccessPolicy) Attributes() iotTimeSeriesInsightsAccessPolicyAttributes {
	return iotTimeSeriesInsightsAccessPolicyAttributes{ref: terra.ReferenceResource(itsiap)}
}

// ImportState imports the given attribute values into [IotTimeSeriesInsightsAccessPolicy]'s state.
func (itsiap *IotTimeSeriesInsightsAccessPolicy) ImportState(av io.Reader) error {
	itsiap.state = &iotTimeSeriesInsightsAccessPolicyState{}
	if err := json.NewDecoder(av).Decode(itsiap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", itsiap.Type(), itsiap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotTimeSeriesInsightsAccessPolicy] has state.
func (itsiap *IotTimeSeriesInsightsAccessPolicy) State() (*iotTimeSeriesInsightsAccessPolicyState, bool) {
	return itsiap.state, itsiap.state != nil
}

// StateMust returns the state for [IotTimeSeriesInsightsAccessPolicy]. Panics if the state is nil.
func (itsiap *IotTimeSeriesInsightsAccessPolicy) StateMust() *iotTimeSeriesInsightsAccessPolicyState {
	if itsiap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", itsiap.Type(), itsiap.LocalName()))
	}
	return itsiap.state
}

// IotTimeSeriesInsightsAccessPolicyArgs contains the configurations for azurerm_iot_time_series_insights_access_policy.
type IotTimeSeriesInsightsAccessPolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrincipalObjectId: string, required
	PrincipalObjectId terra.StringValue `hcl:"principal_object_id,attr" validate:"required"`
	// Roles: set of string, required
	Roles terra.SetValue[terra.StringValue] `hcl:"roles,attr" validate:"required"`
	// TimeSeriesInsightsEnvironmentId: string, required
	TimeSeriesInsightsEnvironmentId terra.StringValue `hcl:"time_series_insights_environment_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *iottimeseriesinsightsaccesspolicy.Timeouts `hcl:"timeouts,block"`
}
type iotTimeSeriesInsightsAccessPolicyAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_iot_time_series_insights_access_policy.
func (itsiap iotTimeSeriesInsightsAccessPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(itsiap.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_iot_time_series_insights_access_policy.
func (itsiap iotTimeSeriesInsightsAccessPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(itsiap.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_iot_time_series_insights_access_policy.
func (itsiap iotTimeSeriesInsightsAccessPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(itsiap.ref.Append("name"))
}

// PrincipalObjectId returns a reference to field principal_object_id of azurerm_iot_time_series_insights_access_policy.
func (itsiap iotTimeSeriesInsightsAccessPolicyAttributes) PrincipalObjectId() terra.StringValue {
	return terra.ReferenceAsString(itsiap.ref.Append("principal_object_id"))
}

// Roles returns a reference to field roles of azurerm_iot_time_series_insights_access_policy.
func (itsiap iotTimeSeriesInsightsAccessPolicyAttributes) Roles() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](itsiap.ref.Append("roles"))
}

// TimeSeriesInsightsEnvironmentId returns a reference to field time_series_insights_environment_id of azurerm_iot_time_series_insights_access_policy.
func (itsiap iotTimeSeriesInsightsAccessPolicyAttributes) TimeSeriesInsightsEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(itsiap.ref.Append("time_series_insights_environment_id"))
}

func (itsiap iotTimeSeriesInsightsAccessPolicyAttributes) Timeouts() iottimeseriesinsightsaccesspolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iottimeseriesinsightsaccesspolicy.TimeoutsAttributes](itsiap.ref.Append("timeouts"))
}

type iotTimeSeriesInsightsAccessPolicyState struct {
	Description                     string                                           `json:"description"`
	Id                              string                                           `json:"id"`
	Name                            string                                           `json:"name"`
	PrincipalObjectId               string                                           `json:"principal_object_id"`
	Roles                           []string                                         `json:"roles"`
	TimeSeriesInsightsEnvironmentId string                                           `json:"time_series_insights_environment_id"`
	Timeouts                        *iottimeseriesinsightsaccesspolicy.TimeoutsState `json:"timeouts"`
}
