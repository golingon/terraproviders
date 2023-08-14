// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	fsxfilecache "github.com/golingon/terraproviders/aws/5.11.0/fsxfilecache"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFsxFileCache creates a new instance of [FsxFileCache].
func NewFsxFileCache(name string, args FsxFileCacheArgs) *FsxFileCache {
	return &FsxFileCache{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FsxFileCache)(nil)

// FsxFileCache represents the Terraform resource aws_fsx_file_cache.
type FsxFileCache struct {
	Name      string
	Args      FsxFileCacheArgs
	state     *fsxFileCacheState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FsxFileCache].
func (ffc *FsxFileCache) Type() string {
	return "aws_fsx_file_cache"
}

// LocalName returns the local name for [FsxFileCache].
func (ffc *FsxFileCache) LocalName() string {
	return ffc.Name
}

// Configuration returns the configuration (args) for [FsxFileCache].
func (ffc *FsxFileCache) Configuration() interface{} {
	return ffc.Args
}

// DependOn is used for other resources to depend on [FsxFileCache].
func (ffc *FsxFileCache) DependOn() terra.Reference {
	return terra.ReferenceResource(ffc)
}

// Dependencies returns the list of resources [FsxFileCache] depends_on.
func (ffc *FsxFileCache) Dependencies() terra.Dependencies {
	return ffc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FsxFileCache].
func (ffc *FsxFileCache) LifecycleManagement() *terra.Lifecycle {
	return ffc.Lifecycle
}

// Attributes returns the attributes for [FsxFileCache].
func (ffc *FsxFileCache) Attributes() fsxFileCacheAttributes {
	return fsxFileCacheAttributes{ref: terra.ReferenceResource(ffc)}
}

// ImportState imports the given attribute values into [FsxFileCache]'s state.
func (ffc *FsxFileCache) ImportState(av io.Reader) error {
	ffc.state = &fsxFileCacheState{}
	if err := json.NewDecoder(av).Decode(ffc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ffc.Type(), ffc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FsxFileCache] has state.
func (ffc *FsxFileCache) State() (*fsxFileCacheState, bool) {
	return ffc.state, ffc.state != nil
}

// StateMust returns the state for [FsxFileCache]. Panics if the state is nil.
func (ffc *FsxFileCache) StateMust() *fsxFileCacheState {
	if ffc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ffc.Type(), ffc.LocalName()))
	}
	return ffc.state
}

