// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_batch_scheduling_policy

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

// Resource represents the Terraform resource aws_batch_scheduling_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *awsBatchSchedulingPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (absp *Resource) Type() string {
	return "aws_batch_scheduling_policy"
}

// LocalName returns the local name for [Resource].
func (absp *Resource) LocalName() string {
	return absp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (absp *Resource) Configuration() interface{} {
	return absp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (absp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(absp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (absp *Resource) Dependencies() terra.Dependencies {
	return absp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (absp *Resource) LifecycleManagement() *terra.Lifecycle {
	return absp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (absp *Resource) Attributes() awsBatchSchedulingPolicyAttributes {
	return awsBatchSchedulingPolicyAttributes{ref: terra.ReferenceResource(absp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (absp *Resource) ImportState(state io.Reader) error {
	absp.state = &awsBatchSchedulingPolicyState{}
	if err := json.NewDecoder(state).Decode(absp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", absp.Type(), absp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (absp *Resource) State() (*awsBatchSchedulingPolicyState, bool) {
	return absp.state, absp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (absp *Resource) StateMust() *awsBatchSchedulingPolicyState {
	if absp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", absp.Type(), absp.LocalName()))
	}
	return absp.state
}

// Args contains the configurations for aws_batch_scheduling_policy.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// FairSharePolicy: optional
	FairSharePolicy *FairSharePolicy `hcl:"fair_share_policy,block"`
}

type awsBatchSchedulingPolicyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_batch_scheduling_policy.
func (absp awsBatchSchedulingPolicyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(absp.ref.Append("arn"))
}

// Id returns a reference to field id of aws_batch_scheduling_policy.
func (absp awsBatchSchedulingPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(absp.ref.Append("id"))
}

// Name returns a reference to field name of aws_batch_scheduling_policy.
func (absp awsBatchSchedulingPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(absp.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_batch_scheduling_policy.
func (absp awsBatchSchedulingPolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](absp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_batch_scheduling_policy.
func (absp awsBatchSchedulingPolicyAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](absp.ref.Append("tags_all"))
}

func (absp awsBatchSchedulingPolicyAttributes) FairSharePolicy() terra.ListValue[FairSharePolicyAttributes] {
	return terra.ReferenceAsList[FairSharePolicyAttributes](absp.ref.Append("fair_share_policy"))
}

type awsBatchSchedulingPolicyState struct {
	Arn             string                 `json:"arn"`
	Id              string                 `json:"id"`
	Name            string                 `json:"name"`
	Tags            map[string]string      `json:"tags"`
	TagsAll         map[string]string      `json:"tags_all"`
	FairSharePolicy []FairSharePolicyState `json:"fair_share_policy"`
}
