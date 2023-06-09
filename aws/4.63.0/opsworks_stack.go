// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	opsworksstack "github.com/golingon/terraproviders/aws/4.63.0/opsworksstack"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOpsworksStack creates a new instance of [OpsworksStack].
func NewOpsworksStack(name string, args OpsworksStackArgs) *OpsworksStack {
	return &OpsworksStack{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpsworksStack)(nil)

// OpsworksStack represents the Terraform resource aws_opsworks_stack.
type OpsworksStack struct {
	Name      string
	Args      OpsworksStackArgs
	state     *opsworksStackState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OpsworksStack].
func (os *OpsworksStack) Type() string {
	return "aws_opsworks_stack"
}

// LocalName returns the local name for [OpsworksStack].
func (os *OpsworksStack) LocalName() string {
	return os.Name
}

// Configuration returns the configuration (args) for [OpsworksStack].
func (os *OpsworksStack) Configuration() interface{} {
	return os.Args
}

// DependOn is used for other resources to depend on [OpsworksStack].
func (os *OpsworksStack) DependOn() terra.Reference {
	return terra.ReferenceResource(os)
}

// Dependencies returns the list of resources [OpsworksStack] depends_on.
func (os *OpsworksStack) Dependencies() terra.Dependencies {
	return os.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OpsworksStack].
func (os *OpsworksStack) LifecycleManagement() *terra.Lifecycle {
	return os.Lifecycle
}

// Attributes returns the attributes for [OpsworksStack].
func (os *OpsworksStack) Attributes() opsworksStackAttributes {
	return opsworksStackAttributes{ref: terra.ReferenceResource(os)}
}

// ImportState imports the given attribute values into [OpsworksStack]'s state.
func (os *OpsworksStack) ImportState(av io.Reader) error {
	os.state = &opsworksStackState{}
	if err := json.NewDecoder(av).Decode(os.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", os.Type(), os.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OpsworksStack] has state.
func (os *OpsworksStack) State() (*opsworksStackState, bool) {
	return os.state, os.state != nil
}

// StateMust returns the state for [OpsworksStack]. Panics if the state is nil.
func (os *OpsworksStack) StateMust() *opsworksStackState {
	if os.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", os.Type(), os.LocalName()))
	}
	return os.state
}

// OpsworksStackArgs contains the configurations for aws_opsworks_stack.
type OpsworksStackArgs struct {
	// AgentVersion: string, optional
	AgentVersion terra.StringValue `hcl:"agent_version,attr"`
	// BerkshelfVersion: string, optional
	BerkshelfVersion terra.StringValue `hcl:"berkshelf_version,attr"`
	// Color: string, optional
	Color terra.StringValue `hcl:"color,attr"`
	// ConfigurationManagerName: string, optional
	ConfigurationManagerName terra.StringValue `hcl:"configuration_manager_name,attr"`
	// ConfigurationManagerVersion: string, optional
	ConfigurationManagerVersion terra.StringValue `hcl:"configuration_manager_version,attr"`
	// CustomJson: string, optional
	CustomJson terra.StringValue `hcl:"custom_json,attr"`
	// DefaultAvailabilityZone: string, optional
	DefaultAvailabilityZone terra.StringValue `hcl:"default_availability_zone,attr"`
	// DefaultInstanceProfileArn: string, required
	DefaultInstanceProfileArn terra.StringValue `hcl:"default_instance_profile_arn,attr" validate:"required"`
	// DefaultOs: string, optional
	DefaultOs terra.StringValue `hcl:"default_os,attr"`
	// DefaultRootDeviceType: string, optional
	DefaultRootDeviceType terra.StringValue `hcl:"default_root_device_type,attr"`
	// DefaultSshKeyName: string, optional
	DefaultSshKeyName terra.StringValue `hcl:"default_ssh_key_name,attr"`
	// DefaultSubnetId: string, optional
	DefaultSubnetId terra.StringValue `hcl:"default_subnet_id,attr"`
	// HostnameTheme: string, optional
	HostnameTheme terra.StringValue `hcl:"hostname_theme,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManageBerkshelf: bool, optional
	ManageBerkshelf terra.BoolValue `hcl:"manage_berkshelf,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
	// ServiceRoleArn: string, required
	ServiceRoleArn terra.StringValue `hcl:"service_role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// UseCustomCookbooks: bool, optional
	UseCustomCookbooks terra.BoolValue `hcl:"use_custom_cookbooks,attr"`
	// UseOpsworksSecurityGroups: bool, optional
	UseOpsworksSecurityGroups terra.BoolValue `hcl:"use_opsworks_security_groups,attr"`
	// VpcId: string, optional
	VpcId terra.StringValue `hcl:"vpc_id,attr"`
	// CustomCookbooksSource: optional
	CustomCookbooksSource *opsworksstack.CustomCookbooksSource `hcl:"custom_cookbooks_source,block"`
	// Timeouts: optional
	Timeouts *opsworksstack.Timeouts `hcl:"timeouts,block"`
}
type opsworksStackAttributes struct {
	ref terra.Reference
}

// AgentVersion returns a reference to field agent_version of aws_opsworks_stack.
func (os opsworksStackAttributes) AgentVersion() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("agent_version"))
}

// Arn returns a reference to field arn of aws_opsworks_stack.
func (os opsworksStackAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("arn"))
}

// BerkshelfVersion returns a reference to field berkshelf_version of aws_opsworks_stack.
func (os opsworksStackAttributes) BerkshelfVersion() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("berkshelf_version"))
}

// Color returns a reference to field color of aws_opsworks_stack.
func (os opsworksStackAttributes) Color() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("color"))
}

// ConfigurationManagerName returns a reference to field configuration_manager_name of aws_opsworks_stack.
func (os opsworksStackAttributes) ConfigurationManagerName() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("configuration_manager_name"))
}

// ConfigurationManagerVersion returns a reference to field configuration_manager_version of aws_opsworks_stack.
func (os opsworksStackAttributes) ConfigurationManagerVersion() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("configuration_manager_version"))
}

// CustomJson returns a reference to field custom_json of aws_opsworks_stack.
func (os opsworksStackAttributes) CustomJson() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("custom_json"))
}

// DefaultAvailabilityZone returns a reference to field default_availability_zone of aws_opsworks_stack.
func (os opsworksStackAttributes) DefaultAvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("default_availability_zone"))
}

// DefaultInstanceProfileArn returns a reference to field default_instance_profile_arn of aws_opsworks_stack.
func (os opsworksStackAttributes) DefaultInstanceProfileArn() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("default_instance_profile_arn"))
}

// DefaultOs returns a reference to field default_os of aws_opsworks_stack.
func (os opsworksStackAttributes) DefaultOs() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("default_os"))
}

// DefaultRootDeviceType returns a reference to field default_root_device_type of aws_opsworks_stack.
func (os opsworksStackAttributes) DefaultRootDeviceType() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("default_root_device_type"))
}

// DefaultSshKeyName returns a reference to field default_ssh_key_name of aws_opsworks_stack.
func (os opsworksStackAttributes) DefaultSshKeyName() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("default_ssh_key_name"))
}

// DefaultSubnetId returns a reference to field default_subnet_id of aws_opsworks_stack.
func (os opsworksStackAttributes) DefaultSubnetId() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("default_subnet_id"))
}

// HostnameTheme returns a reference to field hostname_theme of aws_opsworks_stack.
func (os opsworksStackAttributes) HostnameTheme() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("hostname_theme"))
}

// Id returns a reference to field id of aws_opsworks_stack.
func (os opsworksStackAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("id"))
}

// ManageBerkshelf returns a reference to field manage_berkshelf of aws_opsworks_stack.
func (os opsworksStackAttributes) ManageBerkshelf() terra.BoolValue {
	return terra.ReferenceAsBool(os.ref.Append("manage_berkshelf"))
}

// Name returns a reference to field name of aws_opsworks_stack.
func (os opsworksStackAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("name"))
}

// Region returns a reference to field region of aws_opsworks_stack.
func (os opsworksStackAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("region"))
}

// ServiceRoleArn returns a reference to field service_role_arn of aws_opsworks_stack.
func (os opsworksStackAttributes) ServiceRoleArn() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("service_role_arn"))
}

// StackEndpoint returns a reference to field stack_endpoint of aws_opsworks_stack.
func (os opsworksStackAttributes) StackEndpoint() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("stack_endpoint"))
}

// Tags returns a reference to field tags of aws_opsworks_stack.
func (os opsworksStackAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](os.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_opsworks_stack.
func (os opsworksStackAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](os.ref.Append("tags_all"))
}

// UseCustomCookbooks returns a reference to field use_custom_cookbooks of aws_opsworks_stack.
func (os opsworksStackAttributes) UseCustomCookbooks() terra.BoolValue {
	return terra.ReferenceAsBool(os.ref.Append("use_custom_cookbooks"))
}

// UseOpsworksSecurityGroups returns a reference to field use_opsworks_security_groups of aws_opsworks_stack.
func (os opsworksStackAttributes) UseOpsworksSecurityGroups() terra.BoolValue {
	return terra.ReferenceAsBool(os.ref.Append("use_opsworks_security_groups"))
}

// VpcId returns a reference to field vpc_id of aws_opsworks_stack.
func (os opsworksStackAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("vpc_id"))
}

func (os opsworksStackAttributes) CustomCookbooksSource() terra.ListValue[opsworksstack.CustomCookbooksSourceAttributes] {
	return terra.ReferenceAsList[opsworksstack.CustomCookbooksSourceAttributes](os.ref.Append("custom_cookbooks_source"))
}

func (os opsworksStackAttributes) Timeouts() opsworksstack.TimeoutsAttributes {
	return terra.ReferenceAsSingle[opsworksstack.TimeoutsAttributes](os.ref.Append("timeouts"))
}

type opsworksStackState struct {
	AgentVersion                string                                     `json:"agent_version"`
	Arn                         string                                     `json:"arn"`
	BerkshelfVersion            string                                     `json:"berkshelf_version"`
	Color                       string                                     `json:"color"`
	ConfigurationManagerName    string                                     `json:"configuration_manager_name"`
	ConfigurationManagerVersion string                                     `json:"configuration_manager_version"`
	CustomJson                  string                                     `json:"custom_json"`
	DefaultAvailabilityZone     string                                     `json:"default_availability_zone"`
	DefaultInstanceProfileArn   string                                     `json:"default_instance_profile_arn"`
	DefaultOs                   string                                     `json:"default_os"`
	DefaultRootDeviceType       string                                     `json:"default_root_device_type"`
	DefaultSshKeyName           string                                     `json:"default_ssh_key_name"`
	DefaultSubnetId             string                                     `json:"default_subnet_id"`
	HostnameTheme               string                                     `json:"hostname_theme"`
	Id                          string                                     `json:"id"`
	ManageBerkshelf             bool                                       `json:"manage_berkshelf"`
	Name                        string                                     `json:"name"`
	Region                      string                                     `json:"region"`
	ServiceRoleArn              string                                     `json:"service_role_arn"`
	StackEndpoint               string                                     `json:"stack_endpoint"`
	Tags                        map[string]string                          `json:"tags"`
	TagsAll                     map[string]string                          `json:"tags_all"`
	UseCustomCookbooks          bool                                       `json:"use_custom_cookbooks"`
	UseOpsworksSecurityGroups   bool                                       `json:"use_opsworks_security_groups"`
	VpcId                       string                                     `json:"vpc_id"`
	CustomCookbooksSource       []opsworksstack.CustomCookbooksSourceState `json:"custom_cookbooks_source"`
	Timeouts                    *opsworksstack.TimeoutsState               `json:"timeouts"`
}
