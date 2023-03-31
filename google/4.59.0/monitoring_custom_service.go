// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	monitoringcustomservice "github.com/golingon/terraproviders/google/4.59.0/monitoringcustomservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewMonitoringCustomService(name string, args MonitoringCustomServiceArgs) *MonitoringCustomService {
	return &MonitoringCustomService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitoringCustomService)(nil)

type MonitoringCustomService struct {
	Name  string
	Args  MonitoringCustomServiceArgs
	state *monitoringCustomServiceState
}

func (mcs *MonitoringCustomService) Type() string {
	return "google_monitoring_custom_service"
}

func (mcs *MonitoringCustomService) LocalName() string {
	return mcs.Name
}

func (mcs *MonitoringCustomService) Configuration() interface{} {
	return mcs.Args
}

func (mcs *MonitoringCustomService) Attributes() monitoringCustomServiceAttributes {
	return monitoringCustomServiceAttributes{ref: terra.ReferenceResource(mcs)}
}

func (mcs *MonitoringCustomService) ImportState(av io.Reader) error {
	mcs.state = &monitoringCustomServiceState{}
	if err := json.NewDecoder(av).Decode(mcs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mcs.Type(), mcs.LocalName(), err)
	}
	return nil
}

func (mcs *MonitoringCustomService) State() (*monitoringCustomServiceState, bool) {
	return mcs.state, mcs.state != nil
}

func (mcs *MonitoringCustomService) StateMust() *monitoringCustomServiceState {
	if mcs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mcs.Type(), mcs.LocalName()))
	}
	return mcs.state
}

func (mcs *MonitoringCustomService) DependOn() terra.Reference {
	return terra.ReferenceResource(mcs)
}

type MonitoringCustomServiceArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ServiceId: string, optional
	ServiceId terra.StringValue `hcl:"service_id,attr"`
	// UserLabels: map of string, optional
	UserLabels terra.MapValue[terra.StringValue] `hcl:"user_labels,attr"`
	// Telemetry: optional
	Telemetry *monitoringcustomservice.Telemetry `hcl:"telemetry,block"`
	// Timeouts: optional
	Timeouts *monitoringcustomservice.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that MonitoringCustomService depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type monitoringCustomServiceAttributes struct {
	ref terra.Reference
}

func (mcs monitoringCustomServiceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(mcs.ref.Append("display_name"))
}

func (mcs monitoringCustomServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(mcs.ref.Append("id"))
}

func (mcs monitoringCustomServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(mcs.ref.Append("name"))
}

func (mcs monitoringCustomServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceString(mcs.ref.Append("project"))
}

func (mcs monitoringCustomServiceAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceString(mcs.ref.Append("service_id"))
}

func (mcs monitoringCustomServiceAttributes) UserLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](mcs.ref.Append("user_labels"))
}

func (mcs monitoringCustomServiceAttributes) Telemetry() terra.ListValue[monitoringcustomservice.TelemetryAttributes] {
	return terra.ReferenceList[monitoringcustomservice.TelemetryAttributes](mcs.ref.Append("telemetry"))
}

func (mcs monitoringCustomServiceAttributes) Timeouts() monitoringcustomservice.TimeoutsAttributes {
	return terra.ReferenceSingle[monitoringcustomservice.TimeoutsAttributes](mcs.ref.Append("timeouts"))
}

type monitoringCustomServiceState struct {
	DisplayName string                                   `json:"display_name"`
	Id          string                                   `json:"id"`
	Name        string                                   `json:"name"`
	Project     string                                   `json:"project"`
	ServiceId   string                                   `json:"service_id"`
	UserLabels  map[string]string                        `json:"user_labels"`
	Telemetry   []monitoringcustomservice.TelemetryState `json:"telemetry"`
	Timeouts    *monitoringcustomservice.TimeoutsState   `json:"timeouts"`
}
