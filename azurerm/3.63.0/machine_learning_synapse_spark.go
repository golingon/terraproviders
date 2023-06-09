// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	machinelearningsynapsespark "github.com/golingon/terraproviders/azurerm/3.63.0/machinelearningsynapsespark"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMachineLearningSynapseSpark creates a new instance of [MachineLearningSynapseSpark].
func NewMachineLearningSynapseSpark(name string, args MachineLearningSynapseSparkArgs) *MachineLearningSynapseSpark {
	return &MachineLearningSynapseSpark{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MachineLearningSynapseSpark)(nil)

// MachineLearningSynapseSpark represents the Terraform resource azurerm_machine_learning_synapse_spark.
type MachineLearningSynapseSpark struct {
	Name      string
	Args      MachineLearningSynapseSparkArgs
	state     *machineLearningSynapseSparkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MachineLearningSynapseSpark].
func (mlss *MachineLearningSynapseSpark) Type() string {
	return "azurerm_machine_learning_synapse_spark"
}

// LocalName returns the local name for [MachineLearningSynapseSpark].
func (mlss *MachineLearningSynapseSpark) LocalName() string {
	return mlss.Name
}

// Configuration returns the configuration (args) for [MachineLearningSynapseSpark].
func (mlss *MachineLearningSynapseSpark) Configuration() interface{} {
	return mlss.Args
}

// DependOn is used for other resources to depend on [MachineLearningSynapseSpark].
func (mlss *MachineLearningSynapseSpark) DependOn() terra.Reference {
	return terra.ReferenceResource(mlss)
}

// Dependencies returns the list of resources [MachineLearningSynapseSpark] depends_on.
func (mlss *MachineLearningSynapseSpark) Dependencies() terra.Dependencies {
	return mlss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MachineLearningSynapseSpark].
func (mlss *MachineLearningSynapseSpark) LifecycleManagement() *terra.Lifecycle {
	return mlss.Lifecycle
}

// Attributes returns the attributes for [MachineLearningSynapseSpark].
func (mlss *MachineLearningSynapseSpark) Attributes() machineLearningSynapseSparkAttributes {
	return machineLearningSynapseSparkAttributes{ref: terra.ReferenceResource(mlss)}
}

// ImportState imports the given attribute values into [MachineLearningSynapseSpark]'s state.
func (mlss *MachineLearningSynapseSpark) ImportState(av io.Reader) error {
	mlss.state = &machineLearningSynapseSparkState{}
	if err := json.NewDecoder(av).Decode(mlss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mlss.Type(), mlss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MachineLearningSynapseSpark] has state.
func (mlss *MachineLearningSynapseSpark) State() (*machineLearningSynapseSparkState, bool) {
	return mlss.state, mlss.state != nil
}

// StateMust returns the state for [MachineLearningSynapseSpark]. Panics if the state is nil.
func (mlss *MachineLearningSynapseSpark) StateMust() *machineLearningSynapseSparkState {
	if mlss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mlss.Type(), mlss.LocalName()))
	}
	return mlss.state
}

// MachineLearningSynapseSparkArgs contains the configurations for azurerm_machine_learning_synapse_spark.
type MachineLearningSynapseSparkArgs struct {
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
	// SynapseSparkPoolId: string, required
	SynapseSparkPoolId terra.StringValue `hcl:"synapse_spark_pool_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Identity: optional
	Identity *machinelearningsynapsespark.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *machinelearningsynapsespark.Timeouts `hcl:"timeouts,block"`
}
type machineLearningSynapseSparkAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_machine_learning_synapse_spark.
func (mlss machineLearningSynapseSparkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mlss.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_machine_learning_synapse_spark.
func (mlss machineLearningSynapseSparkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mlss.ref.Append("id"))
}

// LocalAuthEnabled returns a reference to field local_auth_enabled of azurerm_machine_learning_synapse_spark.
func (mlss machineLearningSynapseSparkAttributes) LocalAuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mlss.ref.Append("local_auth_enabled"))
}

// Location returns a reference to field location of azurerm_machine_learning_synapse_spark.
func (mlss machineLearningSynapseSparkAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mlss.ref.Append("location"))
}

// MachineLearningWorkspaceId returns a reference to field machine_learning_workspace_id of azurerm_machine_learning_synapse_spark.
func (mlss machineLearningSynapseSparkAttributes) MachineLearningWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(mlss.ref.Append("machine_learning_workspace_id"))
}

// Name returns a reference to field name of azurerm_machine_learning_synapse_spark.
func (mlss machineLearningSynapseSparkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mlss.ref.Append("name"))
}

// SynapseSparkPoolId returns a reference to field synapse_spark_pool_id of azurerm_machine_learning_synapse_spark.
func (mlss machineLearningSynapseSparkAttributes) SynapseSparkPoolId() terra.StringValue {
	return terra.ReferenceAsString(mlss.ref.Append("synapse_spark_pool_id"))
}

// Tags returns a reference to field tags of azurerm_machine_learning_synapse_spark.
func (mlss machineLearningSynapseSparkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mlss.ref.Append("tags"))
}

func (mlss machineLearningSynapseSparkAttributes) Identity() terra.ListValue[machinelearningsynapsespark.IdentityAttributes] {
	return terra.ReferenceAsList[machinelearningsynapsespark.IdentityAttributes](mlss.ref.Append("identity"))
}

func (mlss machineLearningSynapseSparkAttributes) Timeouts() machinelearningsynapsespark.TimeoutsAttributes {
	return terra.ReferenceAsSingle[machinelearningsynapsespark.TimeoutsAttributes](mlss.ref.Append("timeouts"))
}

type machineLearningSynapseSparkState struct {
	Description                string                                      `json:"description"`
	Id                         string                                      `json:"id"`
	LocalAuthEnabled           bool                                        `json:"local_auth_enabled"`
	Location                   string                                      `json:"location"`
	MachineLearningWorkspaceId string                                      `json:"machine_learning_workspace_id"`
	Name                       string                                      `json:"name"`
	SynapseSparkPoolId         string                                      `json:"synapse_spark_pool_id"`
	Tags                       map[string]string                           `json:"tags"`
	Identity                   []machinelearningsynapsespark.IdentityState `json:"identity"`
	Timeouts                   *machinelearningsynapsespark.TimeoutsState  `json:"timeouts"`
}
