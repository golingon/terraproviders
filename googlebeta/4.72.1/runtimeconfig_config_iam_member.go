// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	runtimeconfigconfigiammember "github.com/golingon/terraproviders/googlebeta/4.72.1/runtimeconfigconfigiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRuntimeconfigConfigIamMember creates a new instance of [RuntimeconfigConfigIamMember].
func NewRuntimeconfigConfigIamMember(name string, args RuntimeconfigConfigIamMemberArgs) *RuntimeconfigConfigIamMember {
	return &RuntimeconfigConfigIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RuntimeconfigConfigIamMember)(nil)

// RuntimeconfigConfigIamMember represents the Terraform resource google_runtimeconfig_config_iam_member.
type RuntimeconfigConfigIamMember struct {
	Name      string
	Args      RuntimeconfigConfigIamMemberArgs
	state     *runtimeconfigConfigIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RuntimeconfigConfigIamMember].
func (rcim *RuntimeconfigConfigIamMember) Type() string {
	return "google_runtimeconfig_config_iam_member"
}

// LocalName returns the local name for [RuntimeconfigConfigIamMember].
func (rcim *RuntimeconfigConfigIamMember) LocalName() string {
	return rcim.Name
}

// Configuration returns the configuration (args) for [RuntimeconfigConfigIamMember].
func (rcim *RuntimeconfigConfigIamMember) Configuration() interface{} {
	return rcim.Args
}

// DependOn is used for other resources to depend on [RuntimeconfigConfigIamMember].
func (rcim *RuntimeconfigConfigIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(rcim)
}

// Dependencies returns the list of resources [RuntimeconfigConfigIamMember] depends_on.
func (rcim *RuntimeconfigConfigIamMember) Dependencies() terra.Dependencies {
	return rcim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RuntimeconfigConfigIamMember].
func (rcim *RuntimeconfigConfigIamMember) LifecycleManagement() *terra.Lifecycle {
	return rcim.Lifecycle
}

// Attributes returns the attributes for [RuntimeconfigConfigIamMember].
func (rcim *RuntimeconfigConfigIamMember) Attributes() runtimeconfigConfigIamMemberAttributes {
	return runtimeconfigConfigIamMemberAttributes{ref: terra.ReferenceResource(rcim)}
}

// ImportState imports the given attribute values into [RuntimeconfigConfigIamMember]'s state.
func (rcim *RuntimeconfigConfigIamMember) ImportState(av io.Reader) error {
	rcim.state = &runtimeconfigConfigIamMemberState{}
	if err := json.NewDecoder(av).Decode(rcim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rcim.Type(), rcim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RuntimeconfigConfigIamMember] has state.
func (rcim *RuntimeconfigConfigIamMember) State() (*runtimeconfigConfigIamMemberState, bool) {
	return rcim.state, rcim.state != nil
}

// StateMust returns the state for [RuntimeconfigConfigIamMember]. Panics if the state is nil.
func (rcim *RuntimeconfigConfigIamMember) StateMust() *runtimeconfigConfigIamMemberState {
	if rcim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rcim.Type(), rcim.LocalName()))
	}
	return rcim.state
}

// RuntimeconfigConfigIamMemberArgs contains the configurations for google_runtimeconfig_config_iam_member.
type RuntimeconfigConfigIamMemberArgs struct {
	// Config: string, required
	Config terra.StringValue `hcl:"config,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *runtimeconfigconfigiammember.Condition `hcl:"condition,block"`
}
type runtimeconfigConfigIamMemberAttributes struct {
	ref terra.Reference
}

// Config returns a reference to field config of google_runtimeconfig_config_iam_member.
func (rcim runtimeconfigConfigIamMemberAttributes) Config() terra.StringValue {
	return terra.ReferenceAsString(rcim.ref.Append("config"))
}

// Etag returns a reference to field etag of google_runtimeconfig_config_iam_member.
func (rcim runtimeconfigConfigIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(rcim.ref.Append("etag"))
}

// Id returns a reference to field id of google_runtimeconfig_config_iam_member.
func (rcim runtimeconfigConfigIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rcim.ref.Append("id"))
}

// Member returns a reference to field member of google_runtimeconfig_config_iam_member.
func (rcim runtimeconfigConfigIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(rcim.ref.Append("member"))
}

// Project returns a reference to field project of google_runtimeconfig_config_iam_member.
func (rcim runtimeconfigConfigIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(rcim.ref.Append("project"))
}

// Role returns a reference to field role of google_runtimeconfig_config_iam_member.
func (rcim runtimeconfigConfigIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(rcim.ref.Append("role"))
}

func (rcim runtimeconfigConfigIamMemberAttributes) Condition() terra.ListValue[runtimeconfigconfigiammember.ConditionAttributes] {
	return terra.ReferenceAsList[runtimeconfigconfigiammember.ConditionAttributes](rcim.ref.Append("condition"))
}

type runtimeconfigConfigIamMemberState struct {
	Config    string                                        `json:"config"`
	Etag      string                                        `json:"etag"`
	Id        string                                        `json:"id"`
	Member    string                                        `json:"member"`
	Project   string                                        `json:"project"`
	Role      string                                        `json:"role"`
	Condition []runtimeconfigconfigiammember.ConditionState `json:"condition"`
}
