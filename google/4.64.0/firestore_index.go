// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	firestoreindex "github.com/golingon/terraproviders/google/4.64.0/firestoreindex"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirestoreIndex creates a new instance of [FirestoreIndex].
func NewFirestoreIndex(name string, args FirestoreIndexArgs) *FirestoreIndex {
	return &FirestoreIndex{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirestoreIndex)(nil)

// FirestoreIndex represents the Terraform resource google_firestore_index.
type FirestoreIndex struct {
	Name      string
	Args      FirestoreIndexArgs
	state     *firestoreIndexState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirestoreIndex].
func (fi *FirestoreIndex) Type() string {
	return "google_firestore_index"
}

// LocalName returns the local name for [FirestoreIndex].
func (fi *FirestoreIndex) LocalName() string {
	return fi.Name
}

// Configuration returns the configuration (args) for [FirestoreIndex].
func (fi *FirestoreIndex) Configuration() interface{} {
	return fi.Args
}

// DependOn is used for other resources to depend on [FirestoreIndex].
func (fi *FirestoreIndex) DependOn() terra.Reference {
	return terra.ReferenceResource(fi)
}

// Dependencies returns the list of resources [FirestoreIndex] depends_on.
func (fi *FirestoreIndex) Dependencies() terra.Dependencies {
	return fi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirestoreIndex].
func (fi *FirestoreIndex) LifecycleManagement() *terra.Lifecycle {
	return fi.Lifecycle
}

// Attributes returns the attributes for [FirestoreIndex].
func (fi *FirestoreIndex) Attributes() firestoreIndexAttributes {
	return firestoreIndexAttributes{ref: terra.ReferenceResource(fi)}
}

// ImportState imports the given attribute values into [FirestoreIndex]'s state.
func (fi *FirestoreIndex) ImportState(av io.Reader) error {
	fi.state = &firestoreIndexState{}
	if err := json.NewDecoder(av).Decode(fi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fi.Type(), fi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirestoreIndex] has state.
func (fi *FirestoreIndex) State() (*firestoreIndexState, bool) {
	return fi.state, fi.state != nil
}

// StateMust returns the state for [FirestoreIndex]. Panics if the state is nil.
func (fi *FirestoreIndex) StateMust() *firestoreIndexState {
	if fi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fi.Type(), fi.LocalName()))
	}
	return fi.state
}

// FirestoreIndexArgs contains the configurations for google_firestore_index.
type FirestoreIndexArgs struct {
	// Collection: string, required
	Collection terra.StringValue `hcl:"collection,attr" validate:"required"`
	// Database: string, optional
	Database terra.StringValue `hcl:"database,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// QueryScope: string, optional
	QueryScope terra.StringValue `hcl:"query_scope,attr"`
	// Fields: min=2
	Fields []firestoreindex.Fields `hcl:"fields,block" validate:"min=2"`
	// Timeouts: optional
	Timeouts *firestoreindex.Timeouts `hcl:"timeouts,block"`
}
type firestoreIndexAttributes struct {
	ref terra.Reference
}

// Collection returns a reference to field collection of google_firestore_index.
func (fi firestoreIndexAttributes) Collection() terra.StringValue {
	return terra.ReferenceAsString(fi.ref.Append("collection"))
}

// Database returns a reference to field database of google_firestore_index.
func (fi firestoreIndexAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(fi.ref.Append("database"))
}

// Id returns a reference to field id of google_firestore_index.
func (fi firestoreIndexAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fi.ref.Append("id"))
}

// Name returns a reference to field name of google_firestore_index.
func (fi firestoreIndexAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fi.ref.Append("name"))
}

// Project returns a reference to field project of google_firestore_index.
func (fi firestoreIndexAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(fi.ref.Append("project"))
}

// QueryScope returns a reference to field query_scope of google_firestore_index.
func (fi firestoreIndexAttributes) QueryScope() terra.StringValue {
	return terra.ReferenceAsString(fi.ref.Append("query_scope"))
}

func (fi firestoreIndexAttributes) Fields() terra.ListValue[firestoreindex.FieldsAttributes] {
	return terra.ReferenceAsList[firestoreindex.FieldsAttributes](fi.ref.Append("fields"))
}

func (fi firestoreIndexAttributes) Timeouts() firestoreindex.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firestoreindex.TimeoutsAttributes](fi.ref.Append("timeouts"))
}

type firestoreIndexState struct {
	Collection string                        `json:"collection"`
	Database   string                        `json:"database"`
	Id         string                        `json:"id"`
	Name       string                        `json:"name"`
	Project    string                        `json:"project"`
	QueryScope string                        `json:"query_scope"`
	Fields     []firestoreindex.FieldsState  `json:"fields"`
	Timeouts   *firestoreindex.TimeoutsState `json:"timeouts"`
}
