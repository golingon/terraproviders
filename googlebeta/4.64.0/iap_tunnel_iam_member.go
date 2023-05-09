// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	iaptunneliammember "github.com/golingon/terraproviders/googlebeta/4.64.0/iaptunneliammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIapTunnelIamMember creates a new instance of [IapTunnelIamMember].
func NewIapTunnelIamMember(name string, args IapTunnelIamMemberArgs) *IapTunnelIamMember {
	return &IapTunnelIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapTunnelIamMember)(nil)

// IapTunnelIamMember represents the Terraform resource google_iap_tunnel_iam_member.
type IapTunnelIamMember struct {
	Name      string
	Args      IapTunnelIamMemberArgs
	state     *iapTunnelIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapTunnelIamMember].
func (itim *IapTunnelIamMember) Type() string {
	return "google_iap_tunnel_iam_member"
}

// LocalName returns the local name for [IapTunnelIamMember].
func (itim *IapTunnelIamMember) LocalName() string {
	return itim.Name
}

// Configuration returns the configuration (args) for [IapTunnelIamMember].
func (itim *IapTunnelIamMember) Configuration() interface{} {
	return itim.Args
}

// DependOn is used for other resources to depend on [IapTunnelIamMember].
func (itim *IapTunnelIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(itim)
}

// Dependencies returns the list of resources [IapTunnelIamMember] depends_on.
func (itim *IapTunnelIamMember) Dependencies() terra.Dependencies {
	return itim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapTunnelIamMember].
func (itim *IapTunnelIamMember) LifecycleManagement() *terra.Lifecycle {
	return itim.Lifecycle
}

// Attributes returns the attributes for [IapTunnelIamMember].
func (itim *IapTunnelIamMember) Attributes() iapTunnelIamMemberAttributes {
	return iapTunnelIamMemberAttributes{ref: terra.ReferenceResource(itim)}
}

// ImportState imports the given attribute values into [IapTunnelIamMember]'s state.
func (itim *IapTunnelIamMember) ImportState(av io.Reader) error {
	itim.state = &iapTunnelIamMemberState{}
	if err := json.NewDecoder(av).Decode(itim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", itim.Type(), itim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapTunnelIamMember] has state.
func (itim *IapTunnelIamMember) State() (*iapTunnelIamMemberState, bool) {
	return itim.state, itim.state != nil
}

// StateMust returns the state for [IapTunnelIamMember]. Panics if the state is nil.
func (itim *IapTunnelIamMember) StateMust() *iapTunnelIamMemberState {
	if itim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", itim.Type(), itim.LocalName()))
	}
	return itim.state
}

// IapTunnelIamMemberArgs contains the configurations for google_iap_tunnel_iam_member.
type IapTunnelIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *iaptunneliammember.Condition `hcl:"condition,block"`
}
type iapTunnelIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_tunnel_iam_member.
func (itim iapTunnelIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(itim.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_tunnel_iam_member.
func (itim iapTunnelIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(itim.ref.Append("id"))
}

// Member returns a reference to field member of google_iap_tunnel_iam_member.
func (itim iapTunnelIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(itim.ref.Append("member"))
}

// Project returns a reference to field project of google_iap_tunnel_iam_member.
func (itim iapTunnelIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(itim.ref.Append("project"))
}

// Role returns a reference to field role of google_iap_tunnel_iam_member.
func (itim iapTunnelIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(itim.ref.Append("role"))
}

func (itim iapTunnelIamMemberAttributes) Condition() terra.ListValue[iaptunneliammember.ConditionAttributes] {
	return terra.ReferenceAsList[iaptunneliammember.ConditionAttributes](itim.ref.Append("condition"))
}

type iapTunnelIamMemberState struct {
	Etag      string                              `json:"etag"`
	Id        string                              `json:"id"`
	Member    string                              `json:"member"`
	Project   string                              `json:"project"`
	Role      string                              `json:"role"`
	Condition []iaptunneliammember.ConditionState `json:"condition"`
}
