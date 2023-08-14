// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudfrontcontinuousdeploymentpolicy "github.com/golingon/terraproviders/aws/5.12.0/cloudfrontcontinuousdeploymentpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudfrontContinuousDeploymentPolicy creates a new instance of [CloudfrontContinuousDeploymentPolicy].
func NewCloudfrontContinuousDeploymentPolicy(name string, args CloudfrontContinuousDeploymentPolicyArgs) *CloudfrontContinuousDeploymentPolicy {
	return &CloudfrontContinuousDeploymentPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudfrontContinuousDeploymentPolicy)(nil)

// CloudfrontContinuousDeploymentPolicy represents the Terraform resource aws_cloudfront_continuous_deployment_policy.
type CloudfrontContinuousDeploymentPolicy struct {
	Name      string
	Args      CloudfrontContinuousDeploymentPolicyArgs
	state     *cloudfrontContinuousDeploymentPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudfrontContinuousDeploymentPolicy].
func (ccdp *CloudfrontContinuousDeploymentPolicy) Type() string {
	return "aws_cloudfront_continuous_deployment_policy"
}

// LocalName returns the local name for [CloudfrontContinuousDeploymentPolicy].
func (ccdp *CloudfrontContinuousDeploymentPolicy) LocalName() string {
	return ccdp.Name
}

// Configuration returns the configuration (args) for [CloudfrontContinuousDeploymentPolicy].
func (ccdp *CloudfrontContinuousDeploymentPolicy) Configuration() interface{} {
	return ccdp.Args
}

// DependOn is used for other resources to depend on [CloudfrontContinuousDeploymentPolicy].
func (ccdp *CloudfrontContinuousDeploymentPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ccdp)
}

// Dependencies returns the list of resources [CloudfrontContinuousDeploymentPolicy] depends_on.
func (ccdp *CloudfrontContinuousDeploymentPolicy) Dependencies() terra.Dependencies {
	return ccdp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudfrontContinuousDeploymentPolicy].
func (ccdp *CloudfrontContinuousDeploymentPolicy) LifecycleManagement() *terra.Lifecycle {
	return ccdp.Lifecycle
}

// Attributes returns the attributes for [CloudfrontContinuousDeploymentPolicy].
func (ccdp *CloudfrontContinuousDeploymentPolicy) Attributes() cloudfrontContinuousDeploymentPolicyAttributes {
	return cloudfrontContinuousDeploymentPolicyAttributes{ref: terra.ReferenceResource(ccdp)}
}

// ImportState imports the given attribute values into [CloudfrontContinuousDeploymentPolicy]'s state.
func (ccdp *CloudfrontContinuousDeploymentPolicy) ImportState(av io.Reader) error {
	ccdp.state = &cloudfrontContinuousDeploymentPolicyState{}
	if err := json.NewDecoder(av).Decode(ccdp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ccdp.Type(), ccdp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudfrontContinuousDeploymentPolicy] has state.
func (ccdp *CloudfrontContinuousDeploymentPolicy) State() (*cloudfrontContinuousDeploymentPolicyState, bool) {
	return ccdp.state, ccdp.state != nil
}

// StateMust returns the state for [CloudfrontContinuousDeploymentPolicy]. Panics if the state is nil.
func (ccdp *CloudfrontContinuousDeploymentPolicy) StateMust() *cloudfrontContinuousDeploymentPolicyState {
	if ccdp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ccdp.Type(), ccdp.LocalName()))
	}
	return ccdp.state
}

// CloudfrontContinuousDeploymentPolicyArgs contains the configurations for aws_cloudfront_continuous_deployment_policy.
type CloudfrontContinuousDeploymentPolicyArgs struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// StagingDistributionDnsNames: min=0
	StagingDistributionDnsNames []cloudfrontcontinuousdeploymentpolicy.StagingDistributionDnsNames `hcl:"staging_distribution_dns_names,block" validate:"min=0"`
	// TrafficConfig: min=0
	TrafficConfig []cloudfrontcontinuousdeploymentpolicy.TrafficConfig `hcl:"traffic_config,block" validate:"min=0"`
}
type cloudfrontContinuousDeploymentPolicyAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of aws_cloudfront_continuous_deployment_policy.
func (ccdp cloudfrontContinuousDeploymentPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ccdp.ref.Append("enabled"))
}

// Etag returns a reference to field etag of aws_cloudfront_continuous_deployment_policy.
func (ccdp cloudfrontContinuousDeploymentPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ccdp.ref.Append("etag"))
}

// Id returns a reference to field id of aws_cloudfront_continuous_deployment_policy.
func (ccdp cloudfrontContinuousDeploymentPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccdp.ref.Append("id"))
}

// LastModifiedTime returns a reference to field last_modified_time of aws_cloudfront_continuous_deployment_policy.
func (ccdp cloudfrontContinuousDeploymentPolicyAttributes) LastModifiedTime() terra.StringValue {
	return terra.ReferenceAsString(ccdp.ref.Append("last_modified_time"))
}

func (ccdp cloudfrontContinuousDeploymentPolicyAttributes) StagingDistributionDnsNames() terra.ListValue[cloudfrontcontinuousdeploymentpolicy.StagingDistributionDnsNamesAttributes] {
	return terra.ReferenceAsList[cloudfrontcontinuousdeploymentpolicy.StagingDistributionDnsNamesAttributes](ccdp.ref.Append("staging_distribution_dns_names"))
}

func (ccdp cloudfrontContinuousDeploymentPolicyAttributes) TrafficConfig() terra.ListValue[cloudfrontcontinuousdeploymentpolicy.TrafficConfigAttributes] {
	return terra.ReferenceAsList[cloudfrontcontinuousdeploymentpolicy.TrafficConfigAttributes](ccdp.ref.Append("traffic_config"))
}

type cloudfrontContinuousDeploymentPolicyState struct {
	Enabled                     bool                                                                    `json:"enabled"`
	Etag                        string                                                                  `json:"etag"`
	Id                          string                                                                  `json:"id"`
	LastModifiedTime            string                                                                  `json:"last_modified_time"`
	StagingDistributionDnsNames []cloudfrontcontinuousdeploymentpolicy.StagingDistributionDnsNamesState `json:"staging_distribution_dns_names"`
	TrafficConfig               []cloudfrontcontinuousdeploymentpolicy.TrafficConfigState               `json:"traffic_config"`
}
