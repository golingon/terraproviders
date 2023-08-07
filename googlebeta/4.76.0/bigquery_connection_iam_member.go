// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	bigqueryconnectioniammember "github.com/golingon/terraproviders/googlebeta/4.76.0/bigqueryconnectioniammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryConnectionIamMember creates a new instance of [BigqueryConnectionIamMember].
func NewBigqueryConnectionIamMember(name string, args BigqueryConnectionIamMemberArgs) *BigqueryConnectionIamMember {
	return &BigqueryConnectionIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryConnectionIamMember)(nil)

// BigqueryConnectionIamMember represents the Terraform resource google_bigquery_connection_iam_member.
type BigqueryConnectionIamMember struct {
	Name      string
	Args      BigqueryConnectionIamMemberArgs
	state     *bigqueryConnectionIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryConnectionIamMember].
func (bcim *BigqueryConnectionIamMember) Type() string {
	return "google_bigquery_connection_iam_member"
}

// LocalName returns the local name for [BigqueryConnectionIamMember].
func (bcim *BigqueryConnectionIamMember) LocalName() string {
	return bcim.Name
}

// Configuration returns the configuration (args) for [BigqueryConnectionIamMember].
func (bcim *BigqueryConnectionIamMember) Configuration() interface{} {
	return bcim.Args
}

// DependOn is used for other resources to depend on [BigqueryConnectionIamMember].
func (bcim *BigqueryConnectionIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(bcim)
}

// Dependencies returns the list of resources [BigqueryConnectionIamMember] depends_on.
func (bcim *BigqueryConnectionIamMember) Dependencies() terra.Dependencies {
	return bcim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryConnectionIamMember].
func (bcim *BigqueryConnectionIamMember) LifecycleManagement() *terra.Lifecycle {
	return bcim.Lifecycle
}

// Attributes returns the attributes for [BigqueryConnectionIamMember].
func (bcim *BigqueryConnectionIamMember) Attributes() bigqueryConnectionIamMemberAttributes {
	return bigqueryConnectionIamMemberAttributes{ref: terra.ReferenceResource(bcim)}
}

// ImportState imports the given attribute values into [BigqueryConnectionIamMember]'s state.
func (bcim *BigqueryConnectionIamMember) ImportState(av io.Reader) error {
	bcim.state = &bigqueryConnectionIamMemberState{}
	if err := json.NewDecoder(av).Decode(bcim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bcim.Type(), bcim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryConnectionIamMember] has state.
func (bcim *BigqueryConnectionIamMember) State() (*bigqueryConnectionIamMemberState, bool) {
	return bcim.state, bcim.state != nil
}

// StateMust returns the state for [BigqueryConnectionIamMember]. Panics if the state is nil.
func (bcim *BigqueryConnectionIamMember) StateMust() *bigqueryConnectionIamMemberState {
	if bcim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bcim.Type(), bcim.LocalName()))
	}
	return bcim.state
}

// BigqueryConnectionIamMemberArgs contains the configurations for google_bigquery_connection_iam_member.
type BigqueryConnectionIamMemberArgs struct {
	// ConnectionId: string, required
	ConnectionId terra.StringValue `hcl:"connection_id,attr" validate:"required"`
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
	Condition *bigqueryconnectioniammember.Condition `hcl:"condition,block"`
}
type bigqueryConnectionIamMemberAttributes struct {
	ref terra.Reference
}

// ConnectionId returns a reference to field connection_id of google_bigquery_connection_iam_member.
func (bcim bigqueryConnectionIamMemberAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(bcim.ref.Append("connection_id"))
}

// Etag returns a reference to field etag of google_bigquery_connection_iam_member.
func (bcim bigqueryConnectionIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(bcim.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_connection_iam_member.
func (bcim bigqueryConnectionIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bcim.ref.Append("id"))
}

// Location returns a reference to field location of google_bigquery_connection_iam_member.
func (bcim bigqueryConnectionIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bcim.ref.Append("location"))
}

// Member returns a reference to field member of google_bigquery_connection_iam_member.
func (bcim bigqueryConnectionIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(bcim.ref.Append("member"))
}

// Project returns a reference to field project of google_bigquery_connection_iam_member.
func (bcim bigqueryConnectionIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bcim.ref.Append("project"))
}

// Role returns a reference to field role of google_bigquery_connection_iam_member.
func (bcim bigqueryConnectionIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(bcim.ref.Append("role"))
}

func (bcim bigqueryConnectionIamMemberAttributes) Condition() terra.ListValue[bigqueryconnectioniammember.ConditionAttributes] {
	return terra.ReferenceAsList[bigqueryconnectioniammember.ConditionAttributes](bcim.ref.Append("condition"))
}

type bigqueryConnectionIamMemberState struct {
	ConnectionId string                                       `json:"connection_id"`
	Etag         string                                       `json:"etag"`
	Id           string                                       `json:"id"`
	Location     string                                       `json:"location"`
	Member       string                                       `json:"member"`
	Project      string                                       `json:"project"`
	Role         string                                       `json:"role"`
	Condition    []bigqueryconnectioniammember.ConditionState `json:"condition"`
}
