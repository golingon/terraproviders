// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	synapsesqlpool "github.com/golingon/terraproviders/azurerm/3.49.0/synapsesqlpool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSynapseSqlPool(name string, args SynapseSqlPoolArgs) *SynapseSqlPool {
	return &SynapseSqlPool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SynapseSqlPool)(nil)

type SynapseSqlPool struct {
	Name  string
	Args  SynapseSqlPoolArgs
	state *synapseSqlPoolState
}

func (ssp *SynapseSqlPool) Type() string {
	return "azurerm_synapse_sql_pool"
}

func (ssp *SynapseSqlPool) LocalName() string {
	return ssp.Name
}

func (ssp *SynapseSqlPool) Configuration() interface{} {
	return ssp.Args
}

func (ssp *SynapseSqlPool) Attributes() synapseSqlPoolAttributes {
	return synapseSqlPoolAttributes{ref: terra.ReferenceResource(ssp)}
}

func (ssp *SynapseSqlPool) ImportState(av io.Reader) error {
	ssp.state = &synapseSqlPoolState{}
	if err := json.NewDecoder(av).Decode(ssp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssp.Type(), ssp.LocalName(), err)
	}
	return nil
}

func (ssp *SynapseSqlPool) State() (*synapseSqlPoolState, bool) {
	return ssp.state, ssp.state != nil
}

func (ssp *SynapseSqlPool) StateMust() *synapseSqlPoolState {
	if ssp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssp.Type(), ssp.LocalName()))
	}
	return ssp.state
}

func (ssp *SynapseSqlPool) DependOn() terra.Reference {
	return terra.ReferenceResource(ssp)
}

type SynapseSqlPoolArgs struct {
	// Collation: string, optional
	Collation terra.StringValue `hcl:"collation,attr"`
	// CreateMode: string, optional
	CreateMode terra.StringValue `hcl:"create_mode,attr"`
	// DataEncrypted: bool, optional
	DataEncrypted terra.BoolValue `hcl:"data_encrypted,attr"`
	// GeoBackupPolicyEnabled: bool, optional
	GeoBackupPolicyEnabled terra.BoolValue `hcl:"geo_backup_policy_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryDatabaseId: string, optional
	RecoveryDatabaseId terra.StringValue `hcl:"recovery_database_id,attr"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// SynapseWorkspaceId: string, required
	SynapseWorkspaceId terra.StringValue `hcl:"synapse_workspace_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Restore: optional
	Restore *synapsesqlpool.Restore `hcl:"restore,block"`
	// Timeouts: optional
	Timeouts *synapsesqlpool.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that SynapseSqlPool depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type synapseSqlPoolAttributes struct {
	ref terra.Reference
}

func (ssp synapseSqlPoolAttributes) Collation() terra.StringValue {
	return terra.ReferenceString(ssp.ref.Append("collation"))
}

func (ssp synapseSqlPoolAttributes) CreateMode() terra.StringValue {
	return terra.ReferenceString(ssp.ref.Append("create_mode"))
}

func (ssp synapseSqlPoolAttributes) DataEncrypted() terra.BoolValue {
	return terra.ReferenceBool(ssp.ref.Append("data_encrypted"))
}

func (ssp synapseSqlPoolAttributes) GeoBackupPolicyEnabled() terra.BoolValue {
	return terra.ReferenceBool(ssp.ref.Append("geo_backup_policy_enabled"))
}

func (ssp synapseSqlPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ssp.ref.Append("id"))
}

func (ssp synapseSqlPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ssp.ref.Append("name"))
}

func (ssp synapseSqlPoolAttributes) RecoveryDatabaseId() terra.StringValue {
	return terra.ReferenceString(ssp.ref.Append("recovery_database_id"))
}

func (ssp synapseSqlPoolAttributes) SkuName() terra.StringValue {
	return terra.ReferenceString(ssp.ref.Append("sku_name"))
}

func (ssp synapseSqlPoolAttributes) SynapseWorkspaceId() terra.StringValue {
	return terra.ReferenceString(ssp.ref.Append("synapse_workspace_id"))
}

func (ssp synapseSqlPoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ssp.ref.Append("tags"))
}

func (ssp synapseSqlPoolAttributes) Restore() terra.ListValue[synapsesqlpool.RestoreAttributes] {
	return terra.ReferenceList[synapsesqlpool.RestoreAttributes](ssp.ref.Append("restore"))
}

func (ssp synapseSqlPoolAttributes) Timeouts() synapsesqlpool.TimeoutsAttributes {
	return terra.ReferenceSingle[synapsesqlpool.TimeoutsAttributes](ssp.ref.Append("timeouts"))
}

type synapseSqlPoolState struct {
	Collation              string                        `json:"collation"`
	CreateMode             string                        `json:"create_mode"`
	DataEncrypted          bool                          `json:"data_encrypted"`
	GeoBackupPolicyEnabled bool                          `json:"geo_backup_policy_enabled"`
	Id                     string                        `json:"id"`
	Name                   string                        `json:"name"`
	RecoveryDatabaseId     string                        `json:"recovery_database_id"`
	SkuName                string                        `json:"sku_name"`
	SynapseWorkspaceId     string                        `json:"synapse_workspace_id"`
	Tags                   map[string]string             `json:"tags"`
	Restore                []synapsesqlpool.RestoreState `json:"restore"`
	Timeouts               *synapsesqlpool.TimeoutsState `json:"timeouts"`
}
