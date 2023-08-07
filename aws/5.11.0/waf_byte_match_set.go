// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafbytematchset "github.com/golingon/terraproviders/aws/5.11.0/wafbytematchset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafByteMatchSet creates a new instance of [WafByteMatchSet].
func NewWafByteMatchSet(name string, args WafByteMatchSetArgs) *WafByteMatchSet {
	return &WafByteMatchSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafByteMatchSet)(nil)

// WafByteMatchSet represents the Terraform resource aws_waf_byte_match_set.
type WafByteMatchSet struct {
	Name      string
	Args      WafByteMatchSetArgs
	state     *wafByteMatchSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafByteMatchSet].
func (wbms *WafByteMatchSet) Type() string {
	return "aws_waf_byte_match_set"
}

// LocalName returns the local name for [WafByteMatchSet].
func (wbms *WafByteMatchSet) LocalName() string {
	return wbms.Name
}

// Configuration returns the configuration (args) for [WafByteMatchSet].
func (wbms *WafByteMatchSet) Configuration() interface{} {
	return wbms.Args
}

// DependOn is used for other resources to depend on [WafByteMatchSet].
func (wbms *WafByteMatchSet) DependOn() terra.Reference {
	return terra.ReferenceResource(wbms)
}

// Dependencies returns the list of resources [WafByteMatchSet] depends_on.
func (wbms *WafByteMatchSet) Dependencies() terra.Dependencies {
	return wbms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafByteMatchSet].
func (wbms *WafByteMatchSet) LifecycleManagement() *terra.Lifecycle {
	return wbms.Lifecycle
}

// Attributes returns the attributes for [WafByteMatchSet].
func (wbms *WafByteMatchSet) Attributes() wafByteMatchSetAttributes {
	return wafByteMatchSetAttributes{ref: terra.ReferenceResource(wbms)}
}

// ImportState imports the given attribute values into [WafByteMatchSet]'s state.
func (wbms *WafByteMatchSet) ImportState(av io.Reader) error {
	wbms.state = &wafByteMatchSetState{}
	if err := json.NewDecoder(av).Decode(wbms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wbms.Type(), wbms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafByteMatchSet] has state.
func (wbms *WafByteMatchSet) State() (*wafByteMatchSetState, bool) {
	return wbms.state, wbms.state != nil
}

// StateMust returns the state for [WafByteMatchSet]. Panics if the state is nil.
func (wbms *WafByteMatchSet) StateMust() *wafByteMatchSetState {
	if wbms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wbms.Type(), wbms.LocalName()))
	}
	return wbms.state
}

// WafByteMatchSetArgs contains the configurations for aws_waf_byte_match_set.
type WafByteMatchSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ByteMatchTuples: min=0
	ByteMatchTuples []wafbytematchset.ByteMatchTuples `hcl:"byte_match_tuples,block" validate:"min=0"`
}
type wafByteMatchSetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_waf_byte_match_set.
func (wbms wafByteMatchSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wbms.ref.Append("id"))
}

// Name returns a reference to field name of aws_waf_byte_match_set.
func (wbms wafByteMatchSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wbms.ref.Append("name"))
}

func (wbms wafByteMatchSetAttributes) ByteMatchTuples() terra.SetValue[wafbytematchset.ByteMatchTuplesAttributes] {
	return terra.ReferenceAsSet[wafbytematchset.ByteMatchTuplesAttributes](wbms.ref.Append("byte_match_tuples"))
}

type wafByteMatchSetState struct {
	Id              string                                 `json:"id"`
	Name            string                                 `json:"name"`
	ByteMatchTuples []wafbytematchset.ByteMatchTuplesState `json:"byte_match_tuples"`
}
