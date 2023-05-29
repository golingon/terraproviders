// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	devtestlinuxvirtualmachine "github.com/golingon/terraproviders/azurerm/3.58.0/devtestlinuxvirtualmachine"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDevTestLinuxVirtualMachine creates a new instance of [DevTestLinuxVirtualMachine].
func NewDevTestLinuxVirtualMachine(name string, args DevTestLinuxVirtualMachineArgs) *DevTestLinuxVirtualMachine {
	return &DevTestLinuxVirtualMachine{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DevTestLinuxVirtualMachine)(nil)

// DevTestLinuxVirtualMachine represents the Terraform resource azurerm_dev_test_linux_virtual_machine.
type DevTestLinuxVirtualMachine struct {
	Name      string
	Args      DevTestLinuxVirtualMachineArgs
	state     *devTestLinuxVirtualMachineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DevTestLinuxVirtualMachine].
func (dtlvm *DevTestLinuxVirtualMachine) Type() string {
	return "azurerm_dev_test_linux_virtual_machine"
}

// LocalName returns the local name for [DevTestLinuxVirtualMachine].
func (dtlvm *DevTestLinuxVirtualMachine) LocalName() string {
	return dtlvm.Name
}

// Configuration returns the configuration (args) for [DevTestLinuxVirtualMachine].
func (dtlvm *DevTestLinuxVirtualMachine) Configuration() interface{} {
	return dtlvm.Args
}

// DependOn is used for other resources to depend on [DevTestLinuxVirtualMachine].
func (dtlvm *DevTestLinuxVirtualMachine) DependOn() terra.Reference {
	return terra.ReferenceResource(dtlvm)
}

// Dependencies returns the list of resources [DevTestLinuxVirtualMachine] depends_on.
func (dtlvm *DevTestLinuxVirtualMachine) Dependencies() terra.Dependencies {
	return dtlvm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DevTestLinuxVirtualMachine].
func (dtlvm *DevTestLinuxVirtualMachine) LifecycleManagement() *terra.Lifecycle {
	return dtlvm.Lifecycle
}

// Attributes returns the attributes for [DevTestLinuxVirtualMachine].
func (dtlvm *DevTestLinuxVirtualMachine) Attributes() devTestLinuxVirtualMachineAttributes {
	return devTestLinuxVirtualMachineAttributes{ref: terra.ReferenceResource(dtlvm)}
}

// ImportState imports the given attribute values into [DevTestLinuxVirtualMachine]'s state.
func (dtlvm *DevTestLinuxVirtualMachine) ImportState(av io.Reader) error {
	dtlvm.state = &devTestLinuxVirtualMachineState{}
	if err := json.NewDecoder(av).Decode(dtlvm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dtlvm.Type(), dtlvm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DevTestLinuxVirtualMachine] has state.
func (dtlvm *DevTestLinuxVirtualMachine) State() (*devTestLinuxVirtualMachineState, bool) {
	return dtlvm.state, dtlvm.state != nil
}

// StateMust returns the state for [DevTestLinuxVirtualMachine]. Panics if the state is nil.
func (dtlvm *DevTestLinuxVirtualMachine) StateMust() *devTestLinuxVirtualMachineState {
	if dtlvm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dtlvm.Type(), dtlvm.LocalName()))
	}
	return dtlvm.state
}

// DevTestLinuxVirtualMachineArgs contains the configurations for azurerm_dev_test_linux_virtual_machine.
type DevTestLinuxVirtualMachineArgs struct {
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
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Size: string, required
	Size terra.StringValue `hcl:"size,attr" validate:"required"`
	// SshKey: string, optional
	SshKey terra.StringValue `hcl:"ssh_key,attr"`
	// StorageType: string, required
	StorageType terra.StringValue `hcl:"storage_type,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
	// GalleryImageReference: required
	GalleryImageReference *devtestlinuxvirtualmachine.GalleryImageReference `hcl:"gallery_image_reference,block" validate:"required"`
	// InboundNatRule: min=0
	InboundNatRule []devtestlinuxvirtualmachine.InboundNatRule `hcl:"inbound_nat_rule,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *devtestlinuxvirtualmachine.Timeouts `hcl:"timeouts,block"`
}
type devTestLinuxVirtualMachineAttributes struct {
	ref terra.Reference
}

// AllowClaim returns a reference to field allow_claim of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) AllowClaim() terra.BoolValue {
	return terra.ReferenceAsBool(dtlvm.ref.Append("allow_claim"))
}

