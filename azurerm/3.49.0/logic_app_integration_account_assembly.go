// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	logicappintegrationaccountassembly "github.com/golingon/terraproviders/azurerm/3.49.0/logicappintegrationaccountassembly"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewLogicAppIntegrationAccountAssembly(name string, args LogicAppIntegrationAccountAssemblyArgs) *LogicAppIntegrationAccountAssembly {
	return &LogicAppIntegrationAccountAssembly{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogicAppIntegrationAccountAssembly)(nil)

type LogicAppIntegrationAccountAssembly struct {
	Name  string
	Args  LogicAppIntegrationAccountAssemblyArgs
	state *logicAppIntegrationAccountAssemblyState
}

func (laiaa *LogicAppIntegrationAccountAssembly) Type() string {
	return "azurerm_logic_app_integration_account_assembly"
}

func (laiaa *LogicAppIntegrationAccountAssembly) LocalName() string {
	return laiaa.Name
}

func (laiaa *LogicAppIntegrationAccountAssembly) Configuration() interface{} {
	return laiaa.Args
}

func (laiaa *LogicAppIntegrationAccountAssembly) Attributes() logicAppIntegrationAccountAssemblyAttributes {
	return logicAppIntegrationAccountAssemblyAttributes{ref: terra.ReferenceResource(laiaa)}
}

func (laiaa *LogicAppIntegrationAccountAssembly) ImportState(av io.Reader) error {
	laiaa.state = &logicAppIntegrationAccountAssemblyState{}
	if err := json.NewDecoder(av).Decode(laiaa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", laiaa.Type(), laiaa.LocalName(), err)
	}
	return nil
}

func (laiaa *LogicAppIntegrationAccountAssembly) State() (*logicAppIntegrationAccountAssemblyState, bool) {
	return laiaa.state, laiaa.state != nil
}

func (laiaa *LogicAppIntegrationAccountAssembly) StateMust() *logicAppIntegrationAccountAssemblyState {
	if laiaa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", laiaa.Type(), laiaa.LocalName()))
	}
	return laiaa.state
}

func (laiaa *LogicAppIntegrationAccountAssembly) DependOn() terra.Reference {
	return terra.ReferenceResource(laiaa)
}

type LogicAppIntegrationAccountAssemblyArgs struct {
	// AssemblyName: string, required
	AssemblyName terra.StringValue `hcl:"assembly_name,attr" validate:"required"`
	// AssemblyVersion: string, optional
	AssemblyVersion terra.StringValue `hcl:"assembly_version,attr"`
	// Content: string, optional
	Content terra.StringValue `hcl:"content,attr"`
	// ContentLinkUri: string, optional
	ContentLinkUri terra.StringValue `hcl:"content_link_uri,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntegrationAccountName: string, required
	IntegrationAccountName terra.StringValue `hcl:"integration_account_name,attr" validate:"required"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *logicappintegrationaccountassembly.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that LogicAppIntegrationAccountAssembly depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type logicAppIntegrationAccountAssemblyAttributes struct {
	ref terra.Reference
}

func (laiaa logicAppIntegrationAccountAssemblyAttributes) AssemblyName() terra.StringValue {
	return terra.ReferenceString(laiaa.ref.Append("assembly_name"))
}

func (laiaa logicAppIntegrationAccountAssemblyAttributes) AssemblyVersion() terra.StringValue {
	return terra.ReferenceString(laiaa.ref.Append("assembly_version"))
}

func (laiaa logicAppIntegrationAccountAssemblyAttributes) Content() terra.StringValue {
	return terra.ReferenceString(laiaa.ref.Append("content"))
}

func (laiaa logicAppIntegrationAccountAssemblyAttributes) ContentLinkUri() terra.StringValue {
	return terra.ReferenceString(laiaa.ref.Append("content_link_uri"))
}

func (laiaa logicAppIntegrationAccountAssemblyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(laiaa.ref.Append("id"))
}

func (laiaa logicAppIntegrationAccountAssemblyAttributes) IntegrationAccountName() terra.StringValue {
	return terra.ReferenceString(laiaa.ref.Append("integration_account_name"))
}

func (laiaa logicAppIntegrationAccountAssemblyAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](laiaa.ref.Append("metadata"))
}

func (laiaa logicAppIntegrationAccountAssemblyAttributes) Name() terra.StringValue {
	return terra.ReferenceString(laiaa.ref.Append("name"))
}

func (laiaa logicAppIntegrationAccountAssemblyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(laiaa.ref.Append("resource_group_name"))
}

func (laiaa logicAppIntegrationAccountAssemblyAttributes) Timeouts() logicappintegrationaccountassembly.TimeoutsAttributes {
	return terra.ReferenceSingle[logicappintegrationaccountassembly.TimeoutsAttributes](laiaa.ref.Append("timeouts"))
}

type logicAppIntegrationAccountAssemblyState struct {
	AssemblyName           string                                            `json:"assembly_name"`
	AssemblyVersion        string                                            `json:"assembly_version"`
	Content                string                                            `json:"content"`
	ContentLinkUri         string                                            `json:"content_link_uri"`
	Id                     string                                            `json:"id"`
	IntegrationAccountName string                                            `json:"integration_account_name"`
	Metadata               map[string]string                                 `json:"metadata"`
	Name                   string                                            `json:"name"`
	ResourceGroupName      string                                            `json:"resource_group_name"`
	Timeouts               *logicappintegrationaccountassembly.TimeoutsState `json:"timeouts"`
}
