// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	synapseroleassignment "github.com/golingon/terraproviders/azurerm/3.58.0/synapseroleassignment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSynapseRoleAssignment creates a new instance of [SynapseRoleAssignment].
func NewSynapseRoleAssignment(name string, args SynapseRoleAssignmentArgs) *SynapseRoleAssignment {
	return &SynapseRoleAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SynapseRoleAssignment)(nil)

// SynapseRoleAssignment represents the Terraform resource azurerm_synapse_role_assignment.
type SynapseRoleAssignment struct {
	Name      string
	Args      SynapseRoleAssignmentArgs
	state     *synapseRoleAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SynapseRoleAssignment].
func (sra *SynapseRoleAssignment) Type() string {
	return "azurerm_synapse_role_assignment"
}

// LocalName returns the local name for [SynapseRoleAssignment].
func (sra *SynapseRoleAssignment) LocalName() string {
	return sra.Name
}

// Configuration returns the configuration (args) for [SynapseRoleAssignment].
func (sra *SynapseRoleAssignment) Configuration() interface{} {
	return sra.Args
}

// DependOn is used for other resources to depend on [SynapseRoleAssignment].
func (sra *SynapseRoleAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(sra)
}

// Dependencies returns the list of resources [SynapseRoleAssignment] depends_on.
func (sra *SynapseRoleAssignment) Dependencies() terra.Dependencies {
	return sra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SynapseRoleAssignment].
func (sra *SynapseRoleAssignment) LifecycleManagement() *terra.Lifecycle {
	return sra.Lifecycle
}

// Attributes returns the attributes for [SynapseRoleAssignment].
func (sra *SynapseRoleAssignment) Attributes() synapseRoleAssignmentAttributes {
	return synapseRoleAssignmentAttributes{ref: terra.ReferenceResource(sra)}
}

// ImportState imports the given attribute values into [SynapseRoleAssignment]'s state.
func (sra *SynapseRoleAssignment) ImportState(av io.Reader) error {
	sra.state = &synapseRoleAssignmentState{}
	if err := json.NewDecoder(av).Decode(sra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sra.Type(), sra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SynapseRoleAssignment] has state.
func (sra *SynapseRoleAssignment) State() (*synapseRoleAssignmentState, bool) {
	return sra.state, sra.state != nil
}

// StateMust returns the state for [SynapseRoleAssignment]. Panics if the state is nil.
func (sra *SynapseRoleAssignment) StateMust() *synapseRoleAssignmentState {
	if sra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sra.Type(), sra.LocalName()))
	}
	return sra.state
}

// SynapseRoleAssignmentArgs contains the configurations for azurerm_synapse_role_assignment.
type SynapseRoleAssignmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PrincipalId: string, required
	PrincipalId terra.StringValue `hcl:"principal_id,attr" validate:"required"`
	// RoleName: string, required
	RoleName terra.StringValue `hcl:"role_name,attr" validate:"required"`
	// SynapseSparkPoolId: string, optional
	SynapseSparkPoolId terra.StringValue `hcl:"synapse_spark_pool_id,attr"`
	// SynapseWorkspaceId: string, optional
	SynapseWorkspaceId terra.StringValue `hcl:"synapse_workspace_id,attr"`
	// Timeouts: optional
	Timeouts *synapseroleassignment.Timeouts `hcl:"timeouts,block"`
}
type synapseRoleAssignmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_synapse_role_assignment.
func (sra synapseRoleAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sra.ref.Append("id"))
}

// PrincipalId returns a reference to field principal_id of azurerm_synapse_role_assignment.
func (sra synapseRoleAssignmentAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(sra.ref.Append("principal_id"))
}

// RoleName returns a reference to field role_name of azurerm_synapse_role_assignment.
func (sra synapseRoleAssignmentAttributes) RoleName() terra.StringValue {
	return terra.ReferenceAsString(sra.ref.Append("role_name"))
}

// SynapseSparkPoolId returns a reference to field synapse_spark_pool_id of azurerm_synapse_role_assignment.
func (sra synapseRoleAssignmentAttributes) SynapseSparkPoolId() terra.StringValue {
	return terra.ReferenceAsString(sra.ref.Append("synapse_spark_pool_id"))
}

// SynapseWorkspaceId returns a reference to field synapse_workspace_id of azurerm_synapse_role_assignment.
func (sra synapseRoleAssignmentAttributes) SynapseWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sra.ref.Append("synapse_workspace_id"))
}

func (sra synapseRoleAssignmentAttributes) Timeouts() synapseroleassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[synapseroleassignment.TimeoutsAttributes](sra.ref.Append("timeouts"))
}

type synapseRoleAssignmentState struct {
	Id                 string                               `json:"id"`
	PrincipalId        string                               `json:"principal_id"`
	RoleName           string                               `json:"role_name"`
	SynapseSparkPoolId string                               `json:"synapse_spark_pool_id"`
	SynapseWorkspaceId string                               `json:"synapse_workspace_id"`
	Timeouts           *synapseroleassignment.TimeoutsState `json:"timeouts"`
}
