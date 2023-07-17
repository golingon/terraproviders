// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafregionalsqlinjectionmatchset "github.com/golingon/terraproviders/aws/5.7.0/wafregionalsqlinjectionmatchset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafregionalSqlInjectionMatchSet creates a new instance of [WafregionalSqlInjectionMatchSet].
func NewWafregionalSqlInjectionMatchSet(name string, args WafregionalSqlInjectionMatchSetArgs) *WafregionalSqlInjectionMatchSet {
	return &WafregionalSqlInjectionMatchSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafregionalSqlInjectionMatchSet)(nil)

// WafregionalSqlInjectionMatchSet represents the Terraform resource aws_wafregional_sql_injection_match_set.
type WafregionalSqlInjectionMatchSet struct {
	Name      string
	Args      WafregionalSqlInjectionMatchSetArgs
	state     *wafregionalSqlInjectionMatchSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafregionalSqlInjectionMatchSet].
func (wsims *WafregionalSqlInjectionMatchSet) Type() string {
	return "aws_wafregional_sql_injection_match_set"
}

// LocalName returns the local name for [WafregionalSqlInjectionMatchSet].
func (wsims *WafregionalSqlInjectionMatchSet) LocalName() string {
	return wsims.Name
}

// Configuration returns the configuration (args) for [WafregionalSqlInjectionMatchSet].
func (wsims *WafregionalSqlInjectionMatchSet) Configuration() interface{} {
	return wsims.Args
}

// DependOn is used for other resources to depend on [WafregionalSqlInjectionMatchSet].
func (wsims *WafregionalSqlInjectionMatchSet) DependOn() terra.Reference {
	return terra.ReferenceResource(wsims)
}

// Dependencies returns the list of resources [WafregionalSqlInjectionMatchSet] depends_on.
func (wsims *WafregionalSqlInjectionMatchSet) Dependencies() terra.Dependencies {
	return wsims.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafregionalSqlInjectionMatchSet].
func (wsims *WafregionalSqlInjectionMatchSet) LifecycleManagement() *terra.Lifecycle {
	return wsims.Lifecycle
}

// Attributes returns the attributes for [WafregionalSqlInjectionMatchSet].
func (wsims *WafregionalSqlInjectionMatchSet) Attributes() wafregionalSqlInjectionMatchSetAttributes {
	return wafregionalSqlInjectionMatchSetAttributes{ref: terra.ReferenceResource(wsims)}
}

// ImportState imports the given attribute values into [WafregionalSqlInjectionMatchSet]'s state.
func (wsims *WafregionalSqlInjectionMatchSet) ImportState(av io.Reader) error {
	wsims.state = &wafregionalSqlInjectionMatchSetState{}
	if err := json.NewDecoder(av).Decode(wsims.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wsims.Type(), wsims.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafregionalSqlInjectionMatchSet] has state.
func (wsims *WafregionalSqlInjectionMatchSet) State() (*wafregionalSqlInjectionMatchSetState, bool) {
	return wsims.state, wsims.state != nil
}

// StateMust returns the state for [WafregionalSqlInjectionMatchSet]. Panics if the state is nil.
func (wsims *WafregionalSqlInjectionMatchSet) StateMust() *wafregionalSqlInjectionMatchSetState {
	if wsims.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wsims.Type(), wsims.LocalName()))
	}
	return wsims.state
}

// WafregionalSqlInjectionMatchSetArgs contains the configurations for aws_wafregional_sql_injection_match_set.
type WafregionalSqlInjectionMatchSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SqlInjectionMatchTuple: min=0
	SqlInjectionMatchTuple []wafregionalsqlinjectionmatchset.SqlInjectionMatchTuple `hcl:"sql_injection_match_tuple,block" validate:"min=0"`
}
type wafregionalSqlInjectionMatchSetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_wafregional_sql_injection_match_set.
func (wsims wafregionalSqlInjectionMatchSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wsims.ref.Append("id"))
}

// Name returns a reference to field name of aws_wafregional_sql_injection_match_set.
func (wsims wafregionalSqlInjectionMatchSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wsims.ref.Append("name"))
}

func (wsims wafregionalSqlInjectionMatchSetAttributes) SqlInjectionMatchTuple() terra.SetValue[wafregionalsqlinjectionmatchset.SqlInjectionMatchTupleAttributes] {
	return terra.ReferenceAsSet[wafregionalsqlinjectionmatchset.SqlInjectionMatchTupleAttributes](wsims.ref.Append("sql_injection_match_tuple"))
}

type wafregionalSqlInjectionMatchSetState struct {
	Id                     string                                                        `json:"id"`
	Name                   string                                                        `json:"name"`
	SqlInjectionMatchTuple []wafregionalsqlinjectionmatchset.SqlInjectionMatchTupleState `json:"sql_injection_match_tuple"`
}