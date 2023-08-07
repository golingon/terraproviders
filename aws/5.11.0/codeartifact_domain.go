// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCodeartifactDomain creates a new instance of [CodeartifactDomain].
func NewCodeartifactDomain(name string, args CodeartifactDomainArgs) *CodeartifactDomain {
	return &CodeartifactDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodeartifactDomain)(nil)

// CodeartifactDomain represents the Terraform resource aws_codeartifact_domain.
type CodeartifactDomain struct {
	Name      string
	Args      CodeartifactDomainArgs
	state     *codeartifactDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CodeartifactDomain].
func (cd *CodeartifactDomain) Type() string {
	return "aws_codeartifact_domain"
}

// LocalName returns the local name for [CodeartifactDomain].
func (cd *CodeartifactDomain) LocalName() string {
	return cd.Name
}

// Configuration returns the configuration (args) for [CodeartifactDomain].
func (cd *CodeartifactDomain) Configuration() interface{} {
	return cd.Args
}

// DependOn is used for other resources to depend on [CodeartifactDomain].
func (cd *CodeartifactDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(cd)
}

// Dependencies returns the list of resources [CodeartifactDomain] depends_on.
func (cd *CodeartifactDomain) Dependencies() terra.Dependencies {
	return cd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CodeartifactDomain].
func (cd *CodeartifactDomain) LifecycleManagement() *terra.Lifecycle {
	return cd.Lifecycle
}

// Attributes returns the attributes for [CodeartifactDomain].
func (cd *CodeartifactDomain) Attributes() codeartifactDomainAttributes {
	return codeartifactDomainAttributes{ref: terra.ReferenceResource(cd)}
}

// ImportState imports the given attribute values into [CodeartifactDomain]'s state.
func (cd *CodeartifactDomain) ImportState(av io.Reader) error {
	cd.state = &codeartifactDomainState{}
	if err := json.NewDecoder(av).Decode(cd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cd.Type(), cd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CodeartifactDomain] has state.
func (cd *CodeartifactDomain) State() (*codeartifactDomainState, bool) {
	return cd.state, cd.state != nil
}

// StateMust returns the state for [CodeartifactDomain]. Panics if the state is nil.
func (cd *CodeartifactDomain) StateMust() *codeartifactDomainState {
	if cd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cd.Type(), cd.LocalName()))
	}
	return cd.state
}

// CodeartifactDomainArgs contains the configurations for aws_codeartifact_domain.
type CodeartifactDomainArgs struct {
	// Domain: string, required
	Domain terra.StringValue `hcl:"domain,attr" validate:"required"`
	// EncryptionKey: string, optional
	EncryptionKey terra.StringValue `hcl:"encryption_key,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type codeartifactDomainAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_codeartifact_domain.
func (cd codeartifactDomainAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("arn"))
}

// AssetSizeBytes returns a reference to field asset_size_bytes of aws_codeartifact_domain.
func (cd codeartifactDomainAttributes) AssetSizeBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(cd.ref.Append("asset_size_bytes"))
}

// CreatedTime returns a reference to field created_time of aws_codeartifact_domain.
func (cd codeartifactDomainAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("created_time"))
}

// Domain returns a reference to field domain of aws_codeartifact_domain.
func (cd codeartifactDomainAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("domain"))
}

// EncryptionKey returns a reference to field encryption_key of aws_codeartifact_domain.
func (cd codeartifactDomainAttributes) EncryptionKey() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("encryption_key"))
}

// Id returns a reference to field id of aws_codeartifact_domain.
func (cd codeartifactDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("id"))
}

// Owner returns a reference to field owner of aws_codeartifact_domain.
func (cd codeartifactDomainAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("owner"))
}

// RepositoryCount returns a reference to field repository_count of aws_codeartifact_domain.
func (cd codeartifactDomainAttributes) RepositoryCount() terra.NumberValue {
	return terra.ReferenceAsNumber(cd.ref.Append("repository_count"))
}

// Tags returns a reference to field tags of aws_codeartifact_domain.
func (cd codeartifactDomainAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_codeartifact_domain.
func (cd codeartifactDomainAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cd.ref.Append("tags_all"))
}

type codeartifactDomainState struct {
	Arn             string            `json:"arn"`
	AssetSizeBytes  float64           `json:"asset_size_bytes"`
	CreatedTime     string            `json:"created_time"`
	Domain          string            `json:"domain"`
	EncryptionKey   string            `json:"encryption_key"`
	Id              string            `json:"id"`
	Owner           string            `json:"owner"`
	RepositoryCount float64           `json:"repository_count"`
	Tags            map[string]string `json:"tags"`
	TagsAll         map[string]string `json:"tags_all"`
}
