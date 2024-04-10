// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	organizationiammember "github.com/golingon/terraproviders/google/5.24.0/organizationiammember"
	"io"
)

// NewOrganizationIamMember creates a new instance of [OrganizationIamMember].
func NewOrganizationIamMember(name string, args OrganizationIamMemberArgs) *OrganizationIamMember {
	return &OrganizationIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OrganizationIamMember)(nil)

// OrganizationIamMember represents the Terraform resource google_organization_iam_member.
type OrganizationIamMember struct {
	Name      string
	Args      OrganizationIamMemberArgs
	state     *organizationIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OrganizationIamMember].
func (oim *OrganizationIamMember) Type() string {
	return "google_organization_iam_member"
}

// LocalName returns the local name for [OrganizationIamMember].
func (oim *OrganizationIamMember) LocalName() string {
	return oim.Name
}

// Configuration returns the configuration (args) for [OrganizationIamMember].
func (oim *OrganizationIamMember) Configuration() interface{} {
	return oim.Args
}

// DependOn is used for other resources to depend on [OrganizationIamMember].
func (oim *OrganizationIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(oim)
}

// Dependencies returns the list of resources [OrganizationIamMember] depends_on.
func (oim *OrganizationIamMember) Dependencies() terra.Dependencies {
	return oim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OrganizationIamMember].
func (oim *OrganizationIamMember) LifecycleManagement() *terra.Lifecycle {
	return oim.Lifecycle
}

// Attributes returns the attributes for [OrganizationIamMember].
func (oim *OrganizationIamMember) Attributes() organizationIamMemberAttributes {
	return organizationIamMemberAttributes{ref: terra.ReferenceResource(oim)}
}

// ImportState imports the given attribute values into [OrganizationIamMember]'s state.
func (oim *OrganizationIamMember) ImportState(av io.Reader) error {
	oim.state = &organizationIamMemberState{}
	if err := json.NewDecoder(av).Decode(oim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oim.Type(), oim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OrganizationIamMember] has state.
func (oim *OrganizationIamMember) State() (*organizationIamMemberState, bool) {
	return oim.state, oim.state != nil
}

// StateMust returns the state for [OrganizationIamMember]. Panics if the state is nil.
func (oim *OrganizationIamMember) StateMust() *organizationIamMemberState {
	if oim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oim.Type(), oim.LocalName()))
	}
	return oim.state
}

// OrganizationIamMemberArgs contains the configurations for google_organization_iam_member.
type OrganizationIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *organizationiammember.Condition `hcl:"condition,block"`
}
type organizationIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_organization_iam_member.
func (oim organizationIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(oim.ref.Append("etag"))
}

// Id returns a reference to field id of google_organization_iam_member.
func (oim organizationIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oim.ref.Append("id"))
}

// Member returns a reference to field member of google_organization_iam_member.
func (oim organizationIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(oim.ref.Append("member"))
}

// OrgId returns a reference to field org_id of google_organization_iam_member.
func (oim organizationIamMemberAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(oim.ref.Append("org_id"))
}

// Role returns a reference to field role of google_organization_iam_member.
func (oim organizationIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(oim.ref.Append("role"))
}

func (oim organizationIamMemberAttributes) Condition() terra.ListValue[organizationiammember.ConditionAttributes] {
	return terra.ReferenceAsList[organizationiammember.ConditionAttributes](oim.ref.Append("condition"))
}

type organizationIamMemberState struct {
	Etag      string                                 `json:"etag"`
	Id        string                                 `json:"id"`
	Member    string                                 `json:"member"`
	OrgId     string                                 `json:"org_id"`
	Role      string                                 `json:"role"`
	Condition []organizationiammember.ConditionState `json:"condition"`
}
