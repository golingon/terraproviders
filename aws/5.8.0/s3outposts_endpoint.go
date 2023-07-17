// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	s3outpostsendpoint "github.com/golingon/terraproviders/aws/5.8.0/s3outpostsendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3OutpostsEndpoint creates a new instance of [S3OutpostsEndpoint].
func NewS3OutpostsEndpoint(name string, args S3OutpostsEndpointArgs) *S3OutpostsEndpoint {
	return &S3OutpostsEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3OutpostsEndpoint)(nil)

// S3OutpostsEndpoint represents the Terraform resource aws_s3outposts_endpoint.
type S3OutpostsEndpoint struct {
	Name      string
	Args      S3OutpostsEndpointArgs
	state     *s3OutpostsEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3OutpostsEndpoint].
func (se *S3OutpostsEndpoint) Type() string {
	return "aws_s3outposts_endpoint"
}

// LocalName returns the local name for [S3OutpostsEndpoint].
func (se *S3OutpostsEndpoint) LocalName() string {
	return se.Name
}

// Configuration returns the configuration (args) for [S3OutpostsEndpoint].
func (se *S3OutpostsEndpoint) Configuration() interface{} {
	return se.Args
}

// DependOn is used for other resources to depend on [S3OutpostsEndpoint].
func (se *S3OutpostsEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(se)
}

// Dependencies returns the list of resources [S3OutpostsEndpoint] depends_on.
func (se *S3OutpostsEndpoint) Dependencies() terra.Dependencies {
	return se.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3OutpostsEndpoint].
func (se *S3OutpostsEndpoint) LifecycleManagement() *terra.Lifecycle {
	return se.Lifecycle
}

// Attributes returns the attributes for [S3OutpostsEndpoint].
func (se *S3OutpostsEndpoint) Attributes() s3OutpostsEndpointAttributes {
	return s3OutpostsEndpointAttributes{ref: terra.ReferenceResource(se)}
}

// ImportState imports the given attribute values into [S3OutpostsEndpoint]'s state.
func (se *S3OutpostsEndpoint) ImportState(av io.Reader) error {
	se.state = &s3OutpostsEndpointState{}
	if err := json.NewDecoder(av).Decode(se.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", se.Type(), se.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3OutpostsEndpoint] has state.
func (se *S3OutpostsEndpoint) State() (*s3OutpostsEndpointState, bool) {
	return se.state, se.state != nil
}

// StateMust returns the state for [S3OutpostsEndpoint]. Panics if the state is nil.
func (se *S3OutpostsEndpoint) StateMust() *s3OutpostsEndpointState {
	if se.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", se.Type(), se.LocalName()))
	}
	return se.state
}

// S3OutpostsEndpointArgs contains the configurations for aws_s3outposts_endpoint.
type S3OutpostsEndpointArgs struct {
	// AccessType: string, optional
	AccessType terra.StringValue `hcl:"access_type,attr"`
	// CustomerOwnedIpv4Pool: string, optional
	CustomerOwnedIpv4Pool terra.StringValue `hcl:"customer_owned_ipv4_pool,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OutpostId: string, required
	OutpostId terra.StringValue `hcl:"outpost_id,attr" validate:"required"`
	// SecurityGroupId: string, required
	SecurityGroupId terra.StringValue `hcl:"security_group_id,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// NetworkInterfaces: min=0
	NetworkInterfaces []s3outpostsendpoint.NetworkInterfaces `hcl:"network_interfaces,block" validate:"min=0"`
}
type s3OutpostsEndpointAttributes struct {
	ref terra.Reference
}

// AccessType returns a reference to field access_type of aws_s3outposts_endpoint.
func (se s3OutpostsEndpointAttributes) AccessType() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("access_type"))
}

// Arn returns a reference to field arn of aws_s3outposts_endpoint.
func (se s3OutpostsEndpointAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("arn"))
}

// CidrBlock returns a reference to field cidr_block of aws_s3outposts_endpoint.
func (se s3OutpostsEndpointAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("cidr_block"))
}

// CreationTime returns a reference to field creation_time of aws_s3outposts_endpoint.
func (se s3OutpostsEndpointAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("creation_time"))
}

// CustomerOwnedIpv4Pool returns a reference to field customer_owned_ipv4_pool of aws_s3outposts_endpoint.
func (se s3OutpostsEndpointAttributes) CustomerOwnedIpv4Pool() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("customer_owned_ipv4_pool"))
}

// Id returns a reference to field id of aws_s3outposts_endpoint.
func (se s3OutpostsEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("id"))
}

// OutpostId returns a reference to field outpost_id of aws_s3outposts_endpoint.
func (se s3OutpostsEndpointAttributes) OutpostId() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("outpost_id"))
}

// SecurityGroupId returns a reference to field security_group_id of aws_s3outposts_endpoint.
func (se s3OutpostsEndpointAttributes) SecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("security_group_id"))
}

// SubnetId returns a reference to field subnet_id of aws_s3outposts_endpoint.
func (se s3OutpostsEndpointAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("subnet_id"))
}

func (se s3OutpostsEndpointAttributes) NetworkInterfaces() terra.SetValue[s3outpostsendpoint.NetworkInterfacesAttributes] {
	return terra.ReferenceAsSet[s3outpostsendpoint.NetworkInterfacesAttributes](se.ref.Append("network_interfaces"))
}

type s3OutpostsEndpointState struct {
	AccessType            string                                      `json:"access_type"`
	Arn                   string                                      `json:"arn"`
	CidrBlock             string                                      `json:"cidr_block"`
	CreationTime          string                                      `json:"creation_time"`
	CustomerOwnedIpv4Pool string                                      `json:"customer_owned_ipv4_pool"`
	Id                    string                                      `json:"id"`
	OutpostId             string                                      `json:"outpost_id"`
	SecurityGroupId       string                                      `json:"security_group_id"`
	SubnetId              string                                      `json:"subnet_id"`
	NetworkInterfaces     []s3outpostsendpoint.NetworkInterfacesState `json:"network_interfaces"`
}
