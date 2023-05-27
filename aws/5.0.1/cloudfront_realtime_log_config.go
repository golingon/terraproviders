// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudfrontrealtimelogconfig "github.com/golingon/terraproviders/aws/5.0.1/cloudfrontrealtimelogconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudfrontRealtimeLogConfig creates a new instance of [CloudfrontRealtimeLogConfig].
func NewCloudfrontRealtimeLogConfig(name string, args CloudfrontRealtimeLogConfigArgs) *CloudfrontRealtimeLogConfig {
	return &CloudfrontRealtimeLogConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudfrontRealtimeLogConfig)(nil)

// CloudfrontRealtimeLogConfig represents the Terraform resource aws_cloudfront_realtime_log_config.
type CloudfrontRealtimeLogConfig struct {
	Name      string
	Args      CloudfrontRealtimeLogConfigArgs
	state     *cloudfrontRealtimeLogConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudfrontRealtimeLogConfig].
func (crlc *CloudfrontRealtimeLogConfig) Type() string {
	return "aws_cloudfront_realtime_log_config"
}

// LocalName returns the local name for [CloudfrontRealtimeLogConfig].
func (crlc *CloudfrontRealtimeLogConfig) LocalName() string {
	return crlc.Name
}

// Configuration returns the configuration (args) for [CloudfrontRealtimeLogConfig].
func (crlc *CloudfrontRealtimeLogConfig) Configuration() interface{} {
	return crlc.Args
}

// DependOn is used for other resources to depend on [CloudfrontRealtimeLogConfig].
func (crlc *CloudfrontRealtimeLogConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(crlc)
}

// Dependencies returns the list of resources [CloudfrontRealtimeLogConfig] depends_on.
func (crlc *CloudfrontRealtimeLogConfig) Dependencies() terra.Dependencies {
	return crlc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudfrontRealtimeLogConfig].
func (crlc *CloudfrontRealtimeLogConfig) LifecycleManagement() *terra.Lifecycle {
	return crlc.Lifecycle
}

// Attributes returns the attributes for [CloudfrontRealtimeLogConfig].
func (crlc *CloudfrontRealtimeLogConfig) Attributes() cloudfrontRealtimeLogConfigAttributes {
	return cloudfrontRealtimeLogConfigAttributes{ref: terra.ReferenceResource(crlc)}
}

// ImportState imports the given attribute values into [CloudfrontRealtimeLogConfig]'s state.
func (crlc *CloudfrontRealtimeLogConfig) ImportState(av io.Reader) error {
	crlc.state = &cloudfrontRealtimeLogConfigState{}
	if err := json.NewDecoder(av).Decode(crlc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crlc.Type(), crlc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudfrontRealtimeLogConfig] has state.
func (crlc *CloudfrontRealtimeLogConfig) State() (*cloudfrontRealtimeLogConfigState, bool) {
	return crlc.state, crlc.state != nil
}

// StateMust returns the state for [CloudfrontRealtimeLogConfig]. Panics if the state is nil.
func (crlc *CloudfrontRealtimeLogConfig) StateMust() *cloudfrontRealtimeLogConfigState {
	if crlc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crlc.Type(), crlc.LocalName()))
	}
	return crlc.state
}

// CloudfrontRealtimeLogConfigArgs contains the configurations for aws_cloudfront_realtime_log_config.
type CloudfrontRealtimeLogConfigArgs struct {
	// Fields: set of string, required
	Fields terra.SetValue[terra.StringValue] `hcl:"fields,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SamplingRate: number, required
	SamplingRate terra.NumberValue `hcl:"sampling_rate,attr" validate:"required"`
	// Endpoint: required
	Endpoint *cloudfrontrealtimelogconfig.Endpoint `hcl:"endpoint,block" validate:"required"`
}
type cloudfrontRealtimeLogConfigAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudfront_realtime_log_config.
func (crlc cloudfrontRealtimeLogConfigAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(crlc.ref.Append("arn"))
}

// Fields returns a reference to field fields of aws_cloudfront_realtime_log_config.
func (crlc cloudfrontRealtimeLogConfigAttributes) Fields() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crlc.ref.Append("fields"))
}

// Id returns a reference to field id of aws_cloudfront_realtime_log_config.
func (crlc cloudfrontRealtimeLogConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crlc.ref.Append("id"))
}

// Name returns a reference to field name of aws_cloudfront_realtime_log_config.
func (crlc cloudfrontRealtimeLogConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crlc.ref.Append("name"))
}

// SamplingRate returns a reference to field sampling_rate of aws_cloudfront_realtime_log_config.
func (crlc cloudfrontRealtimeLogConfigAttributes) SamplingRate() terra.NumberValue {
	return terra.ReferenceAsNumber(crlc.ref.Append("sampling_rate"))
}

func (crlc cloudfrontRealtimeLogConfigAttributes) Endpoint() terra.ListValue[cloudfrontrealtimelogconfig.EndpointAttributes] {
	return terra.ReferenceAsList[cloudfrontrealtimelogconfig.EndpointAttributes](crlc.ref.Append("endpoint"))
}

type cloudfrontRealtimeLogConfigState struct {
	Arn          string                                      `json:"arn"`
	Fields       []string                                    `json:"fields"`
	Id           string                                      `json:"id"`
	Name         string                                      `json:"name"`
	SamplingRate float64                                     `json:"sampling_rate"`
	Endpoint     []cloudfrontrealtimelogconfig.EndpointState `json:"endpoint"`
}
