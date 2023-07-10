// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	storagegatewaynfsfileshare "github.com/golingon/terraproviders/aws/5.7.0/storagegatewaynfsfileshare"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStoragegatewayNfsFileShare creates a new instance of [StoragegatewayNfsFileShare].
func NewStoragegatewayNfsFileShare(name string, args StoragegatewayNfsFileShareArgs) *StoragegatewayNfsFileShare {
	return &StoragegatewayNfsFileShare{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StoragegatewayNfsFileShare)(nil)

// StoragegatewayNfsFileShare represents the Terraform resource aws_storagegateway_nfs_file_share.
type StoragegatewayNfsFileShare struct {
	Name      string
	Args      StoragegatewayNfsFileShareArgs
	state     *storagegatewayNfsFileShareState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StoragegatewayNfsFileShare].
func (snfs *StoragegatewayNfsFileShare) Type() string {
	return "aws_storagegateway_nfs_file_share"
}

// LocalName returns the local name for [StoragegatewayNfsFileShare].
func (snfs *StoragegatewayNfsFileShare) LocalName() string {
	return snfs.Name
}

// Configuration returns the configuration (args) for [StoragegatewayNfsFileShare].
func (snfs *StoragegatewayNfsFileShare) Configuration() interface{} {
	return snfs.Args
}

// DependOn is used for other resources to depend on [StoragegatewayNfsFileShare].
func (snfs *StoragegatewayNfsFileShare) DependOn() terra.Reference {
	return terra.ReferenceResource(snfs)
}

// Dependencies returns the list of resources [StoragegatewayNfsFileShare] depends_on.
func (snfs *StoragegatewayNfsFileShare) Dependencies() terra.Dependencies {
	return snfs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StoragegatewayNfsFileShare].
func (snfs *StoragegatewayNfsFileShare) LifecycleManagement() *terra.Lifecycle {
	return snfs.Lifecycle
}

// Attributes returns the attributes for [StoragegatewayNfsFileShare].
func (snfs *StoragegatewayNfsFileShare) Attributes() storagegatewayNfsFileShareAttributes {
	return storagegatewayNfsFileShareAttributes{ref: terra.ReferenceResource(snfs)}
}

// ImportState imports the given attribute values into [StoragegatewayNfsFileShare]'s state.
func (snfs *StoragegatewayNfsFileShare) ImportState(av io.Reader) error {
	snfs.state = &storagegatewayNfsFileShareState{}
	if err := json.NewDecoder(av).Decode(snfs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", snfs.Type(), snfs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StoragegatewayNfsFileShare] has state.
func (snfs *StoragegatewayNfsFileShare) State() (*storagegatewayNfsFileShareState, bool) {
	return snfs.state, snfs.state != nil
}

// StateMust returns the state for [StoragegatewayNfsFileShare]. Panics if the state is nil.
func (snfs *StoragegatewayNfsFileShare) StateMust() *storagegatewayNfsFileShareState {
	if snfs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", snfs.Type(), snfs.LocalName()))
	}
	return snfs.state
}

// StoragegatewayNfsFileShareArgs contains the configurations for aws_storagegateway_nfs_file_share.
type StoragegatewayNfsFileShareArgs struct {
	// AuditDestinationArn: string, optional
	AuditDestinationArn terra.StringValue `hcl:"audit_destination_arn,attr"`
	// BucketRegion: string, optional
	BucketRegion terra.StringValue `hcl:"bucket_region,attr"`
	// ClientList: set of string, required
	ClientList terra.SetValue[terra.StringValue] `hcl:"client_list,attr" validate:"required"`
	// DefaultStorageClass: string, optional
	DefaultStorageClass terra.StringValue `hcl:"default_storage_class,attr"`
	// FileShareName: string, optional
	FileShareName terra.StringValue `hcl:"file_share_name,attr"`
	// GatewayArn: string, required
	GatewayArn terra.StringValue `hcl:"gateway_arn,attr" validate:"required"`
	// GuessMimeTypeEnabled: bool, optional
	GuessMimeTypeEnabled terra.BoolValue `hcl:"guess_mime_type_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsEncrypted: bool, optional
	KmsEncrypted terra.BoolValue `hcl:"kms_encrypted,attr"`
	// KmsKeyArn: string, optional
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr"`
	// LocationArn: string, required
	LocationArn terra.StringValue `hcl:"location_arn,attr" validate:"required"`
	// NotificationPolicy: string, optional
	NotificationPolicy terra.StringValue `hcl:"notification_policy,attr"`
	// ObjectAcl: string, optional
	ObjectAcl terra.StringValue `hcl:"object_acl,attr"`
	// ReadOnly: bool, optional
	ReadOnly terra.BoolValue `hcl:"read_only,attr"`
	// RequesterPays: bool, optional
	RequesterPays terra.BoolValue `hcl:"requester_pays,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Squash: string, optional
	Squash terra.StringValue `hcl:"squash,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcEndpointDnsName: string, optional
	VpcEndpointDnsName terra.StringValue `hcl:"vpc_endpoint_dns_name,attr"`
	// CacheAttributes: optional
	CacheAttributes *storagegatewaynfsfileshare.CacheAttributes `hcl:"cache_attributes,block"`
	// NfsFileShareDefaults: optional
	NfsFileShareDefaults *storagegatewaynfsfileshare.NfsFileShareDefaults `hcl:"nfs_file_share_defaults,block"`
	// Timeouts: optional
	Timeouts *storagegatewaynfsfileshare.Timeouts `hcl:"timeouts,block"`
}
type storagegatewayNfsFileShareAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(snfs.ref.Append("arn"))
}

// AuditDestinationArn returns a reference to field audit_destination_arn of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) AuditDestinationArn() terra.StringValue {
	return terra.ReferenceAsString(snfs.ref.Append("audit_destination_arn"))
}

