// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appstreamimagebuilder "github.com/golingon/terraproviders/aws/5.7.0/appstreamimagebuilder"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppstreamImageBuilder creates a new instance of [AppstreamImageBuilder].
func NewAppstreamImageBuilder(name string, args AppstreamImageBuilderArgs) *AppstreamImageBuilder {
	return &AppstreamImageBuilder{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppstreamImageBuilder)(nil)

// AppstreamImageBuilder represents the Terraform resource aws_appstream_image_builder.
type AppstreamImageBuilder struct {
	Name      string
	Args      AppstreamImageBuilderArgs
	state     *appstreamImageBuilderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppstreamImageBuilder].
func (aib *AppstreamImageBuilder) Type() string {
	return "aws_appstream_image_builder"
}

// LocalName returns the local name for [AppstreamImageBuilder].
func (aib *AppstreamImageBuilder) LocalName() string {
	return aib.Name
}

// Configuration returns the configuration (args) for [AppstreamImageBuilder].
func (aib *AppstreamImageBuilder) Configuration() interface{} {
	return aib.Args
}

// DependOn is used for other resources to depend on [AppstreamImageBuilder].
func (aib *AppstreamImageBuilder) DependOn() terra.Reference {
	return terra.ReferenceResource(aib)
}

// Dependencies returns the list of resources [AppstreamImageBuilder] depends_on.
func (aib *AppstreamImageBuilder) Dependencies() terra.Dependencies {
	return aib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppstreamImageBuilder].
func (aib *AppstreamImageBuilder) LifecycleManagement() *terra.Lifecycle {
	return aib.Lifecycle
}

// Attributes returns the attributes for [AppstreamImageBuilder].
func (aib *AppstreamImageBuilder) Attributes() appstreamImageBuilderAttributes {
	return appstreamImageBuilderAttributes{ref: terra.ReferenceResource(aib)}
}

// ImportState imports the given attribute values into [AppstreamImageBuilder]'s state.
func (aib *AppstreamImageBuilder) ImportState(av io.Reader) error {
	aib.state = &appstreamImageBuilderState{}
	if err := json.NewDecoder(av).Decode(aib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aib.Type(), aib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppstreamImageBuilder] has state.
func (aib *AppstreamImageBuilder) State() (*appstreamImageBuilderState, bool) {
	return aib.state, aib.state != nil
}

// StateMust returns the state for [AppstreamImageBuilder]. Panics if the state is nil.
func (aib *AppstreamImageBuilder) StateMust() *appstreamImageBuilderState {
	if aib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aib.Type(), aib.LocalName()))
	}
	return aib.state
}

