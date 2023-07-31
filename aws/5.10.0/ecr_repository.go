// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ecrrepository "github.com/golingon/terraproviders/aws/5.10.0/ecrrepository"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEcrRepository creates a new instance of [EcrRepository].
func NewEcrRepository(name string, args EcrRepositoryArgs) *EcrRepository {
	return &EcrRepository{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EcrRepository)(nil)

// EcrRepository represents the Terraform resource aws_ecr_repository.
type EcrRepository struct {
	Name      string
	Args      EcrRepositoryArgs
	state     *ecrRepositoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EcrRepository].
func (er *EcrRepository) Type() string {
	return "aws_ecr_repository"
}

// LocalName returns the local name for [EcrRepository].
func (er *EcrRepository) LocalName() string {
	return er.Name
}

// Configuration returns the configuration (args) for [EcrRepository].
func (er *EcrRepository) Configuration() interface{} {
	return er.Args
}

// DependOn is used for other resources to depend on [EcrRepository].
func (er *EcrRepository) DependOn() terra.Reference {
	return terra.ReferenceResource(er)
}

// Dependencies returns the list of resources [EcrRepository] depends_on.
func (er *EcrRepository) Dependencies() terra.Dependencies {
	return er.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EcrRepository].
func (er *EcrRepository) LifecycleManagement() *terra.Lifecycle {
	return er.Lifecycle
}

// Attributes returns the attributes for [EcrRepository].
func (er *EcrRepository) Attributes() ecrRepositoryAttributes {
	return ecrRepositoryAttributes{ref: terra.ReferenceResource(er)}
}

// ImportState imports the given attribute values into [EcrRepository]'s state.
func (er *EcrRepository) ImportState(av io.Reader) error {
	er.state = &ecrRepositoryState{}
	if err := json.NewDecoder(av).Decode(er.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", er.Type(), er.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EcrRepository] has state.
func (er *EcrRepository) State() (*ecrRepositoryState, bool) {
	return er.state, er.state != nil
}

// StateMust returns the state for [EcrRepository]. Panics if the state is nil.
func (er *EcrRepository) StateMust() *ecrRepositoryState {
	if er.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", er.Type(), er.LocalName()))
	}
	return er.state
}

// EcrRepositoryArgs contains the configurations for aws_ecr_repository.
type EcrRepositoryArgs struct {
	// ForceDelete: bool, optional
	ForceDelete terra.BoolValue `hcl:"force_delete,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImageTagMutability: string, optional
	ImageTagMutability terra.StringValue `hcl:"image_tag_mutability,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// EncryptionConfiguration: min=0
	EncryptionConfiguration []ecrrepository.EncryptionConfiguration `hcl:"encryption_configuration,block" validate:"min=0"`
	// ImageScanningConfiguration: optional
	ImageScanningConfiguration *ecrrepository.ImageScanningConfiguration `hcl:"image_scanning_configuration,block"`
	// Timeouts: optional
	Timeouts *ecrrepository.Timeouts `hcl:"timeouts,block"`
}
type ecrRepositoryAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ecr_repository.
func (er ecrRepositoryAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("arn"))
}

// ForceDelete returns a reference to field force_delete of aws_ecr_repository.
func (er ecrRepositoryAttributes) ForceDelete() terra.BoolValue {
	return terra.ReferenceAsBool(er.ref.Append("force_delete"))
}

// Id returns a reference to field id of aws_ecr_repository.
func (er ecrRepositoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("id"))
}

// ImageTagMutability returns a reference to field image_tag_mutability of aws_ecr_repository.
func (er ecrRepositoryAttributes) ImageTagMutability() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("image_tag_mutability"))
}

// Name returns a reference to field name of aws_ecr_repository.
func (er ecrRepositoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("name"))
}

// RegistryId returns a reference to field registry_id of aws_ecr_repository.
func (er ecrRepositoryAttributes) RegistryId() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("registry_id"))
}

// RepositoryUrl returns a reference to field repository_url of aws_ecr_repository.
func (er ecrRepositoryAttributes) RepositoryUrl() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("repository_url"))
}

// Tags returns a reference to field tags of aws_ecr_repository.
func (er ecrRepositoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](er.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ecr_repository.
func (er ecrRepositoryAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](er.ref.Append("tags_all"))
}

func (er ecrRepositoryAttributes) EncryptionConfiguration() terra.ListValue[ecrrepository.EncryptionConfigurationAttributes] {
	return terra.ReferenceAsList[ecrrepository.EncryptionConfigurationAttributes](er.ref.Append("encryption_configuration"))
}

func (er ecrRepositoryAttributes) ImageScanningConfiguration() terra.ListValue[ecrrepository.ImageScanningConfigurationAttributes] {
	return terra.ReferenceAsList[ecrrepository.ImageScanningConfigurationAttributes](er.ref.Append("image_scanning_configuration"))
}

func (er ecrRepositoryAttributes) Timeouts() ecrrepository.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ecrrepository.TimeoutsAttributes](er.ref.Append("timeouts"))
}

type ecrRepositoryState struct {
	Arn                        string                                          `json:"arn"`
	ForceDelete                bool                                            `json:"force_delete"`
	Id                         string                                          `json:"id"`
	ImageTagMutability         string                                          `json:"image_tag_mutability"`
	Name                       string                                          `json:"name"`
	RegistryId                 string                                          `json:"registry_id"`
	RepositoryUrl              string                                          `json:"repository_url"`
	Tags                       map[string]string                               `json:"tags"`
	TagsAll                    map[string]string                               `json:"tags_all"`
	EncryptionConfiguration    []ecrrepository.EncryptionConfigurationState    `json:"encryption_configuration"`
	ImageScanningConfiguration []ecrrepository.ImageScanningConfigurationState `json:"image_scanning_configuration"`
	Timeouts                   *ecrrepository.TimeoutsState                    `json:"timeouts"`
}
