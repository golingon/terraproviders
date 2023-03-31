// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudiotdevice "github.com/golingon/terraproviders/google/4.59.0/cloudiotdevice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewCloudiotDevice(name string, args CloudiotDeviceArgs) *CloudiotDevice {
	return &CloudiotDevice{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudiotDevice)(nil)

type CloudiotDevice struct {
	Name  string
	Args  CloudiotDeviceArgs
	state *cloudiotDeviceState
}

func (cd *CloudiotDevice) Type() string {
	return "google_cloudiot_device"
}

func (cd *CloudiotDevice) LocalName() string {
	return cd.Name
}

func (cd *CloudiotDevice) Configuration() interface{} {
	return cd.Args
}

func (cd *CloudiotDevice) Attributes() cloudiotDeviceAttributes {
	return cloudiotDeviceAttributes{ref: terra.ReferenceResource(cd)}
}

func (cd *CloudiotDevice) ImportState(av io.Reader) error {
	cd.state = &cloudiotDeviceState{}
	if err := json.NewDecoder(av).Decode(cd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cd.Type(), cd.LocalName(), err)
	}
	return nil
}

func (cd *CloudiotDevice) State() (*cloudiotDeviceState, bool) {
	return cd.state, cd.state != nil
}

func (cd *CloudiotDevice) StateMust() *cloudiotDeviceState {
	if cd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cd.Type(), cd.LocalName()))
	}
	return cd.state
}

func (cd *CloudiotDevice) DependOn() terra.Reference {
	return terra.ReferenceResource(cd)
}

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
	// DependsOn contains resources that CloudiotDevice depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type cloudiotDeviceAttributes struct {
	ref terra.Reference
}

func (cd cloudiotDeviceAttributes) Blocked() terra.BoolValue {
	return terra.ReferenceBool(cd.ref.Append("blocked"))
}

func (cd cloudiotDeviceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("id"))
}

func (cd cloudiotDeviceAttributes) LastConfigAckTime() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("last_config_ack_time"))
}

func (cd cloudiotDeviceAttributes) LastConfigSendTime() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("last_config_send_time"))
}

func (cd cloudiotDeviceAttributes) LastErrorTime() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("last_error_time"))
}

func (cd cloudiotDeviceAttributes) LastEventTime() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("last_event_time"))
}

func (cd cloudiotDeviceAttributes) LastHeartbeatTime() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("last_heartbeat_time"))
}

func (cd cloudiotDeviceAttributes) LastStateTime() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("last_state_time"))
}

func (cd cloudiotDeviceAttributes) LogLevel() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("log_level"))
}

func (cd cloudiotDeviceAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cd.ref.Append("metadata"))
}

func (cd cloudiotDeviceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("name"))
}

func (cd cloudiotDeviceAttributes) NumId() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("num_id"))
}

func (cd cloudiotDeviceAttributes) Registry() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("registry"))
}

func (cd cloudiotDeviceAttributes) Config() terra.ListValue[cloudiotdevice.ConfigAttributes] {
	return terra.ReferenceList[cloudiotdevice.ConfigAttributes](cd.ref.Append("config"))
}

func (cd cloudiotDeviceAttributes) LastErrorStatus() terra.ListValue[cloudiotdevice.LastErrorStatusAttributes] {
	return terra.ReferenceList[cloudiotdevice.LastErrorStatusAttributes](cd.ref.Append("last_error_status"))
}

func (cd cloudiotDeviceAttributes) State() terra.ListValue[cloudiotdevice.StateAttributes] {
	return terra.ReferenceList[cloudiotdevice.StateAttributes](cd.ref.Append("state"))
}

func (cd cloudiotDeviceAttributes) Credentials() terra.ListValue[cloudiotdevice.CredentialsAttributes] {
	return terra.ReferenceList[cloudiotdevice.CredentialsAttributes](cd.ref.Append("credentials"))
}

func (cd cloudiotDeviceAttributes) GatewayConfig() terra.ListValue[cloudiotdevice.GatewayConfigAttributes] {
	return terra.ReferenceList[cloudiotdevice.GatewayConfigAttributes](cd.ref.Append("gateway_config"))
}

func (cd cloudiotDeviceAttributes) Timeouts() cloudiotdevice.TimeoutsAttributes {
	return terra.ReferenceSingle[cloudiotdevice.TimeoutsAttributes](cd.ref.Append("timeouts"))
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
