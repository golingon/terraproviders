// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	workstationsworkstationiammember "github.com/golingon/terraproviders/googlebeta/4.66.0/workstationsworkstationiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWorkstationsWorkstationIamMember creates a new instance of [WorkstationsWorkstationIamMember].
func NewWorkstationsWorkstationIamMember(name string, args WorkstationsWorkstationIamMemberArgs) *WorkstationsWorkstationIamMember {
	return &WorkstationsWorkstationIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WorkstationsWorkstationIamMember)(nil)

// WorkstationsWorkstationIamMember represents the Terraform resource google_workstations_workstation_iam_member.
type WorkstationsWorkstationIamMember struct {
	Name      string
	Args      WorkstationsWorkstationIamMemberArgs
	state     *workstationsWorkstationIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WorkstationsWorkstationIamMember].
func (wwim *WorkstationsWorkstationIamMember) Type() string {
	return "google_workstations_workstation_iam_member"
}

// LocalName returns the local name for [WorkstationsWorkstationIamMember].
func (wwim *WorkstationsWorkstationIamMember) LocalName() string {
	return wwim.Name
}

// Configuration returns the configuration (args) for [WorkstationsWorkstationIamMember].
func (wwim *WorkstationsWorkstationIamMember) Configuration() interface{} {
	return wwim.Args
}

// DependOn is used for other resources to depend on [WorkstationsWorkstationIamMember].
func (wwim *WorkstationsWorkstationIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(wwim)
}

// Dependencies returns the list of resources [WorkstationsWorkstationIamMember] depends_on.
func (wwim *WorkstationsWorkstationIamMember) Dependencies() terra.Dependencies {
	return wwim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WorkstationsWorkstationIamMember].
func (wwim *WorkstationsWorkstationIamMember) LifecycleManagement() *terra.Lifecycle {
	return wwim.Lifecycle
}

// Attributes returns the attributes for [WorkstationsWorkstationIamMember].
func (wwim *WorkstationsWorkstationIamMember) Attributes() workstationsWorkstationIamMemberAttributes {
	return workstationsWorkstationIamMemberAttributes{ref: terra.ReferenceResource(wwim)}
}

// ImportState imports the given attribute values into [WorkstationsWorkstationIamMember]'s state.
func (wwim *WorkstationsWorkstationIamMember) ImportState(av io.Reader) error {
	wwim.state = &workstationsWorkstationIamMemberState{}
	if err := json.NewDecoder(av).Decode(wwim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wwim.Type(), wwim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WorkstationsWorkstationIamMember] has state.
func (wwim *WorkstationsWorkstationIamMember) State() (*workstationsWorkstationIamMemberState, bool) {
	return wwim.state, wwim.state != nil
}

// StateMust returns the state for [WorkstationsWorkstationIamMember]. Panics if the state is nil.
func (wwim *WorkstationsWorkstationIamMember) StateMust() *workstationsWorkstationIamMemberState {
	if wwim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wwim.Type(), wwim.LocalName()))
	}
	return wwim.state
}

// WorkstationsWorkstationIamMemberArgs contains the configurations for google_workstations_workstation_iam_member.
type WorkstationsWorkstationIamMemberArgs struct {
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
	// WorkstationClusterId: string, required
	WorkstationClusterId terra.StringValue `hcl:"workstation_cluster_id,attr" validate:"required"`
	// WorkstationConfigId: string, required
	WorkstationConfigId terra.StringValue `hcl:"workstation_config_id,attr" validate:"required"`
	// WorkstationId: string, required
	WorkstationId terra.StringValue `hcl:"workstation_id,attr" validate:"required"`
	// Condition: optional
	Condition *workstationsworkstationiammember.Condition `hcl:"condition,block"`
}
type workstationsWorkstationIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_workstations_workstation_iam_member.
func (wwim workstationsWorkstationIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(wwim.ref.Append("etag"))
}

// Id returns a reference to field id of google_workstations_workstation_iam_member.
func (wwim workstationsWorkstationIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wwim.ref.Append("id"))
}

// Location returns a reference to field location of google_workstations_workstation_iam_member.
func (wwim workstationsWorkstationIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(wwim.ref.Append("location"))
}

// Member returns a reference to field member of google_workstations_workstation_iam_member.
func (wwim workstationsWorkstationIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(wwim.ref.Append("member"))
}

// Project returns a reference to field project of google_workstations_workstation_iam_member.
func (wwim workstationsWorkstationIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(wwim.ref.Append("project"))
}

// Role returns a reference to field role of google_workstations_workstation_iam_member.
func (wwim workstationsWorkstationIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(wwim.ref.Append("role"))
}

// WorkstationClusterId returns a reference to field workstation_cluster_id of google_workstations_workstation_iam_member.
func (wwim workstationsWorkstationIamMemberAttributes) WorkstationClusterId() terra.StringValue {
	return terra.ReferenceAsString(wwim.ref.Append("workstation_cluster_id"))
}

// WorkstationConfigId returns a reference to field workstation_config_id of google_workstations_workstation_iam_member.
func (wwim workstationsWorkstationIamMemberAttributes) WorkstationConfigId() terra.StringValue {
	return terra.ReferenceAsString(wwim.ref.Append("workstation_config_id"))
}

// WorkstationId returns a reference to field workstation_id of google_workstations_workstation_iam_member.
func (wwim workstationsWorkstationIamMemberAttributes) WorkstationId() terra.StringValue {
	return terra.ReferenceAsString(wwim.ref.Append("workstation_id"))
}

func (wwim workstationsWorkstationIamMemberAttributes) Condition() terra.ListValue[workstationsworkstationiammember.ConditionAttributes] {
	return terra.ReferenceAsList[workstationsworkstationiammember.ConditionAttributes](wwim.ref.Append("condition"))
}

type workstationsWorkstationIamMemberState struct {
	Etag                 string                                            `json:"etag"`
	Id                   string                                            `json:"id"`
	Location             string                                            `json:"location"`
	Member               string                                            `json:"member"`
	Project              string                                            `json:"project"`
	Role                 string                                            `json:"role"`
	WorkstationClusterId string                                            `json:"workstation_cluster_id"`
	WorkstationConfigId  string                                            `json:"workstation_config_id"`
	WorkstationId        string                                            `json:"workstation_id"`
	Condition            []workstationsworkstationiammember.ConditionState `json:"condition"`
}