// AppstreamImageBuilderArgs contains the configurations for aws_appstream_image_builder.
type AppstreamImageBuilderArgs struct {
	// AppstreamAgentVersion: string, optional
	AppstreamAgentVersion terra.StringValue `hcl:"appstream_agent_version,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// EnableDefaultInternetAccess: bool, optional
	EnableDefaultInternetAccess terra.BoolValue `hcl:"enable_default_internet_access,attr"`
	// IamRoleArn: string, optional
	IamRoleArn terra.StringValue `hcl:"iam_role_arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImageArn: string, optional
	ImageArn terra.StringValue `hcl:"image_arn,attr"`
	// ImageName: string, optional
	ImageName terra.StringValue `hcl:"image_name,attr"`
	// InstanceType: string, required
	InstanceType terra.StringValue `hcl:"instance_type,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// AccessEndpoint: min=0,max=4
	AccessEndpoint []appstreamimagebuilder.AccessEndpoint `hcl:"access_endpoint,block" validate:"min=0,max=4"`
	// DomainJoinInfo: optional
	DomainJoinInfo *appstreamimagebuilder.DomainJoinInfo `hcl:"domain_join_info,block"`
	// VpcConfig: optional
	VpcConfig *appstreamimagebuilder.VpcConfig `hcl:"vpc_config,block"`
}
type appstreamImageBuilderAttributes struct {
	ref terra.Reference
}

// AppstreamAgentVersion returns a reference to field appstream_agent_version of aws_appstream_image_builder.
func (aib appstreamImageBuilderAttributes) AppstreamAgentVersion() terra.StringValue {
	return terra.ReferenceAsString(aib.ref.Append("appstream_agent_version"))
}

// Arn returns a reference to field arn of aws_appstream_image_builder.
func (aib appstreamImageBuilderAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aib.ref.Append("arn"))
}

// CreatedTime returns a reference to field created_time of aws_appstream_image_builder.
func (aib appstreamImageBuilderAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(aib.ref.Append("created_time"))
}

// Description returns a reference to field description of aws_appstream_image_builder.
func (aib appstreamImageBuilderAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aib.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of aws_appstream_image_builder.
func (aib appstreamImageBuilderAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(aib.ref.Append("display_name"))
}

// EnableDefaultInternetAccess returns a reference to field enable_default_internet_access of aws_appstream_image_builder.
func (aib appstreamImageBuilderAttributes) EnableDefaultInternetAccess() terra.BoolValue {
	return terra.ReferenceAsBool(aib.ref.Append("enable_default_internet_access"))
}

// IamRoleArn returns a reference to field iam_role_arn of aws_appstream_image_builder.
func (aib appstreamImageBuilderAttributes) IamRoleArn() terra.StringValue {
	return terra.ReferenceAsString(aib.ref.Append("iam_role_arn"))
}

// Id returns a reference to field id of aws_appstream_image_builder.
func (aib appstreamImageBuilderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aib.ref.Append("id"))
}

// ImageArn returns a reference to field image_arn of aws_appstream_image_builder.
func (aib appstreamImageBuilderAttributes) ImageArn() terra.StringValue {
	return terra.ReferenceAsString(aib.ref.Append("image_arn"))
}

// ImageName returns a reference to field image_name of aws_appstream_image_builder.
func (aib appstreamImageBuilderAttributes) ImageName() terra.StringValue {
	return terra.ReferenceAsString(aib.ref.Append("image_name"))
}

// InstanceType returns a reference to field instance_type of aws_appstream_image_builder.
func (aib appstreamImageBuilderAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(aib.ref.Append("instance_type"))
}

// Name returns a reference to field name of aws_appstream_image_builder.
func (aib appstreamImageBuilderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aib.ref.Append("name"))
}

// State returns a reference to field state of aws_appstream_image_builder.
func (aib appstreamImageBuilderAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(aib.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_appstream_image_builder.
func (aib appstreamImageBuilderAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aib.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appstream_image_builder.
func (aib appstreamImageBuilderAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aib.ref.Append("tags_all"))
}

func (aib appstreamImageBuilderAttributes) AccessEndpoint() terra.SetValue[appstreamimagebuilder.AccessEndpointAttributes] {
	return terra.ReferenceAsSet[appstreamimagebuilder.AccessEndpointAttributes](aib.ref.Append("access_endpoint"))
}

func (aib appstreamImageBuilderAttributes) DomainJoinInfo() terra.ListValue[appstreamimagebuilder.DomainJoinInfoAttributes] {
	return terra.ReferenceAsList[appstreamimagebuilder.DomainJoinInfoAttributes](aib.ref.Append("domain_join_info"))
}

func (aib appstreamImageBuilderAttributes) VpcConfig() terra.ListValue[appstreamimagebuilder.VpcConfigAttributes] {
	return terra.ReferenceAsList[appstreamimagebuilder.VpcConfigAttributes](aib.ref.Append("vpc_config"))
}

type appstreamImageBuilderState struct {
	AppstreamAgentVersion       string                                      `json:"appstream_agent_version"`
	Arn                         string                                      `json:"arn"`
	CreatedTime                 string                                      `json:"created_time"`
	Description                 string                                      `json:"description"`
	DisplayName                 string                                      `json:"display_name"`
	EnableDefaultInternetAccess bool                                        `json:"enable_default_internet_access"`
	IamRoleArn                  string                                      `json:"iam_role_arn"`
	Id                          string                                      `json:"id"`
	ImageArn                    string                                      `json:"image_arn"`
	ImageName                   string                                      `json:"image_name"`
	InstanceType                string                                      `json:"instance_type"`
	Name                        string                                      `json:"name"`
	State                       string                                      `json:"state"`
	Tags                        map[string]string                           `json:"tags"`
	TagsAll                     map[string]string                           `json:"tags_all"`
	AccessEndpoint              []appstreamimagebuilder.AccessEndpointState `json:"access_endpoint"`
	DomainJoinInfo              []appstreamimagebuilder.DomainJoinInfoState `json:"domain_join_info"`
	VpcConfig                   []appstreamimagebuilder.VpcConfigState      `json:"vpc_config"`
}
