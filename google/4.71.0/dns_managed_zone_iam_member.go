// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dnsmanagedzoneiammember "github.com/golingon/terraproviders/google/4.71.0/dnsmanagedzoneiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDnsManagedZoneIamMember creates a new instance of [DnsManagedZoneIamMember].
func NewDnsManagedZoneIamMember(name string, args DnsManagedZoneIamMemberArgs) *DnsManagedZoneIamMember {
	return &DnsManagedZoneIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsManagedZoneIamMember)(nil)

// DnsManagedZoneIamMember represents the Terraform resource google_dns_managed_zone_iam_member.
type DnsManagedZoneIamMember struct {
	Name      string
	Args      DnsManagedZoneIamMemberArgs
	state     *dnsManagedZoneIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DnsManagedZoneIamMember].
func (dmzim *DnsManagedZoneIamMember) Type() string {
	return "google_dns_managed_zone_iam_member"
}

// LocalName returns the local name for [DnsManagedZoneIamMember].
func (dmzim *DnsManagedZoneIamMember) LocalName() string {
	return dmzim.Name
}

// Configuration returns the configuration (args) for [DnsManagedZoneIamMember].
func (dmzim *DnsManagedZoneIamMember) Configuration() interface{} {
	return dmzim.Args
}

// DependOn is used for other resources to depend on [DnsManagedZoneIamMember].
func (dmzim *DnsManagedZoneIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(dmzim)
}

// Dependencies returns the list of resources [DnsManagedZoneIamMember] depends_on.
func (dmzim *DnsManagedZoneIamMember) Dependencies() terra.Dependencies {
	return dmzim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DnsManagedZoneIamMember].
func (dmzim *DnsManagedZoneIamMember) LifecycleManagement() *terra.Lifecycle {
	return dmzim.Lifecycle
}

// Attributes returns the attributes for [DnsManagedZoneIamMember].
func (dmzim *DnsManagedZoneIamMember) Attributes() dnsManagedZoneIamMemberAttributes {
	return dnsManagedZoneIamMemberAttributes{ref: terra.ReferenceResource(dmzim)}
}

// ImportState imports the given attribute values into [DnsManagedZoneIamMember]'s state.
func (dmzim *DnsManagedZoneIamMember) ImportState(av io.Reader) error {
	dmzim.state = &dnsManagedZoneIamMemberState{}
	if err := json.NewDecoder(av).Decode(dmzim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dmzim.Type(), dmzim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DnsManagedZoneIamMember] has state.
func (dmzim *DnsManagedZoneIamMember) State() (*dnsManagedZoneIamMemberState, bool) {
	return dmzim.state, dmzim.state != nil
}

// StateMust returns the state for [DnsManagedZoneIamMember]. Panics if the state is nil.
func (dmzim *DnsManagedZoneIamMember) StateMust() *dnsManagedZoneIamMemberState {
	if dmzim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dmzim.Type(), dmzim.LocalName()))
	}
	return dmzim.state
}

// DnsManagedZoneIamMemberArgs contains the configurations for google_dns_managed_zone_iam_member.
type DnsManagedZoneIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagedZone: string, required
	ManagedZone terra.StringValue `hcl:"managed_zone,attr" validate:"required"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *dnsmanagedzoneiammember.Condition `hcl:"condition,block"`
}
type dnsManagedZoneIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dns_managed_zone_iam_member.
func (dmzim dnsManagedZoneIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dmzim.ref.Append("etag"))
}

// Id returns a reference to field id of google_dns_managed_zone_iam_member.
func (dmzim dnsManagedZoneIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmzim.ref.Append("id"))
}

// ManagedZone returns a reference to field managed_zone of google_dns_managed_zone_iam_member.
func (dmzim dnsManagedZoneIamMemberAttributes) ManagedZone() terra.StringValue {
	return terra.ReferenceAsString(dmzim.ref.Append("managed_zone"))
}

// Member returns a reference to field member of google_dns_managed_zone_iam_member.
func (dmzim dnsManagedZoneIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(dmzim.ref.Append("member"))
}

// Project returns a reference to field project of google_dns_managed_zone_iam_member.
func (dmzim dnsManagedZoneIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dmzim.ref.Append("project"))
}

// Role returns a reference to field role of google_dns_managed_zone_iam_member.
func (dmzim dnsManagedZoneIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dmzim.ref.Append("role"))
}

func (dmzim dnsManagedZoneIamMemberAttributes) Condition() terra.ListValue[dnsmanagedzoneiammember.ConditionAttributes] {
	return terra.ReferenceAsList[dnsmanagedzoneiammember.ConditionAttributes](dmzim.ref.Append("condition"))
}

type dnsManagedZoneIamMemberState struct {
	Etag        string                                   `json:"etag"`
	Id          string                                   `json:"id"`
	ManagedZone string                                   `json:"managed_zone"`
	Member      string                                   `json:"member"`
	Project     string                                   `json:"project"`
	Role        string                                   `json:"role"`
	Condition   []dnsmanagedzoneiammember.ConditionState `json:"condition"`
}
