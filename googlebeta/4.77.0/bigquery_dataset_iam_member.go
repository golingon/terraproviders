// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	bigquerydatasetiammember "github.com/golingon/terraproviders/googlebeta/4.77.0/bigquerydatasetiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryDatasetIamMember creates a new instance of [BigqueryDatasetIamMember].
func NewBigqueryDatasetIamMember(name string, args BigqueryDatasetIamMemberArgs) *BigqueryDatasetIamMember {
	return &BigqueryDatasetIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryDatasetIamMember)(nil)

// BigqueryDatasetIamMember represents the Terraform resource google_bigquery_dataset_iam_member.
type BigqueryDatasetIamMember struct {
	Name      string
	Args      BigqueryDatasetIamMemberArgs
	state     *bigqueryDatasetIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryDatasetIamMember].
func (bdim *BigqueryDatasetIamMember) Type() string {
	return "google_bigquery_dataset_iam_member"
}

// LocalName returns the local name for [BigqueryDatasetIamMember].
func (bdim *BigqueryDatasetIamMember) LocalName() string {
	return bdim.Name
}

// Configuration returns the configuration (args) for [BigqueryDatasetIamMember].
func (bdim *BigqueryDatasetIamMember) Configuration() interface{} {
	return bdim.Args
}

// DependOn is used for other resources to depend on [BigqueryDatasetIamMember].
func (bdim *BigqueryDatasetIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(bdim)
}

// Dependencies returns the list of resources [BigqueryDatasetIamMember] depends_on.
func (bdim *BigqueryDatasetIamMember) Dependencies() terra.Dependencies {
	return bdim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryDatasetIamMember].
func (bdim *BigqueryDatasetIamMember) LifecycleManagement() *terra.Lifecycle {
	return bdim.Lifecycle
}

// Attributes returns the attributes for [BigqueryDatasetIamMember].
func (bdim *BigqueryDatasetIamMember) Attributes() bigqueryDatasetIamMemberAttributes {
	return bigqueryDatasetIamMemberAttributes{ref: terra.ReferenceResource(bdim)}
}

// ImportState imports the given attribute values into [BigqueryDatasetIamMember]'s state.
func (bdim *BigqueryDatasetIamMember) ImportState(av io.Reader) error {
	bdim.state = &bigqueryDatasetIamMemberState{}
	if err := json.NewDecoder(av).Decode(bdim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bdim.Type(), bdim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryDatasetIamMember] has state.
func (bdim *BigqueryDatasetIamMember) State() (*bigqueryDatasetIamMemberState, bool) {
	return bdim.state, bdim.state != nil
}

// StateMust returns the state for [BigqueryDatasetIamMember]. Panics if the state is nil.
func (bdim *BigqueryDatasetIamMember) StateMust() *bigqueryDatasetIamMemberState {
	if bdim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bdim.Type(), bdim.LocalName()))
	}
	return bdim.state
}

// BigqueryDatasetIamMemberArgs contains the configurations for google_bigquery_dataset_iam_member.
type BigqueryDatasetIamMemberArgs struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *bigquerydatasetiammember.Condition `hcl:"condition,block"`
}
type bigqueryDatasetIamMemberAttributes struct {
	ref terra.Reference
}

// DatasetId returns a reference to field dataset_id of google_bigquery_dataset_iam_member.
func (bdim bigqueryDatasetIamMemberAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(bdim.ref.Append("dataset_id"))
}

// Etag returns a reference to field etag of google_bigquery_dataset_iam_member.
func (bdim bigqueryDatasetIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(bdim.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_dataset_iam_member.
func (bdim bigqueryDatasetIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bdim.ref.Append("id"))
}

// Member returns a reference to field member of google_bigquery_dataset_iam_member.
func (bdim bigqueryDatasetIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(bdim.ref.Append("member"))
}

// Project returns a reference to field project of google_bigquery_dataset_iam_member.
func (bdim bigqueryDatasetIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bdim.ref.Append("project"))
}

// Role returns a reference to field role of google_bigquery_dataset_iam_member.
func (bdim bigqueryDatasetIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(bdim.ref.Append("role"))
}

func (bdim bigqueryDatasetIamMemberAttributes) Condition() terra.ListValue[bigquerydatasetiammember.ConditionAttributes] {
	return terra.ReferenceAsList[bigquerydatasetiammember.ConditionAttributes](bdim.ref.Append("condition"))
}

type bigqueryDatasetIamMemberState struct {
	DatasetId string                                    `json:"dataset_id"`
	Etag      string                                    `json:"etag"`
	Id        string                                    `json:"id"`
	Member    string                                    `json:"member"`
	Project   string                                    `json:"project"`
	Role      string                                    `json:"role"`
	Condition []bigquerydatasetiammember.ConditionState `json:"condition"`
}
