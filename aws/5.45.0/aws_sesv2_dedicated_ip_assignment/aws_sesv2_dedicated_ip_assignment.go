// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_sesv2_dedicated_ip_assignment

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_sesv2_dedicated_ip_assignment.
type Resource struct {
	Name      string
	Args      Args
	state     *awsSesv2DedicatedIpAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asdia *Resource) Type() string {
	return "aws_sesv2_dedicated_ip_assignment"
}

// LocalName returns the local name for [Resource].
func (asdia *Resource) LocalName() string {
	return asdia.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asdia *Resource) Configuration() interface{} {
	return asdia.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asdia *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asdia)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asdia *Resource) Dependencies() terra.Dependencies {
	return asdia.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asdia *Resource) LifecycleManagement() *terra.Lifecycle {
	return asdia.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asdia *Resource) Attributes() awsSesv2DedicatedIpAssignmentAttributes {
	return awsSesv2DedicatedIpAssignmentAttributes{ref: terra.ReferenceResource(asdia)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asdia *Resource) ImportState(state io.Reader) error {
	asdia.state = &awsSesv2DedicatedIpAssignmentState{}
	if err := json.NewDecoder(state).Decode(asdia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asdia.Type(), asdia.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asdia *Resource) State() (*awsSesv2DedicatedIpAssignmentState, bool) {
	return asdia.state, asdia.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asdia *Resource) StateMust() *awsSesv2DedicatedIpAssignmentState {
	if asdia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asdia.Type(), asdia.LocalName()))
	}
	return asdia.state
}

// Args contains the configurations for aws_sesv2_dedicated_ip_assignment.
type Args struct {
	// DestinationPoolName: string, required
	DestinationPoolName terra.StringValue `hcl:"destination_pool_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Ip: string, required
	Ip terra.StringValue `hcl:"ip,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsSesv2DedicatedIpAssignmentAttributes struct {
	ref terra.Reference
}

// DestinationPoolName returns a reference to field destination_pool_name of aws_sesv2_dedicated_ip_assignment.
func (asdia awsSesv2DedicatedIpAssignmentAttributes) DestinationPoolName() terra.StringValue {
	return terra.ReferenceAsString(asdia.ref.Append("destination_pool_name"))
}

// Id returns a reference to field id of aws_sesv2_dedicated_ip_assignment.
func (asdia awsSesv2DedicatedIpAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asdia.ref.Append("id"))
}

// Ip returns a reference to field ip of aws_sesv2_dedicated_ip_assignment.
func (asdia awsSesv2DedicatedIpAssignmentAttributes) Ip() terra.StringValue {
	return terra.ReferenceAsString(asdia.ref.Append("ip"))
}

func (asdia awsSesv2DedicatedIpAssignmentAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](asdia.ref.Append("timeouts"))
}

type awsSesv2DedicatedIpAssignmentState struct {
	DestinationPoolName string         `json:"destination_pool_name"`
	Id                  string         `json:"id"`
	Ip                  string         `json:"ip"`
	Timeouts            *TimeoutsState `json:"timeouts"`
}
