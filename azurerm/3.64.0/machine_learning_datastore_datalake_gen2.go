// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	machinelearningdatastoredatalakegen2 "github.com/golingon/terraproviders/azurerm/3.64.0/machinelearningdatastoredatalakegen2"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMachineLearningDatastoreDatalakeGen2 creates a new instance of [MachineLearningDatastoreDatalakeGen2].
func NewMachineLearningDatastoreDatalakeGen2(name string, args MachineLearningDatastoreDatalakeGen2Args) *MachineLearningDatastoreDatalakeGen2 {
	return &MachineLearningDatastoreDatalakeGen2{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MachineLearningDatastoreDatalakeGen2)(nil)

// MachineLearningDatastoreDatalakeGen2 represents the Terraform resource azurerm_machine_learning_datastore_datalake_gen2.
type MachineLearningDatastoreDatalakeGen2 struct {
	Name      string
	Args      MachineLearningDatastoreDatalakeGen2Args
	state     *machineLearningDatastoreDatalakeGen2State
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MachineLearningDatastoreDatalakeGen2].
func (mlddg *MachineLearningDatastoreDatalakeGen2) Type() string {
	return "azurerm_machine_learning_datastore_datalake_gen2"
}

// LocalName returns the local name for [MachineLearningDatastoreDatalakeGen2].
func (mlddg *MachineLearningDatastoreDatalakeGen2) LocalName() string {
	return mlddg.Name
}

// Configuration returns the configuration (args) for [MachineLearningDatastoreDatalakeGen2].
func (mlddg *MachineLearningDatastoreDatalakeGen2) Configuration() interface{} {
	return mlddg.Args
}

// DependOn is used for other resources to depend on [MachineLearningDatastoreDatalakeGen2].
func (mlddg *MachineLearningDatastoreDatalakeGen2) DependOn() terra.Reference {
	return terra.ReferenceResource(mlddg)
}

// Dependencies returns the list of resources [MachineLearningDatastoreDatalakeGen2] depends_on.
func (mlddg *MachineLearningDatastoreDatalakeGen2) Dependencies() terra.Dependencies {
	return mlddg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MachineLearningDatastoreDatalakeGen2].
func (mlddg *MachineLearningDatastoreDatalakeGen2) LifecycleManagement() *terra.Lifecycle {
	return mlddg.Lifecycle
}

// Attributes returns the attributes for [MachineLearningDatastoreDatalakeGen2].
func (mlddg *MachineLearningDatastoreDatalakeGen2) Attributes() machineLearningDatastoreDatalakeGen2Attributes {
	return machineLearningDatastoreDatalakeGen2Attributes{ref: terra.ReferenceResource(mlddg)}
}

// ImportState imports the given attribute values into [MachineLearningDatastoreDatalakeGen2]'s state.
func (mlddg *MachineLearningDatastoreDatalakeGen2) ImportState(av io.Reader) error {
	mlddg.state = &machineLearningDatastoreDatalakeGen2State{}
	if err := json.NewDecoder(av).Decode(mlddg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mlddg.Type(), mlddg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MachineLearningDatastoreDatalakeGen2] has state.
func (mlddg *MachineLearningDatastoreDatalakeGen2) State() (*machineLearningDatastoreDatalakeGen2State, bool) {
	return mlddg.state, mlddg.state != nil
}

// StateMust returns the state for [MachineLearningDatastoreDatalakeGen2]. Panics if the state is nil.
func (mlddg *MachineLearningDatastoreDatalakeGen2) StateMust() *machineLearningDatastoreDatalakeGen2State {
	if mlddg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mlddg.Type(), mlddg.LocalName()))
	}
	return mlddg.state
}

// MachineLearningDatastoreDatalakeGen2Args contains the configurations for azurerm_machine_learning_datastore_datalake_gen2.
type MachineLearningDatastoreDatalakeGen2Args struct {
	// AuthorityUrl: string, optional
	AuthorityUrl terra.StringValue `hcl:"authority_url,attr"`
	// ClientId: string, optional
	ClientId terra.StringValue `hcl:"client_id,attr"`
	// ClientSecret: string, optional
	ClientSecret terra.StringValue `hcl:"client_secret,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ServiceDataIdentity: string, optional
	ServiceDataIdentity terra.StringValue `hcl:"service_data_identity,attr"`
	// StorageContainerId: string, required
	StorageContainerId terra.StringValue `hcl:"storage_container_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *machinelearningdatastoredatalakegen2.Timeouts `hcl:"timeouts,block"`
}
type machineLearningDatastoreDatalakeGen2Attributes struct {
	ref terra.Reference
}

