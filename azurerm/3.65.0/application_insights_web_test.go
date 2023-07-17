// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	applicationinsightswebtest "github.com/golingon/terraproviders/azurerm/3.65.0/applicationinsightswebtest"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApplicationInsightsWebTest creates a new instance of [ApplicationInsightsWebTest].
func NewApplicationInsightsWebTest(name string, args ApplicationInsightsWebTestArgs) *ApplicationInsightsWebTest {
	return &ApplicationInsightsWebTest{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApplicationInsightsWebTest)(nil)

// ApplicationInsightsWebTest represents the Terraform resource azurerm_application_insights_web_test.
type ApplicationInsightsWebTest struct {
	Name      string
	Args      ApplicationInsightsWebTestArgs
	state     *applicationInsightsWebTestState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApplicationInsightsWebTest].
func (aiwt *ApplicationInsightsWebTest) Type() string {
	return "azurerm_application_insights_web_test"
}

// LocalName returns the local name for [ApplicationInsightsWebTest].
func (aiwt *ApplicationInsightsWebTest) LocalName() string {
	return aiwt.Name
}

// Configuration returns the configuration (args) for [ApplicationInsightsWebTest].
func (aiwt *ApplicationInsightsWebTest) Configuration() interface{} {
	return aiwt.Args
}

// DependOn is used for other resources to depend on [ApplicationInsightsWebTest].
func (aiwt *ApplicationInsightsWebTest) DependOn() terra.Reference {
	return terra.ReferenceResource(aiwt)
}

// Dependencies returns the list of resources [ApplicationInsightsWebTest] depends_on.
func (aiwt *ApplicationInsightsWebTest) Dependencies() terra.Dependencies {
	return aiwt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApplicationInsightsWebTest].
func (aiwt *ApplicationInsightsWebTest) LifecycleManagement() *terra.Lifecycle {
	return aiwt.Lifecycle
}

// Attributes returns the attributes for [ApplicationInsightsWebTest].
func (aiwt *ApplicationInsightsWebTest) Attributes() applicationInsightsWebTestAttributes {
	return applicationInsightsWebTestAttributes{ref: terra.ReferenceResource(aiwt)}
}

// ImportState imports the given attribute values into [ApplicationInsightsWebTest]'s state.
func (aiwt *ApplicationInsightsWebTest) ImportState(av io.Reader) error {
	aiwt.state = &applicationInsightsWebTestState{}
	if err := json.NewDecoder(av).Decode(aiwt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aiwt.Type(), aiwt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApplicationInsightsWebTest] has state.
func (aiwt *ApplicationInsightsWebTest) State() (*applicationInsightsWebTestState, bool) {
	return aiwt.state, aiwt.state != nil
}

// StateMust returns the state for [ApplicationInsightsWebTest]. Panics if the state is nil.
func (aiwt *ApplicationInsightsWebTest) StateMust() *applicationInsightsWebTestState {
	if aiwt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aiwt.Type(), aiwt.LocalName()))
	}
	return aiwt.state
}

// ApplicationInsightsWebTestArgs contains the configurations for azurerm_application_insights_web_test.
type ApplicationInsightsWebTestArgs struct {
	// ApplicationInsightsId: string, required
	ApplicationInsightsId terra.StringValue `hcl:"application_insights_id,attr" validate:"required"`
	// Configuration: string, required
	Configuration terra.StringValue `hcl:"configuration,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Frequency: number, optional
	Frequency terra.NumberValue `hcl:"frequency,attr"`
	// GeoLocations: list of string, required
	GeoLocations terra.ListValue[terra.StringValue] `hcl:"geo_locations,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Kind: string, required
	Kind terra.StringValue `hcl:"kind,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RetryEnabled: bool, optional
	RetryEnabled terra.BoolValue `hcl:"retry_enabled,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeout: number, optional
	Timeout terra.NumberValue `hcl:"timeout,attr"`
	// Timeouts: optional
	Timeouts *applicationinsightswebtest.Timeouts `hcl:"timeouts,block"`
}
type applicationInsightsWebTestAttributes struct {
	ref terra.Reference
}

