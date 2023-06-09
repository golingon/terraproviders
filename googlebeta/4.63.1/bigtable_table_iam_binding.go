// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	bigtabletableiambinding "github.com/golingon/terraproviders/googlebeta/4.63.1/bigtabletableiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigtableTableIamBinding creates a new instance of [BigtableTableIamBinding].
func NewBigtableTableIamBinding(name string, args BigtableTableIamBindingArgs) *BigtableTableIamBinding {
	return &BigtableTableIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigtableTableIamBinding)(nil)

// BigtableTableIamBinding represents the Terraform resource google_bigtable_table_iam_binding.
type BigtableTableIamBinding struct {
	Name      string
	Args      BigtableTableIamBindingArgs
	state     *bigtableTableIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigtableTableIamBinding].
func (btib *BigtableTableIamBinding) Type() string {
	return "google_bigtable_table_iam_binding"
}

// LocalName returns the local name for [BigtableTableIamBinding].
func (btib *BigtableTableIamBinding) LocalName() string {
	return btib.Name
}

// Configuration returns the configuration (args) for [BigtableTableIamBinding].
func (btib *BigtableTableIamBinding) Configuration() interface{} {
	return btib.Args
}

// DependOn is used for other resources to depend on [BigtableTableIamBinding].
func (btib *BigtableTableIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(btib)
}

// Dependencies returns the list of resources [BigtableTableIamBinding] depends_on.
func (btib *BigtableTableIamBinding) Dependencies() terra.Dependencies {
	return btib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigtableTableIamBinding].
func (btib *BigtableTableIamBinding) LifecycleManagement() *terra.Lifecycle {
	return btib.Lifecycle
}

// Attributes returns the attributes for [BigtableTableIamBinding].
func (btib *BigtableTableIamBinding) Attributes() bigtableTableIamBindingAttributes {
	return bigtableTableIamBindingAttributes{ref: terra.ReferenceResource(btib)}
}

// ImportState imports the given attribute values into [BigtableTableIamBinding]'s state.
func (btib *BigtableTableIamBinding) ImportState(av io.Reader) error {
	btib.state = &bigtableTableIamBindingState{}
	if err := json.NewDecoder(av).Decode(btib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", btib.Type(), btib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigtableTableIamBinding] has state.
func (btib *BigtableTableIamBinding) State() (*bigtableTableIamBindingState, bool) {
	return btib.state, btib.state != nil
}

// StateMust returns the state for [BigtableTableIamBinding]. Panics if the state is nil.
func (btib *BigtableTableIamBinding) StateMust() *bigtableTableIamBindingState {
	if btib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", btib.Type(), btib.LocalName()))
	}
	return btib.state
}

// BigtableTableIamBindingArgs contains the configurations for google_bigtable_table_iam_binding.
type BigtableTableIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Table: string, required
	Table terra.StringValue `hcl:"table,attr" validate:"required"`
	// Condition: optional
	Condition *bigtabletableiambinding.Condition `hcl:"condition,block"`
}
type bigtableTableIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_bigtable_table_iam_binding.
func (btib bigtableTableIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(btib.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigtable_table_iam_binding.
func (btib bigtableTableIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(btib.ref.Append("id"))
}

// Instance returns a reference to field instance of google_bigtable_table_iam_binding.
func (btib bigtableTableIamBindingAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(btib.ref.Append("instance"))
}

// Members returns a reference to field members of google_bigtable_table_iam_binding.
func (btib bigtableTableIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](btib.ref.Append("members"))
}

// Project returns a reference to field project of google_bigtable_table_iam_binding.
func (btib bigtableTableIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(btib.ref.Append("project"))
}

// Role returns a reference to field role of google_bigtable_table_iam_binding.
func (btib bigtableTableIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(btib.ref.Append("role"))
}

// Table returns a reference to field table of google_bigtable_table_iam_binding.
func (btib bigtableTableIamBindingAttributes) Table() terra.StringValue {
	return terra.ReferenceAsString(btib.ref.Append("table"))
}

func (btib bigtableTableIamBindingAttributes) Condition() terra.ListValue[bigtabletableiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[bigtabletableiambinding.ConditionAttributes](btib.ref.Append("condition"))
}

type bigtableTableIamBindingState struct {
	Etag      string                                   `json:"etag"`
	Id        string                                   `json:"id"`
	Instance  string                                   `json:"instance"`
	Members   []string                                 `json:"members"`
	Project   string                                   `json:"project"`
	Role      string                                   `json:"role"`
	Table     string                                   `json:"table"`
	Condition []bigtabletableiambinding.ConditionState `json:"condition"`
}
