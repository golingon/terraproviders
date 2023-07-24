// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	spotfleetrequest "github.com/golingon/terraproviders/aws/5.9.0/spotfleetrequest"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpotFleetRequest creates a new instance of [SpotFleetRequest].
func NewSpotFleetRequest(name string, args SpotFleetRequestArgs) *SpotFleetRequest {
	return &SpotFleetRequest{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpotFleetRequest)(nil)

// SpotFleetRequest represents the Terraform resource aws_spot_fleet_request.
type SpotFleetRequest struct {
	Name      string
	Args      SpotFleetRequestArgs
	state     *spotFleetRequestState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpotFleetRequest].
func (sfr *SpotFleetRequest) Type() string {
	return "aws_spot_fleet_request"
}

// LocalName returns the local name for [SpotFleetRequest].
func (sfr *SpotFleetRequest) LocalName() string {
	return sfr.Name
}

// Configuration returns the configuration (args) for [SpotFleetRequest].
func (sfr *SpotFleetRequest) Configuration() interface{} {
	return sfr.Args
}

// DependOn is used for other resources to depend on [SpotFleetRequest].
func (sfr *SpotFleetRequest) DependOn() terra.Reference {
	return terra.ReferenceResource(sfr)
}

// Dependencies returns the list of resources [SpotFleetRequest] depends_on.
func (sfr *SpotFleetRequest) Dependencies() terra.Dependencies {
	return sfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpotFleetRequest].
func (sfr *SpotFleetRequest) LifecycleManagement() *terra.Lifecycle {
	return sfr.Lifecycle
}

// Attributes returns the attributes for [SpotFleetRequest].
func (sfr *SpotFleetRequest) Attributes() spotFleetRequestAttributes {
	return spotFleetRequestAttributes{ref: terra.ReferenceResource(sfr)}
}

// ImportState imports the given attribute values into [SpotFleetRequest]'s state.
func (sfr *SpotFleetRequest) ImportState(av io.Reader) error {
	sfr.state = &spotFleetRequestState{}
	if err := json.NewDecoder(av).Decode(sfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sfr.Type(), sfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpotFleetRequest] has state.
func (sfr *SpotFleetRequest) State() (*spotFleetRequestState, bool) {
	return sfr.state, sfr.state != nil
}

// StateMust returns the state for [SpotFleetRequest]. Panics if the state is nil.
func (sfr *SpotFleetRequest) StateMust() *spotFleetRequestState {
	if sfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sfr.Type(), sfr.LocalName()))
	}
	return sfr.state
}

// SpotFleetRequestArgs contains the configurations for aws_spot_fleet_request.
type SpotFleetRequestArgs struct {
	// AllocationStrategy: string, optional
	AllocationStrategy terra.StringValue `hcl:"allocation_strategy,attr"`
	// Context: string, optional
	Context terra.StringValue `hcl:"context,attr"`
	// ExcessCapacityTerminationPolicy: string, optional
	ExcessCapacityTerminationPolicy terra.StringValue `hcl:"excess_capacity_termination_policy,attr"`
	// FleetType: string, optional
	FleetType terra.StringValue `hcl:"fleet_type,attr"`
	// IamFleetRole: string, required
	IamFleetRole terra.StringValue `hcl:"iam_fleet_role,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceInterruptionBehaviour: string, optional
	InstanceInterruptionBehaviour terra.StringValue `hcl:"instance_interruption_behaviour,attr"`
	// InstancePoolsToUseCount: number, optional
	InstancePoolsToUseCount terra.NumberValue `hcl:"instance_pools_to_use_count,attr"`
	// LoadBalancers: set of string, optional
	LoadBalancers terra.SetValue[terra.StringValue] `hcl:"load_balancers,attr"`
	// OnDemandAllocationStrategy: string, optional
	OnDemandAllocationStrategy terra.StringValue `hcl:"on_demand_allocation_strategy,attr"`
	// OnDemandMaxTotalPrice: string, optional
	OnDemandMaxTotalPrice terra.StringValue `hcl:"on_demand_max_total_price,attr"`
	// OnDemandTargetCapacity: number, optional
	OnDemandTargetCapacity terra.NumberValue `hcl:"on_demand_target_capacity,attr"`
	// ReplaceUnhealthyInstances: bool, optional
	ReplaceUnhealthyInstances terra.BoolValue `hcl:"replace_unhealthy_instances,attr"`
	// SpotPrice: string, optional
	SpotPrice terra.StringValue `hcl:"spot_price,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TargetCapacity: number, required
	TargetCapacity terra.NumberValue `hcl:"target_capacity,attr" validate:"required"`
	// TargetCapacityUnitType: string, optional
	TargetCapacityUnitType terra.StringValue `hcl:"target_capacity_unit_type,attr"`
	// TargetGroupArns: set of string, optional
	TargetGroupArns terra.SetValue[terra.StringValue] `hcl:"target_group_arns,attr"`
	// TerminateInstancesOnDelete: string, optional
	TerminateInstancesOnDelete terra.StringValue `hcl:"terminate_instances_on_delete,attr"`
	// TerminateInstancesWithExpiration: bool, optional
	TerminateInstancesWithExpiration terra.BoolValue `hcl:"terminate_instances_with_expiration,attr"`
	// ValidFrom: string, optional
	ValidFrom terra.StringValue `hcl:"valid_from,attr"`
	// ValidUntil: string, optional
	ValidUntil terra.StringValue `hcl:"valid_until,attr"`
	// WaitForFulfillment: bool, optional
	WaitForFulfillment terra.BoolValue `hcl:"wait_for_fulfillment,attr"`
	// LaunchSpecification: min=0
	LaunchSpecification []spotfleetrequest.LaunchSpecification `hcl:"launch_specification,block" validate:"min=0"`
	// LaunchTemplateConfig: min=0
	LaunchTemplateConfig []spotfleetrequest.LaunchTemplateConfig `hcl:"launch_template_config,block" validate:"min=0"`
	// SpotMaintenanceStrategies: optional
	SpotMaintenanceStrategies *spotfleetrequest.SpotMaintenanceStrategies `hcl:"spot_maintenance_strategies,block"`
	// Timeouts: optional
	Timeouts *spotfleetrequest.Timeouts `hcl:"timeouts,block"`
}
type spotFleetRequestAttributes struct {
	ref terra.Reference
}

