// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ec2fleet "github.com/golingon/terraproviders/aws/4.63.0/ec2fleet"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2Fleet creates a new instance of [Ec2Fleet].
func NewEc2Fleet(name string, args Ec2FleetArgs) *Ec2Fleet {
	return &Ec2Fleet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2Fleet)(nil)

// Ec2Fleet represents the Terraform resource aws_ec2_fleet.
type Ec2Fleet struct {
	Name      string
	Args      Ec2FleetArgs
	state     *ec2FleetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2Fleet].
func (ef *Ec2Fleet) Type() string {
	return "aws_ec2_fleet"
}

// LocalName returns the local name for [Ec2Fleet].
func (ef *Ec2Fleet) LocalName() string {
	return ef.Name
}

// Configuration returns the configuration (args) for [Ec2Fleet].
func (ef *Ec2Fleet) Configuration() interface{} {
	return ef.Args
}

// DependOn is used for other resources to depend on [Ec2Fleet].
func (ef *Ec2Fleet) DependOn() terra.Reference {
	return terra.ReferenceResource(ef)
}

// Dependencies returns the list of resources [Ec2Fleet] depends_on.
func (ef *Ec2Fleet) Dependencies() terra.Dependencies {
	return ef.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2Fleet].
func (ef *Ec2Fleet) LifecycleManagement() *terra.Lifecycle {
	return ef.Lifecycle
}

// Attributes returns the attributes for [Ec2Fleet].
func (ef *Ec2Fleet) Attributes() ec2FleetAttributes {
	return ec2FleetAttributes{ref: terra.ReferenceResource(ef)}
}

// ImportState imports the given attribute values into [Ec2Fleet]'s state.
func (ef *Ec2Fleet) ImportState(av io.Reader) error {
	ef.state = &ec2FleetState{}
	if err := json.NewDecoder(av).Decode(ef.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ef.Type(), ef.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2Fleet] has state.
func (ef *Ec2Fleet) State() (*ec2FleetState, bool) {
	return ef.state, ef.state != nil
}

// StateMust returns the state for [Ec2Fleet]. Panics if the state is nil.
func (ef *Ec2Fleet) StateMust() *ec2FleetState {
	if ef.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ef.Type(), ef.LocalName()))
	}
	return ef.state
}

