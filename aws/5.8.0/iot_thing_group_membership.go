// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIotThingGroupMembership creates a new instance of [IotThingGroupMembership].
func NewIotThingGroupMembership(name string, args IotThingGroupMembershipArgs) *IotThingGroupMembership {
	return &IotThingGroupMembership{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotThingGroupMembership)(nil)

// IotThingGroupMembership represents the Terraform resource aws_iot_thing_group_membership.
type IotThingGroupMembership struct {
	Name      string
	Args      IotThingGroupMembershipArgs
	state     *iotThingGroupMembershipState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotThingGroupMembership].
func (itgm *IotThingGroupMembership) Type() string {
	return "aws_iot_thing_group_membership"
}

// LocalName returns the local name for [IotThingGroupMembership].
func (itgm *IotThingGroupMembership) LocalName() string {
	return itgm.Name
}

// Configuration returns the configuration (args) for [IotThingGroupMembership].
func (itgm *IotThingGroupMembership) Configuration() interface{} {
	return itgm.Args
}

// DependOn is used for other resources to depend on [IotThingGroupMembership].
func (itgm *IotThingGroupMembership) DependOn() terra.Reference {
	return terra.ReferenceResource(itgm)
}

// Dependencies returns the list of resources [IotThingGroupMembership] depends_on.
func (itgm *IotThingGroupMembership) Dependencies() terra.Dependencies {
	return itgm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotThingGroupMembership].
func (itgm *IotThingGroupMembership) LifecycleManagement() *terra.Lifecycle {
	return itgm.Lifecycle
}

// Attributes returns the attributes for [IotThingGroupMembership].
func (itgm *IotThingGroupMembership) Attributes() iotThingGroupMembershipAttributes {
	return iotThingGroupMembershipAttributes{ref: terra.ReferenceResource(itgm)}
}

// ImportState imports the given attribute values into [IotThingGroupMembership]'s state.
func (itgm *IotThingGroupMembership) ImportState(av io.Reader) error {
	itgm.state = &iotThingGroupMembershipState{}
	if err := json.NewDecoder(av).Decode(itgm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", itgm.Type(), itgm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotThingGroupMembership] has state.
func (itgm *IotThingGroupMembership) State() (*iotThingGroupMembershipState, bool) {
	return itgm.state, itgm.state != nil
}

// StateMust returns the state for [IotThingGroupMembership]. Panics if the state is nil.
func (itgm *IotThingGroupMembership) StateMust() *iotThingGroupMembershipState {
	if itgm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", itgm.Type(), itgm.LocalName()))
	}
	return itgm.state
}

// IotThingGroupMembershipArgs contains the configurations for aws_iot_thing_group_membership.
type IotThingGroupMembershipArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OverrideDynamicGroup: bool, optional
	OverrideDynamicGroup terra.BoolValue `hcl:"override_dynamic_group,attr"`
	// ThingGroupName: string, required
	ThingGroupName terra.StringValue `hcl:"thing_group_name,attr" validate:"required"`
	// ThingName: string, required
	ThingName terra.StringValue `hcl:"thing_name,attr" validate:"required"`
}
type iotThingGroupMembershipAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_iot_thing_group_membership.
func (itgm iotThingGroupMembershipAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(itgm.ref.Append("id"))
}

// OverrideDynamicGroup returns a reference to field override_dynamic_group of aws_iot_thing_group_membership.
func (itgm iotThingGroupMembershipAttributes) OverrideDynamicGroup() terra.BoolValue {
	return terra.ReferenceAsBool(itgm.ref.Append("override_dynamic_group"))
}

// ThingGroupName returns a reference to field thing_group_name of aws_iot_thing_group_membership.
func (itgm iotThingGroupMembershipAttributes) ThingGroupName() terra.StringValue {
	return terra.ReferenceAsString(itgm.ref.Append("thing_group_name"))
}

// ThingName returns a reference to field thing_name of aws_iot_thing_group_membership.
func (itgm iotThingGroupMembershipAttributes) ThingName() terra.StringValue {
	return terra.ReferenceAsString(itgm.ref.Append("thing_name"))
}

type iotThingGroupMembershipState struct {
	Id                   string `json:"id"`
	OverrideDynamicGroup bool   `json:"override_dynamic_group"`
	ThingGroupName       string `json:"thing_group_name"`
	ThingName            string `json:"thing_name"`
}
