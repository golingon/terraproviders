// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	machinelearningdatastoreblobstorage "github.com/golingon/terraproviders/azurerm/3.66.0/machinelearningdatastoreblobstorage"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMachineLearningDatastoreBlobstorage creates a new instance of [MachineLearningDatastoreBlobstorage].
func NewMachineLearningDatastoreBlobstorage(name string, args MachineLearningDatastoreBlobstorageArgs) *MachineLearningDatastoreBlobstorage {
	return &MachineLearningDatastoreBlobstorage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MachineLearningDatastoreBlobstorage)(nil)

// MachineLearningDatastoreBlobstorage represents the Terraform resource azurerm_machine_learning_datastore_blobstorage.
type MachineLearningDatastoreBlobstorage struct {
	Name      string
	Args      MachineLearningDatastoreBlobstorageArgs
	state     *machineLearningDatastoreBlobstorageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MachineLearningDatastoreBlobstorage].
func (mldb *MachineLearningDatastoreBlobstorage) Type() string {
	return "azurerm_machine_learning_datastore_blobstorage"
}

// LocalName returns the local name for [MachineLearningDatastoreBlobstorage].
func (mldb *MachineLearningDatastoreBlobstorage) LocalName() string {
	return mldb.Name
}

// Configuration returns the configuration (args) for [MachineLearningDatastoreBlobstorage].
func (mldb *MachineLearningDatastoreBlobstorage) Configuration() interface{} {
	return mldb.Args
}

// DependOn is used for other resources to depend on [MachineLearningDatastoreBlobstorage].
func (mldb *MachineLearningDatastoreBlobstorage) DependOn() terra.Reference {
	return terra.ReferenceResource(mldb)
}

// Dependencies returns the list of resources [MachineLearningDatastoreBlobstorage] depends_on.
func (mldb *MachineLearningDatastoreBlobstorage) Dependencies() terra.Dependencies {
	return mldb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MachineLearningDatastoreBlobstorage].
func (mldb *MachineLearningDatastoreBlobstorage) LifecycleManagement() *terra.Lifecycle {
	return mldb.Lifecycle
}

// Attributes returns the attributes for [MachineLearningDatastoreBlobstorage].
func (mldb *MachineLearningDatastoreBlobstorage) Attributes() machineLearningDatastoreBlobstorageAttributes {
	return machineLearningDatastoreBlobstorageAttributes{ref: terra.ReferenceResource(mldb)}
}

// ImportState imports the given attribute values into [MachineLearningDatastoreBlobstorage]'s state.
func (mldb *MachineLearningDatastoreBlobstorage) ImportState(av io.Reader) error {
	mldb.state = &machineLearningDatastoreBlobstorageState{}
	if err := json.NewDecoder(av).Decode(mldb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mldb.Type(), mldb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MachineLearningDatastoreBlobstorage] has state.
func (mldb *MachineLearningDatastoreBlobstorage) State() (*machineLearningDatastoreBlobstorageState, bool) {
	return mldb.state, mldb.state != nil
}

// StateMust returns the state for [MachineLearningDatastoreBlobstorage]. Panics if the state is nil.
func (mldb *MachineLearningDatastoreBlobstorage) StateMust() *machineLearningDatastoreBlobstorageState {
	if mldb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mldb.Type(), mldb.LocalName()))
	}
	return mldb.state
}

// MachineLearningDatastoreBlobstorageArgs contains the configurations for azurerm_machine_learning_datastore_blobstorage.
type MachineLearningDatastoreBlobstorageArgs struct {
	// AccountKey: string, optional
	AccountKey terra.StringValue `hcl:"account_key,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsDefault: bool, optional
	IsDefault terra.BoolValue `hcl:"is_default,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ServiceDataAuthIdentity: string, optional
	ServiceDataAuthIdentity terra.StringValue `hcl:"service_data_auth_identity,attr"`
	// SharedAccessSignature: string, optional
	SharedAccessSignature terra.StringValue `hcl:"shared_access_signature,attr"`
	// StorageContainerId: string, required
	StorageContainerId terra.StringValue `hcl:"storage_container_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *machinelearningdatastoreblobstorage.Timeouts `hcl:"timeouts,block"`
}
type machineLearningDatastoreBlobstorageAttributes struct {
	ref terra.Reference
}

// AccountKey returns a reference to field account_key of azurerm_machine_learning_datastore_blobstorage.
func (mldb machineLearningDatastoreBlobstorageAttributes) AccountKey() terra.StringValue {
	return terra.ReferenceAsString(mldb.ref.Append("account_key"))
}

// Description returns a reference to field description of azurerm_machine_learning_datastore_blobstorage.
func (mldb machineLearningDatastoreBlobstorageAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mldb.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_machine_learning_datastore_blobstorage.
func (mldb machineLearningDatastoreBlobstorageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mldb.ref.Append("id"))
}

// IsDefault returns a reference to field is_default of azurerm_machine_learning_datastore_blobstorage.
func (mldb machineLearningDatastoreBlobstorageAttributes) IsDefault() terra.BoolValue {
	return terra.ReferenceAsBool(mldb.ref.Append("is_default"))
}

// Name returns a reference to field name of azurerm_machine_learning_datastore_blobstorage.
func (mldb machineLearningDatastoreBlobstorageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mldb.ref.Append("name"))
}

// ServiceDataAuthIdentity returns a reference to field service_data_auth_identity of azurerm_machine_learning_datastore_blobstorage.
func (mldb machineLearningDatastoreBlobstorageAttributes) ServiceDataAuthIdentity() terra.StringValue {
	return terra.ReferenceAsString(mldb.ref.Append("service_data_auth_identity"))
}

// SharedAccessSignature returns a reference to field shared_access_signature of azurerm_machine_learning_datastore_blobstorage.
func (mldb machineLearningDatastoreBlobstorageAttributes) SharedAccessSignature() terra.StringValue {
	return terra.ReferenceAsString(mldb.ref.Append("shared_access_signature"))
}

// StorageContainerId returns a reference to field storage_container_id of azurerm_machine_learning_datastore_blobstorage.
func (mldb machineLearningDatastoreBlobstorageAttributes) StorageContainerId() terra.StringValue {
	return terra.ReferenceAsString(mldb.ref.Append("storage_container_id"))
}

// Tags returns a reference to field tags of azurerm_machine_learning_datastore_blobstorage.
func (mldb machineLearningDatastoreBlobstorageAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mldb.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_machine_learning_datastore_blobstorage.
func (mldb machineLearningDatastoreBlobstorageAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(mldb.ref.Append("workspace_id"))
}

func (mldb machineLearningDatastoreBlobstorageAttributes) Timeouts() machinelearningdatastoreblobstorage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[machinelearningdatastoreblobstorage.TimeoutsAttributes](mldb.ref.Append("timeouts"))
}

type machineLearningDatastoreBlobstorageState struct {
	AccountKey              string                                             `json:"account_key"`
	Description             string                                             `json:"description"`
	Id                      string                                             `json:"id"`
	IsDefault               bool                                               `json:"is_default"`
	Name                    string                                             `json:"name"`
	ServiceDataAuthIdentity string                                             `json:"service_data_auth_identity"`
	SharedAccessSignature   string                                             `json:"shared_access_signature"`
	StorageContainerId      string                                             `json:"storage_container_id"`
	Tags                    map[string]string                                  `json:"tags"`
	WorkspaceId             string                                             `json:"workspace_id"`
	Timeouts                *machinelearningdatastoreblobstorage.TimeoutsState `json:"timeouts"`
}
