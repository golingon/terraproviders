// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2Host creates a new instance of [Ec2Host].
func NewEc2Host(name string, args Ec2HostArgs) *Ec2Host {
	return &Ec2Host{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2Host)(nil)

// Ec2Host represents the Terraform resource aws_ec2_host.
type Ec2Host struct {
	Name      string
	Args      Ec2HostArgs
	state     *ec2HostState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2Host].
func (eh *Ec2Host) Type() string {
	return "aws_ec2_host"
}

// LocalName returns the local name for [Ec2Host].
func (eh *Ec2Host) LocalName() string {
	return eh.Name
}

// Configuration returns the configuration (args) for [Ec2Host].
func (eh *Ec2Host) Configuration() interface{} {
	return eh.Args
}

// DependOn is used for other resources to depend on [Ec2Host].
func (eh *Ec2Host) DependOn() terra.Reference {
	return terra.ReferenceResource(eh)
}

// Dependencies returns the list of resources [Ec2Host] depends_on.
func (eh *Ec2Host) Dependencies() terra.Dependencies {
	return eh.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2Host].
func (eh *Ec2Host) LifecycleManagement() *terra.Lifecycle {
	return eh.Lifecycle
}

// Attributes returns the attributes for [Ec2Host].
func (eh *Ec2Host) Attributes() ec2HostAttributes {
	return ec2HostAttributes{ref: terra.ReferenceResource(eh)}
}

// ImportState imports the given attribute values into [Ec2Host]'s state.
func (eh *Ec2Host) ImportState(av io.Reader) error {
	eh.state = &ec2HostState{}
	if err := json.NewDecoder(av).Decode(eh.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", eh.Type(), eh.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2Host] has state.
func (eh *Ec2Host) State() (*ec2HostState, bool) {
	return eh.state, eh.state != nil
}

// StateMust returns the state for [Ec2Host]. Panics if the state is nil.
func (eh *Ec2Host) StateMust() *ec2HostState {
	if eh.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", eh.Type(), eh.LocalName()))
	}
	return eh.state
}

// Ec2HostArgs contains the configurations for aws_ec2_host.
type Ec2HostArgs struct {
	// AutoPlacement: string, optional
	AutoPlacement terra.StringValue `hcl:"auto_placement,attr"`
	// AvailabilityZone: string, required
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr" validate:"required"`
	// HostRecovery: string, optional
	HostRecovery terra.StringValue `hcl:"host_recovery,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceFamily: string, optional
	InstanceFamily terra.StringValue `hcl:"instance_family,attr"`
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// OutpostArn: string, optional
	OutpostArn terra.StringValue `hcl:"outpost_arn,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type ec2HostAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_host.
func (eh ec2HostAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("arn"))
}

// AutoPlacement returns a reference to field auto_placement of aws_ec2_host.
func (eh ec2HostAttributes) AutoPlacement() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("auto_placement"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_ec2_host.
func (eh ec2HostAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("availability_zone"))
}

// HostRecovery returns a reference to field host_recovery of aws_ec2_host.
func (eh ec2HostAttributes) HostRecovery() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("host_recovery"))
}

// Id returns a reference to field id of aws_ec2_host.
func (eh ec2HostAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("id"))
}

// InstanceFamily returns a reference to field instance_family of aws_ec2_host.
func (eh ec2HostAttributes) InstanceFamily() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("instance_family"))
}

// InstanceType returns a reference to field instance_type of aws_ec2_host.
func (eh ec2HostAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("instance_type"))
}

// OutpostArn returns a reference to field outpost_arn of aws_ec2_host.
func (eh ec2HostAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("outpost_arn"))
}

// OwnerId returns a reference to field owner_id of aws_ec2_host.
func (eh ec2HostAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("owner_id"))
}

// Tags returns a reference to field tags of aws_ec2_host.
func (eh ec2HostAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](eh.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ec2_host.
func (eh ec2HostAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](eh.ref.Append("tags_all"))
}

type ec2HostState struct {
	Arn              string            `json:"arn"`
	AutoPlacement    string            `json:"auto_placement"`
	AvailabilityZone string            `json:"availability_zone"`
	HostRecovery     string            `json:"host_recovery"`
	Id               string            `json:"id"`
	InstanceFamily   string            `json:"instance_family"`
	InstanceType     string            `json:"instance_type"`
	OutpostArn       string            `json:"outpost_arn"`
	OwnerId          string            `json:"owner_id"`
	Tags             map[string]string `json:"tags"`
	TagsAll          map[string]string `json:"tags_all"`
}
