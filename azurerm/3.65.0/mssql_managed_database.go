// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqlmanageddatabase "github.com/golingon/terraproviders/azurerm/3.65.0/mssqlmanageddatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlManagedDatabase creates a new instance of [MssqlManagedDatabase].
func NewMssqlManagedDatabase(name string, args MssqlManagedDatabaseArgs) *MssqlManagedDatabase {
	return &MssqlManagedDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlManagedDatabase)(nil)

// MssqlManagedDatabase represents the Terraform resource azurerm_mssql_managed_database.
type MssqlManagedDatabase struct {
	Name      string
	Args      MssqlManagedDatabaseArgs
	state     *mssqlManagedDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlManagedDatabase].
func (mmd *MssqlManagedDatabase) Type() string {
	return "azurerm_mssql_managed_database"
}

// LocalName returns the local name for [MssqlManagedDatabase].
func (mmd *MssqlManagedDatabase) LocalName() string {
	return mmd.Name
}

// Configuration returns the configuration (args) for [MssqlManagedDatabase].
func (mmd *MssqlManagedDatabase) Configuration() interface{} {
	return mmd.Args
}

// DependOn is used for other resources to depend on [MssqlManagedDatabase].
func (mmd *MssqlManagedDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(mmd)
}

// Dependencies returns the list of resources [MssqlManagedDatabase] depends_on.
func (mmd *MssqlManagedDatabase) Dependencies() terra.Dependencies {
	return mmd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlManagedDatabase].
func (mmd *MssqlManagedDatabase) LifecycleManagement() *terra.Lifecycle {
	return mmd.Lifecycle
}

// Attributes returns the attributes for [MssqlManagedDatabase].
func (mmd *MssqlManagedDatabase) Attributes() mssqlManagedDatabaseAttributes {
	return mssqlManagedDatabaseAttributes{ref: terra.ReferenceResource(mmd)}
}

// ImportState imports the given attribute values into [MssqlManagedDatabase]'s state.
func (mmd *MssqlManagedDatabase) ImportState(av io.Reader) error {
	mmd.state = &mssqlManagedDatabaseState{}
	if err := json.NewDecoder(av).Decode(mmd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mmd.Type(), mmd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlManagedDatabase] has state.
func (mmd *MssqlManagedDatabase) State() (*mssqlManagedDatabaseState, bool) {
	return mmd.state, mmd.state != nil
}

// StateMust returns the state for [MssqlManagedDatabase]. Panics if the state is nil.
func (mmd *MssqlManagedDatabase) StateMust() *mssqlManagedDatabaseState {
	if mmd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mmd.Type(), mmd.LocalName()))
	}
	return mmd.state
}

// MssqlManagedDatabaseArgs contains the configurations for azurerm_mssql_managed_database.
type MssqlManagedDatabaseArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagedInstanceId: string, required
	ManagedInstanceId terra.StringValue `hcl:"managed_instance_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ShortTermRetentionDays: number, optional
	ShortTermRetentionDays terra.NumberValue `hcl:"short_term_retention_days,attr"`
	// LongTermRetentionPolicy: optional
	LongTermRetentionPolicy *mssqlmanageddatabase.LongTermRetentionPolicy `hcl:"long_term_retention_policy,block"`
	// Timeouts: optional
	Timeouts *mssqlmanageddatabase.Timeouts `hcl:"timeouts,block"`
}
type mssqlManagedDatabaseAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mssql_managed_database.
func (mmd mssqlManagedDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mmd.ref.Append("id"))
}

// ManagedInstanceId returns a reference to field managed_instance_id of azurerm_mssql_managed_database.
func (mmd mssqlManagedDatabaseAttributes) ManagedInstanceId() terra.StringValue {
	return terra.ReferenceAsString(mmd.ref.Append("managed_instance_id"))
}

// Name returns a reference to field name of azurerm_mssql_managed_database.
func (mmd mssqlManagedDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mmd.ref.Append("name"))
}

// ShortTermRetentionDays returns a reference to field short_term_retention_days of azurerm_mssql_managed_database.
func (mmd mssqlManagedDatabaseAttributes) ShortTermRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(mmd.ref.Append("short_term_retention_days"))
}

func (mmd mssqlManagedDatabaseAttributes) LongTermRetentionPolicy() terra.ListValue[mssqlmanageddatabase.LongTermRetentionPolicyAttributes] {
	return terra.ReferenceAsList[mssqlmanageddatabase.LongTermRetentionPolicyAttributes](mmd.ref.Append("long_term_retention_policy"))
}

func (mmd mssqlManagedDatabaseAttributes) Timeouts() mssqlmanageddatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqlmanageddatabase.TimeoutsAttributes](mmd.ref.Append("timeouts"))
}

type mssqlManagedDatabaseState struct {
	Id                      string                                              `json:"id"`
	ManagedInstanceId       string                                              `json:"managed_instance_id"`
	Name                    string                                              `json:"name"`
	ShortTermRetentionDays  float64                                             `json:"short_term_retention_days"`
	LongTermRetentionPolicy []mssqlmanageddatabase.LongTermRetentionPolicyState `json:"long_term_retention_policy"`
	Timeouts                *mssqlmanageddatabase.TimeoutsState                 `json:"timeouts"`
}
