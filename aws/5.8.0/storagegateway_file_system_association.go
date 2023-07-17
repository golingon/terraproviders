// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	storagegatewayfilesystemassociation "github.com/golingon/terraproviders/aws/5.8.0/storagegatewayfilesystemassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStoragegatewayFileSystemAssociation creates a new instance of [StoragegatewayFileSystemAssociation].
func NewStoragegatewayFileSystemAssociation(name string, args StoragegatewayFileSystemAssociationArgs) *StoragegatewayFileSystemAssociation {
	return &StoragegatewayFileSystemAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StoragegatewayFileSystemAssociation)(nil)

// StoragegatewayFileSystemAssociation represents the Terraform resource aws_storagegateway_file_system_association.
type StoragegatewayFileSystemAssociation struct {
	Name      string
	Args      StoragegatewayFileSystemAssociationArgs
	state     *storagegatewayFileSystemAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StoragegatewayFileSystemAssociation].
func (sfsa *StoragegatewayFileSystemAssociation) Type() string {
	return "aws_storagegateway_file_system_association"
}

// LocalName returns the local name for [StoragegatewayFileSystemAssociation].
func (sfsa *StoragegatewayFileSystemAssociation) LocalName() string {
	return sfsa.Name
}

// Configuration returns the configuration (args) for [StoragegatewayFileSystemAssociation].
func (sfsa *StoragegatewayFileSystemAssociation) Configuration() interface{} {
	return sfsa.Args
}

// DependOn is used for other resources to depend on [StoragegatewayFileSystemAssociation].
func (sfsa *StoragegatewayFileSystemAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(sfsa)
}

// Dependencies returns the list of resources [StoragegatewayFileSystemAssociation] depends_on.
func (sfsa *StoragegatewayFileSystemAssociation) Dependencies() terra.Dependencies {
	return sfsa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StoragegatewayFileSystemAssociation].
func (sfsa *StoragegatewayFileSystemAssociation) LifecycleManagement() *terra.Lifecycle {
	return sfsa.Lifecycle
}

// Attributes returns the attributes for [StoragegatewayFileSystemAssociation].
func (sfsa *StoragegatewayFileSystemAssociation) Attributes() storagegatewayFileSystemAssociationAttributes {
	return storagegatewayFileSystemAssociationAttributes{ref: terra.ReferenceResource(sfsa)}
}

// ImportState imports the given attribute values into [StoragegatewayFileSystemAssociation]'s state.
func (sfsa *StoragegatewayFileSystemAssociation) ImportState(av io.Reader) error {
	sfsa.state = &storagegatewayFileSystemAssociationState{}
	if err := json.NewDecoder(av).Decode(sfsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sfsa.Type(), sfsa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StoragegatewayFileSystemAssociation] has state.
func (sfsa *StoragegatewayFileSystemAssociation) State() (*storagegatewayFileSystemAssociationState, bool) {
	return sfsa.state, sfsa.state != nil
}

// StateMust returns the state for [StoragegatewayFileSystemAssociation]. Panics if the state is nil.
func (sfsa *StoragegatewayFileSystemAssociation) StateMust() *storagegatewayFileSystemAssociationState {
	if sfsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sfsa.Type(), sfsa.LocalName()))
	}
	return sfsa.state
}

// StoragegatewayFileSystemAssociationArgs contains the configurations for aws_storagegateway_file_system_association.
type StoragegatewayFileSystemAssociationArgs struct {
	// AuditDestinationArn: string, optional
	AuditDestinationArn terra.StringValue `hcl:"audit_destination_arn,attr"`
	// GatewayArn: string, required
	GatewayArn terra.StringValue `hcl:"gateway_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocationArn: string, required
	LocationArn terra.StringValue `hcl:"location_arn,attr" validate:"required"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
	// CacheAttributes: optional
	CacheAttributes *storagegatewayfilesystemassociation.CacheAttributes `hcl:"cache_attributes,block"`
}
type storagegatewayFileSystemAssociationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_storagegateway_file_system_association.
func (sfsa storagegatewayFileSystemAssociationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sfsa.ref.Append("arn"))
}

// AuditDestinationArn returns a reference to field audit_destination_arn of aws_storagegateway_file_system_association.
func (sfsa storagegatewayFileSystemAssociationAttributes) AuditDestinationArn() terra.StringValue {
	return terra.ReferenceAsString(sfsa.ref.Append("audit_destination_arn"))
}

// GatewayArn returns a reference to field gateway_arn of aws_storagegateway_file_system_association.
func (sfsa storagegatewayFileSystemAssociationAttributes) GatewayArn() terra.StringValue {
	return terra.ReferenceAsString(sfsa.ref.Append("gateway_arn"))
}

// Id returns a reference to field id of aws_storagegateway_file_system_association.
func (sfsa storagegatewayFileSystemAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sfsa.ref.Append("id"))
}

// LocationArn returns a reference to field location_arn of aws_storagegateway_file_system_association.
func (sfsa storagegatewayFileSystemAssociationAttributes) LocationArn() terra.StringValue {
	return terra.ReferenceAsString(sfsa.ref.Append("location_arn"))
}

// Password returns a reference to field password of aws_storagegateway_file_system_association.
func (sfsa storagegatewayFileSystemAssociationAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(sfsa.ref.Append("password"))
}

// Tags returns a reference to field tags of aws_storagegateway_file_system_association.
func (sfsa storagegatewayFileSystemAssociationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sfsa.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_storagegateway_file_system_association.
func (sfsa storagegatewayFileSystemAssociationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sfsa.ref.Append("tags_all"))
}

// Username returns a reference to field username of aws_storagegateway_file_system_association.
func (sfsa storagegatewayFileSystemAssociationAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(sfsa.ref.Append("username"))
}

func (sfsa storagegatewayFileSystemAssociationAttributes) CacheAttributes() terra.ListValue[storagegatewayfilesystemassociation.CacheAttributesAttributes] {
	return terra.ReferenceAsList[storagegatewayfilesystemassociation.CacheAttributesAttributes](sfsa.ref.Append("cache_attributes"))
}

type storagegatewayFileSystemAssociationState struct {
	Arn                 string                                                     `json:"arn"`
	AuditDestinationArn string                                                     `json:"audit_destination_arn"`
	GatewayArn          string                                                     `json:"gateway_arn"`
	Id                  string                                                     `json:"id"`
	LocationArn         string                                                     `json:"location_arn"`
	Password            string                                                     `json:"password"`
	Tags                map[string]string                                          `json:"tags"`
	TagsAll             map[string]string                                          `json:"tags_all"`
	Username            string                                                     `json:"username"`
	CacheAttributes     []storagegatewayfilesystemassociation.CacheAttributesState `json:"cache_attributes"`
}
