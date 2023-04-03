// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigquerytableiambinding "github.com/golingon/terraproviders/google/4.59.0/bigquerytableiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryTableIamBinding creates a new instance of [BigqueryTableIamBinding].
func NewBigqueryTableIamBinding(name string, args BigqueryTableIamBindingArgs) *BigqueryTableIamBinding {
	return &BigqueryTableIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryTableIamBinding)(nil)

// BigqueryTableIamBinding represents the Terraform resource google_bigquery_table_iam_binding.
type BigqueryTableIamBinding struct {
	Name      string
	Args      BigqueryTableIamBindingArgs
	state     *bigqueryTableIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryTableIamBinding].
func (btib *BigqueryTableIamBinding) Type() string {
	return "google_bigquery_table_iam_binding"
}

// LocalName returns the local name for [BigqueryTableIamBinding].
func (btib *BigqueryTableIamBinding) LocalName() string {
	return btib.Name
}

// Configuration returns the configuration (args) for [BigqueryTableIamBinding].
func (btib *BigqueryTableIamBinding) Configuration() interface{} {
	return btib.Args
}

// DependOn is used for other resources to depend on [BigqueryTableIamBinding].
func (btib *BigqueryTableIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(btib)
}

// Dependencies returns the list of resources [BigqueryTableIamBinding] depends_on.
func (btib *BigqueryTableIamBinding) Dependencies() terra.Dependencies {
	return btib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryTableIamBinding].
func (btib *BigqueryTableIamBinding) LifecycleManagement() *terra.Lifecycle {
	return btib.Lifecycle
}

// Attributes returns the attributes for [BigqueryTableIamBinding].
func (btib *BigqueryTableIamBinding) Attributes() bigqueryTableIamBindingAttributes {
	return bigqueryTableIamBindingAttributes{ref: terra.ReferenceResource(btib)}
}

// ImportState imports the given attribute values into [BigqueryTableIamBinding]'s state.
func (btib *BigqueryTableIamBinding) ImportState(av io.Reader) error {
	btib.state = &bigqueryTableIamBindingState{}
	if err := json.NewDecoder(av).Decode(btib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", btib.Type(), btib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryTableIamBinding] has state.
func (btib *BigqueryTableIamBinding) State() (*bigqueryTableIamBindingState, bool) {
	return btib.state, btib.state != nil
}

// StateMust returns the state for [BigqueryTableIamBinding]. Panics if the state is nil.
func (btib *BigqueryTableIamBinding) StateMust() *bigqueryTableIamBindingState {
	if btib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", btib.Type(), btib.LocalName()))
	}
	return btib.state
}

// BigqueryTableIamBindingArgs contains the configurations for google_bigquery_table_iam_binding.
type BigqueryTableIamBindingArgs struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// TableId: string, required
	TableId terra.StringValue `hcl:"table_id,attr" validate:"required"`
	// Condition: optional
	Condition *bigquerytableiambinding.Condition `hcl:"condition,block"`
}
type bigqueryTableIamBindingAttributes struct {
	ref terra.Reference
}

// DatasetId returns a reference to field dataset_id of google_bigquery_table_iam_binding.
func (btib bigqueryTableIamBindingAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(btib.ref.Append("dataset_id"))
}

// Etag returns a reference to field etag of google_bigquery_table_iam_binding.
func (btib bigqueryTableIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(btib.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_table_iam_binding.
func (btib bigqueryTableIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(btib.ref.Append("id"))
}

// Members returns a reference to field members of google_bigquery_table_iam_binding.
func (btib bigqueryTableIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](btib.ref.Append("members"))
}

// Project returns a reference to field project of google_bigquery_table_iam_binding.
func (btib bigqueryTableIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(btib.ref.Append("project"))
}

// Role returns a reference to field role of google_bigquery_table_iam_binding.
func (btib bigqueryTableIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(btib.ref.Append("role"))
}

// TableId returns a reference to field table_id of google_bigquery_table_iam_binding.
func (btib bigqueryTableIamBindingAttributes) TableId() terra.StringValue {
	return terra.ReferenceAsString(btib.ref.Append("table_id"))
}

func (btib bigqueryTableIamBindingAttributes) Condition() terra.ListValue[bigquerytableiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[bigquerytableiambinding.ConditionAttributes](btib.ref.Append("condition"))
}

type bigqueryTableIamBindingState struct {
	DatasetId string                                   `json:"dataset_id"`
	Etag      string                                   `json:"etag"`
	Id        string                                   `json:"id"`
	Members   []string                                 `json:"members"`
	Project   string                                   `json:"project"`
	Role      string                                   `json:"role"`
	TableId   string                                   `json:"table_id"`
	Condition []bigquerytableiambinding.ConditionState `json:"condition"`
}
