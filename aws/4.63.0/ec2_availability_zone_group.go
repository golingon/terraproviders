// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2AvailabilityZoneGroup creates a new instance of [Ec2AvailabilityZoneGroup].
func NewEc2AvailabilityZoneGroup(name string, args Ec2AvailabilityZoneGroupArgs) *Ec2AvailabilityZoneGroup {
	return &Ec2AvailabilityZoneGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2AvailabilityZoneGroup)(nil)

// Ec2AvailabilityZoneGroup represents the Terraform resource aws_ec2_availability_zone_group.
type Ec2AvailabilityZoneGroup struct {
	Name      string
	Args      Ec2AvailabilityZoneGroupArgs
	state     *ec2AvailabilityZoneGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2AvailabilityZoneGroup].
func (eazg *Ec2AvailabilityZoneGroup) Type() string {
	return "aws_ec2_availability_zone_group"
}

// LocalName returns the local name for [Ec2AvailabilityZoneGroup].
func (eazg *Ec2AvailabilityZoneGroup) LocalName() string {
	return eazg.Name
}

// Configuration returns the configuration (args) for [Ec2AvailabilityZoneGroup].
func (eazg *Ec2AvailabilityZoneGroup) Configuration() interface{} {
	return eazg.Args
}

// DependOn is used for other resources to depend on [Ec2AvailabilityZoneGroup].
func (eazg *Ec2AvailabilityZoneGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(eazg)
}

// Dependencies returns the list of resources [Ec2AvailabilityZoneGroup] depends_on.
func (eazg *Ec2AvailabilityZoneGroup) Dependencies() terra.Dependencies {
	return eazg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2AvailabilityZoneGroup].
func (eazg *Ec2AvailabilityZoneGroup) LifecycleManagement() *terra.Lifecycle {
	return eazg.Lifecycle
}

// Attributes returns the attributes for [Ec2AvailabilityZoneGroup].
func (eazg *Ec2AvailabilityZoneGroup) Attributes() ec2AvailabilityZoneGroupAttributes {
	return ec2AvailabilityZoneGroupAttributes{ref: terra.ReferenceResource(eazg)}
}

// ImportState imports the given attribute values into [Ec2AvailabilityZoneGroup]'s state.
func (eazg *Ec2AvailabilityZoneGroup) ImportState(av io.Reader) error {
	eazg.state = &ec2AvailabilityZoneGroupState{}
	if err := json.NewDecoder(av).Decode(eazg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", eazg.Type(), eazg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2AvailabilityZoneGroup] has state.
func (eazg *Ec2AvailabilityZoneGroup) State() (*ec2AvailabilityZoneGroupState, bool) {
	return eazg.state, eazg.state != nil
}

// StateMust returns the state for [Ec2AvailabilityZoneGroup]. Panics if the state is nil.
func (eazg *Ec2AvailabilityZoneGroup) StateMust() *ec2AvailabilityZoneGroupState {
	if eazg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", eazg.Type(), eazg.LocalName()))
	}
	return eazg.state
}

// Ec2AvailabilityZoneGroupArgs contains the configurations for aws_ec2_availability_zone_group.
type Ec2AvailabilityZoneGroupArgs struct {
	// GroupName: string, required
	GroupName terra.StringValue `hcl:"group_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OptInStatus: string, required
	OptInStatus terra.StringValue `hcl:"opt_in_status,attr" validate:"required"`
}
type ec2AvailabilityZoneGroupAttributes struct {
	ref terra.Reference
}

// GroupName returns a reference to field group_name of aws_ec2_availability_zone_group.
func (eazg ec2AvailabilityZoneGroupAttributes) GroupName() terra.StringValue {
	return terra.ReferenceAsString(eazg.ref.Append("group_name"))
}

// Id returns a reference to field id of aws_ec2_availability_zone_group.
func (eazg ec2AvailabilityZoneGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eazg.ref.Append("id"))
}

// OptInStatus returns a reference to field opt_in_status of aws_ec2_availability_zone_group.
func (eazg ec2AvailabilityZoneGroupAttributes) OptInStatus() terra.StringValue {
	return terra.ReferenceAsString(eazg.ref.Append("opt_in_status"))
}

type ec2AvailabilityZoneGroupState struct {
	GroupName   string `json:"group_name"`
	Id          string `json:"id"`
	OptInStatus string `json:"opt_in_status"`
}
