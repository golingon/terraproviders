// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	sourcereporepositoryiambinding "github.com/golingon/terraproviders/google/4.72.1/sourcereporepositoryiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSourcerepoRepositoryIamBinding creates a new instance of [SourcerepoRepositoryIamBinding].
func NewSourcerepoRepositoryIamBinding(name string, args SourcerepoRepositoryIamBindingArgs) *SourcerepoRepositoryIamBinding {
	return &SourcerepoRepositoryIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SourcerepoRepositoryIamBinding)(nil)

// SourcerepoRepositoryIamBinding represents the Terraform resource google_sourcerepo_repository_iam_binding.
type SourcerepoRepositoryIamBinding struct {
	Name      string
	Args      SourcerepoRepositoryIamBindingArgs
	state     *sourcerepoRepositoryIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SourcerepoRepositoryIamBinding].
func (srib *SourcerepoRepositoryIamBinding) Type() string {
	return "google_sourcerepo_repository_iam_binding"
}

// LocalName returns the local name for [SourcerepoRepositoryIamBinding].
func (srib *SourcerepoRepositoryIamBinding) LocalName() string {
	return srib.Name
}

// Configuration returns the configuration (args) for [SourcerepoRepositoryIamBinding].
func (srib *SourcerepoRepositoryIamBinding) Configuration() interface{} {
	return srib.Args
}

// DependOn is used for other resources to depend on [SourcerepoRepositoryIamBinding].
func (srib *SourcerepoRepositoryIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(srib)
}

// Dependencies returns the list of resources [SourcerepoRepositoryIamBinding] depends_on.
func (srib *SourcerepoRepositoryIamBinding) Dependencies() terra.Dependencies {
	return srib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SourcerepoRepositoryIamBinding].
func (srib *SourcerepoRepositoryIamBinding) LifecycleManagement() *terra.Lifecycle {
	return srib.Lifecycle
}

// Attributes returns the attributes for [SourcerepoRepositoryIamBinding].
func (srib *SourcerepoRepositoryIamBinding) Attributes() sourcerepoRepositoryIamBindingAttributes {
	return sourcerepoRepositoryIamBindingAttributes{ref: terra.ReferenceResource(srib)}
}

// ImportState imports the given attribute values into [SourcerepoRepositoryIamBinding]'s state.
func (srib *SourcerepoRepositoryIamBinding) ImportState(av io.Reader) error {
	srib.state = &sourcerepoRepositoryIamBindingState{}
	if err := json.NewDecoder(av).Decode(srib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srib.Type(), srib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SourcerepoRepositoryIamBinding] has state.
func (srib *SourcerepoRepositoryIamBinding) State() (*sourcerepoRepositoryIamBindingState, bool) {
	return srib.state, srib.state != nil
}

// StateMust returns the state for [SourcerepoRepositoryIamBinding]. Panics if the state is nil.
func (srib *SourcerepoRepositoryIamBinding) StateMust() *sourcerepoRepositoryIamBindingState {
	if srib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srib.Type(), srib.LocalName()))
	}
	return srib.state
}

// SourcerepoRepositoryIamBindingArgs contains the configurations for google_sourcerepo_repository_iam_binding.
type SourcerepoRepositoryIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Repository: string, required
	Repository terra.StringValue `hcl:"repository,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *sourcereporepositoryiambinding.Condition `hcl:"condition,block"`
}
type sourcerepoRepositoryIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_sourcerepo_repository_iam_binding.
func (srib sourcerepoRepositoryIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(srib.ref.Append("etag"))
}

// Id returns a reference to field id of google_sourcerepo_repository_iam_binding.
func (srib sourcerepoRepositoryIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srib.ref.Append("id"))
}

// Members returns a reference to field members of google_sourcerepo_repository_iam_binding.
func (srib sourcerepoRepositoryIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](srib.ref.Append("members"))
}

// Project returns a reference to field project of google_sourcerepo_repository_iam_binding.
func (srib sourcerepoRepositoryIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(srib.ref.Append("project"))
}

// Repository returns a reference to field repository of google_sourcerepo_repository_iam_binding.
func (srib sourcerepoRepositoryIamBindingAttributes) Repository() terra.StringValue {
	return terra.ReferenceAsString(srib.ref.Append("repository"))
}

// Role returns a reference to field role of google_sourcerepo_repository_iam_binding.
func (srib sourcerepoRepositoryIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(srib.ref.Append("role"))
}

func (srib sourcerepoRepositoryIamBindingAttributes) Condition() terra.ListValue[sourcereporepositoryiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[sourcereporepositoryiambinding.ConditionAttributes](srib.ref.Append("condition"))
}

type sourcerepoRepositoryIamBindingState struct {
	Etag       string                                          `json:"etag"`
	Id         string                                          `json:"id"`
	Members    []string                                        `json:"members"`
	Project    string                                          `json:"project"`
	Repository string                                          `json:"repository"`
	Role       string                                          `json:"role"`
	Condition  []sourcereporepositoryiambinding.ConditionState `json:"condition"`
}
