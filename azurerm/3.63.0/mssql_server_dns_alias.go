// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqlserverdnsalias "github.com/golingon/terraproviders/azurerm/3.63.0/mssqlserverdnsalias"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlServerDnsAlias creates a new instance of [MssqlServerDnsAlias].
func NewMssqlServerDnsAlias(name string, args MssqlServerDnsAliasArgs) *MssqlServerDnsAlias {
	return &MssqlServerDnsAlias{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlServerDnsAlias)(nil)

// MssqlServerDnsAlias represents the Terraform resource azurerm_mssql_server_dns_alias.
type MssqlServerDnsAlias struct {
	Name      string
	Args      MssqlServerDnsAliasArgs
	state     *mssqlServerDnsAliasState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlServerDnsAlias].
func (msda *MssqlServerDnsAlias) Type() string {
	return "azurerm_mssql_server_dns_alias"
}

// LocalName returns the local name for [MssqlServerDnsAlias].
func (msda *MssqlServerDnsAlias) LocalName() string {
	return msda.Name
}

// Configuration returns the configuration (args) for [MssqlServerDnsAlias].
func (msda *MssqlServerDnsAlias) Configuration() interface{} {
	return msda.Args
}

// DependOn is used for other resources to depend on [MssqlServerDnsAlias].
func (msda *MssqlServerDnsAlias) DependOn() terra.Reference {
	return terra.ReferenceResource(msda)
}

// Dependencies returns the list of resources [MssqlServerDnsAlias] depends_on.
func (msda *MssqlServerDnsAlias) Dependencies() terra.Dependencies {
	return msda.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlServerDnsAlias].
func (msda *MssqlServerDnsAlias) LifecycleManagement() *terra.Lifecycle {
	return msda.Lifecycle
}

// Attributes returns the attributes for [MssqlServerDnsAlias].
func (msda *MssqlServerDnsAlias) Attributes() mssqlServerDnsAliasAttributes {
	return mssqlServerDnsAliasAttributes{ref: terra.ReferenceResource(msda)}
}

// ImportState imports the given attribute values into [MssqlServerDnsAlias]'s state.
func (msda *MssqlServerDnsAlias) ImportState(av io.Reader) error {
	msda.state = &mssqlServerDnsAliasState{}
	if err := json.NewDecoder(av).Decode(msda.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", msda.Type(), msda.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlServerDnsAlias] has state.
func (msda *MssqlServerDnsAlias) State() (*mssqlServerDnsAliasState, bool) {
	return msda.state, msda.state != nil
}

// StateMust returns the state for [MssqlServerDnsAlias]. Panics if the state is nil.
func (msda *MssqlServerDnsAlias) StateMust() *mssqlServerDnsAliasState {
	if msda.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", msda.Type(), msda.LocalName()))
	}
	return msda.state
}

// MssqlServerDnsAliasArgs contains the configurations for azurerm_mssql_server_dns_alias.
type MssqlServerDnsAliasArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MssqlServerId: string, required
	MssqlServerId terra.StringValue `hcl:"mssql_server_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *mssqlserverdnsalias.Timeouts `hcl:"timeouts,block"`
}
type mssqlServerDnsAliasAttributes struct {
	ref terra.Reference
}

// DnsRecord returns a reference to field dns_record of azurerm_mssql_server_dns_alias.
func (msda mssqlServerDnsAliasAttributes) DnsRecord() terra.StringValue {
	return terra.ReferenceAsString(msda.ref.Append("dns_record"))
}

// Id returns a reference to field id of azurerm_mssql_server_dns_alias.
func (msda mssqlServerDnsAliasAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(msda.ref.Append("id"))
}

// MssqlServerId returns a reference to field mssql_server_id of azurerm_mssql_server_dns_alias.
func (msda mssqlServerDnsAliasAttributes) MssqlServerId() terra.StringValue {
	return terra.ReferenceAsString(msda.ref.Append("mssql_server_id"))
}

// Name returns a reference to field name of azurerm_mssql_server_dns_alias.
func (msda mssqlServerDnsAliasAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(msda.ref.Append("name"))
}

func (msda mssqlServerDnsAliasAttributes) Timeouts() mssqlserverdnsalias.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqlserverdnsalias.TimeoutsAttributes](msda.ref.Append("timeouts"))
}

type mssqlServerDnsAliasState struct {
	DnsRecord     string                             `json:"dns_record"`
	Id            string                             `json:"id"`
	MssqlServerId string                             `json:"mssql_server_id"`
	Name          string                             `json:"name"`
	Timeouts      *mssqlserverdnsalias.TimeoutsState `json:"timeouts"`
}
