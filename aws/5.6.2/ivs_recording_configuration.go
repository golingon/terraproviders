// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ivsrecordingconfiguration "github.com/golingon/terraproviders/aws/5.6.2/ivsrecordingconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIvsRecordingConfiguration creates a new instance of [IvsRecordingConfiguration].
func NewIvsRecordingConfiguration(name string, args IvsRecordingConfigurationArgs) *IvsRecordingConfiguration {
	return &IvsRecordingConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IvsRecordingConfiguration)(nil)

// IvsRecordingConfiguration represents the Terraform resource aws_ivs_recording_configuration.
type IvsRecordingConfiguration struct {
	Name      string
	Args      IvsRecordingConfigurationArgs
	state     *ivsRecordingConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IvsRecordingConfiguration].
func (irc *IvsRecordingConfiguration) Type() string {
	return "aws_ivs_recording_configuration"
}

// LocalName returns the local name for [IvsRecordingConfiguration].
func (irc *IvsRecordingConfiguration) LocalName() string {
	return irc.Name
}

// Configuration returns the configuration (args) for [IvsRecordingConfiguration].
func (irc *IvsRecordingConfiguration) Configuration() interface{} {
	return irc.Args
}

// DependOn is used for other resources to depend on [IvsRecordingConfiguration].
func (irc *IvsRecordingConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(irc)
}

// Dependencies returns the list of resources [IvsRecordingConfiguration] depends_on.
func (irc *IvsRecordingConfiguration) Dependencies() terra.Dependencies {
	return irc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IvsRecordingConfiguration].
func (irc *IvsRecordingConfiguration) LifecycleManagement() *terra.Lifecycle {
	return irc.Lifecycle
}

// Attributes returns the attributes for [IvsRecordingConfiguration].
func (irc *IvsRecordingConfiguration) Attributes() ivsRecordingConfigurationAttributes {
	return ivsRecordingConfigurationAttributes{ref: terra.ReferenceResource(irc)}
}

// ImportState imports the given attribute values into [IvsRecordingConfiguration]'s state.
func (irc *IvsRecordingConfiguration) ImportState(av io.Reader) error {
	irc.state = &ivsRecordingConfigurationState{}
	if err := json.NewDecoder(av).Decode(irc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", irc.Type(), irc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IvsRecordingConfiguration] has state.
func (irc *IvsRecordingConfiguration) State() (*ivsRecordingConfigurationState, bool) {
	return irc.state, irc.state != nil
}

// StateMust returns the state for [IvsRecordingConfiguration]. Panics if the state is nil.
func (irc *IvsRecordingConfiguration) StateMust() *ivsRecordingConfigurationState {
	if irc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", irc.Type(), irc.LocalName()))
	}
	return irc.state
}

// IvsRecordingConfigurationArgs contains the configurations for aws_ivs_recording_configuration.
type IvsRecordingConfigurationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// RecordingReconnectWindowSeconds: number, optional
	RecordingReconnectWindowSeconds terra.NumberValue `hcl:"recording_reconnect_window_seconds,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DestinationConfiguration: required
	DestinationConfiguration *ivsrecordingconfiguration.DestinationConfiguration `hcl:"destination_configuration,block" validate:"required"`
	// ThumbnailConfiguration: optional
	ThumbnailConfiguration *ivsrecordingconfiguration.ThumbnailConfiguration `hcl:"thumbnail_configuration,block"`
	// Timeouts: optional
	Timeouts *ivsrecordingconfiguration.Timeouts `hcl:"timeouts,block"`
}
type ivsRecordingConfigurationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ivs_recording_configuration.
func (irc ivsRecordingConfigurationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(irc.ref.Append("arn"))
}

// Id returns a reference to field id of aws_ivs_recording_configuration.
func (irc ivsRecordingConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(irc.ref.Append("id"))
}

// Name returns a reference to field name of aws_ivs_recording_configuration.
func (irc ivsRecordingConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(irc.ref.Append("name"))
}

// RecordingReconnectWindowSeconds returns a reference to field recording_reconnect_window_seconds of aws_ivs_recording_configuration.
func (irc ivsRecordingConfigurationAttributes) RecordingReconnectWindowSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(irc.ref.Append("recording_reconnect_window_seconds"))
}

// State returns a reference to field state of aws_ivs_recording_configuration.
func (irc ivsRecordingConfigurationAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(irc.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_ivs_recording_configuration.
func (irc ivsRecordingConfigurationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](irc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ivs_recording_configuration.
func (irc ivsRecordingConfigurationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](irc.ref.Append("tags_all"))
}

func (irc ivsRecordingConfigurationAttributes) DestinationConfiguration() terra.ListValue[ivsrecordingconfiguration.DestinationConfigurationAttributes] {
	return terra.ReferenceAsList[ivsrecordingconfiguration.DestinationConfigurationAttributes](irc.ref.Append("destination_configuration"))
}

func (irc ivsRecordingConfigurationAttributes) ThumbnailConfiguration() terra.ListValue[ivsrecordingconfiguration.ThumbnailConfigurationAttributes] {
	return terra.ReferenceAsList[ivsrecordingconfiguration.ThumbnailConfigurationAttributes](irc.ref.Append("thumbnail_configuration"))
}

func (irc ivsRecordingConfigurationAttributes) Timeouts() ivsrecordingconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ivsrecordingconfiguration.TimeoutsAttributes](irc.ref.Append("timeouts"))
}

type ivsRecordingConfigurationState struct {
	Arn                             string                                                    `json:"arn"`
	Id                              string                                                    `json:"id"`
	Name                            string                                                    `json:"name"`
	RecordingReconnectWindowSeconds float64                                                   `json:"recording_reconnect_window_seconds"`
	State                           string                                                    `json:"state"`
	Tags                            map[string]string                                         `json:"tags"`
	TagsAll                         map[string]string                                         `json:"tags_all"`
	DestinationConfiguration        []ivsrecordingconfiguration.DestinationConfigurationState `json:"destination_configuration"`
	ThumbnailConfiguration          []ivsrecordingconfiguration.ThumbnailConfigurationState   `json:"thumbnail_configuration"`
	Timeouts                        *ivsrecordingconfiguration.TimeoutsState                  `json:"timeouts"`
}
