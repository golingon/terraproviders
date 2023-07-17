// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqljobagent "github.com/golingon/terraproviders/azurerm/3.65.0/mssqljobagent"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlJobAgent creates a new instance of [MssqlJobAgent].
func NewMssqlJobAgent(name string, args MssqlJobAgentArgs) *MssqlJobAgent {
	return &MssqlJobAgent{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlJobAgent)(nil)

// MssqlJobAgent represents the Terraform resource azurerm_mssql_job_agent.
type MssqlJobAgent struct {
	Name      string
	Args      MssqlJobAgentArgs
	state     *mssqlJobAgentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlJobAgent].
func (mja *MssqlJobAgent) Type() string {
	return "azurerm_mssql_job_agent"
}

// LocalName returns the local name for [MssqlJobAgent].
func (mja *MssqlJobAgent) LocalName() string {
	return mja.Name
}

// Configuration returns the configuration (args) for [MssqlJobAgent].
func (mja *MssqlJobAgent) Configuration() interface{} {
	return mja.Args
}

// DependOn is used for other resources to depend on [MssqlJobAgent].
func (mja *MssqlJobAgent) DependOn() terra.Reference {
	return terra.ReferenceResource(mja)
}

// Dependencies returns the list of resources [MssqlJobAgent] depends_on.
func (mja *MssqlJobAgent) Dependencies() terra.Dependencies {
	return mja.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlJobAgent].
func (mja *MssqlJobAgent) LifecycleManagement() *terra.Lifecycle {
	return mja.Lifecycle
}

// Attributes returns the attributes for [MssqlJobAgent].
func (mja *MssqlJobAgent) Attributes() mssqlJobAgentAttributes {
	return mssqlJobAgentAttributes{ref: terra.ReferenceResource(mja)}
}

// ImportState imports the given attribute values into [MssqlJobAgent]'s state.
func (mja *MssqlJobAgent) ImportState(av io.Reader) error {
	mja.state = &mssqlJobAgentState{}
	if err := json.NewDecoder(av).Decode(mja.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mja.Type(), mja.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlJobAgent] has state.
func (mja *MssqlJobAgent) State() (*mssqlJobAgentState, bool) {
	return mja.state, mja.state != nil
}

// StateMust returns the state for [MssqlJobAgent]. Panics if the state is nil.
func (mja *MssqlJobAgent) StateMust() *mssqlJobAgentState {
	if mja.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mja.Type(), mja.LocalName()))
	}
	return mja.state
}

// MssqlJobAgentArgs contains the configurations for azurerm_mssql_job_agent.
type MssqlJobAgentArgs struct {
	// DatabaseId: string, required
	DatabaseId terra.StringValue `hcl:"database_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *mssqljobagent.Timeouts `hcl:"timeouts,block"`
}
type mssqlJobAgentAttributes struct {
	ref terra.Reference
}

// DatabaseId returns a reference to field database_id of azurerm_mssql_job_agent.
func (mja mssqlJobAgentAttributes) DatabaseId() terra.StringValue {
	return terra.ReferenceAsString(mja.ref.Append("database_id"))
}

// Id returns a reference to field id of azurerm_mssql_job_agent.
func (mja mssqlJobAgentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mja.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mssql_job_agent.
func (mja mssqlJobAgentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mja.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_mssql_job_agent.
func (mja mssqlJobAgentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mja.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_mssql_job_agent.
func (mja mssqlJobAgentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mja.ref.Append("tags"))
}

func (mja mssqlJobAgentAttributes) Timeouts() mssqljobagent.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqljobagent.TimeoutsAttributes](mja.ref.Append("timeouts"))
}

type mssqlJobAgentState struct {
	DatabaseId string                       `json:"database_id"`
	Id         string                       `json:"id"`
	Location   string                       `json:"location"`
	Name       string                       `json:"name"`
	Tags       map[string]string            `json:"tags"`
	Timeouts   *mssqljobagent.TimeoutsState `json:"timeouts"`
}
