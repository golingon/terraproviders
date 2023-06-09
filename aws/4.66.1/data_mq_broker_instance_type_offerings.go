// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datamqbrokerinstancetypeofferings "github.com/golingon/terraproviders/aws/4.66.1/datamqbrokerinstancetypeofferings"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMqBrokerInstanceTypeOfferings creates a new instance of [DataMqBrokerInstanceTypeOfferings].
func NewDataMqBrokerInstanceTypeOfferings(name string, args DataMqBrokerInstanceTypeOfferingsArgs) *DataMqBrokerInstanceTypeOfferings {
	return &DataMqBrokerInstanceTypeOfferings{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMqBrokerInstanceTypeOfferings)(nil)

// DataMqBrokerInstanceTypeOfferings represents the Terraform data resource aws_mq_broker_instance_type_offerings.
type DataMqBrokerInstanceTypeOfferings struct {
	Name string
	Args DataMqBrokerInstanceTypeOfferingsArgs
}

// DataSource returns the Terraform object type for [DataMqBrokerInstanceTypeOfferings].
func (mbito *DataMqBrokerInstanceTypeOfferings) DataSource() string {
	return "aws_mq_broker_instance_type_offerings"
}

// LocalName returns the local name for [DataMqBrokerInstanceTypeOfferings].
func (mbito *DataMqBrokerInstanceTypeOfferings) LocalName() string {
	return mbito.Name
}

// Configuration returns the configuration (args) for [DataMqBrokerInstanceTypeOfferings].
func (mbito *DataMqBrokerInstanceTypeOfferings) Configuration() interface{} {
	return mbito.Args
}

// Attributes returns the attributes for [DataMqBrokerInstanceTypeOfferings].
func (mbito *DataMqBrokerInstanceTypeOfferings) Attributes() dataMqBrokerInstanceTypeOfferingsAttributes {
	return dataMqBrokerInstanceTypeOfferingsAttributes{ref: terra.ReferenceDataResource(mbito)}
}

// DataMqBrokerInstanceTypeOfferingsArgs contains the configurations for aws_mq_broker_instance_type_offerings.
type DataMqBrokerInstanceTypeOfferingsArgs struct {
	// EngineType: string, optional
	EngineType terra.StringValue `hcl:"engine_type,attr"`
	// HostInstanceType: string, optional
	HostInstanceType terra.StringValue `hcl:"host_instance_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// StorageType: string, optional
	StorageType terra.StringValue `hcl:"storage_type,attr"`
	// BrokerInstanceOptions: min=0
	BrokerInstanceOptions []datamqbrokerinstancetypeofferings.BrokerInstanceOptions `hcl:"broker_instance_options,block" validate:"min=0"`
}
type dataMqBrokerInstanceTypeOfferingsAttributes struct {
	ref terra.Reference
}

// EngineType returns a reference to field engine_type of aws_mq_broker_instance_type_offerings.
func (mbito dataMqBrokerInstanceTypeOfferingsAttributes) EngineType() terra.StringValue {
	return terra.ReferenceAsString(mbito.ref.Append("engine_type"))
}

// HostInstanceType returns a reference to field host_instance_type of aws_mq_broker_instance_type_offerings.
func (mbito dataMqBrokerInstanceTypeOfferingsAttributes) HostInstanceType() terra.StringValue {
	return terra.ReferenceAsString(mbito.ref.Append("host_instance_type"))
}

// Id returns a reference to field id of aws_mq_broker_instance_type_offerings.
func (mbito dataMqBrokerInstanceTypeOfferingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mbito.ref.Append("id"))
}

// StorageType returns a reference to field storage_type of aws_mq_broker_instance_type_offerings.
func (mbito dataMqBrokerInstanceTypeOfferingsAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(mbito.ref.Append("storage_type"))
}

func (mbito dataMqBrokerInstanceTypeOfferingsAttributes) BrokerInstanceOptions() terra.ListValue[datamqbrokerinstancetypeofferings.BrokerInstanceOptionsAttributes] {
	return terra.ReferenceAsList[datamqbrokerinstancetypeofferings.BrokerInstanceOptionsAttributes](mbito.ref.Append("broker_instance_options"))
}
