// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googleworkspace

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewDomainAlias creates a new instance of [DomainAlias].
func NewDomainAlias(name string, args DomainAliasArgs) *DomainAlias {
	return &DomainAlias{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DomainAlias)(nil)

// DomainAlias represents the Terraform resource googleworkspace_domain_alias.
type DomainAlias struct {
	Name      string
	Args      DomainAliasArgs
	state     *domainAliasState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DomainAlias].
func (da *DomainAlias) Type() string {
	return "googleworkspace_domain_alias"
}

// LocalName returns the local name for [DomainAlias].
func (da *DomainAlias) LocalName() string {
	return da.Name
}

// Configuration returns the configuration (args) for [DomainAlias].
func (da *DomainAlias) Configuration() interface{} {
	return da.Args
}

// DependOn is used for other resources to depend on [DomainAlias].
func (da *DomainAlias) DependOn() terra.Reference {
	return terra.ReferenceResource(da)
}

// Dependencies returns the list of resources [DomainAlias] depends_on.
func (da *DomainAlias) Dependencies() terra.Dependencies {
	return da.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DomainAlias].
func (da *DomainAlias) LifecycleManagement() *terra.Lifecycle {
	return da.Lifecycle
}

// Attributes returns the attributes for [DomainAlias].
func (da *DomainAlias) Attributes() domainAliasAttributes {
	return domainAliasAttributes{ref: terra.ReferenceResource(da)}
}

// ImportState imports the given attribute values into [DomainAlias]'s state.
func (da *DomainAlias) ImportState(av io.Reader) error {
	da.state = &domainAliasState{}
	if err := json.NewDecoder(av).Decode(da.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", da.Type(), da.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DomainAlias] has state.
func (da *DomainAlias) State() (*domainAliasState, bool) {
	return da.state, da.state != nil
}

// StateMust returns the state for [DomainAlias]. Panics if the state is nil.
func (da *DomainAlias) StateMust() *domainAliasState {
	if da.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", da.Type(), da.LocalName()))
	}
	return da.state
}

// DomainAliasArgs contains the configurations for googleworkspace_domain_alias.
type DomainAliasArgs struct {
	// DomainAliasName: string, required
	DomainAliasName terra.StringValue `hcl:"domain_alias_name,attr" validate:"required"`
	// ParentDomainName: string, optional
	ParentDomainName terra.StringValue `hcl:"parent_domain_name,attr"`
}
type domainAliasAttributes struct {
	ref terra.Reference
}

// CreationTime returns a reference to field creation_time of googleworkspace_domain_alias.
func (da domainAliasAttributes) CreationTime() terra.NumberValue {
	return terra.ReferenceAsNumber(da.ref.Append("creation_time"))
}

// DomainAliasName returns a reference to field domain_alias_name of googleworkspace_domain_alias.
func (da domainAliasAttributes) DomainAliasName() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("domain_alias_name"))
}

// Etag returns a reference to field etag of googleworkspace_domain_alias.
func (da domainAliasAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("etag"))
}

// Id returns a reference to field id of googleworkspace_domain_alias.
func (da domainAliasAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("id"))
}

// ParentDomainName returns a reference to field parent_domain_name of googleworkspace_domain_alias.
func (da domainAliasAttributes) ParentDomainName() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("parent_domain_name"))
}

// Verified returns a reference to field verified of googleworkspace_domain_alias.
func (da domainAliasAttributes) Verified() terra.BoolValue {
	return terra.ReferenceAsBool(da.ref.Append("verified"))
}

type domainAliasState struct {
	CreationTime     float64 `json:"creation_time"`
	DomainAliasName  string  `json:"domain_alias_name"`
	Etag             string  `json:"etag"`
	Id               string  `json:"id"`
	ParentDomainName string  `json:"parent_domain_name"`
	Verified         bool    `json:"verified"`
}
