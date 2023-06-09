// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	machinelearningdatastorefileshare "github.com/golingon/terraproviders/azurerm/3.52.0/machinelearningdatastorefileshare"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMachineLearningDatastoreFileshare creates a new instance of [MachineLearningDatastoreFileshare].
func NewMachineLearningDatastoreFileshare(name string, args MachineLearningDatastoreFileshareArgs) *MachineLearningDatastoreFileshare {
	return &MachineLearningDatastoreFileshare{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MachineLearningDatastoreFileshare)(nil)

// MachineLearningDatastoreFileshare represents the Terraform resource azurerm_machine_learning_datastore_fileshare.
type MachineLearningDatastoreFileshare struct {
	Name      string
	Args      MachineLearningDatastoreFileshareArgs
	state     *machineLearningDatastoreFileshareState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MachineLearningDatastoreFileshare].
func (mldf *MachineLearningDatastoreFileshare) Type() string {
	return "azurerm_machine_learning_datastore_fileshare"
}

// LocalName returns the local name for [MachineLearningDatastoreFileshare].
func (mldf *MachineLearningDatastoreFileshare) LocalName() string {
	return mldf.Name
}

// Configuration returns the configuration (args) for [MachineLearningDatastoreFileshare].
func (mldf *MachineLearningDatastoreFileshare) Configuration() interface{} {
	return mldf.Args
}

// DependOn is used for other resources to depend on [MachineLearningDatastoreFileshare].
func (mldf *MachineLearningDatastoreFileshare) DependOn() terra.Reference {
	return terra.ReferenceResource(mldf)
}

// Dependencies returns the list of resources [MachineLearningDatastoreFileshare] depends_on.
func (mldf *MachineLearningDatastoreFileshare) Dependencies() terra.Dependencies {
	return mldf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MachineLearningDatastoreFileshare].
func (mldf *MachineLearningDatastoreFileshare) LifecycleManagement() *terra.Lifecycle {
	return mldf.Lifecycle
}

// Attributes returns the attributes for [MachineLearningDatastoreFileshare].
func (mldf *MachineLearningDatastoreFileshare) Attributes() machineLearningDatastoreFileshareAttributes {
	return machineLearningDatastoreFileshareAttributes{ref: terra.ReferenceResource(mldf)}
}

// ImportState imports the given attribute values into [MachineLearningDatastoreFileshare]'s state.
func (mldf *MachineLearningDatastoreFileshare) ImportState(av io.Reader) error {
	mldf.state = &machineLearningDatastoreFileshareState{}
	if err := json.NewDecoder(av).Decode(mldf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mldf.Type(), mldf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MachineLearningDatastoreFileshare] has state.
func (mldf *MachineLearningDatastoreFileshare) State() (*machineLearningDatastoreFileshareState, bool) {
	return mldf.state, mldf.state != nil
}

// StateMust returns the state for [MachineLearningDatastoreFileshare]. Panics if the state is nil.
func (mldf *MachineLearningDatastoreFileshare) StateMust() *machineLearningDatastoreFileshareState {
	if mldf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mldf.Type(), mldf.LocalName()))
	}
	return mldf.state
}

// MachineLearningDatastoreFileshareArgs contains the configurations for azurerm_machine_learning_datastore_fileshare.
type MachineLearningDatastoreFileshareArgs struct {
	// AccountKey: string, optional
	AccountKey terra.StringValue `hcl:"account_key,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ServiceDataIdentity: string, optional
	ServiceDataIdentity terra.StringValue `hcl:"service_data_identity,attr"`
	// SharedAccessSignature: string, optional
	SharedAccessSignature terra.StringValue `hcl:"shared_access_signature,attr"`
	// StorageFileshareId: string, required
	StorageFileshareId terra.StringValue `hcl:"storage_fileshare_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *machinelearningdatastorefileshare.Timeouts `hcl:"timeouts,block"`
}
type machineLearningDatastoreFileshareAttributes struct {
	ref terra.Reference
}

// AccountKey returns a reference to field account_key of azurerm_machine_learning_datastore_fileshare.
func (mldf machineLearningDatastoreFileshareAttributes) AccountKey() terra.StringValue {
	return terra.ReferenceAsString(mldf.ref.Append("account_key"))
}

// Description returns a reference to field description of azurerm_machine_learning_datastore_fileshare.
func (mldf machineLearningDatastoreFileshareAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mldf.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_machine_learning_datastore_fileshare.
func (mldf machineLearningDatastoreFileshareAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mldf.ref.Append("id"))
}

// IsDefault returns a reference to field is_default of azurerm_machine_learning_datastore_fileshare.
func (mldf machineLearningDatastoreFileshareAttributes) IsDefault() terra.BoolValue {
	return terra.ReferenceAsBool(mldf.ref.Append("is_default"))
}

// Name returns a reference to field name of azurerm_machine_learning_datastore_fileshare.
func (mldf machineLearningDatastoreFileshareAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mldf.ref.Append("name"))
}

// ServiceDataIdentity returns a reference to field service_data_identity of azurerm_machine_learning_datastore_fileshare.
func (mldf machineLearningDatastoreFileshareAttributes) ServiceDataIdentity() terra.StringValue {
	return terra.ReferenceAsString(mldf.ref.Append("service_data_identity"))
}

// SharedAccessSignature returns a reference to field shared_access_signature of azurerm_machine_learning_datastore_fileshare.
func (mldf machineLearningDatastoreFileshareAttributes) SharedAccessSignature() terra.StringValue {
	return terra.ReferenceAsString(mldf.ref.Append("shared_access_signature"))
}

// StorageFileshareId returns a reference to field storage_fileshare_id of azurerm_machine_learning_datastore_fileshare.
func (mldf machineLearningDatastoreFileshareAttributes) StorageFileshareId() terra.StringValue {
	return terra.ReferenceAsString(mldf.ref.Append("storage_fileshare_id"))
}

// Tags returns a reference to field tags of azurerm_machine_learning_datastore_fileshare.
func (mldf machineLearningDatastoreFileshareAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mldf.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_machine_learning_datastore_fileshare.
func (mldf machineLearningDatastoreFileshareAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(mldf.ref.Append("workspace_id"))
}

func (mldf machineLearningDatastoreFileshareAttributes) Timeouts() machinelearningdatastorefileshare.TimeoutsAttributes {
	return terra.ReferenceAsSingle[machinelearningdatastorefileshare.TimeoutsAttributes](mldf.ref.Append("timeouts"))
}

type machineLearningDatastoreFileshareState struct {
	AccountKey            string                                           `json:"account_key"`
	Description           string                                           `json:"description"`
	Id                    string                                           `json:"id"`
	IsDefault             bool                                             `json:"is_default"`
	Name                  string                                           `json:"name"`
	ServiceDataIdentity   string                                           `json:"service_data_identity"`
	SharedAccessSignature string                                           `json:"shared_access_signature"`
	StorageFileshareId    string                                           `json:"storage_fileshare_id"`
	Tags                  map[string]string                                `json:"tags"`
	WorkspaceId           string                                           `json:"workspace_id"`
	Timeouts              *machinelearningdatastorefileshare.TimeoutsState `json:"timeouts"`
}
