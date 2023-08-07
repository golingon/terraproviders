// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqlelasticpool "github.com/golingon/terraproviders/azurerm/3.68.0/mssqlelasticpool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlElasticpool creates a new instance of [MssqlElasticpool].
func NewMssqlElasticpool(name string, args MssqlElasticpoolArgs) *MssqlElasticpool {
	return &MssqlElasticpool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlElasticpool)(nil)

// MssqlElasticpool represents the Terraform resource azurerm_mssql_elasticpool.
type MssqlElasticpool struct {
	Name      string
	Args      MssqlElasticpoolArgs
	state     *mssqlElasticpoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlElasticpool].
func (me *MssqlElasticpool) Type() string {
	return "azurerm_mssql_elasticpool"
}

// LocalName returns the local name for [MssqlElasticpool].
func (me *MssqlElasticpool) LocalName() string {
	return me.Name
}

// Configuration returns the configuration (args) for [MssqlElasticpool].
func (me *MssqlElasticpool) Configuration() interface{} {
	return me.Args
}

// DependOn is used for other resources to depend on [MssqlElasticpool].
func (me *MssqlElasticpool) DependOn() terra.Reference {
	return terra.ReferenceResource(me)
}

// Dependencies returns the list of resources [MssqlElasticpool] depends_on.
func (me *MssqlElasticpool) Dependencies() terra.Dependencies {
	return me.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlElasticpool].
func (me *MssqlElasticpool) LifecycleManagement() *terra.Lifecycle {
	return me.Lifecycle
}

// Attributes returns the attributes for [MssqlElasticpool].
func (me *MssqlElasticpool) Attributes() mssqlElasticpoolAttributes {
	return mssqlElasticpoolAttributes{ref: terra.ReferenceResource(me)}
}

// ImportState imports the given attribute values into [MssqlElasticpool]'s state.
func (me *MssqlElasticpool) ImportState(av io.Reader) error {
	me.state = &mssqlElasticpoolState{}
	if err := json.NewDecoder(av).Decode(me.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", me.Type(), me.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlElasticpool] has state.
func (me *MssqlElasticpool) State() (*mssqlElasticpoolState, bool) {
	return me.state, me.state != nil
}

// StateMust returns the state for [MssqlElasticpool]. Panics if the state is nil.
func (me *MssqlElasticpool) StateMust() *mssqlElasticpoolState {
	if me.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", me.Type(), me.LocalName()))
	}
	return me.state
}

// MssqlElasticpoolArgs contains the configurations for azurerm_mssql_elasticpool.
type MssqlElasticpoolArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LicenseType: string, optional
	LicenseType terra.StringValue `hcl:"license_type,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MaintenanceConfigurationName: string, optional
	MaintenanceConfigurationName terra.StringValue `hcl:"maintenance_configuration_name,attr"`
	// MaxSizeBytes: number, optional
	MaxSizeBytes terra.NumberValue `hcl:"max_size_bytes,attr"`
	// MaxSizeGb: number, optional
	MaxSizeGb terra.NumberValue `hcl:"max_size_gb,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServerName: string, required
	ServerName terra.StringValue `hcl:"server_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ZoneRedundant: bool, optional
	ZoneRedundant terra.BoolValue `hcl:"zone_redundant,attr"`
	// PerDatabaseSettings: required
	PerDatabaseSettings *mssqlelasticpool.PerDatabaseSettings `hcl:"per_database_settings,block" validate:"required"`
	// Sku: required
	Sku *mssqlelasticpool.Sku `hcl:"sku,block" validate:"required"`
	// Timeouts: optional
	Timeouts *mssqlelasticpool.Timeouts `hcl:"timeouts,block"`
}
type mssqlElasticpoolAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mssql_elasticpool.
func (me mssqlElasticpoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("id"))
}

// LicenseType returns a reference to field license_type of azurerm_mssql_elasticpool.
func (me mssqlElasticpoolAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("license_type"))
}

// Location returns a reference to field location of azurerm_mssql_elasticpool.
func (me mssqlElasticpoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("location"))
}

// MaintenanceConfigurationName returns a reference to field maintenance_configuration_name of azurerm_mssql_elasticpool.
func (me mssqlElasticpoolAttributes) MaintenanceConfigurationName() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("maintenance_configuration_name"))
}

// MaxSizeBytes returns a reference to field max_size_bytes of azurerm_mssql_elasticpool.
func (me mssqlElasticpoolAttributes) MaxSizeBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(me.ref.Append("max_size_bytes"))
}

// MaxSizeGb returns a reference to field max_size_gb of azurerm_mssql_elasticpool.
func (me mssqlElasticpoolAttributes) MaxSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(me.ref.Append("max_size_gb"))
}

// Name returns a reference to field name of azurerm_mssql_elasticpool.
func (me mssqlElasticpoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mssql_elasticpool.
func (me mssqlElasticpoolAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_mssql_elasticpool.
func (me mssqlElasticpoolAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("server_name"))
}

// Tags returns a reference to field tags of azurerm_mssql_elasticpool.
func (me mssqlElasticpoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](me.ref.Append("tags"))
}

// ZoneRedundant returns a reference to field zone_redundant of azurerm_mssql_elasticpool.
func (me mssqlElasticpoolAttributes) ZoneRedundant() terra.BoolValue {
	return terra.ReferenceAsBool(me.ref.Append("zone_redundant"))
}

func (me mssqlElasticpoolAttributes) PerDatabaseSettings() terra.ListValue[mssqlelasticpool.PerDatabaseSettingsAttributes] {
	return terra.ReferenceAsList[mssqlelasticpool.PerDatabaseSettingsAttributes](me.ref.Append("per_database_settings"))
}

func (me mssqlElasticpoolAttributes) Sku() terra.ListValue[mssqlelasticpool.SkuAttributes] {
	return terra.ReferenceAsList[mssqlelasticpool.SkuAttributes](me.ref.Append("sku"))
}

func (me mssqlElasticpoolAttributes) Timeouts() mssqlelasticpool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqlelasticpool.TimeoutsAttributes](me.ref.Append("timeouts"))
}

type mssqlElasticpoolState struct {
	Id                           string                                      `json:"id"`
	LicenseType                  string                                      `json:"license_type"`
	Location                     string                                      `json:"location"`
	MaintenanceConfigurationName string                                      `json:"maintenance_configuration_name"`
	MaxSizeBytes                 float64                                     `json:"max_size_bytes"`
	MaxSizeGb                    float64                                     `json:"max_size_gb"`
	Name                         string                                      `json:"name"`
	ResourceGroupName            string                                      `json:"resource_group_name"`
	ServerName                   string                                      `json:"server_name"`
	Tags                         map[string]string                           `json:"tags"`
	ZoneRedundant                bool                                        `json:"zone_redundant"`
	PerDatabaseSettings          []mssqlelasticpool.PerDatabaseSettingsState `json:"per_database_settings"`
	Sku                          []mssqlelasticpool.SkuState                 `json:"sku"`
	Timeouts                     *mssqlelasticpool.TimeoutsState             `json:"timeouts"`
}