// BucketRegion returns a reference to field bucket_region of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) BucketRegion() terra.StringValue {
	return terra.ReferenceAsString(snfs.ref.Append("bucket_region"))
}

// ClientList returns a reference to field client_list of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) ClientList() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](snfs.ref.Append("client_list"))
}

// DefaultStorageClass returns a reference to field default_storage_class of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) DefaultStorageClass() terra.StringValue {
	return terra.ReferenceAsString(snfs.ref.Append("default_storage_class"))
}

// FileShareName returns a reference to field file_share_name of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) FileShareName() terra.StringValue {
	return terra.ReferenceAsString(snfs.ref.Append("file_share_name"))
}

// FileshareId returns a reference to field fileshare_id of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) FileshareId() terra.StringValue {
	return terra.ReferenceAsString(snfs.ref.Append("fileshare_id"))
}

// GatewayArn returns a reference to field gateway_arn of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) GatewayArn() terra.StringValue {
	return terra.ReferenceAsString(snfs.ref.Append("gateway_arn"))
}

// GuessMimeTypeEnabled returns a reference to field guess_mime_type_enabled of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) GuessMimeTypeEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(snfs.ref.Append("guess_mime_type_enabled"))
}

// Id returns a reference to field id of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(snfs.ref.Append("id"))
}

// KmsEncrypted returns a reference to field kms_encrypted of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) KmsEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(snfs.ref.Append("kms_encrypted"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(snfs.ref.Append("kms_key_arn"))
}

// LocationArn returns a reference to field location_arn of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) LocationArn() terra.StringValue {
	return terra.ReferenceAsString(snfs.ref.Append("location_arn"))
}

// NotificationPolicy returns a reference to field notification_policy of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) NotificationPolicy() terra.StringValue {
	return terra.ReferenceAsString(snfs.ref.Append("notification_policy"))
}

// ObjectAcl returns a reference to field object_acl of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) ObjectAcl() terra.StringValue {
	return terra.ReferenceAsString(snfs.ref.Append("object_acl"))
}

