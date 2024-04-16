// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_verifiedaccess_endpoint

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

// Resource represents the Terraform resource aws_verifiedaccess_endpoint.
type Resource struct {
	Name      string
	Args      Args
	state     *awsVerifiedaccessEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ave *Resource) Type() string {
	return "aws_verifiedaccess_endpoint"
}

// LocalName returns the local name for [Resource].
func (ave *Resource) LocalName() string {
	return ave.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ave *Resource) Configuration() interface{} {
	return ave.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ave *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ave)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ave *Resource) Dependencies() terra.Dependencies {
	return ave.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ave *Resource) LifecycleManagement() *terra.Lifecycle {
	return ave.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ave *Resource) Attributes() awsVerifiedaccessEndpointAttributes {
	return awsVerifiedaccessEndpointAttributes{ref: terra.ReferenceResource(ave)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ave *Resource) ImportState(state io.Reader) error {
	ave.state = &awsVerifiedaccessEndpointState{}
	if err := json.NewDecoder(state).Decode(ave.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ave.Type(), ave.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ave *Resource) State() (*awsVerifiedaccessEndpointState, bool) {
	return ave.state, ave.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ave *Resource) StateMust() *awsVerifiedaccessEndpointState {
	if ave.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ave.Type(), ave.LocalName()))
	}
	return ave.state
}

