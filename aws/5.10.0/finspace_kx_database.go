// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	finspacekxdatabase "github.com/golingon/terraproviders/aws/5.10.0/finspacekxdatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFinspaceKxDatabase creates a new instance of [FinspaceKxDatabase].
func NewFinspaceKxDatabase(name string, args FinspaceKxDatabaseArgs) *FinspaceKxDatabase {
	return &FinspaceKxDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FinspaceKxDatabase)(nil)

// FinspaceKxDatabase represents the Terraform resource aws_finspace_kx_database.
type FinspaceKxDatabase struct {
	Name      string
	Args      FinspaceKxDatabaseArgs
	state     *finspaceKxDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FinspaceKxDatabase].
func (fkd *FinspaceKxDatabase) Type() string {
	return "aws_finspace_kx_database"
}

// LocalName returns the local name for [FinspaceKxDatabase].
func (fkd *FinspaceKxDatabase) LocalName() string {
	return fkd.Name
}

// Configuration returns the configuration (args) for [FinspaceKxDatabase].
func (fkd *FinspaceKxDatabase) Configuration() interface{} {
	return fkd.Args
}

// DependOn is used for other resources to depend on [FinspaceKxDatabase].
func (fkd *FinspaceKxDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(fkd)
}

// Dependencies returns the list of resources [FinspaceKxDatabase] depends_on.
func (fkd *FinspaceKxDatabase) Dependencies() terra.Dependencies {
	return fkd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FinspaceKxDatabase].
func (fkd *FinspaceKxDatabase) LifecycleManagement() *terra.Lifecycle {
	return fkd.Lifecycle
}

// Attributes returns the attributes for [FinspaceKxDatabase].
func (fkd *FinspaceKxDatabase) Attributes() finspaceKxDatabaseAttributes {
	return finspaceKxDatabaseAttributes{ref: terra.ReferenceResource(fkd)}
}

// ImportState imports the given attribute values into [FinspaceKxDatabase]'s state.
func (fkd *FinspaceKxDatabase) ImportState(av io.Reader) error {
	fkd.state = &finspaceKxDatabaseState{}
	if err := json.NewDecoder(av).Decode(fkd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fkd.Type(), fkd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FinspaceKxDatabase] has state.
func (fkd *FinspaceKxDatabase) State() (*finspaceKxDatabaseState, bool) {
	return fkd.state, fkd.state != nil
}

// StateMust returns the state for [FinspaceKxDatabase]. Panics if the state is nil.
func (fkd *FinspaceKxDatabase) StateMust() *finspaceKxDatabaseState {
	if fkd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fkd.Type(), fkd.LocalName()))
	}
	return fkd.state
}

// FinspaceKxDatabaseArgs contains the configurations for aws_finspace_kx_database.
type FinspaceKxDatabaseArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EnvironmentId: string, required
	EnvironmentId terra.StringValue `hcl:"environment_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *finspacekxdatabase.Timeouts `hcl:"timeouts,block"`
}
type finspaceKxDatabaseAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_finspace_kx_database.
func (fkd finspaceKxDatabaseAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(fkd.ref.Append("arn"))
}

// CreatedTimestamp returns a reference to field created_timestamp of aws_finspace_kx_database.
func (fkd finspaceKxDatabaseAttributes) CreatedTimestamp() terra.StringValue {
	return terra.ReferenceAsString(fkd.ref.Append("created_timestamp"))
}

// Description returns a reference to field description of aws_finspace_kx_database.
func (fkd finspaceKxDatabaseAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(fkd.ref.Append("description"))
}

// EnvironmentId returns a reference to field environment_id of aws_finspace_kx_database.
func (fkd finspaceKxDatabaseAttributes) EnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(fkd.ref.Append("environment_id"))
}

// Id returns a reference to field id of aws_finspace_kx_database.
func (fkd finspaceKxDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fkd.ref.Append("id"))
}

// LastModifiedTimestamp returns a reference to field last_modified_timestamp of aws_finspace_kx_database.
func (fkd finspaceKxDatabaseAttributes) LastModifiedTimestamp() terra.StringValue {
	return terra.ReferenceAsString(fkd.ref.Append("last_modified_timestamp"))
}

// Name returns a reference to field name of aws_finspace_kx_database.
func (fkd finspaceKxDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fkd.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_finspace_kx_database.
func (fkd finspaceKxDatabaseAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fkd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_finspace_kx_database.
func (fkd finspaceKxDatabaseAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fkd.ref.Append("tags_all"))
}

func (fkd finspaceKxDatabaseAttributes) Timeouts() finspacekxdatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[finspacekxdatabase.TimeoutsAttributes](fkd.ref.Append("timeouts"))
}

type finspaceKxDatabaseState struct {
	Arn                   string                            `json:"arn"`
	CreatedTimestamp      string                            `json:"created_timestamp"`
	Description           string                            `json:"description"`
	EnvironmentId         string                            `json:"environment_id"`
	Id                    string                            `json:"id"`
	LastModifiedTimestamp string                            `json:"last_modified_timestamp"`
	Name                  string                            `json:"name"`
	Tags                  map[string]string                 `json:"tags"`
	TagsAll               map[string]string                 `json:"tags_all"`
	Timeouts              *finspacekxdatabase.TimeoutsState `json:"timeouts"`
}
