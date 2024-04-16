// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_api_gateway_client_certificate

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_api_gateway_client_certificate.
type Resource struct {
	Name      string
	Args      Args
	state     *awsApiGatewayClientCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aagcc *Resource) Type() string {
	return "aws_api_gateway_client_certificate"
}

// LocalName returns the local name for [Resource].
func (aagcc *Resource) LocalName() string {
	return aagcc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aagcc *Resource) Configuration() interface{} {
	return aagcc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aagcc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aagcc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aagcc *Resource) Dependencies() terra.Dependencies {
	return aagcc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aagcc *Resource) LifecycleManagement() *terra.Lifecycle {
	return aagcc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aagcc *Resource) Attributes() awsApiGatewayClientCertificateAttributes {
	return awsApiGatewayClientCertificateAttributes{ref: terra.ReferenceResource(aagcc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aagcc *Resource) ImportState(state io.Reader) error {
	aagcc.state = &awsApiGatewayClientCertificateState{}
	if err := json.NewDecoder(state).Decode(aagcc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aagcc.Type(), aagcc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aagcc *Resource) State() (*awsApiGatewayClientCertificateState, bool) {
	return aagcc.state, aagcc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aagcc *Resource) StateMust() *awsApiGatewayClientCertificateState {
	if aagcc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aagcc.Type(), aagcc.LocalName()))
	}
	return aagcc.state
}

// Args contains the configurations for aws_api_gateway_client_certificate.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}

type awsApiGatewayClientCertificateAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_api_gateway_client_certificate.
func (aagcc awsApiGatewayClientCertificateAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aagcc.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_api_gateway_client_certificate.
func (aagcc awsApiGatewayClientCertificateAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(aagcc.ref.Append("created_date"))
}

// Description returns a reference to field description of aws_api_gateway_client_certificate.
func (aagcc awsApiGatewayClientCertificateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aagcc.ref.Append("description"))
}

// ExpirationDate returns a reference to field expiration_date of aws_api_gateway_client_certificate.
func (aagcc awsApiGatewayClientCertificateAttributes) ExpirationDate() terra.StringValue {
	return terra.ReferenceAsString(aagcc.ref.Append("expiration_date"))
}

// Id returns a reference to field id of aws_api_gateway_client_certificate.
func (aagcc awsApiGatewayClientCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aagcc.ref.Append("id"))
}

// PemEncodedCertificate returns a reference to field pem_encoded_certificate of aws_api_gateway_client_certificate.
func (aagcc awsApiGatewayClientCertificateAttributes) PemEncodedCertificate() terra.StringValue {
	return terra.ReferenceAsString(aagcc.ref.Append("pem_encoded_certificate"))
}

// Tags returns a reference to field tags of aws_api_gateway_client_certificate.
func (aagcc awsApiGatewayClientCertificateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aagcc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_api_gateway_client_certificate.
func (aagcc awsApiGatewayClientCertificateAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aagcc.ref.Append("tags_all"))
}

type awsApiGatewayClientCertificateState struct {
	Arn                   string            `json:"arn"`
	CreatedDate           string            `json:"created_date"`
	Description           string            `json:"description"`
	ExpirationDate        string            `json:"expiration_date"`
	Id                    string            `json:"id"`
	PemEncodedCertificate string            `json:"pem_encoded_certificate"`
	Tags                  map[string]string `json:"tags"`
	TagsAll               map[string]string `json:"tags_all"`
}
