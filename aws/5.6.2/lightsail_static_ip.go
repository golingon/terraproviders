// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLightsailStaticIp creates a new instance of [LightsailStaticIp].
func NewLightsailStaticIp(name string, args LightsailStaticIpArgs) *LightsailStaticIp {
	return &LightsailStaticIp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LightsailStaticIp)(nil)

// LightsailStaticIp represents the Terraform resource aws_lightsail_static_ip.
type LightsailStaticIp struct {
	Name      string
	Args      LightsailStaticIpArgs
	state     *lightsailStaticIpState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LightsailStaticIp].
func (lsi *LightsailStaticIp) Type() string {
	return "aws_lightsail_static_ip"
}

// LocalName returns the local name for [LightsailStaticIp].
func (lsi *LightsailStaticIp) LocalName() string {
	return lsi.Name
}

// Configuration returns the configuration (args) for [LightsailStaticIp].
func (lsi *LightsailStaticIp) Configuration() interface{} {
	return lsi.Args
}

// DependOn is used for other resources to depend on [LightsailStaticIp].
func (lsi *LightsailStaticIp) DependOn() terra.Reference {
	return terra.ReferenceResource(lsi)
}

// Dependencies returns the list of resources [LightsailStaticIp] depends_on.
func (lsi *LightsailStaticIp) Dependencies() terra.Dependencies {
	return lsi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LightsailStaticIp].
func (lsi *LightsailStaticIp) LifecycleManagement() *terra.Lifecycle {
	return lsi.Lifecycle
}

// Attributes returns the attributes for [LightsailStaticIp].
func (lsi *LightsailStaticIp) Attributes() lightsailStaticIpAttributes {
	return lightsailStaticIpAttributes{ref: terra.ReferenceResource(lsi)}
}

// ImportState imports the given attribute values into [LightsailStaticIp]'s state.
func (lsi *LightsailStaticIp) ImportState(av io.Reader) error {
	lsi.state = &lightsailStaticIpState{}
	if err := json.NewDecoder(av).Decode(lsi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lsi.Type(), lsi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LightsailStaticIp] has state.
func (lsi *LightsailStaticIp) State() (*lightsailStaticIpState, bool) {
	return lsi.state, lsi.state != nil
}

// StateMust returns the state for [LightsailStaticIp]. Panics if the state is nil.
func (lsi *LightsailStaticIp) StateMust() *lightsailStaticIpState {
	if lsi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lsi.Type(), lsi.LocalName()))
	}
	return lsi.state
}

// LightsailStaticIpArgs contains the configurations for aws_lightsail_static_ip.
type LightsailStaticIpArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type lightsailStaticIpAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lightsail_static_ip.
func (lsi lightsailStaticIpAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(lsi.ref.Append("arn"))
}

// Id returns a reference to field id of aws_lightsail_static_ip.
func (lsi lightsailStaticIpAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lsi.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of aws_lightsail_static_ip.
func (lsi lightsailStaticIpAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(lsi.ref.Append("ip_address"))
}

// Name returns a reference to field name of aws_lightsail_static_ip.
func (lsi lightsailStaticIpAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lsi.ref.Append("name"))
}

// SupportCode returns a reference to field support_code of aws_lightsail_static_ip.
func (lsi lightsailStaticIpAttributes) SupportCode() terra.StringValue {
	return terra.ReferenceAsString(lsi.ref.Append("support_code"))
}

type lightsailStaticIpState struct {
	Arn         string `json:"arn"`
	Id          string `json:"id"`
	IpAddress   string `json:"ip_address"`
	Name        string `json:"name"`
	SupportCode string `json:"support_code"`
}
