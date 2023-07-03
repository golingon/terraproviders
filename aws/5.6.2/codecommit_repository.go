// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCodecommitRepository creates a new instance of [CodecommitRepository].
func NewCodecommitRepository(name string, args CodecommitRepositoryArgs) *CodecommitRepository {
	return &CodecommitRepository{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodecommitRepository)(nil)

// CodecommitRepository represents the Terraform resource aws_codecommit_repository.
type CodecommitRepository struct {
	Name      string
	Args      CodecommitRepositoryArgs
	state     *codecommitRepositoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CodecommitRepository].
func (cr *CodecommitRepository) Type() string {
	return "aws_codecommit_repository"
}

// LocalName returns the local name for [CodecommitRepository].
func (cr *CodecommitRepository) LocalName() string {
	return cr.Name
}

// Configuration returns the configuration (args) for [CodecommitRepository].
func (cr *CodecommitRepository) Configuration() interface{} {
	return cr.Args
}

// DependOn is used for other resources to depend on [CodecommitRepository].
func (cr *CodecommitRepository) DependOn() terra.Reference {
	return terra.ReferenceResource(cr)
}

// Dependencies returns the list of resources [CodecommitRepository] depends_on.
func (cr *CodecommitRepository) Dependencies() terra.Dependencies {
	return cr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CodecommitRepository].
func (cr *CodecommitRepository) LifecycleManagement() *terra.Lifecycle {
	return cr.Lifecycle
}

// Attributes returns the attributes for [CodecommitRepository].
func (cr *CodecommitRepository) Attributes() codecommitRepositoryAttributes {
	return codecommitRepositoryAttributes{ref: terra.ReferenceResource(cr)}
}

// ImportState imports the given attribute values into [CodecommitRepository]'s state.
func (cr *CodecommitRepository) ImportState(av io.Reader) error {
	cr.state = &codecommitRepositoryState{}
	if err := json.NewDecoder(av).Decode(cr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cr.Type(), cr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CodecommitRepository] has state.
func (cr *CodecommitRepository) State() (*codecommitRepositoryState, bool) {
	return cr.state, cr.state != nil
}

// StateMust returns the state for [CodecommitRepository]. Panics if the state is nil.
func (cr *CodecommitRepository) StateMust() *codecommitRepositoryState {
	if cr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cr.Type(), cr.LocalName()))
	}
	return cr.state
}

// CodecommitRepositoryArgs contains the configurations for aws_codecommit_repository.
type CodecommitRepositoryArgs struct {
	// DefaultBranch: string, optional
	DefaultBranch terra.StringValue `hcl:"default_branch,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RepositoryName: string, required
	RepositoryName terra.StringValue `hcl:"repository_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type codecommitRepositoryAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_codecommit_repository.
func (cr codecommitRepositoryAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("arn"))
}

// CloneUrlHttp returns a reference to field clone_url_http of aws_codecommit_repository.
func (cr codecommitRepositoryAttributes) CloneUrlHttp() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("clone_url_http"))
}

// CloneUrlSsh returns a reference to field clone_url_ssh of aws_codecommit_repository.
func (cr codecommitRepositoryAttributes) CloneUrlSsh() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("clone_url_ssh"))
}

// DefaultBranch returns a reference to field default_branch of aws_codecommit_repository.
func (cr codecommitRepositoryAttributes) DefaultBranch() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("default_branch"))
}

// Description returns a reference to field description of aws_codecommit_repository.
func (cr codecommitRepositoryAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("description"))
}

// Id returns a reference to field id of aws_codecommit_repository.
func (cr codecommitRepositoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("id"))
}

// RepositoryId returns a reference to field repository_id of aws_codecommit_repository.
func (cr codecommitRepositoryAttributes) RepositoryId() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("repository_id"))
}

// RepositoryName returns a reference to field repository_name of aws_codecommit_repository.
func (cr codecommitRepositoryAttributes) RepositoryName() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("repository_name"))
}

// Tags returns a reference to field tags of aws_codecommit_repository.
func (cr codecommitRepositoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_codecommit_repository.
func (cr codecommitRepositoryAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cr.ref.Append("tags_all"))
}

type codecommitRepositoryState struct {
	Arn            string            `json:"arn"`
	CloneUrlHttp   string            `json:"clone_url_http"`
	CloneUrlSsh    string            `json:"clone_url_ssh"`
	DefaultBranch  string            `json:"default_branch"`
	Description    string            `json:"description"`
	Id             string            `json:"id"`
	RepositoryId   string            `json:"repository_id"`
	RepositoryName string            `json:"repository_name"`
	Tags           map[string]string `json:"tags"`
	TagsAll        map[string]string `json:"tags_all"`
}