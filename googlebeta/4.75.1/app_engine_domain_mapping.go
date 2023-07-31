// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	appenginedomainmapping "github.com/golingon/terraproviders/googlebeta/4.75.1/appenginedomainmapping"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppEngineDomainMapping creates a new instance of [AppEngineDomainMapping].
func NewAppEngineDomainMapping(name string, args AppEngineDomainMappingArgs) *AppEngineDomainMapping {
	return &AppEngineDomainMapping{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppEngineDomainMapping)(nil)

// AppEngineDomainMapping represents the Terraform resource google_app_engine_domain_mapping.
type AppEngineDomainMapping struct {
	Name      string
	Args      AppEngineDomainMappingArgs
	state     *appEngineDomainMappingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppEngineDomainMapping].
func (aedm *AppEngineDomainMapping) Type() string {
	return "google_app_engine_domain_mapping"
}

// LocalName returns the local name for [AppEngineDomainMapping].
func (aedm *AppEngineDomainMapping) LocalName() string {
	return aedm.Name
}

// Configuration returns the configuration (args) for [AppEngineDomainMapping].
func (aedm *AppEngineDomainMapping) Configuration() interface{} {
	return aedm.Args
}

// DependOn is used for other resources to depend on [AppEngineDomainMapping].
func (aedm *AppEngineDomainMapping) DependOn() terra.Reference {
	return terra.ReferenceResource(aedm)
}

// Dependencies returns the list of resources [AppEngineDomainMapping] depends_on.
func (aedm *AppEngineDomainMapping) Dependencies() terra.Dependencies {
	return aedm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppEngineDomainMapping].
func (aedm *AppEngineDomainMapping) LifecycleManagement() *terra.Lifecycle {
	return aedm.Lifecycle
}

// Attributes returns the attributes for [AppEngineDomainMapping].
func (aedm *AppEngineDomainMapping) Attributes() appEngineDomainMappingAttributes {
	return appEngineDomainMappingAttributes{ref: terra.ReferenceResource(aedm)}
}

// ImportState imports the given attribute values into [AppEngineDomainMapping]'s state.
func (aedm *AppEngineDomainMapping) ImportState(av io.Reader) error {
	aedm.state = &appEngineDomainMappingState{}
	if err := json.NewDecoder(av).Decode(aedm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aedm.Type(), aedm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppEngineDomainMapping] has state.
func (aedm *AppEngineDomainMapping) State() (*appEngineDomainMappingState, bool) {
	return aedm.state, aedm.state != nil
}

// StateMust returns the state for [AppEngineDomainMapping]. Panics if the state is nil.
func (aedm *AppEngineDomainMapping) StateMust() *appEngineDomainMappingState {
	if aedm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aedm.Type(), aedm.LocalName()))
	}
	return aedm.state
}

// AppEngineDomainMappingArgs contains the configurations for google_app_engine_domain_mapping.
type AppEngineDomainMappingArgs struct {
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OverrideStrategy: string, optional
	OverrideStrategy terra.StringValue `hcl:"override_strategy,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ResourceRecords: min=0
	ResourceRecords []appenginedomainmapping.ResourceRecords `hcl:"resource_records,block" validate:"min=0"`
	// SslSettings: optional
	SslSettings *appenginedomainmapping.SslSettings `hcl:"ssl_settings,block"`
	// Timeouts: optional
	Timeouts *appenginedomainmapping.Timeouts `hcl:"timeouts,block"`
}
type appEngineDomainMappingAttributes struct {
	ref terra.Reference
}

// DomainName returns a reference to field domain_name of google_app_engine_domain_mapping.
func (aedm appEngineDomainMappingAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(aedm.ref.Append("domain_name"))
}

// Id returns a reference to field id of google_app_engine_domain_mapping.
func (aedm appEngineDomainMappingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aedm.ref.Append("id"))
}

// Name returns a reference to field name of google_app_engine_domain_mapping.
func (aedm appEngineDomainMappingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aedm.ref.Append("name"))
}

// OverrideStrategy returns a reference to field override_strategy of google_app_engine_domain_mapping.
func (aedm appEngineDomainMappingAttributes) OverrideStrategy() terra.StringValue {
	return terra.ReferenceAsString(aedm.ref.Append("override_strategy"))
}

// Project returns a reference to field project of google_app_engine_domain_mapping.
func (aedm appEngineDomainMappingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(aedm.ref.Append("project"))
}

func (aedm appEngineDomainMappingAttributes) ResourceRecords() terra.ListValue[appenginedomainmapping.ResourceRecordsAttributes] {
	return terra.ReferenceAsList[appenginedomainmapping.ResourceRecordsAttributes](aedm.ref.Append("resource_records"))
}

func (aedm appEngineDomainMappingAttributes) SslSettings() terra.ListValue[appenginedomainmapping.SslSettingsAttributes] {
	return terra.ReferenceAsList[appenginedomainmapping.SslSettingsAttributes](aedm.ref.Append("ssl_settings"))
}

func (aedm appEngineDomainMappingAttributes) Timeouts() appenginedomainmapping.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appenginedomainmapping.TimeoutsAttributes](aedm.ref.Append("timeouts"))
}

type appEngineDomainMappingState struct {
	DomainName       string                                        `json:"domain_name"`
	Id               string                                        `json:"id"`
	Name             string                                        `json:"name"`
	OverrideStrategy string                                        `json:"override_strategy"`
	Project          string                                        `json:"project"`
	ResourceRecords  []appenginedomainmapping.ResourceRecordsState `json:"resource_records"`
	SslSettings      []appenginedomainmapping.SslSettingsState     `json:"ssl_settings"`
	Timeouts         *appenginedomainmapping.TimeoutsState         `json:"timeouts"`
}
