// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	devtestwindowsvirtualmachine "github.com/golingon/terraproviders/azurerm/3.64.0/devtestwindowsvirtualmachine"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDevTestWindowsVirtualMachine creates a new instance of [DevTestWindowsVirtualMachine].
func NewDevTestWindowsVirtualMachine(name string, args DevTestWindowsVirtualMachineArgs) *DevTestWindowsVirtualMachine {
	return &DevTestWindowsVirtualMachine{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DevTestWindowsVirtualMachine)(nil)

// DevTestWindowsVirtualMachine represents the Terraform resource azurerm_dev_test_windows_virtual_machine.
type DevTestWindowsVirtualMachine struct {
	Name      string
	Args      DevTestWindowsVirtualMachineArgs
	state     *devTestWindowsVirtualMachineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DevTestWindowsVirtualMachine].
func (dtwvm *DevTestWindowsVirtualMachine) Type() string {
	return "azurerm_dev_test_windows_virtual_machine"
}

// LocalName returns the local name for [DevTestWindowsVirtualMachine].
func (dtwvm *DevTestWindowsVirtualMachine) LocalName() string {
	return dtwvm.Name
}

// Configuration returns the configuration (args) for [DevTestWindowsVirtualMachine].
func (dtwvm *DevTestWindowsVirtualMachine) Configuration() interface{} {
	return dtwvm.Args
}

// DependOn is used for other resources to depend on [DevTestWindowsVirtualMachine].
func (dtwvm *DevTestWindowsVirtualMachine) DependOn() terra.Reference {
	return terra.ReferenceResource(dtwvm)
}

// Dependencies returns the list of resources [DevTestWindowsVirtualMachine] depends_on.
func (dtwvm *DevTestWindowsVirtualMachine) Dependencies() terra.Dependencies {
	return dtwvm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DevTestWindowsVirtualMachine].
func (dtwvm *DevTestWindowsVirtualMachine) LifecycleManagement() *terra.Lifecycle {
	return dtwvm.Lifecycle
}

// Attributes returns the attributes for [DevTestWindowsVirtualMachine].
func (dtwvm *DevTestWindowsVirtualMachine) Attributes() devTestWindowsVirtualMachineAttributes {
	return devTestWindowsVirtualMachineAttributes{ref: terra.ReferenceResource(dtwvm)}
}

// ImportState imports the given attribute values into [DevTestWindowsVirtualMachine]'s state.
func (dtwvm *DevTestWindowsVirtualMachine) ImportState(av io.Reader) error {
	dtwvm.state = &devTestWindowsVirtualMachineState{}
	if err := json.NewDecoder(av).Decode(dtwvm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dtwvm.Type(), dtwvm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DevTestWindowsVirtualMachine] has state.
func (dtwvm *DevTestWindowsVirtualMachine) State() (*devTestWindowsVirtualMachineState, bool) {
	return dtwvm.state, dtwvm.state != nil
}

// StateMust returns the state for [DevTestWindowsVirtualMachine]. Panics if the state is nil.
func (dtwvm *DevTestWindowsVirtualMachine) StateMust() *devTestWindowsVirtualMachineState {
	if dtwvm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dtwvm.Type(), dtwvm.LocalName()))
	}
	return dtwvm.state
}

