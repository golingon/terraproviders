// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_servicebus_topic

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

// Resource represents the Terraform resource azurerm_servicebus_topic.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermServicebusTopicState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ast *Resource) Type() string {
	return "azurerm_servicebus_topic"
}

// LocalName returns the local name for [Resource].
func (ast *Resource) LocalName() string {
	return ast.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ast *Resource) Configuration() interface{} {
	return ast.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ast *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ast)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ast *Resource) Dependencies() terra.Dependencies {
	return ast.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ast *Resource) LifecycleManagement() *terra.Lifecycle {
	return ast.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ast *Resource) Attributes() azurermServicebusTopicAttributes {
	return azurermServicebusTopicAttributes{ref: terra.ReferenceResource(ast)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ast *Resource) ImportState(state io.Reader) error {
	ast.state = &azurermServicebusTopicState{}
	if err := json.NewDecoder(state).Decode(ast.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ast.Type(), ast.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ast *Resource) State() (*azurermServicebusTopicState, bool) {
	return ast.state, ast.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ast *Resource) StateMust() *azurermServicebusTopicState {
	if ast.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ast.Type(), ast.LocalName()))
	}
	return ast.state
}

// Args contains the configurations for azurerm_servicebus_topic.
type Args struct {
	// AutoDeleteOnIdle: string, optional
	AutoDeleteOnIdle terra.StringValue `hcl:"auto_delete_on_idle,attr"`
	// DefaultMessageTtl: string, optional
	DefaultMessageTtl terra.StringValue `hcl:"default_message_ttl,attr"`
	// DuplicateDetectionHistoryTimeWindow: string, optional
	DuplicateDetectionHistoryTimeWindow terra.StringValue `hcl:"duplicate_detection_history_time_window,attr"`
	// EnableBatchedOperations: bool, optional
	EnableBatchedOperations terra.BoolValue `hcl:"enable_batched_operations,attr"`
	// EnableExpress: bool, optional
	EnableExpress terra.BoolValue `hcl:"enable_express,attr"`
	// EnablePartitioning: bool, optional
	EnablePartitioning terra.BoolValue `hcl:"enable_partitioning,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxMessageSizeInKilobytes: number, optional
	MaxMessageSizeInKilobytes terra.NumberValue `hcl:"max_message_size_in_kilobytes,attr"`
	// MaxSizeInMegabytes: number, optional
	MaxSizeInMegabytes terra.NumberValue `hcl:"max_size_in_megabytes,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceId: string, required
	NamespaceId terra.StringValue `hcl:"namespace_id,attr" validate:"required"`
	// RequiresDuplicateDetection: bool, optional
	RequiresDuplicateDetection terra.BoolValue `hcl:"requires_duplicate_detection,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// SupportOrdering: bool, optional
	SupportOrdering terra.BoolValue `hcl:"support_ordering,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermServicebusTopicAttributes struct {
	ref terra.Reference
}

// AutoDeleteOnIdle returns a reference to field auto_delete_on_idle of azurerm_servicebus_topic.
func (ast azurermServicebusTopicAttributes) AutoDeleteOnIdle() terra.StringValue {
	return terra.ReferenceAsString(ast.ref.Append("auto_delete_on_idle"))
}

// DefaultMessageTtl returns a reference to field default_message_ttl of azurerm_servicebus_topic.
func (ast azurermServicebusTopicAttributes) DefaultMessageTtl() terra.StringValue {
	return terra.ReferenceAsString(ast.ref.Append("default_message_ttl"))
}

// DuplicateDetectionHistoryTimeWindow returns a reference to field duplicate_detection_history_time_window of azurerm_servicebus_topic.
func (ast azurermServicebusTopicAttributes) DuplicateDetectionHistoryTimeWindow() terra.StringValue {
	return terra.ReferenceAsString(ast.ref.Append("duplicate_detection_history_time_window"))
}

