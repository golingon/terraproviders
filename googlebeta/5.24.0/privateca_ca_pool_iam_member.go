// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	privatecacapooliammember "github.com/golingon/terraproviders/googlebeta/5.24.0/privatecacapooliammember"
	"io"
)

// NewPrivatecaCaPoolIamMember creates a new instance of [PrivatecaCaPoolIamMember].
func NewPrivatecaCaPoolIamMember(name string, args PrivatecaCaPoolIamMemberArgs) *PrivatecaCaPoolIamMember {
	return &PrivatecaCaPoolIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivatecaCaPoolIamMember)(nil)

// PrivatecaCaPoolIamMember represents the Terraform resource google_privateca_ca_pool_iam_member.
type PrivatecaCaPoolIamMember struct {
	Name      string
	Args      PrivatecaCaPoolIamMemberArgs
	state     *privatecaCaPoolIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivatecaCaPoolIamMember].
func (pcpim *PrivatecaCaPoolIamMember) Type() string {
	return "google_privateca_ca_pool_iam_member"
}

// LocalName returns the local name for [PrivatecaCaPoolIamMember].
func (pcpim *PrivatecaCaPoolIamMember) LocalName() string {
	return pcpim.Name
}

// Configuration returns the configuration (args) for [PrivatecaCaPoolIamMember].
func (pcpim *PrivatecaCaPoolIamMember) Configuration() interface{} {
	return pcpim.Args
}

// DependOn is used for other resources to depend on [PrivatecaCaPoolIamMember].
func (pcpim *PrivatecaCaPoolIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(pcpim)
}

// Dependencies returns the list of resources [PrivatecaCaPoolIamMember] depends_on.
func (pcpim *PrivatecaCaPoolIamMember) Dependencies() terra.Dependencies {
	return pcpim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivatecaCaPoolIamMember].
func (pcpim *PrivatecaCaPoolIamMember) LifecycleManagement() *terra.Lifecycle {
	return pcpim.Lifecycle
}

// Attributes returns the attributes for [PrivatecaCaPoolIamMember].
func (pcpim *PrivatecaCaPoolIamMember) Attributes() privatecaCaPoolIamMemberAttributes {
	return privatecaCaPoolIamMemberAttributes{ref: terra.ReferenceResource(pcpim)}
}

// ImportState imports the given attribute values into [PrivatecaCaPoolIamMember]'s state.
func (pcpim *PrivatecaCaPoolIamMember) ImportState(av io.Reader) error {
	pcpim.state = &privatecaCaPoolIamMemberState{}
	if err := json.NewDecoder(av).Decode(pcpim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pcpim.Type(), pcpim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivatecaCaPoolIamMember] has state.
func (pcpim *PrivatecaCaPoolIamMember) State() (*privatecaCaPoolIamMemberState, bool) {
	return pcpim.state, pcpim.state != nil
}

// StateMust returns the state for [PrivatecaCaPoolIamMember]. Panics if the state is nil.
func (pcpim *PrivatecaCaPoolIamMember) StateMust() *privatecaCaPoolIamMemberState {
	if pcpim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pcpim.Type(), pcpim.LocalName()))
	}
	return pcpim.state
}

// PrivatecaCaPoolIamMemberArgs contains the configurations for google_privateca_ca_pool_iam_member.
type PrivatecaCaPoolIamMemberArgs struct {
	// CaPool: string, required
	CaPool terra.StringValue `hcl:"ca_pool,attr" validate:"required"`
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
	Condition *privatecacapooliammember.Condition `hcl:"condition,block"`
}
type privatecaCaPoolIamMemberAttributes struct {
	ref terra.Reference
}

// CaPool returns a reference to field ca_pool of google_privateca_ca_pool_iam_member.
func (pcpim privatecaCaPoolIamMemberAttributes) CaPool() terra.StringValue {
	return terra.ReferenceAsString(pcpim.ref.Append("ca_pool"))
}

// Etag returns a reference to field etag of google_privateca_ca_pool_iam_member.
func (pcpim privatecaCaPoolIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(pcpim.ref.Append("etag"))
}

// Id returns a reference to field id of google_privateca_ca_pool_iam_member.
func (pcpim privatecaCaPoolIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pcpim.ref.Append("id"))
}

// Location returns a reference to field location of google_privateca_ca_pool_iam_member.
func (pcpim privatecaCaPoolIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pcpim.ref.Append("location"))
}

// Member returns a reference to field member of google_privateca_ca_pool_iam_member.
func (pcpim privatecaCaPoolIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(pcpim.ref.Append("member"))
}

// Project returns a reference to field project of google_privateca_ca_pool_iam_member.
func (pcpim privatecaCaPoolIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pcpim.ref.Append("project"))
}

// Role returns a reference to field role of google_privateca_ca_pool_iam_member.
func (pcpim privatecaCaPoolIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(pcpim.ref.Append("role"))
}

func (pcpim privatecaCaPoolIamMemberAttributes) Condition() terra.ListValue[privatecacapooliammember.ConditionAttributes] {
	return terra.ReferenceAsList[privatecacapooliammember.ConditionAttributes](pcpim.ref.Append("condition"))
}

type privatecaCaPoolIamMemberState struct {
	CaPool    string                                    `json:"ca_pool"`
	Etag      string                                    `json:"etag"`
	Id        string                                    `json:"id"`
	Location  string                                    `json:"location"`
	Member    string                                    `json:"member"`
	Project   string                                    `json:"project"`
	Role      string                                    `json:"role"`
	Condition []privatecacapooliammember.ConditionState `json:"condition"`
}
