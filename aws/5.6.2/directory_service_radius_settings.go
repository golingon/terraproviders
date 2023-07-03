// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	directoryserviceradiussettings "github.com/golingon/terraproviders/aws/5.6.2/directoryserviceradiussettings"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDirectoryServiceRadiusSettings creates a new instance of [DirectoryServiceRadiusSettings].
func NewDirectoryServiceRadiusSettings(name string, args DirectoryServiceRadiusSettingsArgs) *DirectoryServiceRadiusSettings {
	return &DirectoryServiceRadiusSettings{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DirectoryServiceRadiusSettings)(nil)

// DirectoryServiceRadiusSettings represents the Terraform resource aws_directory_service_radius_settings.
type DirectoryServiceRadiusSettings struct {
	Name      string
	Args      DirectoryServiceRadiusSettingsArgs
	state     *directoryServiceRadiusSettingsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DirectoryServiceRadiusSettings].
func (dsrs *DirectoryServiceRadiusSettings) Type() string {
	return "aws_directory_service_radius_settings"
}

// LocalName returns the local name for [DirectoryServiceRadiusSettings].
func (dsrs *DirectoryServiceRadiusSettings) LocalName() string {
	return dsrs.Name
}

// Configuration returns the configuration (args) for [DirectoryServiceRadiusSettings].
func (dsrs *DirectoryServiceRadiusSettings) Configuration() interface{} {
	return dsrs.Args
}

// DependOn is used for other resources to depend on [DirectoryServiceRadiusSettings].
func (dsrs *DirectoryServiceRadiusSettings) DependOn() terra.Reference {
	return terra.ReferenceResource(dsrs)
}

// Dependencies returns the list of resources [DirectoryServiceRadiusSettings] depends_on.
func (dsrs *DirectoryServiceRadiusSettings) Dependencies() terra.Dependencies {
	return dsrs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DirectoryServiceRadiusSettings].
func (dsrs *DirectoryServiceRadiusSettings) LifecycleManagement() *terra.Lifecycle {
	return dsrs.Lifecycle
}

// Attributes returns the attributes for [DirectoryServiceRadiusSettings].
func (dsrs *DirectoryServiceRadiusSettings) Attributes() directoryServiceRadiusSettingsAttributes {
	return directoryServiceRadiusSettingsAttributes{ref: terra.ReferenceResource(dsrs)}
}

// ImportState imports the given attribute values into [DirectoryServiceRadiusSettings]'s state.
func (dsrs *DirectoryServiceRadiusSettings) ImportState(av io.Reader) error {
	dsrs.state = &directoryServiceRadiusSettingsState{}
	if err := json.NewDecoder(av).Decode(dsrs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dsrs.Type(), dsrs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DirectoryServiceRadiusSettings] has state.
func (dsrs *DirectoryServiceRadiusSettings) State() (*directoryServiceRadiusSettingsState, bool) {
	return dsrs.state, dsrs.state != nil
}

// StateMust returns the state for [DirectoryServiceRadiusSettings]. Panics if the state is nil.
func (dsrs *DirectoryServiceRadiusSettings) StateMust() *directoryServiceRadiusSettingsState {
	if dsrs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dsrs.Type(), dsrs.LocalName()))
	}
	return dsrs.state
}

