// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	workstationsworkstationconfigiambinding "github.com/golingon/terraproviders/googlebeta/4.75.1/workstationsworkstationconfigiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWorkstationsWorkstationConfigIamBinding creates a new instance of [WorkstationsWorkstationConfigIamBinding].
func NewWorkstationsWorkstationConfigIamBinding(name string, args WorkstationsWorkstationConfigIamBindingArgs) *WorkstationsWorkstationConfigIamBinding {
	return &WorkstationsWorkstationConfigIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WorkstationsWorkstationConfigIamBinding)(nil)

// WorkstationsWorkstationConfigIamBinding represents the Terraform resource google_workstations_workstation_config_iam_binding.
type WorkstationsWorkstationConfigIamBinding struct {
	Name      string
	Args      WorkstationsWorkstationConfigIamBindingArgs
	state     *workstationsWorkstationConfigIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WorkstationsWorkstationConfigIamBinding].
func (wwcib *WorkstationsWorkstationConfigIamBinding) Type() string {
	return "google_workstations_workstation_config_iam_binding"
}

// LocalName returns the local name for [WorkstationsWorkstationConfigIamBinding].
func (wwcib *WorkstationsWorkstationConfigIamBinding) LocalName() string {
	return wwcib.Name
}

// Configuration returns the configuration (args) for [WorkstationsWorkstationConfigIamBinding].
func (wwcib *WorkstationsWorkstationConfigIamBinding) Configuration() interface{} {
	return wwcib.Args
}

// DependOn is used for other resources to depend on [WorkstationsWorkstationConfigIamBinding].
func (wwcib *WorkstationsWorkstationConfigIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(wwcib)
}

// Dependencies returns the list of resources [WorkstationsWorkstationConfigIamBinding] depends_on.
func (wwcib *WorkstationsWorkstationConfigIamBinding) Dependencies() terra.Dependencies {
	return wwcib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WorkstationsWorkstationConfigIamBinding].
func (wwcib *WorkstationsWorkstationConfigIamBinding) LifecycleManagement() *terra.Lifecycle {
	return wwcib.Lifecycle
}

// Attributes returns the attributes for [WorkstationsWorkstationConfigIamBinding].
func (wwcib *WorkstationsWorkstationConfigIamBinding) Attributes() workstationsWorkstationConfigIamBindingAttributes {
	return workstationsWorkstationConfigIamBindingAttributes{ref: terra.ReferenceResource(wwcib)}
}

// ImportState imports the given attribute values into [WorkstationsWorkstationConfigIamBinding]'s state.
func (wwcib *WorkstationsWorkstationConfigIamBinding) ImportState(av io.Reader) error {
	wwcib.state = &workstationsWorkstationConfigIamBindingState{}
	if err := json.NewDecoder(av).Decode(wwcib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wwcib.Type(), wwcib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WorkstationsWorkstationConfigIamBinding] has state.
func (wwcib *WorkstationsWorkstationConfigIamBinding) State() (*workstationsWorkstationConfigIamBindingState, bool) {
	return wwcib.state, wwcib.state != nil
}

// StateMust returns the state for [WorkstationsWorkstationConfigIamBinding]. Panics if the state is nil.
func (wwcib *WorkstationsWorkstationConfigIamBinding) StateMust() *workstationsWorkstationConfigIamBindingState {
	if wwcib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wwcib.Type(), wwcib.LocalName()))
	}
	return wwcib.state
}

// WorkstationsWorkstationConfigIamBindingArgs contains the configurations for google_workstations_workstation_config_iam_binding.
type WorkstationsWorkstationConfigIamBindingArgs struct {
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
	// Condition: optional
	Condition *workstationsworkstationconfigiambinding.Condition `hcl:"condition,block"`
}
type workstationsWorkstationConfigIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_workstations_workstation_config_iam_binding.
func (wwcib workstationsWorkstationConfigIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(wwcib.ref.Append("etag"))
}

// Id returns a reference to field id of google_workstations_workstation_config_iam_binding.
func (wwcib workstationsWorkstationConfigIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wwcib.ref.Append("id"))
}

// Location returns a reference to field location of google_workstations_workstation_config_iam_binding.
func (wwcib workstationsWorkstationConfigIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(wwcib.ref.Append("location"))
}

// Members returns a reference to field members of google_workstations_workstation_config_iam_binding.
func (wwcib workstationsWorkstationConfigIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](wwcib.ref.Append("members"))
}

// Project returns a reference to field project of google_workstations_workstation_config_iam_binding.
func (wwcib workstationsWorkstationConfigIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(wwcib.ref.Append("project"))
}

// Role returns a reference to field role of google_workstations_workstation_config_iam_binding.
func (wwcib workstationsWorkstationConfigIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(wwcib.ref.Append("role"))
}

// WorkstationClusterId returns a reference to field workstation_cluster_id of google_workstations_workstation_config_iam_binding.
func (wwcib workstationsWorkstationConfigIamBindingAttributes) WorkstationClusterId() terra.StringValue {
	return terra.ReferenceAsString(wwcib.ref.Append("workstation_cluster_id"))
}

// WorkstationConfigId returns a reference to field workstation_config_id of google_workstations_workstation_config_iam_binding.
func (wwcib workstationsWorkstationConfigIamBindingAttributes) WorkstationConfigId() terra.StringValue {
	return terra.ReferenceAsString(wwcib.ref.Append("workstation_config_id"))
}

func (wwcib workstationsWorkstationConfigIamBindingAttributes) Condition() terra.ListValue[workstationsworkstationconfigiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[workstationsworkstationconfigiambinding.ConditionAttributes](wwcib.ref.Append("condition"))
}

type workstationsWorkstationConfigIamBindingState struct {
	Etag                 string                                                   `json:"etag"`
	Id                   string                                                   `json:"id"`
	Location             string                                                   `json:"location"`
	Members              []string                                                 `json:"members"`
	Project              string                                                   `json:"project"`
	Role                 string                                                   `json:"role"`
	WorkstationClusterId string                                                   `json:"workstation_cluster_id"`
	WorkstationConfigId  string                                                   `json:"workstation_config_id"`
	Condition            []workstationsworkstationconfigiambinding.ConditionState `json:"condition"`
}
