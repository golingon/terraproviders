// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCognitoUserPoolDomain creates a new instance of [CognitoUserPoolDomain].
func NewCognitoUserPoolDomain(name string, args CognitoUserPoolDomainArgs) *CognitoUserPoolDomain {
	return &CognitoUserPoolDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CognitoUserPoolDomain)(nil)

// CognitoUserPoolDomain represents the Terraform resource aws_cognito_user_pool_domain.
type CognitoUserPoolDomain struct {
	Name      string
	Args      CognitoUserPoolDomainArgs
	state     *cognitoUserPoolDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CognitoUserPoolDomain].
func (cupd *CognitoUserPoolDomain) Type() string {
	return "aws_cognito_user_pool_domain"
}

// LocalName returns the local name for [CognitoUserPoolDomain].
func (cupd *CognitoUserPoolDomain) LocalName() string {
	return cupd.Name
}

// Configuration returns the configuration (args) for [CognitoUserPoolDomain].
func (cupd *CognitoUserPoolDomain) Configuration() interface{} {
	return cupd.Args
}

// DependOn is used for other resources to depend on [CognitoUserPoolDomain].
func (cupd *CognitoUserPoolDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(cupd)
}

// Dependencies returns the list of resources [CognitoUserPoolDomain] depends_on.
func (cupd *CognitoUserPoolDomain) Dependencies() terra.Dependencies {
	return cupd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CognitoUserPoolDomain].
func (cupd *CognitoUserPoolDomain) LifecycleManagement() *terra.Lifecycle {
	return cupd.Lifecycle
}

// Attributes returns the attributes for [CognitoUserPoolDomain].
func (cupd *CognitoUserPoolDomain) Attributes() cognitoUserPoolDomainAttributes {
	return cognitoUserPoolDomainAttributes{ref: terra.ReferenceResource(cupd)}
}

// ImportState imports the given attribute values into [CognitoUserPoolDomain]'s state.
func (cupd *CognitoUserPoolDomain) ImportState(av io.Reader) error {
	cupd.state = &cognitoUserPoolDomainState{}
	if err := json.NewDecoder(av).Decode(cupd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cupd.Type(), cupd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CognitoUserPoolDomain] has state.
func (cupd *CognitoUserPoolDomain) State() (*cognitoUserPoolDomainState, bool) {
	return cupd.state, cupd.state != nil
}

// StateMust returns the state for [CognitoUserPoolDomain]. Panics if the state is nil.
func (cupd *CognitoUserPoolDomain) StateMust() *cognitoUserPoolDomainState {
	if cupd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cupd.Type(), cupd.LocalName()))
	}
	return cupd.state
}

// CognitoUserPoolDomainArgs contains the configurations for aws_cognito_user_pool_domain.
type CognitoUserPoolDomainArgs struct {
	// CertificateArn: string, optional
	CertificateArn terra.StringValue `hcl:"certificate_arn,attr"`
	// Domain: string, required
	Domain terra.StringValue `hcl:"domain,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// UserPoolId: string, required
	UserPoolId terra.StringValue `hcl:"user_pool_id,attr" validate:"required"`
}
type cognitoUserPoolDomainAttributes struct {
	ref terra.Reference
}

// AwsAccountId returns a reference to field aws_account_id of aws_cognito_user_pool_domain.
func (cupd cognitoUserPoolDomainAttributes) AwsAccountId() terra.StringValue {
	return terra.ReferenceAsString(cupd.ref.Append("aws_account_id"))
}

// CertificateArn returns a reference to field certificate_arn of aws_cognito_user_pool_domain.
func (cupd cognitoUserPoolDomainAttributes) CertificateArn() terra.StringValue {
	return terra.ReferenceAsString(cupd.ref.Append("certificate_arn"))
}

// CloudfrontDistribution returns a reference to field cloudfront_distribution of aws_cognito_user_pool_domain.
func (cupd cognitoUserPoolDomainAttributes) CloudfrontDistribution() terra.StringValue {
	return terra.ReferenceAsString(cupd.ref.Append("cloudfront_distribution"))
}

// CloudfrontDistributionArn returns a reference to field cloudfront_distribution_arn of aws_cognito_user_pool_domain.
func (cupd cognitoUserPoolDomainAttributes) CloudfrontDistributionArn() terra.StringValue {
	return terra.ReferenceAsString(cupd.ref.Append("cloudfront_distribution_arn"))
}

// CloudfrontDistributionZoneId returns a reference to field cloudfront_distribution_zone_id of aws_cognito_user_pool_domain.
func (cupd cognitoUserPoolDomainAttributes) CloudfrontDistributionZoneId() terra.StringValue {
	return terra.ReferenceAsString(cupd.ref.Append("cloudfront_distribution_zone_id"))
}

// Domain returns a reference to field domain of aws_cognito_user_pool_domain.
func (cupd cognitoUserPoolDomainAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(cupd.ref.Append("domain"))
}

// Id returns a reference to field id of aws_cognito_user_pool_domain.
func (cupd cognitoUserPoolDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cupd.ref.Append("id"))
}

// S3Bucket returns a reference to field s3_bucket of aws_cognito_user_pool_domain.
func (cupd cognitoUserPoolDomainAttributes) S3Bucket() terra.StringValue {
	return terra.ReferenceAsString(cupd.ref.Append("s3_bucket"))
}

// UserPoolId returns a reference to field user_pool_id of aws_cognito_user_pool_domain.
func (cupd cognitoUserPoolDomainAttributes) UserPoolId() terra.StringValue {
	return terra.ReferenceAsString(cupd.ref.Append("user_pool_id"))
}

// Version returns a reference to field version of aws_cognito_user_pool_domain.
func (cupd cognitoUserPoolDomainAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(cupd.ref.Append("version"))
}

type cognitoUserPoolDomainState struct {
	AwsAccountId                 string `json:"aws_account_id"`
	CertificateArn               string `json:"certificate_arn"`
	CloudfrontDistribution       string `json:"cloudfront_distribution"`
	CloudfrontDistributionArn    string `json:"cloudfront_distribution_arn"`
	CloudfrontDistributionZoneId string `json:"cloudfront_distribution_zone_id"`
	Domain                       string `json:"domain"`
	Id                           string `json:"id"`
	S3Bucket                     string `json:"s3_bucket"`
	UserPoolId                   string `json:"user_pool_id"`
	Version                      string `json:"version"`
}