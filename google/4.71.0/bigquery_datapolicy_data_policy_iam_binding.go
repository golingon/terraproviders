// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigquerydatapolicydatapolicyiambinding "github.com/golingon/terraproviders/google/4.71.0/bigquerydatapolicydatapolicyiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryDatapolicyDataPolicyIamBinding creates a new instance of [BigqueryDatapolicyDataPolicyIamBinding].
func NewBigqueryDatapolicyDataPolicyIamBinding(name string, args BigqueryDatapolicyDataPolicyIamBindingArgs) *BigqueryDatapolicyDataPolicyIamBinding {
	return &BigqueryDatapolicyDataPolicyIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryDatapolicyDataPolicyIamBinding)(nil)

// BigqueryDatapolicyDataPolicyIamBinding represents the Terraform resource google_bigquery_datapolicy_data_policy_iam_binding.
type BigqueryDatapolicyDataPolicyIamBinding struct {
	Name      string
	Args      BigqueryDatapolicyDataPolicyIamBindingArgs
	state     *bigqueryDatapolicyDataPolicyIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryDatapolicyDataPolicyIamBinding].
func (bddpib *BigqueryDatapolicyDataPolicyIamBinding) Type() string {
	return "google_bigquery_datapolicy_data_policy_iam_binding"
}

// LocalName returns the local name for [BigqueryDatapolicyDataPolicyIamBinding].
func (bddpib *BigqueryDatapolicyDataPolicyIamBinding) LocalName() string {
	return bddpib.Name
}

// Configuration returns the configuration (args) for [BigqueryDatapolicyDataPolicyIamBinding].
func (bddpib *BigqueryDatapolicyDataPolicyIamBinding) Configuration() interface{} {
	return bddpib.Args
}

// DependOn is used for other resources to depend on [BigqueryDatapolicyDataPolicyIamBinding].
func (bddpib *BigqueryDatapolicyDataPolicyIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(bddpib)
}

// Dependencies returns the list of resources [BigqueryDatapolicyDataPolicyIamBinding] depends_on.
func (bddpib *BigqueryDatapolicyDataPolicyIamBinding) Dependencies() terra.Dependencies {
	return bddpib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryDatapolicyDataPolicyIamBinding].
func (bddpib *BigqueryDatapolicyDataPolicyIamBinding) LifecycleManagement() *terra.Lifecycle {
	return bddpib.Lifecycle
}

// Attributes returns the attributes for [BigqueryDatapolicyDataPolicyIamBinding].
func (bddpib *BigqueryDatapolicyDataPolicyIamBinding) Attributes() bigqueryDatapolicyDataPolicyIamBindingAttributes {
	return bigqueryDatapolicyDataPolicyIamBindingAttributes{ref: terra.ReferenceResource(bddpib)}
}

// ImportState imports the given attribute values into [BigqueryDatapolicyDataPolicyIamBinding]'s state.
func (bddpib *BigqueryDatapolicyDataPolicyIamBinding) ImportState(av io.Reader) error {
	bddpib.state = &bigqueryDatapolicyDataPolicyIamBindingState{}
	if err := json.NewDecoder(av).Decode(bddpib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bddpib.Type(), bddpib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryDatapolicyDataPolicyIamBinding] has state.
func (bddpib *BigqueryDatapolicyDataPolicyIamBinding) State() (*bigqueryDatapolicyDataPolicyIamBindingState, bool) {
	return bddpib.state, bddpib.state != nil
}

// StateMust returns the state for [BigqueryDatapolicyDataPolicyIamBinding]. Panics if the state is nil.
func (bddpib *BigqueryDatapolicyDataPolicyIamBinding) StateMust() *bigqueryDatapolicyDataPolicyIamBindingState {
	if bddpib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bddpib.Type(), bddpib.LocalName()))
	}
	return bddpib.state
}

// BigqueryDatapolicyDataPolicyIamBindingArgs contains the configurations for google_bigquery_datapolicy_data_policy_iam_binding.
type BigqueryDatapolicyDataPolicyIamBindingArgs struct {
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
	Condition *bigquerydatapolicydatapolicyiambinding.Condition `hcl:"condition,block"`
}
type bigqueryDatapolicyDataPolicyIamBindingAttributes struct {
	ref terra.Reference
}

// DataPolicyId returns a reference to field data_policy_id of google_bigquery_datapolicy_data_policy_iam_binding.
func (bddpib bigqueryDatapolicyDataPolicyIamBindingAttributes) DataPolicyId() terra.StringValue {
	return terra.ReferenceAsString(bddpib.ref.Append("data_policy_id"))
}

// Etag returns a reference to field etag of google_bigquery_datapolicy_data_policy_iam_binding.
func (bddpib bigqueryDatapolicyDataPolicyIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(bddpib.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_datapolicy_data_policy_iam_binding.
func (bddpib bigqueryDatapolicyDataPolicyIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bddpib.ref.Append("id"))
}

// Location returns a reference to field location of google_bigquery_datapolicy_data_policy_iam_binding.
func (bddpib bigqueryDatapolicyDataPolicyIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bddpib.ref.Append("location"))
}

// Members returns a reference to field members of google_bigquery_datapolicy_data_policy_iam_binding.
func (bddpib bigqueryDatapolicyDataPolicyIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](bddpib.ref.Append("members"))
}

// Project returns a reference to field project of google_bigquery_datapolicy_data_policy_iam_binding.
func (bddpib bigqueryDatapolicyDataPolicyIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bddpib.ref.Append("project"))
}

// Role returns a reference to field role of google_bigquery_datapolicy_data_policy_iam_binding.
func (bddpib bigqueryDatapolicyDataPolicyIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(bddpib.ref.Append("role"))
}

func (bddpib bigqueryDatapolicyDataPolicyIamBindingAttributes) Condition() terra.ListValue[bigquerydatapolicydatapolicyiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[bigquerydatapolicydatapolicyiambinding.ConditionAttributes](bddpib.ref.Append("condition"))
}

type bigqueryDatapolicyDataPolicyIamBindingState struct {
	DataPolicyId string                                                  `json:"data_policy_id"`
	Etag         string                                                  `json:"etag"`
	Id           string                                                  `json:"id"`
	Location     string                                                  `json:"location"`
	Members      []string                                                `json:"members"`
	Project      string                                                  `json:"project"`
	Role         string                                                  `json:"role"`
	Condition    []bigquerydatapolicydatapolicyiambinding.ConditionState `json:"condition"`
}
