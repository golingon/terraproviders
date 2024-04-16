// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_devicefarm_network_profile

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

// Resource represents the Terraform resource aws_devicefarm_network_profile.
type Resource struct {
	Name      string
	Args      Args
	state     *awsDevicefarmNetworkProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adnp *Resource) Type() string {
	return "aws_devicefarm_network_profile"
}

// LocalName returns the local name for [Resource].
func (adnp *Resource) LocalName() string {
	return adnp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adnp *Resource) Configuration() interface{} {
	return adnp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adnp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adnp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adnp *Resource) Dependencies() terra.Dependencies {
	return adnp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adnp *Resource) LifecycleManagement() *terra.Lifecycle {
	return adnp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adnp *Resource) Attributes() awsDevicefarmNetworkProfileAttributes {
	return awsDevicefarmNetworkProfileAttributes{ref: terra.ReferenceResource(adnp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adnp *Resource) ImportState(state io.Reader) error {
	adnp.state = &awsDevicefarmNetworkProfileState{}
	if err := json.NewDecoder(state).Decode(adnp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adnp.Type(), adnp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adnp *Resource) State() (*awsDevicefarmNetworkProfileState, bool) {
	return adnp.state, adnp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adnp *Resource) StateMust() *awsDevicefarmNetworkProfileState {
	if adnp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adnp.Type(), adnp.LocalName()))
	}
	return adnp.state
}

// Args contains the configurations for aws_devicefarm_network_profile.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DownlinkBandwidthBits: number, optional
	DownlinkBandwidthBits terra.NumberValue `hcl:"downlink_bandwidth_bits,attr"`
	// DownlinkDelayMs: number, optional
	DownlinkDelayMs terra.NumberValue `hcl:"downlink_delay_ms,attr"`
	// DownlinkJitterMs: number, optional
	DownlinkJitterMs terra.NumberValue `hcl:"downlink_jitter_ms,attr"`
	// DownlinkLossPercent: number, optional
	DownlinkLossPercent terra.NumberValue `hcl:"downlink_loss_percent,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProjectArn: string, required
	ProjectArn terra.StringValue `hcl:"project_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// UplinkBandwidthBits: number, optional
	UplinkBandwidthBits terra.NumberValue `hcl:"uplink_bandwidth_bits,attr"`
	// UplinkDelayMs: number, optional
	UplinkDelayMs terra.NumberValue `hcl:"uplink_delay_ms,attr"`
	// UplinkJitterMs: number, optional
	UplinkJitterMs terra.NumberValue `hcl:"uplink_jitter_ms,attr"`
	// UplinkLossPercent: number, optional
	UplinkLossPercent terra.NumberValue `hcl:"uplink_loss_percent,attr"`
}

type awsDevicefarmNetworkProfileAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_devicefarm_network_profile.
func (adnp awsDevicefarmNetworkProfileAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(adnp.ref.Append("arn"))
}

// Description returns a reference to field description of aws_devicefarm_network_profile.
func (adnp awsDevicefarmNetworkProfileAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(adnp.ref.Append("description"))
}

// DownlinkBandwidthBits returns a reference to field downlink_bandwidth_bits of aws_devicefarm_network_profile.
func (adnp awsDevicefarmNetworkProfileAttributes) DownlinkBandwidthBits() terra.NumberValue {
	return terra.ReferenceAsNumber(adnp.ref.Append("downlink_bandwidth_bits"))
}

// DownlinkDelayMs returns a reference to field downlink_delay_ms of aws_devicefarm_network_profile.
func (adnp awsDevicefarmNetworkProfileAttributes) DownlinkDelayMs() terra.NumberValue {
	return terra.ReferenceAsNumber(adnp.ref.Append("downlink_delay_ms"))
}

// DownlinkJitterMs returns a reference to field downlink_jitter_ms of aws_devicefarm_network_profile.
func (adnp awsDevicefarmNetworkProfileAttributes) DownlinkJitterMs() terra.NumberValue {
	return terra.ReferenceAsNumber(adnp.ref.Append("downlink_jitter_ms"))
}

// DownlinkLossPercent returns a reference to field downlink_loss_percent of aws_devicefarm_network_profile.
func (adnp awsDevicefarmNetworkProfileAttributes) DownlinkLossPercent() terra.NumberValue {
	return terra.ReferenceAsNumber(adnp.ref.Append("downlink_loss_percent"))
}

// Id returns a reference to field id of aws_devicefarm_network_profile.
func (adnp awsDevicefarmNetworkProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adnp.ref.Append("id"))
}

// Name returns a reference to field name of aws_devicefarm_network_profile.
func (adnp awsDevicefarmNetworkProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adnp.ref.Append("name"))
}

// ProjectArn returns a reference to field project_arn of aws_devicefarm_network_profile.
func (adnp awsDevicefarmNetworkProfileAttributes) ProjectArn() terra.StringValue {
	return terra.ReferenceAsString(adnp.ref.Append("project_arn"))
}

// Tags returns a reference to field tags of aws_devicefarm_network_profile.
func (adnp awsDevicefarmNetworkProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adnp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_devicefarm_network_profile.
func (adnp awsDevicefarmNetworkProfileAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adnp.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_devicefarm_network_profile.
func (adnp awsDevicefarmNetworkProfileAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(adnp.ref.Append("type"))
}

// UplinkBandwidthBits returns a reference to field uplink_bandwidth_bits of aws_devicefarm_network_profile.
func (adnp awsDevicefarmNetworkProfileAttributes) UplinkBandwidthBits() terra.NumberValue {
	return terra.ReferenceAsNumber(adnp.ref.Append("uplink_bandwidth_bits"))
}

// UplinkDelayMs returns a reference to field uplink_delay_ms of aws_devicefarm_network_profile.
func (adnp awsDevicefarmNetworkProfileAttributes) UplinkDelayMs() terra.NumberValue {
	return terra.ReferenceAsNumber(adnp.ref.Append("uplink_delay_ms"))
}

// UplinkJitterMs returns a reference to field uplink_jitter_ms of aws_devicefarm_network_profile.
func (adnp awsDevicefarmNetworkProfileAttributes) UplinkJitterMs() terra.NumberValue {
	return terra.ReferenceAsNumber(adnp.ref.Append("uplink_jitter_ms"))
}

// UplinkLossPercent returns a reference to field uplink_loss_percent of aws_devicefarm_network_profile.
func (adnp awsDevicefarmNetworkProfileAttributes) UplinkLossPercent() terra.NumberValue {
	return terra.ReferenceAsNumber(adnp.ref.Append("uplink_loss_percent"))
}

type awsDevicefarmNetworkProfileState struct {
	Arn                   string            `json:"arn"`
	Description           string            `json:"description"`
	DownlinkBandwidthBits float64           `json:"downlink_bandwidth_bits"`
	DownlinkDelayMs       float64           `json:"downlink_delay_ms"`
	DownlinkJitterMs      float64           `json:"downlink_jitter_ms"`
	DownlinkLossPercent   float64           `json:"downlink_loss_percent"`
	Id                    string            `json:"id"`
	Name                  string            `json:"name"`
	ProjectArn            string            `json:"project_arn"`
	Tags                  map[string]string `json:"tags"`
	TagsAll               map[string]string `json:"tags_all"`
	Type                  string            `json:"type"`
	UplinkBandwidthBits   float64           `json:"uplink_bandwidth_bits"`
	UplinkDelayMs         float64           `json:"uplink_delay_ms"`
	UplinkJitterMs        float64           `json:"uplink_jitter_ms"`
	UplinkLossPercent     float64           `json:"uplink_loss_percent"`
}
