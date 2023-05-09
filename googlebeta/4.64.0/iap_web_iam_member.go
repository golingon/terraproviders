// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	iapwebiammember "github.com/golingon/terraproviders/googlebeta/4.64.0/iapwebiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapWebIamMember creates a new instance of [IapWebIamMember].
func NewIapWebIamMember(name string, args IapWebIamMemberArgs) *IapWebIamMember {
	return &IapWebIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapWebIamMember)(nil)

// IapWebIamMember represents the Terraform resource google_iap_web_iam_member.
type IapWebIamMember struct {
	Name      string
	Args      IapWebIamMemberArgs
	state     *iapWebIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapWebIamMember].
func (iwim *IapWebIamMember) Type() string {
	return "google_iap_web_iam_member"
}

// LocalName returns the local name for [IapWebIamMember].
func (iwim *IapWebIamMember) LocalName() string {
	return iwim.Name
}

// Configuration returns the configuration (args) for [IapWebIamMember].
func (iwim *IapWebIamMember) Configuration() interface{} {
	return iwim.Args
}

// DependOn is used for other resources to depend on [IapWebIamMember].
func (iwim *IapWebIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(iwim)
}

// Dependencies returns the list of resources [IapWebIamMember] depends_on.
func (iwim *IapWebIamMember) Dependencies() terra.Dependencies {
	return iwim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapWebIamMember].
func (iwim *IapWebIamMember) LifecycleManagement() *terra.Lifecycle {
	return iwim.Lifecycle
}

// Attributes returns the attributes for [IapWebIamMember].
func (iwim *IapWebIamMember) Attributes() iapWebIamMemberAttributes {
	return iapWebIamMemberAttributes{ref: terra.ReferenceResource(iwim)}
}

// ImportState imports the given attribute values into [IapWebIamMember]'s state.
func (iwim *IapWebIamMember) ImportState(av io.Reader) error {
	iwim.state = &iapWebIamMemberState{}
	if err := json.NewDecoder(av).Decode(iwim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iwim.Type(), iwim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapWebIamMember] has state.
func (iwim *IapWebIamMember) State() (*iapWebIamMemberState, bool) {
	return iwim.state, iwim.state != nil
}

// StateMust returns the state for [IapWebIamMember]. Panics if the state is nil.
func (iwim *IapWebIamMember) StateMust() *iapWebIamMemberState {
	if iwim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iwim.Type(), iwim.LocalName()))
	}
	return iwim.state
}

// IapWebIamMemberArgs contains the configurations for google_iap_web_iam_member.
type IapWebIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *iapwebiammember.Condition `hcl:"condition,block"`
}
type iapWebIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_web_iam_member.
func (iwim iapWebIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iwim.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_web_iam_member.
func (iwim iapWebIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwim.ref.Append("id"))
}

// Member returns a reference to field member of google_iap_web_iam_member.
func (iwim iapWebIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(iwim.ref.Append("member"))
}

// Project returns a reference to field project of google_iap_web_iam_member.
func (iwim iapWebIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwim.ref.Append("project"))
}

// Role returns a reference to field role of google_iap_web_iam_member.
func (iwim iapWebIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(iwim.ref.Append("role"))
}

func (iwim iapWebIamMemberAttributes) Condition() terra.ListValue[iapwebiammember.ConditionAttributes] {
	return terra.ReferenceAsList[iapwebiammember.ConditionAttributes](iwim.ref.Append("condition"))
}

type iapWebIamMemberState struct {
	Etag      string                           `json:"etag"`
	Id        string                           `json:"id"`
	Member    string                           `json:"member"`
	Project   string                           `json:"project"`
	Role      string                           `json:"role"`
	Condition []iapwebiammember.ConditionState `json:"condition"`
}
