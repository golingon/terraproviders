// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datapubsubtopic "github.com/golingon/terraproviders/googlebeta/4.66.0/datapubsubtopic"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPubsubTopic creates a new instance of [DataPubsubTopic].
func NewDataPubsubTopic(name string, args DataPubsubTopicArgs) *DataPubsubTopic {
	return &DataPubsubTopic{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPubsubTopic)(nil)

// DataPubsubTopic represents the Terraform data resource google_pubsub_topic.
type DataPubsubTopic struct {
	Name string
	Args DataPubsubTopicArgs
}

// DataSource returns the Terraform object type for [DataPubsubTopic].
func (pt *DataPubsubTopic) DataSource() string {
	return "google_pubsub_topic"
}

// LocalName returns the local name for [DataPubsubTopic].
func (pt *DataPubsubTopic) LocalName() string {
	return pt.Name
}

// Configuration returns the configuration (args) for [DataPubsubTopic].
func (pt *DataPubsubTopic) Configuration() interface{} {
	return pt.Args
}

// Attributes returns the attributes for [DataPubsubTopic].
func (pt *DataPubsubTopic) Attributes() dataPubsubTopicAttributes {
	return dataPubsubTopicAttributes{ref: terra.ReferenceDataResource(pt)}
}

// DataPubsubTopicArgs contains the configurations for google_pubsub_topic.
type DataPubsubTopicArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// MessageStoragePolicy: min=0
	MessageStoragePolicy []datapubsubtopic.MessageStoragePolicy `hcl:"message_storage_policy,block" validate:"min=0"`
	// SchemaSettings: min=0
	SchemaSettings []datapubsubtopic.SchemaSettings `hcl:"schema_settings,block" validate:"min=0"`
}
type dataPubsubTopicAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_pubsub_topic.
func (pt dataPubsubTopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("id"))
}

// KmsKeyName returns a reference to field kms_key_name of google_pubsub_topic.
func (pt dataPubsubTopicAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("kms_key_name"))
}

// Labels returns a reference to field labels of google_pubsub_topic.
func (pt dataPubsubTopicAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pt.ref.Append("labels"))
}

// MessageRetentionDuration returns a reference to field message_retention_duration of google_pubsub_topic.
func (pt dataPubsubTopicAttributes) MessageRetentionDuration() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("message_retention_duration"))
}

// Name returns a reference to field name of google_pubsub_topic.
func (pt dataPubsubTopicAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("name"))
}

// Project returns a reference to field project of google_pubsub_topic.
func (pt dataPubsubTopicAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("project"))
}

func (pt dataPubsubTopicAttributes) MessageStoragePolicy() terra.ListValue[datapubsubtopic.MessageStoragePolicyAttributes] {
	return terra.ReferenceAsList[datapubsubtopic.MessageStoragePolicyAttributes](pt.ref.Append("message_storage_policy"))
}

func (pt dataPubsubTopicAttributes) SchemaSettings() terra.ListValue[datapubsubtopic.SchemaSettingsAttributes] {
	return terra.ReferenceAsList[datapubsubtopic.SchemaSettingsAttributes](pt.ref.Append("schema_settings"))
}
