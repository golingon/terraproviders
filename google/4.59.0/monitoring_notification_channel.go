// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	monitoringnotificationchannel "github.com/golingon/terraproviders/google/4.59.0/monitoringnotificationchannel"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitoringNotificationChannel creates a new instance of [MonitoringNotificationChannel].
func NewMonitoringNotificationChannel(name string, args MonitoringNotificationChannelArgs) *MonitoringNotificationChannel {
	return &MonitoringNotificationChannel{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitoringNotificationChannel)(nil)

// MonitoringNotificationChannel represents the Terraform resource google_monitoring_notification_channel.
type MonitoringNotificationChannel struct {
	Name      string
	Args      MonitoringNotificationChannelArgs
	state     *monitoringNotificationChannelState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitoringNotificationChannel].
func (mnc *MonitoringNotificationChannel) Type() string {
	return "google_monitoring_notification_channel"
}

// LocalName returns the local name for [MonitoringNotificationChannel].
func (mnc *MonitoringNotificationChannel) LocalName() string {
	return mnc.Name
}

// Configuration returns the configuration (args) for [MonitoringNotificationChannel].
func (mnc *MonitoringNotificationChannel) Configuration() interface{} {
	return mnc.Args
}

// DependOn is used for other resources to depend on [MonitoringNotificationChannel].
func (mnc *MonitoringNotificationChannel) DependOn() terra.Reference {
	return terra.ReferenceResource(mnc)
}

// Dependencies returns the list of resources [MonitoringNotificationChannel] depends_on.
func (mnc *MonitoringNotificationChannel) Dependencies() terra.Dependencies {
	return mnc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitoringNotificationChannel].
func (mnc *MonitoringNotificationChannel) LifecycleManagement() *terra.Lifecycle {
	return mnc.Lifecycle
}

// Attributes returns the attributes for [MonitoringNotificationChannel].
func (mnc *MonitoringNotificationChannel) Attributes() monitoringNotificationChannelAttributes {
	return monitoringNotificationChannelAttributes{ref: terra.ReferenceResource(mnc)}
}

// ImportState imports the given attribute values into [MonitoringNotificationChannel]'s state.
func (mnc *MonitoringNotificationChannel) ImportState(av io.Reader) error {
	mnc.state = &monitoringNotificationChannelState{}
	if err := json.NewDecoder(av).Decode(mnc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mnc.Type(), mnc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitoringNotificationChannel] has state.
func (mnc *MonitoringNotificationChannel) State() (*monitoringNotificationChannelState, bool) {
	return mnc.state, mnc.state != nil
}

// StateMust returns the state for [MonitoringNotificationChannel]. Panics if the state is nil.
func (mnc *MonitoringNotificationChannel) StateMust() *monitoringNotificationChannelState {
	if mnc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mnc.Type(), mnc.LocalName()))
	}
	return mnc.state
}

// MonitoringNotificationChannelArgs contains the configurations for google_monitoring_notification_channel.
type MonitoringNotificationChannelArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// ForceDelete: bool, optional
	ForceDelete terra.BoolValue `hcl:"force_delete,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// UserLabels: map of string, optional
	UserLabels terra.MapValue[terra.StringValue] `hcl:"user_labels,attr"`
	// SensitiveLabels: optional
	SensitiveLabels *monitoringnotificationchannel.SensitiveLabels `hcl:"sensitive_labels,block"`
	// Timeouts: optional
	Timeouts *monitoringnotificationchannel.Timeouts `hcl:"timeouts,block"`
}
type monitoringNotificationChannelAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_monitoring_notification_channel.
func (mnc monitoringNotificationChannelAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mnc.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_monitoring_notification_channel.
func (mnc monitoringNotificationChannelAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(mnc.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of google_monitoring_notification_channel.
func (mnc monitoringNotificationChannelAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(mnc.ref.Append("enabled"))
}

// ForceDelete returns a reference to field force_delete of google_monitoring_notification_channel.
func (mnc monitoringNotificationChannelAttributes) ForceDelete() terra.BoolValue {
	return terra.ReferenceAsBool(mnc.ref.Append("force_delete"))
}

// Id returns a reference to field id of google_monitoring_notification_channel.
func (mnc monitoringNotificationChannelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mnc.ref.Append("id"))
}

// Labels returns a reference to field labels of google_monitoring_notification_channel.
func (mnc monitoringNotificationChannelAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mnc.ref.Append("labels"))
}

// Name returns a reference to field name of google_monitoring_notification_channel.
func (mnc monitoringNotificationChannelAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mnc.ref.Append("name"))
}

// Project returns a reference to field project of google_monitoring_notification_channel.
func (mnc monitoringNotificationChannelAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(mnc.ref.Append("project"))
}

// Type returns a reference to field type of google_monitoring_notification_channel.
func (mnc monitoringNotificationChannelAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(mnc.ref.Append("type"))
}

// UserLabels returns a reference to field user_labels of google_monitoring_notification_channel.
func (mnc monitoringNotificationChannelAttributes) UserLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mnc.ref.Append("user_labels"))
}

// VerificationStatus returns a reference to field verification_status of google_monitoring_notification_channel.
func (mnc monitoringNotificationChannelAttributes) VerificationStatus() terra.StringValue {
	return terra.ReferenceAsString(mnc.ref.Append("verification_status"))
}

func (mnc monitoringNotificationChannelAttributes) SensitiveLabels() terra.ListValue[monitoringnotificationchannel.SensitiveLabelsAttributes] {
	return terra.ReferenceAsList[monitoringnotificationchannel.SensitiveLabelsAttributes](mnc.ref.Append("sensitive_labels"))
}

func (mnc monitoringNotificationChannelAttributes) Timeouts() monitoringnotificationchannel.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitoringnotificationchannel.TimeoutsAttributes](mnc.ref.Append("timeouts"))
}

type monitoringNotificationChannelState struct {
	Description        string                                               `json:"description"`
	DisplayName        string                                               `json:"display_name"`
	Enabled            bool                                                 `json:"enabled"`
	ForceDelete        bool                                                 `json:"force_delete"`
	Id                 string                                               `json:"id"`
	Labels             map[string]string                                    `json:"labels"`
	Name               string                                               `json:"name"`
	Project            string                                               `json:"project"`
	Type               string                                               `json:"type"`
	UserLabels         map[string]string                                    `json:"user_labels"`
	VerificationStatus string                                               `json:"verification_status"`
	SensitiveLabels    []monitoringnotificationchannel.SensitiveLabelsState `json:"sensitive_labels"`
	Timeouts           *monitoringnotificationchannel.TimeoutsState         `json:"timeouts"`
}
