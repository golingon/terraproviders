// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	kustoscript "github.com/golingon/terraproviders/azurerm/3.67.0/kustoscript"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKustoScript creates a new instance of [KustoScript].
func NewKustoScript(name string, args KustoScriptArgs) *KustoScript {
	return &KustoScript{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KustoScript)(nil)

// KustoScript represents the Terraform resource azurerm_kusto_script.
type KustoScript struct {
	Name      string
	Args      KustoScriptArgs
	state     *kustoScriptState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KustoScript].
func (ks *KustoScript) Type() string {
	return "azurerm_kusto_script"
}

// LocalName returns the local name for [KustoScript].
func (ks *KustoScript) LocalName() string {
	return ks.Name
}

// Configuration returns the configuration (args) for [KustoScript].
func (ks *KustoScript) Configuration() interface{} {
	return ks.Args
}

// DependOn is used for other resources to depend on [KustoScript].
func (ks *KustoScript) DependOn() terra.Reference {
	return terra.ReferenceResource(ks)
}

// Dependencies returns the list of resources [KustoScript] depends_on.
func (ks *KustoScript) Dependencies() terra.Dependencies {
	return ks.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KustoScript].
func (ks *KustoScript) LifecycleManagement() *terra.Lifecycle {
	return ks.Lifecycle
}

// Attributes returns the attributes for [KustoScript].
func (ks *KustoScript) Attributes() kustoScriptAttributes {
	return kustoScriptAttributes{ref: terra.ReferenceResource(ks)}
}

// ImportState imports the given attribute values into [KustoScript]'s state.
func (ks *KustoScript) ImportState(av io.Reader) error {
	ks.state = &kustoScriptState{}
	if err := json.NewDecoder(av).Decode(ks.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ks.Type(), ks.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KustoScript] has state.
func (ks *KustoScript) State() (*kustoScriptState, bool) {
	return ks.state, ks.state != nil
}

// StateMust returns the state for [KustoScript]. Panics if the state is nil.
func (ks *KustoScript) StateMust() *kustoScriptState {
	if ks.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ks.Type(), ks.LocalName()))
	}
	return ks.state
}

// KustoScriptArgs contains the configurations for azurerm_kusto_script.
type KustoScriptArgs struct {
	// ContinueOnErrorsEnabled: bool, optional
	ContinueOnErrorsEnabled terra.BoolValue `hcl:"continue_on_errors_enabled,attr"`
	// DatabaseId: string, required
	DatabaseId terra.StringValue `hcl:"database_id,attr" validate:"required"`
	// ForceAnUpdateWhenValueChanged: string, optional
	ForceAnUpdateWhenValueChanged terra.StringValue `hcl:"force_an_update_when_value_changed,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SasToken: string, optional
	SasToken terra.StringValue `hcl:"sas_token,attr"`
	// ScriptContent: string, optional
	ScriptContent terra.StringValue `hcl:"script_content,attr"`
	// Url: string, optional
	Url terra.StringValue `hcl:"url,attr"`
	// Timeouts: optional
	Timeouts *kustoscript.Timeouts `hcl:"timeouts,block"`
}
type kustoScriptAttributes struct {
	ref terra.Reference
}

// ContinueOnErrorsEnabled returns a reference to field continue_on_errors_enabled of azurerm_kusto_script.
func (ks kustoScriptAttributes) ContinueOnErrorsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ks.ref.Append("continue_on_errors_enabled"))
}

// DatabaseId returns a reference to field database_id of azurerm_kusto_script.
func (ks kustoScriptAttributes) DatabaseId() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("database_id"))
}

// ForceAnUpdateWhenValueChanged returns a reference to field force_an_update_when_value_changed of azurerm_kusto_script.
func (ks kustoScriptAttributes) ForceAnUpdateWhenValueChanged() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("force_an_update_when_value_changed"))
}

// Id returns a reference to field id of azurerm_kusto_script.
func (ks kustoScriptAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_kusto_script.
func (ks kustoScriptAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("name"))
}

// SasToken returns a reference to field sas_token of azurerm_kusto_script.
func (ks kustoScriptAttributes) SasToken() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("sas_token"))
}

// ScriptContent returns a reference to field script_content of azurerm_kusto_script.
func (ks kustoScriptAttributes) ScriptContent() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("script_content"))
}

// Url returns a reference to field url of azurerm_kusto_script.
func (ks kustoScriptAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("url"))
}

func (ks kustoScriptAttributes) Timeouts() kustoscript.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kustoscript.TimeoutsAttributes](ks.ref.Append("timeouts"))
}

type kustoScriptState struct {
	ContinueOnErrorsEnabled       bool                       `json:"continue_on_errors_enabled"`
	DatabaseId                    string                     `json:"database_id"`
	ForceAnUpdateWhenValueChanged string                     `json:"force_an_update_when_value_changed"`
	Id                            string                     `json:"id"`
	Name                          string                     `json:"name"`
	SasToken                      string                     `json:"sas_token"`
	ScriptContent                 string                     `json:"script_content"`
	Url                           string                     `json:"url"`
	Timeouts                      *kustoscript.TimeoutsState `json:"timeouts"`
}
