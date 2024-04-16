// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_app_engine_application

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

// Resource represents the Terraform resource google_app_engine_application.
type Resource struct {
	Name      string
	Args      Args
	state     *googleAppEngineApplicationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gaea *Resource) Type() string {
	return "google_app_engine_application"
}

// LocalName returns the local name for [Resource].
func (gaea *Resource) LocalName() string {
	return gaea.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gaea *Resource) Configuration() interface{} {
	return gaea.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gaea *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gaea)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gaea *Resource) Dependencies() terra.Dependencies {
	return gaea.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gaea *Resource) LifecycleManagement() *terra.Lifecycle {
	return gaea.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gaea *Resource) Attributes() googleAppEngineApplicationAttributes {
	return googleAppEngineApplicationAttributes{ref: terra.ReferenceResource(gaea)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gaea *Resource) ImportState(state io.Reader) error {
	gaea.state = &googleAppEngineApplicationState{}
	if err := json.NewDecoder(state).Decode(gaea.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gaea.Type(), gaea.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gaea *Resource) State() (*googleAppEngineApplicationState, bool) {
	return gaea.state, gaea.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gaea *Resource) StateMust() *googleAppEngineApplicationState {
	if gaea.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gaea.Type(), gaea.LocalName()))
	}
	return gaea.state
}

// Args contains the configurations for google_app_engine_application.
type Args struct {
	// AuthDomain: string, optional
	AuthDomain terra.StringValue `hcl:"auth_domain,attr"`
	// DatabaseType: string, optional
	DatabaseType terra.StringValue `hcl:"database_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocationId: string, required
	LocationId terra.StringValue `hcl:"location_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ServingStatus: string, optional
	ServingStatus terra.StringValue `hcl:"serving_status,attr"`
	// FeatureSettings: optional
	FeatureSettings *FeatureSettings `hcl:"feature_settings,block"`
	// Iap: optional
	Iap *Iap `hcl:"iap,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleAppEngineApplicationAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of google_app_engine_application.
func (gaea googleAppEngineApplicationAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(gaea.ref.Append("app_id"))
}

// AuthDomain returns a reference to field auth_domain of google_app_engine_application.
func (gaea googleAppEngineApplicationAttributes) AuthDomain() terra.StringValue {
	return terra.ReferenceAsString(gaea.ref.Append("auth_domain"))
}

// CodeBucket returns a reference to field code_bucket of google_app_engine_application.
func (gaea googleAppEngineApplicationAttributes) CodeBucket() terra.StringValue {
	return terra.ReferenceAsString(gaea.ref.Append("code_bucket"))
}

// DatabaseType returns a reference to field database_type of google_app_engine_application.
func (gaea googleAppEngineApplicationAttributes) DatabaseType() terra.StringValue {
	return terra.ReferenceAsString(gaea.ref.Append("database_type"))
}

// DefaultBucket returns a reference to field default_bucket of google_app_engine_application.
func (gaea googleAppEngineApplicationAttributes) DefaultBucket() terra.StringValue {
	return terra.ReferenceAsString(gaea.ref.Append("default_bucket"))
}

// DefaultHostname returns a reference to field default_hostname of google_app_engine_application.
func (gaea googleAppEngineApplicationAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(gaea.ref.Append("default_hostname"))
}

// GcrDomain returns a reference to field gcr_domain of google_app_engine_application.
func (gaea googleAppEngineApplicationAttributes) GcrDomain() terra.StringValue {
	return terra.ReferenceAsString(gaea.ref.Append("gcr_domain"))
}

// Id returns a reference to field id of google_app_engine_application.
func (gaea googleAppEngineApplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gaea.ref.Append("id"))
}

// LocationId returns a reference to field location_id of google_app_engine_application.
func (gaea googleAppEngineApplicationAttributes) LocationId() terra.StringValue {
	return terra.ReferenceAsString(gaea.ref.Append("location_id"))
}

// Name returns a reference to field name of google_app_engine_application.
func (gaea googleAppEngineApplicationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gaea.ref.Append("name"))
}

// Project returns a reference to field project of google_app_engine_application.
func (gaea googleAppEngineApplicationAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gaea.ref.Append("project"))
}

// ServingStatus returns a reference to field serving_status of google_app_engine_application.
func (gaea googleAppEngineApplicationAttributes) ServingStatus() terra.StringValue {
	return terra.ReferenceAsString(gaea.ref.Append("serving_status"))
}

func (gaea googleAppEngineApplicationAttributes) UrlDispatchRule() terra.ListValue[UrlDispatchRuleAttributes] {
	return terra.ReferenceAsList[UrlDispatchRuleAttributes](gaea.ref.Append("url_dispatch_rule"))
}

func (gaea googleAppEngineApplicationAttributes) FeatureSettings() terra.ListValue[FeatureSettingsAttributes] {
	return terra.ReferenceAsList[FeatureSettingsAttributes](gaea.ref.Append("feature_settings"))
}

func (gaea googleAppEngineApplicationAttributes) Iap() terra.ListValue[IapAttributes] {
	return terra.ReferenceAsList[IapAttributes](gaea.ref.Append("iap"))
}

func (gaea googleAppEngineApplicationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gaea.ref.Append("timeouts"))
}

type googleAppEngineApplicationState struct {
	AppId           string                 `json:"app_id"`
	AuthDomain      string                 `json:"auth_domain"`
	CodeBucket      string                 `json:"code_bucket"`
	DatabaseType    string                 `json:"database_type"`
	DefaultBucket   string                 `json:"default_bucket"`
	DefaultHostname string                 `json:"default_hostname"`
	GcrDomain       string                 `json:"gcr_domain"`
	Id              string                 `json:"id"`
	LocationId      string                 `json:"location_id"`
	Name            string                 `json:"name"`
	Project         string                 `json:"project"`
	ServingStatus   string                 `json:"serving_status"`
	UrlDispatchRule []UrlDispatchRuleState `json:"url_dispatch_rule"`
	FeatureSettings []FeatureSettingsState `json:"feature_settings"`
	Iap             []IapState             `json:"iap"`
	Timeouts        *TimeoutsState         `json:"timeouts"`
}