// EnableBatchedOperations returns a reference to field enable_batched_operations of azurerm_servicebus_topic.
func (ast azurermServicebusTopicAttributes) EnableBatchedOperations() terra.BoolValue {
	return terra.ReferenceAsBool(ast.ref.Append("enable_batched_operations"))
}

// EnableExpress returns a reference to field enable_express of azurerm_servicebus_topic.
func (ast azurermServicebusTopicAttributes) EnableExpress() terra.BoolValue {
	return terra.ReferenceAsBool(ast.ref.Append("enable_express"))
}

// EnablePartitioning returns a reference to field enable_partitioning of azurerm_servicebus_topic.
func (ast azurermServicebusTopicAttributes) EnablePartitioning() terra.BoolValue {
	return terra.ReferenceAsBool(ast.ref.Append("enable_partitioning"))
}

// Id returns a reference to field id of azurerm_servicebus_topic.
func (ast azurermServicebusTopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ast.ref.Append("id"))
}

// MaxMessageSizeInKilobytes returns a reference to field max_message_size_in_kilobytes of azurerm_servicebus_topic.
func (ast azurermServicebusTopicAttributes) MaxMessageSizeInKilobytes() terra.NumberValue {
	return terra.ReferenceAsNumber(ast.ref.Append("max_message_size_in_kilobytes"))
}

// MaxSizeInMegabytes returns a reference to field max_size_in_megabytes of azurerm_servicebus_topic.
func (ast azurermServicebusTopicAttributes) MaxSizeInMegabytes() terra.NumberValue {
	return terra.ReferenceAsNumber(ast.ref.Append("max_size_in_megabytes"))
}

// Name returns a reference to field name of azurerm_servicebus_topic.
func (ast azurermServicebusTopicAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ast.ref.Append("name"))
}

// NamespaceId returns a reference to field namespace_id of azurerm_servicebus_topic.
func (ast azurermServicebusTopicAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(ast.ref.Append("namespace_id"))
}

// RequiresDuplicateDetection returns a reference to field requires_duplicate_detection of azurerm_servicebus_topic.
func (ast azurermServicebusTopicAttributes) RequiresDuplicateDetection() terra.BoolValue {
	return terra.ReferenceAsBool(ast.ref.Append("requires_duplicate_detection"))
}

// Status returns a reference to field status of azurerm_servicebus_topic.
func (ast azurermServicebusTopicAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ast.ref.Append("status"))
}

// SupportOrdering returns a reference to field support_ordering of azurerm_servicebus_topic.
func (ast azurermServicebusTopicAttributes) SupportOrdering() terra.BoolValue {
	return terra.ReferenceAsBool(ast.ref.Append("support_ordering"))
}

func (ast azurermServicebusTopicAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ast.ref.Append("timeouts"))
}

type azurermServicebusTopicState struct {
	AutoDeleteOnIdle                    string         `json:"auto_delete_on_idle"`
	DefaultMessageTtl                   string         `json:"default_message_ttl"`
	DuplicateDetectionHistoryTimeWindow string         `json:"duplicate_detection_history_time_window"`
	EnableBatchedOperations             bool           `json:"enable_batched_operations"`
	EnableExpress                       bool           `json:"enable_express"`
	EnablePartitioning                  bool           `json:"enable_partitioning"`
	Id                                  string         `json:"id"`
	MaxMessageSizeInKilobytes           float64        `json:"max_message_size_in_kilobytes"`
	MaxSizeInMegabytes                  float64        `json:"max_size_in_megabytes"`
	Name                                string         `json:"name"`
	NamespaceId                         string         `json:"namespace_id"`
	RequiresDuplicateDetection          bool           `json:"requires_duplicate_detection"`
	Status                              string         `json:"status"`
	SupportOrdering                     bool           `json:"support_ordering"`
	Timeouts                            *TimeoutsState `json:"timeouts"`
}
