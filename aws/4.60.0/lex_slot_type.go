// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lexslottype "github.com/golingon/terraproviders/aws/4.60.0/lexslottype"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLexSlotType creates a new instance of [LexSlotType].
func NewLexSlotType(name string, args LexSlotTypeArgs) *LexSlotType {
	return &LexSlotType{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LexSlotType)(nil)

// LexSlotType represents the Terraform resource aws_lex_slot_type.
type LexSlotType struct {
	Name      string
	Args      LexSlotTypeArgs
	state     *lexSlotTypeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LexSlotType].
func (lst *LexSlotType) Type() string {
	return "aws_lex_slot_type"
}

// LocalName returns the local name for [LexSlotType].
func (lst *LexSlotType) LocalName() string {
	return lst.Name
}

// Configuration returns the configuration (args) for [LexSlotType].
func (lst *LexSlotType) Configuration() interface{} {
	return lst.Args
}

// DependOn is used for other resources to depend on [LexSlotType].
func (lst *LexSlotType) DependOn() terra.Reference {
	return terra.ReferenceResource(lst)
}

// Dependencies returns the list of resources [LexSlotType] depends_on.
func (lst *LexSlotType) Dependencies() terra.Dependencies {
	return lst.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LexSlotType].
func (lst *LexSlotType) LifecycleManagement() *terra.Lifecycle {
	return lst.Lifecycle
}

// Attributes returns the attributes for [LexSlotType].
func (lst *LexSlotType) Attributes() lexSlotTypeAttributes {
	return lexSlotTypeAttributes{ref: terra.ReferenceResource(lst)}
}

// ImportState imports the given attribute values into [LexSlotType]'s state.
func (lst *LexSlotType) ImportState(av io.Reader) error {
	lst.state = &lexSlotTypeState{}
	if err := json.NewDecoder(av).Decode(lst.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lst.Type(), lst.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LexSlotType] has state.
func (lst *LexSlotType) State() (*lexSlotTypeState, bool) {
	return lst.state, lst.state != nil
}

// StateMust returns the state for [LexSlotType]. Panics if the state is nil.
func (lst *LexSlotType) StateMust() *lexSlotTypeState {
	if lst.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lst.Type(), lst.LocalName()))
	}
	return lst.state
}

// LexSlotTypeArgs contains the configurations for aws_lex_slot_type.
type LexSlotTypeArgs struct {
	// CreateVersion: bool, optional
	CreateVersion terra.BoolValue `hcl:"create_version,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ValueSelectionStrategy: string, optional
	ValueSelectionStrategy terra.StringValue `hcl:"value_selection_strategy,attr"`
	// EnumerationValue: min=1,max=10000
	EnumerationValue []lexslottype.EnumerationValue `hcl:"enumeration_value,block" validate:"min=1,max=10000"`
	// Timeouts: optional
	Timeouts *lexslottype.Timeouts `hcl:"timeouts,block"`
}
type lexSlotTypeAttributes struct {
	ref terra.Reference
}

// Checksum returns a reference to field checksum of aws_lex_slot_type.
func (lst lexSlotTypeAttributes) Checksum() terra.StringValue {
	return terra.ReferenceAsString(lst.ref.Append("checksum"))
}

// CreateVersion returns a reference to field create_version of aws_lex_slot_type.
func (lst lexSlotTypeAttributes) CreateVersion() terra.BoolValue {
	return terra.ReferenceAsBool(lst.ref.Append("create_version"))
}

// CreatedDate returns a reference to field created_date of aws_lex_slot_type.
func (lst lexSlotTypeAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(lst.ref.Append("created_date"))
}

// Description returns a reference to field description of aws_lex_slot_type.
func (lst lexSlotTypeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lst.ref.Append("description"))
}

// Id returns a reference to field id of aws_lex_slot_type.
func (lst lexSlotTypeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lst.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_lex_slot_type.
func (lst lexSlotTypeAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(lst.ref.Append("last_updated_date"))
}

// Name returns a reference to field name of aws_lex_slot_type.
func (lst lexSlotTypeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lst.ref.Append("name"))
}

// ValueSelectionStrategy returns a reference to field value_selection_strategy of aws_lex_slot_type.
func (lst lexSlotTypeAttributes) ValueSelectionStrategy() terra.StringValue {
	return terra.ReferenceAsString(lst.ref.Append("value_selection_strategy"))
}

// Version returns a reference to field version of aws_lex_slot_type.
func (lst lexSlotTypeAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(lst.ref.Append("version"))
}

func (lst lexSlotTypeAttributes) EnumerationValue() terra.SetValue[lexslottype.EnumerationValueAttributes] {
	return terra.ReferenceAsSet[lexslottype.EnumerationValueAttributes](lst.ref.Append("enumeration_value"))
}

func (lst lexSlotTypeAttributes) Timeouts() lexslottype.TimeoutsAttributes {
	return terra.ReferenceAsSingle[lexslottype.TimeoutsAttributes](lst.ref.Append("timeouts"))
}

type lexSlotTypeState struct {
	Checksum               string                              `json:"checksum"`
	CreateVersion          bool                                `json:"create_version"`
	CreatedDate            string                              `json:"created_date"`
	Description            string                              `json:"description"`
	Id                     string                              `json:"id"`
	LastUpdatedDate        string                              `json:"last_updated_date"`
	Name                   string                              `json:"name"`
	ValueSelectionStrategy string                              `json:"value_selection_strategy"`
	Version                string                              `json:"version"`
	EnumerationValue       []lexslottype.EnumerationValueState `json:"enumeration_value"`
	Timeouts               *lexslottype.TimeoutsState          `json:"timeouts"`
}
