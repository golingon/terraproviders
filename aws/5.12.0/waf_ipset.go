// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafipset "github.com/golingon/terraproviders/aws/5.12.0/wafipset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafIpset creates a new instance of [WafIpset].
func NewWafIpset(name string, args WafIpsetArgs) *WafIpset {
	return &WafIpset{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafIpset)(nil)

// WafIpset represents the Terraform resource aws_waf_ipset.
type WafIpset struct {
	Name      string
	Args      WafIpsetArgs
	state     *wafIpsetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafIpset].
func (wi *WafIpset) Type() string {
	return "aws_waf_ipset"
}

// LocalName returns the local name for [WafIpset].
func (wi *WafIpset) LocalName() string {
	return wi.Name
}

// Configuration returns the configuration (args) for [WafIpset].
func (wi *WafIpset) Configuration() interface{} {
	return wi.Args
}

// DependOn is used for other resources to depend on [WafIpset].
func (wi *WafIpset) DependOn() terra.Reference {
	return terra.ReferenceResource(wi)
}

// Dependencies returns the list of resources [WafIpset] depends_on.
func (wi *WafIpset) Dependencies() terra.Dependencies {
	return wi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafIpset].
func (wi *WafIpset) LifecycleManagement() *terra.Lifecycle {
	return wi.Lifecycle
}

// Attributes returns the attributes for [WafIpset].
func (wi *WafIpset) Attributes() wafIpsetAttributes {
	return wafIpsetAttributes{ref: terra.ReferenceResource(wi)}
}

// ImportState imports the given attribute values into [WafIpset]'s state.
func (wi *WafIpset) ImportState(av io.Reader) error {
	wi.state = &wafIpsetState{}
	if err := json.NewDecoder(av).Decode(wi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wi.Type(), wi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafIpset] has state.
func (wi *WafIpset) State() (*wafIpsetState, bool) {
	return wi.state, wi.state != nil
}

// StateMust returns the state for [WafIpset]. Panics if the state is nil.
func (wi *WafIpset) StateMust() *wafIpsetState {
	if wi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wi.Type(), wi.LocalName()))
	}
	return wi.state
}

// WafIpsetArgs contains the configurations for aws_waf_ipset.
type WafIpsetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// IpSetDescriptors: min=0
	IpSetDescriptors []wafipset.IpSetDescriptors `hcl:"ip_set_descriptors,block" validate:"min=0"`
}
type wafIpsetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_waf_ipset.
func (wi wafIpsetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wi.ref.Append("arn"))
}

// Id returns a reference to field id of aws_waf_ipset.
func (wi wafIpsetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wi.ref.Append("id"))
}

// Name returns a reference to field name of aws_waf_ipset.
func (wi wafIpsetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wi.ref.Append("name"))
}

func (wi wafIpsetAttributes) IpSetDescriptors() terra.SetValue[wafipset.IpSetDescriptorsAttributes] {
	return terra.ReferenceAsSet[wafipset.IpSetDescriptorsAttributes](wi.ref.Append("ip_set_descriptors"))
}

type wafIpsetState struct {
	Arn              string                           `json:"arn"`
	Id               string                           `json:"id"`
	Name             string                           `json:"name"`
	IpSetDescriptors []wafipset.IpSetDescriptorsState `json:"ip_set_descriptors"`
}
