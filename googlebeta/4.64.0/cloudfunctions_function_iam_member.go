// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	cloudfunctionsfunctioniammember "github.com/golingon/terraproviders/googlebeta/4.64.0/cloudfunctionsfunctioniammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudfunctionsFunctionIamMember creates a new instance of [CloudfunctionsFunctionIamMember].
func NewCloudfunctionsFunctionIamMember(name string, args CloudfunctionsFunctionIamMemberArgs) *CloudfunctionsFunctionIamMember {
	return &CloudfunctionsFunctionIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudfunctionsFunctionIamMember)(nil)

// CloudfunctionsFunctionIamMember represents the Terraform resource google_cloudfunctions_function_iam_member.
type CloudfunctionsFunctionIamMember struct {
	Name      string
	Args      CloudfunctionsFunctionIamMemberArgs
	state     *cloudfunctionsFunctionIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudfunctionsFunctionIamMember].
func (cfim *CloudfunctionsFunctionIamMember) Type() string {
	return "google_cloudfunctions_function_iam_member"
}

// LocalName returns the local name for [CloudfunctionsFunctionIamMember].
func (cfim *CloudfunctionsFunctionIamMember) LocalName() string {
	return cfim.Name
}

// Configuration returns the configuration (args) for [CloudfunctionsFunctionIamMember].
func (cfim *CloudfunctionsFunctionIamMember) Configuration() interface{} {
	return cfim.Args
}

// DependOn is used for other resources to depend on [CloudfunctionsFunctionIamMember].
func (cfim *CloudfunctionsFunctionIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(cfim)
}

// Dependencies returns the list of resources [CloudfunctionsFunctionIamMember] depends_on.
func (cfim *CloudfunctionsFunctionIamMember) Dependencies() terra.Dependencies {
	return cfim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudfunctionsFunctionIamMember].
func (cfim *CloudfunctionsFunctionIamMember) LifecycleManagement() *terra.Lifecycle {
	return cfim.Lifecycle
}

// Attributes returns the attributes for [CloudfunctionsFunctionIamMember].
func (cfim *CloudfunctionsFunctionIamMember) Attributes() cloudfunctionsFunctionIamMemberAttributes {
	return cloudfunctionsFunctionIamMemberAttributes{ref: terra.ReferenceResource(cfim)}
}

// ImportState imports the given attribute values into [CloudfunctionsFunctionIamMember]'s state.
func (cfim *CloudfunctionsFunctionIamMember) ImportState(av io.Reader) error {
	cfim.state = &cloudfunctionsFunctionIamMemberState{}
	if err := json.NewDecoder(av).Decode(cfim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfim.Type(), cfim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudfunctionsFunctionIamMember] has state.
func (cfim *CloudfunctionsFunctionIamMember) State() (*cloudfunctionsFunctionIamMemberState, bool) {
	return cfim.state, cfim.state != nil
}

// StateMust returns the state for [CloudfunctionsFunctionIamMember]. Panics if the state is nil.
func (cfim *CloudfunctionsFunctionIamMember) StateMust() *cloudfunctionsFunctionIamMemberState {
	if cfim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfim.Type(), cfim.LocalName()))
	}
	return cfim.state
}

// CloudfunctionsFunctionIamMemberArgs contains the configurations for google_cloudfunctions_function_iam_member.
type CloudfunctionsFunctionIamMemberArgs struct {
	// CloudFunction: string, required
	CloudFunction terra.StringValue `hcl:"cloud_function,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *cloudfunctionsfunctioniammember.Condition `hcl:"condition,block"`
}
type cloudfunctionsFunctionIamMemberAttributes struct {
	ref terra.Reference
}

// CloudFunction returns a reference to field cloud_function of google_cloudfunctions_function_iam_member.
func (cfim cloudfunctionsFunctionIamMemberAttributes) CloudFunction() terra.StringValue {
	return terra.ReferenceAsString(cfim.ref.Append("cloud_function"))
}

// Etag returns a reference to field etag of google_cloudfunctions_function_iam_member.
func (cfim cloudfunctionsFunctionIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cfim.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloudfunctions_function_iam_member.
func (cfim cloudfunctionsFunctionIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfim.ref.Append("id"))
}

// Member returns a reference to field member of google_cloudfunctions_function_iam_member.
func (cfim cloudfunctionsFunctionIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(cfim.ref.Append("member"))
}

// Project returns a reference to field project of google_cloudfunctions_function_iam_member.
func (cfim cloudfunctionsFunctionIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cfim.ref.Append("project"))
}

// Region returns a reference to field region of google_cloudfunctions_function_iam_member.
func (cfim cloudfunctionsFunctionIamMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cfim.ref.Append("region"))
}

// Role returns a reference to field role of google_cloudfunctions_function_iam_member.
func (cfim cloudfunctionsFunctionIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(cfim.ref.Append("role"))
}

func (cfim cloudfunctionsFunctionIamMemberAttributes) Condition() terra.ListValue[cloudfunctionsfunctioniammember.ConditionAttributes] {
	return terra.ReferenceAsList[cloudfunctionsfunctioniammember.ConditionAttributes](cfim.ref.Append("condition"))
}

type cloudfunctionsFunctionIamMemberState struct {
	CloudFunction string                                           `json:"cloud_function"`
	Etag          string                                           `json:"etag"`
	Id            string                                           `json:"id"`
	Member        string                                           `json:"member"`
	Project       string                                           `json:"project"`
	Region        string                                           `json:"region"`
	Role          string                                           `json:"role"`
	Condition     []cloudfunctionsfunctioniammember.ConditionState `json:"condition"`
}
