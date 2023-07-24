// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafregionalsizeconstraintset "github.com/golingon/terraproviders/aws/5.9.0/wafregionalsizeconstraintset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafregionalSizeConstraintSet creates a new instance of [WafregionalSizeConstraintSet].
func NewWafregionalSizeConstraintSet(name string, args WafregionalSizeConstraintSetArgs) *WafregionalSizeConstraintSet {
	return &WafregionalSizeConstraintSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafregionalSizeConstraintSet)(nil)

// WafregionalSizeConstraintSet represents the Terraform resource aws_wafregional_size_constraint_set.
type WafregionalSizeConstraintSet struct {
	Name      string
	Args      WafregionalSizeConstraintSetArgs
	state     *wafregionalSizeConstraintSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafregionalSizeConstraintSet].
func (wscs *WafregionalSizeConstraintSet) Type() string {
	return "aws_wafregional_size_constraint_set"
}

// LocalName returns the local name for [WafregionalSizeConstraintSet].
func (wscs *WafregionalSizeConstraintSet) LocalName() string {
	return wscs.Name
}

// Configuration returns the configuration (args) for [WafregionalSizeConstraintSet].
func (wscs *WafregionalSizeConstraintSet) Configuration() interface{} {
	return wscs.Args
}

// DependOn is used for other resources to depend on [WafregionalSizeConstraintSet].
func (wscs *WafregionalSizeConstraintSet) DependOn() terra.Reference {
	return terra.ReferenceResource(wscs)
}

// Dependencies returns the list of resources [WafregionalSizeConstraintSet] depends_on.
func (wscs *WafregionalSizeConstraintSet) Dependencies() terra.Dependencies {
	return wscs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafregionalSizeConstraintSet].
func (wscs *WafregionalSizeConstraintSet) LifecycleManagement() *terra.Lifecycle {
	return wscs.Lifecycle
}

// Attributes returns the attributes for [WafregionalSizeConstraintSet].
func (wscs *WafregionalSizeConstraintSet) Attributes() wafregionalSizeConstraintSetAttributes {
	return wafregionalSizeConstraintSetAttributes{ref: terra.ReferenceResource(wscs)}
}

// ImportState imports the given attribute values into [WafregionalSizeConstraintSet]'s state.
func (wscs *WafregionalSizeConstraintSet) ImportState(av io.Reader) error {
	wscs.state = &wafregionalSizeConstraintSetState{}
	if err := json.NewDecoder(av).Decode(wscs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wscs.Type(), wscs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafregionalSizeConstraintSet] has state.
func (wscs *WafregionalSizeConstraintSet) State() (*wafregionalSizeConstraintSetState, bool) {
	return wscs.state, wscs.state != nil
}

// StateMust returns the state for [WafregionalSizeConstraintSet]. Panics if the state is nil.
func (wscs *WafregionalSizeConstraintSet) StateMust() *wafregionalSizeConstraintSetState {
	if wscs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wscs.Type(), wscs.LocalName()))
	}
	return wscs.state
}

// WafregionalSizeConstraintSetArgs contains the configurations for aws_wafregional_size_constraint_set.
type WafregionalSizeConstraintSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SizeConstraints: min=0
	SizeConstraints []wafregionalsizeconstraintset.SizeConstraints `hcl:"size_constraints,block" validate:"min=0"`
}
type wafregionalSizeConstraintSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_wafregional_size_constraint_set.
func (wscs wafregionalSizeConstraintSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wscs.ref.Append("arn"))
}

// Id returns a reference to field id of aws_wafregional_size_constraint_set.
func (wscs wafregionalSizeConstraintSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wscs.ref.Append("id"))
}

// Name returns a reference to field name of aws_wafregional_size_constraint_set.
func (wscs wafregionalSizeConstraintSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wscs.ref.Append("name"))
}

func (wscs wafregionalSizeConstraintSetAttributes) SizeConstraints() terra.SetValue[wafregionalsizeconstraintset.SizeConstraintsAttributes] {
	return terra.ReferenceAsSet[wafregionalsizeconstraintset.SizeConstraintsAttributes](wscs.ref.Append("size_constraints"))
}

type wafregionalSizeConstraintSetState struct {
	Arn             string                                              `json:"arn"`
	Id              string                                              `json:"id"`
	Name            string                                              `json:"name"`
	SizeConstraints []wafregionalsizeconstraintset.SizeConstraintsState `json:"size_constraints"`
}
