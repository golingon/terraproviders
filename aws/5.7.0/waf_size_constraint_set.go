// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafsizeconstraintset "github.com/golingon/terraproviders/aws/5.7.0/wafsizeconstraintset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafSizeConstraintSet creates a new instance of [WafSizeConstraintSet].
func NewWafSizeConstraintSet(name string, args WafSizeConstraintSetArgs) *WafSizeConstraintSet {
	return &WafSizeConstraintSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafSizeConstraintSet)(nil)

// WafSizeConstraintSet represents the Terraform resource aws_waf_size_constraint_set.
type WafSizeConstraintSet struct {
	Name      string
	Args      WafSizeConstraintSetArgs
	state     *wafSizeConstraintSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafSizeConstraintSet].
func (wscs *WafSizeConstraintSet) Type() string {
	return "aws_waf_size_constraint_set"
}

// LocalName returns the local name for [WafSizeConstraintSet].
func (wscs *WafSizeConstraintSet) LocalName() string {
	return wscs.Name
}

// Configuration returns the configuration (args) for [WafSizeConstraintSet].
func (wscs *WafSizeConstraintSet) Configuration() interface{} {
	return wscs.Args
}

// DependOn is used for other resources to depend on [WafSizeConstraintSet].
func (wscs *WafSizeConstraintSet) DependOn() terra.Reference {
	return terra.ReferenceResource(wscs)
}

// Dependencies returns the list of resources [WafSizeConstraintSet] depends_on.
func (wscs *WafSizeConstraintSet) Dependencies() terra.Dependencies {
	return wscs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafSizeConstraintSet].
func (wscs *WafSizeConstraintSet) LifecycleManagement() *terra.Lifecycle {
	return wscs.Lifecycle
}

// Attributes returns the attributes for [WafSizeConstraintSet].
func (wscs *WafSizeConstraintSet) Attributes() wafSizeConstraintSetAttributes {
	return wafSizeConstraintSetAttributes{ref: terra.ReferenceResource(wscs)}
}

// ImportState imports the given attribute values into [WafSizeConstraintSet]'s state.
func (wscs *WafSizeConstraintSet) ImportState(av io.Reader) error {
	wscs.state = &wafSizeConstraintSetState{}
	if err := json.NewDecoder(av).Decode(wscs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wscs.Type(), wscs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafSizeConstraintSet] has state.
func (wscs *WafSizeConstraintSet) State() (*wafSizeConstraintSetState, bool) {
	return wscs.state, wscs.state != nil
}

// StateMust returns the state for [WafSizeConstraintSet]. Panics if the state is nil.
func (wscs *WafSizeConstraintSet) StateMust() *wafSizeConstraintSetState {
	if wscs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wscs.Type(), wscs.LocalName()))
	}
	return wscs.state
}

// WafSizeConstraintSetArgs contains the configurations for aws_waf_size_constraint_set.
type WafSizeConstraintSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SizeConstraints: min=0
	SizeConstraints []wafsizeconstraintset.SizeConstraints `hcl:"size_constraints,block" validate:"min=0"`
}
type wafSizeConstraintSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_waf_size_constraint_set.
func (wscs wafSizeConstraintSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wscs.ref.Append("arn"))
}

// Id returns a reference to field id of aws_waf_size_constraint_set.
func (wscs wafSizeConstraintSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wscs.ref.Append("id"))
}

// Name returns a reference to field name of aws_waf_size_constraint_set.
func (wscs wafSizeConstraintSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wscs.ref.Append("name"))
}

func (wscs wafSizeConstraintSetAttributes) SizeConstraints() terra.SetValue[wafsizeconstraintset.SizeConstraintsAttributes] {
	return terra.ReferenceAsSet[wafsizeconstraintset.SizeConstraintsAttributes](wscs.ref.Append("size_constraints"))
}

type wafSizeConstraintSetState struct {
	Arn             string                                      `json:"arn"`
	Id              string                                      `json:"id"`
	Name            string                                      `json:"name"`
	SizeConstraints []wafsizeconstraintset.SizeConstraintsState `json:"size_constraints"`
}
