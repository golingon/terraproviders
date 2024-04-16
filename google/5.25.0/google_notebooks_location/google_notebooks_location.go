// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_notebooks_location

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_notebooks_location.
type Resource struct {
	Name      string
	Args      Args
	state     *googleNotebooksLocationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gnl *Resource) Type() string {
	return "google_notebooks_location"
}

// LocalName returns the local name for [Resource].
func (gnl *Resource) LocalName() string {
	return gnl.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gnl *Resource) Configuration() interface{} {
	return gnl.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gnl *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gnl)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gnl *Resource) Dependencies() terra.Dependencies {
	return gnl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gnl *Resource) LifecycleManagement() *terra.Lifecycle {
	return gnl.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gnl *Resource) Attributes() googleNotebooksLocationAttributes {
	return googleNotebooksLocationAttributes{ref: terra.ReferenceResource(gnl)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gnl *Resource) ImportState(state io.Reader) error {
	gnl.state = &googleNotebooksLocationState{}
	if err := json.NewDecoder(state).Decode(gnl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gnl.Type(), gnl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gnl *Resource) State() (*googleNotebooksLocationState, bool) {
	return gnl.state, gnl.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gnl *Resource) StateMust() *googleNotebooksLocationState {
	if gnl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gnl.Type(), gnl.LocalName()))
	}
	return gnl.state
}

// Args contains the configurations for google_notebooks_location.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleNotebooksLocationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_notebooks_location.
func (gnl googleNotebooksLocationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gnl.ref.Append("id"))
}

// Name returns a reference to field name of google_notebooks_location.
func (gnl googleNotebooksLocationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gnl.ref.Append("name"))
}

// Project returns a reference to field project of google_notebooks_location.
func (gnl googleNotebooksLocationAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gnl.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_notebooks_location.
func (gnl googleNotebooksLocationAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gnl.ref.Append("self_link"))
}

func (gnl googleNotebooksLocationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gnl.ref.Append("timeouts"))
}

type googleNotebooksLocationState struct {
	Id       string         `json:"id"`
	Name     string         `json:"name"`
	Project  string         `json:"project"`
	SelfLink string         `json:"self_link"`
	Timeouts *TimeoutsState `json:"timeouts"`
}
