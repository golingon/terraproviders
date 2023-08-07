// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dataprocmetastoreserviceiammember "github.com/golingon/terraproviders/googlebeta/4.76.0/dataprocmetastoreserviceiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataprocMetastoreServiceIamMember creates a new instance of [DataprocMetastoreServiceIamMember].
func NewDataprocMetastoreServiceIamMember(name string, args DataprocMetastoreServiceIamMemberArgs) *DataprocMetastoreServiceIamMember {
	return &DataprocMetastoreServiceIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocMetastoreServiceIamMember)(nil)

// DataprocMetastoreServiceIamMember represents the Terraform resource google_dataproc_metastore_service_iam_member.
type DataprocMetastoreServiceIamMember struct {
	Name      string
	Args      DataprocMetastoreServiceIamMemberArgs
	state     *dataprocMetastoreServiceIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocMetastoreServiceIamMember].
func (dmsim *DataprocMetastoreServiceIamMember) Type() string {
	return "google_dataproc_metastore_service_iam_member"
}

// LocalName returns the local name for [DataprocMetastoreServiceIamMember].
func (dmsim *DataprocMetastoreServiceIamMember) LocalName() string {
	return dmsim.Name
}

// Configuration returns the configuration (args) for [DataprocMetastoreServiceIamMember].
func (dmsim *DataprocMetastoreServiceIamMember) Configuration() interface{} {
	return dmsim.Args
}

// DependOn is used for other resources to depend on [DataprocMetastoreServiceIamMember].
func (dmsim *DataprocMetastoreServiceIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(dmsim)
}

// Dependencies returns the list of resources [DataprocMetastoreServiceIamMember] depends_on.
func (dmsim *DataprocMetastoreServiceIamMember) Dependencies() terra.Dependencies {
	return dmsim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocMetastoreServiceIamMember].
func (dmsim *DataprocMetastoreServiceIamMember) LifecycleManagement() *terra.Lifecycle {
	return dmsim.Lifecycle
}

// Attributes returns the attributes for [DataprocMetastoreServiceIamMember].
func (dmsim *DataprocMetastoreServiceIamMember) Attributes() dataprocMetastoreServiceIamMemberAttributes {
	return dataprocMetastoreServiceIamMemberAttributes{ref: terra.ReferenceResource(dmsim)}
}

// ImportState imports the given attribute values into [DataprocMetastoreServiceIamMember]'s state.
func (dmsim *DataprocMetastoreServiceIamMember) ImportState(av io.Reader) error {
	dmsim.state = &dataprocMetastoreServiceIamMemberState{}
	if err := json.NewDecoder(av).Decode(dmsim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dmsim.Type(), dmsim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocMetastoreServiceIamMember] has state.
func (dmsim *DataprocMetastoreServiceIamMember) State() (*dataprocMetastoreServiceIamMemberState, bool) {
	return dmsim.state, dmsim.state != nil
}

// StateMust returns the state for [DataprocMetastoreServiceIamMember]. Panics if the state is nil.
func (dmsim *DataprocMetastoreServiceIamMember) StateMust() *dataprocMetastoreServiceIamMemberState {
	if dmsim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dmsim.Type(), dmsim.LocalName()))
	}
	return dmsim.state
}

// DataprocMetastoreServiceIamMemberArgs contains the configurations for google_dataproc_metastore_service_iam_member.
type DataprocMetastoreServiceIamMemberArgs struct {
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
	// ServiceId: string, required
	ServiceId terra.StringValue `hcl:"service_id,attr" validate:"required"`
	// Condition: optional
	Condition *dataprocmetastoreserviceiammember.Condition `hcl:"condition,block"`
}
type dataprocMetastoreServiceIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataproc_metastore_service_iam_member.
func (dmsim dataprocMetastoreServiceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dmsim.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_metastore_service_iam_member.
func (dmsim dataprocMetastoreServiceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmsim.ref.Append("id"))
}

// Location returns a reference to field location of google_dataproc_metastore_service_iam_member.
func (dmsim dataprocMetastoreServiceIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dmsim.ref.Append("location"))
}

// Member returns a reference to field member of google_dataproc_metastore_service_iam_member.
func (dmsim dataprocMetastoreServiceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(dmsim.ref.Append("member"))
}

// Project returns a reference to field project of google_dataproc_metastore_service_iam_member.
func (dmsim dataprocMetastoreServiceIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dmsim.ref.Append("project"))
}

// Role returns a reference to field role of google_dataproc_metastore_service_iam_member.
func (dmsim dataprocMetastoreServiceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dmsim.ref.Append("role"))
}

// ServiceId returns a reference to field service_id of google_dataproc_metastore_service_iam_member.
func (dmsim dataprocMetastoreServiceIamMemberAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(dmsim.ref.Append("service_id"))
}

func (dmsim dataprocMetastoreServiceIamMemberAttributes) Condition() terra.ListValue[dataprocmetastoreserviceiammember.ConditionAttributes] {
	return terra.ReferenceAsList[dataprocmetastoreserviceiammember.ConditionAttributes](dmsim.ref.Append("condition"))
}

type dataprocMetastoreServiceIamMemberState struct {
	Etag      string                                             `json:"etag"`
	Id        string                                             `json:"id"`
	Location  string                                             `json:"location"`
	Member    string                                             `json:"member"`
	Project   string                                             `json:"project"`
	Role      string                                             `json:"role"`
	ServiceId string                                             `json:"service_id"`
	Condition []dataprocmetastoreserviceiammember.ConditionState `json:"condition"`
}
