// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewLightsailLbStickinessPolicy(name string, args LightsailLbStickinessPolicyArgs) *LightsailLbStickinessPolicy {
	return &LightsailLbStickinessPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LightsailLbStickinessPolicy)(nil)

type LightsailLbStickinessPolicy struct {
	Name  string
	Args  LightsailLbStickinessPolicyArgs
	state *lightsailLbStickinessPolicyState
}

func (llsp *LightsailLbStickinessPolicy) Type() string {
	return "aws_lightsail_lb_stickiness_policy"
}

func (llsp *LightsailLbStickinessPolicy) LocalName() string {
	return llsp.Name
}

func (llsp *LightsailLbStickinessPolicy) Configuration() interface{} {
	return llsp.Args
}

func (llsp *LightsailLbStickinessPolicy) Attributes() lightsailLbStickinessPolicyAttributes {
	return lightsailLbStickinessPolicyAttributes{ref: terra.ReferenceResource(llsp)}
}

func (llsp *LightsailLbStickinessPolicy) ImportState(av io.Reader) error {
	llsp.state = &lightsailLbStickinessPolicyState{}
	if err := json.NewDecoder(av).Decode(llsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", llsp.Type(), llsp.LocalName(), err)
	}
	return nil
}

func (llsp *LightsailLbStickinessPolicy) State() (*lightsailLbStickinessPolicyState, bool) {
	return llsp.state, llsp.state != nil
}

func (llsp *LightsailLbStickinessPolicy) StateMust() *lightsailLbStickinessPolicyState {
	if llsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", llsp.Type(), llsp.LocalName()))
	}
	return llsp.state
}

func (llsp *LightsailLbStickinessPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(llsp)
}

type LightsailLbStickinessPolicyArgs struct {
	// CookieDuration: number, required
	CookieDuration terra.NumberValue `hcl:"cookie_duration,attr" validate:"required"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LbName: string, required
	LbName terra.StringValue `hcl:"lb_name,attr" validate:"required"`
	// DependsOn contains resources that LightsailLbStickinessPolicy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type lightsailLbStickinessPolicyAttributes struct {
	ref terra.Reference
}

func (llsp lightsailLbStickinessPolicyAttributes) CookieDuration() terra.NumberValue {
	return terra.ReferenceNumber(llsp.ref.Append("cookie_duration"))
}

func (llsp lightsailLbStickinessPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(llsp.ref.Append("enabled"))
}

func (llsp lightsailLbStickinessPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(llsp.ref.Append("id"))
}

func (llsp lightsailLbStickinessPolicyAttributes) LbName() terra.StringValue {
	return terra.ReferenceString(llsp.ref.Append("lb_name"))
}

type lightsailLbStickinessPolicyState struct {
	CookieDuration float64 `json:"cookie_duration"`
	Enabled        bool    `json:"enabled"`
	Id             string  `json:"id"`
	LbName         string  `json:"lb_name"`
}
