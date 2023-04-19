// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	fsxontapstoragevirtualmachine "github.com/golingon/terraproviders/aws/4.63.0/fsxontapstoragevirtualmachine"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFsxOntapStorageVirtualMachine creates a new instance of [FsxOntapStorageVirtualMachine].
func NewFsxOntapStorageVirtualMachine(name string, args FsxOntapStorageVirtualMachineArgs) *FsxOntapStorageVirtualMachine {
	return &FsxOntapStorageVirtualMachine{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FsxOntapStorageVirtualMachine)(nil)

// FsxOntapStorageVirtualMachine represents the Terraform resource aws_fsx_ontap_storage_virtual_machine.
type FsxOntapStorageVirtualMachine struct {
	Name      string
	Args      FsxOntapStorageVirtualMachineArgs
	state     *fsxOntapStorageVirtualMachineState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FsxOntapStorageVirtualMachine].
func (fosvm *FsxOntapStorageVirtualMachine) Type() string {
	return "aws_fsx_ontap_storage_virtual_machine"
}

// LocalName returns the local name for [FsxOntapStorageVirtualMachine].
func (fosvm *FsxOntapStorageVirtualMachine) LocalName() string {
	return fosvm.Name
}

// Configuration returns the configuration (args) for [FsxOntapStorageVirtualMachine].
func (fosvm *FsxOntapStorageVirtualMachine) Configuration() interface{} {
	return fosvm.Args
}

// DependOn is used for other resources to depend on [FsxOntapStorageVirtualMachine].
func (fosvm *FsxOntapStorageVirtualMachine) DependOn() terra.Reference {
	return terra.ReferenceResource(fosvm)
}

// Dependencies returns the list of resources [FsxOntapStorageVirtualMachine] depends_on.
func (fosvm *FsxOntapStorageVirtualMachine) Dependencies() terra.Dependencies {
	return fosvm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FsxOntapStorageVirtualMachine].
func (fosvm *FsxOntapStorageVirtualMachine) LifecycleManagement() *terra.Lifecycle {
	return fosvm.Lifecycle
}

// Attributes returns the attributes for [FsxOntapStorageVirtualMachine].
func (fosvm *FsxOntapStorageVirtualMachine) Attributes() fsxOntapStorageVirtualMachineAttributes {
	return fsxOntapStorageVirtualMachineAttributes{ref: terra.ReferenceResource(fosvm)}
}

// ImportState imports the given attribute values into [FsxOntapStorageVirtualMachine]'s state.
func (fosvm *FsxOntapStorageVirtualMachine) ImportState(av io.Reader) error {
	fosvm.state = &fsxOntapStorageVirtualMachineState{}
	if err := json.NewDecoder(av).Decode(fosvm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fosvm.Type(), fosvm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FsxOntapStorageVirtualMachine] has state.
func (fosvm *FsxOntapStorageVirtualMachine) State() (*fsxOntapStorageVirtualMachineState, bool) {
	return fosvm.state, fosvm.state != nil
}

// StateMust returns the state for [FsxOntapStorageVirtualMachine]. Panics if the state is nil.
func (fosvm *FsxOntapStorageVirtualMachine) StateMust() *fsxOntapStorageVirtualMachineState {
	if fosvm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fosvm.Type(), fosvm.LocalName()))
	}
	return fosvm.state
}

// FsxOntapStorageVirtualMachineArgs contains the configurations for aws_fsx_ontap_storage_virtual_machine.
type FsxOntapStorageVirtualMachineArgs struct {
	// FileSystemId: string, required
	FileSystemId terra.StringValue `hcl:"file_system_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RootVolumeSecurityStyle: string, optional
	RootVolumeSecurityStyle terra.StringValue `hcl:"root_volume_security_style,attr"`
	// SvmAdminPassword: string, optional
	SvmAdminPassword terra.StringValue `hcl:"svm_admin_password,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Endpoints: min=0
	Endpoints []fsxontapstoragevirtualmachine.Endpoints `hcl:"endpoints,block" validate:"min=0"`
	// ActiveDirectoryConfiguration: optional
	ActiveDirectoryConfiguration *fsxontapstoragevirtualmachine.ActiveDirectoryConfiguration `hcl:"active_directory_configuration,block"`
	// Timeouts: optional
	Timeouts *fsxontapstoragevirtualmachine.Timeouts `hcl:"timeouts,block"`
}
type fsxOntapStorageVirtualMachineAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_fsx_ontap_storage_virtual_machine.
func (fosvm fsxOntapStorageVirtualMachineAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(fosvm.ref.Append("arn"))
}

