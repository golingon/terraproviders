// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAwsAuthBackendClient creates a new instance of [AwsAuthBackendClient].
func NewAwsAuthBackendClient(name string, args AwsAuthBackendClientArgs) *AwsAuthBackendClient {
	return &AwsAuthBackendClient{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AwsAuthBackendClient)(nil)

// AwsAuthBackendClient represents the Terraform resource vault_aws_auth_backend_client.
type AwsAuthBackendClient struct {
	Name      string
	Args      AwsAuthBackendClientArgs
	state     *awsAuthBackendClientState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AwsAuthBackendClient].
func (aabc *AwsAuthBackendClient) Type() string {
	return "vault_aws_auth_backend_client"
}

// LocalName returns the local name for [AwsAuthBackendClient].
func (aabc *AwsAuthBackendClient) LocalName() string {
	return aabc.Name
}

// Configuration returns the configuration (args) for [AwsAuthBackendClient].
func (aabc *AwsAuthBackendClient) Configuration() interface{} {
	return aabc.Args
}

// DependOn is used for other resources to depend on [AwsAuthBackendClient].
func (aabc *AwsAuthBackendClient) DependOn() terra.Reference {
	return terra.ReferenceResource(aabc)
}

// Dependencies returns the list of resources [AwsAuthBackendClient] depends_on.
func (aabc *AwsAuthBackendClient) Dependencies() terra.Dependencies {
	return aabc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AwsAuthBackendClient].
func (aabc *AwsAuthBackendClient) LifecycleManagement() *terra.Lifecycle {
	return aabc.Lifecycle
}

// Attributes returns the attributes for [AwsAuthBackendClient].
func (aabc *AwsAuthBackendClient) Attributes() awsAuthBackendClientAttributes {
	return awsAuthBackendClientAttributes{ref: terra.ReferenceResource(aabc)}
}

// ImportState imports the given attribute values into [AwsAuthBackendClient]'s state.
func (aabc *AwsAuthBackendClient) ImportState(av io.Reader) error {
	aabc.state = &awsAuthBackendClientState{}
	if err := json.NewDecoder(av).Decode(aabc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aabc.Type(), aabc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AwsAuthBackendClient] has state.
func (aabc *AwsAuthBackendClient) State() (*awsAuthBackendClientState, bool) {
	return aabc.state, aabc.state != nil
}

// StateMust returns the state for [AwsAuthBackendClient]. Panics if the state is nil.
func (aabc *AwsAuthBackendClient) StateMust() *awsAuthBackendClientState {
	if aabc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aabc.Type(), aabc.LocalName()))
	}
	return aabc.state
}

// AwsAuthBackendClientArgs contains the configurations for vault_aws_auth_backend_client.
type AwsAuthBackendClientArgs struct {
	// AccessKey: string, optional
	AccessKey terra.StringValue `hcl:"access_key,attr"`
	// Backend: string, optional
	Backend terra.StringValue `hcl:"backend,attr"`
	// Ec2Endpoint: string, optional
	Ec2Endpoint terra.StringValue `hcl:"ec2_endpoint,attr"`
	// IamEndpoint: string, optional
	IamEndpoint terra.StringValue `hcl:"iam_endpoint,attr"`
	// IamServerIdHeaderValue: string, optional
	IamServerIdHeaderValue terra.StringValue `hcl:"iam_server_id_header_value,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// SecretKey: string, optional
	SecretKey terra.StringValue `hcl:"secret_key,attr"`
	// StsEndpoint: string, optional
	StsEndpoint terra.StringValue `hcl:"sts_endpoint,attr"`
	// StsRegion: string, optional
	StsRegion terra.StringValue `hcl:"sts_region,attr"`
}
type awsAuthBackendClientAttributes struct {
	ref terra.Reference
}

// AccessKey returns a reference to field access_key of vault_aws_auth_backend_client.
func (aabc awsAuthBackendClientAttributes) AccessKey() terra.StringValue {
	return terra.ReferenceAsString(aabc.ref.Append("access_key"))
}

// Backend returns a reference to field backend of vault_aws_auth_backend_client.
func (aabc awsAuthBackendClientAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(aabc.ref.Append("backend"))
}

// Ec2Endpoint returns a reference to field ec2_endpoint of vault_aws_auth_backend_client.
func (aabc awsAuthBackendClientAttributes) Ec2Endpoint() terra.StringValue {
	return terra.ReferenceAsString(aabc.ref.Append("ec2_endpoint"))
}

// IamEndpoint returns a reference to field iam_endpoint of vault_aws_auth_backend_client.
func (aabc awsAuthBackendClientAttributes) IamEndpoint() terra.StringValue {
	return terra.ReferenceAsString(aabc.ref.Append("iam_endpoint"))
}

// IamServerIdHeaderValue returns a reference to field iam_server_id_header_value of vault_aws_auth_backend_client.
func (aabc awsAuthBackendClientAttributes) IamServerIdHeaderValue() terra.StringValue {
	return terra.ReferenceAsString(aabc.ref.Append("iam_server_id_header_value"))
}

// Id returns a reference to field id of vault_aws_auth_backend_client.
func (aabc awsAuthBackendClientAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aabc.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_aws_auth_backend_client.
func (aabc awsAuthBackendClientAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(aabc.ref.Append("namespace"))
}

// SecretKey returns a reference to field secret_key of vault_aws_auth_backend_client.
func (aabc awsAuthBackendClientAttributes) SecretKey() terra.StringValue {
	return terra.ReferenceAsString(aabc.ref.Append("secret_key"))
}

// StsEndpoint returns a reference to field sts_endpoint of vault_aws_auth_backend_client.
func (aabc awsAuthBackendClientAttributes) StsEndpoint() terra.StringValue {
	return terra.ReferenceAsString(aabc.ref.Append("sts_endpoint"))
}

// StsRegion returns a reference to field sts_region of vault_aws_auth_backend_client.
func (aabc awsAuthBackendClientAttributes) StsRegion() terra.StringValue {
	return terra.ReferenceAsString(aabc.ref.Append("sts_region"))
}

type awsAuthBackendClientState struct {
	AccessKey              string `json:"access_key"`
	Backend                string `json:"backend"`
	Ec2Endpoint            string `json:"ec2_endpoint"`
	IamEndpoint            string `json:"iam_endpoint"`
	IamServerIdHeaderValue string `json:"iam_server_id_header_value"`
	Id                     string `json:"id"`
	Namespace              string `json:"namespace"`
	SecretKey              string `json:"secret_key"`
	StsEndpoint            string `json:"sts_endpoint"`
	StsRegion              string `json:"sts_region"`
}
