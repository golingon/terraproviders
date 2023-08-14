// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	datafusioninstanceiambinding "github.com/golingon/terraproviders/googlebeta/4.77.0/datafusioninstanceiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFusionInstanceIamBinding creates a new instance of [DataFusionInstanceIamBinding].
func NewDataFusionInstanceIamBinding(name string, args DataFusionInstanceIamBindingArgs) *DataFusionInstanceIamBinding {
	return &DataFusionInstanceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFusionInstanceIamBinding)(nil)

// DataFusionInstanceIamBinding represents the Terraform resource google_data_fusion_instance_iam_binding.
type DataFusionInstanceIamBinding struct {
	Name      string
	Args      DataFusionInstanceIamBindingArgs
	state     *dataFusionInstanceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFusionInstanceIamBinding].
func (dfiib *DataFusionInstanceIamBinding) Type() string {
	return "google_data_fusion_instance_iam_binding"
}

// LocalName returns the local name for [DataFusionInstanceIamBinding].
func (dfiib *DataFusionInstanceIamBinding) LocalName() string {
	return dfiib.Name
}

// Configuration returns the configuration (args) for [DataFusionInstanceIamBinding].
func (dfiib *DataFusionInstanceIamBinding) Configuration() interface{} {
	return dfiib.Args
}

// DependOn is used for other resources to depend on [DataFusionInstanceIamBinding].
func (dfiib *DataFusionInstanceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(dfiib)
}

// Dependencies returns the list of resources [DataFusionInstanceIamBinding] depends_on.
func (dfiib *DataFusionInstanceIamBinding) Dependencies() terra.Dependencies {
	return dfiib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFusionInstanceIamBinding].
func (dfiib *DataFusionInstanceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return dfiib.Lifecycle
}

// Attributes returns the attributes for [DataFusionInstanceIamBinding].
func (dfiib *DataFusionInstanceIamBinding) Attributes() dataFusionInstanceIamBindingAttributes {
	return dataFusionInstanceIamBindingAttributes{ref: terra.ReferenceResource(dfiib)}
}

// ImportState imports the given attribute values into [DataFusionInstanceIamBinding]'s state.
func (dfiib *DataFusionInstanceIamBinding) ImportState(av io.Reader) error {
	dfiib.state = &dataFusionInstanceIamBindingState{}
	if err := json.NewDecoder(av).Decode(dfiib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfiib.Type(), dfiib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFusionInstanceIamBinding] has state.
func (dfiib *DataFusionInstanceIamBinding) State() (*dataFusionInstanceIamBindingState, bool) {
	return dfiib.state, dfiib.state != nil
}

// StateMust returns the state for [DataFusionInstanceIamBinding]. Panics if the state is nil.
func (dfiib *DataFusionInstanceIamBinding) StateMust() *dataFusionInstanceIamBindingState {
	if dfiib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfiib.Type(), dfiib.LocalName()))
	}
	return dfiib.state
}

// DataFusionInstanceIamBindingArgs contains the configurations for google_data_fusion_instance_iam_binding.
type DataFusionInstanceIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *datafusioninstanceiambinding.Condition `hcl:"condition,block"`
}
type dataFusionInstanceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_data_fusion_instance_iam_binding.
func (dfiib dataFusionInstanceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dfiib.ref.Append("etag"))
}

// Id returns a reference to field id of google_data_fusion_instance_iam_binding.
func (dfiib dataFusionInstanceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfiib.ref.Append("id"))
}

// Members returns a reference to field members of google_data_fusion_instance_iam_binding.
func (dfiib dataFusionInstanceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dfiib.ref.Append("members"))
}

// Name returns a reference to field name of google_data_fusion_instance_iam_binding.
func (dfiib dataFusionInstanceIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfiib.ref.Append("name"))
}

// Project returns a reference to field project of google_data_fusion_instance_iam_binding.
func (dfiib dataFusionInstanceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dfiib.ref.Append("project"))
}

// Region returns a reference to field region of google_data_fusion_instance_iam_binding.
func (dfiib dataFusionInstanceIamBindingAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dfiib.ref.Append("region"))
}

// Role returns a reference to field role of google_data_fusion_instance_iam_binding.
func (dfiib dataFusionInstanceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dfiib.ref.Append("role"))
}

func (dfiib dataFusionInstanceIamBindingAttributes) Condition() terra.ListValue[datafusioninstanceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[datafusioninstanceiambinding.ConditionAttributes](dfiib.ref.Append("condition"))
}

type dataFusionInstanceIamBindingState struct {
	Etag      string                                        `json:"etag"`
	Id        string                                        `json:"id"`
	Members   []string                                      `json:"members"`
	Name      string                                        `json:"name"`
	Project   string                                        `json:"project"`
	Region    string                                        `json:"region"`
	Role      string                                        `json:"role"`
	Condition []datafusioninstanceiambinding.ConditionState `json:"condition"`
}
