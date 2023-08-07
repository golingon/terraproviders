// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	sccnotificationconfig "github.com/golingon/terraproviders/google/4.76.0/sccnotificationconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSccNotificationConfig creates a new instance of [SccNotificationConfig].
func NewSccNotificationConfig(name string, args SccNotificationConfigArgs) *SccNotificationConfig {
	return &SccNotificationConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SccNotificationConfig)(nil)

// SccNotificationConfig represents the Terraform resource google_scc_notification_config.
type SccNotificationConfig struct {
	Name      string
	Args      SccNotificationConfigArgs
	state     *sccNotificationConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SccNotificationConfig].
func (snc *SccNotificationConfig) Type() string {
	return "google_scc_notification_config"
}

// LocalName returns the local name for [SccNotificationConfig].
func (snc *SccNotificationConfig) LocalName() string {
	return snc.Name
}

// Configuration returns the configuration (args) for [SccNotificationConfig].
func (snc *SccNotificationConfig) Configuration() interface{} {
	return snc.Args
}

// DependOn is used for other resources to depend on [SccNotificationConfig].
func (snc *SccNotificationConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(snc)
}

// Dependencies returns the list of resources [SccNotificationConfig] depends_on.
func (snc *SccNotificationConfig) Dependencies() terra.Dependencies {
	return snc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SccNotificationConfig].
func (snc *SccNotificationConfig) LifecycleManagement() *terra.Lifecycle {
	return snc.Lifecycle
}

// Attributes returns the attributes for [SccNotificationConfig].
func (snc *SccNotificationConfig) Attributes() sccNotificationConfigAttributes {
	return sccNotificationConfigAttributes{ref: terra.ReferenceResource(snc)}
}

// ImportState imports the given attribute values into [SccNotificationConfig]'s state.
func (snc *SccNotificationConfig) ImportState(av io.Reader) error {
	snc.state = &sccNotificationConfigState{}
	if err := json.NewDecoder(av).Decode(snc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", snc.Type(), snc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SccNotificationConfig] has state.
func (snc *SccNotificationConfig) State() (*sccNotificationConfigState, bool) {
	return snc.state, snc.state != nil
}

// StateMust returns the state for [SccNotificationConfig]. Panics if the state is nil.
func (snc *SccNotificationConfig) StateMust() *sccNotificationConfigState {
	if snc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", snc.Type(), snc.LocalName()))
	}
	return snc.state
}

// SccNotificationConfigArgs contains the configurations for google_scc_notification_config.
type SccNotificationConfigArgs struct {
	// ConfigId: string, required
	ConfigId terra.StringValue `hcl:"config_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Organization: string, required
	Organization terra.StringValue `hcl:"organization,attr" validate:"required"`
	// PubsubTopic: string, required
	PubsubTopic terra.StringValue `hcl:"pubsub_topic,attr" validate:"required"`
	// StreamingConfig: required
	StreamingConfig *sccnotificationconfig.StreamingConfig `hcl:"streaming_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *sccnotificationconfig.Timeouts `hcl:"timeouts,block"`
}
type sccNotificationConfigAttributes struct {
	ref terra.Reference
}

// ConfigId returns a reference to field config_id of google_scc_notification_config.
func (snc sccNotificationConfigAttributes) ConfigId() terra.StringValue {
	return terra.ReferenceAsString(snc.ref.Append("config_id"))
}

// Description returns a reference to field description of google_scc_notification_config.
func (snc sccNotificationConfigAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(snc.ref.Append("description"))
}

// Id returns a reference to field id of google_scc_notification_config.
func (snc sccNotificationConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(snc.ref.Append("id"))
}

// Name returns a reference to field name of google_scc_notification_config.
func (snc sccNotificationConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(snc.ref.Append("name"))
}

// Organization returns a reference to field organization of google_scc_notification_config.
func (snc sccNotificationConfigAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(snc.ref.Append("organization"))
}

// PubsubTopic returns a reference to field pubsub_topic of google_scc_notification_config.
func (snc sccNotificationConfigAttributes) PubsubTopic() terra.StringValue {
	return terra.ReferenceAsString(snc.ref.Append("pubsub_topic"))
}

// ServiceAccount returns a reference to field service_account of google_scc_notification_config.
func (snc sccNotificationConfigAttributes) ServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(snc.ref.Append("service_account"))
}

func (snc sccNotificationConfigAttributes) StreamingConfig() terra.ListValue[sccnotificationconfig.StreamingConfigAttributes] {
	return terra.ReferenceAsList[sccnotificationconfig.StreamingConfigAttributes](snc.ref.Append("streaming_config"))
}

func (snc sccNotificationConfigAttributes) Timeouts() sccnotificationconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sccnotificationconfig.TimeoutsAttributes](snc.ref.Append("timeouts"))
}

type sccNotificationConfigState struct {
	ConfigId        string                                       `json:"config_id"`
	Description     string                                       `json:"description"`
	Id              string                                       `json:"id"`
	Name            string                                       `json:"name"`
	Organization    string                                       `json:"organization"`
	PubsubTopic     string                                       `json:"pubsub_topic"`
	ServiceAccount  string                                       `json:"service_account"`
	StreamingConfig []sccnotificationconfig.StreamingConfigState `json:"streaming_config"`
	Timeouts        *sccnotificationconfig.TimeoutsState         `json:"timeouts"`
}
