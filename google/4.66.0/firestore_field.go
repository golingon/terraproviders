// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	firestorefield "github.com/golingon/terraproviders/google/4.66.0/firestorefield"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirestoreField creates a new instance of [FirestoreField].
func NewFirestoreField(name string, args FirestoreFieldArgs) *FirestoreField {
	return &FirestoreField{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirestoreField)(nil)

// FirestoreField represents the Terraform resource google_firestore_field.
type FirestoreField struct {
	Name      string
	Args      FirestoreFieldArgs
	state     *firestoreFieldState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirestoreField].
func (ff *FirestoreField) Type() string {
	return "google_firestore_field"
}

// LocalName returns the local name for [FirestoreField].
func (ff *FirestoreField) LocalName() string {
	return ff.Name
}

// Configuration returns the configuration (args) for [FirestoreField].
func (ff *FirestoreField) Configuration() interface{} {
	return ff.Args
}

// DependOn is used for other resources to depend on [FirestoreField].
func (ff *FirestoreField) DependOn() terra.Reference {
	return terra.ReferenceResource(ff)
}

// Dependencies returns the list of resources [FirestoreField] depends_on.
func (ff *FirestoreField) Dependencies() terra.Dependencies {
	return ff.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirestoreField].
func (ff *FirestoreField) LifecycleManagement() *terra.Lifecycle {
	return ff.Lifecycle
}

// Attributes returns the attributes for [FirestoreField].
func (ff *FirestoreField) Attributes() firestoreFieldAttributes {
	return firestoreFieldAttributes{ref: terra.ReferenceResource(ff)}
}

// ImportState imports the given attribute values into [FirestoreField]'s state.
func (ff *FirestoreField) ImportState(av io.Reader) error {
	ff.state = &firestoreFieldState{}
	if err := json.NewDecoder(av).Decode(ff.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ff.Type(), ff.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirestoreField] has state.
func (ff *FirestoreField) State() (*firestoreFieldState, bool) {
	return ff.state, ff.state != nil
}

// StateMust returns the state for [FirestoreField]. Panics if the state is nil.
func (ff *FirestoreField) StateMust() *firestoreFieldState {
	if ff.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ff.Type(), ff.LocalName()))
	}
	return ff.state
}

// FirestoreFieldArgs contains the configurations for google_firestore_field.
type FirestoreFieldArgs struct {
	// Collection: string, required
	Collection terra.StringValue `hcl:"collection,attr" validate:"required"`
	// Database: string, optional
	Database terra.StringValue `hcl:"database,attr"`
	// Field: string, required
	Field terra.StringValue `hcl:"field,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// IndexConfig: optional
	IndexConfig *firestorefield.IndexConfig `hcl:"index_config,block"`
	// Timeouts: optional
	Timeouts *firestorefield.Timeouts `hcl:"timeouts,block"`
	// TtlConfig: optional
	TtlConfig *firestorefield.TtlConfig `hcl:"ttl_config,block"`
}
type firestoreFieldAttributes struct {
	ref terra.Reference
}

// Collection returns a reference to field collection of google_firestore_field.
func (ff firestoreFieldAttributes) Collection() terra.StringValue {
	return terra.ReferenceAsString(ff.ref.Append("collection"))
}

// Database returns a reference to field database of google_firestore_field.
func (ff firestoreFieldAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(ff.ref.Append("database"))
}

// Field returns a reference to field field of google_firestore_field.
func (ff firestoreFieldAttributes) Field() terra.StringValue {
	return terra.ReferenceAsString(ff.ref.Append("field"))
}

// Id returns a reference to field id of google_firestore_field.
func (ff firestoreFieldAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ff.ref.Append("id"))
}

// Name returns a reference to field name of google_firestore_field.
func (ff firestoreFieldAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ff.ref.Append("name"))
}

// Project returns a reference to field project of google_firestore_field.
func (ff firestoreFieldAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ff.ref.Append("project"))
}

func (ff firestoreFieldAttributes) IndexConfig() terra.ListValue[firestorefield.IndexConfigAttributes] {
	return terra.ReferenceAsList[firestorefield.IndexConfigAttributes](ff.ref.Append("index_config"))
}

func (ff firestoreFieldAttributes) Timeouts() firestorefield.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firestorefield.TimeoutsAttributes](ff.ref.Append("timeouts"))
}

func (ff firestoreFieldAttributes) TtlConfig() terra.ListValue[firestorefield.TtlConfigAttributes] {
	return terra.ReferenceAsList[firestorefield.TtlConfigAttributes](ff.ref.Append("ttl_config"))
}

type firestoreFieldState struct {
	Collection  string                            `json:"collection"`
	Database    string                            `json:"database"`
	Field       string                            `json:"field"`
	Id          string                            `json:"id"`
	Name        string                            `json:"name"`
	Project     string                            `json:"project"`
	IndexConfig []firestorefield.IndexConfigState `json:"index_config"`
	Timeouts    *firestorefield.TimeoutsState     `json:"timeouts"`
	TtlConfig   []firestorefield.TtlConfigState   `json:"ttl_config"`
}
