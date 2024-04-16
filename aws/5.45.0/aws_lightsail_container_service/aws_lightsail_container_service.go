// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_lightsail_container_service

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

// Resource represents the Terraform resource aws_lightsail_container_service.
type Resource struct {
	Name      string
	Args      Args
	state     *awsLightsailContainerServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (alcs *Resource) Type() string {
	return "aws_lightsail_container_service"
}

// LocalName returns the local name for [Resource].
func (alcs *Resource) LocalName() string {
	return alcs.Name
}

// Configuration returns the configuration (args) for [Resource].
func (alcs *Resource) Configuration() interface{} {
	return alcs.Args
}

// DependOn is used for other resources to depend on [Resource].
func (alcs *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(alcs)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (alcs *Resource) Dependencies() terra.Dependencies {
	return alcs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (alcs *Resource) LifecycleManagement() *terra.Lifecycle {
	return alcs.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (alcs *Resource) Attributes() awsLightsailContainerServiceAttributes {
	return awsLightsailContainerServiceAttributes{ref: terra.ReferenceResource(alcs)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (alcs *Resource) ImportState(state io.Reader) error {
	alcs.state = &awsLightsailContainerServiceState{}
	if err := json.NewDecoder(state).Decode(alcs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", alcs.Type(), alcs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (alcs *Resource) State() (*awsLightsailContainerServiceState, bool) {
	return alcs.state, alcs.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (alcs *Resource) StateMust() *awsLightsailContainerServiceState {
	if alcs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", alcs.Type(), alcs.LocalName()))
	}
	return alcs.state
}

// Args contains the configurations for aws_lightsail_container_service.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsDisabled: bool, optional
	IsDisabled terra.BoolValue `hcl:"is_disabled,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Power: string, required
	Power terra.StringValue `hcl:"power,attr" validate:"required"`
	// Scale: number, required
	Scale terra.NumberValue `hcl:"scale,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// PrivateRegistryAccess: optional
	PrivateRegistryAccess *PrivateRegistryAccess `hcl:"private_registry_access,block"`
	// PublicDomainNames: optional
	PublicDomainNames *PublicDomainNames `hcl:"public_domain_names,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsLightsailContainerServiceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lightsail_container_service.
func (alcs awsLightsailContainerServiceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(alcs.ref.Append("arn"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_lightsail_container_service.
func (alcs awsLightsailContainerServiceAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(alcs.ref.Append("availability_zone"))
}

// CreatedAt returns a reference to field created_at of aws_lightsail_container_service.
func (alcs awsLightsailContainerServiceAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(alcs.ref.Append("created_at"))
}

// Id returns a reference to field id of aws_lightsail_container_service.
func (alcs awsLightsailContainerServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(alcs.ref.Append("id"))
}

// IsDisabled returns a reference to field is_disabled of aws_lightsail_container_service.
func (alcs awsLightsailContainerServiceAttributes) IsDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(alcs.ref.Append("is_disabled"))
}

// Name returns a reference to field name of aws_lightsail_container_service.
func (alcs awsLightsailContainerServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(alcs.ref.Append("name"))
}

// Power returns a reference to field power of aws_lightsail_container_service.
func (alcs awsLightsailContainerServiceAttributes) Power() terra.StringValue {
	return terra.ReferenceAsString(alcs.ref.Append("power"))
}

// PowerId returns a reference to field power_id of aws_lightsail_container_service.
func (alcs awsLightsailContainerServiceAttributes) PowerId() terra.StringValue {
	return terra.ReferenceAsString(alcs.ref.Append("power_id"))
}

// PrincipalArn returns a reference to field principal_arn of aws_lightsail_container_service.
func (alcs awsLightsailContainerServiceAttributes) PrincipalArn() terra.StringValue {
	return terra.ReferenceAsString(alcs.ref.Append("principal_arn"))
}

// PrivateDomainName returns a reference to field private_domain_name of aws_lightsail_container_service.
func (alcs awsLightsailContainerServiceAttributes) PrivateDomainName() terra.StringValue {
	return terra.ReferenceAsString(alcs.ref.Append("private_domain_name"))
}

// ResourceType returns a reference to field resource_type of aws_lightsail_container_service.
func (alcs awsLightsailContainerServiceAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(alcs.ref.Append("resource_type"))
}

// Scale returns a reference to field scale of aws_lightsail_container_service.
func (alcs awsLightsailContainerServiceAttributes) Scale() terra.NumberValue {
	return terra.ReferenceAsNumber(alcs.ref.Append("scale"))
}

// State returns a reference to field state of aws_lightsail_container_service.
func (alcs awsLightsailContainerServiceAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(alcs.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_lightsail_container_service.
func (alcs awsLightsailContainerServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](alcs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_lightsail_container_service.
func (alcs awsLightsailContainerServiceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](alcs.ref.Append("tags_all"))
}

// Url returns a reference to field url of aws_lightsail_container_service.
func (alcs awsLightsailContainerServiceAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(alcs.ref.Append("url"))
}

func (alcs awsLightsailContainerServiceAttributes) PrivateRegistryAccess() terra.ListValue[PrivateRegistryAccessAttributes] {
	return terra.ReferenceAsList[PrivateRegistryAccessAttributes](alcs.ref.Append("private_registry_access"))
}

func (alcs awsLightsailContainerServiceAttributes) PublicDomainNames() terra.ListValue[PublicDomainNamesAttributes] {
	return terra.ReferenceAsList[PublicDomainNamesAttributes](alcs.ref.Append("public_domain_names"))
}

func (alcs awsLightsailContainerServiceAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](alcs.ref.Append("timeouts"))
}

type awsLightsailContainerServiceState struct {
	Arn                   string                       `json:"arn"`
	AvailabilityZone      string                       `json:"availability_zone"`
	CreatedAt             string                       `json:"created_at"`
	Id                    string                       `json:"id"`
	IsDisabled            bool                         `json:"is_disabled"`
	Name                  string                       `json:"name"`
	Power                 string                       `json:"power"`
	PowerId               string                       `json:"power_id"`
	PrincipalArn          string                       `json:"principal_arn"`
	PrivateDomainName     string                       `json:"private_domain_name"`
	ResourceType          string                       `json:"resource_type"`
	Scale                 float64                      `json:"scale"`
	State                 string                       `json:"state"`
	Tags                  map[string]string            `json:"tags"`
	TagsAll               map[string]string            `json:"tags_all"`
	Url                   string                       `json:"url"`
	PrivateRegistryAccess []PrivateRegistryAccessState `json:"private_registry_access"`
	PublicDomainNames     []PublicDomainNamesState     `json:"public_domain_names"`
	Timeouts              *TimeoutsState               `json:"timeouts"`
}
