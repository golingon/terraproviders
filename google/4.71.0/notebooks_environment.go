// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	notebooksenvironment "github.com/golingon/terraproviders/google/4.71.0/notebooksenvironment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNotebooksEnvironment creates a new instance of [NotebooksEnvironment].
func NewNotebooksEnvironment(name string, args NotebooksEnvironmentArgs) *NotebooksEnvironment {
	return &NotebooksEnvironment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NotebooksEnvironment)(nil)

// NotebooksEnvironment represents the Terraform resource google_notebooks_environment.
type NotebooksEnvironment struct {
	Name      string
	Args      NotebooksEnvironmentArgs
	state     *notebooksEnvironmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NotebooksEnvironment].
func (ne *NotebooksEnvironment) Type() string {
	return "google_notebooks_environment"
}

// LocalName returns the local name for [NotebooksEnvironment].
func (ne *NotebooksEnvironment) LocalName() string {
	return ne.Name
}

// Configuration returns the configuration (args) for [NotebooksEnvironment].
func (ne *NotebooksEnvironment) Configuration() interface{} {
	return ne.Args
}

// DependOn is used for other resources to depend on [NotebooksEnvironment].
func (ne *NotebooksEnvironment) DependOn() terra.Reference {
	return terra.ReferenceResource(ne)
}

// Dependencies returns the list of resources [NotebooksEnvironment] depends_on.
func (ne *NotebooksEnvironment) Dependencies() terra.Dependencies {
	return ne.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NotebooksEnvironment].
func (ne *NotebooksEnvironment) LifecycleManagement() *terra.Lifecycle {
	return ne.Lifecycle
}

// Attributes returns the attributes for [NotebooksEnvironment].
func (ne *NotebooksEnvironment) Attributes() notebooksEnvironmentAttributes {
	return notebooksEnvironmentAttributes{ref: terra.ReferenceResource(ne)}
}

// ImportState imports the given attribute values into [NotebooksEnvironment]'s state.
func (ne *NotebooksEnvironment) ImportState(av io.Reader) error {
	ne.state = &notebooksEnvironmentState{}
	if err := json.NewDecoder(av).Decode(ne.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ne.Type(), ne.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NotebooksEnvironment] has state.
func (ne *NotebooksEnvironment) State() (*notebooksEnvironmentState, bool) {
	return ne.state, ne.state != nil
}

// StateMust returns the state for [NotebooksEnvironment]. Panics if the state is nil.
func (ne *NotebooksEnvironment) StateMust() *notebooksEnvironmentState {
	if ne.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ne.Type(), ne.LocalName()))
	}
	return ne.state
}

// NotebooksEnvironmentArgs contains the configurations for google_notebooks_environment.
type NotebooksEnvironmentArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PostStartupScript: string, optional
	PostStartupScript terra.StringValue `hcl:"post_startup_script,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ContainerImage: optional
	ContainerImage *notebooksenvironment.ContainerImage `hcl:"container_image,block"`
	// Timeouts: optional
	Timeouts *notebooksenvironment.Timeouts `hcl:"timeouts,block"`
	// VmImage: optional
	VmImage *notebooksenvironment.VmImage `hcl:"vm_image,block"`
}
type notebooksEnvironmentAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_notebooks_environment.
func (ne notebooksEnvironmentAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ne.ref.Append("create_time"))
}

// Description returns a reference to field description of google_notebooks_environment.
func (ne notebooksEnvironmentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ne.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_notebooks_environment.
func (ne notebooksEnvironmentAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ne.ref.Append("display_name"))
}

// Id returns a reference to field id of google_notebooks_environment.
func (ne notebooksEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ne.ref.Append("id"))
}

// Location returns a reference to field location of google_notebooks_environment.
func (ne notebooksEnvironmentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ne.ref.Append("location"))
}

// Name returns a reference to field name of google_notebooks_environment.
func (ne notebooksEnvironmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ne.ref.Append("name"))
}

// PostStartupScript returns a reference to field post_startup_script of google_notebooks_environment.
func (ne notebooksEnvironmentAttributes) PostStartupScript() terra.StringValue {
	return terra.ReferenceAsString(ne.ref.Append("post_startup_script"))
}

// Project returns a reference to field project of google_notebooks_environment.
func (ne notebooksEnvironmentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ne.ref.Append("project"))
}

func (ne notebooksEnvironmentAttributes) ContainerImage() terra.ListValue[notebooksenvironment.ContainerImageAttributes] {
	return terra.ReferenceAsList[notebooksenvironment.ContainerImageAttributes](ne.ref.Append("container_image"))
}

func (ne notebooksEnvironmentAttributes) Timeouts() notebooksenvironment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[notebooksenvironment.TimeoutsAttributes](ne.ref.Append("timeouts"))
}

func (ne notebooksEnvironmentAttributes) VmImage() terra.ListValue[notebooksenvironment.VmImageAttributes] {
	return terra.ReferenceAsList[notebooksenvironment.VmImageAttributes](ne.ref.Append("vm_image"))
}

type notebooksEnvironmentState struct {
	CreateTime        string                                     `json:"create_time"`
	Description       string                                     `json:"description"`
	DisplayName       string                                     `json:"display_name"`
	Id                string                                     `json:"id"`
	Location          string                                     `json:"location"`
	Name              string                                     `json:"name"`
	PostStartupScript string                                     `json:"post_startup_script"`
	Project           string                                     `json:"project"`
	ContainerImage    []notebooksenvironment.ContainerImageState `json:"container_image"`
	Timeouts          *notebooksenvironment.TimeoutsState        `json:"timeouts"`
	VmImage           []notebooksenvironment.VmImageState        `json:"vm_image"`
}