// Ec2FleetArgs contains the configurations for aws_ec2_fleet.
type Ec2FleetArgs struct {
	// Context: string, optional
	Context terra.StringValue `hcl:"context,attr"`
	// ExcessCapacityTerminationPolicy: string, optional
	ExcessCapacityTerminationPolicy terra.StringValue `hcl:"excess_capacity_termination_policy,attr"`
	// FleetState: string, optional
	FleetState terra.StringValue `hcl:"fleet_state,attr"`
	// FulfilledCapacity: number, optional
	FulfilledCapacity terra.NumberValue `hcl:"fulfilled_capacity,attr"`
	// FulfilledOnDemandCapacity: number, optional
	FulfilledOnDemandCapacity terra.NumberValue `hcl:"fulfilled_on_demand_capacity,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ReplaceUnhealthyInstances: bool, optional
	ReplaceUnhealthyInstances terra.BoolValue `hcl:"replace_unhealthy_instances,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TerminateInstances: bool, optional
	TerminateInstances terra.BoolValue `hcl:"terminate_instances,attr"`
	// TerminateInstancesWithExpiration: bool, optional
	TerminateInstancesWithExpiration terra.BoolValue `hcl:"terminate_instances_with_expiration,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// ValidFrom: string, optional
	ValidFrom terra.StringValue `hcl:"valid_from,attr"`
	// ValidUntil: string, optional
	ValidUntil terra.StringValue `hcl:"valid_until,attr"`
	// FleetInstanceSet: min=0
	FleetInstanceSet []ec2fleet.FleetInstanceSet `hcl:"fleet_instance_set,block" validate:"min=0"`
	// LaunchTemplateConfig: min=1,max=50
	LaunchTemplateConfig []ec2fleet.LaunchTemplateConfig `hcl:"launch_template_config,block" validate:"min=1,max=50"`
	// OnDemandOptions: optional
	OnDemandOptions *ec2fleet.OnDemandOptions `hcl:"on_demand_options,block"`
	// SpotOptions: optional
	SpotOptions *ec2fleet.SpotOptions `hcl:"spot_options,block"`
	// TargetCapacitySpecification: required
	TargetCapacitySpecification *ec2fleet.TargetCapacitySpecification `hcl:"target_capacity_specification,block" validate:"required"`
	// Timeouts: optional
	Timeouts *ec2fleet.Timeouts `hcl:"timeouts,block"`
}
type ec2FleetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_fleet.
func (ef ec2FleetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("arn"))
}

// Context returns a reference to field context of aws_ec2_fleet.
func (ef ec2FleetAttributes) Context() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("context"))
}

// ExcessCapacityTerminationPolicy returns a reference to field excess_capacity_termination_policy of aws_ec2_fleet.
func (ef ec2FleetAttributes) ExcessCapacityTerminationPolicy() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("excess_capacity_termination_policy"))
}

// FleetState returns a reference to field fleet_state of aws_ec2_fleet.
func (ef ec2FleetAttributes) FleetState() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("fleet_state"))
}

// FulfilledCapacity returns a reference to field fulfilled_capacity of aws_ec2_fleet.
func (ef ec2FleetAttributes) FulfilledCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(ef.ref.Append("fulfilled_capacity"))
}

// FulfilledOnDemandCapacity returns a reference to field fulfilled_on_demand_capacity of aws_ec2_fleet.
func (ef ec2FleetAttributes) FulfilledOnDemandCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(ef.ref.Append("fulfilled_on_demand_capacity"))
}

// Id returns a reference to field id of aws_ec2_fleet.
func (ef ec2FleetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("id"))
}

// ReplaceUnhealthyInstances returns a reference to field replace_unhealthy_instances of aws_ec2_fleet.
func (ef ec2FleetAttributes) ReplaceUnhealthyInstances() terra.BoolValue {
	return terra.ReferenceAsBool(ef.ref.Append("replace_unhealthy_instances"))
}

// Tags returns a reference to field tags of aws_ec2_fleet.
func (ef ec2FleetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ef.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ec2_fleet.
func (ef ec2FleetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ef.ref.Append("tags_all"))
}

// TerminateInstances returns a reference to field terminate_instances of aws_ec2_fleet.
func (ef ec2FleetAttributes) TerminateInstances() terra.BoolValue {
	return terra.ReferenceAsBool(ef.ref.Append("terminate_instances"))
}

// TerminateInstancesWithExpiration returns a reference to field terminate_instances_with_expiration of aws_ec2_fleet.
func (ef ec2FleetAttributes) TerminateInstancesWithExpiration() terra.BoolValue {
	return terra.ReferenceAsBool(ef.ref.Append("terminate_instances_with_expiration"))
}

// Type returns a reference to field type of aws_ec2_fleet.
func (ef ec2FleetAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("type"))
}

// ValidFrom returns a reference to field valid_from of aws_ec2_fleet.
func (ef ec2FleetAttributes) ValidFrom() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("valid_from"))
}

// ValidUntil returns a reference to field valid_until of aws_ec2_fleet.
func (ef ec2FleetAttributes) ValidUntil() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("valid_until"))
}

func (ef ec2FleetAttributes) FleetInstanceSet() terra.ListValue[ec2fleet.FleetInstanceSetAttributes] {
	return terra.ReferenceAsList[ec2fleet.FleetInstanceSetAttributes](ef.ref.Append("fleet_instance_set"))
}

func (ef ec2FleetAttributes) LaunchTemplateConfig() terra.ListValue[ec2fleet.LaunchTemplateConfigAttributes] {
	return terra.ReferenceAsList[ec2fleet.LaunchTemplateConfigAttributes](ef.ref.Append("launch_template_config"))
}

func (ef ec2FleetAttributes) OnDemandOptions() terra.ListValue[ec2fleet.OnDemandOptionsAttributes] {
	return terra.ReferenceAsList[ec2fleet.OnDemandOptionsAttributes](ef.ref.Append("on_demand_options"))
}

func (ef ec2FleetAttributes) SpotOptions() terra.ListValue[ec2fleet.SpotOptionsAttributes] {
	return terra.ReferenceAsList[ec2fleet.SpotOptionsAttributes](ef.ref.Append("spot_options"))
}

func (ef ec2FleetAttributes) TargetCapacitySpecification() terra.ListValue[ec2fleet.TargetCapacitySpecificationAttributes] {
	return terra.ReferenceAsList[ec2fleet.TargetCapacitySpecificationAttributes](ef.ref.Append("target_capacity_specification"))
}

func (ef ec2FleetAttributes) Timeouts() ec2fleet.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ec2fleet.TimeoutsAttributes](ef.ref.Append("timeouts"))
}

type ec2FleetState struct {
	Arn                              string                                      `json:"arn"`
	Context                          string                                      `json:"context"`
	ExcessCapacityTerminationPolicy  string                                      `json:"excess_capacity_termination_policy"`
	FleetState                       string                                      `json:"fleet_state"`
	FulfilledCapacity                float64                                     `json:"fulfilled_capacity"`
	FulfilledOnDemandCapacity        float64                                     `json:"fulfilled_on_demand_capacity"`
	Id                               string                                      `json:"id"`
	ReplaceUnhealthyInstances        bool                                        `json:"replace_unhealthy_instances"`
	Tags                             map[string]string                           `json:"tags"`
	TagsAll                          map[string]string                           `json:"tags_all"`
	TerminateInstances               bool                                        `json:"terminate_instances"`
	TerminateInstancesWithExpiration bool                                        `json:"terminate_instances_with_expiration"`
	Type                             string                                      `json:"type"`
	ValidFrom                        string                                      `json:"valid_from"`
	ValidUntil                       string                                      `json:"valid_until"`
	FleetInstanceSet                 []ec2fleet.FleetInstanceSetState            `json:"fleet_instance_set"`
	LaunchTemplateConfig             []ec2fleet.LaunchTemplateConfigState        `json:"launch_template_config"`
	OnDemandOptions                  []ec2fleet.OnDemandOptionsState             `json:"on_demand_options"`
	SpotOptions                      []ec2fleet.SpotOptionsState                 `json:"spot_options"`
	TargetCapacitySpecification      []ec2fleet.TargetCapacitySpecificationState `json:"target_capacity_specification"`
	Timeouts                         *ec2fleet.TimeoutsState                     `json:"timeouts"`
}
