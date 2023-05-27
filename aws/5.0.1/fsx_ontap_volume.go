// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	fsxontapvolume "github.com/golingon/terraproviders/aws/5.0.1/fsxontapvolume"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFsxOntapVolume creates a new instance of [FsxOntapVolume].
func NewFsxOntapVolume(name string, args FsxOntapVolumeArgs) *FsxOntapVolume {
	return &FsxOntapVolume{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FsxOntapVolume)(nil)

// FsxOntapVolume represents the Terraform resource aws_fsx_ontap_volume.
type FsxOntapVolume struct {
	Name      string
	Args      FsxOntapVolumeArgs
	state     *fsxOntapVolumeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FsxOntapVolume].
func (fov *FsxOntapVolume) Type() string {
	return "aws_fsx_ontap_volume"
}

// LocalName returns the local name for [FsxOntapVolume].
func (fov *FsxOntapVolume) LocalName() string {
	return fov.Name
}

// Configuration returns the configuration (args) for [FsxOntapVolume].
func (fov *FsxOntapVolume) Configuration() interface{} {
	return fov.Args
}

// DependOn is used for other resources to depend on [FsxOntapVolume].
func (fov *FsxOntapVolume) DependOn() terra.Reference {
	return terra.ReferenceResource(fov)
}

// Dependencies returns the list of resources [FsxOntapVolume] depends_on.
func (fov *FsxOntapVolume) Dependencies() terra.Dependencies {
	return fov.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FsxOntapVolume].
func (fov *FsxOntapVolume) LifecycleManagement() *terra.Lifecycle {
	return fov.Lifecycle
}

// Attributes returns the attributes for [FsxOntapVolume].
func (fov *FsxOntapVolume) Attributes() fsxOntapVolumeAttributes {
	return fsxOntapVolumeAttributes{ref: terra.ReferenceResource(fov)}
}

// ImportState imports the given attribute values into [FsxOntapVolume]'s state.
func (fov *FsxOntapVolume) ImportState(av io.Reader) error {
	fov.state = &fsxOntapVolumeState{}
	if err := json.NewDecoder(av).Decode(fov.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fov.Type(), fov.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FsxOntapVolume] has state.
func (fov *FsxOntapVolume) State() (*fsxOntapVolumeState, bool) {
	return fov.state, fov.state != nil
}

// StateMust returns the state for [FsxOntapVolume]. Panics if the state is nil.
func (fov *FsxOntapVolume) StateMust() *fsxOntapVolumeState {
	if fov.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fov.Type(), fov.LocalName()))
	}
	return fov.state
}

// FsxOntapVolumeArgs contains the configurations for aws_fsx_ontap_volume.
type FsxOntapVolumeArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// JunctionPath: string, required
	JunctionPath terra.StringValue `hcl:"junction_path,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SecurityStyle: string, optional
	SecurityStyle terra.StringValue `hcl:"security_style,attr"`
	// SizeInMegabytes: number, required
	SizeInMegabytes terra.NumberValue `hcl:"size_in_megabytes,attr" validate:"required"`
	// StorageEfficiencyEnabled: bool, required
	StorageEfficiencyEnabled terra.BoolValue `hcl:"storage_efficiency_enabled,attr" validate:"required"`
	// StorageVirtualMachineId: string, required
	StorageVirtualMachineId terra.StringValue `hcl:"storage_virtual_machine_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VolumeType: string, optional
	VolumeType terra.StringValue `hcl:"volume_type,attr"`
	// TieringPolicy: optional
	TieringPolicy *fsxontapvolume.TieringPolicy `hcl:"tiering_policy,block"`
	// Timeouts: optional
	Timeouts *fsxontapvolume.Timeouts `hcl:"timeouts,block"`
}
type fsxOntapVolumeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_fsx_ontap_volume.
func (fov fsxOntapVolumeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("arn"))
}

// FileSystemId returns a reference to field file_system_id of aws_fsx_ontap_volume.
func (fov fsxOntapVolumeAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("file_system_id"))
}

