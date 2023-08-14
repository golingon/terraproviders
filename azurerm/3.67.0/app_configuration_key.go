// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appconfigurationkey "github.com/golingon/terraproviders/azurerm/3.67.0/appconfigurationkey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppConfigurationKey creates a new instance of [AppConfigurationKey].
func NewAppConfigurationKey(name string, args AppConfigurationKeyArgs) *AppConfigurationKey {
	return &AppConfigurationKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppConfigurationKey)(nil)

// AppConfigurationKey represents the Terraform resource azurerm_app_configuration_key.
type AppConfigurationKey struct {
	Name      string
	Args      AppConfigurationKeyArgs
	state     *appConfigurationKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppConfigurationKey].
func (ack *AppConfigurationKey) Type() string {
	return "azurerm_app_configuration_key"
}

// LocalName returns the local name for [AppConfigurationKey].
func (ack *AppConfigurationKey) LocalName() string {
	return ack.Name
}

// Configuration returns the configuration (args) for [AppConfigurationKey].
func (ack *AppConfigurationKey) Configuration() interface{} {
	return ack.Args
}

// DependOn is used for other resources to depend on [AppConfigurationKey].
func (ack *AppConfigurationKey) DependOn() terra.Reference {
	return terra.ReferenceResource(ack)
}

// Dependencies returns the list of resources [AppConfigurationKey] depends_on.
func (ack *AppConfigurationKey) Dependencies() terra.Dependencies {
	return ack.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppConfigurationKey].
func (ack *AppConfigurationKey) LifecycleManagement() *terra.Lifecycle {
	return ack.Lifecycle
}

// Attributes returns the attributes for [AppConfigurationKey].
func (ack *AppConfigurationKey) Attributes() appConfigurationKeyAttributes {
	return appConfigurationKeyAttributes{ref: terra.ReferenceResource(ack)}
}

// ImportState imports the given attribute values into [AppConfigurationKey]'s state.
func (ack *AppConfigurationKey) ImportState(av io.Reader) error {
	ack.state = &appConfigurationKeyState{}
	if err := json.NewDecoder(av).Decode(ack.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ack.Type(), ack.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppConfigurationKey] has state.
func (ack *AppConfigurationKey) State() (*appConfigurationKeyState, bool) {
	return ack.state, ack.state != nil
}

// StateMust returns the state for [AppConfigurationKey]. Panics if the state is nil.
func (ack *AppConfigurationKey) StateMust() *appConfigurationKeyState {
	if ack.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ack.Type(), ack.LocalName()))
	}
	return ack.state
}

// AppConfigurationKeyArgs contains the configurations for azurerm_app_configuration_key.
type AppConfigurationKeyArgs struct {
	// ConfigurationStoreId: string, required
	ConfigurationStoreId terra.StringValue `hcl:"configuration_store_id,attr" validate:"required"`
	// ContentType: string, optional
	ContentType terra.StringValue `hcl:"content_type,attr"`
	// Etag: string, optional
	Etag terra.StringValue `hcl:"etag,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// Locked: bool, optional
	Locked terra.BoolValue `hcl:"locked,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
	// VaultKeyReference: string, optional
	VaultKeyReference terra.StringValue `hcl:"vault_key_reference,attr"`
	// Timeouts: optional
	Timeouts *appconfigurationkey.Timeouts `hcl:"timeouts,block"`
}
type appConfigurationKeyAttributes struct {
	ref terra.Reference
}

// ConfigurationStoreId returns a reference to field configuration_store_id of azurerm_app_configuration_key.
func (ack appConfigurationKeyAttributes) ConfigurationStoreId() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("configuration_store_id"))
}

// ContentType returns a reference to field content_type of azurerm_app_configuration_key.
func (ack appConfigurationKeyAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("content_type"))
}

// Etag returns a reference to field etag of azurerm_app_configuration_key.
func (ack appConfigurationKeyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("etag"))
}

// Id returns a reference to field id of azurerm_app_configuration_key.
func (ack appConfigurationKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("id"))
}

// Key returns a reference to field key of azurerm_app_configuration_key.
func (ack appConfigurationKeyAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("key"))
}

// Label returns a reference to field label of azurerm_app_configuration_key.
func (ack appConfigurationKeyAttributes) Label() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("label"))
}

// Locked returns a reference to field locked of azurerm_app_configuration_key.
func (ack appConfigurationKeyAttributes) Locked() terra.BoolValue {
	return terra.ReferenceAsBool(ack.ref.Append("locked"))
}

// Tags returns a reference to field tags of azurerm_app_configuration_key.
func (ack appConfigurationKeyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ack.ref.Append("tags"))
}

// Type returns a reference to field type of azurerm_app_configuration_key.
func (ack appConfigurationKeyAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("type"))
}

// Value returns a reference to field value of azurerm_app_configuration_key.
func (ack appConfigurationKeyAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("value"))
}

// VaultKeyReference returns a reference to field vault_key_reference of azurerm_app_configuration_key.
func (ack appConfigurationKeyAttributes) VaultKeyReference() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("vault_key_reference"))
}

func (ack appConfigurationKeyAttributes) Timeouts() appconfigurationkey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appconfigurationkey.TimeoutsAttributes](ack.ref.Append("timeouts"))
}

type appConfigurationKeyState struct {
	ConfigurationStoreId string                             `json:"configuration_store_id"`
	ContentType          string                             `json:"content_type"`
	Etag                 string                             `json:"etag"`
	Id                   string                             `json:"id"`
	Key                  string                             `json:"key"`
	Label                string                             `json:"label"`
	Locked               bool                               `json:"locked"`
	Tags                 map[string]string                  `json:"tags"`
	Type                 string                             `json:"type"`
	Value                string                             `json:"value"`
	VaultKeyReference    string                             `json:"vault_key_reference"`
	Timeouts             *appconfigurationkey.TimeoutsState `json:"timeouts"`
}