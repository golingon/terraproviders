// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigquerydatapolicydatapolicyiammember "github.com/golingon/terraproviders/google/4.72.1/bigquerydatapolicydatapolicyiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryDatapolicyDataPolicyIamMember creates a new instance of [BigqueryDatapolicyDataPolicyIamMember].
func NewBigqueryDatapolicyDataPolicyIamMember(name string, args BigqueryDatapolicyDataPolicyIamMemberArgs) *BigqueryDatapolicyDataPolicyIamMember {
	return &BigqueryDatapolicyDataPolicyIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryDatapolicyDataPolicyIamMember)(nil)

// BigqueryDatapolicyDataPolicyIamMember represents the Terraform resource google_bigquery_datapolicy_data_policy_iam_member.
type BigqueryDatapolicyDataPolicyIamMember struct {
	Name      string
	Args      BigqueryDatapolicyDataPolicyIamMemberArgs
	state     *bigqueryDatapolicyDataPolicyIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryDatapolicyDataPolicyIamMember].
func (bddpim *BigqueryDatapolicyDataPolicyIamMember) Type() string {
	return "google_bigquery_datapolicy_data_policy_iam_member"
}

// LocalName returns the local name for [BigqueryDatapolicyDataPolicyIamMember].
func (bddpim *BigqueryDatapolicyDataPolicyIamMember) LocalName() string {
	return bddpim.Name
}

// Configuration returns the configuration (args) for [BigqueryDatapolicyDataPolicyIamMember].
func (bddpim *BigqueryDatapolicyDataPolicyIamMember) Configuration() interface{} {
	return bddpim.Args
}

// DependOn is used for other resources to depend on [BigqueryDatapolicyDataPolicyIamMember].
func (bddpim *BigqueryDatapolicyDataPolicyIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(bddpim)
}

// Dependencies returns the list of resources [BigqueryDatapolicyDataPolicyIamMember] depends_on.
func (bddpim *BigqueryDatapolicyDataPolicyIamMember) Dependencies() terra.Dependencies {
	return bddpim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryDatapolicyDataPolicyIamMember].
func (bddpim *BigqueryDatapolicyDataPolicyIamMember) LifecycleManagement() *terra.Lifecycle {
	return bddpim.Lifecycle
}

// Attributes returns the attributes for [BigqueryDatapolicyDataPolicyIamMember].
func (bddpim *BigqueryDatapolicyDataPolicyIamMember) Attributes() bigqueryDatapolicyDataPolicyIamMemberAttributes {
	return bigqueryDatapolicyDataPolicyIamMemberAttributes{ref: terra.ReferenceResource(bddpim)}
}

// ImportState imports the given attribute values into [BigqueryDatapolicyDataPolicyIamMember]'s state.
func (bddpim *BigqueryDatapolicyDataPolicyIamMember) ImportState(av io.Reader) error {
	bddpim.state = &bigqueryDatapolicyDataPolicyIamMemberState{}
	if err := json.NewDecoder(av).Decode(bddpim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bddpim.Type(), bddpim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryDatapolicyDataPolicyIamMember] has state.
func (bddpim *BigqueryDatapolicyDataPolicyIamMember) State() (*bigqueryDatapolicyDataPolicyIamMemberState, bool) {
	return bddpim.state, bddpim.state != nil
}

// StateMust returns the state for [BigqueryDatapolicyDataPolicyIamMember]. Panics if the state is nil.
func (bddpim *BigqueryDatapolicyDataPolicyIamMember) StateMust() *bigqueryDatapolicyDataPolicyIamMemberState {
	if bddpim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bddpim.Type(), bddpim.LocalName()))
	}
	return bddpim.state
}

// BigqueryDatapolicyDataPolicyIamMemberArgs contains the configurations for google_bigquery_datapolicy_data_policy_iam_member.
type BigqueryDatapolicyDataPolicyIamMemberArgs struct {
	// DataPolicyId: string, required
	DataPolicyId terra.StringValue `hcl:"data_policy_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *bigquerydatapolicydatapolicyiammember.Condition `hcl:"condition,block"`
}
type bigqueryDatapolicyDataPolicyIamMemberAttributes struct {
	ref terra.Reference
}

// DataPolicyId returns a reference to field data_policy_id of google_bigquery_datapolicy_data_policy_iam_member.
func (bddpim bigqueryDatapolicyDataPolicyIamMemberAttributes) DataPolicyId() terra.StringValue {
	return terra.ReferenceAsString(bddpim.ref.Append("data_policy_id"))
}

// Etag returns a reference to field etag of google_bigquery_datapolicy_data_policy_iam_member.
func (bddpim bigqueryDatapolicyDataPolicyIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(bddpim.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_datapolicy_data_policy_iam_member.
func (bddpim bigqueryDatapolicyDataPolicyIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bddpim.ref.Append("id"))
}

// Location returns a reference to field location of google_bigquery_datapolicy_data_policy_iam_member.
func (bddpim bigqueryDatapolicyDataPolicyIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bddpim.ref.Append("location"))
}

// Member returns a reference to field member of google_bigquery_datapolicy_data_policy_iam_member.
func (bddpim bigqueryDatapolicyDataPolicyIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(bddpim.ref.Append("member"))
}

// Project returns a reference to field project of google_bigquery_datapolicy_data_policy_iam_member.
func (bddpim bigqueryDatapolicyDataPolicyIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bddpim.ref.Append("project"))
}

// Role returns a reference to field role of google_bigquery_datapolicy_data_policy_iam_member.
func (bddpim bigqueryDatapolicyDataPolicyIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(bddpim.ref.Append("role"))
}

func (bddpim bigqueryDatapolicyDataPolicyIamMemberAttributes) Condition() terra.ListValue[bigquerydatapolicydatapolicyiammember.ConditionAttributes] {
	return terra.ReferenceAsList[bigquerydatapolicydatapolicyiammember.ConditionAttributes](bddpim.ref.Append("condition"))
}

type bigqueryDatapolicyDataPolicyIamMemberState struct {
	DataPolicyId string                                                 `json:"data_policy_id"`
	Etag         string                                                 `json:"etag"`
	Id           string                                                 `json:"id"`
	Location     string                                                 `json:"location"`
	Member       string                                                 `json:"member"`
	Project      string                                                 `json:"project"`
	Role         string                                                 `json:"role"`
	Condition    []bigquerydatapolicydatapolicyiammember.ConditionState `json:"condition"`
}
