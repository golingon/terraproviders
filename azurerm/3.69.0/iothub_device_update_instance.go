// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iothubdeviceupdateinstance "github.com/golingon/terraproviders/azurerm/3.69.0/iothubdeviceupdateinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIothubDeviceUpdateInstance creates a new instance of [IothubDeviceUpdateInstance].
func NewIothubDeviceUpdateInstance(name string, args IothubDeviceUpdateInstanceArgs) *IothubDeviceUpdateInstance {
	return &IothubDeviceUpdateInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IothubDeviceUpdateInstance)(nil)

// IothubDeviceUpdateInstance represents the Terraform resource azurerm_iothub_device_update_instance.
type IothubDeviceUpdateInstance struct {
	Name      string
	Args      IothubDeviceUpdateInstanceArgs
	state     *iothubDeviceUpdateInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IothubDeviceUpdateInstance].
func (idui *IothubDeviceUpdateInstance) Type() string {
	return "azurerm_iothub_device_update_instance"
}

// LocalName returns the local name for [IothubDeviceUpdateInstance].
func (idui *IothubDeviceUpdateInstance) LocalName() string {
	return idui.Name
}

// Configuration returns the configuration (args) for [IothubDeviceUpdateInstance].
func (idui *IothubDeviceUpdateInstance) Configuration() interface{} {
	return idui.Args
}

// DependOn is used for other resources to depend on [IothubDeviceUpdateInstance].
func (idui *IothubDeviceUpdateInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(idui)
}

// Dependencies returns the list of resources [IothubDeviceUpdateInstance] depends_on.
func (idui *IothubDeviceUpdateInstance) Dependencies() terra.Dependencies {
	return idui.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IothubDeviceUpdateInstance].
func (idui *IothubDeviceUpdateInstance) LifecycleManagement() *terra.Lifecycle {
	return idui.Lifecycle
}

// Attributes returns the attributes for [IothubDeviceUpdateInstance].
func (idui *IothubDeviceUpdateInstance) Attributes() iothubDeviceUpdateInstanceAttributes {
	return iothubDeviceUpdateInstanceAttributes{ref: terra.ReferenceResource(idui)}
}

// ImportState imports the given attribute values into [IothubDeviceUpdateInstance]'s state.
func (idui *IothubDeviceUpdateInstance) ImportState(av io.Reader) error {
	idui.state = &iothubDeviceUpdateInstanceState{}
	if err := json.NewDecoder(av).Decode(idui.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", idui.Type(), idui.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IothubDeviceUpdateInstance] has state.
func (idui *IothubDeviceUpdateInstance) State() (*iothubDeviceUpdateInstanceState, bool) {
	return idui.state, idui.state != nil
}

// StateMust returns the state for [IothubDeviceUpdateInstance]. Panics if the state is nil.
func (idui *IothubDeviceUpdateInstance) StateMust() *iothubDeviceUpdateInstanceState {
	if idui.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", idui.Type(), idui.LocalName()))
	}
	return idui.state
}

// IothubDeviceUpdateInstanceArgs contains the configurations for azurerm_iothub_device_update_instance.
type IothubDeviceUpdateInstanceArgs struct {
	// DeviceUpdateAccountId: string, required
	DeviceUpdateAccountId terra.StringValue `hcl:"device_update_account_id,attr" validate:"required"`
	// DiagnosticEnabled: bool, optional
	DiagnosticEnabled terra.BoolValue `hcl:"diagnostic_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IothubId: string, required
	IothubId terra.StringValue `hcl:"iothub_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// DiagnosticStorageAccount: optional
	DiagnosticStorageAccount *iothubdeviceupdateinstance.DiagnosticStorageAccount `hcl:"diagnostic_storage_account,block"`
	// Timeouts: optional
	Timeouts *iothubdeviceupdateinstance.Timeouts `hcl:"timeouts,block"`
}
type iothubDeviceUpdateInstanceAttributes struct {
	ref terra.Reference
}

// DeviceUpdateAccountId returns a reference to field device_update_account_id of azurerm_iothub_device_update_instance.
func (idui iothubDeviceUpdateInstanceAttributes) DeviceUpdateAccountId() terra.StringValue {
	return terra.ReferenceAsString(idui.ref.Append("device_update_account_id"))
}

// DiagnosticEnabled returns a reference to field diagnostic_enabled of azurerm_iothub_device_update_instance.
func (idui iothubDeviceUpdateInstanceAttributes) DiagnosticEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(idui.ref.Append("diagnostic_enabled"))
}

// Id returns a reference to field id of azurerm_iothub_device_update_instance.
func (idui iothubDeviceUpdateInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(idui.ref.Append("id"))
}

// IothubId returns a reference to field iothub_id of azurerm_iothub_device_update_instance.
func (idui iothubDeviceUpdateInstanceAttributes) IothubId() terra.StringValue {
	return terra.ReferenceAsString(idui.ref.Append("iothub_id"))
}

// Name returns a reference to field name of azurerm_iothub_device_update_instance.
func (idui iothubDeviceUpdateInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(idui.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_iothub_device_update_instance.
func (idui iothubDeviceUpdateInstanceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](idui.ref.Append("tags"))
}

func (idui iothubDeviceUpdateInstanceAttributes) DiagnosticStorageAccount() terra.ListValue[iothubdeviceupdateinstance.DiagnosticStorageAccountAttributes] {
	return terra.ReferenceAsList[iothubdeviceupdateinstance.DiagnosticStorageAccountAttributes](idui.ref.Append("diagnostic_storage_account"))
}

func (idui iothubDeviceUpdateInstanceAttributes) Timeouts() iothubdeviceupdateinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iothubdeviceupdateinstance.TimeoutsAttributes](idui.ref.Append("timeouts"))
}

type iothubDeviceUpdateInstanceState struct {
	DeviceUpdateAccountId    string                                                     `json:"device_update_account_id"`
	DiagnosticEnabled        bool                                                       `json:"diagnostic_enabled"`
	Id                       string                                                     `json:"id"`
	IothubId                 string                                                     `json:"iothub_id"`
	Name                     string                                                     `json:"name"`
	Tags                     map[string]string                                          `json:"tags"`
	DiagnosticStorageAccount []iothubdeviceupdateinstance.DiagnosticStorageAccountState `json:"diagnostic_storage_account"`
	Timeouts                 *iothubdeviceupdateinstance.TimeoutsState                  `json:"timeouts"`
}