// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	iapwebbackendserviceiammember "github.com/golingon/terraproviders/googlebeta/4.71.0/iapwebbackendserviceiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapWebBackendServiceIamMember creates a new instance of [IapWebBackendServiceIamMember].
func NewIapWebBackendServiceIamMember(name string, args IapWebBackendServiceIamMemberArgs) *IapWebBackendServiceIamMember {
	return &IapWebBackendServiceIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapWebBackendServiceIamMember)(nil)

// IapWebBackendServiceIamMember represents the Terraform resource google_iap_web_backend_service_iam_member.
type IapWebBackendServiceIamMember struct {
	Name      string
	Args      IapWebBackendServiceIamMemberArgs
	state     *iapWebBackendServiceIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapWebBackendServiceIamMember].
func (iwbsim *IapWebBackendServiceIamMember) Type() string {
	return "google_iap_web_backend_service_iam_member"
}

// LocalName returns the local name for [IapWebBackendServiceIamMember].
func (iwbsim *IapWebBackendServiceIamMember) LocalName() string {
	return iwbsim.Name
}

// Configuration returns the configuration (args) for [IapWebBackendServiceIamMember].
func (iwbsim *IapWebBackendServiceIamMember) Configuration() interface{} {
	return iwbsim.Args
}

// DependOn is used for other resources to depend on [IapWebBackendServiceIamMember].
func (iwbsim *IapWebBackendServiceIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(iwbsim)
}

// Dependencies returns the list of resources [IapWebBackendServiceIamMember] depends_on.
func (iwbsim *IapWebBackendServiceIamMember) Dependencies() terra.Dependencies {
	return iwbsim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapWebBackendServiceIamMember].
func (iwbsim *IapWebBackendServiceIamMember) LifecycleManagement() *terra.Lifecycle {
	return iwbsim.Lifecycle
}

// Attributes returns the attributes for [IapWebBackendServiceIamMember].
func (iwbsim *IapWebBackendServiceIamMember) Attributes() iapWebBackendServiceIamMemberAttributes {
	return iapWebBackendServiceIamMemberAttributes{ref: terra.ReferenceResource(iwbsim)}
}

// ImportState imports the given attribute values into [IapWebBackendServiceIamMember]'s state.
func (iwbsim *IapWebBackendServiceIamMember) ImportState(av io.Reader) error {
	iwbsim.state = &iapWebBackendServiceIamMemberState{}
	if err := json.NewDecoder(av).Decode(iwbsim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iwbsim.Type(), iwbsim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapWebBackendServiceIamMember] has state.
func (iwbsim *IapWebBackendServiceIamMember) State() (*iapWebBackendServiceIamMemberState, bool) {
	return iwbsim.state, iwbsim.state != nil
}

// StateMust returns the state for [IapWebBackendServiceIamMember]. Panics if the state is nil.
func (iwbsim *IapWebBackendServiceIamMember) StateMust() *iapWebBackendServiceIamMemberState {
	if iwbsim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iwbsim.Type(), iwbsim.LocalName()))
	}
	return iwbsim.state
}

// IapWebBackendServiceIamMemberArgs contains the configurations for google_iap_web_backend_service_iam_member.
type IapWebBackendServiceIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// WebBackendService: string, required
	WebBackendService terra.StringValue `hcl:"web_backend_service,attr" validate:"required"`
	// Condition: optional
	Condition *iapwebbackendserviceiammember.Condition `hcl:"condition,block"`
}
type iapWebBackendServiceIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_web_backend_service_iam_member.
func (iwbsim iapWebBackendServiceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iwbsim.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_web_backend_service_iam_member.
func (iwbsim iapWebBackendServiceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwbsim.ref.Append("id"))
}

// Member returns a reference to field member of google_iap_web_backend_service_iam_member.
func (iwbsim iapWebBackendServiceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(iwbsim.ref.Append("member"))
}

// Project returns a reference to field project of google_iap_web_backend_service_iam_member.
func (iwbsim iapWebBackendServiceIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwbsim.ref.Append("project"))
}

// Role returns a reference to field role of google_iap_web_backend_service_iam_member.
func (iwbsim iapWebBackendServiceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(iwbsim.ref.Append("role"))
}

// WebBackendService returns a reference to field web_backend_service of google_iap_web_backend_service_iam_member.
func (iwbsim iapWebBackendServiceIamMemberAttributes) WebBackendService() terra.StringValue {
	return terra.ReferenceAsString(iwbsim.ref.Append("web_backend_service"))
}

func (iwbsim iapWebBackendServiceIamMemberAttributes) Condition() terra.ListValue[iapwebbackendserviceiammember.ConditionAttributes] {
	return terra.ReferenceAsList[iapwebbackendserviceiammember.ConditionAttributes](iwbsim.ref.Append("condition"))
}

type iapWebBackendServiceIamMemberState struct {
	Etag              string                                         `json:"etag"`
	Id                string                                         `json:"id"`
	Member            string                                         `json:"member"`
	Project           string                                         `json:"project"`
	Role              string                                         `json:"role"`
	WebBackendService string                                         `json:"web_backend_service"`
	Condition         []iapwebbackendserviceiammember.ConditionState `json:"condition"`
}