// AllocationStrategy returns a reference to field allocation_strategy of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) AllocationStrategy() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("allocation_strategy"))
}

// ClientToken returns a reference to field client_token of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) ClientToken() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("client_token"))
}

// Context returns a reference to field context of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) Context() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("context"))
}

// ExcessCapacityTerminationPolicy returns a reference to field excess_capacity_termination_policy of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) ExcessCapacityTerminationPolicy() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("excess_capacity_termination_policy"))
}

// FleetType returns a reference to field fleet_type of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) FleetType() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("fleet_type"))
}

// IamFleetRole returns a reference to field iam_fleet_role of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) IamFleetRole() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("iam_fleet_role"))
}

// Id returns a reference to field id of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("id"))
}

// InstanceInterruptionBehaviour returns a reference to field instance_interruption_behaviour of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) InstanceInterruptionBehaviour() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("instance_interruption_behaviour"))
}

// InstancePoolsToUseCount returns a reference to field instance_pools_to_use_count of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) InstancePoolsToUseCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sfr.ref.Append("instance_pools_to_use_count"))
}

// LoadBalancers returns a reference to field load_balancers of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) LoadBalancers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sfr.ref.Append("load_balancers"))
}

// OnDemandAllocationStrategy returns a reference to field on_demand_allocation_strategy of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) OnDemandAllocationStrategy() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("on_demand_allocation_strategy"))
}

// OnDemandMaxTotalPrice returns a reference to field on_demand_max_total_price of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) OnDemandMaxTotalPrice() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("on_demand_max_total_price"))
}

// OnDemandTargetCapacity returns a reference to field on_demand_target_capacity of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) OnDemandTargetCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(sfr.ref.Append("on_demand_target_capacity"))
}

// ReplaceUnhealthyInstances returns a reference to field replace_unhealthy_instances of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) ReplaceUnhealthyInstances() terra.BoolValue {
	return terra.ReferenceAsBool(sfr.ref.Append("replace_unhealthy_instances"))
}

// SpotPrice returns a reference to field spot_price of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) SpotPrice() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("spot_price"))
}

// SpotRequestState returns a reference to field spot_request_state of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) SpotRequestState() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("spot_request_state"))
}

// Tags returns a reference to field tags of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sfr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sfr.ref.Append("tags_all"))
}

// TargetCapacity returns a reference to field target_capacity of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) TargetCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(sfr.ref.Append("target_capacity"))
}

// TargetCapacityUnitType returns a reference to field target_capacity_unit_type of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) TargetCapacityUnitType() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("target_capacity_unit_type"))
}

// TargetGroupArns returns a reference to field target_group_arns of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) TargetGroupArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sfr.ref.Append("target_group_arns"))
}

// TerminateInstancesOnDelete returns a reference to field terminate_instances_on_delete of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) TerminateInstancesOnDelete() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("terminate_instances_on_delete"))
}

