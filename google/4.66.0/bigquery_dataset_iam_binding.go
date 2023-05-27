// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigquerydatasetiambinding "github.com/golingon/terraproviders/google/4.66.0/bigquerydatasetiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryDatasetIamBinding creates a new instance of [BigqueryDatasetIamBinding].
func NewBigqueryDatasetIamBinding(name string, args BigqueryDatasetIamBindingArgs) *BigqueryDatasetIamBinding {
	return &BigqueryDatasetIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryDatasetIamBinding)(nil)

// BigqueryDatasetIamBinding represents the Terraform resource google_bigquery_dataset_iam_binding.
type BigqueryDatasetIamBinding struct {
	Name      string
	Args      BigqueryDatasetIamBindingArgs
	state     *bigqueryDatasetIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryDatasetIamBinding].
func (bdib *BigqueryDatasetIamBinding) Type() string {
	return "google_bigquery_dataset_iam_binding"
}

// LocalName returns the local name for [BigqueryDatasetIamBinding].
func (bdib *BigqueryDatasetIamBinding) LocalName() string {
	return bdib.Name
}

// Configuration returns the configuration (args) for [BigqueryDatasetIamBinding].
func (bdib *BigqueryDatasetIamBinding) Configuration() interface{} {
	return bdib.Args
}

// DependOn is used for other resources to depend on [BigqueryDatasetIamBinding].
func (bdib *BigqueryDatasetIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(bdib)
}

// Dependencies returns the list of resources [BigqueryDatasetIamBinding] depends_on.
func (bdib *BigqueryDatasetIamBinding) Dependencies() terra.Dependencies {
	return bdib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryDatasetIamBinding].
func (bdib *BigqueryDatasetIamBinding) LifecycleManagement() *terra.Lifecycle {
	return bdib.Lifecycle
}

// Attributes returns the attributes for [BigqueryDatasetIamBinding].
func (bdib *BigqueryDatasetIamBinding) Attributes() bigqueryDatasetIamBindingAttributes {
	return bigqueryDatasetIamBindingAttributes{ref: terra.ReferenceResource(bdib)}
}

// ImportState imports the given attribute values into [BigqueryDatasetIamBinding]'s state.
func (bdib *BigqueryDatasetIamBinding) ImportState(av io.Reader) error {
	bdib.state = &bigqueryDatasetIamBindingState{}
	if err := json.NewDecoder(av).Decode(bdib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bdib.Type(), bdib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryDatasetIamBinding] has state.
func (bdib *BigqueryDatasetIamBinding) State() (*bigqueryDatasetIamBindingState, bool) {
	return bdib.state, bdib.state != nil
}

// StateMust returns the state for [BigqueryDatasetIamBinding]. Panics if the state is nil.
func (bdib *BigqueryDatasetIamBinding) StateMust() *bigqueryDatasetIamBindingState {
	if bdib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bdib.Type(), bdib.LocalName()))
	}
	return bdib.state
}

// BigqueryDatasetIamBindingArgs contains the configurations for google_bigquery_dataset_iam_binding.
type BigqueryDatasetIamBindingArgs struct {
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
	// Condition: optional
	Condition *bigquerydatasetiambinding.Condition `hcl:"condition,block"`
}
type bigqueryDatasetIamBindingAttributes struct {
	ref terra.Reference
}

// DatasetId returns a reference to field dataset_id of google_bigquery_dataset_iam_binding.
func (bdib bigqueryDatasetIamBindingAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(bdib.ref.Append("dataset_id"))
}

// Etag returns a reference to field etag of google_bigquery_dataset_iam_binding.
func (bdib bigqueryDatasetIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(bdib.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_dataset_iam_binding.
func (bdib bigqueryDatasetIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bdib.ref.Append("id"))
}

// Members returns a reference to field members of google_bigquery_dataset_iam_binding.
func (bdib bigqueryDatasetIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](bdib.ref.Append("members"))
}

// Project returns a reference to field project of google_bigquery_dataset_iam_binding.
func (bdib bigqueryDatasetIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bdib.ref.Append("project"))
}

// Role returns a reference to field role of google_bigquery_dataset_iam_binding.
func (bdib bigqueryDatasetIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(bdib.ref.Append("role"))
}

func (bdib bigqueryDatasetIamBindingAttributes) Condition() terra.ListValue[bigquerydatasetiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[bigquerydatasetiambinding.ConditionAttributes](bdib.ref.Append("condition"))
}

type bigqueryDatasetIamBindingState struct {
	DatasetId string                                     `json:"dataset_id"`
	Etag      string                                     `json:"etag"`
	Id        string                                     `json:"id"`
	Members   []string                                   `json:"members"`
	Project   string                                     `json:"project"`
	Role      string                                     `json:"role"`
	Condition []bigquerydatasetiambinding.ConditionState `json:"condition"`
}
