// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewEcrRepositoryPolicy(name string, args EcrRepositoryPolicyArgs) *EcrRepositoryPolicy {
	return &EcrRepositoryPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EcrRepositoryPolicy)(nil)

type EcrRepositoryPolicy struct {
	Name  string
	Args  EcrRepositoryPolicyArgs
	state *ecrRepositoryPolicyState
}

func (erp *EcrRepositoryPolicy) Type() string {
	return "aws_ecr_repository_policy"
}

func (erp *EcrRepositoryPolicy) LocalName() string {
	return erp.Name
}

func (erp *EcrRepositoryPolicy) Configuration() interface{} {
	return erp.Args
}

func (erp *EcrRepositoryPolicy) Attributes() ecrRepositoryPolicyAttributes {
	return ecrRepositoryPolicyAttributes{ref: terra.ReferenceResource(erp)}
}

func (erp *EcrRepositoryPolicy) ImportState(av io.Reader) error {
	erp.state = &ecrRepositoryPolicyState{}
	if err := json.NewDecoder(av).Decode(erp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", erp.Type(), erp.LocalName(), err)
	}
	return nil
}

func (erp *EcrRepositoryPolicy) State() (*ecrRepositoryPolicyState, bool) {
	return erp.state, erp.state != nil
}

func (erp *EcrRepositoryPolicy) StateMust() *ecrRepositoryPolicyState {
	if erp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", erp.Type(), erp.LocalName()))
	}
	return erp.state
}

func (erp *EcrRepositoryPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(erp)
}

type EcrRepositoryPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
	// Repository: string, required
	Repository terra.StringValue `hcl:"repository,attr" validate:"required"`
	// DependsOn contains resources that EcrRepositoryPolicy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type ecrRepositoryPolicyAttributes struct {
	ref terra.Reference
}

func (erp ecrRepositoryPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(erp.ref.Append("id"))
}

func (erp ecrRepositoryPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceString(erp.ref.Append("policy"))
}

func (erp ecrRepositoryPolicyAttributes) RegistryId() terra.StringValue {
	return terra.ReferenceString(erp.ref.Append("registry_id"))
}

func (erp ecrRepositoryPolicyAttributes) Repository() terra.StringValue {
	return terra.ReferenceString(erp.ref.Append("repository"))
}

type ecrRepositoryPolicyState struct {
	Id         string `json:"id"`
	Policy     string `json:"policy"`
	RegistryId string `json:"registry_id"`
	Repository string `json:"repository"`
}
