// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	eventhubnamespaceschemagroup "github.com/golingon/terraproviders/azurerm/3.49.0/eventhubnamespaceschemagroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewEventhubNamespaceSchemaGroup(name string, args EventhubNamespaceSchemaGroupArgs) *EventhubNamespaceSchemaGroup {
	return &EventhubNamespaceSchemaGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EventhubNamespaceSchemaGroup)(nil)

type EventhubNamespaceSchemaGroup struct {
	Name  string
	Args  EventhubNamespaceSchemaGroupArgs
	state *eventhubNamespaceSchemaGroupState
}

func (ensg *EventhubNamespaceSchemaGroup) Type() string {
	return "azurerm_eventhub_namespace_schema_group"
}

func (ensg *EventhubNamespaceSchemaGroup) LocalName() string {
	return ensg.Name
}

func (ensg *EventhubNamespaceSchemaGroup) Configuration() interface{} {
	return ensg.Args
}

func (ensg *EventhubNamespaceSchemaGroup) Attributes() eventhubNamespaceSchemaGroupAttributes {
	return eventhubNamespaceSchemaGroupAttributes{ref: terra.ReferenceResource(ensg)}
}

func (ensg *EventhubNamespaceSchemaGroup) ImportState(av io.Reader) error {
	ensg.state = &eventhubNamespaceSchemaGroupState{}
	if err := json.NewDecoder(av).Decode(ensg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ensg.Type(), ensg.LocalName(), err)
	}
	return nil
}

func (ensg *EventhubNamespaceSchemaGroup) State() (*eventhubNamespaceSchemaGroupState, bool) {
	return ensg.state, ensg.state != nil
}

func (ensg *EventhubNamespaceSchemaGroup) StateMust() *eventhubNamespaceSchemaGroupState {
	if ensg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ensg.Type(), ensg.LocalName()))
	}
	return ensg.state
}

func (ensg *EventhubNamespaceSchemaGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(ensg)
}

type EventhubNamespaceSchemaGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceId: string, required
	NamespaceId terra.StringValue `hcl:"namespace_id,attr" validate:"required"`
	// SchemaCompatibility: string, required
	SchemaCompatibility terra.StringValue `hcl:"schema_compatibility,attr" validate:"required"`
	// SchemaType: string, required
	SchemaType terra.StringValue `hcl:"schema_type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *eventhubnamespaceschemagroup.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that EventhubNamespaceSchemaGroup depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type eventhubNamespaceSchemaGroupAttributes struct {
	ref terra.Reference
}

func (ensg eventhubNamespaceSchemaGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ensg.ref.Append("id"))
}

func (ensg eventhubNamespaceSchemaGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ensg.ref.Append("name"))
}

func (ensg eventhubNamespaceSchemaGroupAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceString(ensg.ref.Append("namespace_id"))
}

func (ensg eventhubNamespaceSchemaGroupAttributes) SchemaCompatibility() terra.StringValue {
	return terra.ReferenceString(ensg.ref.Append("schema_compatibility"))
}

func (ensg eventhubNamespaceSchemaGroupAttributes) SchemaType() terra.StringValue {
	return terra.ReferenceString(ensg.ref.Append("schema_type"))
}

func (ensg eventhubNamespaceSchemaGroupAttributes) Timeouts() eventhubnamespaceschemagroup.TimeoutsAttributes {
	return terra.ReferenceSingle[eventhubnamespaceschemagroup.TimeoutsAttributes](ensg.ref.Append("timeouts"))
}

type eventhubNamespaceSchemaGroupState struct {
	Id                  string                                      `json:"id"`
	Name                string                                      `json:"name"`
	NamespaceId         string                                      `json:"namespace_id"`
	SchemaCompatibility string                                      `json:"schema_compatibility"`
	SchemaType          string                                      `json:"schema_type"`
	Timeouts            *eventhubnamespaceschemagroup.TimeoutsState `json:"timeouts"`
}
