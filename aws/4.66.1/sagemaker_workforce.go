// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakerworkforce "github.com/golingon/terraproviders/aws/4.66.1/sagemakerworkforce"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSagemakerWorkforce creates a new instance of [SagemakerWorkforce].
func NewSagemakerWorkforce(name string, args SagemakerWorkforceArgs) *SagemakerWorkforce {
	return &SagemakerWorkforce{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerWorkforce)(nil)

// SagemakerWorkforce represents the Terraform resource aws_sagemaker_workforce.
type SagemakerWorkforce struct {
	Name      string
	Args      SagemakerWorkforceArgs
	state     *sagemakerWorkforceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerWorkforce].
func (sw *SagemakerWorkforce) Type() string {
	return "aws_sagemaker_workforce"
}

// LocalName returns the local name for [SagemakerWorkforce].
func (sw *SagemakerWorkforce) LocalName() string {
	return sw.Name
}

// Configuration returns the configuration (args) for [SagemakerWorkforce].
func (sw *SagemakerWorkforce) Configuration() interface{} {
	return sw.Args
}

// DependOn is used for other resources to depend on [SagemakerWorkforce].
func (sw *SagemakerWorkforce) DependOn() terra.Reference {
	return terra.ReferenceResource(sw)
}

// Dependencies returns the list of resources [SagemakerWorkforce] depends_on.
func (sw *SagemakerWorkforce) Dependencies() terra.Dependencies {
	return sw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerWorkforce].
func (sw *SagemakerWorkforce) LifecycleManagement() *terra.Lifecycle {
	return sw.Lifecycle
}

// Attributes returns the attributes for [SagemakerWorkforce].
func (sw *SagemakerWorkforce) Attributes() sagemakerWorkforceAttributes {
	return sagemakerWorkforceAttributes{ref: terra.ReferenceResource(sw)}
}

// ImportState imports the given attribute values into [SagemakerWorkforce]'s state.
func (sw *SagemakerWorkforce) ImportState(av io.Reader) error {
	sw.state = &sagemakerWorkforceState{}
	if err := json.NewDecoder(av).Decode(sw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sw.Type(), sw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerWorkforce] has state.
func (sw *SagemakerWorkforce) State() (*sagemakerWorkforceState, bool) {
	return sw.state, sw.state != nil
}

// StateMust returns the state for [SagemakerWorkforce]. Panics if the state is nil.
func (sw *SagemakerWorkforce) StateMust() *sagemakerWorkforceState {
	if sw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sw.Type(), sw.LocalName()))
	}
	return sw.state
}

// SagemakerWorkforceArgs contains the configurations for aws_sagemaker_workforce.
type SagemakerWorkforceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// WorkforceName: string, required
	WorkforceName terra.StringValue `hcl:"workforce_name,attr" validate:"required"`
	// CognitoConfig: optional
	CognitoConfig *sagemakerworkforce.CognitoConfig `hcl:"cognito_config,block"`
	// OidcConfig: optional
	OidcConfig *sagemakerworkforce.OidcConfig `hcl:"oidc_config,block"`
	// SourceIpConfig: optional
	SourceIpConfig *sagemakerworkforce.SourceIpConfig `hcl:"source_ip_config,block"`
	// WorkforceVpcConfig: optional
	WorkforceVpcConfig *sagemakerworkforce.WorkforceVpcConfig `hcl:"workforce_vpc_config,block"`
}
type sagemakerWorkforceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sagemaker_workforce.
func (sw sagemakerWorkforceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("arn"))
}

// Id returns a reference to field id of aws_sagemaker_workforce.
func (sw sagemakerWorkforceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("id"))
}

// Subdomain returns a reference to field subdomain of aws_sagemaker_workforce.
func (sw sagemakerWorkforceAttributes) Subdomain() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("subdomain"))
}

// WorkforceName returns a reference to field workforce_name of aws_sagemaker_workforce.
func (sw sagemakerWorkforceAttributes) WorkforceName() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("workforce_name"))
}

func (sw sagemakerWorkforceAttributes) CognitoConfig() terra.ListValue[sagemakerworkforce.CognitoConfigAttributes] {
	return terra.ReferenceAsList[sagemakerworkforce.CognitoConfigAttributes](sw.ref.Append("cognito_config"))
}

func (sw sagemakerWorkforceAttributes) OidcConfig() terra.ListValue[sagemakerworkforce.OidcConfigAttributes] {
	return terra.ReferenceAsList[sagemakerworkforce.OidcConfigAttributes](sw.ref.Append("oidc_config"))
}

func (sw sagemakerWorkforceAttributes) SourceIpConfig() terra.ListValue[sagemakerworkforce.SourceIpConfigAttributes] {
	return terra.ReferenceAsList[sagemakerworkforce.SourceIpConfigAttributes](sw.ref.Append("source_ip_config"))
}

func (sw sagemakerWorkforceAttributes) WorkforceVpcConfig() terra.ListValue[sagemakerworkforce.WorkforceVpcConfigAttributes] {
	return terra.ReferenceAsList[sagemakerworkforce.WorkforceVpcConfigAttributes](sw.ref.Append("workforce_vpc_config"))
}

type sagemakerWorkforceState struct {
	Arn                string                                       `json:"arn"`
	Id                 string                                       `json:"id"`
	Subdomain          string                                       `json:"subdomain"`
	WorkforceName      string                                       `json:"workforce_name"`
	CognitoConfig      []sagemakerworkforce.CognitoConfigState      `json:"cognito_config"`
	OidcConfig         []sagemakerworkforce.OidcConfigState         `json:"oidc_config"`
	SourceIpConfig     []sagemakerworkforce.SourceIpConfigState     `json:"source_ip_config"`
	WorkforceVpcConfig []sagemakerworkforce.WorkforceVpcConfigState `json:"workforce_vpc_config"`
}
