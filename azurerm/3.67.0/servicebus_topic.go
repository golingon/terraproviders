// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	servicebustopic "github.com/golingon/terraproviders/azurerm/3.67.0/servicebustopic"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicebusTopic creates a new instance of [ServicebusTopic].
func NewServicebusTopic(name string, args ServicebusTopicArgs) *ServicebusTopic {
	return &ServicebusTopic{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicebusTopic)(nil)

// ServicebusTopic represents the Terraform resource azurerm_servicebus_topic.
type ServicebusTopic struct {
	Name      string
	Args      ServicebusTopicArgs
	state     *servicebusTopicState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicebusTopic].
func (st *ServicebusTopic) Type() string {
	return "azurerm_servicebus_topic"
}

// LocalName returns the local name for [ServicebusTopic].
func (st *ServicebusTopic) LocalName() string {
	return st.Name
}

// Configuration returns the configuration (args) for [ServicebusTopic].
func (st *ServicebusTopic) Configuration() interface{} {
	return st.Args
}

// DependOn is used for other resources to depend on [ServicebusTopic].
func (st *ServicebusTopic) DependOn() terra.Reference {
	return terra.ReferenceResource(st)
}

// Dependencies returns the list of resources [ServicebusTopic] depends_on.
func (st *ServicebusTopic) Dependencies() terra.Dependencies {
	return st.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicebusTopic].
func (st *ServicebusTopic) LifecycleManagement() *terra.Lifecycle {
	return st.Lifecycle
}

// Attributes returns the attributes for [ServicebusTopic].
func (st *ServicebusTopic) Attributes() servicebusTopicAttributes {
	return servicebusTopicAttributes{ref: terra.ReferenceResource(st)}
}

// ImportState imports the given attribute values into [ServicebusTopic]'s state.
func (st *ServicebusTopic) ImportState(av io.Reader) error {
	st.state = &servicebusTopicState{}
	if err := json.NewDecoder(av).Decode(st.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", st.Type(), st.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicebusTopic] has state.
func (st *ServicebusTopic) State() (*servicebusTopicState, bool) {
	return st.state, st.state != nil
}

// StateMust returns the state for [ServicebusTopic]. Panics if the state is nil.
func (st *ServicebusTopic) StateMust() *servicebusTopicState {
	if st.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", st.Type(), st.LocalName()))
	}
	return st.state
}

// ServicebusTopicArgs contains the configurations for azurerm_servicebus_topic.
type ServicebusTopicArgs struct {
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
	Timeouts *servicebustopic.Timeouts `hcl:"timeouts,block"`
}
type servicebusTopicAttributes struct {
	ref terra.Reference
}

// AutoDeleteOnIdle returns a reference to field auto_delete_on_idle of azurerm_servicebus_topic.
func (st servicebusTopicAttributes) AutoDeleteOnIdle() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("auto_delete_on_idle"))
}

// DefaultMessageTtl returns a reference to field default_message_ttl of azurerm_servicebus_topic.
func (st servicebusTopicAttributes) DefaultMessageTtl() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("default_message_ttl"))
}

// DuplicateDetectionHistoryTimeWindow returns a reference to field duplicate_detection_history_time_window of azurerm_servicebus_topic.
func (st servicebusTopicAttributes) DuplicateDetectionHistoryTimeWindow() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("duplicate_detection_history_time_window"))
}

// EnableBatchedOperations returns a reference to field enable_batched_operations of azurerm_servicebus_topic.
func (st servicebusTopicAttributes) EnableBatchedOperations() terra.BoolValue {
	return terra.ReferenceAsBool(st.ref.Append("enable_batched_operations"))
}

// EnableExpress returns a reference to field enable_express of azurerm_servicebus_topic.
func (st servicebusTopicAttributes) EnableExpress() terra.BoolValue {
	return terra.ReferenceAsBool(st.ref.Append("enable_express"))
}

// EnablePartitioning returns a reference to field enable_partitioning of azurerm_servicebus_topic.
func (st servicebusTopicAttributes) EnablePartitioning() terra.BoolValue {
	return terra.ReferenceAsBool(st.ref.Append("enable_partitioning"))
}

// Id returns a reference to field id of azurerm_servicebus_topic.
func (st servicebusTopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("id"))
}

// MaxMessageSizeInKilobytes returns a reference to field max_message_size_in_kilobytes of azurerm_servicebus_topic.
func (st servicebusTopicAttributes) MaxMessageSizeInKilobytes() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("max_message_size_in_kilobytes"))
}

// MaxSizeInMegabytes returns a reference to field max_size_in_megabytes of azurerm_servicebus_topic.
func (st servicebusTopicAttributes) MaxSizeInMegabytes() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("max_size_in_megabytes"))
}

// Name returns a reference to field name of azurerm_servicebus_topic.
func (st servicebusTopicAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("name"))
}

// NamespaceId returns a reference to field namespace_id of azurerm_servicebus_topic.
func (st servicebusTopicAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("namespace_id"))
}

// RequiresDuplicateDetection returns a reference to field requires_duplicate_detection of azurerm_servicebus_topic.
func (st servicebusTopicAttributes) RequiresDuplicateDetection() terra.BoolValue {
	return terra.ReferenceAsBool(st.ref.Append("requires_duplicate_detection"))
}

// Status returns a reference to field status of azurerm_servicebus_topic.
func (st servicebusTopicAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("status"))
}

// SupportOrdering returns a reference to field support_ordering of azurerm_servicebus_topic.
func (st servicebusTopicAttributes) SupportOrdering() terra.BoolValue {
	return terra.ReferenceAsBool(st.ref.Append("support_ordering"))
}

func (st servicebusTopicAttributes) Timeouts() servicebustopic.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicebustopic.TimeoutsAttributes](st.ref.Append("timeouts"))
}

type servicebusTopicState struct {
	AutoDeleteOnIdle                    string                         `json:"auto_delete_on_idle"`
	DefaultMessageTtl                   string                         `json:"default_message_ttl"`
	DuplicateDetectionHistoryTimeWindow string                         `json:"duplicate_detection_history_time_window"`
	EnableBatchedOperations             bool                           `json:"enable_batched_operations"`
	EnableExpress                       bool                           `json:"enable_express"`
	EnablePartitioning                  bool                           `json:"enable_partitioning"`
	Id                                  string                         `json:"id"`
	MaxMessageSizeInKilobytes           float64                        `json:"max_message_size_in_kilobytes"`
	MaxSizeInMegabytes                  float64                        `json:"max_size_in_megabytes"`
	Name                                string                         `json:"name"`
	NamespaceId                         string                         `json:"namespace_id"`
	RequiresDuplicateDetection          bool                           `json:"requires_duplicate_detection"`
	Status                              string                         `json:"status"`
	SupportOrdering                     bool                           `json:"support_ordering"`
	Timeouts                            *servicebustopic.TimeoutsState `json:"timeouts"`
}
