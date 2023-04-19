// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	apigatewaydomainname "github.com/golingon/terraproviders/aws/4.63.0/apigatewaydomainname"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayDomainName creates a new instance of [ApiGatewayDomainName].
func NewApiGatewayDomainName(name string, args ApiGatewayDomainNameArgs) *ApiGatewayDomainName {
	return &ApiGatewayDomainName{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayDomainName)(nil)

// ApiGatewayDomainName represents the Terraform resource aws_api_gateway_domain_name.
type ApiGatewayDomainName struct {
	Name      string
	Args      ApiGatewayDomainNameArgs
	state     *apiGatewayDomainNameState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayDomainName].
func (agdn *ApiGatewayDomainName) Type() string {
	return "aws_api_gateway_domain_name"
}

// LocalName returns the local name for [ApiGatewayDomainName].
func (agdn *ApiGatewayDomainName) LocalName() string {
	return agdn.Name
}

// Configuration returns the configuration (args) for [ApiGatewayDomainName].
func (agdn *ApiGatewayDomainName) Configuration() interface{} {
	return agdn.Args
}

// DependOn is used for other resources to depend on [ApiGatewayDomainName].
func (agdn *ApiGatewayDomainName) DependOn() terra.Reference {
	return terra.ReferenceResource(agdn)
}

// Dependencies returns the list of resources [ApiGatewayDomainName] depends_on.
func (agdn *ApiGatewayDomainName) Dependencies() terra.Dependencies {
	return agdn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayDomainName].
func (agdn *ApiGatewayDomainName) LifecycleManagement() *terra.Lifecycle {
	return agdn.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayDomainName].
func (agdn *ApiGatewayDomainName) Attributes() apiGatewayDomainNameAttributes {
	return apiGatewayDomainNameAttributes{ref: terra.ReferenceResource(agdn)}
}

// ImportState imports the given attribute values into [ApiGatewayDomainName]'s state.
func (agdn *ApiGatewayDomainName) ImportState(av io.Reader) error {
	agdn.state = &apiGatewayDomainNameState{}
	if err := json.NewDecoder(av).Decode(agdn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agdn.Type(), agdn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayDomainName] has state.
func (agdn *ApiGatewayDomainName) State() (*apiGatewayDomainNameState, bool) {
	return agdn.state, agdn.state != nil
}

// StateMust returns the state for [ApiGatewayDomainName]. Panics if the state is nil.
func (agdn *ApiGatewayDomainName) StateMust() *apiGatewayDomainNameState {
	if agdn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agdn.Type(), agdn.LocalName()))
	}
	return agdn.state
}

// ApiGatewayDomainNameArgs contains the configurations for aws_api_gateway_domain_name.
type ApiGatewayDomainNameArgs struct {
	// CertificateArn: string, optional
	CertificateArn terra.StringValue `hcl:"certificate_arn,attr"`
	// CertificateBody: string, optional
	CertificateBody terra.StringValue `hcl:"certificate_body,attr"`
	// CertificateChain: string, optional
	CertificateChain terra.StringValue `hcl:"certificate_chain,attr"`
	// CertificateName: string, optional
	CertificateName terra.StringValue `hcl:"certificate_name,attr"`
	// CertificatePrivateKey: string, optional
	CertificatePrivateKey terra.StringValue `hcl:"certificate_private_key,attr"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OwnershipVerificationCertificateArn: string, optional
	OwnershipVerificationCertificateArn terra.StringValue `hcl:"ownership_verification_certificate_arn,attr"`
	// RegionalCertificateArn: string, optional
	RegionalCertificateArn terra.StringValue `hcl:"regional_certificate_arn,attr"`
	// RegionalCertificateName: string, optional
	RegionalCertificateName terra.StringValue `hcl:"regional_certificate_name,attr"`
	// SecurityPolicy: string, optional
	SecurityPolicy terra.StringValue `hcl:"security_policy,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// EndpointConfiguration: optional
	EndpointConfiguration *apigatewaydomainname.EndpointConfiguration `hcl:"endpoint_configuration,block"`
	// MutualTlsAuthentication: optional
	MutualTlsAuthentication *apigatewaydomainname.MutualTlsAuthentication `hcl:"mutual_tls_authentication,block"`
}
type apiGatewayDomainNameAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("arn"))
}

