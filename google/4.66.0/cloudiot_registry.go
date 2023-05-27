// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudiotregistry "github.com/golingon/terraproviders/google/4.66.0/cloudiotregistry"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudiotRegistry creates a new instance of [CloudiotRegistry].
func NewCloudiotRegistry(name string, args CloudiotRegistryArgs) *CloudiotRegistry {
	return &CloudiotRegistry{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudiotRegistry)(nil)

// CloudiotRegistry represents the Terraform resource google_cloudiot_registry.
type CloudiotRegistry struct {
	Name      string
	Args      CloudiotRegistryArgs
	state     *cloudiotRegistryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudiotRegistry].
func (cr *CloudiotRegistry) Type() string {
	return "google_cloudiot_registry"
}

// LocalName returns the local name for [CloudiotRegistry].
func (cr *CloudiotRegistry) LocalName() string {
	return cr.Name
}

// Configuration returns the configuration (args) for [CloudiotRegistry].
func (cr *CloudiotRegistry) Configuration() interface{} {
	return cr.Args
}

// DependOn is used for other resources to depend on [CloudiotRegistry].
func (cr *CloudiotRegistry) DependOn() terra.Reference {
	return terra.ReferenceResource(cr)
}

// Dependencies returns the list of resources [CloudiotRegistry] depends_on.
func (cr *CloudiotRegistry) Dependencies() terra.Dependencies {
	return cr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudiotRegistry].
func (cr *CloudiotRegistry) LifecycleManagement() *terra.Lifecycle {
	return cr.Lifecycle
}

// Attributes returns the attributes for [CloudiotRegistry].
func (cr *CloudiotRegistry) Attributes() cloudiotRegistryAttributes {
	return cloudiotRegistryAttributes{ref: terra.ReferenceResource(cr)}
}

// ImportState imports the given attribute values into [CloudiotRegistry]'s state.
func (cr *CloudiotRegistry) ImportState(av io.Reader) error {
	cr.state = &cloudiotRegistryState{}
	if err := json.NewDecoder(av).Decode(cr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cr.Type(), cr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudiotRegistry] has state.
func (cr *CloudiotRegistry) State() (*cloudiotRegistryState, bool) {
	return cr.state, cr.state != nil
}

// StateMust returns the state for [CloudiotRegistry]. Panics if the state is nil.
func (cr *CloudiotRegistry) StateMust() *cloudiotRegistryState {
	if cr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cr.Type(), cr.LocalName()))
	}
	return cr.state
}

// CloudiotRegistryArgs contains the configurations for google_cloudiot_registry.
type CloudiotRegistryArgs struct {
	// HttpConfig: map of string, optional
	HttpConfig terra.MapValue[terra.StringValue] `hcl:"http_config,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogLevel: string, optional
	LogLevel terra.StringValue `hcl:"log_level,attr"`
	// MqttConfig: map of string, optional
	MqttConfig terra.MapValue[terra.StringValue] `hcl:"mqtt_config,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// StateNotificationConfig: map of string, optional
	StateNotificationConfig terra.MapValue[terra.StringValue] `hcl:"state_notification_config,attr"`
	// Credentials: min=0,max=10
	Credentials []cloudiotregistry.Credentials `hcl:"credentials,block" validate:"min=0,max=10"`
	// EventNotificationConfigs: min=0,max=10
	EventNotificationConfigs []cloudiotregistry.EventNotificationConfigs `hcl:"event_notification_configs,block" validate:"min=0,max=10"`
	// Timeouts: optional
	Timeouts *cloudiotregistry.Timeouts `hcl:"timeouts,block"`
}
type cloudiotRegistryAttributes struct {
	ref terra.Reference
}

// HttpConfig returns a reference to field http_config of google_cloudiot_registry.
func (cr cloudiotRegistryAttributes) HttpConfig() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cr.ref.Append("http_config"))
}

// Id returns a reference to field id of google_cloudiot_registry.
func (cr cloudiotRegistryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("id"))
}

// LogLevel returns a reference to field log_level of google_cloudiot_registry.
func (cr cloudiotRegistryAttributes) LogLevel() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("log_level"))
}

// MqttConfig returns a reference to field mqtt_config of google_cloudiot_registry.
func (cr cloudiotRegistryAttributes) MqttConfig() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cr.ref.Append("mqtt_config"))
}

// Name returns a reference to field name of google_cloudiot_registry.
func (cr cloudiotRegistryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("name"))
}

// Project returns a reference to field project of google_cloudiot_registry.
func (cr cloudiotRegistryAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("project"))
}

// Region returns a reference to field region of google_cloudiot_registry.
func (cr cloudiotRegistryAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("region"))
}

// StateNotificationConfig returns a reference to field state_notification_config of google_cloudiot_registry.
func (cr cloudiotRegistryAttributes) StateNotificationConfig() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cr.ref.Append("state_notification_config"))
}

func (cr cloudiotRegistryAttributes) Credentials() terra.ListValue[cloudiotregistry.CredentialsAttributes] {
	return terra.ReferenceAsList[cloudiotregistry.CredentialsAttributes](cr.ref.Append("credentials"))
}

func (cr cloudiotRegistryAttributes) EventNotificationConfigs() terra.ListValue[cloudiotregistry.EventNotificationConfigsAttributes] {
	return terra.ReferenceAsList[cloudiotregistry.EventNotificationConfigsAttributes](cr.ref.Append("event_notification_configs"))
}

func (cr cloudiotRegistryAttributes) Timeouts() cloudiotregistry.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudiotregistry.TimeoutsAttributes](cr.ref.Append("timeouts"))
}

type cloudiotRegistryState struct {
	HttpConfig               map[string]string                                `json:"http_config"`
	Id                       string                                           `json:"id"`
	LogLevel                 string                                           `json:"log_level"`
	MqttConfig               map[string]string                                `json:"mqtt_config"`
	Name                     string                                           `json:"name"`
	Project                  string                                           `json:"project"`
	Region                   string                                           `json:"region"`
	StateNotificationConfig  map[string]string                                `json:"state_notification_config"`
	Credentials              []cloudiotregistry.CredentialsState              `json:"credentials"`
	EventNotificationConfigs []cloudiotregistry.EventNotificationConfigsState `json:"event_notification_configs"`
	Timeouts                 *cloudiotregistry.TimeoutsState                  `json:"timeouts"`
}
