// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	bigtableinstanceiambinding "github.com/golingon/terraproviders/google/5.24.0/bigtableinstanceiambinding"
	"io"
)

// NewBigtableInstanceIamBinding creates a new instance of [BigtableInstanceIamBinding].
func NewBigtableInstanceIamBinding(name string, args BigtableInstanceIamBindingArgs) *BigtableInstanceIamBinding {
	return &BigtableInstanceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigtableInstanceIamBinding)(nil)

// BigtableInstanceIamBinding represents the Terraform resource google_bigtable_instance_iam_binding.
type BigtableInstanceIamBinding struct {
	Name      string
	Args      BigtableInstanceIamBindingArgs
	state     *bigtableInstanceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigtableInstanceIamBinding].
func (biib *BigtableInstanceIamBinding) Type() string {
	return "google_bigtable_instance_iam_binding"
}

// LocalName returns the local name for [BigtableInstanceIamBinding].
func (biib *BigtableInstanceIamBinding) LocalName() string {
	return biib.Name
}

// Configuration returns the configuration (args) for [BigtableInstanceIamBinding].
func (biib *BigtableInstanceIamBinding) Configuration() interface{} {
	return biib.Args
}

// DependOn is used for other resources to depend on [BigtableInstanceIamBinding].
func (biib *BigtableInstanceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(biib)
}

// Dependencies returns the list of resources [BigtableInstanceIamBinding] depends_on.
func (biib *BigtableInstanceIamBinding) Dependencies() terra.Dependencies {
	return biib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigtableInstanceIamBinding].
func (biib *BigtableInstanceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return biib.Lifecycle
}

// Attributes returns the attributes for [BigtableInstanceIamBinding].
func (biib *BigtableInstanceIamBinding) Attributes() bigtableInstanceIamBindingAttributes {
	return bigtableInstanceIamBindingAttributes{ref: terra.ReferenceResource(biib)}
}

// ImportState imports the given attribute values into [BigtableInstanceIamBinding]'s state.
func (biib *BigtableInstanceIamBinding) ImportState(av io.Reader) error {
	biib.state = &bigtableInstanceIamBindingState{}
	if err := json.NewDecoder(av).Decode(biib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", biib.Type(), biib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigtableInstanceIamBinding] has state.
func (biib *BigtableInstanceIamBinding) State() (*bigtableInstanceIamBindingState, bool) {
	return biib.state, biib.state != nil
}

// StateMust returns the state for [BigtableInstanceIamBinding]. Panics if the state is nil.
func (biib *BigtableInstanceIamBinding) StateMust() *bigtableInstanceIamBindingState {
	if biib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", biib.Type(), biib.LocalName()))
	}
	return biib.state
}

// BigtableInstanceIamBindingArgs contains the configurations for google_bigtable_instance_iam_binding.
type BigtableInstanceIamBindingArgs struct {
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
	// Condition: optional
	Condition *bigtableinstanceiambinding.Condition `hcl:"condition,block"`
}
type bigtableInstanceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_bigtable_instance_iam_binding.
func (biib bigtableInstanceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(biib.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigtable_instance_iam_binding.
func (biib bigtableInstanceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(biib.ref.Append("id"))
}

// Instance returns a reference to field instance of google_bigtable_instance_iam_binding.
func (biib bigtableInstanceIamBindingAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(biib.ref.Append("instance"))
}

// Members returns a reference to field members of google_bigtable_instance_iam_binding.
func (biib bigtableInstanceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](biib.ref.Append("members"))
}

// Project returns a reference to field project of google_bigtable_instance_iam_binding.
func (biib bigtableInstanceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(biib.ref.Append("project"))
}

// Role returns a reference to field role of google_bigtable_instance_iam_binding.
func (biib bigtableInstanceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(biib.ref.Append("role"))
}

func (biib bigtableInstanceIamBindingAttributes) Condition() terra.ListValue[bigtableinstanceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[bigtableinstanceiambinding.ConditionAttributes](biib.ref.Append("condition"))
}

type bigtableInstanceIamBindingState struct {
	Etag      string                                      `json:"etag"`
	Id        string                                      `json:"id"`
	Instance  string                                      `json:"instance"`
	Members   []string                                    `json:"members"`
	Project   string                                      `json:"project"`
	Role      string                                      `json:"role"`
	Condition []bigtableinstanceiambinding.ConditionState `json:"condition"`
}