// DirectoryServiceRadiusSettingsArgs contains the configurations for aws_directory_service_radius_settings.
type DirectoryServiceRadiusSettingsArgs struct {
	// AuthenticationProtocol: string, required
	AuthenticationProtocol terra.StringValue `hcl:"authentication_protocol,attr" validate:"required"`
	// DirectoryId: string, required
	DirectoryId terra.StringValue `hcl:"directory_id,attr" validate:"required"`
	// DisplayLabel: string, required
	DisplayLabel terra.StringValue `hcl:"display_label,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RadiusPort: number, required
	RadiusPort terra.NumberValue `hcl:"radius_port,attr" validate:"required"`
	// RadiusRetries: number, required
	RadiusRetries terra.NumberValue `hcl:"radius_retries,attr" validate:"required"`
	// RadiusServers: set of string, required
	RadiusServers terra.SetValue[terra.StringValue] `hcl:"radius_servers,attr" validate:"required"`
	// RadiusTimeout: number, required
	RadiusTimeout terra.NumberValue `hcl:"radius_timeout,attr" validate:"required"`
	// SharedSecret: string, required
	SharedSecret terra.StringValue `hcl:"shared_secret,attr" validate:"required"`
	// UseSameUsername: bool, optional
	UseSameUsername terra.BoolValue `hcl:"use_same_username,attr"`
	// Timeouts: optional
	Timeouts *directoryserviceradiussettings.Timeouts `hcl:"timeouts,block"`
}
type directoryServiceRadiusSettingsAttributes struct {
	ref terra.Reference
}

// AuthenticationProtocol returns a reference to field authentication_protocol of aws_directory_service_radius_settings.
func (dsrs directoryServiceRadiusSettingsAttributes) AuthenticationProtocol() terra.StringValue {
	return terra.ReferenceAsString(dsrs.ref.Append("authentication_protocol"))
}

// DirectoryId returns a reference to field directory_id of aws_directory_service_radius_settings.
func (dsrs directoryServiceRadiusSettingsAttributes) DirectoryId() terra.StringValue {
	return terra.ReferenceAsString(dsrs.ref.Append("directory_id"))
}

// DisplayLabel returns a reference to field display_label of aws_directory_service_radius_settings.
func (dsrs directoryServiceRadiusSettingsAttributes) DisplayLabel() terra.StringValue {
	return terra.ReferenceAsString(dsrs.ref.Append("display_label"))
}

// Id returns a reference to field id of aws_directory_service_radius_settings.
func (dsrs directoryServiceRadiusSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsrs.ref.Append("id"))
}

// RadiusPort returns a reference to field radius_port of aws_directory_service_radius_settings.
func (dsrs directoryServiceRadiusSettingsAttributes) RadiusPort() terra.NumberValue {
	return terra.ReferenceAsNumber(dsrs.ref.Append("radius_port"))
}

// RadiusRetries returns a reference to field radius_retries of aws_directory_service_radius_settings.
func (dsrs directoryServiceRadiusSettingsAttributes) RadiusRetries() terra.NumberValue {
	return terra.ReferenceAsNumber(dsrs.ref.Append("radius_retries"))
}

// RadiusServers returns a reference to field radius_servers of aws_directory_service_radius_settings.
func (dsrs directoryServiceRadiusSettingsAttributes) RadiusServers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dsrs.ref.Append("radius_servers"))
}

// RadiusTimeout returns a reference to field radius_timeout of aws_directory_service_radius_settings.
func (dsrs directoryServiceRadiusSettingsAttributes) RadiusTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(dsrs.ref.Append("radius_timeout"))
}

// SharedSecret returns a reference to field shared_secret of aws_directory_service_radius_settings.
func (dsrs directoryServiceRadiusSettingsAttributes) SharedSecret() terra.StringValue {
	return terra.ReferenceAsString(dsrs.ref.Append("shared_secret"))
}

// UseSameUsername returns a reference to field use_same_username of aws_directory_service_radius_settings.
func (dsrs directoryServiceRadiusSettingsAttributes) UseSameUsername() terra.BoolValue {
	return terra.ReferenceAsBool(dsrs.ref.Append("use_same_username"))
}

func (dsrs directoryServiceRadiusSettingsAttributes) Timeouts() directoryserviceradiussettings.TimeoutsAttributes {
	return terra.ReferenceAsSingle[directoryserviceradiussettings.TimeoutsAttributes](dsrs.ref.Append("timeouts"))
}

type directoryServiceRadiusSettingsState struct {
	AuthenticationProtocol string                                        `json:"authentication_protocol"`
	DirectoryId            string                                        `json:"directory_id"`
	DisplayLabel           string                                        `json:"display_label"`
	Id                     string                                        `json:"id"`
	RadiusPort             float64                                       `json:"radius_port"`
	RadiusRetries          float64                                       `json:"radius_retries"`
	RadiusServers          []string                                      `json:"radius_servers"`
	RadiusTimeout          float64                                       `json:"radius_timeout"`
	SharedSecret           string                                        `json:"shared_secret"`
	UseSameUsername        bool                                          `json:"use_same_username"`
	Timeouts               *directoryserviceradiussettings.TimeoutsState `json:"timeouts"`
}
