// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datacloudfrontrealtimelogconfig "github.com/golingon/terraproviders/aws/5.0.1/datacloudfrontrealtimelogconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCloudfrontRealtimeLogConfig creates a new instance of [DataCloudfrontRealtimeLogConfig].
func NewDataCloudfrontRealtimeLogConfig(name string, args DataCloudfrontRealtimeLogConfigArgs) *DataCloudfrontRealtimeLogConfig {
	return &DataCloudfrontRealtimeLogConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudfrontRealtimeLogConfig)(nil)

// DataCloudfrontRealtimeLogConfig represents the Terraform data resource aws_cloudfront_realtime_log_config.
type DataCloudfrontRealtimeLogConfig struct {
	Name string
	Args DataCloudfrontRealtimeLogConfigArgs
}

// DataSource returns the Terraform object type for [DataCloudfrontRealtimeLogConfig].
func (crlc *DataCloudfrontRealtimeLogConfig) DataSource() string {
	return "aws_cloudfront_realtime_log_config"
}

// LocalName returns the local name for [DataCloudfrontRealtimeLogConfig].
func (crlc *DataCloudfrontRealtimeLogConfig) LocalName() string {
	return crlc.Name
}

// Configuration returns the configuration (args) for [DataCloudfrontRealtimeLogConfig].
func (crlc *DataCloudfrontRealtimeLogConfig) Configuration() interface{} {
	return crlc.Args
}

// Attributes returns the attributes for [DataCloudfrontRealtimeLogConfig].
func (crlc *DataCloudfrontRealtimeLogConfig) Attributes() dataCloudfrontRealtimeLogConfigAttributes {
	return dataCloudfrontRealtimeLogConfigAttributes{ref: terra.ReferenceDataResource(crlc)}
}

// DataCloudfrontRealtimeLogConfigArgs contains the configurations for aws_cloudfront_realtime_log_config.
type DataCloudfrontRealtimeLogConfigArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Endpoint: min=0
	Endpoint []datacloudfrontrealtimelogconfig.Endpoint `hcl:"endpoint,block" validate:"min=0"`
}
type dataCloudfrontRealtimeLogConfigAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudfront_realtime_log_config.
func (crlc dataCloudfrontRealtimeLogConfigAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(crlc.ref.Append("arn"))
}

// Fields returns a reference to field fields of aws_cloudfront_realtime_log_config.
func (crlc dataCloudfrontRealtimeLogConfigAttributes) Fields() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crlc.ref.Append("fields"))
}

// Id returns a reference to field id of aws_cloudfront_realtime_log_config.
func (crlc dataCloudfrontRealtimeLogConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crlc.ref.Append("id"))
}

// Name returns a reference to field name of aws_cloudfront_realtime_log_config.
func (crlc dataCloudfrontRealtimeLogConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crlc.ref.Append("name"))
}

// SamplingRate returns a reference to field sampling_rate of aws_cloudfront_realtime_log_config.
func (crlc dataCloudfrontRealtimeLogConfigAttributes) SamplingRate() terra.NumberValue {
	return terra.ReferenceAsNumber(crlc.ref.Append("sampling_rate"))
}

func (crlc dataCloudfrontRealtimeLogConfigAttributes) Endpoint() terra.ListValue[datacloudfrontrealtimelogconfig.EndpointAttributes] {
	return terra.ReferenceAsList[datacloudfrontrealtimelogconfig.EndpointAttributes](crlc.ref.Append("endpoint"))
}
