// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	monitoringcustomservice "github.com/golingon/terraproviders/googlebeta/5.24.0/monitoringcustomservice"
	"io"
)

// NewMonitoringCustomService creates a new instance of [MonitoringCustomService].
func NewMonitoringCustomService(name string, args MonitoringCustomServiceArgs) *MonitoringCustomService {
	return &MonitoringCustomService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitoringCustomService)(nil)

// MonitoringCustomService represents the Terraform resource google_monitoring_custom_service.
type MonitoringCustomService struct {
	Name      string
	Args      MonitoringCustomServiceArgs
	state     *monitoringCustomServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitoringCustomService].
func (mcs *MonitoringCustomService) Type() string {
	return "google_monitoring_custom_service"
}

// LocalName returns the local name for [MonitoringCustomService].
func (mcs *MonitoringCustomService) LocalName() string {
	return mcs.Name
}

// Configuration returns the configuration (args) for [MonitoringCustomService].
func (mcs *MonitoringCustomService) Configuration() interface{} {
	return mcs.Args
}

// DependOn is used for other resources to depend on [MonitoringCustomService].
func (mcs *MonitoringCustomService) DependOn() terra.Reference {
	return terra.ReferenceResource(mcs)
}

// Dependencies returns the list of resources [MonitoringCustomService] depends_on.
func (mcs *MonitoringCustomService) Dependencies() terra.Dependencies {
	return mcs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitoringCustomService].
func (mcs *MonitoringCustomService) LifecycleManagement() *terra.Lifecycle {
	return mcs.Lifecycle
}

// Attributes returns the attributes for [MonitoringCustomService].
func (mcs *MonitoringCustomService) Attributes() monitoringCustomServiceAttributes {
	return monitoringCustomServiceAttributes{ref: terra.ReferenceResource(mcs)}
}

// ImportState imports the given attribute values into [MonitoringCustomService]'s state.
func (mcs *MonitoringCustomService) ImportState(av io.Reader) error {
	mcs.state = &monitoringCustomServiceState{}
	if err := json.NewDecoder(av).Decode(mcs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mcs.Type(), mcs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitoringCustomService] has state.
func (mcs *MonitoringCustomService) State() (*monitoringCustomServiceState, bool) {
	return mcs.state, mcs.state != nil
}

// StateMust returns the state for [MonitoringCustomService]. Panics if the state is nil.
func (mcs *MonitoringCustomService) StateMust() *monitoringCustomServiceState {
	if mcs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mcs.Type(), mcs.LocalName()))
	}
	return mcs.state
}

// MonitoringCustomServiceArgs contains the configurations for google_monitoring_custom_service.
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
}
type monitoringCustomServiceAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_monitoring_custom_service.
func (mcs monitoringCustomServiceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(mcs.ref.Append("display_name"))
}

// Id returns a reference to field id of google_monitoring_custom_service.
func (mcs monitoringCustomServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mcs.ref.Append("id"))
}

// Name returns a reference to field name of google_monitoring_custom_service.
func (mcs monitoringCustomServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mcs.ref.Append("name"))
}

// Project returns a reference to field project of google_monitoring_custom_service.
func (mcs monitoringCustomServiceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(mcs.ref.Append("project"))
}

// ServiceId returns a reference to field service_id of google_monitoring_custom_service.
func (mcs monitoringCustomServiceAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(mcs.ref.Append("service_id"))
}

// UserLabels returns a reference to field user_labels of google_monitoring_custom_service.
func (mcs monitoringCustomServiceAttributes) UserLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mcs.ref.Append("user_labels"))
}

func (mcs monitoringCustomServiceAttributes) Telemetry() terra.ListValue[monitoringcustomservice.TelemetryAttributes] {
	return terra.ReferenceAsList[monitoringcustomservice.TelemetryAttributes](mcs.ref.Append("telemetry"))
}

func (mcs monitoringCustomServiceAttributes) Timeouts() monitoringcustomservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitoringcustomservice.TimeoutsAttributes](mcs.ref.Append("timeouts"))
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
