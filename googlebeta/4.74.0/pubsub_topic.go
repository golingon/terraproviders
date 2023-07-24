// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	pubsubtopic "github.com/golingon/terraproviders/googlebeta/4.74.0/pubsubtopic"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPubsubTopic creates a new instance of [PubsubTopic].
func NewPubsubTopic(name string, args PubsubTopicArgs) *PubsubTopic {
	return &PubsubTopic{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PubsubTopic)(nil)

// PubsubTopic represents the Terraform resource google_pubsub_topic.
type PubsubTopic struct {
	Name      string
	Args      PubsubTopicArgs
	state     *pubsubTopicState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PubsubTopic].
func (pt *PubsubTopic) Type() string {
	return "google_pubsub_topic"
}

// LocalName returns the local name for [PubsubTopic].
func (pt *PubsubTopic) LocalName() string {
	return pt.Name
}

// Configuration returns the configuration (args) for [PubsubTopic].
func (pt *PubsubTopic) Configuration() interface{} {
	return pt.Args
}

// DependOn is used for other resources to depend on [PubsubTopic].
func (pt *PubsubTopic) DependOn() terra.Reference {
	return terra.ReferenceResource(pt)
}

// Dependencies returns the list of resources [PubsubTopic] depends_on.
func (pt *PubsubTopic) Dependencies() terra.Dependencies {
	return pt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PubsubTopic].
func (pt *PubsubTopic) LifecycleManagement() *terra.Lifecycle {
	return pt.Lifecycle
}

// Attributes returns the attributes for [PubsubTopic].
func (pt *PubsubTopic) Attributes() pubsubTopicAttributes {
	return pubsubTopicAttributes{ref: terra.ReferenceResource(pt)}
}

// ImportState imports the given attribute values into [PubsubTopic]'s state.
func (pt *PubsubTopic) ImportState(av io.Reader) error {
	pt.state = &pubsubTopicState{}
	if err := json.NewDecoder(av).Decode(pt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pt.Type(), pt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PubsubTopic] has state.
func (pt *PubsubTopic) State() (*pubsubTopicState, bool) {
	return pt.state, pt.state != nil
}

// StateMust returns the state for [PubsubTopic]. Panics if the state is nil.
func (pt *PubsubTopic) StateMust() *pubsubTopicState {
	if pt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pt.Type(), pt.LocalName()))
	}
	return pt.state
}

// PubsubTopicArgs contains the configurations for google_pubsub_topic.
type PubsubTopicArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyName: string, optional
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// MessageRetentionDuration: string, optional
	MessageRetentionDuration terra.StringValue `hcl:"message_retention_duration,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// MessageStoragePolicy: optional
	MessageStoragePolicy *pubsubtopic.MessageStoragePolicy `hcl:"message_storage_policy,block"`
	// SchemaSettings: optional
	SchemaSettings *pubsubtopic.SchemaSettings `hcl:"schema_settings,block"`
	// Timeouts: optional
	Timeouts *pubsubtopic.Timeouts `hcl:"timeouts,block"`
}
type pubsubTopicAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_pubsub_topic.
func (pt pubsubTopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("id"))
}

// KmsKeyName returns a reference to field kms_key_name of google_pubsub_topic.
func (pt pubsubTopicAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("kms_key_name"))
}

// Labels returns a reference to field labels of google_pubsub_topic.
func (pt pubsubTopicAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pt.ref.Append("labels"))
}

// MessageRetentionDuration returns a reference to field message_retention_duration of google_pubsub_topic.
func (pt pubsubTopicAttributes) MessageRetentionDuration() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("message_retention_duration"))
}

// Name returns a reference to field name of google_pubsub_topic.
func (pt pubsubTopicAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("name"))
}

// Project returns a reference to field project of google_pubsub_topic.
func (pt pubsubTopicAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("project"))
}

func (pt pubsubTopicAttributes) MessageStoragePolicy() terra.ListValue[pubsubtopic.MessageStoragePolicyAttributes] {
	return terra.ReferenceAsList[pubsubtopic.MessageStoragePolicyAttributes](pt.ref.Append("message_storage_policy"))
}

func (pt pubsubTopicAttributes) SchemaSettings() terra.ListValue[pubsubtopic.SchemaSettingsAttributes] {
	return terra.ReferenceAsList[pubsubtopic.SchemaSettingsAttributes](pt.ref.Append("schema_settings"))
}

func (pt pubsubTopicAttributes) Timeouts() pubsubtopic.TimeoutsAttributes {
	return terra.ReferenceAsSingle[pubsubtopic.TimeoutsAttributes](pt.ref.Append("timeouts"))
}

type pubsubTopicState struct {
	Id                       string                                  `json:"id"`
	KmsKeyName               string                                  `json:"kms_key_name"`
	Labels                   map[string]string                       `json:"labels"`
	MessageRetentionDuration string                                  `json:"message_retention_duration"`
	Name                     string                                  `json:"name"`
	Project                  string                                  `json:"project"`
	MessageStoragePolicy     []pubsubtopic.MessageStoragePolicyState `json:"message_storage_policy"`
	SchemaSettings           []pubsubtopic.SchemaSettingsState       `json:"schema_settings"`
	Timeouts                 *pubsubtopic.TimeoutsState              `json:"timeouts"`
}
