// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	redshiftendpointaccess "github.com/golingon/terraproviders/aws/4.60.0/redshiftendpointaccess"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewRedshiftEndpointAccess(name string, args RedshiftEndpointAccessArgs) *RedshiftEndpointAccess {
	return &RedshiftEndpointAccess{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedshiftEndpointAccess)(nil)

type RedshiftEndpointAccess struct {
	Name  string
	Args  RedshiftEndpointAccessArgs
	state *redshiftEndpointAccessState
}

func (rea *RedshiftEndpointAccess) Type() string {
	return "aws_redshift_endpoint_access"
}

func (rea *RedshiftEndpointAccess) LocalName() string {
	return rea.Name
}

func (rea *RedshiftEndpointAccess) Configuration() interface{} {
	return rea.Args
}

func (rea *RedshiftEndpointAccess) Attributes() redshiftEndpointAccessAttributes {
	return redshiftEndpointAccessAttributes{ref: terra.ReferenceResource(rea)}
}

func (rea *RedshiftEndpointAccess) ImportState(av io.Reader) error {
	rea.state = &redshiftEndpointAccessState{}
	if err := json.NewDecoder(av).Decode(rea.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rea.Type(), rea.LocalName(), err)
	}
	return nil
}

func (rea *RedshiftEndpointAccess) State() (*redshiftEndpointAccessState, bool) {
	return rea.state, rea.state != nil
}

func (rea *RedshiftEndpointAccess) StateMust() *redshiftEndpointAccessState {
	if rea.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rea.Type(), rea.LocalName()))
	}
	return rea.state
}

func (rea *RedshiftEndpointAccess) DependOn() terra.Reference {
	return terra.ReferenceResource(rea)
}

type RedshiftEndpointAccessArgs struct {
	// ClusterIdentifier: string, required
	ClusterIdentifier terra.StringValue `hcl:"cluster_identifier,attr" validate:"required"`
	// EndpointName: string, required
	EndpointName terra.StringValue `hcl:"endpoint_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceOwner: string, optional
	ResourceOwner terra.StringValue `hcl:"resource_owner,attr"`
	// SubnetGroupName: string, required
	SubnetGroupName terra.StringValue `hcl:"subnet_group_name,attr" validate:"required"`
	// VpcSecurityGroupIds: set of string, optional
	VpcSecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"vpc_security_group_ids,attr"`
	// VpcEndpoint: min=0
	VpcEndpoint []redshiftendpointaccess.VpcEndpoint `hcl:"vpc_endpoint,block" validate:"min=0"`
	// DependsOn contains resources that RedshiftEndpointAccess depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type redshiftEndpointAccessAttributes struct {
	ref terra.Reference
}

func (rea redshiftEndpointAccessAttributes) Address() terra.StringValue {
	return terra.ReferenceString(rea.ref.Append("address"))
}

func (rea redshiftEndpointAccessAttributes) ClusterIdentifier() terra.StringValue {
	return terra.ReferenceString(rea.ref.Append("cluster_identifier"))
}

func (rea redshiftEndpointAccessAttributes) EndpointName() terra.StringValue {
	return terra.ReferenceString(rea.ref.Append("endpoint_name"))
}

func (rea redshiftEndpointAccessAttributes) Id() terra.StringValue {
	return terra.ReferenceString(rea.ref.Append("id"))
}

func (rea redshiftEndpointAccessAttributes) Port() terra.NumberValue {
	return terra.ReferenceNumber(rea.ref.Append("port"))
}

func (rea redshiftEndpointAccessAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceString(rea.ref.Append("resource_owner"))
}

func (rea redshiftEndpointAccessAttributes) SubnetGroupName() terra.StringValue {
	return terra.ReferenceString(rea.ref.Append("subnet_group_name"))
}

func (rea redshiftEndpointAccessAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](rea.ref.Append("vpc_security_group_ids"))
}

func (rea redshiftEndpointAccessAttributes) VpcEndpoint() terra.ListValue[redshiftendpointaccess.VpcEndpointAttributes] {
	return terra.ReferenceList[redshiftendpointaccess.VpcEndpointAttributes](rea.ref.Append("vpc_endpoint"))
}

type redshiftEndpointAccessState struct {
	Address             string                                    `json:"address"`
	ClusterIdentifier   string                                    `json:"cluster_identifier"`
	EndpointName        string                                    `json:"endpoint_name"`
	Id                  string                                    `json:"id"`
	Port                float64                                   `json:"port"`
	ResourceOwner       string                                    `json:"resource_owner"`
	SubnetGroupName     string                                    `json:"subnet_group_name"`
	VpcSecurityGroupIds []string                                  `json:"vpc_security_group_ids"`
	VpcEndpoint         []redshiftendpointaccess.VpcEndpointState `json:"vpc_endpoint"`
}
