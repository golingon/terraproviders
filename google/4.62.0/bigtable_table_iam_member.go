// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigtabletableiammember "github.com/golingon/terraproviders/google/4.62.0/bigtabletableiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigtableTableIamMember creates a new instance of [BigtableTableIamMember].
func NewBigtableTableIamMember(name string, args BigtableTableIamMemberArgs) *BigtableTableIamMember {
	return &BigtableTableIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigtableTableIamMember)(nil)

// BigtableTableIamMember represents the Terraform resource google_bigtable_table_iam_member.
type BigtableTableIamMember struct {
	Name      string
	Args      BigtableTableIamMemberArgs
	state     *bigtableTableIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigtableTableIamMember].
func (btim *BigtableTableIamMember) Type() string {
	return "google_bigtable_table_iam_member"
}

// LocalName returns the local name for [BigtableTableIamMember].
func (btim *BigtableTableIamMember) LocalName() string {
	return btim.Name
}

// Configuration returns the configuration (args) for [BigtableTableIamMember].
func (btim *BigtableTableIamMember) Configuration() interface{} {
	return btim.Args
}

// DependOn is used for other resources to depend on [BigtableTableIamMember].
func (btim *BigtableTableIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(btim)
}

// Dependencies returns the list of resources [BigtableTableIamMember] depends_on.
func (btim *BigtableTableIamMember) Dependencies() terra.Dependencies {
	return btim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigtableTableIamMember].
func (btim *BigtableTableIamMember) LifecycleManagement() *terra.Lifecycle {
	return btim.Lifecycle
}

// Attributes returns the attributes for [BigtableTableIamMember].
func (btim *BigtableTableIamMember) Attributes() bigtableTableIamMemberAttributes {
	return bigtableTableIamMemberAttributes{ref: terra.ReferenceResource(btim)}
}

// ImportState imports the given attribute values into [BigtableTableIamMember]'s state.
func (btim *BigtableTableIamMember) ImportState(av io.Reader) error {
	btim.state = &bigtableTableIamMemberState{}
	if err := json.NewDecoder(av).Decode(btim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", btim.Type(), btim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigtableTableIamMember] has state.
func (btim *BigtableTableIamMember) State() (*bigtableTableIamMemberState, bool) {
	return btim.state, btim.state != nil
}

// StateMust returns the state for [BigtableTableIamMember]. Panics if the state is nil.
func (btim *BigtableTableIamMember) StateMust() *bigtableTableIamMemberState {
	if btim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", btim.Type(), btim.LocalName()))
	}
	return btim.state
}

// BigtableTableIamMemberArgs contains the configurations for google_bigtable_table_iam_member.
type BigtableTableIamMemberArgs struct {
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
	// Table: string, required
	Table terra.StringValue `hcl:"table,attr" validate:"required"`
	// Condition: optional
	Condition *bigtabletableiammember.Condition `hcl:"condition,block"`
}
type bigtableTableIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_bigtable_table_iam_member.
func (btim bigtableTableIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(btim.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigtable_table_iam_member.
func (btim bigtableTableIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(btim.ref.Append("id"))
}

// Instance returns a reference to field instance of google_bigtable_table_iam_member.
func (btim bigtableTableIamMemberAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(btim.ref.Append("instance"))
}

// Member returns a reference to field member of google_bigtable_table_iam_member.
func (btim bigtableTableIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(btim.ref.Append("member"))
}

// Project returns a reference to field project of google_bigtable_table_iam_member.
func (btim bigtableTableIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(btim.ref.Append("project"))
}

// Role returns a reference to field role of google_bigtable_table_iam_member.
func (btim bigtableTableIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(btim.ref.Append("role"))
}

// Table returns a reference to field table of google_bigtable_table_iam_member.
func (btim bigtableTableIamMemberAttributes) Table() terra.StringValue {
	return terra.ReferenceAsString(btim.ref.Append("table"))
}

func (btim bigtableTableIamMemberAttributes) Condition() terra.ListValue[bigtabletableiammember.ConditionAttributes] {
	return terra.ReferenceAsList[bigtabletableiammember.ConditionAttributes](btim.ref.Append("condition"))
}

type bigtableTableIamMemberState struct {
	Etag      string                                  `json:"etag"`
	Id        string                                  `json:"id"`
	Instance  string                                  `json:"instance"`
	Member    string                                  `json:"member"`
	Project   string                                  `json:"project"`
	Role      string                                  `json:"role"`
	Table     string                                  `json:"table"`
	Condition []bigtabletableiammember.ConditionState `json:"condition"`
}
