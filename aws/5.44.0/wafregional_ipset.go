// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	wafregionalipset "github.com/golingon/terraproviders/aws/5.44.0/wafregionalipset"
	"io"
)

// NewWafregionalIpset creates a new instance of [WafregionalIpset].
func NewWafregionalIpset(name string, args WafregionalIpsetArgs) *WafregionalIpset {
	return &WafregionalIpset{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafregionalIpset)(nil)

// WafregionalIpset represents the Terraform resource aws_wafregional_ipset.
type WafregionalIpset struct {
	Name      string
	Args      WafregionalIpsetArgs
	state     *wafregionalIpsetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafregionalIpset].
func (wi *WafregionalIpset) Type() string {
	return "aws_wafregional_ipset"
}

// LocalName returns the local name for [WafregionalIpset].
func (wi *WafregionalIpset) LocalName() string {
	return wi.Name
}

// Configuration returns the configuration (args) for [WafregionalIpset].
func (wi *WafregionalIpset) Configuration() interface{} {
	return wi.Args
}

// DependOn is used for other resources to depend on [WafregionalIpset].
func (wi *WafregionalIpset) DependOn() terra.Reference {
	return terra.ReferenceResource(wi)
}

// Dependencies returns the list of resources [WafregionalIpset] depends_on.
func (wi *WafregionalIpset) Dependencies() terra.Dependencies {
	return wi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafregionalIpset].
func (wi *WafregionalIpset) LifecycleManagement() *terra.Lifecycle {
	return wi.Lifecycle
}

// Attributes returns the attributes for [WafregionalIpset].
func (wi *WafregionalIpset) Attributes() wafregionalIpsetAttributes {
	return wafregionalIpsetAttributes{ref: terra.ReferenceResource(wi)}
}

// ImportState imports the given attribute values into [WafregionalIpset]'s state.
func (wi *WafregionalIpset) ImportState(av io.Reader) error {
	wi.state = &wafregionalIpsetState{}
	if err := json.NewDecoder(av).Decode(wi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wi.Type(), wi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafregionalIpset] has state.
func (wi *WafregionalIpset) State() (*wafregionalIpsetState, bool) {
	return wi.state, wi.state != nil
}

// StateMust returns the state for [WafregionalIpset]. Panics if the state is nil.
func (wi *WafregionalIpset) StateMust() *wafregionalIpsetState {
	if wi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wi.Type(), wi.LocalName()))
	}
	return wi.state
}

// WafregionalIpsetArgs contains the configurations for aws_wafregional_ipset.
type WafregionalIpsetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// IpSetDescriptor: min=0
	IpSetDescriptor []wafregionalipset.IpSetDescriptor `hcl:"ip_set_descriptor,block" validate:"min=0"`
}
type wafregionalIpsetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_wafregional_ipset.
func (wi wafregionalIpsetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wi.ref.Append("arn"))
}

// Id returns a reference to field id of aws_wafregional_ipset.
func (wi wafregionalIpsetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wi.ref.Append("id"))
}

// Name returns a reference to field name of aws_wafregional_ipset.
func (wi wafregionalIpsetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wi.ref.Append("name"))
}

func (wi wafregionalIpsetAttributes) IpSetDescriptor() terra.SetValue[wafregionalipset.IpSetDescriptorAttributes] {
	return terra.ReferenceAsSet[wafregionalipset.IpSetDescriptorAttributes](wi.ref.Append("ip_set_descriptor"))
}

type wafregionalIpsetState struct {
	Arn             string                                  `json:"arn"`
	Id              string                                  `json:"id"`
	Name            string                                  `json:"name"`
	IpSetDescriptor []wafregionalipset.IpSetDescriptorState `json:"ip_set_descriptor"`
}
