// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	storagegatewaysmbfileshare "github.com/golingon/terraproviders/aws/5.11.0/storagegatewaysmbfileshare"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStoragegatewaySmbFileShare creates a new instance of [StoragegatewaySmbFileShare].
func NewStoragegatewaySmbFileShare(name string, args StoragegatewaySmbFileShareArgs) *StoragegatewaySmbFileShare {
	return &StoragegatewaySmbFileShare{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StoragegatewaySmbFileShare)(nil)

// StoragegatewaySmbFileShare represents the Terraform resource aws_storagegateway_smb_file_share.
type StoragegatewaySmbFileShare struct {
	Name      string
	Args      StoragegatewaySmbFileShareArgs
	state     *storagegatewaySmbFileShareState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StoragegatewaySmbFileShare].
func (ssfs *StoragegatewaySmbFileShare) Type() string {
	return "aws_storagegateway_smb_file_share"
}

// LocalName returns the local name for [StoragegatewaySmbFileShare].
func (ssfs *StoragegatewaySmbFileShare) LocalName() string {
	return ssfs.Name
}

// Configuration returns the configuration (args) for [StoragegatewaySmbFileShare].
func (ssfs *StoragegatewaySmbFileShare) Configuration() interface{} {
	return ssfs.Args
}

// DependOn is used for other resources to depend on [StoragegatewaySmbFileShare].
func (ssfs *StoragegatewaySmbFileShare) DependOn() terra.Reference {
	return terra.ReferenceResource(ssfs)
}

// Dependencies returns the list of resources [StoragegatewaySmbFileShare] depends_on.
func (ssfs *StoragegatewaySmbFileShare) Dependencies() terra.Dependencies {
	return ssfs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StoragegatewaySmbFileShare].
func (ssfs *StoragegatewaySmbFileShare) LifecycleManagement() *terra.Lifecycle {
	return ssfs.Lifecycle
}

// Attributes returns the attributes for [StoragegatewaySmbFileShare].
func (ssfs *StoragegatewaySmbFileShare) Attributes() storagegatewaySmbFileShareAttributes {
	return storagegatewaySmbFileShareAttributes{ref: terra.ReferenceResource(ssfs)}
}

// ImportState imports the given attribute values into [StoragegatewaySmbFileShare]'s state.
func (ssfs *StoragegatewaySmbFileShare) ImportState(av io.Reader) error {
	ssfs.state = &storagegatewaySmbFileShareState{}
	if err := json.NewDecoder(av).Decode(ssfs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssfs.Type(), ssfs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StoragegatewaySmbFileShare] has state.
func (ssfs *StoragegatewaySmbFileShare) State() (*storagegatewaySmbFileShareState, bool) {
	return ssfs.state, ssfs.state != nil
}

// StateMust returns the state for [StoragegatewaySmbFileShare]. Panics if the state is nil.
func (ssfs *StoragegatewaySmbFileShare) StateMust() *storagegatewaySmbFileShareState {
	if ssfs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssfs.Type(), ssfs.LocalName()))
	}
	return ssfs.state
}

// StoragegatewaySmbFileShareArgs contains the configurations for aws_storagegateway_smb_file_share.
type StoragegatewaySmbFileShareArgs struct {
	// AccessBasedEnumeration: bool, optional
	AccessBasedEnumeration terra.BoolValue `hcl:"access_based_enumeration,attr"`
	// AdminUserList: set of string, optional
	AdminUserList terra.SetValue[terra.StringValue] `hcl:"admin_user_list,attr"`
	// AuditDestinationArn: string, optional
	AuditDestinationArn terra.StringValue `hcl:"audit_destination_arn,attr"`
	// Authentication: string, optional
	Authentication terra.StringValue `hcl:"authentication,attr"`
	// BucketRegion: string, optional
	BucketRegion terra.StringValue `hcl:"bucket_region,attr"`
	// CaseSensitivity: string, optional
	CaseSensitivity terra.StringValue `hcl:"case_sensitivity,attr"`
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
	// InvalidUserList: set of string, optional
	InvalidUserList terra.SetValue[terra.StringValue] `hcl:"invalid_user_list,attr"`
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
	// OplocksEnabled: bool, optional
	OplocksEnabled terra.BoolValue `hcl:"oplocks_enabled,attr"`
	// ReadOnly: bool, optional
	ReadOnly terra.BoolValue `hcl:"read_only,attr"`
	// RequesterPays: bool, optional
	RequesterPays terra.BoolValue `hcl:"requester_pays,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// SmbAclEnabled: bool, optional
	SmbAclEnabled terra.BoolValue `hcl:"smb_acl_enabled,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ValidUserList: set of string, optional
	ValidUserList terra.SetValue[terra.StringValue] `hcl:"valid_user_list,attr"`
	// VpcEndpointDnsName: string, optional
	VpcEndpointDnsName terra.StringValue `hcl:"vpc_endpoint_dns_name,attr"`
	// CacheAttributes: optional
	CacheAttributes *storagegatewaysmbfileshare.CacheAttributes `hcl:"cache_attributes,block"`
	// Timeouts: optional
	Timeouts *storagegatewaysmbfileshare.Timeouts `hcl:"timeouts,block"`
}
type storagegatewaySmbFileShareAttributes struct {
	ref terra.Reference
}

// AccessBasedEnumeration returns a reference to field access_based_enumeration of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) AccessBasedEnumeration() terra.BoolValue {
	return terra.ReferenceAsBool(ssfs.ref.Append("access_based_enumeration"))
}

// AdminUserList returns a reference to field admin_user_list of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) AdminUserList() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ssfs.ref.Append("admin_user_list"))
}

// Arn returns a reference to field arn of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("arn"))
}

