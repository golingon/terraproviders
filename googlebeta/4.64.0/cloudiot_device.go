// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	cloudiotdevice "github.com/golingon/terraproviders/googlebeta/4.64.0/cloudiotdevice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudiotDevice creates a new instance of [CloudiotDevice].
func NewCloudiotDevice(name string, args CloudiotDeviceArgs) *CloudiotDevice {
	return &CloudiotDevice{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudiotDevice)(nil)

// CloudiotDevice represents the Terraform resource google_cloudiot_device.
type CloudiotDevice struct {
	Name      string
	Args      CloudiotDeviceArgs
	state     *cloudiotDeviceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudiotDevice].
func (cd *CloudiotDevice) Type() string {
	return "google_cloudiot_device"
}

// LocalName returns the local name for [CloudiotDevice].
func (cd *CloudiotDevice) LocalName() string {
	return cd.Name
}

// Configuration returns the configuration (args) for [CloudiotDevice].
func (cd *CloudiotDevice) Configuration() interface{} {
	return cd.Args
}

// DependOn is used for other resources to depend on [CloudiotDevice].
func (cd *CloudiotDevice) DependOn() terra.Reference {
	return terra.ReferenceResource(cd)
}

// Dependencies returns the list of resources [CloudiotDevice] depends_on.
func (cd *CloudiotDevice) Dependencies() terra.Dependencies {
	return cd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudiotDevice].
func (cd *CloudiotDevice) LifecycleManagement() *terra.Lifecycle {
	return cd.Lifecycle
}

// Attributes returns the attributes for [CloudiotDevice].
func (cd *CloudiotDevice) Attributes() cloudiotDeviceAttributes {
	return cloudiotDeviceAttributes{ref: terra.ReferenceResource(cd)}
}

// ImportState imports the given attribute values into [CloudiotDevice]'s state.
func (cd *CloudiotDevice) ImportState(av io.Reader) error {
	cd.state = &cloudiotDeviceState{}
	if err := json.NewDecoder(av).Decode(cd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cd.Type(), cd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudiotDevice] has state.
func (cd *CloudiotDevice) State() (*cloudiotDeviceState, bool) {
	return cd.state, cd.state != nil
}

// StateMust returns the state for [CloudiotDevice]. Panics if the state is nil.
func (cd *CloudiotDevice) StateMust() *cloudiotDeviceState {
	if cd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cd.Type(), cd.LocalName()))
	}
	return cd.state
}

// CloudiotDeviceArgs contains the configurations for google_cloudiot_device.
type CloudiotDeviceArgs struct {
	// Blocked: bool, optional
	Blocked terra.BoolValue `hcl:"blocked,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogLevel: string, optional
	LogLevel terra.StringValue `hcl:"log_level,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Registry: string, required
	Registry terra.StringValue `hcl:"registry,attr" validate:"required"`
	// Config: min=0
	Config []cloudiotdevice.Config `hcl:"config,block" validate:"min=0"`
	// LastErrorStatus: min=0
	LastErrorStatus []cloudiotdevice.LastErrorStatus `hcl:"last_error_status,block" validate:"min=0"`
	// State: min=0
	State []cloudiotdevice.State `hcl:"state,block" validate:"min=0"`
	// Credentials: min=0,max=3
	Credentials []cloudiotdevice.Credentials `hcl:"credentials,block" validate:"min=0,max=3"`
	// GatewayConfig: optional
	GatewayConfig *cloudiotdevice.GatewayConfig `hcl:"gateway_config,block"`
	// Timeouts: optional
	Timeouts *cloudiotdevice.Timeouts `hcl:"timeouts,block"`
}
type cloudiotDeviceAttributes struct {
	ref terra.Reference
}

// Blocked returns a reference to field blocked of google_cloudiot_device.
func (cd cloudiotDeviceAttributes) Blocked() terra.BoolValue {
	return terra.ReferenceAsBool(cd.ref.Append("blocked"))
}

// Id returns a reference to field id of google_cloudiot_device.
func (cd cloudiotDeviceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("id"))
}

// LastConfigAckTime returns a reference to field last_config_ack_time of google_cloudiot_device.
func (cd cloudiotDeviceAttributes) LastConfigAckTime() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("last_config_ack_time"))
}

