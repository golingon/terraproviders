// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iotsecuritydevicegroup "github.com/golingon/terraproviders/azurerm/3.67.0/iotsecuritydevicegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIotSecurityDeviceGroup creates a new instance of [IotSecurityDeviceGroup].
func NewIotSecurityDeviceGroup(name string, args IotSecurityDeviceGroupArgs) *IotSecurityDeviceGroup {
	return &IotSecurityDeviceGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotSecurityDeviceGroup)(nil)

// IotSecurityDeviceGroup represents the Terraform resource azurerm_iot_security_device_group.
type IotSecurityDeviceGroup struct {
	Name      string
	Args      IotSecurityDeviceGroupArgs
	state     *iotSecurityDeviceGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotSecurityDeviceGroup].
func (isdg *IotSecurityDeviceGroup) Type() string {
	return "azurerm_iot_security_device_group"
}

// LocalName returns the local name for [IotSecurityDeviceGroup].
func (isdg *IotSecurityDeviceGroup) LocalName() string {
	return isdg.Name
}

// Configuration returns the configuration (args) for [IotSecurityDeviceGroup].
func (isdg *IotSecurityDeviceGroup) Configuration() interface{} {
	return isdg.Args
}

// DependOn is used for other resources to depend on [IotSecurityDeviceGroup].
func (isdg *IotSecurityDeviceGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(isdg)
}

// Dependencies returns the list of resources [IotSecurityDeviceGroup] depends_on.
func (isdg *IotSecurityDeviceGroup) Dependencies() terra.Dependencies {
	return isdg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotSecurityDeviceGroup].
func (isdg *IotSecurityDeviceGroup) LifecycleManagement() *terra.Lifecycle {
	return isdg.Lifecycle
}

// Attributes returns the attributes for [IotSecurityDeviceGroup].
func (isdg *IotSecurityDeviceGroup) Attributes() iotSecurityDeviceGroupAttributes {
	return iotSecurityDeviceGroupAttributes{ref: terra.ReferenceResource(isdg)}
}

// ImportState imports the given attribute values into [IotSecurityDeviceGroup]'s state.
func (isdg *IotSecurityDeviceGroup) ImportState(av io.Reader) error {
	isdg.state = &iotSecurityDeviceGroupState{}
	if err := json.NewDecoder(av).Decode(isdg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", isdg.Type(), isdg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotSecurityDeviceGroup] has state.
func (isdg *IotSecurityDeviceGroup) State() (*iotSecurityDeviceGroupState, bool) {
	return isdg.state, isdg.state != nil
}

// StateMust returns the state for [IotSecurityDeviceGroup]. Panics if the state is nil.
func (isdg *IotSecurityDeviceGroup) StateMust() *iotSecurityDeviceGroupState {
	if isdg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", isdg.Type(), isdg.LocalName()))
	}
	return isdg.state
}

// IotSecurityDeviceGroupArgs contains the configurations for azurerm_iot_security_device_group.
type IotSecurityDeviceGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IothubId: string, required
	IothubId terra.StringValue `hcl:"iothub_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// AllowRule: optional
	AllowRule *iotsecuritydevicegroup.AllowRule `hcl:"allow_rule,block"`
	// RangeRule: min=0
	RangeRule []iotsecuritydevicegroup.RangeRule `hcl:"range_rule,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *iotsecuritydevicegroup.Timeouts `hcl:"timeouts,block"`
}
type iotSecurityDeviceGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_iot_security_device_group.
func (isdg iotSecurityDeviceGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(isdg.ref.Append("id"))
}

// IothubId returns a reference to field iothub_id of azurerm_iot_security_device_group.
func (isdg iotSecurityDeviceGroupAttributes) IothubId() terra.StringValue {
	return terra.ReferenceAsString(isdg.ref.Append("iothub_id"))
}

// Name returns a reference to field name of azurerm_iot_security_device_group.
func (isdg iotSecurityDeviceGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(isdg.ref.Append("name"))
}

func (isdg iotSecurityDeviceGroupAttributes) AllowRule() terra.ListValue[iotsecuritydevicegroup.AllowRuleAttributes] {
	return terra.ReferenceAsList[iotsecuritydevicegroup.AllowRuleAttributes](isdg.ref.Append("allow_rule"))
}

func (isdg iotSecurityDeviceGroupAttributes) RangeRule() terra.SetValue[iotsecuritydevicegroup.RangeRuleAttributes] {
	return terra.ReferenceAsSet[iotsecuritydevicegroup.RangeRuleAttributes](isdg.ref.Append("range_rule"))
}

func (isdg iotSecurityDeviceGroupAttributes) Timeouts() iotsecuritydevicegroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iotsecuritydevicegroup.TimeoutsAttributes](isdg.ref.Append("timeouts"))
}

type iotSecurityDeviceGroupState struct {
	Id        string                                  `json:"id"`
	IothubId  string                                  `json:"iothub_id"`
	Name      string                                  `json:"name"`
	AllowRule []iotsecuritydevicegroup.AllowRuleState `json:"allow_rule"`
	RangeRule []iotsecuritydevicegroup.RangeRuleState `json:"range_rule"`
	Timeouts  *iotsecuritydevicegroup.TimeoutsState   `json:"timeouts"`
}
