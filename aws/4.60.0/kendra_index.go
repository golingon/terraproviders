// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	kendraindex "github.com/golingon/terraproviders/aws/4.60.0/kendraindex"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKendraIndex creates a new instance of [KendraIndex].
func NewKendraIndex(name string, args KendraIndexArgs) *KendraIndex {
	return &KendraIndex{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KendraIndex)(nil)

// KendraIndex represents the Terraform resource aws_kendra_index.
type KendraIndex struct {
	Name      string
	Args      KendraIndexArgs
	state     *kendraIndexState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KendraIndex].
func (ki *KendraIndex) Type() string {
	return "aws_kendra_index"
}

// LocalName returns the local name for [KendraIndex].
func (ki *KendraIndex) LocalName() string {
	return ki.Name
}

// Configuration returns the configuration (args) for [KendraIndex].
func (ki *KendraIndex) Configuration() interface{} {
	return ki.Args
}

// DependOn is used for other resources to depend on [KendraIndex].
func (ki *KendraIndex) DependOn() terra.Reference {
	return terra.ReferenceResource(ki)
}

// Dependencies returns the list of resources [KendraIndex] depends_on.
func (ki *KendraIndex) Dependencies() terra.Dependencies {
	return ki.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KendraIndex].
func (ki *KendraIndex) LifecycleManagement() *terra.Lifecycle {
	return ki.Lifecycle
}

// Attributes returns the attributes for [KendraIndex].
func (ki *KendraIndex) Attributes() kendraIndexAttributes {
	return kendraIndexAttributes{ref: terra.ReferenceResource(ki)}
}

// ImportState imports the given attribute values into [KendraIndex]'s state.
func (ki *KendraIndex) ImportState(av io.Reader) error {
	ki.state = &kendraIndexState{}
	if err := json.NewDecoder(av).Decode(ki.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ki.Type(), ki.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KendraIndex] has state.
func (ki *KendraIndex) State() (*kendraIndexState, bool) {
	return ki.state, ki.state != nil
}

// StateMust returns the state for [KendraIndex]. Panics if the state is nil.
func (ki *KendraIndex) StateMust() *kendraIndexState {
	if ki.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ki.Type(), ki.LocalName()))
	}
	return ki.state
}

