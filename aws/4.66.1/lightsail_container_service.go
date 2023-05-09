// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lightsailcontainerservice "github.com/golingon/terraproviders/aws/4.66.1/lightsailcontainerservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLightsailContainerService creates a new instance of [LightsailContainerService].
func NewLightsailContainerService(name string, args LightsailContainerServiceArgs) *LightsailContainerService {
	return &LightsailContainerService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LightsailContainerService)(nil)

// LightsailContainerService represents the Terraform resource aws_lightsail_container_service.
type LightsailContainerService struct {
	Name      string
	Args      LightsailContainerServiceArgs
	state     *lightsailContainerServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LightsailContainerService].
func (lcs *LightsailContainerService) Type() string {
	return "aws_lightsail_container_service"
}

// LocalName returns the local name for [LightsailContainerService].
func (lcs *LightsailContainerService) LocalName() string {
	return lcs.Name
}

// Configuration returns the configuration (args) for [LightsailContainerService].
func (lcs *LightsailContainerService) Configuration() interface{} {
	return lcs.Args
}

// DependOn is used for other resources to depend on [LightsailContainerService].
func (lcs *LightsailContainerService) DependOn() terra.Reference {
	return terra.ReferenceResource(lcs)
}

// Dependencies returns the list of resources [LightsailContainerService] depends_on.
func (lcs *LightsailContainerService) Dependencies() terra.Dependencies {
	return lcs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LightsailContainerService].
func (lcs *LightsailContainerService) LifecycleManagement() *terra.Lifecycle {
	return lcs.Lifecycle
}

// Attributes returns the attributes for [LightsailContainerService].
func (lcs *LightsailContainerService) Attributes() lightsailContainerServiceAttributes {
	return lightsailContainerServiceAttributes{ref: terra.ReferenceResource(lcs)}
}

// ImportState imports the given attribute values into [LightsailContainerService]'s state.
func (lcs *LightsailContainerService) ImportState(av io.Reader) error {
	lcs.state = &lightsailContainerServiceState{}
	if err := json.NewDecoder(av).Decode(lcs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lcs.Type(), lcs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LightsailContainerService] has state.
func (lcs *LightsailContainerService) State() (*lightsailContainerServiceState, bool) {
	return lcs.state, lcs.state != nil
}

// StateMust returns the state for [LightsailContainerService]. Panics if the state is nil.
func (lcs *LightsailContainerService) StateMust() *lightsailContainerServiceState {
	if lcs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lcs.Type(), lcs.LocalName()))
	}
	return lcs.state
}

// LightsailContainerServiceArgs contains the configurations for aws_lightsail_container_service.
type LightsailContainerServiceArgs struct {
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
	PrivateRegistryAccess *lightsailcontainerservice.PrivateRegistryAccess `hcl:"private_registry_access,block"`
	// PublicDomainNames: optional
	PublicDomainNames *lightsailcontainerservice.PublicDomainNames `hcl:"public_domain_names,block"`
	// Timeouts: optional
	Timeouts *lightsailcontainerservice.Timeouts `hcl:"timeouts,block"`
}
type lightsailContainerServiceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lightsail_container_service.
func (lcs lightsailContainerServiceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(lcs.ref.Append("arn"))
}

// AvailabilityZone returns a reference to field availability_zone of aws_lightsail_container_service.
func (lcs lightsailContainerServiceAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(lcs.ref.Append("availability_zone"))
}

// CreatedAt returns a reference to field created_at of aws_lightsail_container_service.
func (lcs lightsailContainerServiceAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(lcs.ref.Append("created_at"))
}

// Id returns a reference to field id of aws_lightsail_container_service.
func (lcs lightsailContainerServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lcs.ref.Append("id"))
}

// IsDisabled returns a reference to field is_disabled of aws_lightsail_container_service.
func (lcs lightsailContainerServiceAttributes) IsDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(lcs.ref.Append("is_disabled"))
}

// Name returns a reference to field name of aws_lightsail_container_service.
func (lcs lightsailContainerServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lcs.ref.Append("name"))
}

