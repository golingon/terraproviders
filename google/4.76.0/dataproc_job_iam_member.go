// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dataprocjobiammember "github.com/golingon/terraproviders/google/4.76.0/dataprocjobiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataprocJobIamMember creates a new instance of [DataprocJobIamMember].
func NewDataprocJobIamMember(name string, args DataprocJobIamMemberArgs) *DataprocJobIamMember {
	return &DataprocJobIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocJobIamMember)(nil)

// DataprocJobIamMember represents the Terraform resource google_dataproc_job_iam_member.
type DataprocJobIamMember struct {
	Name      string
	Args      DataprocJobIamMemberArgs
	state     *dataprocJobIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocJobIamMember].
func (djim *DataprocJobIamMember) Type() string {
	return "google_dataproc_job_iam_member"
}

// LocalName returns the local name for [DataprocJobIamMember].
func (djim *DataprocJobIamMember) LocalName() string {
	return djim.Name
}

// Configuration returns the configuration (args) for [DataprocJobIamMember].
func (djim *DataprocJobIamMember) Configuration() interface{} {
	return djim.Args
}

// DependOn is used for other resources to depend on [DataprocJobIamMember].
func (djim *DataprocJobIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(djim)
}

// Dependencies returns the list of resources [DataprocJobIamMember] depends_on.
func (djim *DataprocJobIamMember) Dependencies() terra.Dependencies {
	return djim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocJobIamMember].
func (djim *DataprocJobIamMember) LifecycleManagement() *terra.Lifecycle {
	return djim.Lifecycle
}

// Attributes returns the attributes for [DataprocJobIamMember].
func (djim *DataprocJobIamMember) Attributes() dataprocJobIamMemberAttributes {
	return dataprocJobIamMemberAttributes{ref: terra.ReferenceResource(djim)}
}

// ImportState imports the given attribute values into [DataprocJobIamMember]'s state.
func (djim *DataprocJobIamMember) ImportState(av io.Reader) error {
	djim.state = &dataprocJobIamMemberState{}
	if err := json.NewDecoder(av).Decode(djim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", djim.Type(), djim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocJobIamMember] has state.
func (djim *DataprocJobIamMember) State() (*dataprocJobIamMemberState, bool) {
	return djim.state, djim.state != nil
}

// StateMust returns the state for [DataprocJobIamMember]. Panics if the state is nil.
func (djim *DataprocJobIamMember) StateMust() *dataprocJobIamMemberState {
	if djim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", djim.Type(), djim.LocalName()))
	}
	return djim.state
}

// DataprocJobIamMemberArgs contains the configurations for google_dataproc_job_iam_member.
type DataprocJobIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// JobId: string, required
	JobId terra.StringValue `hcl:"job_id,attr" validate:"required"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *dataprocjobiammember.Condition `hcl:"condition,block"`
}
type dataprocJobIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataproc_job_iam_member.
func (djim dataprocJobIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(djim.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_job_iam_member.
func (djim dataprocJobIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(djim.ref.Append("id"))
}

// JobId returns a reference to field job_id of google_dataproc_job_iam_member.
func (djim dataprocJobIamMemberAttributes) JobId() terra.StringValue {
	return terra.ReferenceAsString(djim.ref.Append("job_id"))
}

// Member returns a reference to field member of google_dataproc_job_iam_member.
func (djim dataprocJobIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(djim.ref.Append("member"))
}

// Project returns a reference to field project of google_dataproc_job_iam_member.
func (djim dataprocJobIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(djim.ref.Append("project"))
}

// Region returns a reference to field region of google_dataproc_job_iam_member.
func (djim dataprocJobIamMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(djim.ref.Append("region"))
}

// Role returns a reference to field role of google_dataproc_job_iam_member.
func (djim dataprocJobIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(djim.ref.Append("role"))
}

func (djim dataprocJobIamMemberAttributes) Condition() terra.ListValue[dataprocjobiammember.ConditionAttributes] {
	return terra.ReferenceAsList[dataprocjobiammember.ConditionAttributes](djim.ref.Append("condition"))
}

type dataprocJobIamMemberState struct {
	Etag      string                                `json:"etag"`
	Id        string                                `json:"id"`
	JobId     string                                `json:"job_id"`
	Member    string                                `json:"member"`
	Project   string                                `json:"project"`
	Region    string                                `json:"region"`
	Role      string                                `json:"role"`
	Condition []dataprocjobiammember.ConditionState `json:"condition"`
}
