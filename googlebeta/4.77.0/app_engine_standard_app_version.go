// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	appenginestandardappversion "github.com/golingon/terraproviders/googlebeta/4.77.0/appenginestandardappversion"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppEngineStandardAppVersion creates a new instance of [AppEngineStandardAppVersion].
func NewAppEngineStandardAppVersion(name string, args AppEngineStandardAppVersionArgs) *AppEngineStandardAppVersion {
	return &AppEngineStandardAppVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppEngineStandardAppVersion)(nil)

// AppEngineStandardAppVersion represents the Terraform resource google_app_engine_standard_app_version.
type AppEngineStandardAppVersion struct {
	Name      string
	Args      AppEngineStandardAppVersionArgs
	state     *appEngineStandardAppVersionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppEngineStandardAppVersion].
func (aesav *AppEngineStandardAppVersion) Type() string {
	return "google_app_engine_standard_app_version"
}

// LocalName returns the local name for [AppEngineStandardAppVersion].
func (aesav *AppEngineStandardAppVersion) LocalName() string {
	return aesav.Name
}

// Configuration returns the configuration (args) for [AppEngineStandardAppVersion].
func (aesav *AppEngineStandardAppVersion) Configuration() interface{} {
	return aesav.Args
}

// DependOn is used for other resources to depend on [AppEngineStandardAppVersion].
func (aesav *AppEngineStandardAppVersion) DependOn() terra.Reference {
	return terra.ReferenceResource(aesav)
}

// Dependencies returns the list of resources [AppEngineStandardAppVersion] depends_on.
func (aesav *AppEngineStandardAppVersion) Dependencies() terra.Dependencies {
	return aesav.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppEngineStandardAppVersion].
func (aesav *AppEngineStandardAppVersion) LifecycleManagement() *terra.Lifecycle {
	return aesav.Lifecycle
}

// Attributes returns the attributes for [AppEngineStandardAppVersion].
func (aesav *AppEngineStandardAppVersion) Attributes() appEngineStandardAppVersionAttributes {
	return appEngineStandardAppVersionAttributes{ref: terra.ReferenceResource(aesav)}
}

// ImportState imports the given attribute values into [AppEngineStandardAppVersion]'s state.
func (aesav *AppEngineStandardAppVersion) ImportState(av io.Reader) error {
	aesav.state = &appEngineStandardAppVersionState{}
	if err := json.NewDecoder(av).Decode(aesav.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aesav.Type(), aesav.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppEngineStandardAppVersion] has state.
func (aesav *AppEngineStandardAppVersion) State() (*appEngineStandardAppVersionState, bool) {
	return aesav.state, aesav.state != nil
}

// StateMust returns the state for [AppEngineStandardAppVersion]. Panics if the state is nil.
func (aesav *AppEngineStandardAppVersion) StateMust() *appEngineStandardAppVersionState {
	if aesav.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aesav.Type(), aesav.LocalName()))
	}
	return aesav.state
}

