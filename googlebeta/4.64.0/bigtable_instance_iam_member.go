// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	bigtableinstanceiammember "github.com/golingon/terraproviders/googlebeta/4.64.0/bigtableinstanceiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigtableInstanceIamMember creates a new instance of [BigtableInstanceIamMember].
func NewBigtableInstanceIamMember(name string, args BigtableInstanceIamMemberArgs) *BigtableInstanceIamMember {
	return &BigtableInstanceIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigtableInstanceIamMember)(nil)

// BigtableInstanceIamMember represents the Terraform resource google_bigtable_instance_iam_member.
type BigtableInstanceIamMember struct {
	Name      string
	Args      BigtableInstanceIamMemberArgs
	state     *bigtableInstanceIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigtableInstanceIamMember].
func (biim *BigtableInstanceIamMember) Type() string {
	return "google_bigtable_instance_iam_member"
}

// LocalName returns the local name for [BigtableInstanceIamMember].
func (biim *BigtableInstanceIamMember) LocalName() string {
	return biim.Name
}

// Configuration returns the configuration (args) for [BigtableInstanceIamMember].
func (biim *BigtableInstanceIamMember) Configuration() interface{} {
	return biim.Args
}

// DependOn is used for other resources to depend on [BigtableInstanceIamMember].
func (biim *BigtableInstanceIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(biim)
}

// Dependencies returns the list of resources [BigtableInstanceIamMember] depends_on.
func (biim *BigtableInstanceIamMember) Dependencies() terra.Dependencies {
	return biim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigtableInstanceIamMember].
func (biim *BigtableInstanceIamMember) LifecycleManagement() *terra.Lifecycle {
	return biim.Lifecycle
}

// Attributes returns the attributes for [BigtableInstanceIamMember].
func (biim *BigtableInstanceIamMember) Attributes() bigtableInstanceIamMemberAttributes {
	return bigtableInstanceIamMemberAttributes{ref: terra.ReferenceResource(biim)}
}

// ImportState imports the given attribute values into [BigtableInstanceIamMember]'s state.
func (biim *BigtableInstanceIamMember) ImportState(av io.Reader) error {
	biim.state = &bigtableInstanceIamMemberState{}
	if err := json.NewDecoder(av).Decode(biim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", biim.Type(), biim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigtableInstanceIamMember] has state.
func (biim *BigtableInstanceIamMember) State() (*bigtableInstanceIamMemberState, bool) {
	return biim.state, biim.state != nil
}

// StateMust returns the state for [BigtableInstanceIamMember]. Panics if the state is nil.
func (biim *BigtableInstanceIamMember) StateMust() *bigtableInstanceIamMemberState {
	if biim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", biim.Type(), biim.LocalName()))
	}
	return biim.state
}

// BigtableInstanceIamMemberArgs contains the configurations for google_bigtable_instance_iam_member.
type BigtableInstanceIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *bigtableinstanceiammember.Condition `hcl:"condition,block"`
}
type bigtableInstanceIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_bigtable_instance_iam_member.
func (biim bigtableInstanceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(biim.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigtable_instance_iam_member.
func (biim bigtableInstanceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(biim.ref.Append("id"))
}

// Instance returns a reference to field instance of google_bigtable_instance_iam_member.
func (biim bigtableInstanceIamMemberAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(biim.ref.Append("instance"))
}

// Member returns a reference to field member of google_bigtable_instance_iam_member.
func (biim bigtableInstanceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(biim.ref.Append("member"))
}

// Project returns a reference to field project of google_bigtable_instance_iam_member.
func (biim bigtableInstanceIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(biim.ref.Append("project"))
}

// Role returns a reference to field role of google_bigtable_instance_iam_member.
func (biim bigtableInstanceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(biim.ref.Append("role"))
}

func (biim bigtableInstanceIamMemberAttributes) Condition() terra.ListValue[bigtableinstanceiammember.ConditionAttributes] {
	return terra.ReferenceAsList[bigtableinstanceiammember.ConditionAttributes](biim.ref.Append("condition"))
}

type bigtableInstanceIamMemberState struct {
	Etag      string                                     `json:"etag"`
	Id        string                                     `json:"id"`
	Instance  string                                     `json:"instance"`
	Member    string                                     `json:"member"`
	Project   string                                     `json:"project"`
	Role      string                                     `json:"role"`
	Condition []bigtableinstanceiammember.ConditionState `json:"condition"`
}