// CertificateArn returns a reference to field certificate_arn of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) CertificateArn() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("certificate_arn"))
}

// CertificateBody returns a reference to field certificate_body of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) CertificateBody() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("certificate_body"))
}

// CertificateChain returns a reference to field certificate_chain of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) CertificateChain() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("certificate_chain"))
}

// CertificateName returns a reference to field certificate_name of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) CertificateName() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("certificate_name"))
}

// CertificatePrivateKey returns a reference to field certificate_private_key of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) CertificatePrivateKey() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("certificate_private_key"))
}

// CertificateUploadDate returns a reference to field certificate_upload_date of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) CertificateUploadDate() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("certificate_upload_date"))
}

// CloudfrontDomainName returns a reference to field cloudfront_domain_name of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) CloudfrontDomainName() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("cloudfront_domain_name"))
}

// CloudfrontZoneId returns a reference to field cloudfront_zone_id of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) CloudfrontZoneId() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("cloudfront_zone_id"))
}

// DomainName returns a reference to field domain_name of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("domain_name"))
}

// Id returns a reference to field id of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("id"))
}

// OwnershipVerificationCertificateArn returns a reference to field ownership_verification_certificate_arn of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) OwnershipVerificationCertificateArn() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("ownership_verification_certificate_arn"))
}

// RegionalCertificateArn returns a reference to field regional_certificate_arn of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) RegionalCertificateArn() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("regional_certificate_arn"))
}

// RegionalCertificateName returns a reference to field regional_certificate_name of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) RegionalCertificateName() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("regional_certificate_name"))
}

// RegionalDomainName returns a reference to field regional_domain_name of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) RegionalDomainName() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("regional_domain_name"))
}

// RegionalZoneId returns a reference to field regional_zone_id of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) RegionalZoneId() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("regional_zone_id"))
}

// SecurityPolicy returns a reference to field security_policy of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) SecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(agdn.ref.Append("security_policy"))
}

// Tags returns a reference to field tags of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agdn.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_api_gateway_domain_name.
func (agdn apiGatewayDomainNameAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agdn.ref.Append("tags_all"))
}

func (agdn apiGatewayDomainNameAttributes) EndpointConfiguration() terra.ListValue[apigatewaydomainname.EndpointConfigurationAttributes] {
	return terra.ReferenceAsList[apigatewaydomainname.EndpointConfigurationAttributes](agdn.ref.Append("endpoint_configuration"))
}

func (agdn apiGatewayDomainNameAttributes) MutualTlsAuthentication() terra.ListValue[apigatewaydomainname.MutualTlsAuthenticationAttributes] {
	return terra.ReferenceAsList[apigatewaydomainname.MutualTlsAuthenticationAttributes](agdn.ref.Append("mutual_tls_authentication"))
}

type apiGatewayDomainNameState struct {
	Arn                                 string                                              `json:"arn"`
	CertificateArn                      string                                              `json:"certificate_arn"`
	CertificateBody                     string                                              `json:"certificate_body"`
	CertificateChain                    string                                              `json:"certificate_chain"`
	CertificateName                     string                                              `json:"certificate_name"`
	CertificatePrivateKey               string                                              `json:"certificate_private_key"`
	CertificateUploadDate               string                                              `json:"certificate_upload_date"`
	CloudfrontDomainName                string                                              `json:"cloudfront_domain_name"`
	CloudfrontZoneId                    string                                              `json:"cloudfront_zone_id"`
	DomainName                          string                                              `json:"domain_name"`
	Id                                  string                                              `json:"id"`
	OwnershipVerificationCertificateArn string                                              `json:"ownership_verification_certificate_arn"`
	RegionalCertificateArn              string                                              `json:"regional_certificate_arn"`
	RegionalCertificateName             string                                              `json:"regional_certificate_name"`
	RegionalDomainName                  string                                              `json:"regional_domain_name"`
	RegionalZoneId                      string                                              `json:"regional_zone_id"`
	SecurityPolicy                      string                                              `json:"security_policy"`
	Tags                                map[string]string                                   `json:"tags"`
	TagsAll                             map[string]string                                   `json:"tags_all"`
	EndpointConfiguration               []apigatewaydomainname.EndpointConfigurationState   `json:"endpoint_configuration"`
	MutualTlsAuthentication             []apigatewaydomainname.MutualTlsAuthenticationState `json:"mutual_tls_authentication"`
}
