// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	databasemigrationservice "github.com/golingon/terraproviders/azurerm/3.69.0/databasemigrationservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatabaseMigrationService creates a new instance of [DatabaseMigrationService].
func NewDatabaseMigrationService(name string, args DatabaseMigrationServiceArgs) *DatabaseMigrationService {
	return &DatabaseMigrationService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatabaseMigrationService)(nil)

// DatabaseMigrationService represents the Terraform resource azurerm_database_migration_service.
type DatabaseMigrationService struct {
	Name      string
	Args      DatabaseMigrationServiceArgs
	state     *databaseMigrationServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatabaseMigrationService].
func (dms *DatabaseMigrationService) Type() string {
	return "azurerm_database_migration_service"
}

// LocalName returns the local name for [DatabaseMigrationService].
func (dms *DatabaseMigrationService) LocalName() string {
	return dms.Name
}

// Configuration returns the configuration (args) for [DatabaseMigrationService].
func (dms *DatabaseMigrationService) Configuration() interface{} {
	return dms.Args
}

// DependOn is used for other resources to depend on [DatabaseMigrationService].
func (dms *DatabaseMigrationService) DependOn() terra.Reference {
	return terra.ReferenceResource(dms)
}

// Dependencies returns the list of resources [DatabaseMigrationService] depends_on.
func (dms *DatabaseMigrationService) Dependencies() terra.Dependencies {
	return dms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatabaseMigrationService].
func (dms *DatabaseMigrationService) LifecycleManagement() *terra.Lifecycle {
	return dms.Lifecycle
}

// Attributes returns the attributes for [DatabaseMigrationService].
func (dms *DatabaseMigrationService) Attributes() databaseMigrationServiceAttributes {
	return databaseMigrationServiceAttributes{ref: terra.ReferenceResource(dms)}
}

// ImportState imports the given attribute values into [DatabaseMigrationService]'s state.
func (dms *DatabaseMigrationService) ImportState(av io.Reader) error {
	dms.state = &databaseMigrationServiceState{}
	if err := json.NewDecoder(av).Decode(dms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dms.Type(), dms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatabaseMigrationService] has state.
func (dms *DatabaseMigrationService) State() (*databaseMigrationServiceState, bool) {
	return dms.state, dms.state != nil
}

// StateMust returns the state for [DatabaseMigrationService]. Panics if the state is nil.
func (dms *DatabaseMigrationService) StateMust() *databaseMigrationServiceState {
	if dms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dms.Type(), dms.LocalName()))
	}
	return dms.state
}

// DatabaseMigrationServiceArgs contains the configurations for azurerm_database_migration_service.
type DatabaseMigrationServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *databasemigrationservice.Timeouts `hcl:"timeouts,block"`
}
type databaseMigrationServiceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_database_migration_service.
func (dms databaseMigrationServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_database_migration_service.
func (dms databaseMigrationServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_database_migration_service.
func (dms databaseMigrationServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_database_migration_service.
func (dms databaseMigrationServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_database_migration_service.
func (dms databaseMigrationServiceAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("sku_name"))
}

// SubnetId returns a reference to field subnet_id of azurerm_database_migration_service.
func (dms databaseMigrationServiceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(dms.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_database_migration_service.
func (dms databaseMigrationServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dms.ref.Append("tags"))
}

func (dms databaseMigrationServiceAttributes) Timeouts() databasemigrationservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databasemigrationservice.TimeoutsAttributes](dms.ref.Append("timeouts"))
}

type databaseMigrationServiceState struct {
	Id                string                                  `json:"id"`
	Location          string                                  `json:"location"`
	Name              string                                  `json:"name"`
	ResourceGroupName string                                  `json:"resource_group_name"`
	SkuName           string                                  `json:"sku_name"`
	SubnetId          string                                  `json:"subnet_id"`
	Tags              map[string]string                       `json:"tags"`
	Timeouts          *databasemigrationservice.TimeoutsState `json:"timeouts"`
}