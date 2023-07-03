// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dataprocclusteriammember "github.com/golingon/terraproviders/googlebeta/4.71.0/dataprocclusteriammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataprocClusterIamMember creates a new instance of [DataprocClusterIamMember].
func NewDataprocClusterIamMember(name string, args DataprocClusterIamMemberArgs) *DataprocClusterIamMember {
	return &DataprocClusterIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocClusterIamMember)(nil)

// DataprocClusterIamMember represents the Terraform resource google_dataproc_cluster_iam_member.
type DataprocClusterIamMember struct {
	Name      string
	Args      DataprocClusterIamMemberArgs
	state     *dataprocClusterIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocClusterIamMember].
func (dcim *DataprocClusterIamMember) Type() string {
	return "google_dataproc_cluster_iam_member"
}

// LocalName returns the local name for [DataprocClusterIamMember].
func (dcim *DataprocClusterIamMember) LocalName() string {
	return dcim.Name
}

// Configuration returns the configuration (args) for [DataprocClusterIamMember].
func (dcim *DataprocClusterIamMember) Configuration() interface{} {
	return dcim.Args
}

// DependOn is used for other resources to depend on [DataprocClusterIamMember].
func (dcim *DataprocClusterIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(dcim)
}

// Dependencies returns the list of resources [DataprocClusterIamMember] depends_on.
func (dcim *DataprocClusterIamMember) Dependencies() terra.Dependencies {
	return dcim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocClusterIamMember].
func (dcim *DataprocClusterIamMember) LifecycleManagement() *terra.Lifecycle {
	return dcim.Lifecycle
}

// Attributes returns the attributes for [DataprocClusterIamMember].
func (dcim *DataprocClusterIamMember) Attributes() dataprocClusterIamMemberAttributes {
	return dataprocClusterIamMemberAttributes{ref: terra.ReferenceResource(dcim)}
}

// ImportState imports the given attribute values into [DataprocClusterIamMember]'s state.
func (dcim *DataprocClusterIamMember) ImportState(av io.Reader) error {
	dcim.state = &dataprocClusterIamMemberState{}
	if err := json.NewDecoder(av).Decode(dcim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcim.Type(), dcim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocClusterIamMember] has state.
func (dcim *DataprocClusterIamMember) State() (*dataprocClusterIamMemberState, bool) {
	return dcim.state, dcim.state != nil
}

// StateMust returns the state for [DataprocClusterIamMember]. Panics if the state is nil.
func (dcim *DataprocClusterIamMember) StateMust() *dataprocClusterIamMemberState {
	if dcim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcim.Type(), dcim.LocalName()))
	}
	return dcim.state
}

// DataprocClusterIamMemberArgs contains the configurations for google_dataproc_cluster_iam_member.
type DataprocClusterIamMemberArgs struct {
	// Cluster: string, required
	Cluster terra.StringValue `hcl:"cluster,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *dataprocclusteriammember.Condition `hcl:"condition,block"`
}
type dataprocClusterIamMemberAttributes struct {
	ref terra.Reference
}

// Cluster returns a reference to field cluster of google_dataproc_cluster_iam_member.
func (dcim dataprocClusterIamMemberAttributes) Cluster() terra.StringValue {
	return terra.ReferenceAsString(dcim.ref.Append("cluster"))
}

// Etag returns a reference to field etag of google_dataproc_cluster_iam_member.
func (dcim dataprocClusterIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dcim.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_cluster_iam_member.
func (dcim dataprocClusterIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcim.ref.Append("id"))
}

// Member returns a reference to field member of google_dataproc_cluster_iam_member.
func (dcim dataprocClusterIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(dcim.ref.Append("member"))
}

// Project returns a reference to field project of google_dataproc_cluster_iam_member.
func (dcim dataprocClusterIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dcim.ref.Append("project"))
}

// Region returns a reference to field region of google_dataproc_cluster_iam_member.
func (dcim dataprocClusterIamMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dcim.ref.Append("region"))
}

// Role returns a reference to field role of google_dataproc_cluster_iam_member.
func (dcim dataprocClusterIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dcim.ref.Append("role"))
}

func (dcim dataprocClusterIamMemberAttributes) Condition() terra.ListValue[dataprocclusteriammember.ConditionAttributes] {
	return terra.ReferenceAsList[dataprocclusteriammember.ConditionAttributes](dcim.ref.Append("condition"))
}

type dataprocClusterIamMemberState struct {
	Cluster   string                                    `json:"cluster"`
	Etag      string                                    `json:"etag"`
	Id        string                                    `json:"id"`
	Member    string                                    `json:"member"`
	Project   string                                    `json:"project"`
	Region    string                                    `json:"region"`
	Role      string                                    `json:"role"`
	Condition []dataprocclusteriammember.ConditionState `json:"condition"`
}
