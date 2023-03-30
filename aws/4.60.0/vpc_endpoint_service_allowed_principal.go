// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewVpcEndpointServiceAllowedPrincipal(name string, args VpcEndpointServiceAllowedPrincipalArgs) *VpcEndpointServiceAllowedPrincipal {
	return &VpcEndpointServiceAllowedPrincipal{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcEndpointServiceAllowedPrincipal)(nil)

type VpcEndpointServiceAllowedPrincipal struct {
	Name  string
	Args  VpcEndpointServiceAllowedPrincipalArgs
	state *vpcEndpointServiceAllowedPrincipalState
}

func (vesap *VpcEndpointServiceAllowedPrincipal) Type() string {
	return "aws_vpc_endpoint_service_allowed_principal"
}

func (vesap *VpcEndpointServiceAllowedPrincipal) LocalName() string {
	return vesap.Name
}

func (vesap *VpcEndpointServiceAllowedPrincipal) Configuration() interface{} {
	return vesap.Args
}

func (vesap *VpcEndpointServiceAllowedPrincipal) Attributes() vpcEndpointServiceAllowedPrincipalAttributes {
	return vpcEndpointServiceAllowedPrincipalAttributes{ref: terra.ReferenceResource(vesap)}
}

func (vesap *VpcEndpointServiceAllowedPrincipal) ImportState(av io.Reader) error {
	vesap.state = &vpcEndpointServiceAllowedPrincipalState{}
	if err := json.NewDecoder(av).Decode(vesap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vesap.Type(), vesap.LocalName(), err)
	}
	return nil
}

func (vesap *VpcEndpointServiceAllowedPrincipal) State() (*vpcEndpointServiceAllowedPrincipalState, bool) {
	return vesap.state, vesap.state != nil
}

func (vesap *VpcEndpointServiceAllowedPrincipal) StateMust() *vpcEndpointServiceAllowedPrincipalState {
	if vesap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vesap.Type(), vesap.LocalName()))
	}
	return vesap.state
}

func (vesap *VpcEndpointServiceAllowedPrincipal) DependOn() terra.Reference {
	return terra.ReferenceResource(vesap)
}

type VpcEndpointServiceAllowedPrincipalArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PrincipalArn: string, required
	PrincipalArn terra.StringValue `hcl:"principal_arn,attr" validate:"required"`
	// VpcEndpointServiceId: string, required
	VpcEndpointServiceId terra.StringValue `hcl:"vpc_endpoint_service_id,attr" validate:"required"`
	// DependsOn contains resources that VpcEndpointServiceAllowedPrincipal depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type vpcEndpointServiceAllowedPrincipalAttributes struct {
	ref terra.Reference
}

func (vesap vpcEndpointServiceAllowedPrincipalAttributes) Id() terra.StringValue {
	return terra.ReferenceString(vesap.ref.Append("id"))
}

func (vesap vpcEndpointServiceAllowedPrincipalAttributes) PrincipalArn() terra.StringValue {
	return terra.ReferenceString(vesap.ref.Append("principal_arn"))
}

func (vesap vpcEndpointServiceAllowedPrincipalAttributes) VpcEndpointServiceId() terra.StringValue {
	return terra.ReferenceString(vesap.ref.Append("vpc_endpoint_service_id"))
}

type vpcEndpointServiceAllowedPrincipalState struct {
	Id                   string `json:"id"`
	PrincipalArn         string `json:"principal_arn"`
	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`
}
