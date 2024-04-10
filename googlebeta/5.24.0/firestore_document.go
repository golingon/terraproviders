// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	firestoredocument "github.com/golingon/terraproviders/googlebeta/5.24.0/firestoredocument"
	"io"
)

// NewFirestoreDocument creates a new instance of [FirestoreDocument].
func NewFirestoreDocument(name string, args FirestoreDocumentArgs) *FirestoreDocument {
	return &FirestoreDocument{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirestoreDocument)(nil)

// FirestoreDocument represents the Terraform resource google_firestore_document.
type FirestoreDocument struct {
	Name      string
	Args      FirestoreDocumentArgs
	state     *firestoreDocumentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirestoreDocument].
func (fd *FirestoreDocument) Type() string {
	return "google_firestore_document"
}

// LocalName returns the local name for [FirestoreDocument].
func (fd *FirestoreDocument) LocalName() string {
	return fd.Name
}

// Configuration returns the configuration (args) for [FirestoreDocument].
func (fd *FirestoreDocument) Configuration() interface{} {
	return fd.Args
}

// DependOn is used for other resources to depend on [FirestoreDocument].
func (fd *FirestoreDocument) DependOn() terra.Reference {
	return terra.ReferenceResource(fd)
}

// Dependencies returns the list of resources [FirestoreDocument] depends_on.
func (fd *FirestoreDocument) Dependencies() terra.Dependencies {
	return fd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirestoreDocument].
func (fd *FirestoreDocument) LifecycleManagement() *terra.Lifecycle {
	return fd.Lifecycle
}

// Attributes returns the attributes for [FirestoreDocument].
func (fd *FirestoreDocument) Attributes() firestoreDocumentAttributes {
	return firestoreDocumentAttributes{ref: terra.ReferenceResource(fd)}
}

// ImportState imports the given attribute values into [FirestoreDocument]'s state.
func (fd *FirestoreDocument) ImportState(av io.Reader) error {
	fd.state = &firestoreDocumentState{}
	if err := json.NewDecoder(av).Decode(fd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fd.Type(), fd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirestoreDocument] has state.
func (fd *FirestoreDocument) State() (*firestoreDocumentState, bool) {
	return fd.state, fd.state != nil
}

// StateMust returns the state for [FirestoreDocument]. Panics if the state is nil.
func (fd *FirestoreDocument) StateMust() *firestoreDocumentState {
	if fd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fd.Type(), fd.LocalName()))
	}
	return fd.state
}

// FirestoreDocumentArgs contains the configurations for google_firestore_document.
type FirestoreDocumentArgs struct {
	// Collection: string, required
	Collection terra.StringValue `hcl:"collection,attr" validate:"required"`
	// Database: string, optional
	Database terra.StringValue `hcl:"database,attr"`
	// DocumentId: string, required
	DocumentId terra.StringValue `hcl:"document_id,attr" validate:"required"`
	// Fields: string, required
	Fields terra.StringValue `hcl:"fields,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *firestoredocument.Timeouts `hcl:"timeouts,block"`
}
type firestoreDocumentAttributes struct {
	ref terra.Reference
}

// Collection returns a reference to field collection of google_firestore_document.
func (fd firestoreDocumentAttributes) Collection() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("collection"))
}

// CreateTime returns a reference to field create_time of google_firestore_document.
func (fd firestoreDocumentAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("create_time"))
}

// Database returns a reference to field database of google_firestore_document.
func (fd firestoreDocumentAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("database"))
}

// DocumentId returns a reference to field document_id of google_firestore_document.
func (fd firestoreDocumentAttributes) DocumentId() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("document_id"))
}

// Fields returns a reference to field fields of google_firestore_document.
func (fd firestoreDocumentAttributes) Fields() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("fields"))
}

// Id returns a reference to field id of google_firestore_document.
func (fd firestoreDocumentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("id"))
}

// Name returns a reference to field name of google_firestore_document.
func (fd firestoreDocumentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("name"))
}

// Path returns a reference to field path of google_firestore_document.
func (fd firestoreDocumentAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("path"))
}

// Project returns a reference to field project of google_firestore_document.
func (fd firestoreDocumentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("project"))
}

// UpdateTime returns a reference to field update_time of google_firestore_document.
func (fd firestoreDocumentAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("update_time"))
}

func (fd firestoreDocumentAttributes) Timeouts() firestoredocument.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firestoredocument.TimeoutsAttributes](fd.ref.Append("timeouts"))
}

type firestoreDocumentState struct {
	Collection string                           `json:"collection"`
	CreateTime string                           `json:"create_time"`
	Database   string                           `json:"database"`
	DocumentId string                           `json:"document_id"`
	Fields     string                           `json:"fields"`
	Id         string                           `json:"id"`
	Name       string                           `json:"name"`
	Path       string                           `json:"path"`
	Project    string                           `json:"project"`
	UpdateTime string                           `json:"update_time"`
	Timeouts   *firestoredocument.TimeoutsState `json:"timeouts"`
}