// DisallowPublicIpAddress returns a reference to field disallow_public_ip_address of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) DisallowPublicIpAddress() terra.BoolValue {
	return terra.ReferenceAsBool(dtlvm.ref.Append("disallow_public_ip_address"))
}

// Fqdn returns a reference to field fqdn of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dtlvm.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dtlvm.ref.Append("id"))
}

// LabName returns a reference to field lab_name of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) LabName() terra.StringValue {
	return terra.ReferenceAsString(dtlvm.ref.Append("lab_name"))
}

// LabSubnetName returns a reference to field lab_subnet_name of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) LabSubnetName() terra.StringValue {
	return terra.ReferenceAsString(dtlvm.ref.Append("lab_subnet_name"))
}

// LabVirtualNetworkId returns a reference to field lab_virtual_network_id of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) LabVirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(dtlvm.ref.Append("lab_virtual_network_id"))
}

// Location returns a reference to field location of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dtlvm.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dtlvm.ref.Append("name"))
}

// Notes returns a reference to field notes of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) Notes() terra.StringValue {
	return terra.ReferenceAsString(dtlvm.ref.Append("notes"))
}

// Password returns a reference to field password of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(dtlvm.ref.Append("password"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dtlvm.ref.Append("resource_group_name"))
}

// Size returns a reference to field size of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) Size() terra.StringValue {
	return terra.ReferenceAsString(dtlvm.ref.Append("size"))
}

// SshKey returns a reference to field ssh_key of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) SshKey() terra.StringValue {
	return terra.ReferenceAsString(dtlvm.ref.Append("ssh_key"))
}

// StorageType returns a reference to field storage_type of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(dtlvm.ref.Append("storage_type"))
}

// Tags returns a reference to field tags of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dtlvm.ref.Append("tags"))
}

// UniqueIdentifier returns a reference to field unique_identifier of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) UniqueIdentifier() terra.StringValue {
	return terra.ReferenceAsString(dtlvm.ref.Append("unique_identifier"))
}

// Username returns a reference to field username of azurerm_dev_test_linux_virtual_machine.
func (dtlvm devTestLinuxVirtualMachineAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(dtlvm.ref.Append("username"))
}

func (dtlvm devTestLinuxVirtualMachineAttributes) GalleryImageReference() terra.ListValue[devtestlinuxvirtualmachine.GalleryImageReferenceAttributes] {
	return terra.ReferenceAsList[devtestlinuxvirtualmachine.GalleryImageReferenceAttributes](dtlvm.ref.Append("gallery_image_reference"))
}

func (dtlvm devTestLinuxVirtualMachineAttributes) InboundNatRule() terra.SetValue[devtestlinuxvirtualmachine.InboundNatRuleAttributes] {
	return terra.ReferenceAsSet[devtestlinuxvirtualmachine.InboundNatRuleAttributes](dtlvm.ref.Append("inbound_nat_rule"))
}

func (dtlvm devTestLinuxVirtualMachineAttributes) Timeouts() devtestlinuxvirtualmachine.TimeoutsAttributes {
	return terra.ReferenceAsSingle[devtestlinuxvirtualmachine.TimeoutsAttributes](dtlvm.ref.Append("timeouts"))
}

type devTestLinuxVirtualMachineState struct {
	AllowClaim              bool                                                    `json:"allow_claim"`
	DisallowPublicIpAddress bool                                                    `json:"disallow_public_ip_address"`
	Fqdn                    string                                                  `json:"fqdn"`
	Id                      string                                                  `json:"id"`
	LabName                 string                                                  `json:"lab_name"`
	LabSubnetName           string                                                  `json:"lab_subnet_name"`
	LabVirtualNetworkId     string                                                  `json:"lab_virtual_network_id"`
	Location                string                                                  `json:"location"`
	Name                    string                                                  `json:"name"`
	Notes                   string                                                  `json:"notes"`
	Password                string                                                  `json:"password"`
	ResourceGroupName       string                                                  `json:"resource_group_name"`
	Size                    string                                                  `json:"size"`
	SshKey                  string                                                  `json:"ssh_key"`
	StorageType             string                                                  `json:"storage_type"`
	Tags                    map[string]string                                       `json:"tags"`
	UniqueIdentifier        string                                                  `json:"unique_identifier"`
	Username                string                                                  `json:"username"`
	GalleryImageReference   []devtestlinuxvirtualmachine.GalleryImageReferenceState `json:"gallery_image_reference"`
	InboundNatRule          []devtestlinuxvirtualmachine.InboundNatRuleState        `json:"inbound_nat_rule"`
	Timeouts                *devtestlinuxvirtualmachine.TimeoutsState               `json:"timeouts"`
}