// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafsqlinjectionmatchset "github.com/golingon/terraproviders/aws/4.63.0/wafsqlinjectionmatchset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafSqlInjectionMatchSet creates a new instance of [WafSqlInjectionMatchSet].
func NewWafSqlInjectionMatchSet(name string, args WafSqlInjectionMatchSetArgs) *WafSqlInjectionMatchSet {
	return &WafSqlInjectionMatchSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafSqlInjectionMatchSet)(nil)

// WafSqlInjectionMatchSet represents the Terraform resource aws_waf_sql_injection_match_set.
type WafSqlInjectionMatchSet struct {
	Name      string
	Args      WafSqlInjectionMatchSetArgs
	state     *wafSqlInjectionMatchSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafSqlInjectionMatchSet].
func (wsims *WafSqlInjectionMatchSet) Type() string {
	return "aws_waf_sql_injection_match_set"
}

// LocalName returns the local name for [WafSqlInjectionMatchSet].
func (wsims *WafSqlInjectionMatchSet) LocalName() string {
	return wsims.Name
}

// Configuration returns the configuration (args) for [WafSqlInjectionMatchSet].
func (wsims *WafSqlInjectionMatchSet) Configuration() interface{} {
	return wsims.Args
}

// DependOn is used for other resources to depend on [WafSqlInjectionMatchSet].
func (wsims *WafSqlInjectionMatchSet) DependOn() terra.Reference {
	return terra.ReferenceResource(wsims)
}

// Dependencies returns the list of resources [WafSqlInjectionMatchSet] depends_on.
func (wsims *WafSqlInjectionMatchSet) Dependencies() terra.Dependencies {
	return wsims.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafSqlInjectionMatchSet].
func (wsims *WafSqlInjectionMatchSet) LifecycleManagement() *terra.Lifecycle {
	return wsims.Lifecycle
}

// Attributes returns the attributes for [WafSqlInjectionMatchSet].
func (wsims *WafSqlInjectionMatchSet) Attributes() wafSqlInjectionMatchSetAttributes {
	return wafSqlInjectionMatchSetAttributes{ref: terra.ReferenceResource(wsims)}
}

// ImportState imports the given attribute values into [WafSqlInjectionMatchSet]'s state.
func (wsims *WafSqlInjectionMatchSet) ImportState(av io.Reader) error {
	wsims.state = &wafSqlInjectionMatchSetState{}
	if err := json.NewDecoder(av).Decode(wsims.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wsims.Type(), wsims.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafSqlInjectionMatchSet] has state.
func (wsims *WafSqlInjectionMatchSet) State() (*wafSqlInjectionMatchSetState, bool) {
	return wsims.state, wsims.state != nil
}

// StateMust returns the state for [WafSqlInjectionMatchSet]. Panics if the state is nil.
func (wsims *WafSqlInjectionMatchSet) StateMust() *wafSqlInjectionMatchSetState {
	if wsims.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wsims.Type(), wsims.LocalName()))
	}
	return wsims.state
}

// WafSqlInjectionMatchSetArgs contains the configurations for aws_waf_sql_injection_match_set.
type WafSqlInjectionMatchSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SqlInjectionMatchTuples: min=0
	SqlInjectionMatchTuples []wafsqlinjectionmatchset.SqlInjectionMatchTuples `hcl:"sql_injection_match_tuples,block" validate:"min=0"`
}
type wafSqlInjectionMatchSetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_waf_sql_injection_match_set.
func (wsims wafSqlInjectionMatchSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wsims.ref.Append("id"))
}

// Name returns a reference to field name of aws_waf_sql_injection_match_set.
func (wsims wafSqlInjectionMatchSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wsims.ref.Append("name"))
}

func (wsims wafSqlInjectionMatchSetAttributes) SqlInjectionMatchTuples() terra.SetValue[wafsqlinjectionmatchset.SqlInjectionMatchTuplesAttributes] {
	return terra.ReferenceAsSet[wafsqlinjectionmatchset.SqlInjectionMatchTuplesAttributes](wsims.ref.Append("sql_injection_match_tuples"))
}

type wafSqlInjectionMatchSetState struct {
	Id                      string                                                 `json:"id"`
	Name                    string                                                 `json:"name"`
	SqlInjectionMatchTuples []wafsqlinjectionmatchset.SqlInjectionMatchTuplesState `json:"sql_injection_match_tuples"`
}
