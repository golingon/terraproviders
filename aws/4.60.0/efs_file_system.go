// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	efsfilesystem "github.com/golingon/terraproviders/aws/4.60.0/efsfilesystem"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEfsFileSystem creates a new instance of [EfsFileSystem].
func NewEfsFileSystem(name string, args EfsFileSystemArgs) *EfsFileSystem {
	return &EfsFileSystem{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EfsFileSystem)(nil)

// EfsFileSystem represents the Terraform resource aws_efs_file_system.
type EfsFileSystem struct {
	Name      string
	Args      EfsFileSystemArgs
	state     *efsFileSystemState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EfsFileSystem].
func (efs *EfsFileSystem) Type() string {
	return "aws_efs_file_system"
}

// LocalName returns the local name for [EfsFileSystem].
func (efs *EfsFileSystem) LocalName() string {
	return efs.Name
}

// Configuration returns the configuration (args) for [EfsFileSystem].
func (efs *EfsFileSystem) Configuration() interface{} {
	return efs.Args
}

// DependOn is used for other resources to depend on [EfsFileSystem].
func (efs *EfsFileSystem) DependOn() terra.Reference {
	return terra.ReferenceResource(efs)
}

// Dependencies returns the list of resources [EfsFileSystem] depends_on.
func (efs *EfsFileSystem) Dependencies() terra.Dependencies {
	return efs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EfsFileSystem].
func (efs *EfsFileSystem) LifecycleManagement() *terra.Lifecycle {
	return efs.Lifecycle
}

// Attributes returns the attributes for [EfsFileSystem].
func (efs *EfsFileSystem) Attributes() efsFileSystemAttributes {
	return efsFileSystemAttributes{ref: terra.ReferenceResource(efs)}
}

// ImportState imports the given attribute values into [EfsFileSystem]'s state.
func (efs *EfsFileSystem) ImportState(av io.Reader) error {
	efs.state = &efsFileSystemState{}
	if err := json.NewDecoder(av).Decode(efs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", efs.Type(), efs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EfsFileSystem] has state.
func (efs *EfsFileSystem) State() (*efsFileSystemState, bool) {
	return efs.state, efs.state != nil
}

// StateMust returns the state for [EfsFileSystem]. Panics if the state is nil.
func (efs *EfsFileSystem) StateMust() *efsFileSystemState {
	if efs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", efs.Type(), efs.LocalName()))
	}
	return efs.state
}

// EfsFileSystemArgs contains the configurations for aws_efs_file_system.
type EfsFileSystemArgs struct {
	// AvailabilityZoneName: string, optional
	AvailabilityZoneName terra.StringValue `hcl:"availability_zone_name,attr"`
	// CreationToken: string, optional
	CreationToken terra.StringValue `hcl:"creation_token,attr"`
	// Encrypted: bool, optional
	Encrypted terra.BoolValue `hcl:"encrypted,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// PerformanceMode: string, optional
	PerformanceMode terra.StringValue `hcl:"performance_mode,attr"`
	// ProvisionedThroughputInMibps: number, optional
	ProvisionedThroughputInMibps terra.NumberValue `hcl:"provisioned_throughput_in_mibps,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ThroughputMode: string, optional
	ThroughputMode terra.StringValue `hcl:"throughput_mode,attr"`
	// SizeInBytes: min=0
	SizeInBytes []efsfilesystem.SizeInBytes `hcl:"size_in_bytes,block" validate:"min=0"`
	// LifecyclePolicy: min=0,max=2
	LifecyclePolicy []efsfilesystem.LifecyclePolicy `hcl:"lifecycle_policy,block" validate:"min=0,max=2"`
}
type efsFileSystemAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_efs_file_system.
func (efs efsFileSystemAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("arn"))
}

// AvailabilityZoneId returns a reference to field availability_zone_id of aws_efs_file_system.
func (efs efsFileSystemAttributes) AvailabilityZoneId() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("availability_zone_id"))
}