// FlexcacheEndpointType returns a reference to field flexcache_endpoint_type of aws_fsx_ontap_volume.
func (fov fsxOntapVolumeAttributes) FlexcacheEndpointType() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("flexcache_endpoint_type"))
}

// Id returns a reference to field id of aws_fsx_ontap_volume.
func (fov fsxOntapVolumeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("id"))
}

// JunctionPath returns a reference to field junction_path of aws_fsx_ontap_volume.
func (fov fsxOntapVolumeAttributes) JunctionPath() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("junction_path"))
}

// Name returns a reference to field name of aws_fsx_ontap_volume.
func (fov fsxOntapVolumeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("name"))
}

// OntapVolumeType returns a reference to field ontap_volume_type of aws_fsx_ontap_volume.
func (fov fsxOntapVolumeAttributes) OntapVolumeType() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("ontap_volume_type"))
}

// SecurityStyle returns a reference to field security_style of aws_fsx_ontap_volume.
func (fov fsxOntapVolumeAttributes) SecurityStyle() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("security_style"))
}

// SizeInMegabytes returns a reference to field size_in_megabytes of aws_fsx_ontap_volume.
func (fov fsxOntapVolumeAttributes) SizeInMegabytes() terra.NumberValue {
	return terra.ReferenceAsNumber(fov.ref.Append("size_in_megabytes"))
}

// StorageEfficiencyEnabled returns a reference to field storage_efficiency_enabled of aws_fsx_ontap_volume.
func (fov fsxOntapVolumeAttributes) StorageEfficiencyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(fov.ref.Append("storage_efficiency_enabled"))
}

// StorageVirtualMachineId returns a reference to field storage_virtual_machine_id of aws_fsx_ontap_volume.
func (fov fsxOntapVolumeAttributes) StorageVirtualMachineId() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("storage_virtual_machine_id"))
}

// Tags returns a reference to field tags of aws_fsx_ontap_volume.
func (fov fsxOntapVolumeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fov.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_fsx_ontap_volume.
func (fov fsxOntapVolumeAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fov.ref.Append("tags_all"))
}

// Uuid returns a reference to field uuid of aws_fsx_ontap_volume.
func (fov fsxOntapVolumeAttributes) Uuid() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("uuid"))
}

// VolumeType returns a reference to field volume_type of aws_fsx_ontap_volume.
func (fov fsxOntapVolumeAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceAsString(fov.ref.Append("volume_type"))
}

func (fov fsxOntapVolumeAttributes) TieringPolicy() terra.ListValue[fsxontapvolume.TieringPolicyAttributes] {
	return terra.ReferenceAsList[fsxontapvolume.TieringPolicyAttributes](fov.ref.Append("tiering_policy"))
}

func (fov fsxOntapVolumeAttributes) Timeouts() fsxontapvolume.TimeoutsAttributes {
	return terra.ReferenceAsSingle[fsxontapvolume.TimeoutsAttributes](fov.ref.Append("timeouts"))
}

type fsxOntapVolumeState struct {
	Arn                      string                              `json:"arn"`
	FileSystemId             string                              `json:"file_system_id"`
	FlexcacheEndpointType    string                              `json:"flexcache_endpoint_type"`
	Id                       string                              `json:"id"`
	JunctionPath             string                              `json:"junction_path"`
	Name                     string                              `json:"name"`
	OntapVolumeType          string                              `json:"ontap_volume_type"`
	SecurityStyle            string                              `json:"security_style"`
	SizeInMegabytes          float64                             `json:"size_in_megabytes"`
	StorageEfficiencyEnabled bool                                `json:"storage_efficiency_enabled"`
	StorageVirtualMachineId  string                              `json:"storage_virtual_machine_id"`
	Tags                     map[string]string                   `json:"tags"`
	TagsAll                  map[string]string                   `json:"tags_all"`
	Uuid                     string                              `json:"uuid"`
	VolumeType               string                              `json:"volume_type"`
	TieringPolicy            []fsxontapvolume.TieringPolicyState `json:"tiering_policy"`
	Timeouts                 *fsxontapvolume.TimeoutsState       `json:"timeouts"`
}
