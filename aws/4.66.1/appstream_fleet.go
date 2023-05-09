// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appstreamfleet "github.com/golingon/terraproviders/aws/4.66.1/appstreamfleet"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppstreamFleet creates a new instance of [AppstreamFleet].
func NewAppstreamFleet(name string, args AppstreamFleetArgs) *AppstreamFleet {
	return &AppstreamFleet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppstreamFleet)(nil)

// AppstreamFleet represents the Terraform resource aws_appstream_fleet.
type AppstreamFleet struct {
	Name      string
	Args      AppstreamFleetArgs
	state     *appstreamFleetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppstreamFleet].
func (af *AppstreamFleet) Type() string {
	return "aws_appstream_fleet"
}

// LocalName returns the local name for [AppstreamFleet].
func (af *AppstreamFleet) LocalName() string {
	return af.Name
}

// Configuration returns the configuration (args) for [AppstreamFleet].
func (af *AppstreamFleet) Configuration() interface{} {
	return af.Args
}

// DependOn is used for other resources to depend on [AppstreamFleet].
func (af *AppstreamFleet) DependOn() terra.Reference {
	return terra.ReferenceResource(af)
}

// Dependencies returns the list of resources [AppstreamFleet] depends_on.
func (af *AppstreamFleet) Dependencies() terra.Dependencies {
	return af.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppstreamFleet].
func (af *AppstreamFleet) LifecycleManagement() *terra.Lifecycle {
	return af.Lifecycle
}

// Attributes returns the attributes for [AppstreamFleet].
func (af *AppstreamFleet) Attributes() appstreamFleetAttributes {
	return appstreamFleetAttributes{ref: terra.ReferenceResource(af)}
}

// ImportState imports the given attribute values into [AppstreamFleet]'s state.
func (af *AppstreamFleet) ImportState(av io.Reader) error {
	af.state = &appstreamFleetState{}
	if err := json.NewDecoder(av).Decode(af.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", af.Type(), af.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppstreamFleet] has state.
func (af *AppstreamFleet) State() (*appstreamFleetState, bool) {
	return af.state, af.state != nil
}

// StateMust returns the state for [AppstreamFleet]. Panics if the state is nil.
func (af *AppstreamFleet) StateMust() *appstreamFleetState {
	if af.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", af.Type(), af.LocalName()))
	}
	return af.state
}

// AppstreamFleetArgs contains the configurations for aws_appstream_fleet.
type AppstreamFleetArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisconnectTimeoutInSeconds: number, optional
	DisconnectTimeoutInSeconds terra.NumberValue `hcl:"disconnect_timeout_in_seconds,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// EnableDefaultInternetAccess: bool, optional
	EnableDefaultInternetAccess terra.BoolValue `hcl:"enable_default_internet_access,attr"`
	// FleetType: string, optional
	FleetType terra.StringValue `hcl:"fleet_type,attr"`
	// IamRoleArn: string, optional
	IamRoleArn terra.StringValue `hcl:"iam_role_arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdleDisconnectTimeoutInSeconds: number, optional
	IdleDisconnectTimeoutInSeconds terra.NumberValue `hcl:"idle_disconnect_timeout_in_seconds,attr"`
	// ImageArn: string, optional
	ImageArn terra.StringValue `hcl:"image_arn,attr"`
	// ImageName: string, optional
	ImageName terra.StringValue `hcl:"image_name,attr"`
	// InstanceType: string, required
	InstanceType terra.StringValue `hcl:"instance_type,attr" validate:"required"`
	// MaxUserDurationInSeconds: number, optional
	MaxUserDurationInSeconds terra.NumberValue `hcl:"max_user_duration_in_seconds,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StreamView: string, optional
	StreamView terra.StringValue `hcl:"stream_view,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ComputeCapacity: required
	ComputeCapacity *appstreamfleet.ComputeCapacity `hcl:"compute_capacity,block" validate:"required"`
	// DomainJoinInfo: optional
	DomainJoinInfo *appstreamfleet.DomainJoinInfo `hcl:"domain_join_info,block"`
	// VpcConfig: optional
	VpcConfig *appstreamfleet.VpcConfig `hcl:"vpc_config,block"`
}
type appstreamFleetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appstream_fleet.
func (af appstreamFleetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("arn"))
}

// CreatedTime returns a reference to field created_time of aws_appstream_fleet.
func (af appstreamFleetAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("created_time"))
}

// Description returns a reference to field description of aws_appstream_fleet.
func (af appstreamFleetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("description"))
}

// DisconnectTimeoutInSeconds returns a reference to field disconnect_timeout_in_seconds of aws_appstream_fleet.
func (af appstreamFleetAttributes) DisconnectTimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(af.ref.Append("disconnect_timeout_in_seconds"))
}

