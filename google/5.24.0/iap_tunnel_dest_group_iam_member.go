// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	iaptunneldestgroupiammember "github.com/golingon/terraproviders/google/5.24.0/iaptunneldestgroupiammember"
	"io"
)

// NewIapTunnelDestGroupIamMember creates a new instance of [IapTunnelDestGroupIamMember].
func NewIapTunnelDestGroupIamMember(name string, args IapTunnelDestGroupIamMemberArgs) *IapTunnelDestGroupIamMember {
	return &IapTunnelDestGroupIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IapTunnelDestGroupIamMember)(nil)

// IapTunnelDestGroupIamMember represents the Terraform resource google_iap_tunnel_dest_group_iam_member.
type IapTunnelDestGroupIamMember struct {
	Name      string
	Args      IapTunnelDestGroupIamMemberArgs
	state     *iapTunnelDestGroupIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IapTunnelDestGroupIamMember].
func (itdgim *IapTunnelDestGroupIamMember) Type() string {
	return "google_iap_tunnel_dest_group_iam_member"
}

// LocalName returns the local name for [IapTunnelDestGroupIamMember].
func (itdgim *IapTunnelDestGroupIamMember) LocalName() string {
	return itdgim.Name
}

// Configuration returns the configuration (args) for [IapTunnelDestGroupIamMember].
func (itdgim *IapTunnelDestGroupIamMember) Configuration() interface{} {
	return itdgim.Args
}

// DependOn is used for other resources to depend on [IapTunnelDestGroupIamMember].
func (itdgim *IapTunnelDestGroupIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(itdgim)
}

// Dependencies returns the list of resources [IapTunnelDestGroupIamMember] depends_on.
func (itdgim *IapTunnelDestGroupIamMember) Dependencies() terra.Dependencies {
	return itdgim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IapTunnelDestGroupIamMember].
func (itdgim *IapTunnelDestGroupIamMember) LifecycleManagement() *terra.Lifecycle {
	return itdgim.Lifecycle
}

// Attributes returns the attributes for [IapTunnelDestGroupIamMember].
func (itdgim *IapTunnelDestGroupIamMember) Attributes() iapTunnelDestGroupIamMemberAttributes {
	return iapTunnelDestGroupIamMemberAttributes{ref: terra.ReferenceResource(itdgim)}
}

// ImportState imports the given attribute values into [IapTunnelDestGroupIamMember]'s state.
func (itdgim *IapTunnelDestGroupIamMember) ImportState(av io.Reader) error {
	itdgim.state = &iapTunnelDestGroupIamMemberState{}
	if err := json.NewDecoder(av).Decode(itdgim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", itdgim.Type(), itdgim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IapTunnelDestGroupIamMember] has state.
func (itdgim *IapTunnelDestGroupIamMember) State() (*iapTunnelDestGroupIamMemberState, bool) {
	return itdgim.state, itdgim.state != nil
}

// StateMust returns the state for [IapTunnelDestGroupIamMember]. Panics if the state is nil.
func (itdgim *IapTunnelDestGroupIamMember) StateMust() *iapTunnelDestGroupIamMemberState {
	if itdgim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", itdgim.Type(), itdgim.LocalName()))
	}
	return itdgim.state
}

// IapTunnelDestGroupIamMemberArgs contains the configurations for google_iap_tunnel_dest_group_iam_member.
type IapTunnelDestGroupIamMemberArgs struct {
	// DestGroup: string, required
	DestGroup terra.StringValue `hcl:"dest_group,attr" validate:"required"`
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
	Condition *iaptunneldestgroupiammember.Condition `hcl:"condition,block"`
}
type iapTunnelDestGroupIamMemberAttributes struct {
	ref terra.Reference
}

// DestGroup returns a reference to field dest_group of google_iap_tunnel_dest_group_iam_member.
func (itdgim iapTunnelDestGroupIamMemberAttributes) DestGroup() terra.StringValue {
	return terra.ReferenceAsString(itdgim.ref.Append("dest_group"))
}

// Etag returns a reference to field etag of google_iap_tunnel_dest_group_iam_member.
func (itdgim iapTunnelDestGroupIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(itdgim.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_tunnel_dest_group_iam_member.
func (itdgim iapTunnelDestGroupIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(itdgim.ref.Append("id"))
}

// Member returns a reference to field member of google_iap_tunnel_dest_group_iam_member.
func (itdgim iapTunnelDestGroupIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(itdgim.ref.Append("member"))
}

// Project returns a reference to field project of google_iap_tunnel_dest_group_iam_member.
func (itdgim iapTunnelDestGroupIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(itdgim.ref.Append("project"))
}

// Region returns a reference to field region of google_iap_tunnel_dest_group_iam_member.
func (itdgim iapTunnelDestGroupIamMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(itdgim.ref.Append("region"))
}

// Role returns a reference to field role of google_iap_tunnel_dest_group_iam_member.
func (itdgim iapTunnelDestGroupIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(itdgim.ref.Append("role"))
}

func (itdgim iapTunnelDestGroupIamMemberAttributes) Condition() terra.ListValue[iaptunneldestgroupiammember.ConditionAttributes] {
	return terra.ReferenceAsList[iaptunneldestgroupiammember.ConditionAttributes](itdgim.ref.Append("condition"))
}

type iapTunnelDestGroupIamMemberState struct {
	DestGroup string                                       `json:"dest_group"`
	Etag      string                                       `json:"etag"`
	Id        string                                       `json:"id"`
	Member    string                                       `json:"member"`
	Project   string                                       `json:"project"`
	Region    string                                       `json:"region"`
	Role      string                                       `json:"role"`
	Condition []iaptunneldestgroupiammember.ConditionState `json:"condition"`
}
