// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_storagegateway_file_system_association

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_storagegateway_file_system_association.
type Resource struct {
	Name      string
	Args      Args
	state     *awsStoragegatewayFileSystemAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asfsa *Resource) Type() string {
	return "aws_storagegateway_file_system_association"
}

// LocalName returns the local name for [Resource].
func (asfsa *Resource) LocalName() string {
	return asfsa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asfsa *Resource) Configuration() interface{} {
	return asfsa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asfsa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asfsa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asfsa *Resource) Dependencies() terra.Dependencies {
	return asfsa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asfsa *Resource) LifecycleManagement() *terra.Lifecycle {
	return asfsa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asfsa *Resource) Attributes() awsStoragegatewayFileSystemAssociationAttributes {
	return awsStoragegatewayFileSystemAssociationAttributes{ref: terra.ReferenceResource(asfsa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asfsa *Resource) ImportState(state io.Reader) error {
	asfsa.state = &awsStoragegatewayFileSystemAssociationState{}
	if err := json.NewDecoder(state).Decode(asfsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asfsa.Type(), asfsa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asfsa *Resource) State() (*awsStoragegatewayFileSystemAssociationState, bool) {
	return asfsa.state, asfsa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asfsa *Resource) StateMust() *awsStoragegatewayFileSystemAssociationState {
	if asfsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asfsa.Type(), asfsa.LocalName()))
	}
	return asfsa.state
}

// Args contains the configurations for aws_storagegateway_file_system_association.
type Args struct {
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
	CacheAttributes *CacheAttributes `hcl:"cache_attributes,block"`
}

type awsStoragegatewayFileSystemAssociationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_storagegateway_file_system_association.
func (asfsa awsStoragegatewayFileSystemAssociationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(asfsa.ref.Append("arn"))
}

// AuditDestinationArn returns a reference to field audit_destination_arn of aws_storagegateway_file_system_association.
func (asfsa awsStoragegatewayFileSystemAssociationAttributes) AuditDestinationArn() terra.StringValue {
	return terra.ReferenceAsString(asfsa.ref.Append("audit_destination_arn"))
}

// GatewayArn returns a reference to field gateway_arn of aws_storagegateway_file_system_association.
func (asfsa awsStoragegatewayFileSystemAssociationAttributes) GatewayArn() terra.StringValue {
	return terra.ReferenceAsString(asfsa.ref.Append("gateway_arn"))
}

// Id returns a reference to field id of aws_storagegateway_file_system_association.
func (asfsa awsStoragegatewayFileSystemAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asfsa.ref.Append("id"))
}

// LocationArn returns a reference to field location_arn of aws_storagegateway_file_system_association.
func (asfsa awsStoragegatewayFileSystemAssociationAttributes) LocationArn() terra.StringValue {
	return terra.ReferenceAsString(asfsa.ref.Append("location_arn"))
}

// Password returns a reference to field password of aws_storagegateway_file_system_association.
func (asfsa awsStoragegatewayFileSystemAssociationAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(asfsa.ref.Append("password"))
}

// Tags returns a reference to field tags of aws_storagegateway_file_system_association.
func (asfsa awsStoragegatewayFileSystemAssociationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asfsa.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_storagegateway_file_system_association.
func (asfsa awsStoragegatewayFileSystemAssociationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asfsa.ref.Append("tags_all"))
}

// Username returns a reference to field username of aws_storagegateway_file_system_association.
func (asfsa awsStoragegatewayFileSystemAssociationAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(asfsa.ref.Append("username"))
}

func (asfsa awsStoragegatewayFileSystemAssociationAttributes) CacheAttributes() terra.ListValue[CacheAttributesAttributes] {
	return terra.ReferenceAsList[CacheAttributesAttributes](asfsa.ref.Append("cache_attributes"))
}

type awsStoragegatewayFileSystemAssociationState struct {
	Arn                 string                 `json:"arn"`
	AuditDestinationArn string                 `json:"audit_destination_arn"`
	GatewayArn          string                 `json:"gateway_arn"`
	Id                  string                 `json:"id"`
	LocationArn         string                 `json:"location_arn"`
	Password            string                 `json:"password"`
	Tags                map[string]string      `json:"tags"`
	TagsAll             map[string]string      `json:"tags_all"`
	Username            string                 `json:"username"`
	CacheAttributes     []CacheAttributesState `json:"cache_attributes"`
}
