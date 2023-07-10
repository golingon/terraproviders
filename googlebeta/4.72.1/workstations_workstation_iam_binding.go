// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	workstationsworkstationiambinding "github.com/golingon/terraproviders/googlebeta/4.72.1/workstationsworkstationiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWorkstationsWorkstationIamBinding creates a new instance of [WorkstationsWorkstationIamBinding].
func NewWorkstationsWorkstationIamBinding(name string, args WorkstationsWorkstationIamBindingArgs) *WorkstationsWorkstationIamBinding {
	return &WorkstationsWorkstationIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WorkstationsWorkstationIamBinding)(nil)

// WorkstationsWorkstationIamBinding represents the Terraform resource google_workstations_workstation_iam_binding.
type WorkstationsWorkstationIamBinding struct {
	Name      string
	Args      WorkstationsWorkstationIamBindingArgs
	state     *workstationsWorkstationIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WorkstationsWorkstationIamBinding].
func (wwib *WorkstationsWorkstationIamBinding) Type() string {
	return "google_workstations_workstation_iam_binding"
}

// LocalName returns the local name for [WorkstationsWorkstationIamBinding].
func (wwib *WorkstationsWorkstationIamBinding) LocalName() string {
	return wwib.Name
}

// Configuration returns the configuration (args) for [WorkstationsWorkstationIamBinding].
func (wwib *WorkstationsWorkstationIamBinding) Configuration() interface{} {
	return wwib.Args
}

// DependOn is used for other resources to depend on [WorkstationsWorkstationIamBinding].
func (wwib *WorkstationsWorkstationIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(wwib)
}

// Dependencies returns the list of resources [WorkstationsWorkstationIamBinding] depends_on.
func (wwib *WorkstationsWorkstationIamBinding) Dependencies() terra.Dependencies {
	return wwib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WorkstationsWorkstationIamBinding].
func (wwib *WorkstationsWorkstationIamBinding) LifecycleManagement() *terra.Lifecycle {
	return wwib.Lifecycle
}

// Attributes returns the attributes for [WorkstationsWorkstationIamBinding].
func (wwib *WorkstationsWorkstationIamBinding) Attributes() workstationsWorkstationIamBindingAttributes {
	return workstationsWorkstationIamBindingAttributes{ref: terra.ReferenceResource(wwib)}
}

// ImportState imports the given attribute values into [WorkstationsWorkstationIamBinding]'s state.
func (wwib *WorkstationsWorkstationIamBinding) ImportState(av io.Reader) error {
	wwib.state = &workstationsWorkstationIamBindingState{}
	if err := json.NewDecoder(av).Decode(wwib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wwib.Type(), wwib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WorkstationsWorkstationIamBinding] has state.
func (wwib *WorkstationsWorkstationIamBinding) State() (*workstationsWorkstationIamBindingState, bool) {
	return wwib.state, wwib.state != nil
}

// StateMust returns the state for [WorkstationsWorkstationIamBinding]. Panics if the state is nil.
func (wwib *WorkstationsWorkstationIamBinding) StateMust() *workstationsWorkstationIamBindingState {
	if wwib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wwib.Type(), wwib.LocalName()))
	}
	return wwib.state
}

// WorkstationsWorkstationIamBindingArgs contains the configurations for google_workstations_workstation_iam_binding.
type WorkstationsWorkstationIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// WorkstationClusterId: string, required
	WorkstationClusterId terra.StringValue `hcl:"workstation_cluster_id,attr" validate:"required"`
	// WorkstationConfigId: string, required
	WorkstationConfigId terra.StringValue `hcl:"workstation_config_id,attr" validate:"required"`
	// WorkstationId: string, required
	WorkstationId terra.StringValue `hcl:"workstation_id,attr" validate:"required"`
	// Condition: optional
	Condition *workstationsworkstationiambinding.Condition `hcl:"condition,block"`
}
type workstationsWorkstationIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_workstations_workstation_iam_binding.
func (wwib workstationsWorkstationIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(wwib.ref.Append("etag"))
}

// Id returns a reference to field id of google_workstations_workstation_iam_binding.
func (wwib workstationsWorkstationIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wwib.ref.Append("id"))
}

// Location returns a reference to field location of google_workstations_workstation_iam_binding.
func (wwib workstationsWorkstationIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(wwib.ref.Append("location"))
}

// Members returns a reference to field members of google_workstations_workstation_iam_binding.
func (wwib workstationsWorkstationIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](wwib.ref.Append("members"))
}

// Project returns a reference to field project of google_workstations_workstation_iam_binding.
func (wwib workstationsWorkstationIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(wwib.ref.Append("project"))
}

// Role returns a reference to field role of google_workstations_workstation_iam_binding.
func (wwib workstationsWorkstationIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(wwib.ref.Append("role"))
}

// WorkstationClusterId returns a reference to field workstation_cluster_id of google_workstations_workstation_iam_binding.
func (wwib workstationsWorkstationIamBindingAttributes) WorkstationClusterId() terra.StringValue {
	return terra.ReferenceAsString(wwib.ref.Append("workstation_cluster_id"))
}

// WorkstationConfigId returns a reference to field workstation_config_id of google_workstations_workstation_iam_binding.
func (wwib workstationsWorkstationIamBindingAttributes) WorkstationConfigId() terra.StringValue {
	return terra.ReferenceAsString(wwib.ref.Append("workstation_config_id"))
}

// WorkstationId returns a reference to field workstation_id of google_workstations_workstation_iam_binding.
func (wwib workstationsWorkstationIamBindingAttributes) WorkstationId() terra.StringValue {
	return terra.ReferenceAsString(wwib.ref.Append("workstation_id"))
}

func (wwib workstationsWorkstationIamBindingAttributes) Condition() terra.ListValue[workstationsworkstationiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[workstationsworkstationiambinding.ConditionAttributes](wwib.ref.Append("condition"))
}

type workstationsWorkstationIamBindingState struct {
	Etag                 string                                             `json:"etag"`
	Id                   string                                             `json:"id"`
	Location             string                                             `json:"location"`
	Members              []string                                           `json:"members"`
	Project              string                                             `json:"project"`
	Role                 string                                             `json:"role"`
	WorkstationClusterId string                                             `json:"workstation_cluster_id"`
	WorkstationConfigId  string                                             `json:"workstation_config_id"`
	WorkstationId        string                                             `json:"workstation_id"`
	Condition            []workstationsworkstationiambinding.ConditionState `json:"condition"`
}