// FileSystemId returns a reference to field file_system_id of aws_fsx_ontap_storage_virtual_machine.
func (fosvm fsxOntapStorageVirtualMachineAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceAsString(fosvm.ref.Append("file_system_id"))
}

// Id returns a reference to field id of aws_fsx_ontap_storage_virtual_machine.
func (fosvm fsxOntapStorageVirtualMachineAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fosvm.ref.Append("id"))
}

// Name returns a reference to field name of aws_fsx_ontap_storage_virtual_machine.
func (fosvm fsxOntapStorageVirtualMachineAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fosvm.ref.Append("name"))
}

// RootVolumeSecurityStyle returns a reference to field root_volume_security_style of aws_fsx_ontap_storage_virtual_machine.
func (fosvm fsxOntapStorageVirtualMachineAttributes) RootVolumeSecurityStyle() terra.StringValue {
	return terra.ReferenceAsString(fosvm.ref.Append("root_volume_security_style"))
}

// Subtype returns a reference to field subtype of aws_fsx_ontap_storage_virtual_machine.
func (fosvm fsxOntapStorageVirtualMachineAttributes) Subtype() terra.StringValue {
	return terra.ReferenceAsString(fosvm.ref.Append("subtype"))
}

// SvmAdminPassword returns a reference to field svm_admin_password of aws_fsx_ontap_storage_virtual_machine.
func (fosvm fsxOntapStorageVirtualMachineAttributes) SvmAdminPassword() terra.StringValue {
	return terra.ReferenceAsString(fosvm.ref.Append("svm_admin_password"))
}

// Tags returns a reference to field tags of aws_fsx_ontap_storage_virtual_machine.
func (fosvm fsxOntapStorageVirtualMachineAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fosvm.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_fsx_ontap_storage_virtual_machine.
func (fosvm fsxOntapStorageVirtualMachineAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fosvm.ref.Append("tags_all"))
}

// Uuid returns a reference to field uuid of aws_fsx_ontap_storage_virtual_machine.
func (fosvm fsxOntapStorageVirtualMachineAttributes) Uuid() terra.StringValue {
	return terra.ReferenceAsString(fosvm.ref.Append("uuid"))
}

func (fosvm fsxOntapStorageVirtualMachineAttributes) Endpoints() terra.ListValue[fsxontapstoragevirtualmachine.EndpointsAttributes] {
	return terra.ReferenceAsList[fsxontapstoragevirtualmachine.EndpointsAttributes](fosvm.ref.Append("endpoints"))
}

func (fosvm fsxOntapStorageVirtualMachineAttributes) ActiveDirectoryConfiguration() terra.ListValue[fsxontapstoragevirtualmachine.ActiveDirectoryConfigurationAttributes] {
	return terra.ReferenceAsList[fsxontapstoragevirtualmachine.ActiveDirectoryConfigurationAttributes](fosvm.ref.Append("active_directory_configuration"))
}

func (fosvm fsxOntapStorageVirtualMachineAttributes) Timeouts() fsxontapstoragevirtualmachine.TimeoutsAttributes {
	return terra.ReferenceAsSingle[fsxontapstoragevirtualmachine.TimeoutsAttributes](fosvm.ref.Append("timeouts"))
}

type fsxOntapStorageVirtualMachineState struct {
	Arn                          string                                                            `json:"arn"`
	FileSystemId                 string                                                            `json:"file_system_id"`
	Id                           string                                                            `json:"id"`
	Name                         string                                                            `json:"name"`
	RootVolumeSecurityStyle      string                                                            `json:"root_volume_security_style"`
	Subtype                      string                                                            `json:"subtype"`
	SvmAdminPassword             string                                                            `json:"svm_admin_password"`
	Tags                         map[string]string                                                 `json:"tags"`
	TagsAll                      map[string]string                                                 `json:"tags_all"`
	Uuid                         string                                                            `json:"uuid"`
	Endpoints                    []fsxontapstoragevirtualmachine.EndpointsState                    `json:"endpoints"`
	ActiveDirectoryConfiguration []fsxontapstoragevirtualmachine.ActiveDirectoryConfigurationState `json:"active_directory_configuration"`
	Timeouts                     *fsxontapstoragevirtualmachine.TimeoutsState                      `json:"timeouts"`
}
