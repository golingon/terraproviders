// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_bigquery_datapolicy_data_policy_iam_binding

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

// Resource represents the Terraform resource google_bigquery_datapolicy_data_policy_iam_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleBigqueryDatapolicyDataPolicyIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gbddpib *Resource) Type() string {
	return "google_bigquery_datapolicy_data_policy_iam_binding"
}

// LocalName returns the local name for [Resource].
func (gbddpib *Resource) LocalName() string {
	return gbddpib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gbddpib *Resource) Configuration() interface{} {
	return gbddpib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gbddpib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gbddpib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gbddpib *Resource) Dependencies() terra.Dependencies {
	return gbddpib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gbddpib *Resource) LifecycleManagement() *terra.Lifecycle {
	return gbddpib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gbddpib *Resource) Attributes() googleBigqueryDatapolicyDataPolicyIamBindingAttributes {
	return googleBigqueryDatapolicyDataPolicyIamBindingAttributes{ref: terra.ReferenceResource(gbddpib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gbddpib *Resource) ImportState(state io.Reader) error {
	gbddpib.state = &googleBigqueryDatapolicyDataPolicyIamBindingState{}
	if err := json.NewDecoder(state).Decode(gbddpib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gbddpib.Type(), gbddpib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gbddpib *Resource) State() (*googleBigqueryDatapolicyDataPolicyIamBindingState, bool) {
	return gbddpib.state, gbddpib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gbddpib *Resource) StateMust() *googleBigqueryDatapolicyDataPolicyIamBindingState {
	if gbddpib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gbddpib.Type(), gbddpib.LocalName()))
	}
	return gbddpib.state
}

// Args contains the configurations for google_bigquery_datapolicy_data_policy_iam_binding.
type Args struct {
	// DataPolicyId: string, required
	DataPolicyId terra.StringValue `hcl:"data_policy_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleBigqueryDatapolicyDataPolicyIamBindingAttributes struct {
	ref terra.Reference
}

// DataPolicyId returns a reference to field data_policy_id of google_bigquery_datapolicy_data_policy_iam_binding.
func (gbddpib googleBigqueryDatapolicyDataPolicyIamBindingAttributes) DataPolicyId() terra.StringValue {
	return terra.ReferenceAsString(gbddpib.ref.Append("data_policy_id"))
}

// Etag returns a reference to field etag of google_bigquery_datapolicy_data_policy_iam_binding.
func (gbddpib googleBigqueryDatapolicyDataPolicyIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gbddpib.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_datapolicy_data_policy_iam_binding.
func (gbddpib googleBigqueryDatapolicyDataPolicyIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gbddpib.ref.Append("id"))
}

// Location returns a reference to field location of google_bigquery_datapolicy_data_policy_iam_binding.
func (gbddpib googleBigqueryDatapolicyDataPolicyIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gbddpib.ref.Append("location"))
}

// Members returns a reference to field members of google_bigquery_datapolicy_data_policy_iam_binding.
func (gbddpib googleBigqueryDatapolicyDataPolicyIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gbddpib.ref.Append("members"))
}

// Project returns a reference to field project of google_bigquery_datapolicy_data_policy_iam_binding.
func (gbddpib googleBigqueryDatapolicyDataPolicyIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gbddpib.ref.Append("project"))
}

// Role returns a reference to field role of google_bigquery_datapolicy_data_policy_iam_binding.
func (gbddpib googleBigqueryDatapolicyDataPolicyIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gbddpib.ref.Append("role"))
}

func (gbddpib googleBigqueryDatapolicyDataPolicyIamBindingAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gbddpib.ref.Append("condition"))
}

type googleBigqueryDatapolicyDataPolicyIamBindingState struct {
	DataPolicyId string           `json:"data_policy_id"`
	Etag         string           `json:"etag"`
	Id           string           `json:"id"`
	Location     string           `json:"location"`
	Members      []string         `json:"members"`
	Project      string           `json:"project"`
	Role         string           `json:"role"`
	Condition    []ConditionState `json:"condition"`
}
