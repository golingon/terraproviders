// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataMskKafkaVersion creates a new instance of [DataMskKafkaVersion].
func NewDataMskKafkaVersion(name string, args DataMskKafkaVersionArgs) *DataMskKafkaVersion {
	return &DataMskKafkaVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMskKafkaVersion)(nil)

// DataMskKafkaVersion represents the Terraform data resource aws_msk_kafka_version.
type DataMskKafkaVersion struct {
	Name string
	Args DataMskKafkaVersionArgs
}

// DataSource returns the Terraform object type for [DataMskKafkaVersion].
func (mkv *DataMskKafkaVersion) DataSource() string {
	return "aws_msk_kafka_version"
}

// LocalName returns the local name for [DataMskKafkaVersion].
func (mkv *DataMskKafkaVersion) LocalName() string {
	return mkv.Name
}

// Configuration returns the configuration (args) for [DataMskKafkaVersion].
func (mkv *DataMskKafkaVersion) Configuration() interface{} {
	return mkv.Args
}

// Attributes returns the attributes for [DataMskKafkaVersion].
func (mkv *DataMskKafkaVersion) Attributes() dataMskKafkaVersionAttributes {
	return dataMskKafkaVersionAttributes{ref: terra.ReferenceDataResource(mkv)}
}

// DataMskKafkaVersionArgs contains the configurations for aws_msk_kafka_version.
type DataMskKafkaVersionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PreferredVersions: list of string, optional
	PreferredVersions terra.ListValue[terra.StringValue] `hcl:"preferred_versions,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
}
type dataMskKafkaVersionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_msk_kafka_version.
func (mkv dataMskKafkaVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mkv.ref.Append("id"))
}

// PreferredVersions returns a reference to field preferred_versions of aws_msk_kafka_version.
func (mkv dataMskKafkaVersionAttributes) PreferredVersions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mkv.ref.Append("preferred_versions"))
}

// Status returns a reference to field status of aws_msk_kafka_version.
func (mkv dataMskKafkaVersionAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(mkv.ref.Append("status"))
}

// Version returns a reference to field version of aws_msk_kafka_version.
func (mkv dataMskKafkaVersionAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(mkv.ref.Append("version"))
}