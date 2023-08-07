// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	sourcereporepositoryiammember "github.com/golingon/terraproviders/google/4.76.0/sourcereporepositoryiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSourcerepoRepositoryIamMember creates a new instance of [SourcerepoRepositoryIamMember].
func NewSourcerepoRepositoryIamMember(name string, args SourcerepoRepositoryIamMemberArgs) *SourcerepoRepositoryIamMember {
	return &SourcerepoRepositoryIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SourcerepoRepositoryIamMember)(nil)

// SourcerepoRepositoryIamMember represents the Terraform resource google_sourcerepo_repository_iam_member.
type SourcerepoRepositoryIamMember struct {
	Name      string
	Args      SourcerepoRepositoryIamMemberArgs
	state     *sourcerepoRepositoryIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SourcerepoRepositoryIamMember].
func (srim *SourcerepoRepositoryIamMember) Type() string {
	return "google_sourcerepo_repository_iam_member"
}

// LocalName returns the local name for [SourcerepoRepositoryIamMember].
func (srim *SourcerepoRepositoryIamMember) LocalName() string {
	return srim.Name
}

// Configuration returns the configuration (args) for [SourcerepoRepositoryIamMember].
func (srim *SourcerepoRepositoryIamMember) Configuration() interface{} {
	return srim.Args
}

// DependOn is used for other resources to depend on [SourcerepoRepositoryIamMember].
func (srim *SourcerepoRepositoryIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(srim)
}

// Dependencies returns the list of resources [SourcerepoRepositoryIamMember] depends_on.
func (srim *SourcerepoRepositoryIamMember) Dependencies() terra.Dependencies {
	return srim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SourcerepoRepositoryIamMember].
func (srim *SourcerepoRepositoryIamMember) LifecycleManagement() *terra.Lifecycle {
	return srim.Lifecycle
}

// Attributes returns the attributes for [SourcerepoRepositoryIamMember].
func (srim *SourcerepoRepositoryIamMember) Attributes() sourcerepoRepositoryIamMemberAttributes {
	return sourcerepoRepositoryIamMemberAttributes{ref: terra.ReferenceResource(srim)}
}

// ImportState imports the given attribute values into [SourcerepoRepositoryIamMember]'s state.
func (srim *SourcerepoRepositoryIamMember) ImportState(av io.Reader) error {
	srim.state = &sourcerepoRepositoryIamMemberState{}
	if err := json.NewDecoder(av).Decode(srim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srim.Type(), srim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SourcerepoRepositoryIamMember] has state.
func (srim *SourcerepoRepositoryIamMember) State() (*sourcerepoRepositoryIamMemberState, bool) {
	return srim.state, srim.state != nil
}

// StateMust returns the state for [SourcerepoRepositoryIamMember]. Panics if the state is nil.
func (srim *SourcerepoRepositoryIamMember) StateMust() *sourcerepoRepositoryIamMemberState {
	if srim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srim.Type(), srim.LocalName()))
	}
	return srim.state
}

// SourcerepoRepositoryIamMemberArgs contains the configurations for google_sourcerepo_repository_iam_member.
type SourcerepoRepositoryIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Repository: string, required
	Repository terra.StringValue `hcl:"repository,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *sourcereporepositoryiammember.Condition `hcl:"condition,block"`
}
type sourcerepoRepositoryIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_sourcerepo_repository_iam_member.
func (srim sourcerepoRepositoryIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(srim.ref.Append("etag"))
}

// Id returns a reference to field id of google_sourcerepo_repository_iam_member.
func (srim sourcerepoRepositoryIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srim.ref.Append("id"))
}

// Member returns a reference to field member of google_sourcerepo_repository_iam_member.
func (srim sourcerepoRepositoryIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(srim.ref.Append("member"))
}

// Project returns a reference to field project of google_sourcerepo_repository_iam_member.
func (srim sourcerepoRepositoryIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(srim.ref.Append("project"))
}

// Repository returns a reference to field repository of google_sourcerepo_repository_iam_member.
func (srim sourcerepoRepositoryIamMemberAttributes) Repository() terra.StringValue {
	return terra.ReferenceAsString(srim.ref.Append("repository"))
}

// Role returns a reference to field role of google_sourcerepo_repository_iam_member.
func (srim sourcerepoRepositoryIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(srim.ref.Append("role"))
}

func (srim sourcerepoRepositoryIamMemberAttributes) Condition() terra.ListValue[sourcereporepositoryiammember.ConditionAttributes] {
	return terra.ReferenceAsList[sourcereporepositoryiammember.ConditionAttributes](srim.ref.Append("condition"))
}

type sourcerepoRepositoryIamMemberState struct {
	Etag       string                                         `json:"etag"`
	Id         string                                         `json:"id"`
	Member     string                                         `json:"member"`
	Project    string                                         `json:"project"`
	Repository string                                         `json:"repository"`
	Role       string                                         `json:"role"`
	Condition  []sourcereporepositoryiammember.ConditionState `json:"condition"`
}
