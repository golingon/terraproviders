// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	machinelearningcomputeinstance "github.com/golingon/terraproviders/azurerm/3.49.0/machinelearningcomputeinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMachineLearningComputeInstance creates a new instance of [MachineLearningComputeInstance].
func NewMachineLearningComputeInstance(name string, args MachineLearningComputeInstanceArgs) *MachineLearningComputeInstance {
	return &MachineLearningComputeInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MachineLearningComputeInstance)(nil)

// MachineLearningComputeInstance represents the Terraform resource azurerm_machine_learning_compute_instance.
type MachineLearningComputeInstance struct {
	Name      string
	Args      MachineLearningComputeInstanceArgs
	state     *machineLearningComputeInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MachineLearningComputeInstance].
func (mlci *MachineLearningComputeInstance) Type() string {
	return "azurerm_machine_learning_compute_instance"
}

// LocalName returns the local name for [MachineLearningComputeInstance].
func (mlci *MachineLearningComputeInstance) LocalName() string {
	return mlci.Name
}

// Configuration returns the configuration (args) for [MachineLearningComputeInstance].
func (mlci *MachineLearningComputeInstance) Configuration() interface{} {
	return mlci.Args
}

// DependOn is used for other resources to depend on [MachineLearningComputeInstance].
func (mlci *MachineLearningComputeInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(mlci)
}

// Dependencies returns the list of resources [MachineLearningComputeInstance] depends_on.
func (mlci *MachineLearningComputeInstance) Dependencies() terra.Dependencies {
	return mlci.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MachineLearningComputeInstance].
func (mlci *MachineLearningComputeInstance) LifecycleManagement() *terra.Lifecycle {
	return mlci.Lifecycle
}

// Attributes returns the attributes for [MachineLearningComputeInstance].
func (mlci *MachineLearningComputeInstance) Attributes() machineLearningComputeInstanceAttributes {
	return machineLearningComputeInstanceAttributes{ref: terra.ReferenceResource(mlci)}
}

// ImportState imports the given attribute values into [MachineLearningComputeInstance]'s state.
func (mlci *MachineLearningComputeInstance) ImportState(av io.Reader) error {
	mlci.state = &machineLearningComputeInstanceState{}
	if err := json.NewDecoder(av).Decode(mlci.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mlci.Type(), mlci.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MachineLearningComputeInstance] has state.
func (mlci *MachineLearningComputeInstance) State() (*machineLearningComputeInstanceState, bool) {
	return mlci.state, mlci.state != nil
}

// StateMust returns the state for [MachineLearningComputeInstance]. Panics if the state is nil.
func (mlci *MachineLearningComputeInstance) StateMust() *machineLearningComputeInstanceState {
	if mlci.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mlci.Type(), mlci.LocalName()))
	}
	return mlci.state
}

// MachineLearningComputeInstanceArgs contains the configurations for azurerm_machine_learning_compute_instance.
type MachineLearningComputeInstanceArgs struct {
	// AuthorizationType: string, optional
	AuthorizationType terra.StringValue `hcl:"authorization_type,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocalAuthEnabled: bool, optional
	LocalAuthEnabled terra.BoolValue `hcl:"local_auth_enabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MachineLearningWorkspaceId: string, required
	MachineLearningWorkspaceId terra.StringValue `hcl:"machine_learning_workspace_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SubnetResourceId: string, optional
	SubnetResourceId terra.StringValue `hcl:"subnet_resource_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VirtualMachineSize: string, required
	VirtualMachineSize terra.StringValue `hcl:"virtual_machine_size,attr" validate:"required"`
	// AssignToUser: optional
	AssignToUser *machinelearningcomputeinstance.AssignToUser `hcl:"assign_to_user,block"`
	// Identity: optional
	Identity *machinelearningcomputeinstance.Identity `hcl:"identity,block"`
	// Ssh: optional
	Ssh *machinelearningcomputeinstance.Ssh `hcl:"ssh,block"`
	// Timeouts: optional
	Timeouts *machinelearningcomputeinstance.Timeouts `hcl:"timeouts,block"`
}
type machineLearningComputeInstanceAttributes struct {
	ref terra.Reference
}

