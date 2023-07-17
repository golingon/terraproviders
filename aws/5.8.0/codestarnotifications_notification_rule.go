// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	codestarnotificationsnotificationrule "github.com/golingon/terraproviders/aws/5.8.0/codestarnotificationsnotificationrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCodestarnotificationsNotificationRule creates a new instance of [CodestarnotificationsNotificationRule].
func NewCodestarnotificationsNotificationRule(name string, args CodestarnotificationsNotificationRuleArgs) *CodestarnotificationsNotificationRule {
	return &CodestarnotificationsNotificationRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodestarnotificationsNotificationRule)(nil)

// CodestarnotificationsNotificationRule represents the Terraform resource aws_codestarnotifications_notification_rule.
type CodestarnotificationsNotificationRule struct {
	Name      string
	Args      CodestarnotificationsNotificationRuleArgs
	state     *codestarnotificationsNotificationRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CodestarnotificationsNotificationRule].
func (cnr *CodestarnotificationsNotificationRule) Type() string {
	return "aws_codestarnotifications_notification_rule"
}

// LocalName returns the local name for [CodestarnotificationsNotificationRule].
func (cnr *CodestarnotificationsNotificationRule) LocalName() string {
	return cnr.Name
}

// Configuration returns the configuration (args) for [CodestarnotificationsNotificationRule].
func (cnr *CodestarnotificationsNotificationRule) Configuration() interface{} {
	return cnr.Args
}

// DependOn is used for other resources to depend on [CodestarnotificationsNotificationRule].
func (cnr *CodestarnotificationsNotificationRule) DependOn() terra.Reference {
	return terra.ReferenceResource(cnr)
}

// Dependencies returns the list of resources [CodestarnotificationsNotificationRule] depends_on.
func (cnr *CodestarnotificationsNotificationRule) Dependencies() terra.Dependencies {
	return cnr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CodestarnotificationsNotificationRule].
func (cnr *CodestarnotificationsNotificationRule) LifecycleManagement() *terra.Lifecycle {
	return cnr.Lifecycle
}

// Attributes returns the attributes for [CodestarnotificationsNotificationRule].
func (cnr *CodestarnotificationsNotificationRule) Attributes() codestarnotificationsNotificationRuleAttributes {
	return codestarnotificationsNotificationRuleAttributes{ref: terra.ReferenceResource(cnr)}
}

// ImportState imports the given attribute values into [CodestarnotificationsNotificationRule]'s state.
func (cnr *CodestarnotificationsNotificationRule) ImportState(av io.Reader) error {
	cnr.state = &codestarnotificationsNotificationRuleState{}
	if err := json.NewDecoder(av).Decode(cnr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cnr.Type(), cnr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CodestarnotificationsNotificationRule] has state.
func (cnr *CodestarnotificationsNotificationRule) State() (*codestarnotificationsNotificationRuleState, bool) {
	return cnr.state, cnr.state != nil
}

// StateMust returns the state for [CodestarnotificationsNotificationRule]. Panics if the state is nil.
func (cnr *CodestarnotificationsNotificationRule) StateMust() *codestarnotificationsNotificationRuleState {
	if cnr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cnr.Type(), cnr.LocalName()))
	}
	return cnr.state
}

// CodestarnotificationsNotificationRuleArgs contains the configurations for aws_codestarnotifications_notification_rule.
type CodestarnotificationsNotificationRuleArgs struct {
	// DetailType: string, required
	DetailType terra.StringValue `hcl:"detail_type,attr" validate:"required"`
	// EventTypeIds: set of string, required
	EventTypeIds terra.SetValue[terra.StringValue] `hcl:"event_type_ids,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Resource: string, required
	Resource terra.StringValue `hcl:"resource,attr" validate:"required"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Target: min=0,max=10
	Target []codestarnotificationsnotificationrule.Target `hcl:"target,block" validate:"min=0,max=10"`
}
type codestarnotificationsNotificationRuleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_codestarnotifications_notification_rule.
func (cnr codestarnotificationsNotificationRuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cnr.ref.Append("arn"))
}

// DetailType returns a reference to field detail_type of aws_codestarnotifications_notification_rule.
func (cnr codestarnotificationsNotificationRuleAttributes) DetailType() terra.StringValue {
	return terra.ReferenceAsString(cnr.ref.Append("detail_type"))
}

// EventTypeIds returns a reference to field event_type_ids of aws_codestarnotifications_notification_rule.
func (cnr codestarnotificationsNotificationRuleAttributes) EventTypeIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cnr.ref.Append("event_type_ids"))
}

// Id returns a reference to field id of aws_codestarnotifications_notification_rule.
func (cnr codestarnotificationsNotificationRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cnr.ref.Append("id"))
}

// Name returns a reference to field name of aws_codestarnotifications_notification_rule.
func (cnr codestarnotificationsNotificationRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cnr.ref.Append("name"))
}

// Resource returns a reference to field resource of aws_codestarnotifications_notification_rule.
func (cnr codestarnotificationsNotificationRuleAttributes) Resource() terra.StringValue {
	return terra.ReferenceAsString(cnr.ref.Append("resource"))
}

// Status returns a reference to field status of aws_codestarnotifications_notification_rule.
func (cnr codestarnotificationsNotificationRuleAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(cnr.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_codestarnotifications_notification_rule.
func (cnr codestarnotificationsNotificationRuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cnr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_codestarnotifications_notification_rule.
func (cnr codestarnotificationsNotificationRuleAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cnr.ref.Append("tags_all"))
}

func (cnr codestarnotificationsNotificationRuleAttributes) Target() terra.SetValue[codestarnotificationsnotificationrule.TargetAttributes] {
	return terra.ReferenceAsSet[codestarnotificationsnotificationrule.TargetAttributes](cnr.ref.Append("target"))
}

type codestarnotificationsNotificationRuleState struct {
	Arn          string                                              `json:"arn"`
	DetailType   string                                              `json:"detail_type"`
	EventTypeIds []string                                            `json:"event_type_ids"`
	Id           string                                              `json:"id"`
	Name         string                                              `json:"name"`
	Resource     string                                              `json:"resource"`
	Status       string                                              `json:"status"`
	Tags         map[string]string                                   `json:"tags"`
	TagsAll      map[string]string                                   `json:"tags_all"`
	Target       []codestarnotificationsnotificationrule.TargetState `json:"target"`
}
