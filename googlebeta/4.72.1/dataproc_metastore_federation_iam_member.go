// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dataprocmetastorefederationiammember "github.com/golingon/terraproviders/googlebeta/4.72.1/dataprocmetastorefederationiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataprocMetastoreFederationIamMember creates a new instance of [DataprocMetastoreFederationIamMember].
func NewDataprocMetastoreFederationIamMember(name string, args DataprocMetastoreFederationIamMemberArgs) *DataprocMetastoreFederationIamMember {
	return &DataprocMetastoreFederationIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataprocMetastoreFederationIamMember)(nil)

// DataprocMetastoreFederationIamMember represents the Terraform resource google_dataproc_metastore_federation_iam_member.
type DataprocMetastoreFederationIamMember struct {
	Name      string
	Args      DataprocMetastoreFederationIamMemberArgs
	state     *dataprocMetastoreFederationIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataprocMetastoreFederationIamMember].
func (dmfim *DataprocMetastoreFederationIamMember) Type() string {
	return "google_dataproc_metastore_federation_iam_member"
}

// LocalName returns the local name for [DataprocMetastoreFederationIamMember].
func (dmfim *DataprocMetastoreFederationIamMember) LocalName() string {
	return dmfim.Name
}

// Configuration returns the configuration (args) for [DataprocMetastoreFederationIamMember].
func (dmfim *DataprocMetastoreFederationIamMember) Configuration() interface{} {
	return dmfim.Args
}

// DependOn is used for other resources to depend on [DataprocMetastoreFederationIamMember].
func (dmfim *DataprocMetastoreFederationIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(dmfim)
}

// Dependencies returns the list of resources [DataprocMetastoreFederationIamMember] depends_on.
func (dmfim *DataprocMetastoreFederationIamMember) Dependencies() terra.Dependencies {
	return dmfim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataprocMetastoreFederationIamMember].
func (dmfim *DataprocMetastoreFederationIamMember) LifecycleManagement() *terra.Lifecycle {
	return dmfim.Lifecycle
}

// Attributes returns the attributes for [DataprocMetastoreFederationIamMember].
func (dmfim *DataprocMetastoreFederationIamMember) Attributes() dataprocMetastoreFederationIamMemberAttributes {
	return dataprocMetastoreFederationIamMemberAttributes{ref: terra.ReferenceResource(dmfim)}
}

// ImportState imports the given attribute values into [DataprocMetastoreFederationIamMember]'s state.
func (dmfim *DataprocMetastoreFederationIamMember) ImportState(av io.Reader) error {
	dmfim.state = &dataprocMetastoreFederationIamMemberState{}
	if err := json.NewDecoder(av).Decode(dmfim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dmfim.Type(), dmfim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataprocMetastoreFederationIamMember] has state.
func (dmfim *DataprocMetastoreFederationIamMember) State() (*dataprocMetastoreFederationIamMemberState, bool) {
	return dmfim.state, dmfim.state != nil
}

// StateMust returns the state for [DataprocMetastoreFederationIamMember]. Panics if the state is nil.
func (dmfim *DataprocMetastoreFederationIamMember) StateMust() *dataprocMetastoreFederationIamMemberState {
	if dmfim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dmfim.Type(), dmfim.LocalName()))
	}
	return dmfim.state
}

// DataprocMetastoreFederationIamMemberArgs contains the configurations for google_dataproc_metastore_federation_iam_member.
type DataprocMetastoreFederationIamMemberArgs struct {
	// FederationId: string, required
	FederationId terra.StringValue `hcl:"federation_id,attr" validate:"required"`
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
	Condition *dataprocmetastorefederationiammember.Condition `hcl:"condition,block"`
}
type dataprocMetastoreFederationIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataproc_metastore_federation_iam_member.
func (dmfim dataprocMetastoreFederationIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dmfim.ref.Append("etag"))
}

// FederationId returns a reference to field federation_id of google_dataproc_metastore_federation_iam_member.
func (dmfim dataprocMetastoreFederationIamMemberAttributes) FederationId() terra.StringValue {
	return terra.ReferenceAsString(dmfim.ref.Append("federation_id"))
}

// Id returns a reference to field id of google_dataproc_metastore_federation_iam_member.
func (dmfim dataprocMetastoreFederationIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmfim.ref.Append("id"))
}

// Location returns a reference to field location of google_dataproc_metastore_federation_iam_member.
func (dmfim dataprocMetastoreFederationIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dmfim.ref.Append("location"))
}

// Member returns a reference to field member of google_dataproc_metastore_federation_iam_member.
func (dmfim dataprocMetastoreFederationIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(dmfim.ref.Append("member"))
}

// Project returns a reference to field project of google_dataproc_metastore_federation_iam_member.
func (dmfim dataprocMetastoreFederationIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dmfim.ref.Append("project"))
}

// Role returns a reference to field role of google_dataproc_metastore_federation_iam_member.
func (dmfim dataprocMetastoreFederationIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dmfim.ref.Append("role"))
}

func (dmfim dataprocMetastoreFederationIamMemberAttributes) Condition() terra.ListValue[dataprocmetastorefederationiammember.ConditionAttributes] {
	return terra.ReferenceAsList[dataprocmetastorefederationiammember.ConditionAttributes](dmfim.ref.Append("condition"))
}

type dataprocMetastoreFederationIamMemberState struct {
	Etag         string                                                `json:"etag"`
	FederationId string                                                `json:"federation_id"`
	Id           string                                                `json:"id"`
	Location     string                                                `json:"location"`
	Member       string                                                `json:"member"`
	Project      string                                                `json:"project"`
	Role         string                                                `json:"role"`
	Condition    []dataprocmetastorefederationiammember.ConditionState `json:"condition"`
}
