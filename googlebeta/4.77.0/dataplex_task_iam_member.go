// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dataplextaskiammember "github.com/golingon/terraproviders/googlebeta/4.77.0/dataplextaskiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataplexTaskIamMember creates a new instance of [DataplexTaskIamMember].
func NewDataplexTaskIamMember(name string, args DataplexTaskIamMemberArgs) *DataplexTaskIamMember {
	return &DataplexTaskIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataplexTaskIamMember)(nil)

// DataplexTaskIamMember represents the Terraform resource google_dataplex_task_iam_member.
type DataplexTaskIamMember struct {
	Name      string
	Args      DataplexTaskIamMemberArgs
	state     *dataplexTaskIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataplexTaskIamMember].
func (dtim *DataplexTaskIamMember) Type() string {
	return "google_dataplex_task_iam_member"
}

// LocalName returns the local name for [DataplexTaskIamMember].
func (dtim *DataplexTaskIamMember) LocalName() string {
	return dtim.Name
}

// Configuration returns the configuration (args) for [DataplexTaskIamMember].
func (dtim *DataplexTaskIamMember) Configuration() interface{} {
	return dtim.Args
}

// DependOn is used for other resources to depend on [DataplexTaskIamMember].
func (dtim *DataplexTaskIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(dtim)
}

// Dependencies returns the list of resources [DataplexTaskIamMember] depends_on.
func (dtim *DataplexTaskIamMember) Dependencies() terra.Dependencies {
	return dtim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataplexTaskIamMember].
func (dtim *DataplexTaskIamMember) LifecycleManagement() *terra.Lifecycle {
	return dtim.Lifecycle
}

// Attributes returns the attributes for [DataplexTaskIamMember].
func (dtim *DataplexTaskIamMember) Attributes() dataplexTaskIamMemberAttributes {
	return dataplexTaskIamMemberAttributes{ref: terra.ReferenceResource(dtim)}
}

// ImportState imports the given attribute values into [DataplexTaskIamMember]'s state.
func (dtim *DataplexTaskIamMember) ImportState(av io.Reader) error {
	dtim.state = &dataplexTaskIamMemberState{}
	if err := json.NewDecoder(av).Decode(dtim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dtim.Type(), dtim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataplexTaskIamMember] has state.
func (dtim *DataplexTaskIamMember) State() (*dataplexTaskIamMemberState, bool) {
	return dtim.state, dtim.state != nil
}

// StateMust returns the state for [DataplexTaskIamMember]. Panics if the state is nil.
func (dtim *DataplexTaskIamMember) StateMust() *dataplexTaskIamMemberState {
	if dtim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dtim.Type(), dtim.LocalName()))
	}
	return dtim.state
}

// DataplexTaskIamMemberArgs contains the configurations for google_dataplex_task_iam_member.
type DataplexTaskIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Lake: string, required
	Lake terra.StringValue `hcl:"lake,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// TaskId: string, required
	TaskId terra.StringValue `hcl:"task_id,attr" validate:"required"`
	// Condition: optional
	Condition *dataplextaskiammember.Condition `hcl:"condition,block"`
}
type dataplexTaskIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataplex_task_iam_member.
func (dtim dataplexTaskIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dtim.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_task_iam_member.
func (dtim dataplexTaskIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dtim.ref.Append("id"))
}

// Lake returns a reference to field lake of google_dataplex_task_iam_member.
func (dtim dataplexTaskIamMemberAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(dtim.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_task_iam_member.
func (dtim dataplexTaskIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dtim.ref.Append("location"))
}

// Member returns a reference to field member of google_dataplex_task_iam_member.
func (dtim dataplexTaskIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(dtim.ref.Append("member"))
}

// Project returns a reference to field project of google_dataplex_task_iam_member.
func (dtim dataplexTaskIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dtim.ref.Append("project"))
}

// Role returns a reference to field role of google_dataplex_task_iam_member.
func (dtim dataplexTaskIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dtim.ref.Append("role"))
}

// TaskId returns a reference to field task_id of google_dataplex_task_iam_member.
func (dtim dataplexTaskIamMemberAttributes) TaskId() terra.StringValue {
	return terra.ReferenceAsString(dtim.ref.Append("task_id"))
}

func (dtim dataplexTaskIamMemberAttributes) Condition() terra.ListValue[dataplextaskiammember.ConditionAttributes] {
	return terra.ReferenceAsList[dataplextaskiammember.ConditionAttributes](dtim.ref.Append("condition"))
}

type dataplexTaskIamMemberState struct {
	Etag      string                                 `json:"etag"`
	Id        string                                 `json:"id"`
	Lake      string                                 `json:"lake"`
	Location  string                                 `json:"location"`
	Member    string                                 `json:"member"`
	Project   string                                 `json:"project"`
	Role      string                                 `json:"role"`
	TaskId    string                                 `json:"task_id"`
	Condition []dataplextaskiammember.ConditionState `json:"condition"`
}
