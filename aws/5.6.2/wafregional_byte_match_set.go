// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafregionalbytematchset "github.com/golingon/terraproviders/aws/5.6.2/wafregionalbytematchset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafregionalByteMatchSet creates a new instance of [WafregionalByteMatchSet].
func NewWafregionalByteMatchSet(name string, args WafregionalByteMatchSetArgs) *WafregionalByteMatchSet {
	return &WafregionalByteMatchSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafregionalByteMatchSet)(nil)

// WafregionalByteMatchSet represents the Terraform resource aws_wafregional_byte_match_set.
type WafregionalByteMatchSet struct {
	Name      string
	Args      WafregionalByteMatchSetArgs
	state     *wafregionalByteMatchSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafregionalByteMatchSet].
func (wbms *WafregionalByteMatchSet) Type() string {
	return "aws_wafregional_byte_match_set"
}

// LocalName returns the local name for [WafregionalByteMatchSet].
func (wbms *WafregionalByteMatchSet) LocalName() string {
	return wbms.Name
}

// Configuration returns the configuration (args) for [WafregionalByteMatchSet].
func (wbms *WafregionalByteMatchSet) Configuration() interface{} {
	return wbms.Args
}

// DependOn is used for other resources to depend on [WafregionalByteMatchSet].
func (wbms *WafregionalByteMatchSet) DependOn() terra.Reference {
	return terra.ReferenceResource(wbms)
}

// Dependencies returns the list of resources [WafregionalByteMatchSet] depends_on.
func (wbms *WafregionalByteMatchSet) Dependencies() terra.Dependencies {
	return wbms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafregionalByteMatchSet].
func (wbms *WafregionalByteMatchSet) LifecycleManagement() *terra.Lifecycle {
	return wbms.Lifecycle
}

// Attributes returns the attributes for [WafregionalByteMatchSet].
func (wbms *WafregionalByteMatchSet) Attributes() wafregionalByteMatchSetAttributes {
	return wafregionalByteMatchSetAttributes{ref: terra.ReferenceResource(wbms)}
}

// ImportState imports the given attribute values into [WafregionalByteMatchSet]'s state.
func (wbms *WafregionalByteMatchSet) ImportState(av io.Reader) error {
	wbms.state = &wafregionalByteMatchSetState{}
	if err := json.NewDecoder(av).Decode(wbms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wbms.Type(), wbms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafregionalByteMatchSet] has state.
func (wbms *WafregionalByteMatchSet) State() (*wafregionalByteMatchSetState, bool) {
	return wbms.state, wbms.state != nil
}

// StateMust returns the state for [WafregionalByteMatchSet]. Panics if the state is nil.
func (wbms *WafregionalByteMatchSet) StateMust() *wafregionalByteMatchSetState {
	if wbms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wbms.Type(), wbms.LocalName()))
	}
	return wbms.state
}

// WafregionalByteMatchSetArgs contains the configurations for aws_wafregional_byte_match_set.
type WafregionalByteMatchSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ByteMatchTuples: min=0
	ByteMatchTuples []wafregionalbytematchset.ByteMatchTuples `hcl:"byte_match_tuples,block" validate:"min=0"`
}
type wafregionalByteMatchSetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_wafregional_byte_match_set.
func (wbms wafregionalByteMatchSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wbms.ref.Append("id"))
}

// Name returns a reference to field name of aws_wafregional_byte_match_set.
func (wbms wafregionalByteMatchSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wbms.ref.Append("name"))
}

func (wbms wafregionalByteMatchSetAttributes) ByteMatchTuples() terra.SetValue[wafregionalbytematchset.ByteMatchTuplesAttributes] {
	return terra.ReferenceAsSet[wafregionalbytematchset.ByteMatchTuplesAttributes](wbms.ref.Append("byte_match_tuples"))
}

type wafregionalByteMatchSetState struct {
	Id              string                                         `json:"id"`
	Name            string                                         `json:"name"`
	ByteMatchTuples []wafregionalbytematchset.ByteMatchTuplesState `json:"byte_match_tuples"`
}