// DevTestWindowsVirtualMachineArgs contains the configurations for azurerm_dev_test_windows_virtual_machine.
type DevTestWindowsVirtualMachineArgs struct {
	// AllowClaim: bool, optional
	AllowClaim terra.BoolValue `hcl:"allow_claim,attr"`
	// DisallowPublicIpAddress: bool, optional
	DisallowPublicIpAddress terra.BoolValue `hcl:"disallow_public_ip_address,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LabName: string, required
	LabName terra.StringValue `hcl:"lab_name,attr" validate:"required"`
	// LabSubnetName: string, required
	LabSubnetName terra.StringValue `hcl:"lab_subnet_name,attr" validate:"required"`
	// LabVirtualNetworkId: string, required
	LabVirtualNetworkId terra.StringValue `hcl:"lab_virtual_network_id,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Notes: string, optional
	Notes terra.StringValue `hcl:"notes,attr"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Size: string, required
	Size terra.StringValue `hcl:"size,attr" validate:"required"`
	// StorageType: string, required
	StorageType terra.StringValue `hcl:"storage_type,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
	// GalleryImageReference: required
	GalleryImageReference *devtestwindowsvirtualmachine.GalleryImageReference `hcl:"gallery_image_reference,block" validate:"required"`
	// InboundNatRule: min=0
	InboundNatRule []devtestwindowsvirtualmachine.InboundNatRule `hcl:"inbound_nat_rule,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *devtestwindowsvirtualmachine.Timeouts `hcl:"timeouts,block"`
}
type devTestWindowsVirtualMachineAttributes struct {
	ref terra.Reference
}

// AllowClaim returns a reference to field allow_claim of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) AllowClaim() terra.BoolValue {
	return terra.ReferenceAsBool(dtwvm.ref.Append("allow_claim"))
}

// DisallowPublicIpAddress returns a reference to field disallow_public_ip_address of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) DisallowPublicIpAddress() terra.BoolValue {
	return terra.ReferenceAsBool(dtwvm.ref.Append("disallow_public_ip_address"))
}

// Fqdn returns a reference to field fqdn of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dtwvm.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dtwvm.ref.Append("id"))
}

// LabName returns a reference to field lab_name of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) LabName() terra.StringValue {
	return terra.ReferenceAsString(dtwvm.ref.Append("lab_name"))
}

// LabSubnetName returns a reference to field lab_subnet_name of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) LabSubnetName() terra.StringValue {
	return terra.ReferenceAsString(dtwvm.ref.Append("lab_subnet_name"))
}

// LabVirtualNetworkId returns a reference to field lab_virtual_network_id of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) LabVirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(dtwvm.ref.Append("lab_virtual_network_id"))
}

// Location returns a reference to field location of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dtwvm.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dtwvm.ref.Append("name"))
}

// Notes returns a reference to field notes of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) Notes() terra.StringValue {
	return terra.ReferenceAsString(dtwvm.ref.Append("notes"))
}

// Password returns a reference to field password of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(dtwvm.ref.Append("password"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dtwvm.ref.Append("resource_group_name"))
}

// Size returns a reference to field size of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) Size() terra.StringValue {
	return terra.ReferenceAsString(dtwvm.ref.Append("size"))
}

// StorageType returns a reference to field storage_type of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(dtwvm.ref.Append("storage_type"))
}

// Tags returns a reference to field tags of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dtwvm.ref.Append("tags"))
}

// UniqueIdentifier returns a reference to field unique_identifier of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) UniqueIdentifier() terra.StringValue {
	return terra.ReferenceAsString(dtwvm.ref.Append("unique_identifier"))
}

// Username returns a reference to field username of azurerm_dev_test_windows_virtual_machine.
func (dtwvm devTestWindowsVirtualMachineAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(dtwvm.ref.Append("username"))
}

func (dtwvm devTestWindowsVirtualMachineAttributes) GalleryImageReference() terra.ListValue[devtestwindowsvirtualmachine.GalleryImageReferenceAttributes] {
	return terra.ReferenceAsList[devtestwindowsvirtualmachine.GalleryImageReferenceAttributes](dtwvm.ref.Append("gallery_image_reference"))
}

func (dtwvm devTestWindowsVirtualMachineAttributes) InboundNatRule() terra.SetValue[devtestwindowsvirtualmachine.InboundNatRuleAttributes] {
	return terra.ReferenceAsSet[devtestwindowsvirtualmachine.InboundNatRuleAttributes](dtwvm.ref.Append("inbound_nat_rule"))
}

func (dtwvm devTestWindowsVirtualMachineAttributes) Timeouts() devtestwindowsvirtualmachine.TimeoutsAttributes {
	return terra.ReferenceAsSingle[devtestwindowsvirtualmachine.TimeoutsAttributes](dtwvm.ref.Append("timeouts"))
}

type devTestWindowsVirtualMachineState struct {
	AllowClaim              bool                                                      `json:"allow_claim"`
	DisallowPublicIpAddress bool                                                      `json:"disallow_public_ip_address"`
	Fqdn                    string                                                    `json:"fqdn"`
	Id                      string                                                    `json:"id"`
	LabName                 string                                                    `json:"lab_name"`
	LabSubnetName           string                                                    `json:"lab_subnet_name"`
	LabVirtualNetworkId     string                                                    `json:"lab_virtual_network_id"`
	Location                string                                                    `json:"location"`
	Name                    string                                                    `json:"name"`
	Notes                   string                                                    `json:"notes"`
	Password                string                                                    `json:"password"`
	ResourceGroupName       string                                                    `json:"resource_group_name"`
	Size                    string                                                    `json:"size"`
	StorageType             string                                                    `json:"storage_type"`
	Tags                    map[string]string                                         `json:"tags"`
	UniqueIdentifier        string                                                    `json:"unique_identifier"`
	Username                string                                                    `json:"username"`
	GalleryImageReference   []devtestwindowsvirtualmachine.GalleryImageReferenceState `json:"gallery_image_reference"`
	InboundNatRule          []devtestwindowsvirtualmachine.InboundNatRuleState        `json:"inbound_nat_rule"`
	Timeouts                *devtestwindowsvirtualmachine.TimeoutsState               `json:"timeouts"`
}
