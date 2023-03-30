// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudfrontrealtimelogconfig "github.com/golingon/terraproviders/aws/4.60.0/cloudfrontrealtimelogconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewCloudfrontRealtimeLogConfig(name string, args CloudfrontRealtimeLogConfigArgs) *CloudfrontRealtimeLogConfig {
	return &CloudfrontRealtimeLogConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudfrontRealtimeLogConfig)(nil)

type CloudfrontRealtimeLogConfig struct {
	Name  string
	Args  CloudfrontRealtimeLogConfigArgs
	state *cloudfrontRealtimeLogConfigState
}

func (crlc *CloudfrontRealtimeLogConfig) Type() string {
	return "aws_cloudfront_realtime_log_config"
}

func (crlc *CloudfrontRealtimeLogConfig) LocalName() string {
	return crlc.Name
}

func (crlc *CloudfrontRealtimeLogConfig) Configuration() interface{} {
	return crlc.Args
}

func (crlc *CloudfrontRealtimeLogConfig) Attributes() cloudfrontRealtimeLogConfigAttributes {
	return cloudfrontRealtimeLogConfigAttributes{ref: terra.ReferenceResource(crlc)}
}

func (crlc *CloudfrontRealtimeLogConfig) ImportState(av io.Reader) error {
	crlc.state = &cloudfrontRealtimeLogConfigState{}
	if err := json.NewDecoder(av).Decode(crlc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crlc.Type(), crlc.LocalName(), err)
	}
	return nil
}

func (crlc *CloudfrontRealtimeLogConfig) State() (*cloudfrontRealtimeLogConfigState, bool) {
	return crlc.state, crlc.state != nil
}

func (crlc *CloudfrontRealtimeLogConfig) StateMust() *cloudfrontRealtimeLogConfigState {
	if crlc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crlc.Type(), crlc.LocalName()))
	}
	return crlc.state
}

func (crlc *CloudfrontRealtimeLogConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(crlc)
}

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
	// DependsOn contains resources that CloudfrontRealtimeLogConfig depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type cloudfrontRealtimeLogConfigAttributes struct {
	ref terra.Reference
}

func (crlc cloudfrontRealtimeLogConfigAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(crlc.ref.Append("arn"))
}

func (crlc cloudfrontRealtimeLogConfigAttributes) Fields() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](crlc.ref.Append("fields"))
}

func (crlc cloudfrontRealtimeLogConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceString(crlc.ref.Append("id"))
}

func (crlc cloudfrontRealtimeLogConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceString(crlc.ref.Append("name"))
}

func (crlc cloudfrontRealtimeLogConfigAttributes) SamplingRate() terra.NumberValue {
	return terra.ReferenceNumber(crlc.ref.Append("sampling_rate"))
}

func (crlc cloudfrontRealtimeLogConfigAttributes) Endpoint() terra.ListValue[cloudfrontrealtimelogconfig.EndpointAttributes] {
	return terra.ReferenceList[cloudfrontrealtimelogconfig.EndpointAttributes](crlc.ref.Append("endpoint"))
}

type cloudfrontRealtimeLogConfigState struct {
	Arn          string                                      `json:"arn"`
	Fields       []string                                    `json:"fields"`
	Id           string                                      `json:"id"`
	Name         string                                      `json:"name"`
	SamplingRate float64                                     `json:"sampling_rate"`
	Endpoint     []cloudfrontrealtimelogconfig.EndpointState `json:"endpoint"`
}