// Power returns a reference to field power of aws_lightsail_container_service.
func (lcs lightsailContainerServiceAttributes) Power() terra.StringValue {
	return terra.ReferenceAsString(lcs.ref.Append("power"))
}

// PowerId returns a reference to field power_id of aws_lightsail_container_service.
func (lcs lightsailContainerServiceAttributes) PowerId() terra.StringValue {
	return terra.ReferenceAsString(lcs.ref.Append("power_id"))
}

// PrincipalArn returns a reference to field principal_arn of aws_lightsail_container_service.
func (lcs lightsailContainerServiceAttributes) PrincipalArn() terra.StringValue {
	return terra.ReferenceAsString(lcs.ref.Append("principal_arn"))
}

// PrivateDomainName returns a reference to field private_domain_name of aws_lightsail_container_service.
func (lcs lightsailContainerServiceAttributes) PrivateDomainName() terra.StringValue {
	return terra.ReferenceAsString(lcs.ref.Append("private_domain_name"))
}

// ResourceType returns a reference to field resource_type of aws_lightsail_container_service.
func (lcs lightsailContainerServiceAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(lcs.ref.Append("resource_type"))
}

// Scale returns a reference to field scale of aws_lightsail_container_service.
func (lcs lightsailContainerServiceAttributes) Scale() terra.NumberValue {
	return terra.ReferenceAsNumber(lcs.ref.Append("scale"))
}

// State returns a reference to field state of aws_lightsail_container_service.
func (lcs lightsailContainerServiceAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(lcs.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_lightsail_container_service.
func (lcs lightsailContainerServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lcs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_lightsail_container_service.
func (lcs lightsailContainerServiceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lcs.ref.Append("tags_all"))
}

// Url returns a reference to field url of aws_lightsail_container_service.
func (lcs lightsailContainerServiceAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(lcs.ref.Append("url"))
}

func (lcs lightsailContainerServiceAttributes) PrivateRegistryAccess() terra.ListValue[lightsailcontainerservice.PrivateRegistryAccessAttributes] {
	return terra.ReferenceAsList[lightsailcontainerservice.PrivateRegistryAccessAttributes](lcs.ref.Append("private_registry_access"))
}

func (lcs lightsailContainerServiceAttributes) PublicDomainNames() terra.ListValue[lightsailcontainerservice.PublicDomainNamesAttributes] {
	return terra.ReferenceAsList[lightsailcontainerservice.PublicDomainNamesAttributes](lcs.ref.Append("public_domain_names"))
}

func (lcs lightsailContainerServiceAttributes) Timeouts() lightsailcontainerservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[lightsailcontainerservice.TimeoutsAttributes](lcs.ref.Append("timeouts"))
}

type lightsailContainerServiceState struct {
	Arn                   string                                                 `json:"arn"`
	AvailabilityZone      string                                                 `json:"availability_zone"`
	CreatedAt             string                                                 `json:"created_at"`
	Id                    string                                                 `json:"id"`
	IsDisabled            bool                                                   `json:"is_disabled"`
	Name                  string                                                 `json:"name"`
	Power                 string                                                 `json:"power"`
	PowerId               string                                                 `json:"power_id"`
	PrincipalArn          string                                                 `json:"principal_arn"`
	PrivateDomainName     string                                                 `json:"private_domain_name"`
	ResourceType          string                                                 `json:"resource_type"`
	Scale                 float64                                                `json:"scale"`
	State                 string                                                 `json:"state"`
	Tags                  map[string]string                                      `json:"tags"`
	TagsAll               map[string]string                                      `json:"tags_all"`
	Url                   string                                                 `json:"url"`
	PrivateRegistryAccess []lightsailcontainerservice.PrivateRegistryAccessState `json:"private_registry_access"`
	PublicDomainNames     []lightsailcontainerservice.PublicDomainNamesState     `json:"public_domain_names"`
	Timeouts              *lightsailcontainerservice.TimeoutsState               `json:"timeouts"`
}