// KendraIndexArgs contains the configurations for aws_kendra_index.
type KendraIndexArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Edition: string, optional
	Edition terra.StringValue `hcl:"edition,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// UserContextPolicy: string, optional
	UserContextPolicy terra.StringValue `hcl:"user_context_policy,attr"`
	// IndexStatistics: min=0
	IndexStatistics []kendraindex.IndexStatistics `hcl:"index_statistics,block" validate:"min=0"`
	// CapacityUnits: optional
	CapacityUnits *kendraindex.CapacityUnits `hcl:"capacity_units,block"`
	// DocumentMetadataConfigurationUpdates: min=0,max=500
	DocumentMetadataConfigurationUpdates []kendraindex.DocumentMetadataConfigurationUpdates `hcl:"document_metadata_configuration_updates,block" validate:"min=0,max=500"`
	// ServerSideEncryptionConfiguration: optional
	ServerSideEncryptionConfiguration *kendraindex.ServerSideEncryptionConfiguration `hcl:"server_side_encryption_configuration,block"`
	// Timeouts: optional
	Timeouts *kendraindex.Timeouts `hcl:"timeouts,block"`
	// UserGroupResolutionConfiguration: optional
	UserGroupResolutionConfiguration *kendraindex.UserGroupResolutionConfiguration `hcl:"user_group_resolution_configuration,block"`
	// UserTokenConfigurations: optional
	UserTokenConfigurations *kendraindex.UserTokenConfigurations `hcl:"user_token_configurations,block"`
}
type kendraIndexAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kendra_index.
func (ki kendraIndexAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("arn"))
}

// CreatedAt returns a reference to field created_at of aws_kendra_index.
func (ki kendraIndexAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("created_at"))
}

// Description returns a reference to field description of aws_kendra_index.
func (ki kendraIndexAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("description"))
}

// Edition returns a reference to field edition of aws_kendra_index.
func (ki kendraIndexAttributes) Edition() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("edition"))
}

// ErrorMessage returns a reference to field error_message of aws_kendra_index.
func (ki kendraIndexAttributes) ErrorMessage() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("error_message"))
}

// Id returns a reference to field id of aws_kendra_index.
func (ki kendraIndexAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("id"))
}

// Name returns a reference to field name of aws_kendra_index.
func (ki kendraIndexAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("name"))
}

// RoleArn returns a reference to field role_arn of aws_kendra_index.
func (ki kendraIndexAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("role_arn"))
}

// Status returns a reference to field status of aws_kendra_index.
func (ki kendraIndexAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_kendra_index.
func (ki kendraIndexAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ki.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_kendra_index.
func (ki kendraIndexAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ki.ref.Append("tags_all"))
}

// UpdatedAt returns a reference to field updated_at of aws_kendra_index.
func (ki kendraIndexAttributes) UpdatedAt() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("updated_at"))
}

// UserContextPolicy returns a reference to field user_context_policy of aws_kendra_index.
func (ki kendraIndexAttributes) UserContextPolicy() terra.StringValue {
	return terra.ReferenceAsString(ki.ref.Append("user_context_policy"))
}

func (ki kendraIndexAttributes) IndexStatistics() terra.ListValue[kendraindex.IndexStatisticsAttributes] {
	return terra.ReferenceAsList[kendraindex.IndexStatisticsAttributes](ki.ref.Append("index_statistics"))
}

func (ki kendraIndexAttributes) CapacityUnits() terra.ListValue[kendraindex.CapacityUnitsAttributes] {
	return terra.ReferenceAsList[kendraindex.CapacityUnitsAttributes](ki.ref.Append("capacity_units"))
}

func (ki kendraIndexAttributes) DocumentMetadataConfigurationUpdates() terra.SetValue[kendraindex.DocumentMetadataConfigurationUpdatesAttributes] {
	return terra.ReferenceAsSet[kendraindex.DocumentMetadataConfigurationUpdatesAttributes](ki.ref.Append("document_metadata_configuration_updates"))
}

func (ki kendraIndexAttributes) ServerSideEncryptionConfiguration() terra.ListValue[kendraindex.ServerSideEncryptionConfigurationAttributes] {
	return terra.ReferenceAsList[kendraindex.ServerSideEncryptionConfigurationAttributes](ki.ref.Append("server_side_encryption_configuration"))
}

func (ki kendraIndexAttributes) Timeouts() kendraindex.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kendraindex.TimeoutsAttributes](ki.ref.Append("timeouts"))
}

func (ki kendraIndexAttributes) UserGroupResolutionConfiguration() terra.ListValue[kendraindex.UserGroupResolutionConfigurationAttributes] {
	return terra.ReferenceAsList[kendraindex.UserGroupResolutionConfigurationAttributes](ki.ref.Append("user_group_resolution_configuration"))
}

func (ki kendraIndexAttributes) UserTokenConfigurations() terra.ListValue[kendraindex.UserTokenConfigurationsAttributes] {
	return terra.ReferenceAsList[kendraindex.UserTokenConfigurationsAttributes](ki.ref.Append("user_token_configurations"))
}

type kendraIndexState struct {
	Arn                                  string                                                  `json:"arn"`
	CreatedAt                            string                                                  `json:"created_at"`
	Description                          string                                                  `json:"description"`
	Edition                              string                                                  `json:"edition"`
	ErrorMessage                         string                                                  `json:"error_message"`
	Id                                   string                                                  `json:"id"`
	Name                                 string                                                  `json:"name"`
	RoleArn                              string                                                  `json:"role_arn"`
	Status                               string                                                  `json:"status"`
	Tags                                 map[string]string                                       `json:"tags"`
	TagsAll                              map[string]string                                       `json:"tags_all"`
	UpdatedAt                            string                                                  `json:"updated_at"`
	UserContextPolicy                    string                                                  `json:"user_context_policy"`
	IndexStatistics                      []kendraindex.IndexStatisticsState                      `json:"index_statistics"`
	CapacityUnits                        []kendraindex.CapacityUnitsState                        `json:"capacity_units"`
	DocumentMetadataConfigurationUpdates []kendraindex.DocumentMetadataConfigurationUpdatesState `json:"document_metadata_configuration_updates"`
	ServerSideEncryptionConfiguration    []kendraindex.ServerSideEncryptionConfigurationState    `json:"server_side_encryption_configuration"`
	Timeouts                             *kendraindex.TimeoutsState                              `json:"timeouts"`
	UserGroupResolutionConfiguration     []kendraindex.UserGroupResolutionConfigurationState     `json:"user_group_resolution_configuration"`
	UserTokenConfigurations              []kendraindex.UserTokenConfigurationsState              `json:"user_token_configurations"`
}
