// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	appengineapplication "github.com/golingon/terraproviders/googlebeta/4.71.0/appengineapplication"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppEngineApplication creates a new instance of [AppEngineApplication].
func NewAppEngineApplication(name string, args AppEngineApplicationArgs) *AppEngineApplication {
	return &AppEngineApplication{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppEngineApplication)(nil)

// AppEngineApplication represents the Terraform resource google_app_engine_application.
type AppEngineApplication struct {
	Name      string
	Args      AppEngineApplicationArgs
	state     *appEngineApplicationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppEngineApplication].
func (aea *AppEngineApplication) Type() string {
	return "google_app_engine_application"
}

// LocalName returns the local name for [AppEngineApplication].
func (aea *AppEngineApplication) LocalName() string {
	return aea.Name
}

// Configuration returns the configuration (args) for [AppEngineApplication].
func (aea *AppEngineApplication) Configuration() interface{} {
	return aea.Args
}

// DependOn is used for other resources to depend on [AppEngineApplication].
func (aea *AppEngineApplication) DependOn() terra.Reference {
	return terra.ReferenceResource(aea)
}

// Dependencies returns the list of resources [AppEngineApplication] depends_on.
func (aea *AppEngineApplication) Dependencies() terra.Dependencies {
	return aea.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppEngineApplication].
func (aea *AppEngineApplication) LifecycleManagement() *terra.Lifecycle {
	return aea.Lifecycle
}

// Attributes returns the attributes for [AppEngineApplication].
func (aea *AppEngineApplication) Attributes() appEngineApplicationAttributes {
	return appEngineApplicationAttributes{ref: terra.ReferenceResource(aea)}
}

// ImportState imports the given attribute values into [AppEngineApplication]'s state.
func (aea *AppEngineApplication) ImportState(av io.Reader) error {
	aea.state = &appEngineApplicationState{}
	if err := json.NewDecoder(av).Decode(aea.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aea.Type(), aea.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppEngineApplication] has state.
func (aea *AppEngineApplication) State() (*appEngineApplicationState, bool) {
	return aea.state, aea.state != nil
}

// StateMust returns the state for [AppEngineApplication]. Panics if the state is nil.
func (aea *AppEngineApplication) StateMust() *appEngineApplicationState {
	if aea.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aea.Type(), aea.LocalName()))
	}
	return aea.state
}

// AppEngineApplicationArgs contains the configurations for google_app_engine_application.
type AppEngineApplicationArgs struct {
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
	// UrlDispatchRule: min=0
	UrlDispatchRule []appengineapplication.UrlDispatchRule `hcl:"url_dispatch_rule,block" validate:"min=0"`
	// FeatureSettings: optional
	FeatureSettings *appengineapplication.FeatureSettings `hcl:"feature_settings,block"`
	// Iap: optional
	Iap *appengineapplication.Iap `hcl:"iap,block"`
	// Timeouts: optional
	Timeouts *appengineapplication.Timeouts `hcl:"timeouts,block"`
}
type appEngineApplicationAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of google_app_engine_application.
func (aea appEngineApplicationAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("app_id"))
}

// AuthDomain returns a reference to field auth_domain of google_app_engine_application.
func (aea appEngineApplicationAttributes) AuthDomain() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("auth_domain"))
}

// CodeBucket returns a reference to field code_bucket of google_app_engine_application.
func (aea appEngineApplicationAttributes) CodeBucket() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("code_bucket"))
}

// DatabaseType returns a reference to field database_type of google_app_engine_application.
func (aea appEngineApplicationAttributes) DatabaseType() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("database_type"))
}

// DefaultBucket returns a reference to field default_bucket of google_app_engine_application.
func (aea appEngineApplicationAttributes) DefaultBucket() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("default_bucket"))
}

// DefaultHostname returns a reference to field default_hostname of google_app_engine_application.
func (aea appEngineApplicationAttributes) DefaultHostname() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("default_hostname"))
}

// GcrDomain returns a reference to field gcr_domain of google_app_engine_application.
func (aea appEngineApplicationAttributes) GcrDomain() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("gcr_domain"))
}

// Id returns a reference to field id of google_app_engine_application.
func (aea appEngineApplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("id"))
}

// LocationId returns a reference to field location_id of google_app_engine_application.
func (aea appEngineApplicationAttributes) LocationId() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("location_id"))
}

// Name returns a reference to field name of google_app_engine_application.
func (aea appEngineApplicationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("name"))
}

// Project returns a reference to field project of google_app_engine_application.
func (aea appEngineApplicationAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("project"))
}

// ServingStatus returns a reference to field serving_status of google_app_engine_application.
func (aea appEngineApplicationAttributes) ServingStatus() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("serving_status"))
}

func (aea appEngineApplicationAttributes) UrlDispatchRule() terra.ListValue[appengineapplication.UrlDispatchRuleAttributes] {
	return terra.ReferenceAsList[appengineapplication.UrlDispatchRuleAttributes](aea.ref.Append("url_dispatch_rule"))
}

func (aea appEngineApplicationAttributes) FeatureSettings() terra.ListValue[appengineapplication.FeatureSettingsAttributes] {
	return terra.ReferenceAsList[appengineapplication.FeatureSettingsAttributes](aea.ref.Append("feature_settings"))
}

func (aea appEngineApplicationAttributes) Iap() terra.ListValue[appengineapplication.IapAttributes] {
	return terra.ReferenceAsList[appengineapplication.IapAttributes](aea.ref.Append("iap"))
}

func (aea appEngineApplicationAttributes) Timeouts() appengineapplication.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appengineapplication.TimeoutsAttributes](aea.ref.Append("timeouts"))
}

type appEngineApplicationState struct {
	AppId           string                                      `json:"app_id"`
	AuthDomain      string                                      `json:"auth_domain"`
	CodeBucket      string                                      `json:"code_bucket"`
	DatabaseType    string                                      `json:"database_type"`
	DefaultBucket   string                                      `json:"default_bucket"`
	DefaultHostname string                                      `json:"default_hostname"`
	GcrDomain       string                                      `json:"gcr_domain"`
	Id              string                                      `json:"id"`
	LocationId      string                                      `json:"location_id"`
	Name            string                                      `json:"name"`
	Project         string                                      `json:"project"`
	ServingStatus   string                                      `json:"serving_status"`
	UrlDispatchRule []appengineapplication.UrlDispatchRuleState `json:"url_dispatch_rule"`
	FeatureSettings []appengineapplication.FeatureSettingsState `json:"feature_settings"`
	Iap             []appengineapplication.IapState             `json:"iap"`
	Timeouts        *appengineapplication.TimeoutsState         `json:"timeouts"`
}
