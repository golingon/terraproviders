// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	appengineapplicationurldispatchrules "github.com/golingon/terraproviders/googlebeta/4.64.0/appengineapplicationurldispatchrules"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppEngineApplicationUrlDispatchRules creates a new instance of [AppEngineApplicationUrlDispatchRules].
func NewAppEngineApplicationUrlDispatchRules(name string, args AppEngineApplicationUrlDispatchRulesArgs) *AppEngineApplicationUrlDispatchRules {
	return &AppEngineApplicationUrlDispatchRules{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppEngineApplicationUrlDispatchRules)(nil)

// AppEngineApplicationUrlDispatchRules represents the Terraform resource google_app_engine_application_url_dispatch_rules.
type AppEngineApplicationUrlDispatchRules struct {
	Name      string
	Args      AppEngineApplicationUrlDispatchRulesArgs
	state     *appEngineApplicationUrlDispatchRulesState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppEngineApplicationUrlDispatchRules].
func (aeaudr *AppEngineApplicationUrlDispatchRules) Type() string {
	return "google_app_engine_application_url_dispatch_rules"
}

// LocalName returns the local name for [AppEngineApplicationUrlDispatchRules].
func (aeaudr *AppEngineApplicationUrlDispatchRules) LocalName() string {
	return aeaudr.Name
}

// Configuration returns the configuration (args) for [AppEngineApplicationUrlDispatchRules].
func (aeaudr *AppEngineApplicationUrlDispatchRules) Configuration() interface{} {
	return aeaudr.Args
}

// DependOn is used for other resources to depend on [AppEngineApplicationUrlDispatchRules].
func (aeaudr *AppEngineApplicationUrlDispatchRules) DependOn() terra.Reference {
	return terra.ReferenceResource(aeaudr)
}

// Dependencies returns the list of resources [AppEngineApplicationUrlDispatchRules] depends_on.
func (aeaudr *AppEngineApplicationUrlDispatchRules) Dependencies() terra.Dependencies {
	return aeaudr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppEngineApplicationUrlDispatchRules].
func (aeaudr *AppEngineApplicationUrlDispatchRules) LifecycleManagement() *terra.Lifecycle {
	return aeaudr.Lifecycle
}

// Attributes returns the attributes for [AppEngineApplicationUrlDispatchRules].
func (aeaudr *AppEngineApplicationUrlDispatchRules) Attributes() appEngineApplicationUrlDispatchRulesAttributes {
	return appEngineApplicationUrlDispatchRulesAttributes{ref: terra.ReferenceResource(aeaudr)}
}

// ImportState imports the given attribute values into [AppEngineApplicationUrlDispatchRules]'s state.
func (aeaudr *AppEngineApplicationUrlDispatchRules) ImportState(av io.Reader) error {
	aeaudr.state = &appEngineApplicationUrlDispatchRulesState{}
	if err := json.NewDecoder(av).Decode(aeaudr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aeaudr.Type(), aeaudr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppEngineApplicationUrlDispatchRules] has state.
func (aeaudr *AppEngineApplicationUrlDispatchRules) State() (*appEngineApplicationUrlDispatchRulesState, bool) {
	return aeaudr.state, aeaudr.state != nil
}

// StateMust returns the state for [AppEngineApplicationUrlDispatchRules]. Panics if the state is nil.
func (aeaudr *AppEngineApplicationUrlDispatchRules) StateMust() *appEngineApplicationUrlDispatchRulesState {
	if aeaudr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aeaudr.Type(), aeaudr.LocalName()))
	}
	return aeaudr.state
}

// AppEngineApplicationUrlDispatchRulesArgs contains the configurations for google_app_engine_application_url_dispatch_rules.
type AppEngineApplicationUrlDispatchRulesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// DispatchRules: min=1
	DispatchRules []appengineapplicationurldispatchrules.DispatchRules `hcl:"dispatch_rules,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *appengineapplicationurldispatchrules.Timeouts `hcl:"timeouts,block"`
}
type appEngineApplicationUrlDispatchRulesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_app_engine_application_url_dispatch_rules.
func (aeaudr appEngineApplicationUrlDispatchRulesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aeaudr.ref.Append("id"))
}

// Project returns a reference to field project of google_app_engine_application_url_dispatch_rules.
func (aeaudr appEngineApplicationUrlDispatchRulesAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(aeaudr.ref.Append("project"))
}

func (aeaudr appEngineApplicationUrlDispatchRulesAttributes) DispatchRules() terra.ListValue[appengineapplicationurldispatchrules.DispatchRulesAttributes] {
	return terra.ReferenceAsList[appengineapplicationurldispatchrules.DispatchRulesAttributes](aeaudr.ref.Append("dispatch_rules"))
}

func (aeaudr appEngineApplicationUrlDispatchRulesAttributes) Timeouts() appengineapplicationurldispatchrules.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appengineapplicationurldispatchrules.TimeoutsAttributes](aeaudr.ref.Append("timeouts"))
}

type appEngineApplicationUrlDispatchRulesState struct {
	Id            string                                                    `json:"id"`
	Project       string                                                    `json:"project"`
	DispatchRules []appengineapplicationurldispatchrules.DispatchRulesState `json:"dispatch_rules"`
	Timeouts      *appengineapplicationurldispatchrules.TimeoutsState       `json:"timeouts"`
}