// DisplayName returns a reference to field display_name of aws_appstream_fleet.
func (af appstreamFleetAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("display_name"))
}

// EnableDefaultInternetAccess returns a reference to field enable_default_internet_access of aws_appstream_fleet.
func (af appstreamFleetAttributes) EnableDefaultInternetAccess() terra.BoolValue {
	return terra.ReferenceAsBool(af.ref.Append("enable_default_internet_access"))
}

// FleetType returns a reference to field fleet_type of aws_appstream_fleet.
func (af appstreamFleetAttributes) FleetType() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("fleet_type"))
}

// IamRoleArn returns a reference to field iam_role_arn of aws_appstream_fleet.
func (af appstreamFleetAttributes) IamRoleArn() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("iam_role_arn"))
}

// Id returns a reference to field id of aws_appstream_fleet.
func (af appstreamFleetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("id"))
}

// IdleDisconnectTimeoutInSeconds returns a reference to field idle_disconnect_timeout_in_seconds of aws_appstream_fleet.
func (af appstreamFleetAttributes) IdleDisconnectTimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(af.ref.Append("idle_disconnect_timeout_in_seconds"))
}

// ImageArn returns a reference to field image_arn of aws_appstream_fleet.
func (af appstreamFleetAttributes) ImageArn() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("image_arn"))
}

// ImageName returns a reference to field image_name of aws_appstream_fleet.
func (af appstreamFleetAttributes) ImageName() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("image_name"))
}

// InstanceType returns a reference to field instance_type of aws_appstream_fleet.
func (af appstreamFleetAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("instance_type"))
}

// MaxUserDurationInSeconds returns a reference to field max_user_duration_in_seconds of aws_appstream_fleet.
func (af appstreamFleetAttributes) MaxUserDurationInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(af.ref.Append("max_user_duration_in_seconds"))
}

// Name returns a reference to field name of aws_appstream_fleet.
func (af appstreamFleetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("name"))
}

// State returns a reference to field state of aws_appstream_fleet.
func (af appstreamFleetAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("state"))
}

// StreamView returns a reference to field stream_view of aws_appstream_fleet.
func (af appstreamFleetAttributes) StreamView() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("stream_view"))
}

// Tags returns a reference to field tags of aws_appstream_fleet.
func (af appstreamFleetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](af.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appstream_fleet.
func (af appstreamFleetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](af.ref.Append("tags_all"))
}

func (af appstreamFleetAttributes) ComputeCapacity() terra.ListValue[appstreamfleet.ComputeCapacityAttributes] {
	return terra.ReferenceAsList[appstreamfleet.ComputeCapacityAttributes](af.ref.Append("compute_capacity"))
}

func (af appstreamFleetAttributes) DomainJoinInfo() terra.ListValue[appstreamfleet.DomainJoinInfoAttributes] {
	return terra.ReferenceAsList[appstreamfleet.DomainJoinInfoAttributes](af.ref.Append("domain_join_info"))
}

func (af appstreamFleetAttributes) VpcConfig() terra.ListValue[appstreamfleet.VpcConfigAttributes] {
	return terra.ReferenceAsList[appstreamfleet.VpcConfigAttributes](af.ref.Append("vpc_config"))
}

type appstreamFleetState struct {
	Arn                            string                                `json:"arn"`
	CreatedTime                    string                                `json:"created_time"`
	Description                    string                                `json:"description"`
	DisconnectTimeoutInSeconds     float64                               `json:"disconnect_timeout_in_seconds"`
	DisplayName                    string                                `json:"display_name"`
	EnableDefaultInternetAccess    bool                                  `json:"enable_default_internet_access"`
	FleetType                      string                                `json:"fleet_type"`
	IamRoleArn                     string                                `json:"iam_role_arn"`
	Id                             string                                `json:"id"`
	IdleDisconnectTimeoutInSeconds float64                               `json:"idle_disconnect_timeout_in_seconds"`
	ImageArn                       string                                `json:"image_arn"`
	ImageName                      string                                `json:"image_name"`
	InstanceType                   string                                `json:"instance_type"`
	MaxUserDurationInSeconds       float64                               `json:"max_user_duration_in_seconds"`
	Name                           string                                `json:"name"`
	State                          string                                `json:"state"`
	StreamView                     string                                `json:"stream_view"`
	Tags                           map[string]string                     `json:"tags"`
	TagsAll                        map[string]string                     `json:"tags_all"`
	ComputeCapacity                []appstreamfleet.ComputeCapacityState `json:"compute_capacity"`
	DomainJoinInfo                 []appstreamfleet.DomainJoinInfoState  `json:"domain_join_info"`
	VpcConfig                      []appstreamfleet.VpcConfigState       `json:"vpc_config"`
}