// FsxFileCacheArgs contains the configurations for aws_fsx_file_cache.
type FsxFileCacheArgs struct {
	// CopyTagsToDataRepositoryAssociations: bool, optional
	CopyTagsToDataRepositoryAssociations terra.BoolValue `hcl:"copy_tags_to_data_repository_associations,attr"`
	// FileCacheType: string, required
	FileCacheType terra.StringValue `hcl:"file_cache_type,attr" validate:"required"`
	// FileCacheTypeVersion: string, required
	FileCacheTypeVersion terra.StringValue `hcl:"file_cache_type_version,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// SecurityGroupIds: set of string, optional
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// StorageCapacity: number, required
	StorageCapacity terra.NumberValue `hcl:"storage_capacity,attr" validate:"required"`
	// SubnetIds: list of string, required
	SubnetIds terra.ListValue[terra.StringValue] `hcl:"subnet_ids,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DataRepositoryAssociation: min=0,max=8
	DataRepositoryAssociation []fsxfilecache.DataRepositoryAssociation `hcl:"data_repository_association,block" validate:"min=0,max=8"`
	// LustreConfiguration: min=0
	LustreConfiguration []fsxfilecache.LustreConfiguration `hcl:"lustre_configuration,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *fsxfilecache.Timeouts `hcl:"timeouts,block"`
}
type fsxFileCacheAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ffc.ref.Append("arn"))
}

// CopyTagsToDataRepositoryAssociations returns a reference to field copy_tags_to_data_repository_associations of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) CopyTagsToDataRepositoryAssociations() terra.BoolValue {
	return terra.ReferenceAsBool(ffc.ref.Append("copy_tags_to_data_repository_associations"))
}

// DataRepositoryAssociationIds returns a reference to field data_repository_association_ids of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) DataRepositoryAssociationIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ffc.ref.Append("data_repository_association_ids"))
}

// DnsName returns a reference to field dns_name of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(ffc.ref.Append("dns_name"))
}

// FileCacheId returns a reference to field file_cache_id of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) FileCacheId() terra.StringValue {
	return terra.ReferenceAsString(ffc.ref.Append("file_cache_id"))
}

// FileCacheType returns a reference to field file_cache_type of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) FileCacheType() terra.StringValue {
	return terra.ReferenceAsString(ffc.ref.Append("file_cache_type"))
}

// FileCacheTypeVersion returns a reference to field file_cache_type_version of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) FileCacheTypeVersion() terra.StringValue {
	return terra.ReferenceAsString(ffc.ref.Append("file_cache_type_version"))
}

// Id returns a reference to field id of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ffc.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ffc.ref.Append("kms_key_id"))
}

// NetworkInterfaceIds returns a reference to field network_interface_ids of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) NetworkInterfaceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ffc.ref.Append("network_interface_ids"))
}

// OwnerId returns a reference to field owner_id of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(ffc.ref.Append("owner_id"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ffc.ref.Append("security_group_ids"))
}

// StorageCapacity returns a reference to field storage_capacity of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) StorageCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(ffc.ref.Append("storage_capacity"))
}

// SubnetIds returns a reference to field subnet_ids of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) SubnetIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ffc.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ffc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ffc.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_fsx_file_cache.
func (ffc fsxFileCacheAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ffc.ref.Append("vpc_id"))
}

func (ffc fsxFileCacheAttributes) DataRepositoryAssociation() terra.SetValue[fsxfilecache.DataRepositoryAssociationAttributes] {
	return terra.ReferenceAsSet[fsxfilecache.DataRepositoryAssociationAttributes](ffc.ref.Append("data_repository_association"))
}

func (ffc fsxFileCacheAttributes) LustreConfiguration() terra.SetValue[fsxfilecache.LustreConfigurationAttributes] {
	return terra.ReferenceAsSet[fsxfilecache.LustreConfigurationAttributes](ffc.ref.Append("lustre_configuration"))
}

func (ffc fsxFileCacheAttributes) Timeouts() fsxfilecache.TimeoutsAttributes {
	return terra.ReferenceAsSingle[fsxfilecache.TimeoutsAttributes](ffc.ref.Append("timeouts"))
}

type fsxFileCacheState struct {
	Arn                                  string                                        `json:"arn"`
	CopyTagsToDataRepositoryAssociations bool                                          `json:"copy_tags_to_data_repository_associations"`
	DataRepositoryAssociationIds         []string                                      `json:"data_repository_association_ids"`
	DnsName                              string                                        `json:"dns_name"`
	FileCacheId                          string                                        `json:"file_cache_id"`
	FileCacheType                        string                                        `json:"file_cache_type"`
	FileCacheTypeVersion                 string                                        `json:"file_cache_type_version"`
	Id                                   string                                        `json:"id"`
	KmsKeyId                             string                                        `json:"kms_key_id"`
	NetworkInterfaceIds                  []string                                      `json:"network_interface_ids"`
	OwnerId                              string                                        `json:"owner_id"`
	SecurityGroupIds                     []string                                      `json:"security_group_ids"`
	StorageCapacity                      float64                                       `json:"storage_capacity"`
	SubnetIds                            []string                                      `json:"subnet_ids"`
	Tags                                 map[string]string                             `json:"tags"`
	TagsAll                              map[string]string                             `json:"tags_all"`
	VpcId                                string                                        `json:"vpc_id"`
	DataRepositoryAssociation            []fsxfilecache.DataRepositoryAssociationState `json:"data_repository_association"`
	LustreConfiguration                  []fsxfilecache.LustreConfigurationState       `json:"lustre_configuration"`
	Timeouts                             *fsxfilecache.TimeoutsState                   `json:"timeouts"`
}