// Args contains the configurations for aws_verifiedaccess_endpoint.
type Args struct {
	// ApplicationDomain: string, required
	ApplicationDomain terra.StringValue `hcl:"application_domain,attr" validate:"required"`
	// AttachmentType: string, required
	AttachmentType terra.StringValue `hcl:"attachment_type,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DomainCertificateArn: string, required
	DomainCertificateArn terra.StringValue `hcl:"domain_certificate_arn,attr" validate:"required"`
	// EndpointDomainPrefix: string, required
	EndpointDomainPrefix terra.StringValue `hcl:"endpoint_domain_prefix,attr" validate:"required"`
	// EndpointType: string, required
	EndpointType terra.StringValue `hcl:"endpoint_type,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyDocument: string, optional
	PolicyDocument terra.StringValue `hcl:"policy_document,attr"`
	// SecurityGroupIds: set of string, optional
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VerifiedAccessGroupId: string, required
	VerifiedAccessGroupId terra.StringValue `hcl:"verified_access_group_id,attr" validate:"required"`
	// LoadBalancerOptions: optional
	LoadBalancerOptions *LoadBalancerOptions `hcl:"load_balancer_options,block"`
	// NetworkInterfaceOptions: optional
	NetworkInterfaceOptions *NetworkInterfaceOptions `hcl:"network_interface_options,block"`
	// SseSpecification: optional
	SseSpecification *SseSpecification `hcl:"sse_specification,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsVerifiedaccessEndpointAttributes struct {
	ref terra.Reference
}

// ApplicationDomain returns a reference to field application_domain of aws_verifiedaccess_endpoint.
func (ave awsVerifiedaccessEndpointAttributes) ApplicationDomain() terra.StringValue {
	return terra.ReferenceAsString(ave.ref.Append("application_domain"))
}

// AttachmentType returns a reference to field attachment_type of aws_verifiedaccess_endpoint.
func (ave awsVerifiedaccessEndpointAttributes) AttachmentType() terra.StringValue {
	return terra.ReferenceAsString(ave.ref.Append("attachment_type"))
}

// Description returns a reference to field description of aws_verifiedaccess_endpoint.
func (ave awsVerifiedaccessEndpointAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ave.ref.Append("description"))
}

// DeviceValidationDomain returns a reference to field device_validation_domain of aws_verifiedaccess_endpoint.
func (ave awsVerifiedaccessEndpointAttributes) DeviceValidationDomain() terra.StringValue {
	return terra.ReferenceAsString(ave.ref.Append("device_validation_domain"))
}

// DomainCertificateArn returns a reference to field domain_certificate_arn of aws_verifiedaccess_endpoint.
func (ave awsVerifiedaccessEndpointAttributes) DomainCertificateArn() terra.StringValue {
	return terra.ReferenceAsString(ave.ref.Append("domain_certificate_arn"))
}

// EndpointDomain returns a reference to field endpoint_domain of aws_verifiedaccess_endpoint.
func (ave awsVerifiedaccessEndpointAttributes) EndpointDomain() terra.StringValue {
	return terra.ReferenceAsString(ave.ref.Append("endpoint_domain"))
}

// EndpointDomainPrefix returns a reference to field endpoint_domain_prefix of aws_verifiedaccess_endpoint.
func (ave awsVerifiedaccessEndpointAttributes) EndpointDomainPrefix() terra.StringValue {
	return terra.ReferenceAsString(ave.ref.Append("endpoint_domain_prefix"))
}

// EndpointType returns a reference to field endpoint_type of aws_verifiedaccess_endpoint.
func (ave awsVerifiedaccessEndpointAttributes) EndpointType() terra.StringValue {
	return terra.ReferenceAsString(ave.ref.Append("endpoint_type"))
}

// Id returns a reference to field id of aws_verifiedaccess_endpoint.
func (ave awsVerifiedaccessEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ave.ref.Append("id"))
}

// PolicyDocument returns a reference to field policy_document of aws_verifiedaccess_endpoint.
func (ave awsVerifiedaccessEndpointAttributes) PolicyDocument() terra.StringValue {
	return terra.ReferenceAsString(ave.ref.Append("policy_document"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_verifiedaccess_endpoint.
func (ave awsVerifiedaccessEndpointAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ave.ref.Append("security_group_ids"))
}

// Tags returns a reference to field tags of aws_verifiedaccess_endpoint.
func (ave awsVerifiedaccessEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ave.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_verifiedaccess_endpoint.
func (ave awsVerifiedaccessEndpointAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ave.ref.Append("tags_all"))
}

// VerifiedAccessGroupId returns a reference to field verified_access_group_id of aws_verifiedaccess_endpoint.
func (ave awsVerifiedaccessEndpointAttributes) VerifiedAccessGroupId() terra.StringValue {
	return terra.ReferenceAsString(ave.ref.Append("verified_access_group_id"))
}

// VerifiedAccessInstanceId returns a reference to field verified_access_instance_id of aws_verifiedaccess_endpoint.
func (ave awsVerifiedaccessEndpointAttributes) VerifiedAccessInstanceId() terra.StringValue {
	return terra.ReferenceAsString(ave.ref.Append("verified_access_instance_id"))
}

func (ave awsVerifiedaccessEndpointAttributes) LoadBalancerOptions() terra.ListValue[LoadBalancerOptionsAttributes] {
	return terra.ReferenceAsList[LoadBalancerOptionsAttributes](ave.ref.Append("load_balancer_options"))
}

func (ave awsVerifiedaccessEndpointAttributes) NetworkInterfaceOptions() terra.ListValue[NetworkInterfaceOptionsAttributes] {
	return terra.ReferenceAsList[NetworkInterfaceOptionsAttributes](ave.ref.Append("network_interface_options"))
}

func (ave awsVerifiedaccessEndpointAttributes) SseSpecification() terra.ListValue[SseSpecificationAttributes] {
	return terra.ReferenceAsList[SseSpecificationAttributes](ave.ref.Append("sse_specification"))
}

func (ave awsVerifiedaccessEndpointAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ave.ref.Append("timeouts"))
}

type awsVerifiedaccessEndpointState struct {
	ApplicationDomain        string                         `json:"application_domain"`
	AttachmentType           string                         `json:"attachment_type"`
	Description              string                         `json:"description"`
	DeviceValidationDomain   string                         `json:"device_validation_domain"`
	DomainCertificateArn     string                         `json:"domain_certificate_arn"`
	EndpointDomain           string                         `json:"endpoint_domain"`
	EndpointDomainPrefix     string                         `json:"endpoint_domain_prefix"`
	EndpointType             string                         `json:"endpoint_type"`
	Id                       string                         `json:"id"`
	PolicyDocument           string                         `json:"policy_document"`
	SecurityGroupIds         []string                       `json:"security_group_ids"`
	Tags                     map[string]string              `json:"tags"`
	TagsAll                  map[string]string              `json:"tags_all"`
	VerifiedAccessGroupId    string                         `json:"verified_access_group_id"`
	VerifiedAccessInstanceId string                         `json:"verified_access_instance_id"`
	LoadBalancerOptions      []LoadBalancerOptionsState     `json:"load_balancer_options"`
	NetworkInterfaceOptions  []NetworkInterfaceOptionsState `json:"network_interface_options"`
	SseSpecification         []SseSpecificationState        `json:"sse_specification"`
	Timeouts                 *TimeoutsState                 `json:"timeouts"`
}
