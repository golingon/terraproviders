// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	appengineservicesplittraffic "github.com/golingon/terraproviders/google/4.73.1/appengineservicesplittraffic"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppEngineServiceSplitTraffic creates a new instance of [AppEngineServiceSplitTraffic].
func NewAppEngineServiceSplitTraffic(name string, args AppEngineServiceSplitTrafficArgs) *AppEngineServiceSplitTraffic {
	return &AppEngineServiceSplitTraffic{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppEngineServiceSplitTraffic)(nil)

// AppEngineServiceSplitTraffic represents the Terraform resource google_app_engine_service_split_traffic.
type AppEngineServiceSplitTraffic struct {
	Name      string
	Args      AppEngineServiceSplitTrafficArgs
	state     *appEngineServiceSplitTrafficState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppEngineServiceSplitTraffic].
func (aesst *AppEngineServiceSplitTraffic) Type() string {
	return "google_app_engine_service_split_traffic"
}

// LocalName returns the local name for [AppEngineServiceSplitTraffic].
func (aesst *AppEngineServiceSplitTraffic) LocalName() string {
	return aesst.Name
}

// Configuration returns the configuration (args) for [AppEngineServiceSplitTraffic].
func (aesst *AppEngineServiceSplitTraffic) Configuration() interface{} {
	return aesst.Args
}

// DependOn is used for other resources to depend on [AppEngineServiceSplitTraffic].
func (aesst *AppEngineServiceSplitTraffic) DependOn() terra.Reference {
	return terra.ReferenceResource(aesst)
}

// Dependencies returns the list of resources [AppEngineServiceSplitTraffic] depends_on.
func (aesst *AppEngineServiceSplitTraffic) Dependencies() terra.Dependencies {
	return aesst.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppEngineServiceSplitTraffic].
func (aesst *AppEngineServiceSplitTraffic) LifecycleManagement() *terra.Lifecycle {
	return aesst.Lifecycle
}

// Attributes returns the attributes for [AppEngineServiceSplitTraffic].
func (aesst *AppEngineServiceSplitTraffic) Attributes() appEngineServiceSplitTrafficAttributes {
	return appEngineServiceSplitTrafficAttributes{ref: terra.ReferenceResource(aesst)}
}

// ImportState imports the given attribute values into [AppEngineServiceSplitTraffic]'s state.
func (aesst *AppEngineServiceSplitTraffic) ImportState(av io.Reader) error {
	aesst.state = &appEngineServiceSplitTrafficState{}
	if err := json.NewDecoder(av).Decode(aesst.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aesst.Type(), aesst.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppEngineServiceSplitTraffic] has state.
func (aesst *AppEngineServiceSplitTraffic) State() (*appEngineServiceSplitTrafficState, bool) {
	return aesst.state, aesst.state != nil
}

// StateMust returns the state for [AppEngineServiceSplitTraffic]. Panics if the state is nil.
func (aesst *AppEngineServiceSplitTraffic) StateMust() *appEngineServiceSplitTrafficState {
	if aesst.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aesst.Type(), aesst.LocalName()))
	}
	return aesst.state
}

// AppEngineServiceSplitTrafficArgs contains the configurations for google_app_engine_service_split_traffic.
type AppEngineServiceSplitTrafficArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MigrateTraffic: bool, optional
	MigrateTraffic terra.BoolValue `hcl:"migrate_traffic,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// Split: required
	Split *appengineservicesplittraffic.Split `hcl:"split,block" validate:"required"`
	// Timeouts: optional
	Timeouts *appengineservicesplittraffic.Timeouts `hcl:"timeouts,block"`
}
type appEngineServiceSplitTrafficAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_app_engine_service_split_traffic.
func (aesst appEngineServiceSplitTrafficAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aesst.ref.Append("id"))
}

// MigrateTraffic returns a reference to field migrate_traffic of google_app_engine_service_split_traffic.
func (aesst appEngineServiceSplitTrafficAttributes) MigrateTraffic() terra.BoolValue {
	return terra.ReferenceAsBool(aesst.ref.Append("migrate_traffic"))
}

// Project returns a reference to field project of google_app_engine_service_split_traffic.
func (aesst appEngineServiceSplitTrafficAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(aesst.ref.Append("project"))
}

// Service returns a reference to field service of google_app_engine_service_split_traffic.
func (aesst appEngineServiceSplitTrafficAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(aesst.ref.Append("service"))
}

func (aesst appEngineServiceSplitTrafficAttributes) Split() terra.ListValue[appengineservicesplittraffic.SplitAttributes] {
	return terra.ReferenceAsList[appengineservicesplittraffic.SplitAttributes](aesst.ref.Append("split"))
}

func (aesst appEngineServiceSplitTrafficAttributes) Timeouts() appengineservicesplittraffic.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appengineservicesplittraffic.TimeoutsAttributes](aesst.ref.Append("timeouts"))
}

type appEngineServiceSplitTrafficState struct {
	Id             string                                      `json:"id"`
	MigrateTraffic bool                                        `json:"migrate_traffic"`
	Project        string                                      `json:"project"`
	Service        string                                      `json:"service"`
	Split          []appengineservicesplittraffic.SplitState   `json:"split"`
	Timeouts       *appengineservicesplittraffic.TimeoutsState `json:"timeouts"`
}
