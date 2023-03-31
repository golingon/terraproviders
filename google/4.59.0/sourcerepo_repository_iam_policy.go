// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSourcerepoRepositoryIamPolicy(name string, args SourcerepoRepositoryIamPolicyArgs) *SourcerepoRepositoryIamPolicy {
	return &SourcerepoRepositoryIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SourcerepoRepositoryIamPolicy)(nil)

type SourcerepoRepositoryIamPolicy struct {
	Name  string
	Args  SourcerepoRepositoryIamPolicyArgs
	state *sourcerepoRepositoryIamPolicyState
}

func (srip *SourcerepoRepositoryIamPolicy) Type() string {
	return "google_sourcerepo_repository_iam_policy"
}

func (srip *SourcerepoRepositoryIamPolicy) LocalName() string {
	return srip.Name
}

func (srip *SourcerepoRepositoryIamPolicy) Configuration() interface{} {
	return srip.Args
}

func (srip *SourcerepoRepositoryIamPolicy) Attributes() sourcerepoRepositoryIamPolicyAttributes {
	return sourcerepoRepositoryIamPolicyAttributes{ref: terra.ReferenceResource(srip)}
}

func (srip *SourcerepoRepositoryIamPolicy) ImportState(av io.Reader) error {
	srip.state = &sourcerepoRepositoryIamPolicyState{}
	if err := json.NewDecoder(av).Decode(srip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srip.Type(), srip.LocalName(), err)
	}
	return nil
}

func (srip *SourcerepoRepositoryIamPolicy) State() (*sourcerepoRepositoryIamPolicyState, bool) {
	return srip.state, srip.state != nil
}

func (srip *SourcerepoRepositoryIamPolicy) StateMust() *sourcerepoRepositoryIamPolicyState {
	if srip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srip.Type(), srip.LocalName()))
	}
	return srip.state
}

func (srip *SourcerepoRepositoryIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(srip)
}

type SourcerepoRepositoryIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Repository: string, required
	Repository terra.StringValue `hcl:"repository,attr" validate:"required"`
	// DependsOn contains resources that SourcerepoRepositoryIamPolicy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type sourcerepoRepositoryIamPolicyAttributes struct {
	ref terra.Reference
}

func (srip sourcerepoRepositoryIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(srip.ref.Append("etag"))
}

func (srip sourcerepoRepositoryIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(srip.ref.Append("id"))
}

func (srip sourcerepoRepositoryIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceString(srip.ref.Append("policy_data"))
}

func (srip sourcerepoRepositoryIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceString(srip.ref.Append("project"))
}

func (srip sourcerepoRepositoryIamPolicyAttributes) Repository() terra.StringValue {
	return terra.ReferenceString(srip.ref.Append("repository"))
}

type sourcerepoRepositoryIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	Repository string `json:"repository"`
}