// LastConfigSendTime returns a reference to field last_config_send_time of google_cloudiot_device.
func (cd cloudiotDeviceAttributes) LastConfigSendTime() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("last_config_send_time"))
}

// LastErrorTime returns a reference to field last_error_time of google_cloudiot_device.
func (cd cloudiotDeviceAttributes) LastErrorTime() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("last_error_time"))
}

// LastEventTime returns a reference to field last_event_time of google_cloudiot_device.
func (cd cloudiotDeviceAttributes) LastEventTime() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("last_event_time"))
}

// LastHeartbeatTime returns a reference to field last_heartbeat_time of google_cloudiot_device.
func (cd cloudiotDeviceAttributes) LastHeartbeatTime() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("last_heartbeat_time"))
}

// LastStateTime returns a reference to field last_state_time of google_cloudiot_device.
func (cd cloudiotDeviceAttributes) LastStateTime() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("last_state_time"))
}

// LogLevel returns a reference to field log_level of google_cloudiot_device.
func (cd cloudiotDeviceAttributes) LogLevel() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("log_level"))
}

// Metadata returns a reference to field metadata of google_cloudiot_device.
func (cd cloudiotDeviceAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cd.ref.Append("metadata"))
}

// Name returns a reference to field name of google_cloudiot_device.
func (cd cloudiotDeviceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("name"))
}

// NumId returns a reference to field num_id of google_cloudiot_device.
func (cd cloudiotDeviceAttributes) NumId() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("num_id"))
}

// Registry returns a reference to field registry of google_cloudiot_device.
func (cd cloudiotDeviceAttributes) Registry() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("registry"))
}

func (cd cloudiotDeviceAttributes) Config() terra.ListValue[cloudiotdevice.ConfigAttributes] {
	return terra.ReferenceAsList[cloudiotdevice.ConfigAttributes](cd.ref.Append("config"))
}

func (cd cloudiotDeviceAttributes) LastErrorStatus() terra.ListValue[cloudiotdevice.LastErrorStatusAttributes] {
	return terra.ReferenceAsList[cloudiotdevice.LastErrorStatusAttributes](cd.ref.Append("last_error_status"))
}

func (cd cloudiotDeviceAttributes) State() terra.ListValue[cloudiotdevice.StateAttributes] {
	return terra.ReferenceAsList[cloudiotdevice.StateAttributes](cd.ref.Append("state"))
}

func (cd cloudiotDeviceAttributes) Credentials() terra.ListValue[cloudiotdevice.CredentialsAttributes] {
	return terra.ReferenceAsList[cloudiotdevice.CredentialsAttributes](cd.ref.Append("credentials"))
}

func (cd cloudiotDeviceAttributes) GatewayConfig() terra.ListValue[cloudiotdevice.GatewayConfigAttributes] {
	return terra.ReferenceAsList[cloudiotdevice.GatewayConfigAttributes](cd.ref.Append("gateway_config"))
}

func (cd cloudiotDeviceAttributes) Timeouts() cloudiotdevice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudiotdevice.TimeoutsAttributes](cd.ref.Append("timeouts"))
}

type cloudiotDeviceState struct {
	Blocked            bool                                  `json:"blocked"`
	Id                 string                                `json:"id"`
	LastConfigAckTime  string                                `json:"last_config_ack_time"`
	LastConfigSendTime string                                `json:"last_config_send_time"`
	LastErrorTime      string                                `json:"last_error_time"`
	LastEventTime      string                                `json:"last_event_time"`
	LastHeartbeatTime  string                                `json:"last_heartbeat_time"`
	LastStateTime      string                                `json:"last_state_time"`
	LogLevel           string                                `json:"log_level"`
	Metadata           map[string]string                     `json:"metadata"`
	Name               string                                `json:"name"`
	NumId              string                                `json:"num_id"`
	Registry           string                                `json:"registry"`
	Config             []cloudiotdevice.ConfigState          `json:"config"`
	LastErrorStatus    []cloudiotdevice.LastErrorStatusState `json:"last_error_status"`
	State              []cloudiotdevice.StateState           `json:"state"`
	Credentials        []cloudiotdevice.CredentialsState     `json:"credentials"`
	GatewayConfig      []cloudiotdevice.GatewayConfigState   `json:"gateway_config"`
	Timeouts           *cloudiotdevice.TimeoutsState         `json:"timeouts"`
}
