// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	notebookslocation "github.com/golingon/terraproviders/googlebeta/4.62.0/notebookslocation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNotebooksLocation creates a new instance of [NotebooksLocation].
func NewNotebooksLocation(name string, args NotebooksLocationArgs) *NotebooksLocation {
	return &NotebooksLocation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NotebooksLocation)(nil)

// NotebooksLocation represents the Terraform resource google_notebooks_location.
type NotebooksLocation struct {
	Name      string
	Args      NotebooksLocationArgs
	state     *notebooksLocationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NotebooksLocation].
func (nl *NotebooksLocation) Type() string {
	return "google_notebooks_location"
}

// LocalName returns the local name for [NotebooksLocation].
func (nl *NotebooksLocation) LocalName() string {
	return nl.Name
}

// Configuration returns the configuration (args) for [NotebooksLocation].
func (nl *NotebooksLocation) Configuration() interface{} {
	return nl.Args
}

// DependOn is used for other resources to depend on [NotebooksLocation].
func (nl *NotebooksLocation) DependOn() terra.Reference {
	return terra.ReferenceResource(nl)
}

// Dependencies returns the list of resources [NotebooksLocation] depends_on.
func (nl *NotebooksLocation) Dependencies() terra.Dependencies {
	return nl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NotebooksLocation].
func (nl *NotebooksLocation) LifecycleManagement() *terra.Lifecycle {
	return nl.Lifecycle
}

// Attributes returns the attributes for [NotebooksLocation].
func (nl *NotebooksLocation) Attributes() notebooksLocationAttributes {
	return notebooksLocationAttributes{ref: terra.ReferenceResource(nl)}
}

// ImportState imports the given attribute values into [NotebooksLocation]'s state.
func (nl *NotebooksLocation) ImportState(av io.Reader) error {
	nl.state = &notebooksLocationState{}
	if err := json.NewDecoder(av).Decode(nl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nl.Type(), nl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NotebooksLocation] has state.
func (nl *NotebooksLocation) State() (*notebooksLocationState, bool) {
	return nl.state, nl.state != nil
}

// StateMust returns the state for [NotebooksLocation]. Panics if the state is nil.
func (nl *NotebooksLocation) StateMust() *notebooksLocationState {
	if nl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nl.Type(), nl.LocalName()))
	}
	return nl.state
}

// NotebooksLocationArgs contains the configurations for google_notebooks_location.
type NotebooksLocationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *notebookslocation.Timeouts `hcl:"timeouts,block"`
}
type notebooksLocationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_notebooks_location.
func (nl notebooksLocationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("id"))
}

// Name returns a reference to field name of google_notebooks_location.
func (nl notebooksLocationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("name"))
}

// Project returns a reference to field project of google_notebooks_location.
func (nl notebooksLocationAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_notebooks_location.
func (nl notebooksLocationAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(nl.ref.Append("self_link"))
}

func (nl notebooksLocationAttributes) Timeouts() notebookslocation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[notebookslocation.TimeoutsAttributes](nl.ref.Append("timeouts"))
}

type notebooksLocationState struct {
	Id       string                           `json:"id"`
	Name     string                           `json:"name"`
	Project  string                           `json:"project"`
	SelfLink string                           `json:"self_link"`
	Timeouts *notebookslocation.TimeoutsState `json:"timeouts"`
}
