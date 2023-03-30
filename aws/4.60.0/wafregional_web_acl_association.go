// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewWafregionalWebAclAssociation(name string, args WafregionalWebAclAssociationArgs) *WafregionalWebAclAssociation {
	return &WafregionalWebAclAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafregionalWebAclAssociation)(nil)

type WafregionalWebAclAssociation struct {
	Name  string
	Args  WafregionalWebAclAssociationArgs
	state *wafregionalWebAclAssociationState
}

func (wwaa *WafregionalWebAclAssociation) Type() string {
	return "aws_wafregional_web_acl_association"
}

func (wwaa *WafregionalWebAclAssociation) LocalName() string {
	return wwaa.Name
}

func (wwaa *WafregionalWebAclAssociation) Configuration() interface{} {
	return wwaa.Args
}

func (wwaa *WafregionalWebAclAssociation) Attributes() wafregionalWebAclAssociationAttributes {
	return wafregionalWebAclAssociationAttributes{ref: terra.ReferenceResource(wwaa)}
}

func (wwaa *WafregionalWebAclAssociation) ImportState(av io.Reader) error {
	wwaa.state = &wafregionalWebAclAssociationState{}
	if err := json.NewDecoder(av).Decode(wwaa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wwaa.Type(), wwaa.LocalName(), err)
	}
	return nil
}

func (wwaa *WafregionalWebAclAssociation) State() (*wafregionalWebAclAssociationState, bool) {
	return wwaa.state, wwaa.state != nil
}

func (wwaa *WafregionalWebAclAssociation) StateMust() *wafregionalWebAclAssociationState {
	if wwaa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wwaa.Type(), wwaa.LocalName()))
	}
	return wwaa.state
}

func (wwaa *WafregionalWebAclAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(wwaa)
}

type WafregionalWebAclAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
	// WebAclId: string, required
	WebAclId terra.StringValue `hcl:"web_acl_id,attr" validate:"required"`
	// DependsOn contains resources that WafregionalWebAclAssociation depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type wafregionalWebAclAssociationAttributes struct {
	ref terra.Reference
}

func (wwaa wafregionalWebAclAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(wwaa.ref.Append("id"))
}

func (wwaa wafregionalWebAclAssociationAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceString(wwaa.ref.Append("resource_arn"))
}

func (wwaa wafregionalWebAclAssociationAttributes) WebAclId() terra.StringValue {
	return terra.ReferenceString(wwaa.ref.Append("web_acl_id"))
}

type wafregionalWebAclAssociationState struct {
	Id          string `json:"id"`
	ResourceArn string `json:"resource_arn"`
	WebAclId    string `json:"web_acl_id"`
}
