// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewIamAccountAlias(name string, args IamAccountAliasArgs) *IamAccountAlias {
	return &IamAccountAlias{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IamAccountAlias)(nil)

type IamAccountAlias struct {
	Name  string
	Args  IamAccountAliasArgs
	state *iamAccountAliasState
}

func (iaa *IamAccountAlias) Type() string {
	return "aws_iam_account_alias"
}

func (iaa *IamAccountAlias) LocalName() string {
	return iaa.Name
}

func (iaa *IamAccountAlias) Configuration() interface{} {
	return iaa.Args
}

func (iaa *IamAccountAlias) Attributes() iamAccountAliasAttributes {
	return iamAccountAliasAttributes{ref: terra.ReferenceResource(iaa)}
}

func (iaa *IamAccountAlias) ImportState(av io.Reader) error {
	iaa.state = &iamAccountAliasState{}
	if err := json.NewDecoder(av).Decode(iaa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iaa.Type(), iaa.LocalName(), err)
	}
	return nil
}

func (iaa *IamAccountAlias) State() (*iamAccountAliasState, bool) {
	return iaa.state, iaa.state != nil
}

func (iaa *IamAccountAlias) StateMust() *iamAccountAliasState {
	if iaa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iaa.Type(), iaa.LocalName()))
	}
	return iaa.state
}

func (iaa *IamAccountAlias) DependOn() terra.Reference {
	return terra.ReferenceResource(iaa)
}

type IamAccountAliasArgs struct {
	// AccountAlias: string, required
	AccountAlias terra.StringValue `hcl:"account_alias,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// DependsOn contains resources that IamAccountAlias depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type iamAccountAliasAttributes struct {
	ref terra.Reference
}

func (iaa iamAccountAliasAttributes) AccountAlias() terra.StringValue {
	return terra.ReferenceString(iaa.ref.Append("account_alias"))
}

func (iaa iamAccountAliasAttributes) Id() terra.StringValue {
	return terra.ReferenceString(iaa.ref.Append("id"))
}

type iamAccountAliasState struct {
	AccountAlias string `json:"account_alias"`
	Id           string `json:"id"`
}
