// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigquerytableiammember "github.com/golingon/terraproviders/google/4.59.0/bigquerytableiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryTableIamMember creates a new instance of [BigqueryTableIamMember].
func NewBigqueryTableIamMember(name string, args BigqueryTableIamMemberArgs) *BigqueryTableIamMember {
	return &BigqueryTableIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryTableIamMember)(nil)

// BigqueryTableIamMember represents the Terraform resource google_bigquery_table_iam_member.
type BigqueryTableIamMember struct {
	Name      string
	Args      BigqueryTableIamMemberArgs
	state     *bigqueryTableIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryTableIamMember].
func (btim *BigqueryTableIamMember) Type() string {
	return "google_bigquery_table_iam_member"
}

// LocalName returns the local name for [BigqueryTableIamMember].
func (btim *BigqueryTableIamMember) LocalName() string {
	return btim.Name
}

// Configuration returns the configuration (args) for [BigqueryTableIamMember].
func (btim *BigqueryTableIamMember) Configuration() interface{} {
	return btim.Args
}

// DependOn is used for other resources to depend on [BigqueryTableIamMember].
func (btim *BigqueryTableIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(btim)
}

// Dependencies returns the list of resources [BigqueryTableIamMember] depends_on.
func (btim *BigqueryTableIamMember) Dependencies() terra.Dependencies {
	return btim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryTableIamMember].
func (btim *BigqueryTableIamMember) LifecycleManagement() *terra.Lifecycle {
	return btim.Lifecycle
}

// Attributes returns the attributes for [BigqueryTableIamMember].
func (btim *BigqueryTableIamMember) Attributes() bigqueryTableIamMemberAttributes {
	return bigqueryTableIamMemberAttributes{ref: terra.ReferenceResource(btim)}
}

// ImportState imports the given attribute values into [BigqueryTableIamMember]'s state.
func (btim *BigqueryTableIamMember) ImportState(av io.Reader) error {
	btim.state = &bigqueryTableIamMemberState{}
	if err := json.NewDecoder(av).Decode(btim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", btim.Type(), btim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryTableIamMember] has state.
func (btim *BigqueryTableIamMember) State() (*bigqueryTableIamMemberState, bool) {
	return btim.state, btim.state != nil
}

// StateMust returns the state for [BigqueryTableIamMember]. Panics if the state is nil.
func (btim *BigqueryTableIamMember) StateMust() *bigqueryTableIamMemberState {
	if btim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", btim.Type(), btim.LocalName()))
	}
	return btim.state
}

// BigqueryTableIamMemberArgs contains the configurations for google_bigquery_table_iam_member.
type BigqueryTableIamMemberArgs struct {
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
	// TableId: string, required
	TableId terra.StringValue `hcl:"table_id,attr" validate:"required"`
	// Condition: optional
	Condition *bigquerytableiammember.Condition `hcl:"condition,block"`
}
type bigqueryTableIamMemberAttributes struct {
	ref terra.Reference
}

// DatasetId returns a reference to field dataset_id of google_bigquery_table_iam_member.
func (btim bigqueryTableIamMemberAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(btim.ref.Append("dataset_id"))
}

// Etag returns a reference to field etag of google_bigquery_table_iam_member.
func (btim bigqueryTableIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(btim.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_table_iam_member.
func (btim bigqueryTableIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(btim.ref.Append("id"))
}

// Member returns a reference to field member of google_bigquery_table_iam_member.
func (btim bigqueryTableIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(btim.ref.Append("member"))
}

// Project returns a reference to field project of google_bigquery_table_iam_member.
func (btim bigqueryTableIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(btim.ref.Append("project"))
}

// Role returns a reference to field role of google_bigquery_table_iam_member.
func (btim bigqueryTableIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(btim.ref.Append("role"))
}

// TableId returns a reference to field table_id of google_bigquery_table_iam_member.
func (btim bigqueryTableIamMemberAttributes) TableId() terra.StringValue {
	return terra.ReferenceAsString(btim.ref.Append("table_id"))
}

func (btim bigqueryTableIamMemberAttributes) Condition() terra.ListValue[bigquerytableiammember.ConditionAttributes] {
	return terra.ReferenceAsList[bigquerytableiammember.ConditionAttributes](btim.ref.Append("condition"))
}

type bigqueryTableIamMemberState struct {
	DatasetId string                                  `json:"dataset_id"`
	Etag      string                                  `json:"etag"`
	Id        string                                  `json:"id"`
	Member    string                                  `json:"member"`
	Project   string                                  `json:"project"`
	Role      string                                  `json:"role"`
	TableId   string                                  `json:"table_id"`
	Condition []bigquerytableiammember.ConditionState `json:"condition"`
}
