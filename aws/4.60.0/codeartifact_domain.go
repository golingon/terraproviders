// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewCodeartifactDomain(name string, args CodeartifactDomainArgs) *CodeartifactDomain {
	return &CodeartifactDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodeartifactDomain)(nil)

type CodeartifactDomain struct {
	Name  string
	Args  CodeartifactDomainArgs
	state *codeartifactDomainState
}

func (cd *CodeartifactDomain) Type() string {
	return "aws_codeartifact_domain"
}

func (cd *CodeartifactDomain) LocalName() string {
	return cd.Name
}

func (cd *CodeartifactDomain) Configuration() interface{} {
	return cd.Args
}

func (cd *CodeartifactDomain) Attributes() codeartifactDomainAttributes {
	return codeartifactDomainAttributes{ref: terra.ReferenceResource(cd)}
}

func (cd *CodeartifactDomain) ImportState(av io.Reader) error {
	cd.state = &codeartifactDomainState{}
	if err := json.NewDecoder(av).Decode(cd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cd.Type(), cd.LocalName(), err)
	}
	return nil
}

func (cd *CodeartifactDomain) State() (*codeartifactDomainState, bool) {
	return cd.state, cd.state != nil
}

func (cd *CodeartifactDomain) StateMust() *codeartifactDomainState {
	if cd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cd.Type(), cd.LocalName()))
	}
	return cd.state
}

func (cd *CodeartifactDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(cd)
}

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
	// DependsOn contains resources that CodeartifactDomain depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type codeartifactDomainAttributes struct {
	ref terra.Reference
}

func (cd codeartifactDomainAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("arn"))
}

func (cd codeartifactDomainAttributes) AssetSizeBytes() terra.NumberValue {
	return terra.ReferenceNumber(cd.ref.Append("asset_size_bytes"))
}

func (cd codeartifactDomainAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("created_time"))
}

func (cd codeartifactDomainAttributes) Domain() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("domain"))
}

func (cd codeartifactDomainAttributes) EncryptionKey() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("encryption_key"))
}

func (cd codeartifactDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("id"))
}

func (cd codeartifactDomainAttributes) Owner() terra.StringValue {
	return terra.ReferenceString(cd.ref.Append("owner"))
}

func (cd codeartifactDomainAttributes) RepositoryCount() terra.NumberValue {
	return terra.ReferenceNumber(cd.ref.Append("repository_count"))
}

func (cd codeartifactDomainAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cd.ref.Append("tags"))
}

func (cd codeartifactDomainAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cd.ref.Append("tags_all"))
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
