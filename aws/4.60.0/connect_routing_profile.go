// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	connectroutingprofile "github.com/golingon/terraproviders/aws/4.60.0/connectroutingprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConnectRoutingProfile creates a new instance of [ConnectRoutingProfile].
func NewConnectRoutingProfile(name string, args ConnectRoutingProfileArgs) *ConnectRoutingProfile {
	return &ConnectRoutingProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConnectRoutingProfile)(nil)

// ConnectRoutingProfile represents the Terraform resource aws_connect_routing_profile.
type ConnectRoutingProfile struct {
	Name      string
	Args      ConnectRoutingProfileArgs
	state     *connectRoutingProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConnectRoutingProfile].
func (crp *ConnectRoutingProfile) Type() string {
	return "aws_connect_routing_profile"
}

// LocalName returns the local name for [ConnectRoutingProfile].
func (crp *ConnectRoutingProfile) LocalName() string {
	return crp.Name
}

// Configuration returns the configuration (args) for [ConnectRoutingProfile].
func (crp *ConnectRoutingProfile) Configuration() interface{} {
	return crp.Args
}

// DependOn is used for other resources to depend on [ConnectRoutingProfile].
func (crp *ConnectRoutingProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(crp)
}

// Dependencies returns the list of resources [ConnectRoutingProfile] depends_on.
func (crp *ConnectRoutingProfile) Dependencies() terra.Dependencies {
	return crp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConnectRoutingProfile].
func (crp *ConnectRoutingProfile) LifecycleManagement() *terra.Lifecycle {
	return crp.Lifecycle
}

// Attributes returns the attributes for [ConnectRoutingProfile].
func (crp *ConnectRoutingProfile) Attributes() connectRoutingProfileAttributes {
	return connectRoutingProfileAttributes{ref: terra.ReferenceResource(crp)}
}

// ImportState imports the given attribute values into [ConnectRoutingProfile]'s state.
func (crp *ConnectRoutingProfile) ImportState(av io.Reader) error {
	crp.state = &connectRoutingProfileState{}
	if err := json.NewDecoder(av).Decode(crp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crp.Type(), crp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConnectRoutingProfile] has state.
func (crp *ConnectRoutingProfile) State() (*connectRoutingProfileState, bool) {
	return crp.state, crp.state != nil
}

// StateMust returns the state for [ConnectRoutingProfile]. Panics if the state is nil.
func (crp *ConnectRoutingProfile) StateMust() *connectRoutingProfileState {
	if crp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crp.Type(), crp.LocalName()))
	}
	return crp.state
}

// ConnectRoutingProfileArgs contains the configurations for aws_connect_routing_profile.
type ConnectRoutingProfileArgs struct {
	// DefaultOutboundQueueId: string, required
	DefaultOutboundQueueId terra.StringValue `hcl:"default_outbound_queue_id,attr" validate:"required"`
	// Description: string, required
	Description terra.StringValue `hcl:"description,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// QueueConfigsAssociated: min=0
	QueueConfigsAssociated []connectroutingprofile.QueueConfigsAssociated `hcl:"queue_configs_associated,block" validate:"min=0"`
	// MediaConcurrencies: min=1
	MediaConcurrencies []connectroutingprofile.MediaConcurrencies `hcl:"media_concurrencies,block" validate:"min=1"`
	// QueueConfigs: min=0,max=10
	QueueConfigs []connectroutingprofile.QueueConfigs `hcl:"queue_configs,block" validate:"min=0,max=10"`
}
type connectRoutingProfileAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_routing_profile.
func (crp connectRoutingProfileAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("arn"))
}

// DefaultOutboundQueueId returns a reference to field default_outbound_queue_id of aws_connect_routing_profile.
func (crp connectRoutingProfileAttributes) DefaultOutboundQueueId() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("default_outbound_queue_id"))
}

// Description returns a reference to field description of aws_connect_routing_profile.
func (crp connectRoutingProfileAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("description"))
}

// Id returns a reference to field id of aws_connect_routing_profile.
func (crp connectRoutingProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_routing_profile.
func (crp connectRoutingProfileAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("instance_id"))
}

// Name returns a reference to field name of aws_connect_routing_profile.
func (crp connectRoutingProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("name"))
}

// RoutingProfileId returns a reference to field routing_profile_id of aws_connect_routing_profile.
func (crp connectRoutingProfileAttributes) RoutingProfileId() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("routing_profile_id"))
}

// Tags returns a reference to field tags of aws_connect_routing_profile.
func (crp connectRoutingProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_connect_routing_profile.
func (crp connectRoutingProfileAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crp.ref.Append("tags_all"))
}

func (crp connectRoutingProfileAttributes) QueueConfigsAssociated() terra.SetValue[connectroutingprofile.QueueConfigsAssociatedAttributes] {
	return terra.ReferenceAsSet[connectroutingprofile.QueueConfigsAssociatedAttributes](crp.ref.Append("queue_configs_associated"))
}

func (crp connectRoutingProfileAttributes) MediaConcurrencies() terra.SetValue[connectroutingprofile.MediaConcurrenciesAttributes] {
	return terra.ReferenceAsSet[connectroutingprofile.MediaConcurrenciesAttributes](crp.ref.Append("media_concurrencies"))
}

func (crp connectRoutingProfileAttributes) QueueConfigs() terra.SetValue[connectroutingprofile.QueueConfigsAttributes] {
	return terra.ReferenceAsSet[connectroutingprofile.QueueConfigsAttributes](crp.ref.Append("queue_configs"))
}

type connectRoutingProfileState struct {
	Arn                    string                                              `json:"arn"`
	DefaultOutboundQueueId string                                              `json:"default_outbound_queue_id"`
	Description            string                                              `json:"description"`
	Id                     string                                              `json:"id"`
	InstanceId             string                                              `json:"instance_id"`
	Name                   string                                              `json:"name"`
	RoutingProfileId       string                                              `json:"routing_profile_id"`
	Tags                   map[string]string                                   `json:"tags"`
	TagsAll                map[string]string                                   `json:"tags_all"`
	QueueConfigsAssociated []connectroutingprofile.QueueConfigsAssociatedState `json:"queue_configs_associated"`
	MediaConcurrencies     []connectroutingprofile.MediaConcurrenciesState     `json:"media_concurrencies"`
	QueueConfigs           []connectroutingprofile.QueueConfigsState           `json:"queue_configs"`
}
