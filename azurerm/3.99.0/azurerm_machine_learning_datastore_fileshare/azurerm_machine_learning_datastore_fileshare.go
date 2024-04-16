// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_machine_learning_datastore_fileshare

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_machine_learning_datastore_fileshare.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermMachineLearningDatastoreFileshareState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (amldf *Resource) Type() string {
	return "azurerm_machine_learning_datastore_fileshare"
}

// LocalName returns the local name for [Resource].
func (amldf *Resource) LocalName() string {
	return amldf.Name
}

// Configuration returns the configuration (args) for [Resource].
func (amldf *Resource) Configuration() interface{} {
	return amldf.Args
}

// DependOn is used for other resources to depend on [Resource].
func (amldf *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(amldf)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (amldf *Resource) Dependencies() terra.Dependencies {
	return amldf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (amldf *Resource) LifecycleManagement() *terra.Lifecycle {
	return amldf.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (amldf *Resource) Attributes() azurermMachineLearningDatastoreFileshareAttributes {
	return azurermMachineLearningDatastoreFileshareAttributes{ref: terra.ReferenceResource(amldf)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (amldf *Resource) ImportState(state io.Reader) error {
	amldf.state = &azurermMachineLearningDatastoreFileshareState{}
	if err := json.NewDecoder(state).Decode(amldf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amldf.Type(), amldf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (amldf *Resource) State() (*azurermMachineLearningDatastoreFileshareState, bool) {
	return amldf.state, amldf.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (amldf *Resource) StateMust() *azurermMachineLearningDatastoreFileshareState {
	if amldf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amldf.Type(), amldf.LocalName()))
	}
	return amldf.state
}

// Args contains the configurations for azurerm_machine_learning_datastore_fileshare.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermMachineLearningDatastoreFileshareAttributes struct {
	ref terra.Reference
}

// AccountKey returns a reference to field account_key of azurerm_machine_learning_datastore_fileshare.
func (amldf azurermMachineLearningDatastoreFileshareAttributes) AccountKey() terra.StringValue {
	return terra.ReferenceAsString(amldf.ref.Append("account_key"))
}

// Description returns a reference to field description of azurerm_machine_learning_datastore_fileshare.
func (amldf azurermMachineLearningDatastoreFileshareAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amldf.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_machine_learning_datastore_fileshare.
func (amldf azurermMachineLearningDatastoreFileshareAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amldf.ref.Append("id"))
}

// IsDefault returns a reference to field is_default of azurerm_machine_learning_datastore_fileshare.
func (amldf azurermMachineLearningDatastoreFileshareAttributes) IsDefault() terra.BoolValue {
	return terra.ReferenceAsBool(amldf.ref.Append("is_default"))
}

// Name returns a reference to field name of azurerm_machine_learning_datastore_fileshare.
func (amldf azurermMachineLearningDatastoreFileshareAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amldf.ref.Append("name"))
}

// ServiceDataIdentity returns a reference to field service_data_identity of azurerm_machine_learning_datastore_fileshare.
func (amldf azurermMachineLearningDatastoreFileshareAttributes) ServiceDataIdentity() terra.StringValue {
	return terra.ReferenceAsString(amldf.ref.Append("service_data_identity"))
}

// SharedAccessSignature returns a reference to field shared_access_signature of azurerm_machine_learning_datastore_fileshare.
func (amldf azurermMachineLearningDatastoreFileshareAttributes) SharedAccessSignature() terra.StringValue {
	return terra.ReferenceAsString(amldf.ref.Append("shared_access_signature"))
}

// StorageFileshareId returns a reference to field storage_fileshare_id of azurerm_machine_learning_datastore_fileshare.
func (amldf azurermMachineLearningDatastoreFileshareAttributes) StorageFileshareId() terra.StringValue {
	return terra.ReferenceAsString(amldf.ref.Append("storage_fileshare_id"))
}

// Tags returns a reference to field tags of azurerm_machine_learning_datastore_fileshare.
func (amldf azurermMachineLearningDatastoreFileshareAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](amldf.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_machine_learning_datastore_fileshare.
func (amldf azurermMachineLearningDatastoreFileshareAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(amldf.ref.Append("workspace_id"))
}

func (amldf azurermMachineLearningDatastoreFileshareAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](amldf.ref.Append("timeouts"))
}

type azurermMachineLearningDatastoreFileshareState struct {
	AccountKey            string            `json:"account_key"`
	Description           string            `json:"description"`
	Id                    string            `json:"id"`
	IsDefault             bool              `json:"is_default"`
	Name                  string            `json:"name"`
	ServiceDataIdentity   string            `json:"service_data_identity"`
	SharedAccessSignature string            `json:"shared_access_signature"`
	StorageFileshareId    string            `json:"storage_fileshare_id"`
	Tags                  map[string]string `json:"tags"`
	WorkspaceId           string            `json:"workspace_id"`
	Timeouts              *TimeoutsState    `json:"timeouts"`
}
