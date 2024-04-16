// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_cloudhsm_v2_hsm

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

// Resource represents the Terraform resource aws_cloudhsm_v2_hsm.
type Resource struct {
	Name      string
	Args      Args
	state     *awsCloudhsmV2HsmState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acvh *Resource) Type() string {
	return "aws_cloudhsm_v2_hsm"
}

// LocalName returns the local name for [Resource].
func (acvh *Resource) LocalName() string {
	return acvh.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acvh *Resource) Configuration() interface{} {
	return acvh.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acvh *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acvh)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acvh *Resource) Dependencies() terra.Dependencies {
	return acvh.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acvh *Resource) LifecycleManagement() *terra.Lifecycle {
	return acvh.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acvh *Resource) Attributes() awsCloudhsmV2HsmAttributes {
	return awsCloudhsmV2HsmAttributes{ref: terra.ReferenceResource(acvh)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acvh *Resource) ImportState(state io.Reader) error {
	acvh.state = &awsCloudhsmV2HsmState{}
	if err := json.NewDecoder(state).Decode(acvh.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acvh.Type(), acvh.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acvh *Resource) State() (*awsCloudhsmV2HsmState, bool) {
	return acvh.state, acvh.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acvh *Resource) StateMust() *awsCloudhsmV2HsmState {
	if acvh.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acvh.Type(), acvh.LocalName()))
	}
	return acvh.state
}

// Args contains the configurations for aws_cloudhsm_v2_hsm.
type Args struct {
	// AvailabilityZone: string, optional
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr"`
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpAddress: string, optional
	IpAddress terra.StringValue `hcl:"ip_address,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsCloudhsmV2HsmAttributes struct {
	ref terra.Reference
}

// AvailabilityZone returns a reference to field availability_zone of aws_cloudhsm_v2_hsm.
func (acvh awsCloudhsmV2HsmAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(acvh.ref.Append("availability_zone"))
}

// ClusterId returns a reference to field cluster_id of aws_cloudhsm_v2_hsm.
func (acvh awsCloudhsmV2HsmAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(acvh.ref.Append("cluster_id"))
}

// HsmEniId returns a reference to field hsm_eni_id of aws_cloudhsm_v2_hsm.
func (acvh awsCloudhsmV2HsmAttributes) HsmEniId() terra.StringValue {
	return terra.ReferenceAsString(acvh.ref.Append("hsm_eni_id"))
}

// HsmId returns a reference to field hsm_id of aws_cloudhsm_v2_hsm.
func (acvh awsCloudhsmV2HsmAttributes) HsmId() terra.StringValue {
	return terra.ReferenceAsString(acvh.ref.Append("hsm_id"))
}

// HsmState returns a reference to field hsm_state of aws_cloudhsm_v2_hsm.
func (acvh awsCloudhsmV2HsmAttributes) HsmState() terra.StringValue {
	return terra.ReferenceAsString(acvh.ref.Append("hsm_state"))
}

// Id returns a reference to field id of aws_cloudhsm_v2_hsm.
func (acvh awsCloudhsmV2HsmAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acvh.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of aws_cloudhsm_v2_hsm.
func (acvh awsCloudhsmV2HsmAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(acvh.ref.Append("ip_address"))
}

// SubnetId returns a reference to field subnet_id of aws_cloudhsm_v2_hsm.
func (acvh awsCloudhsmV2HsmAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(acvh.ref.Append("subnet_id"))
}

func (acvh awsCloudhsmV2HsmAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](acvh.ref.Append("timeouts"))
}

type awsCloudhsmV2HsmState struct {
	AvailabilityZone string         `json:"availability_zone"`
	ClusterId        string         `json:"cluster_id"`
	HsmEniId         string         `json:"hsm_eni_id"`
	HsmId            string         `json:"hsm_id"`
	HsmState         string         `json:"hsm_state"`
	Id               string         `json:"id"`
	IpAddress        string         `json:"ip_address"`
	SubnetId         string         `json:"subnet_id"`
	Timeouts         *TimeoutsState `json:"timeouts"`
}