// Path returns a reference to field path of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(snfs.ref.Append("path"))
}

// ReadOnly returns a reference to field read_only of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) ReadOnly() terra.BoolValue {
	return terra.ReferenceAsBool(snfs.ref.Append("read_only"))
}

// RequesterPays returns a reference to field requester_pays of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) RequesterPays() terra.BoolValue {
	return terra.ReferenceAsBool(snfs.ref.Append("requester_pays"))
}

// RoleArn returns a reference to field role_arn of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(snfs.ref.Append("role_arn"))
}

// Squash returns a reference to field squash of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) Squash() terra.StringValue {
	return terra.ReferenceAsString(snfs.ref.Append("squash"))
}

// Tags returns a reference to field tags of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](snfs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](snfs.ref.Append("tags_all"))
}

// VpcEndpointDnsName returns a reference to field vpc_endpoint_dns_name of aws_storagegateway_nfs_file_share.
func (snfs storagegatewayNfsFileShareAttributes) VpcEndpointDnsName() terra.StringValue {
	return terra.ReferenceAsString(snfs.ref.Append("vpc_endpoint_dns_name"))
}

func (snfs storagegatewayNfsFileShareAttributes) CacheAttributes() terra.ListValue[storagegatewaynfsfileshare.CacheAttributesAttributes] {
	return terra.ReferenceAsList[storagegatewaynfsfileshare.CacheAttributesAttributes](snfs.ref.Append("cache_attributes"))
}

func (snfs storagegatewayNfsFileShareAttributes) NfsFileShareDefaults() terra.ListValue[storagegatewaynfsfileshare.NfsFileShareDefaultsAttributes] {
	return terra.ReferenceAsList[storagegatewaynfsfileshare.NfsFileShareDefaultsAttributes](snfs.ref.Append("nfs_file_share_defaults"))
}

func (snfs storagegatewayNfsFileShareAttributes) Timeouts() storagegatewaynfsfileshare.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagegatewaynfsfileshare.TimeoutsAttributes](snfs.ref.Append("timeouts"))
}

type storagegatewayNfsFileShareState struct {
	Arn                  string                                                 `json:"arn"`
	AuditDestinationArn  string                                                 `json:"audit_destination_arn"`
	BucketRegion         string                                                 `json:"bucket_region"`
	ClientList           []string                                               `json:"client_list"`
	DefaultStorageClass  string                                                 `json:"default_storage_class"`
	FileShareName        string                                                 `json:"file_share_name"`
	FileshareId          string                                                 `json:"fileshare_id"`
	GatewayArn           string                                                 `json:"gateway_arn"`
	GuessMimeTypeEnabled bool                                                   `json:"guess_mime_type_enabled"`
	Id                   string                                                 `json:"id"`
	KmsEncrypted         bool                                                   `json:"kms_encrypted"`
	KmsKeyArn            string                                                 `json:"kms_key_arn"`
	LocationArn          string                                                 `json:"location_arn"`
	NotificationPolicy   string                                                 `json:"notification_policy"`
	ObjectAcl            string                                                 `json:"object_acl"`
	Path                 string                                                 `json:"path"`
	ReadOnly             bool                                                   `json:"read_only"`
	RequesterPays        bool                                                   `json:"requester_pays"`
	RoleArn              string                                                 `json:"role_arn"`
	Squash               string                                                 `json:"squash"`
	Tags                 map[string]string                                      `json:"tags"`
	TagsAll              map[string]string                                      `json:"tags_all"`
	VpcEndpointDnsName   string                                                 `json:"vpc_endpoint_dns_name"`
	CacheAttributes      []storagegatewaynfsfileshare.CacheAttributesState      `json:"cache_attributes"`
	NfsFileShareDefaults []storagegatewaynfsfileshare.NfsFileShareDefaultsState `json:"nfs_file_share_defaults"`
	Timeouts             *storagegatewaynfsfileshare.TimeoutsState              `json:"timeouts"`
}
