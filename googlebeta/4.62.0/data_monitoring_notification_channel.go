// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datamonitoringnotificationchannel "github.com/golingon/terraproviders/googlebeta/4.62.0/datamonitoringnotificationchannel"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMonitoringNotificationChannel creates a new instance of [DataMonitoringNotificationChannel].
func NewDataMonitoringNotificationChannel(name string, args DataMonitoringNotificationChannelArgs) *DataMonitoringNotificationChannel {
	return &DataMonitoringNotificationChannel{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMonitoringNotificationChannel)(nil)

// DataMonitoringNotificationChannel represents the Terraform data resource google_monitoring_notification_channel.
type DataMonitoringNotificationChannel struct {
	Name string
	Args DataMonitoringNotificationChannelArgs
}

// DataSource returns the Terraform object type for [DataMonitoringNotificationChannel].
func (mnc *DataMonitoringNotificationChannel) DataSource() string {
	return "google_monitoring_notification_channel"
}

// LocalName returns the local name for [DataMonitoringNotificationChannel].
func (mnc *DataMonitoringNotificationChannel) LocalName() string {
	return mnc.Name
}

// Configuration returns the configuration (args) for [DataMonitoringNotificationChannel].
func (mnc *DataMonitoringNotificationChannel) Configuration() interface{} {
	return mnc.Args
}

// Attributes returns the attributes for [DataMonitoringNotificationChannel].
func (mnc *DataMonitoringNotificationChannel) Attributes() dataMonitoringNotificationChannelAttributes {
	return dataMonitoringNotificationChannelAttributes{ref: terra.ReferenceDataResource(mnc)}
}

// DataMonitoringNotificationChannelArgs contains the configurations for google_monitoring_notification_channel.
type DataMonitoringNotificationChannelArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// UserLabels: map of string, optional
	UserLabels terra.MapValue[terra.StringValue] `hcl:"user_labels,attr"`
	// SensitiveLabels: min=0
	SensitiveLabels []datamonitoringnotificationchannel.SensitiveLabels `hcl:"sensitive_labels,block" validate:"min=0"`
}
type dataMonitoringNotificationChannelAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_monitoring_notification_channel.
func (mnc dataMonitoringNotificationChannelAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mnc.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_monitoring_notification_channel.
func (mnc dataMonitoringNotificationChannelAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(mnc.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of google_monitoring_notification_channel.
func (mnc dataMonitoringNotificationChannelAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(mnc.ref.Append("enabled"))
}

// ForceDelete returns a reference to field force_delete of google_monitoring_notification_channel.
func (mnc dataMonitoringNotificationChannelAttributes) ForceDelete() terra.BoolValue {
	return terra.ReferenceAsBool(mnc.ref.Append("force_delete"))
}

// Id returns a reference to field id of google_monitoring_notification_channel.
func (mnc dataMonitoringNotificationChannelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mnc.ref.Append("id"))
}

// Labels returns a reference to field labels of google_monitoring_notification_channel.
func (mnc dataMonitoringNotificationChannelAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mnc.ref.Append("labels"))
}

// Name returns a reference to field name of google_monitoring_notification_channel.
func (mnc dataMonitoringNotificationChannelAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mnc.ref.Append("name"))
}

// Project returns a reference to field project of google_monitoring_notification_channel.
func (mnc dataMonitoringNotificationChannelAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(mnc.ref.Append("project"))
}

// Type returns a reference to field type of google_monitoring_notification_channel.
func (mnc dataMonitoringNotificationChannelAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(mnc.ref.Append("type"))
}

// UserLabels returns a reference to field user_labels of google_monitoring_notification_channel.
func (mnc dataMonitoringNotificationChannelAttributes) UserLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mnc.ref.Append("user_labels"))
}

// VerificationStatus returns a reference to field verification_status of google_monitoring_notification_channel.
func (mnc dataMonitoringNotificationChannelAttributes) VerificationStatus() terra.StringValue {
	return terra.ReferenceAsString(mnc.ref.Append("verification_status"))
}

func (mnc dataMonitoringNotificationChannelAttributes) SensitiveLabels() terra.ListValue[datamonitoringnotificationchannel.SensitiveLabelsAttributes] {
	return terra.ReferenceAsList[datamonitoringnotificationchannel.SensitiveLabelsAttributes](mnc.ref.Append("sensitive_labels"))
}
