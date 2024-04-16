// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_codeartifact_repository

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_codeartifact_repository.
type Resource struct {
	Name      string
	Args      Args
	state     *awsCodeartifactRepositoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acr *Resource) Type() string {
	return "aws_codeartifact_repository"
}

// LocalName returns the local name for [Resource].
func (acr *Resource) LocalName() string {
	return acr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acr *Resource) Configuration() interface{} {
	return acr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acr *Resource) Dependencies() terra.Dependencies {
	return acr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acr *Resource) LifecycleManagement() *terra.Lifecycle {
	return acr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acr *Resource) Attributes() awsCodeartifactRepositoryAttributes {
	return awsCodeartifactRepositoryAttributes{ref: terra.ReferenceResource(acr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acr *Resource) ImportState(state io.Reader) error {
	acr.state = &awsCodeartifactRepositoryState{}
	if err := json.NewDecoder(state).Decode(acr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acr.Type(), acr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acr *Resource) State() (*awsCodeartifactRepositoryState, bool) {
	return acr.state, acr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acr *Resource) StateMust() *awsCodeartifactRepositoryState {
	if acr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acr.Type(), acr.LocalName()))
	}
	return acr.state
}

// Args contains the configurations for aws_codeartifact_repository.
type Args struct {
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
	ExternalConnections *ExternalConnections `hcl:"external_connections,block"`
	// Upstream: min=0
	Upstream []Upstream `hcl:"upstream,block" validate:"min=0"`
}

type awsCodeartifactRepositoryAttributes struct {
	ref terra.Reference
}

// AdministratorAccount returns a reference to field administrator_account of aws_codeartifact_repository.
func (acr awsCodeartifactRepositoryAttributes) AdministratorAccount() terra.StringValue {
	return terra.ReferenceAsString(acr.ref.Append("administrator_account"))
}

// Arn returns a reference to field arn of aws_codeartifact_repository.
func (acr awsCodeartifactRepositoryAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(acr.ref.Append("arn"))
}

// Description returns a reference to field description of aws_codeartifact_repository.
func (acr awsCodeartifactRepositoryAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(acr.ref.Append("description"))
}

// Domain returns a reference to field domain of aws_codeartifact_repository.
func (acr awsCodeartifactRepositoryAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(acr.ref.Append("domain"))
}

// DomainOwner returns a reference to field domain_owner of aws_codeartifact_repository.
func (acr awsCodeartifactRepositoryAttributes) DomainOwner() terra.StringValue {
	return terra.ReferenceAsString(acr.ref.Append("domain_owner"))
}

// Id returns a reference to field id of aws_codeartifact_repository.
func (acr awsCodeartifactRepositoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acr.ref.Append("id"))
}

// Repository returns a reference to field repository of aws_codeartifact_repository.
func (acr awsCodeartifactRepositoryAttributes) Repository() terra.StringValue {
	return terra.ReferenceAsString(acr.ref.Append("repository"))
}

// Tags returns a reference to field tags of aws_codeartifact_repository.
func (acr awsCodeartifactRepositoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_codeartifact_repository.
func (acr awsCodeartifactRepositoryAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acr.ref.Append("tags_all"))
}

func (acr awsCodeartifactRepositoryAttributes) ExternalConnections() terra.ListValue[ExternalConnectionsAttributes] {
	return terra.ReferenceAsList[ExternalConnectionsAttributes](acr.ref.Append("external_connections"))
}

func (acr awsCodeartifactRepositoryAttributes) Upstream() terra.ListValue[UpstreamAttributes] {
	return terra.ReferenceAsList[UpstreamAttributes](acr.ref.Append("upstream"))
}

type awsCodeartifactRepositoryState struct {
	AdministratorAccount string                     `json:"administrator_account"`
	Arn                  string                     `json:"arn"`
	Description          string                     `json:"description"`
	Domain               string                     `json:"domain"`
	DomainOwner          string                     `json:"domain_owner"`
	Id                   string                     `json:"id"`
	Repository           string                     `json:"repository"`
	Tags                 map[string]string          `json:"tags"`
	TagsAll              map[string]string          `json:"tags_all"`
	ExternalConnections  []ExternalConnectionsState `json:"external_connections"`
	Upstream             []UpstreamState            `json:"upstream"`
}
