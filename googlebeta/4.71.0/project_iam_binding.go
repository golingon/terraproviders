// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	projectiambinding "github.com/golingon/terraproviders/googlebeta/4.71.0/projectiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewProjectIamBinding creates a new instance of [ProjectIamBinding].
func NewProjectIamBinding(name string, args ProjectIamBindingArgs) *ProjectIamBinding {
	return &ProjectIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ProjectIamBinding)(nil)

// ProjectIamBinding represents the Terraform resource google_project_iam_binding.
type ProjectIamBinding struct {
	Name      string
	Args      ProjectIamBindingArgs
	state     *projectIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ProjectIamBinding].
func (pib *ProjectIamBinding) Type() string {
	return "google_project_iam_binding"
}

// LocalName returns the local name for [ProjectIamBinding].
func (pib *ProjectIamBinding) LocalName() string {
	return pib.Name
}

// Configuration returns the configuration (args) for [ProjectIamBinding].
func (pib *ProjectIamBinding) Configuration() interface{} {
	return pib.Args
}

// DependOn is used for other resources to depend on [ProjectIamBinding].
func (pib *ProjectIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(pib)
}

// Dependencies returns the list of resources [ProjectIamBinding] depends_on.
func (pib *ProjectIamBinding) Dependencies() terra.Dependencies {
	return pib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ProjectIamBinding].
func (pib *ProjectIamBinding) LifecycleManagement() *terra.Lifecycle {
	return pib.Lifecycle
}

// Attributes returns the attributes for [ProjectIamBinding].
func (pib *ProjectIamBinding) Attributes() projectIamBindingAttributes {
	return projectIamBindingAttributes{ref: terra.ReferenceResource(pib)}
}

// ImportState imports the given attribute values into [ProjectIamBinding]'s state.
func (pib *ProjectIamBinding) ImportState(av io.Reader) error {
	pib.state = &projectIamBindingState{}
	if err := json.NewDecoder(av).Decode(pib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pib.Type(), pib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ProjectIamBinding] has state.
func (pib *ProjectIamBinding) State() (*projectIamBindingState, bool) {
	return pib.state, pib.state != nil
}

// StateMust returns the state for [ProjectIamBinding]. Panics if the state is nil.
func (pib *ProjectIamBinding) StateMust() *projectIamBindingState {
	if pib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pib.Type(), pib.LocalName()))
	}
	return pib.state
}

// ProjectIamBindingArgs contains the configurations for google_project_iam_binding.
type ProjectIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *projectiambinding.Condition `hcl:"condition,block"`
}
type projectIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_project_iam_binding.
func (pib projectIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(pib.ref.Append("etag"))
}

// Id returns a reference to field id of google_project_iam_binding.
func (pib projectIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pib.ref.Append("id"))
}

// Members returns a reference to field members of google_project_iam_binding.
func (pib projectIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pib.ref.Append("members"))
}

// Project returns a reference to field project of google_project_iam_binding.
func (pib projectIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pib.ref.Append("project"))
}

// Role returns a reference to field role of google_project_iam_binding.
func (pib projectIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(pib.ref.Append("role"))
}

func (pib projectIamBindingAttributes) Condition() terra.ListValue[projectiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[projectiambinding.ConditionAttributes](pib.ref.Append("condition"))
}

type projectIamBindingState struct {
	Etag      string                             `json:"etag"`
	Id        string                             `json:"id"`
	Members   []string                           `json:"members"`
	Project   string                             `json:"project"`
	Role      string                             `json:"role"`
	Condition []projectiambinding.ConditionState `json:"condition"`
}