// AuthorityUrl returns a reference to field authority_url of azurerm_machine_learning_datastore_datalake_gen2.
func (mlddg machineLearningDatastoreDatalakeGen2Attributes) AuthorityUrl() terra.StringValue {
	return terra.ReferenceAsString(mlddg.ref.Append("authority_url"))
}

// ClientId returns a reference to field client_id of azurerm_machine_learning_datastore_datalake_gen2.
func (mlddg machineLearningDatastoreDatalakeGen2Attributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(mlddg.ref.Append("client_id"))
}

// ClientSecret returns a reference to field client_secret of azurerm_machine_learning_datastore_datalake_gen2.
func (mlddg machineLearningDatastoreDatalakeGen2Attributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(mlddg.ref.Append("client_secret"))
}

// Description returns a reference to field description of azurerm_machine_learning_datastore_datalake_gen2.
func (mlddg machineLearningDatastoreDatalakeGen2Attributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mlddg.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_machine_learning_datastore_datalake_gen2.
func (mlddg machineLearningDatastoreDatalakeGen2Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mlddg.ref.Append("id"))
}

// IsDefault returns a reference to field is_default of azurerm_machine_learning_datastore_datalake_gen2.
func (mlddg machineLearningDatastoreDatalakeGen2Attributes) IsDefault() terra.BoolValue {
	return terra.ReferenceAsBool(mlddg.ref.Append("is_default"))
}

// Name returns a reference to field name of azurerm_machine_learning_datastore_datalake_gen2.
func (mlddg machineLearningDatastoreDatalakeGen2Attributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mlddg.ref.Append("name"))
}

// ServiceDataIdentity returns a reference to field service_data_identity of azurerm_machine_learning_datastore_datalake_gen2.
func (mlddg machineLearningDatastoreDatalakeGen2Attributes) ServiceDataIdentity() terra.StringValue {
	return terra.ReferenceAsString(mlddg.ref.Append("service_data_identity"))
}

// StorageContainerId returns a reference to field storage_container_id of azurerm_machine_learning_datastore_datalake_gen2.
func (mlddg machineLearningDatastoreDatalakeGen2Attributes) StorageContainerId() terra.StringValue {
	return terra.ReferenceAsString(mlddg.ref.Append("storage_container_id"))
}

// Tags returns a reference to field tags of azurerm_machine_learning_datastore_datalake_gen2.
func (mlddg machineLearningDatastoreDatalakeGen2Attributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mlddg.ref.Append("tags"))
}

// TenantId returns a reference to field tenant_id of azurerm_machine_learning_datastore_datalake_gen2.
func (mlddg machineLearningDatastoreDatalakeGen2Attributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(mlddg.ref.Append("tenant_id"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_machine_learning_datastore_datalake_gen2.
func (mlddg machineLearningDatastoreDatalakeGen2Attributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(mlddg.ref.Append("workspace_id"))
}

func (mlddg machineLearningDatastoreDatalakeGen2Attributes) Timeouts() machinelearningdatastoredatalakegen2.TimeoutsAttributes {
	return terra.ReferenceAsSingle[machinelearningdatastoredatalakegen2.TimeoutsAttributes](mlddg.ref.Append("timeouts"))
}

type machineLearningDatastoreDatalakeGen2State struct {
	AuthorityUrl        string                                              `json:"authority_url"`
	ClientId            string                                              `json:"client_id"`
	ClientSecret        string                                              `json:"client_secret"`
	Description         string                                              `json:"description"`
	Id                  string                                              `json:"id"`
	IsDefault           bool                                                `json:"is_default"`
	Name                string                                              `json:"name"`
	ServiceDataIdentity string                                              `json:"service_data_identity"`
	StorageContainerId  string                                              `json:"storage_container_id"`
	Tags                map[string]string                                   `json:"tags"`
	TenantId            string                                              `json:"tenant_id"`
	WorkspaceId         string                                              `json:"workspace_id"`
	Timeouts            *machinelearningdatastoredatalakegen2.TimeoutsState `json:"timeouts"`
}
