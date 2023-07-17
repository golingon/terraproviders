// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigqueryroutine "github.com/golingon/terraproviders/google/4.73.1/bigqueryroutine"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryRoutine creates a new instance of [BigqueryRoutine].
func NewBigqueryRoutine(name string, args BigqueryRoutineArgs) *BigqueryRoutine {
	return &BigqueryRoutine{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryRoutine)(nil)

// BigqueryRoutine represents the Terraform resource google_bigquery_routine.
type BigqueryRoutine struct {
	Name      string
	Args      BigqueryRoutineArgs
	state     *bigqueryRoutineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryRoutine].
func (br *BigqueryRoutine) Type() string {
	return "google_bigquery_routine"
}

// LocalName returns the local name for [BigqueryRoutine].
func (br *BigqueryRoutine) LocalName() string {
	return br.Name
}

// Configuration returns the configuration (args) for [BigqueryRoutine].
func (br *BigqueryRoutine) Configuration() interface{} {
	return br.Args
}

// DependOn is used for other resources to depend on [BigqueryRoutine].
func (br *BigqueryRoutine) DependOn() terra.Reference {
	return terra.ReferenceResource(br)
}

// Dependencies returns the list of resources [BigqueryRoutine] depends_on.
func (br *BigqueryRoutine) Dependencies() terra.Dependencies {
	return br.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryRoutine].
func (br *BigqueryRoutine) LifecycleManagement() *terra.Lifecycle {
	return br.Lifecycle
}

// Attributes returns the attributes for [BigqueryRoutine].
func (br *BigqueryRoutine) Attributes() bigqueryRoutineAttributes {
	return bigqueryRoutineAttributes{ref: terra.ReferenceResource(br)}
}

// ImportState imports the given attribute values into [BigqueryRoutine]'s state.
func (br *BigqueryRoutine) ImportState(av io.Reader) error {
	br.state = &bigqueryRoutineState{}
	if err := json.NewDecoder(av).Decode(br.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", br.Type(), br.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryRoutine] has state.
func (br *BigqueryRoutine) State() (*bigqueryRoutineState, bool) {
	return br.state, br.state != nil
}

// StateMust returns the state for [BigqueryRoutine]. Panics if the state is nil.
func (br *BigqueryRoutine) StateMust() *bigqueryRoutineState {
	if br.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", br.Type(), br.LocalName()))
	}
	return br.state
}

// BigqueryRoutineArgs contains the configurations for google_bigquery_routine.
type BigqueryRoutineArgs struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// DefinitionBody: string, required
	DefinitionBody terra.StringValue `hcl:"definition_body,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DeterminismLevel: string, optional
	DeterminismLevel terra.StringValue `hcl:"determinism_level,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImportedLibraries: list of string, optional
	ImportedLibraries terra.ListValue[terra.StringValue] `hcl:"imported_libraries,attr"`
	// Language: string, optional
	Language terra.StringValue `hcl:"language,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ReturnTableType: string, optional
	ReturnTableType terra.StringValue `hcl:"return_table_type,attr"`
	// ReturnType: string, optional
	ReturnType terra.StringValue `hcl:"return_type,attr"`
	// RoutineId: string, required
	RoutineId terra.StringValue `hcl:"routine_id,attr" validate:"required"`
	// RoutineType: string, optional
	RoutineType terra.StringValue `hcl:"routine_type,attr"`
	// Arguments: min=0
	Arguments []bigqueryroutine.Arguments `hcl:"arguments,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *bigqueryroutine.Timeouts `hcl:"timeouts,block"`
}
type bigqueryRoutineAttributes struct {
	ref terra.Reference
}

// CreationTime returns a reference to field creation_time of google_bigquery_routine.
func (br bigqueryRoutineAttributes) CreationTime() terra.NumberValue {
	return terra.ReferenceAsNumber(br.ref.Append("creation_time"))
}

// DatasetId returns a reference to field dataset_id of google_bigquery_routine.
func (br bigqueryRoutineAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("dataset_id"))
}

// DefinitionBody returns a reference to field definition_body of google_bigquery_routine.
func (br bigqueryRoutineAttributes) DefinitionBody() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("definition_body"))
}

// Description returns a reference to field description of google_bigquery_routine.
func (br bigqueryRoutineAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("description"))
}

// DeterminismLevel returns a reference to field determinism_level of google_bigquery_routine.
func (br bigqueryRoutineAttributes) DeterminismLevel() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("determinism_level"))
}

// Id returns a reference to field id of google_bigquery_routine.
func (br bigqueryRoutineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("id"))
}

// ImportedLibraries returns a reference to field imported_libraries of google_bigquery_routine.
func (br bigqueryRoutineAttributes) ImportedLibraries() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](br.ref.Append("imported_libraries"))
}

// Language returns a reference to field language of google_bigquery_routine.
func (br bigqueryRoutineAttributes) Language() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("language"))
}

// LastModifiedTime returns a reference to field last_modified_time of google_bigquery_routine.
func (br bigqueryRoutineAttributes) LastModifiedTime() terra.NumberValue {
	return terra.ReferenceAsNumber(br.ref.Append("last_modified_time"))
}

// Project returns a reference to field project of google_bigquery_routine.
func (br bigqueryRoutineAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("project"))
}

// ReturnTableType returns a reference to field return_table_type of google_bigquery_routine.
func (br bigqueryRoutineAttributes) ReturnTableType() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("return_table_type"))
}

// ReturnType returns a reference to field return_type of google_bigquery_routine.
func (br bigqueryRoutineAttributes) ReturnType() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("return_type"))
}

// RoutineId returns a reference to field routine_id of google_bigquery_routine.
func (br bigqueryRoutineAttributes) RoutineId() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("routine_id"))
}

// RoutineType returns a reference to field routine_type of google_bigquery_routine.
func (br bigqueryRoutineAttributes) RoutineType() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("routine_type"))
}

func (br bigqueryRoutineAttributes) Arguments() terra.ListValue[bigqueryroutine.ArgumentsAttributes] {
	return terra.ReferenceAsList[bigqueryroutine.ArgumentsAttributes](br.ref.Append("arguments"))
}

func (br bigqueryRoutineAttributes) Timeouts() bigqueryroutine.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bigqueryroutine.TimeoutsAttributes](br.ref.Append("timeouts"))
}

type bigqueryRoutineState struct {
	CreationTime      float64                          `json:"creation_time"`
	DatasetId         string                           `json:"dataset_id"`
	DefinitionBody    string                           `json:"definition_body"`
	Description       string                           `json:"description"`
	DeterminismLevel  string                           `json:"determinism_level"`
	Id                string                           `json:"id"`
	ImportedLibraries []string                         `json:"imported_libraries"`
	Language          string                           `json:"language"`
	LastModifiedTime  float64                          `json:"last_modified_time"`
	Project           string                           `json:"project"`
	ReturnTableType   string                           `json:"return_table_type"`
	ReturnType        string                           `json:"return_type"`
	RoutineId         string                           `json:"routine_id"`
	RoutineType       string                           `json:"routine_type"`
	Arguments         []bigqueryroutine.ArgumentsState `json:"arguments"`
	Timeouts          *bigqueryroutine.TimeoutsState   `json:"timeouts"`
}
