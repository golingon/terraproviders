// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudfunctions2functioniammember "github.com/golingon/terraproviders/google/4.62.0/cloudfunctions2functioniammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudfunctions2FunctionIamMember creates a new instance of [Cloudfunctions2FunctionIamMember].
func NewCloudfunctions2FunctionIamMember(name string, args Cloudfunctions2FunctionIamMemberArgs) *Cloudfunctions2FunctionIamMember {
	return &Cloudfunctions2FunctionIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Cloudfunctions2FunctionIamMember)(nil)

// Cloudfunctions2FunctionIamMember represents the Terraform resource google_cloudfunctions2_function_iam_member.
type Cloudfunctions2FunctionIamMember struct {
	Name      string
	Args      Cloudfunctions2FunctionIamMemberArgs
	state     *cloudfunctions2FunctionIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Cloudfunctions2FunctionIamMember].
func (cfim *Cloudfunctions2FunctionIamMember) Type() string {
	return "google_cloudfunctions2_function_iam_member"
}

// LocalName returns the local name for [Cloudfunctions2FunctionIamMember].
func (cfim *Cloudfunctions2FunctionIamMember) LocalName() string {
	return cfim.Name
}

// Configuration returns the configuration (args) for [Cloudfunctions2FunctionIamMember].
func (cfim *Cloudfunctions2FunctionIamMember) Configuration() interface{} {
	return cfim.Args
}

// DependOn is used for other resources to depend on [Cloudfunctions2FunctionIamMember].
func (cfim *Cloudfunctions2FunctionIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(cfim)
}

// Dependencies returns the list of resources [Cloudfunctions2FunctionIamMember] depends_on.
func (cfim *Cloudfunctions2FunctionIamMember) Dependencies() terra.Dependencies {
	return cfim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Cloudfunctions2FunctionIamMember].
func (cfim *Cloudfunctions2FunctionIamMember) LifecycleManagement() *terra.Lifecycle {
	return cfim.Lifecycle
}

// Attributes returns the attributes for [Cloudfunctions2FunctionIamMember].
func (cfim *Cloudfunctions2FunctionIamMember) Attributes() cloudfunctions2FunctionIamMemberAttributes {
	return cloudfunctions2FunctionIamMemberAttributes{ref: terra.ReferenceResource(cfim)}
}

// ImportState imports the given attribute values into [Cloudfunctions2FunctionIamMember]'s state.
func (cfim *Cloudfunctions2FunctionIamMember) ImportState(av io.Reader) error {
	cfim.state = &cloudfunctions2FunctionIamMemberState{}
	if err := json.NewDecoder(av).Decode(cfim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfim.Type(), cfim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Cloudfunctions2FunctionIamMember] has state.
func (cfim *Cloudfunctions2FunctionIamMember) State() (*cloudfunctions2FunctionIamMemberState, bool) {
	return cfim.state, cfim.state != nil
}

// StateMust returns the state for [Cloudfunctions2FunctionIamMember]. Panics if the state is nil.
func (cfim *Cloudfunctions2FunctionIamMember) StateMust() *cloudfunctions2FunctionIamMemberState {
	if cfim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfim.Type(), cfim.LocalName()))
	}
	return cfim.state
}

// Cloudfunctions2FunctionIamMemberArgs contains the configurations for google_cloudfunctions2_function_iam_member.
type Cloudfunctions2FunctionIamMemberArgs struct {
	// CloudFunction: string, required
	CloudFunction terra.StringValue `hcl:"cloud_function,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *cloudfunctions2functioniammember.Condition `hcl:"condition,block"`
}
type cloudfunctions2FunctionIamMemberAttributes struct {
	ref terra.Reference
}

// CloudFunction returns a reference to field cloud_function of google_cloudfunctions2_function_iam_member.
func (cfim cloudfunctions2FunctionIamMemberAttributes) CloudFunction() terra.StringValue {
	return terra.ReferenceAsString(cfim.ref.Append("cloud_function"))
}

// Etag returns a reference to field etag of google_cloudfunctions2_function_iam_member.
func (cfim cloudfunctions2FunctionIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cfim.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloudfunctions2_function_iam_member.
func (cfim cloudfunctions2FunctionIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfim.ref.Append("id"))
}

// Location returns a reference to field location of google_cloudfunctions2_function_iam_member.
func (cfim cloudfunctions2FunctionIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cfim.ref.Append("location"))
}

// Member returns a reference to field member of google_cloudfunctions2_function_iam_member.
func (cfim cloudfunctions2FunctionIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(cfim.ref.Append("member"))
}

// Project returns a reference to field project of google_cloudfunctions2_function_iam_member.
func (cfim cloudfunctions2FunctionIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cfim.ref.Append("project"))
}

// Role returns a reference to field role of google_cloudfunctions2_function_iam_member.
func (cfim cloudfunctions2FunctionIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(cfim.ref.Append("role"))
}

func (cfim cloudfunctions2FunctionIamMemberAttributes) Condition() terra.ListValue[cloudfunctions2functioniammember.ConditionAttributes] {
	return terra.ReferenceAsList[cloudfunctions2functioniammember.ConditionAttributes](cfim.ref.Append("condition"))
}

type cloudfunctions2FunctionIamMemberState struct {
	CloudFunction string                                            `json:"cloud_function"`
	Etag          string                                            `json:"etag"`
	Id            string                                            `json:"id"`
	Location      string                                            `json:"location"`
	Member        string                                            `json:"member"`
	Project       string                                            `json:"project"`
	Role          string                                            `json:"role"`
	Condition     []cloudfunctions2functioniammember.ConditionState `json:"condition"`
}
