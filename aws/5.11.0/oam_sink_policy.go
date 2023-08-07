// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	oamsinkpolicy "github.com/golingon/terraproviders/aws/5.11.0/oamsinkpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOamSinkPolicy creates a new instance of [OamSinkPolicy].
func NewOamSinkPolicy(name string, args OamSinkPolicyArgs) *OamSinkPolicy {
	return &OamSinkPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OamSinkPolicy)(nil)

// OamSinkPolicy represents the Terraform resource aws_oam_sink_policy.
type OamSinkPolicy struct {
	Name      string
	Args      OamSinkPolicyArgs
	state     *oamSinkPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OamSinkPolicy].
func (osp *OamSinkPolicy) Type() string {
	return "aws_oam_sink_policy"
}

// LocalName returns the local name for [OamSinkPolicy].
func (osp *OamSinkPolicy) LocalName() string {
	return osp.Name
}

// Configuration returns the configuration (args) for [OamSinkPolicy].
func (osp *OamSinkPolicy) Configuration() interface{} {
	return osp.Args
}

// DependOn is used for other resources to depend on [OamSinkPolicy].
func (osp *OamSinkPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(osp)
}

// Dependencies returns the list of resources [OamSinkPolicy] depends_on.
func (osp *OamSinkPolicy) Dependencies() terra.Dependencies {
	return osp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OamSinkPolicy].
func (osp *OamSinkPolicy) LifecycleManagement() *terra.Lifecycle {
	return osp.Lifecycle
}

// Attributes returns the attributes for [OamSinkPolicy].
func (osp *OamSinkPolicy) Attributes() oamSinkPolicyAttributes {
	return oamSinkPolicyAttributes{ref: terra.ReferenceResource(osp)}
}

// ImportState imports the given attribute values into [OamSinkPolicy]'s state.
func (osp *OamSinkPolicy) ImportState(av io.Reader) error {
	osp.state = &oamSinkPolicyState{}
	if err := json.NewDecoder(av).Decode(osp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", osp.Type(), osp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OamSinkPolicy] has state.
func (osp *OamSinkPolicy) State() (*oamSinkPolicyState, bool) {
	return osp.state, osp.state != nil
}

// StateMust returns the state for [OamSinkPolicy]. Panics if the state is nil.
func (osp *OamSinkPolicy) StateMust() *oamSinkPolicyState {
	if osp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", osp.Type(), osp.LocalName()))
	}
	return osp.state
}

// OamSinkPolicyArgs contains the configurations for aws_oam_sink_policy.
type OamSinkPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
	// SinkIdentifier: string, required
	SinkIdentifier terra.StringValue `hcl:"sink_identifier,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *oamsinkpolicy.Timeouts `hcl:"timeouts,block"`
}
type oamSinkPolicyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_oam_sink_policy.
func (osp oamSinkPolicyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(osp.ref.Append("arn"))
}

// Id returns a reference to field id of aws_oam_sink_policy.
func (osp oamSinkPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(osp.ref.Append("id"))
}

// Policy returns a reference to field policy of aws_oam_sink_policy.
func (osp oamSinkPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(osp.ref.Append("policy"))
}

// SinkId returns a reference to field sink_id of aws_oam_sink_policy.
func (osp oamSinkPolicyAttributes) SinkId() terra.StringValue {
	return terra.ReferenceAsString(osp.ref.Append("sink_id"))
}

// SinkIdentifier returns a reference to field sink_identifier of aws_oam_sink_policy.
func (osp oamSinkPolicyAttributes) SinkIdentifier() terra.StringValue {
	return terra.ReferenceAsString(osp.ref.Append("sink_identifier"))
}

func (osp oamSinkPolicyAttributes) Timeouts() oamsinkpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[oamsinkpolicy.TimeoutsAttributes](osp.ref.Append("timeouts"))
}

type oamSinkPolicyState struct {
	Arn            string                       `json:"arn"`
	Id             string                       `json:"id"`
	Policy         string                       `json:"policy"`
	SinkId         string                       `json:"sink_id"`
	SinkIdentifier string                       `json:"sink_identifier"`
	Timeouts       *oamsinkpolicy.TimeoutsState `json:"timeouts"`
}