// AvailabilityZoneName returns a reference to field availability_zone_name of aws_efs_file_system.
func (efs efsFileSystemAttributes) AvailabilityZoneName() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("availability_zone_name"))
}

// CreationToken returns a reference to field creation_token of aws_efs_file_system.
func (efs efsFileSystemAttributes) CreationToken() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("creation_token"))
}

// DnsName returns a reference to field dns_name of aws_efs_file_system.
func (efs efsFileSystemAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("dns_name"))
}

// Encrypted returns a reference to field encrypted of aws_efs_file_system.
func (efs efsFileSystemAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(efs.ref.Append("encrypted"))
}

// Id returns a reference to field id of aws_efs_file_system.
func (efs efsFileSystemAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_efs_file_system.
func (efs efsFileSystemAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("kms_key_id"))
}

// NumberOfMountTargets returns a reference to field number_of_mount_targets of aws_efs_file_system.
func (efs efsFileSystemAttributes) NumberOfMountTargets() terra.NumberValue {
	return terra.ReferenceAsNumber(efs.ref.Append("number_of_mount_targets"))
}

// OwnerId returns a reference to field owner_id of aws_efs_file_system.
func (efs efsFileSystemAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("owner_id"))
}

// PerformanceMode returns a reference to field performance_mode of aws_efs_file_system.
func (efs efsFileSystemAttributes) PerformanceMode() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("performance_mode"))
}

// ProvisionedThroughputInMibps returns a reference to field provisioned_throughput_in_mibps of aws_efs_file_system.
func (efs efsFileSystemAttributes) ProvisionedThroughputInMibps() terra.NumberValue {
	return terra.ReferenceAsNumber(efs.ref.Append("provisioned_throughput_in_mibps"))
}

// Tags returns a reference to field tags of aws_efs_file_system.
func (efs efsFileSystemAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](efs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_efs_file_system.
func (efs efsFileSystemAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](efs.ref.Append("tags_all"))
}

// ThroughputMode returns a reference to field throughput_mode of aws_efs_file_system.
func (efs efsFileSystemAttributes) ThroughputMode() terra.StringValue {
	return terra.ReferenceAsString(efs.ref.Append("throughput_mode"))
}

func (efs efsFileSystemAttributes) SizeInBytes() terra.ListValue[efsfilesystem.SizeInBytesAttributes] {
	return terra.ReferenceAsList[efsfilesystem.SizeInBytesAttributes](efs.ref.Append("size_in_bytes"))
}

func (efs efsFileSystemAttributes) LifecyclePolicy() terra.ListValue[efsfilesystem.LifecyclePolicyAttributes] {
	return terra.ReferenceAsList[efsfilesystem.LifecyclePolicyAttributes](efs.ref.Append("lifecycle_policy"))
}

type efsFileSystemState struct {
	Arn                          string                               `json:"arn"`
	AvailabilityZoneId           string                               `json:"availability_zone_id"`
	AvailabilityZoneName         string                               `json:"availability_zone_name"`
	CreationToken                string                               `json:"creation_token"`
	DnsName                      string                               `json:"dns_name"`
	Encrypted                    bool                                 `json:"encrypted"`
	Id                           string                               `json:"id"`
	KmsKeyId                     string                               `json:"kms_key_id"`
	NumberOfMountTargets         float64                              `json:"number_of_mount_targets"`
	OwnerId                      string                               `json:"owner_id"`
	PerformanceMode              string                               `json:"performance_mode"`
	ProvisionedThroughputInMibps float64                              `json:"provisioned_throughput_in_mibps"`
	Tags                         map[string]string                    `json:"tags"`
	TagsAll                      map[string]string                    `json:"tags_all"`
	ThroughputMode               string                               `json:"throughput_mode"`
	SizeInBytes                  []efsfilesystem.SizeInBytesState     `json:"size_in_bytes"`
	LifecyclePolicy              []efsfilesystem.LifecyclePolicyState `json:"lifecycle_policy"`
}
