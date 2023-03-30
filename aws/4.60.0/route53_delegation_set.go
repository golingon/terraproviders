// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewRoute53DelegationSet(name string, args Route53DelegationSetArgs) *Route53DelegationSet {
	return &Route53DelegationSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53DelegationSet)(nil)

type Route53DelegationSet struct {
	Name  string
	Args  Route53DelegationSetArgs
	state *route53DelegationSetState
}

func (rds *Route53DelegationSet) Type() string {
	return "aws_route53_delegation_set"
}

func (rds *Route53DelegationSet) LocalName() string {
	return rds.Name
}

func (rds *Route53DelegationSet) Configuration() interface{} {
	return rds.Args
}

func (rds *Route53DelegationSet) Attributes() route53DelegationSetAttributes {
	return route53DelegationSetAttributes{ref: terra.ReferenceResource(rds)}
}

func (rds *Route53DelegationSet) ImportState(av io.Reader) error {
	rds.state = &route53DelegationSetState{}
	if err := json.NewDecoder(av).Decode(rds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rds.Type(), rds.LocalName(), err)
	}
	return nil
}

func (rds *Route53DelegationSet) State() (*route53DelegationSetState, bool) {
	return rds.state, rds.state != nil
}

func (rds *Route53DelegationSet) StateMust() *route53DelegationSetState {
	if rds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rds.Type(), rds.LocalName()))
	}
	return rds.state
}

func (rds *Route53DelegationSet) DependOn() terra.Reference {
	return terra.ReferenceResource(rds)
}

type Route53DelegationSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ReferenceName: string, optional
	ReferenceName terra.StringValue `hcl:"reference_name,attr"`
	// DependsOn contains resources that Route53DelegationSet depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type route53DelegationSetAttributes struct {
	ref terra.Reference
}

func (rds route53DelegationSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(rds.ref.Append("arn"))
}

func (rds route53DelegationSetAttributes) Id() terra.StringValue {
	return terra.ReferenceString(rds.ref.Append("id"))
}

func (rds route53DelegationSetAttributes) NameServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](rds.ref.Append("name_servers"))
}

func (rds route53DelegationSetAttributes) ReferenceName() terra.StringValue {
	return terra.ReferenceString(rds.ref.Append("reference_name"))
}

type route53DelegationSetState struct {
	Arn           string   `json:"arn"`
	Id            string   `json:"id"`
	NameServers   []string `json:"name_servers"`
	ReferenceName string   `json:"reference_name"`
}