// AuditDestinationArn returns a reference to field audit_destination_arn of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) AuditDestinationArn() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("audit_destination_arn"))
}

// Authentication returns a reference to field authentication of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) Authentication() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("authentication"))
}

// BucketRegion returns a reference to field bucket_region of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) BucketRegion() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("bucket_region"))
}

// CaseSensitivity returns a reference to field case_sensitivity of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) CaseSensitivity() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("case_sensitivity"))
}

// DefaultStorageClass returns a reference to field default_storage_class of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) DefaultStorageClass() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("default_storage_class"))
}

// FileShareName returns a reference to field file_share_name of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) FileShareName() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("file_share_name"))
}

// FileshareId returns a reference to field fileshare_id of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) FileshareId() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("fileshare_id"))
}

// GatewayArn returns a reference to field gateway_arn of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) GatewayArn() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("gateway_arn"))
}

// GuessMimeTypeEnabled returns a reference to field guess_mime_type_enabled of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) GuessMimeTypeEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ssfs.ref.Append("guess_mime_type_enabled"))
}

// Id returns a reference to field id of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("id"))
}

// InvalidUserList returns a reference to field invalid_user_list of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) InvalidUserList() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ssfs.ref.Append("invalid_user_list"))
}

// KmsEncrypted returns a reference to field kms_encrypted of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) KmsEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(ssfs.ref.Append("kms_encrypted"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("kms_key_arn"))
}

// LocationArn returns a reference to field location_arn of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) LocationArn() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("location_arn"))
}

// NotificationPolicy returns a reference to field notification_policy of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) NotificationPolicy() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("notification_policy"))
}

// ObjectAcl returns a reference to field object_acl of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) ObjectAcl() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("object_acl"))
}

// OplocksEnabled returns a reference to field oplocks_enabled of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) OplocksEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ssfs.ref.Append("oplocks_enabled"))
}

// Path returns a reference to field path of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("path"))
}

// ReadOnly returns a reference to field read_only of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) ReadOnly() terra.BoolValue {
	return terra.ReferenceAsBool(ssfs.ref.Append("read_only"))
}

// RequesterPays returns a reference to field requester_pays of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) RequesterPays() terra.BoolValue {
	return terra.ReferenceAsBool(ssfs.ref.Append("requester_pays"))
}

// RoleArn returns a reference to field role_arn of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("role_arn"))
}

// SmbAclEnabled returns a reference to field smb_acl_enabled of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) SmbAclEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ssfs.ref.Append("smb_acl_enabled"))
}

// Tags returns a reference to field tags of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ssfs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ssfs.ref.Append("tags_all"))
}

// ValidUserList returns a reference to field valid_user_list of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) ValidUserList() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ssfs.ref.Append("valid_user_list"))
}

// VpcEndpointDnsName returns a reference to field vpc_endpoint_dns_name of aws_storagegateway_smb_file_share.
func (ssfs storagegatewaySmbFileShareAttributes) VpcEndpointDnsName() terra.StringValue {
	return terra.ReferenceAsString(ssfs.ref.Append("vpc_endpoint_dns_name"))
}

func (ssfs storagegatewaySmbFileShareAttributes) CacheAttributes() terra.ListValue[storagegatewaysmbfileshare.CacheAttributesAttributes] {
	return terra.ReferenceAsList[storagegatewaysmbfileshare.CacheAttributesAttributes](ssfs.ref.Append("cache_attributes"))
}

func (ssfs storagegatewaySmbFileShareAttributes) Timeouts() storagegatewaysmbfileshare.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagegatewaysmbfileshare.TimeoutsAttributes](ssfs.ref.Append("timeouts"))
}

type storagegatewaySmbFileShareState struct {
	AccessBasedEnumeration bool                                              `json:"access_based_enumeration"`
	AdminUserList          []string                                          `json:"admin_user_list"`
	Arn                    string                                            `json:"arn"`
	AuditDestinationArn    string                                            `json:"audit_destination_arn"`
	Authentication         string                                            `json:"authentication"`
	BucketRegion           string                                            `json:"bucket_region"`
	CaseSensitivity        string                                            `json:"case_sensitivity"`
	DefaultStorageClass    string                                            `json:"default_storage_class"`
	FileShareName          string                                            `json:"file_share_name"`
	FileshareId            string                                            `json:"fileshare_id"`
	GatewayArn             string                                            `json:"gateway_arn"`
	GuessMimeTypeEnabled   bool                                              `json:"guess_mime_type_enabled"`
	Id                     string                                            `json:"id"`
	InvalidUserList        []string                                          `json:"invalid_user_list"`
	KmsEncrypted           bool                                              `json:"kms_encrypted"`
	KmsKeyArn              string                                            `json:"kms_key_arn"`
	LocationArn            string                                            `json:"location_arn"`
	NotificationPolicy     string                                            `json:"notification_policy"`
	ObjectAcl              string                                            `json:"object_acl"`
	OplocksEnabled         bool                                              `json:"oplocks_enabled"`
	Path                   string                                            `json:"path"`
	ReadOnly               bool                                              `json:"read_only"`
	RequesterPays          bool                                              `json:"requester_pays"`
	RoleArn                string                                            `json:"role_arn"`
	SmbAclEnabled          bool                                              `json:"smb_acl_enabled"`
	Tags                   map[string]string                                 `json:"tags"`
	TagsAll                map[string]string                                 `json:"tags_all"`
	ValidUserList          []string                                          `json:"valid_user_list"`
	VpcEndpointDnsName     string                                            `json:"vpc_endpoint_dns_name"`
	CacheAttributes        []storagegatewaysmbfileshare.CacheAttributesState `json:"cache_attributes"`
	Timeouts               *storagegatewaysmbfileshare.TimeoutsState         `json:"timeouts"`
}