// ApplicationInsightsId returns a reference to field application_insights_id of azurerm_application_insights_web_test.
func (aiwt applicationInsightsWebTestAttributes) ApplicationInsightsId() terra.StringValue {
	return terra.ReferenceAsString(aiwt.ref.Append("application_insights_id"))
}

// Configuration returns a reference to field configuration of azurerm_application_insights_web_test.
func (aiwt applicationInsightsWebTestAttributes) Configuration() terra.StringValue {
	return terra.ReferenceAsString(aiwt.ref.Append("configuration"))
}

// Description returns a reference to field description of azurerm_application_insights_web_test.
func (aiwt applicationInsightsWebTestAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aiwt.ref.Append("description"))
}

// Enabled returns a reference to field enabled of azurerm_application_insights_web_test.
func (aiwt applicationInsightsWebTestAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(aiwt.ref.Append("enabled"))
}

// Frequency returns a reference to field frequency of azurerm_application_insights_web_test.
func (aiwt applicationInsightsWebTestAttributes) Frequency() terra.NumberValue {
	return terra.ReferenceAsNumber(aiwt.ref.Append("frequency"))
}

// GeoLocations returns a reference to field geo_locations of azurerm_application_insights_web_test.
func (aiwt applicationInsightsWebTestAttributes) GeoLocations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aiwt.ref.Append("geo_locations"))
}

// Id returns a reference to field id of azurerm_application_insights_web_test.
func (aiwt applicationInsightsWebTestAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aiwt.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_application_insights_web_test.
func (aiwt applicationInsightsWebTestAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(aiwt.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_application_insights_web_test.
func (aiwt applicationInsightsWebTestAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aiwt.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_application_insights_web_test.
func (aiwt applicationInsightsWebTestAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aiwt.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_application_insights_web_test.
func (aiwt applicationInsightsWebTestAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aiwt.ref.Append("resource_group_name"))
}

// RetryEnabled returns a reference to field retry_enabled of azurerm_application_insights_web_test.
func (aiwt applicationInsightsWebTestAttributes) RetryEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aiwt.ref.Append("retry_enabled"))
}

// SyntheticMonitorId returns a reference to field synthetic_monitor_id of azurerm_application_insights_web_test.
func (aiwt applicationInsightsWebTestAttributes) SyntheticMonitorId() terra.StringValue {
	return terra.ReferenceAsString(aiwt.ref.Append("synthetic_monitor_id"))
}

// Tags returns a reference to field tags of azurerm_application_insights_web_test.
func (aiwt applicationInsightsWebTestAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aiwt.ref.Append("tags"))
}

// Timeout returns a reference to field timeout of azurerm_application_insights_web_test.
func (aiwt applicationInsightsWebTestAttributes) Timeout() terra.NumberValue {
	return terra.ReferenceAsNumber(aiwt.ref.Append("timeout"))
}

func (aiwt applicationInsightsWebTestAttributes) Timeouts() applicationinsightswebtest.TimeoutsAttributes {
	return terra.ReferenceAsSingle[applicationinsightswebtest.TimeoutsAttributes](aiwt.ref.Append("timeouts"))
}

type applicationInsightsWebTestState struct {
	ApplicationInsightsId string                                    `json:"application_insights_id"`
	Configuration         string                                    `json:"configuration"`
	Description           string                                    `json:"description"`
	Enabled               bool                                      `json:"enabled"`
	Frequency             float64                                   `json:"frequency"`
	GeoLocations          []string                                  `json:"geo_locations"`
	Id                    string                                    `json:"id"`
	Kind                  string                                    `json:"kind"`
	Location              string                                    `json:"location"`
	Name                  string                                    `json:"name"`
	ResourceGroupName     string                                    `json:"resource_group_name"`
	RetryEnabled          bool                                      `json:"retry_enabled"`
	SyntheticMonitorId    string                                    `json:"synthetic_monitor_id"`
	Tags                  map[string]string                         `json:"tags"`
	Timeout               float64                                   `json:"timeout"`
	Timeouts              *applicationinsightswebtest.TimeoutsState `json:"timeouts"`
}