// AppEngineStandardAppVersionArgs contains the configurations for google_app_engine_standard_app_version.
type AppEngineStandardAppVersionArgs struct {
	// AppEngineApis: bool, optional
	AppEngineApis terra.BoolValue `hcl:"app_engine_apis,attr"`
	// DeleteServiceOnDestroy: bool, optional
	DeleteServiceOnDestroy terra.BoolValue `hcl:"delete_service_on_destroy,attr"`
	// EnvVariables: map of string, optional
	EnvVariables terra.MapValue[terra.StringValue] `hcl:"env_variables,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InboundServices: set of string, optional
	InboundServices terra.SetValue[terra.StringValue] `hcl:"inbound_services,attr"`
	// InstanceClass: string, optional
	InstanceClass terra.StringValue `hcl:"instance_class,attr"`
	// NoopOnDestroy: bool, optional
	NoopOnDestroy terra.BoolValue `hcl:"noop_on_destroy,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Runtime: string, required
	Runtime terra.StringValue `hcl:"runtime,attr" validate:"required"`
	// RuntimeApiVersion: string, optional
	RuntimeApiVersion terra.StringValue `hcl:"runtime_api_version,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// ServiceAccount: string, optional
	ServiceAccount terra.StringValue `hcl:"service_account,attr"`
	// Threadsafe: bool, optional
	Threadsafe terra.BoolValue `hcl:"threadsafe,attr"`
	// VersionId: string, optional
	VersionId terra.StringValue `hcl:"version_id,attr"`
	// AutomaticScaling: optional
	AutomaticScaling *appenginestandardappversion.AutomaticScaling `hcl:"automatic_scaling,block"`
	// BasicScaling: optional
	BasicScaling *appenginestandardappversion.BasicScaling `hcl:"basic_scaling,block"`
	// Deployment: required
	Deployment *appenginestandardappversion.Deployment `hcl:"deployment,block" validate:"required"`
	// Entrypoint: required
	Entrypoint *appenginestandardappversion.Entrypoint `hcl:"entrypoint,block" validate:"required"`
	// Handlers: min=0
	Handlers []appenginestandardappversion.Handlers `hcl:"handlers,block" validate:"min=0"`
	// Libraries: min=0
	Libraries []appenginestandardappversion.Libraries `hcl:"libraries,block" validate:"min=0"`
	// ManualScaling: optional
	ManualScaling *appenginestandardappversion.ManualScaling `hcl:"manual_scaling,block"`
	// Timeouts: optional
	Timeouts *appenginestandardappversion.Timeouts `hcl:"timeouts,block"`
	// VpcAccessConnector: optional
	VpcAccessConnector *appenginestandardappversion.VpcAccessConnector `hcl:"vpc_access_connector,block"`
}
type appEngineStandardAppVersionAttributes struct {
	ref terra.Reference
}

// AppEngineApis returns a reference to field app_engine_apis of google_app_engine_standard_app_version.
func (aesav appEngineStandardAppVersionAttributes) AppEngineApis() terra.BoolValue {
	return terra.ReferenceAsBool(aesav.ref.Append("app_engine_apis"))
}

// DeleteServiceOnDestroy returns a reference to field delete_service_on_destroy of google_app_engine_standard_app_version.
func (aesav appEngineStandardAppVersionAttributes) DeleteServiceOnDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(aesav.ref.Append("delete_service_on_destroy"))
}

// EnvVariables returns a reference to field env_variables of google_app_engine_standard_app_version.
func (aesav appEngineStandardAppVersionAttributes) EnvVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aesav.ref.Append("env_variables"))
}

// Id returns a reference to field id of google_app_engine_standard_app_version.
func (aesav appEngineStandardAppVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aesav.ref.Append("id"))
}

// InboundServices returns a reference to field inbound_services of google_app_engine_standard_app_version.
func (aesav appEngineStandardAppVersionAttributes) InboundServices() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aesav.ref.Append("inbound_services"))
}

// InstanceClass returns a reference to field instance_class of google_app_engine_standard_app_version.
func (aesav appEngineStandardAppVersionAttributes) InstanceClass() terra.StringValue {
	return terra.ReferenceAsString(aesav.ref.Append("instance_class"))
}

// Name returns a reference to field name of google_app_engine_standard_app_version.
func (aesav appEngineStandardAppVersionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aesav.ref.Append("name"))
}

// NoopOnDestroy returns a reference to field noop_on_destroy of google_app_engine_standard_app_version.
func (aesav appEngineStandardAppVersionAttributes) NoopOnDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(aesav.ref.Append("noop_on_destroy"))
}

// Project returns a reference to field project of google_app_engine_standard_app_version.
func (aesav appEngineStandardAppVersionAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(aesav.ref.Append("project"))
}

// Runtime returns a reference to field runtime of google_app_engine_standard_app_version.
func (aesav appEngineStandardAppVersionAttributes) Runtime() terra.StringValue {
	return terra.ReferenceAsString(aesav.ref.Append("runtime"))
}

// RuntimeApiVersion returns a reference to field runtime_api_version of google_app_engine_standard_app_version.
func (aesav appEngineStandardAppVersionAttributes) RuntimeApiVersion() terra.StringValue {
	return terra.ReferenceAsString(aesav.ref.Append("runtime_api_version"))
}

// Service returns a reference to field service of google_app_engine_standard_app_version.
func (aesav appEngineStandardAppVersionAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(aesav.ref.Append("service"))
}

// ServiceAccount returns a reference to field service_account of google_app_engine_standard_app_version.
func (aesav appEngineStandardAppVersionAttributes) ServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(aesav.ref.Append("service_account"))
}

// Threadsafe returns a reference to field threadsafe of google_app_engine_standard_app_version.
func (aesav appEngineStandardAppVersionAttributes) Threadsafe() terra.BoolValue {
	return terra.ReferenceAsBool(aesav.ref.Append("threadsafe"))
}

// VersionId returns a reference to field version_id of google_app_engine_standard_app_version.
func (aesav appEngineStandardAppVersionAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(aesav.ref.Append("version_id"))
}

func (aesav appEngineStandardAppVersionAttributes) AutomaticScaling() terra.ListValue[appenginestandardappversion.AutomaticScalingAttributes] {
	return terra.ReferenceAsList[appenginestandardappversion.AutomaticScalingAttributes](aesav.ref.Append("automatic_scaling"))
}

func (aesav appEngineStandardAppVersionAttributes) BasicScaling() terra.ListValue[appenginestandardappversion.BasicScalingAttributes] {
	return terra.ReferenceAsList[appenginestandardappversion.BasicScalingAttributes](aesav.ref.Append("basic_scaling"))
}

func (aesav appEngineStandardAppVersionAttributes) Deployment() terra.ListValue[appenginestandardappversion.DeploymentAttributes] {
	return terra.ReferenceAsList[appenginestandardappversion.DeploymentAttributes](aesav.ref.Append("deployment"))
}

func (aesav appEngineStandardAppVersionAttributes) Entrypoint() terra.ListValue[appenginestandardappversion.EntrypointAttributes] {
	return terra.ReferenceAsList[appenginestandardappversion.EntrypointAttributes](aesav.ref.Append("entrypoint"))
}

func (aesav appEngineStandardAppVersionAttributes) Handlers() terra.ListValue[appenginestandardappversion.HandlersAttributes] {
	return terra.ReferenceAsList[appenginestandardappversion.HandlersAttributes](aesav.ref.Append("handlers"))
}

func (aesav appEngineStandardAppVersionAttributes) Libraries() terra.ListValue[appenginestandardappversion.LibrariesAttributes] {
	return terra.ReferenceAsList[appenginestandardappversion.LibrariesAttributes](aesav.ref.Append("libraries"))
}

func (aesav appEngineStandardAppVersionAttributes) ManualScaling() terra.ListValue[appenginestandardappversion.ManualScalingAttributes] {
	return terra.ReferenceAsList[appenginestandardappversion.ManualScalingAttributes](aesav.ref.Append("manual_scaling"))
}

func (aesav appEngineStandardAppVersionAttributes) Timeouts() appenginestandardappversion.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appenginestandardappversion.TimeoutsAttributes](aesav.ref.Append("timeouts"))
}

func (aesav appEngineStandardAppVersionAttributes) VpcAccessConnector() terra.ListValue[appenginestandardappversion.VpcAccessConnectorAttributes] {
	return terra.ReferenceAsList[appenginestandardappversion.VpcAccessConnectorAttributes](aesav.ref.Append("vpc_access_connector"))
}

type appEngineStandardAppVersionState struct {
	AppEngineApis          bool                                                  `json:"app_engine_apis"`
	DeleteServiceOnDestroy bool                                                  `json:"delete_service_on_destroy"`
	EnvVariables           map[string]string                                     `json:"env_variables"`
	Id                     string                                                `json:"id"`
	InboundServices        []string                                              `json:"inbound_services"`
	InstanceClass          string                                                `json:"instance_class"`
	Name                   string                                                `json:"name"`
	NoopOnDestroy          bool                                                  `json:"noop_on_destroy"`
	Project                string                                                `json:"project"`
	Runtime                string                                                `json:"runtime"`
	RuntimeApiVersion      string                                                `json:"runtime_api_version"`
	Service                string                                                `json:"service"`
	ServiceAccount         string                                                `json:"service_account"`
	Threadsafe             bool                                                  `json:"threadsafe"`
	VersionId              string                                                `json:"version_id"`
	AutomaticScaling       []appenginestandardappversion.AutomaticScalingState   `json:"automatic_scaling"`
	BasicScaling           []appenginestandardappversion.BasicScalingState       `json:"basic_scaling"`
	Deployment             []appenginestandardappversion.DeploymentState         `json:"deployment"`
	Entrypoint             []appenginestandardappversion.EntrypointState         `json:"entrypoint"`
	Handlers               []appenginestandardappversion.HandlersState           `json:"handlers"`
	Libraries              []appenginestandardappversion.LibrariesState          `json:"libraries"`
	ManualScaling          []appenginestandardappversion.ManualScalingState      `json:"manual_scaling"`
	Timeouts               *appenginestandardappversion.TimeoutsState            `json:"timeouts"`
	VpcAccessConnector     []appenginestandardappversion.VpcAccessConnectorState `json:"vpc_access_connector"`
}
