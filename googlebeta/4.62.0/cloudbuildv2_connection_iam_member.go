// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	cloudbuildv2connectioniammember "github.com/golingon/terraproviders/googlebeta/4.62.0/cloudbuildv2connectioniammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudbuildv2ConnectionIamMember creates a new instance of [Cloudbuildv2ConnectionIamMember].
func NewCloudbuildv2ConnectionIamMember(name string, args Cloudbuildv2ConnectionIamMemberArgs) *Cloudbuildv2ConnectionIamMember {
	return &Cloudbuildv2ConnectionIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Cloudbuildv2ConnectionIamMember)(nil)

// Cloudbuildv2ConnectionIamMember represents the Terraform resource google_cloudbuildv2_connection_iam_member.
type Cloudbuildv2ConnectionIamMember struct {
	Name      string
	Args      Cloudbuildv2ConnectionIamMemberArgs
	state     *cloudbuildv2ConnectionIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Cloudbuildv2ConnectionIamMember].
func (ccim *Cloudbuildv2ConnectionIamMember) Type() string {
	return "google_cloudbuildv2_connection_iam_member"
}

// LocalName returns the local name for [Cloudbuildv2ConnectionIamMember].
func (ccim *Cloudbuildv2ConnectionIamMember) LocalName() string {
	return ccim.Name
}

// Configuration returns the configuration (args) for [Cloudbuildv2ConnectionIamMember].
func (ccim *Cloudbuildv2ConnectionIamMember) Configuration() interface{} {
	return ccim.Args
}

// DependOn is used for other resources to depend on [Cloudbuildv2ConnectionIamMember].
func (ccim *Cloudbuildv2ConnectionIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(ccim)
}

// Dependencies returns the list of resources [Cloudbuildv2ConnectionIamMember] depends_on.
func (ccim *Cloudbuildv2ConnectionIamMember) Dependencies() terra.Dependencies {
	return ccim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Cloudbuildv2ConnectionIamMember].
func (ccim *Cloudbuildv2ConnectionIamMember) LifecycleManagement() *terra.Lifecycle {
	return ccim.Lifecycle
}

// Attributes returns the attributes for [Cloudbuildv2ConnectionIamMember].
func (ccim *Cloudbuildv2ConnectionIamMember) Attributes() cloudbuildv2ConnectionIamMemberAttributes {
	return cloudbuildv2ConnectionIamMemberAttributes{ref: terra.ReferenceResource(ccim)}
}

// ImportState imports the given attribute values into [Cloudbuildv2ConnectionIamMember]'s state.
func (ccim *Cloudbuildv2ConnectionIamMember) ImportState(av io.Reader) error {
	ccim.state = &cloudbuildv2ConnectionIamMemberState{}
	if err := json.NewDecoder(av).Decode(ccim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ccim.Type(), ccim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Cloudbuildv2ConnectionIamMember] has state.
func (ccim *Cloudbuildv2ConnectionIamMember) State() (*cloudbuildv2ConnectionIamMemberState, bool) {
	return ccim.state, ccim.state != nil
}

// StateMust returns the state for [Cloudbuildv2ConnectionIamMember]. Panics if the state is nil.
func (ccim *Cloudbuildv2ConnectionIamMember) StateMust() *cloudbuildv2ConnectionIamMemberState {
	if ccim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ccim.Type(), ccim.LocalName()))
	}
	return ccim.state
}

// Cloudbuildv2ConnectionIamMemberArgs contains the configurations for google_cloudbuildv2_connection_iam_member.
type Cloudbuildv2ConnectionIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *cloudbuildv2connectioniammember.Condition `hcl:"condition,block"`
}
type cloudbuildv2ConnectionIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloudbuildv2_connection_iam_member.
func (ccim cloudbuildv2ConnectionIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ccim.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloudbuildv2_connection_iam_member.
func (ccim cloudbuildv2ConnectionIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccim.ref.Append("id"))
}

// Location returns a reference to field location of google_cloudbuildv2_connection_iam_member.
func (ccim cloudbuildv2ConnectionIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ccim.ref.Append("location"))
}

// Member returns a reference to field member of google_cloudbuildv2_connection_iam_member.
func (ccim cloudbuildv2ConnectionIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(ccim.ref.Append("member"))
}

// Name returns a reference to field name of google_cloudbuildv2_connection_iam_member.
func (ccim cloudbuildv2ConnectionIamMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ccim.ref.Append("name"))
}

// Project returns a reference to field project of google_cloudbuildv2_connection_iam_member.
func (ccim cloudbuildv2ConnectionIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ccim.ref.Append("project"))
}

// Role returns a reference to field role of google_cloudbuildv2_connection_iam_member.
func (ccim cloudbuildv2ConnectionIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ccim.ref.Append("role"))
}

func (ccim cloudbuildv2ConnectionIamMemberAttributes) Condition() terra.ListValue[cloudbuildv2connectioniammember.ConditionAttributes] {
	return terra.ReferenceAsList[cloudbuildv2connectioniammember.ConditionAttributes](ccim.ref.Append("condition"))
}

type cloudbuildv2ConnectionIamMemberState struct {
	Etag      string                                           `json:"etag"`
	Id        string                                           `json:"id"`
	Location  string                                           `json:"location"`
	Member    string                                           `json:"member"`
	Name      string                                           `json:"name"`
	Project   string                                           `json:"project"`
	Role      string                                           `json:"role"`
	Condition []cloudbuildv2connectioniammember.ConditionState `json:"condition"`
}