// AuthorizationType returns a reference to field authorization_type of azurerm_machine_learning_compute_instance.
func (mlci machineLearningComputeInstanceAttributes) AuthorizationType() terra.StringValue {
	return terra.ReferenceAsString(mlci.ref.Append("authorization_type"))
}

// Description returns a reference to field description of azurerm_machine_learning_compute_instance.
func (mlci machineLearningComputeInstanceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mlci.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_machine_learning_compute_instance.
func (mlci machineLearningComputeInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mlci.ref.Append("id"))
}

// LocalAuthEnabled returns a reference to field local_auth_enabled of azurerm_machine_learning_compute_instance.
func (mlci machineLearningComputeInstanceAttributes) LocalAuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mlci.ref.Append("local_auth_enabled"))
}

// Location returns a reference to field location of azurerm_machine_learning_compute_instance.
func (mlci machineLearningComputeInstanceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mlci.ref.Append("location"))
}

// MachineLearningWorkspaceId returns a reference to field machine_learning_workspace_id of azurerm_machine_learning_compute_instance.
func (mlci machineLearningComputeInstanceAttributes) MachineLearningWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(mlci.ref.Append("machine_learning_workspace_id"))
}

// Name returns a reference to field name of azurerm_machine_learning_compute_instance.
func (mlci machineLearningComputeInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mlci.ref.Append("name"))
}

// SubnetResourceId returns a reference to field subnet_resource_id of azurerm_machine_learning_compute_instance.
func (mlci machineLearningComputeInstanceAttributes) SubnetResourceId() terra.StringValue {
	return terra.ReferenceAsString(mlci.ref.Append("subnet_resource_id"))
}

// Tags returns a reference to field tags of azurerm_machine_learning_compute_instance.
func (mlci machineLearningComputeInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mlci.ref.Append("tags"))
}

// VirtualMachineSize returns a reference to field virtual_machine_size of azurerm_machine_learning_compute_instance.
func (mlci machineLearningComputeInstanceAttributes) VirtualMachineSize() terra.StringValue {
	return terra.ReferenceAsString(mlci.ref.Append("virtual_machine_size"))
}

func (mlci machineLearningComputeInstanceAttributes) AssignToUser() terra.ListValue[machinelearningcomputeinstance.AssignToUserAttributes] {
	return terra.ReferenceAsList[machinelearningcomputeinstance.AssignToUserAttributes](mlci.ref.Append("assign_to_user"))
}

func (mlci machineLearningComputeInstanceAttributes) Identity() terra.ListValue[machinelearningcomputeinstance.IdentityAttributes] {
	return terra.ReferenceAsList[machinelearningcomputeinstance.IdentityAttributes](mlci.ref.Append("identity"))
}

func (mlci machineLearningComputeInstanceAttributes) Ssh() terra.ListValue[machinelearningcomputeinstance.SshAttributes] {
	return terra.ReferenceAsList[machinelearningcomputeinstance.SshAttributes](mlci.ref.Append("ssh"))
}

func (mlci machineLearningComputeInstanceAttributes) Timeouts() machinelearningcomputeinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[machinelearningcomputeinstance.TimeoutsAttributes](mlci.ref.Append("timeouts"))
}

type machineLearningComputeInstanceState struct {
	AuthorizationType          string                                             `json:"authorization_type"`
	Description                string                                             `json:"description"`
	Id                         string                                             `json:"id"`
	LocalAuthEnabled           bool                                               `json:"local_auth_enabled"`
	Location                   string                                             `json:"location"`
	MachineLearningWorkspaceId string                                             `json:"machine_learning_workspace_id"`
	Name                       string                                             `json:"name"`
	SubnetResourceId           string                                             `json:"subnet_resource_id"`
	Tags                       map[string]string                                  `json:"tags"`
	VirtualMachineSize         string                                             `json:"virtual_machine_size"`
	AssignToUser               []machinelearningcomputeinstance.AssignToUserState `json:"assign_to_user"`
	Identity                   []machinelearningcomputeinstance.IdentityState     `json:"identity"`
	Ssh                        []machinelearningcomputeinstance.SshState          `json:"ssh"`
	Timeouts                   *machinelearningcomputeinstance.TimeoutsState      `json:"timeouts"`
}
