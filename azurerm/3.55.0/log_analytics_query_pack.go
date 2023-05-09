// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	loganalyticsquerypack "github.com/golingon/terraproviders/azurerm/3.55.0/loganalyticsquerypack"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogAnalyticsQueryPack creates a new instance of [LogAnalyticsQueryPack].
func NewLogAnalyticsQueryPack(name string, args LogAnalyticsQueryPackArgs) *LogAnalyticsQueryPack {
	return &LogAnalyticsQueryPack{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogAnalyticsQueryPack)(nil)

// LogAnalyticsQueryPack represents the Terraform resource azurerm_log_analytics_query_pack.
type LogAnalyticsQueryPack struct {
	Name      string
	Args      LogAnalyticsQueryPackArgs
	state     *logAnalyticsQueryPackState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogAnalyticsQueryPack].
func (laqp *LogAnalyticsQueryPack) Type() string {
	return "azurerm_log_analytics_query_pack"
}

// LocalName returns the local name for [LogAnalyticsQueryPack].
func (laqp *LogAnalyticsQueryPack) LocalName() string {
	return laqp.Name
}

// Configuration returns the configuration (args) for [LogAnalyticsQueryPack].
func (laqp *LogAnalyticsQueryPack) Configuration() interface{} {
	return laqp.Args
}

// DependOn is used for other resources to depend on [LogAnalyticsQueryPack].
func (laqp *LogAnalyticsQueryPack) DependOn() terra.Reference {
	return terra.ReferenceResource(laqp)
}

// Dependencies returns the list of resources [LogAnalyticsQueryPack] depends_on.
func (laqp *LogAnalyticsQueryPack) Dependencies() terra.Dependencies {
	return laqp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogAnalyticsQueryPack].
func (laqp *LogAnalyticsQueryPack) LifecycleManagement() *terra.Lifecycle {
	return laqp.Lifecycle
}

// Attributes returns the attributes for [LogAnalyticsQueryPack].
func (laqp *LogAnalyticsQueryPack) Attributes() logAnalyticsQueryPackAttributes {
	return logAnalyticsQueryPackAttributes{ref: terra.ReferenceResource(laqp)}
}

// ImportState imports the given attribute values into [LogAnalyticsQueryPack]'s state.
func (laqp *LogAnalyticsQueryPack) ImportState(av io.Reader) error {
	laqp.state = &logAnalyticsQueryPackState{}
	if err := json.NewDecoder(av).Decode(laqp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", laqp.Type(), laqp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogAnalyticsQueryPack] has state.
func (laqp *LogAnalyticsQueryPack) State() (*logAnalyticsQueryPackState, bool) {
	return laqp.state, laqp.state != nil
}

// StateMust returns the state for [LogAnalyticsQueryPack]. Panics if the state is nil.
func (laqp *LogAnalyticsQueryPack) StateMust() *logAnalyticsQueryPackState {
	if laqp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", laqp.Type(), laqp.LocalName()))
	}
	return laqp.state
}

// LogAnalyticsQueryPackArgs contains the configurations for azurerm_log_analytics_query_pack.
type LogAnalyticsQueryPackArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *loganalyticsquerypack.Timeouts `hcl:"timeouts,block"`
}
type logAnalyticsQueryPackAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_log_analytics_query_pack.
func (laqp logAnalyticsQueryPackAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(laqp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_log_analytics_query_pack.
func (laqp logAnalyticsQueryPackAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(laqp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_log_analytics_query_pack.
func (laqp logAnalyticsQueryPackAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(laqp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_log_analytics_query_pack.
func (laqp logAnalyticsQueryPackAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(laqp.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_log_analytics_query_pack.
func (laqp logAnalyticsQueryPackAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](laqp.ref.Append("tags"))
}

func (laqp logAnalyticsQueryPackAttributes) Timeouts() loganalyticsquerypack.TimeoutsAttributes {
	return terra.ReferenceAsSingle[loganalyticsquerypack.TimeoutsAttributes](laqp.ref.Append("timeouts"))
}

type logAnalyticsQueryPackState struct {
	Id                string                               `json:"id"`
	Location          string                               `json:"location"`
	Name              string                               `json:"name"`
	ResourceGroupName string                               `json:"resource_group_name"`
	Tags              map[string]string                    `json:"tags"`
	Timeouts          *loganalyticsquerypack.TimeoutsState `json:"timeouts"`
}