// TerminateInstancesWithExpiration returns a reference to field terminate_instances_with_expiration of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) TerminateInstancesWithExpiration() terra.BoolValue {
	return terra.ReferenceAsBool(sfr.ref.Append("terminate_instances_with_expiration"))
}

// ValidFrom returns a reference to field valid_from of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) ValidFrom() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("valid_from"))
}

// ValidUntil returns a reference to field valid_until of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) ValidUntil() terra.StringValue {
	return terra.ReferenceAsString(sfr.ref.Append("valid_until"))
}

// WaitForFulfillment returns a reference to field wait_for_fulfillment of aws_spot_fleet_request.
func (sfr spotFleetRequestAttributes) WaitForFulfillment() terra.BoolValue {
	return terra.ReferenceAsBool(sfr.ref.Append("wait_for_fulfillment"))
}

func (sfr spotFleetRequestAttributes) LaunchSpecification() terra.SetValue[spotfleetrequest.LaunchSpecificationAttributes] {
	return terra.ReferenceAsSet[spotfleetrequest.LaunchSpecificationAttributes](sfr.ref.Append("launch_specification"))
}

func (sfr spotFleetRequestAttributes) LaunchTemplateConfig() terra.SetValue[spotfleetrequest.LaunchTemplateConfigAttributes] {
	return terra.ReferenceAsSet[spotfleetrequest.LaunchTemplateConfigAttributes](sfr.ref.Append("launch_template_config"))
}

func (sfr spotFleetRequestAttributes) SpotMaintenanceStrategies() terra.ListValue[spotfleetrequest.SpotMaintenanceStrategiesAttributes] {
	return terra.ReferenceAsList[spotfleetrequest.SpotMaintenanceStrategiesAttributes](sfr.ref.Append("spot_maintenance_strategies"))
}

func (sfr spotFleetRequestAttributes) Timeouts() spotfleetrequest.TimeoutsAttributes {
	return terra.ReferenceAsSingle[spotfleetrequest.TimeoutsAttributes](sfr.ref.Append("timeouts"))
}

type spotFleetRequestState struct {
	AllocationStrategy               string                                            `json:"allocation_strategy"`
	ClientToken                      string                                            `json:"client_token"`
	Context                          string                                            `json:"context"`
	ExcessCapacityTerminationPolicy  string                                            `json:"excess_capacity_termination_policy"`
	FleetType                        string                                            `json:"fleet_type"`
	IamFleetRole                     string                                            `json:"iam_fleet_role"`
	Id                               string                                            `json:"id"`
	InstanceInterruptionBehaviour    string                                            `json:"instance_interruption_behaviour"`
	InstancePoolsToUseCount          float64                                           `json:"instance_pools_to_use_count"`
	LoadBalancers                    []string                                          `json:"load_balancers"`
	OnDemandAllocationStrategy       string                                            `json:"on_demand_allocation_strategy"`
	OnDemandMaxTotalPrice            string                                            `json:"on_demand_max_total_price"`
	OnDemandTargetCapacity           float64                                           `json:"on_demand_target_capacity"`
	ReplaceUnhealthyInstances        bool                                              `json:"replace_unhealthy_instances"`
	SpotPrice                        string                                            `json:"spot_price"`
	SpotRequestState                 string                                            `json:"spot_request_state"`
	Tags                             map[string]string                                 `json:"tags"`
	TagsAll                          map[string]string                                 `json:"tags_all"`
	TargetCapacity                   float64                                           `json:"target_capacity"`
	TargetCapacityUnitType           string                                            `json:"target_capacity_unit_type"`
	TargetGroupArns                  []string                                          `json:"target_group_arns"`
	TerminateInstancesOnDelete       string                                            `json:"terminate_instances_on_delete"`
	TerminateInstancesWithExpiration bool                                              `json:"terminate_instances_with_expiration"`
	ValidFrom                        string                                            `json:"valid_from"`
	ValidUntil                       string                                            `json:"valid_until"`
	WaitForFulfillment               bool                                              `json:"wait_for_fulfillment"`
	LaunchSpecification              []spotfleetrequest.LaunchSpecificationState       `json:"launch_specification"`
	LaunchTemplateConfig             []spotfleetrequest.LaunchTemplateConfigState      `json:"launch_template_config"`
	SpotMaintenanceStrategies        []spotfleetrequest.SpotMaintenanceStrategiesState `json:"spot_maintenance_strategies"`
	Timeouts                         *spotfleetrequest.TimeoutsState                   `json:"timeouts"`
}
