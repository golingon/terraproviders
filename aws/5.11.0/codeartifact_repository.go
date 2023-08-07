// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	codeartifactrepository "github.com/golingon/terraproviders/aws/5.11.0/codeartifactrepository"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCodeartifactRepository creates a new instance of [CodeartifactRepository].
func NewCodeartifactRepository(name string, args CodeartifactRepositoryArgs) *CodeartifactRepository {
	return &CodeartifactRepository{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodeartifactRepository)(nil)

// CodeartifactRepository represents the Terraform resource aws_codeartifact_repository.
type CodeartifactRepository struct {
	Name      string
	Args      CodeartifactRepositoryArgs
	state     *codeartifactRepositoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CodeartifactRepository].
func (cr *CodeartifactRepository) Type() string {
	return "aws_codeartifact_repository"
}

// LocalName returns the local name for [CodeartifactRepository].
func (cr *CodeartifactRepository) LocalName() string {
	return cr.Name
}

// Configuration returns the configuration (args) for [CodeartifactRepository].
func (cr *CodeartifactRepository) Configuration() interface{} {
	return cr.Args
}

// DependOn is used for other resources to depend on [CodeartifactRepository].
func (cr *CodeartifactRepository) DependOn() terra.Reference {
	return terra.ReferenceResource(cr)
}

// Dependencies returns the list of resources [CodeartifactRepository] depends_on.
func (cr *CodeartifactRepository) Dependencies() terra.Dependencies {
	return cr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CodeartifactRepository].
func (cr *CodeartifactRepository) LifecycleManagement() *terra.Lifecycle {
	return cr.Lifecycle
}

// Attributes returns the attributes for [CodeartifactRepository].
func (cr *CodeartifactRepository) Attributes() codeartifactRepositoryAttributes {
	return codeartifactRepositoryAttributes{ref: terra.ReferenceResource(cr)}
}

// ImportState imports the given attribute values into [CodeartifactRepository]'s state.
func (cr *CodeartifactRepository) ImportState(av io.Reader) error {
	cr.state = &codeartifactRepositoryState{}
	if err := json.NewDecoder(av).Decode(cr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cr.Type(), cr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CodeartifactRepository] has state.
func (cr *CodeartifactRepository) State() (*codeartifactRepositoryState, bool) {
	return cr.state, cr.state != nil
}

// StateMust returns the state for [CodeartifactRepository]. Panics if the state is nil.
func (cr *CodeartifactRepository) StateMust() *codeartifactRepositoryState {
	if cr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cr.Type(), cr.LocalName()))
	}
	return cr.state
}

// CodeartifactRepositoryArgs contains the configurations for aws_codeartifact_repository.
type CodeartifactRepositoryArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Domain: string, required
	Domain terra.StringValue `hcl:"domain,attr" validate:"required"`
	// DomainOwner: string, optional
	DomainOwner terra.StringValue `hcl:"domain_owner,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Repository: string, required
	Repository terra.StringValue `hcl:"repository,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ExternalConnections: optional
	ExternalConnections *codeartifactrepository.ExternalConnections `hcl:"external_connections,block"`
	// Upstream: min=0
	Upstream []codeartifactrepository.Upstream `hcl:"upstream,block" validate:"min=0"`
}
type codeartifactRepositoryAttributes struct {
	ref terra.Reference
}

// AdministratorAccount returns a reference to field administrator_account of aws_codeartifact_repository.
func (cr codeartifactRepositoryAttributes) AdministratorAccount() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("administrator_account"))
}

// Arn returns a reference to field arn of aws_codeartifact_repository.
func (cr codeartifactRepositoryAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("arn"))
}

// Description returns a reference to field description of aws_codeartifact_repository.
func (cr codeartifactRepositoryAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("description"))
}

// Domain returns a reference to field domain of aws_codeartifact_repository.
func (cr codeartifactRepositoryAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("domain"))
}

// DomainOwner returns a reference to field domain_owner of aws_codeartifact_repository.
func (cr codeartifactRepositoryAttributes) DomainOwner() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("domain_owner"))
}

// Id returns a reference to field id of aws_codeartifact_repository.
func (cr codeartifactRepositoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("id"))
}

// Repository returns a reference to field repository of aws_codeartifact_repository.
func (cr codeartifactRepositoryAttributes) Repository() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("repository"))
}

// Tags returns a reference to field tags of aws_codeartifact_repository.
func (cr codeartifactRepositoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_codeartifact_repository.
func (cr codeartifactRepositoryAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cr.ref.Append("tags_all"))
}

func (cr codeartifactRepositoryAttributes) ExternalConnections() terra.ListValue[codeartifactrepository.ExternalConnectionsAttributes] {
	return terra.ReferenceAsList[codeartifactrepository.ExternalConnectionsAttributes](cr.ref.Append("external_connections"))
}

func (cr codeartifactRepositoryAttributes) Upstream() terra.ListValue[codeartifactrepository.UpstreamAttributes] {
	return terra.ReferenceAsList[codeartifactrepository.UpstreamAttributes](cr.ref.Append("upstream"))
}

type codeartifactRepositoryState struct {
	AdministratorAccount string                                            `json:"administrator_account"`
	Arn                  string                                            `json:"arn"`
	Description          string                                            `json:"description"`
	Domain               string                                            `json:"domain"`
	DomainOwner          string                                            `json:"domain_owner"`
	Id                   string                                            `json:"id"`
	Repository           string                                            `json:"repository"`
	Tags                 map[string]string                                 `json:"tags"`
	TagsAll              map[string]string                                 `json:"tags_all"`
	ExternalConnections  []codeartifactrepository.ExternalConnectionsState `json:"external_connections"`
	Upstream             []codeartifactrepository.UpstreamState            `json:"upstream"`
}
