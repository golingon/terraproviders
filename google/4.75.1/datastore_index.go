// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	datastoreindex "github.com/golingon/terraproviders/google/4.75.1/datastoreindex"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatastoreIndex creates a new instance of [DatastoreIndex].
func NewDatastoreIndex(name string, args DatastoreIndexArgs) *DatastoreIndex {
	return &DatastoreIndex{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatastoreIndex)(nil)

// DatastoreIndex represents the Terraform resource google_datastore_index.
type DatastoreIndex struct {
	Name      string
	Args      DatastoreIndexArgs
	state     *datastoreIndexState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatastoreIndex].
func (di *DatastoreIndex) Type() string {
	return "google_datastore_index"
}

// LocalName returns the local name for [DatastoreIndex].
func (di *DatastoreIndex) LocalName() string {
	return di.Name
}

// Configuration returns the configuration (args) for [DatastoreIndex].
func (di *DatastoreIndex) Configuration() interface{} {
	return di.Args
}

// DependOn is used for other resources to depend on [DatastoreIndex].
func (di *DatastoreIndex) DependOn() terra.Reference {
	return terra.ReferenceResource(di)
}

// Dependencies returns the list of resources [DatastoreIndex] depends_on.
func (di *DatastoreIndex) Dependencies() terra.Dependencies {
	return di.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatastoreIndex].
func (di *DatastoreIndex) LifecycleManagement() *terra.Lifecycle {
	return di.Lifecycle
}

// Attributes returns the attributes for [DatastoreIndex].
func (di *DatastoreIndex) Attributes() datastoreIndexAttributes {
	return datastoreIndexAttributes{ref: terra.ReferenceResource(di)}
}

// ImportState imports the given attribute values into [DatastoreIndex]'s state.
func (di *DatastoreIndex) ImportState(av io.Reader) error {
	di.state = &datastoreIndexState{}
	if err := json.NewDecoder(av).Decode(di.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", di.Type(), di.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatastoreIndex] has state.
func (di *DatastoreIndex) State() (*datastoreIndexState, bool) {
	return di.state, di.state != nil
}

// StateMust returns the state for [DatastoreIndex]. Panics if the state is nil.
func (di *DatastoreIndex) StateMust() *datastoreIndexState {
	if di.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", di.Type(), di.LocalName()))
	}
	return di.state
}

// DatastoreIndexArgs contains the configurations for google_datastore_index.
type DatastoreIndexArgs struct {
	// Ancestor: string, optional
	Ancestor terra.StringValue `hcl:"ancestor,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Kind: string, required
	Kind terra.StringValue `hcl:"kind,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Properties: min=0
	Properties []datastoreindex.Properties `hcl:"properties,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datastoreindex.Timeouts `hcl:"timeouts,block"`
}
type datastoreIndexAttributes struct {
	ref terra.Reference
}

// Ancestor returns a reference to field ancestor of google_datastore_index.
func (di datastoreIndexAttributes) Ancestor() terra.StringValue {
	return terra.ReferenceAsString(di.ref.Append("ancestor"))
}

// Id returns a reference to field id of google_datastore_index.
func (di datastoreIndexAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(di.ref.Append("id"))
}

// IndexId returns a reference to field index_id of google_datastore_index.
func (di datastoreIndexAttributes) IndexId() terra.StringValue {
	return terra.ReferenceAsString(di.ref.Append("index_id"))
}

// Kind returns a reference to field kind of google_datastore_index.
func (di datastoreIndexAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(di.ref.Append("kind"))
}

// Project returns a reference to field project of google_datastore_index.
func (di datastoreIndexAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(di.ref.Append("project"))
}

func (di datastoreIndexAttributes) Properties() terra.ListValue[datastoreindex.PropertiesAttributes] {
	return terra.ReferenceAsList[datastoreindex.PropertiesAttributes](di.ref.Append("properties"))
}

func (di datastoreIndexAttributes) Timeouts() datastoreindex.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datastoreindex.TimeoutsAttributes](di.ref.Append("timeouts"))
}

type datastoreIndexState struct {
	Ancestor   string                           `json:"ancestor"`
	Id         string                           `json:"id"`
	IndexId    string                           `json:"index_id"`
	Kind       string                           `json:"kind"`
	Project    string                           `json:"project"`
	Properties []datastoreindex.PropertiesState `json:"properties"`
	Timeouts   *datastoreindex.TimeoutsState    `json:"timeouts"`
}
