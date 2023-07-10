// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataconnectroutingprofile "github.com/golingon/terraproviders/aws/5.7.0/dataconnectroutingprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataConnectRoutingProfile creates a new instance of [DataConnectRoutingProfile].
func NewDataConnectRoutingProfile(name string, args DataConnectRoutingProfileArgs) *DataConnectRoutingProfile {
	return &DataConnectRoutingProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataConnectRoutingProfile)(nil)

// DataConnectRoutingProfile represents the Terraform data resource aws_connect_routing_profile.
type DataConnectRoutingProfile struct {
	Name string
	Args DataConnectRoutingProfileArgs
}

// DataSource returns the Terraform object type for [DataConnectRoutingProfile].
func (crp *DataConnectRoutingProfile) DataSource() string {
	return "aws_connect_routing_profile"
}

// LocalName returns the local name for [DataConnectRoutingProfile].
func (crp *DataConnectRoutingProfile) LocalName() string {
	return crp.Name
}

// Configuration returns the configuration (args) for [DataConnectRoutingProfile].
func (crp *DataConnectRoutingProfile) Configuration() interface{} {
	return crp.Args
}

// Attributes returns the attributes for [DataConnectRoutingProfile].
func (crp *DataConnectRoutingProfile) Attributes() dataConnectRoutingProfileAttributes {
	return dataConnectRoutingProfileAttributes{ref: terra.ReferenceDataResource(crp)}
}

// DataConnectRoutingProfileArgs contains the configurations for aws_connect_routing_profile.
type DataConnectRoutingProfileArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// RoutingProfileId: string, optional
	RoutingProfileId terra.StringValue `hcl:"routing_profile_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// MediaConcurrencies: min=0
	MediaConcurrencies []dataconnectroutingprofile.MediaConcurrencies `hcl:"media_concurrencies,block" validate:"min=0"`
	// QueueConfigs: min=0
	QueueConfigs []dataconnectroutingprofile.QueueConfigs `hcl:"queue_configs,block" validate:"min=0"`
}
type dataConnectRoutingProfileAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_routing_profile.
func (crp dataConnectRoutingProfileAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("arn"))
}

// DefaultOutboundQueueId returns a reference to field default_outbound_queue_id of aws_connect_routing_profile.
func (crp dataConnectRoutingProfileAttributes) DefaultOutboundQueueId() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("default_outbound_queue_id"))
}

// Description returns a reference to field description of aws_connect_routing_profile.
func (crp dataConnectRoutingProfileAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("description"))
}

// Id returns a reference to field id of aws_connect_routing_profile.
func (crp dataConnectRoutingProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_routing_profile.
func (crp dataConnectRoutingProfileAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("instance_id"))
}

// Name returns a reference to field name of aws_connect_routing_profile.
func (crp dataConnectRoutingProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("name"))
}

// RoutingProfileId returns a reference to field routing_profile_id of aws_connect_routing_profile.
func (crp dataConnectRoutingProfileAttributes) RoutingProfileId() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("routing_profile_id"))
}

// Tags returns a reference to field tags of aws_connect_routing_profile.
func (crp dataConnectRoutingProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crp.ref.Append("tags"))
}

func (crp dataConnectRoutingProfileAttributes) MediaConcurrencies() terra.SetValue[dataconnectroutingprofile.MediaConcurrenciesAttributes] {
	return terra.ReferenceAsSet[dataconnectroutingprofile.MediaConcurrenciesAttributes](crp.ref.Append("media_concurrencies"))
}

func (crp dataConnectRoutingProfileAttributes) QueueConfigs() terra.SetValue[dataconnectroutingprofile.QueueConfigsAttributes] {
	return terra.ReferenceAsSet[dataconnectroutingprofile.QueueConfigsAttributes](crp.ref.Append("queue_configs"))
}
