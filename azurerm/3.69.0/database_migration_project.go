// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	databasemigrationproject "github.com/golingon/terraproviders/azurerm/3.69.0/databasemigrationproject"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatabaseMigrationProject creates a new instance of [DatabaseMigrationProject].
func NewDatabaseMigrationProject(name string, args DatabaseMigrationProjectArgs) *DatabaseMigrationProject {
	return &DatabaseMigrationProject{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatabaseMigrationProject)(nil)

// DatabaseMigrationProject represents the Terraform resource azurerm_database_migration_project.
type DatabaseMigrationProject struct {
	Name      string
	Args      DatabaseMigrationProjectArgs
	state     *databaseMigrationProjectState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatabaseMigrationProject].
func (dmp *DatabaseMigrationProject) Type() string {
	return "azurerm_database_migration_project"
}

// LocalName returns the local name for [DatabaseMigrationProject].
func (dmp *DatabaseMigrationProject) LocalName() string {
	return dmp.Name
}

// Configuration returns the configuration (args) for [DatabaseMigrationProject].
func (dmp *DatabaseMigrationProject) Configuration() interface{} {
	return dmp.Args
}

// DependOn is used for other resources to depend on [DatabaseMigrationProject].
func (dmp *DatabaseMigrationProject) DependOn() terra.Reference {
	return terra.ReferenceResource(dmp)
}

// Dependencies returns the list of resources [DatabaseMigrationProject] depends_on.
func (dmp *DatabaseMigrationProject) Dependencies() terra.Dependencies {
	return dmp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatabaseMigrationProject].
func (dmp *DatabaseMigrationProject) LifecycleManagement() *terra.Lifecycle {
	return dmp.Lifecycle
}

// Attributes returns the attributes for [DatabaseMigrationProject].
func (dmp *DatabaseMigrationProject) Attributes() databaseMigrationProjectAttributes {
	return databaseMigrationProjectAttributes{ref: terra.ReferenceResource(dmp)}
}

// ImportState imports the given attribute values into [DatabaseMigrationProject]'s state.
func (dmp *DatabaseMigrationProject) ImportState(av io.Reader) error {
	dmp.state = &databaseMigrationProjectState{}
	if err := json.NewDecoder(av).Decode(dmp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dmp.Type(), dmp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatabaseMigrationProject] has state.
func (dmp *DatabaseMigrationProject) State() (*databaseMigrationProjectState, bool) {
	return dmp.state, dmp.state != nil
}

// StateMust returns the state for [DatabaseMigrationProject]. Panics if the state is nil.
func (dmp *DatabaseMigrationProject) StateMust() *databaseMigrationProjectState {
	if dmp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dmp.Type(), dmp.LocalName()))
	}
	return dmp.state
}

// DatabaseMigrationProjectArgs contains the configurations for azurerm_database_migration_project.
type DatabaseMigrationProjectArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// SourcePlatform: string, required
	SourcePlatform terra.StringValue `hcl:"source_platform,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TargetPlatform: string, required
	TargetPlatform terra.StringValue `hcl:"target_platform,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *databasemigrationproject.Timeouts `hcl:"timeouts,block"`
}
type databaseMigrationProjectAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_database_migration_project.
func (dmp databaseMigrationProjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_database_migration_project.
func (dmp databaseMigrationProjectAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_database_migration_project.
func (dmp databaseMigrationProjectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_database_migration_project.
func (dmp databaseMigrationProjectAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("resource_group_name"))
}

// ServiceName returns a reference to field service_name of azurerm_database_migration_project.
func (dmp databaseMigrationProjectAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("service_name"))
}

// SourcePlatform returns a reference to field source_platform of azurerm_database_migration_project.
func (dmp databaseMigrationProjectAttributes) SourcePlatform() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("source_platform"))
}

// Tags returns a reference to field tags of azurerm_database_migration_project.
func (dmp databaseMigrationProjectAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dmp.ref.Append("tags"))
}

// TargetPlatform returns a reference to field target_platform of azurerm_database_migration_project.
func (dmp databaseMigrationProjectAttributes) TargetPlatform() terra.StringValue {
	return terra.ReferenceAsString(dmp.ref.Append("target_platform"))
}

func (dmp databaseMigrationProjectAttributes) Timeouts() databasemigrationproject.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databasemigrationproject.TimeoutsAttributes](dmp.ref.Append("timeouts"))
}

type databaseMigrationProjectState struct {
	Id                string                                  `json:"id"`
	Location          string                                  `json:"location"`
	Name              string                                  `json:"name"`
	ResourceGroupName string                                  `json:"resource_group_name"`
	ServiceName       string                                  `json:"service_name"`
	SourcePlatform    string                                  `json:"source_platform"`
	Tags              map[string]string                       `json:"tags"`
	TargetPlatform    string                                  `json:"target_platform"`
	Timeouts          *databasemigrationproject.TimeoutsState `json:"timeouts"`
}
