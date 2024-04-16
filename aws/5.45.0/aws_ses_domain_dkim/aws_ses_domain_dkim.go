// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ses_domain_dkim

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

// Resource represents the Terraform resource aws_ses_domain_dkim.
type Resource struct {
	Name      string
	Args      Args
	state     *awsSesDomainDkimState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asdd *Resource) Type() string {
	return "aws_ses_domain_dkim"
}

// LocalName returns the local name for [Resource].
func (asdd *Resource) LocalName() string {
	return asdd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asdd *Resource) Configuration() interface{} {
	return asdd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asdd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asdd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asdd *Resource) Dependencies() terra.Dependencies {
	return asdd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asdd *Resource) LifecycleManagement() *terra.Lifecycle {
	return asdd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asdd *Resource) Attributes() awsSesDomainDkimAttributes {
	return awsSesDomainDkimAttributes{ref: terra.ReferenceResource(asdd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asdd *Resource) ImportState(state io.Reader) error {
	asdd.state = &awsSesDomainDkimState{}
	if err := json.NewDecoder(state).Decode(asdd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asdd.Type(), asdd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asdd *Resource) State() (*awsSesDomainDkimState, bool) {
	return asdd.state, asdd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asdd *Resource) StateMust() *awsSesDomainDkimState {
	if asdd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asdd.Type(), asdd.LocalName()))
	}
	return asdd.state
}

// Args contains the configurations for aws_ses_domain_dkim.
type Args struct {
	// Domain: string, required
	Domain terra.StringValue `hcl:"domain,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}

type awsSesDomainDkimAttributes struct {
	ref terra.Reference
}

// DkimTokens returns a reference to field dkim_tokens of aws_ses_domain_dkim.
func (asdd awsSesDomainDkimAttributes) DkimTokens() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asdd.ref.Append("dkim_tokens"))
}

// Domain returns a reference to field domain of aws_ses_domain_dkim.
func (asdd awsSesDomainDkimAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(asdd.ref.Append("domain"))
}

// Id returns a reference to field id of aws_ses_domain_dkim.
func (asdd awsSesDomainDkimAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asdd.ref.Append("id"))
}

type awsSesDomainDkimState struct {
	DkimTokens []string `json:"dkim_tokens"`
	Domain     string   `json:"domain"`
	Id         string   `json:"id"`
}
