// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	auditmanagerframework "github.com/golingon/terraproviders/aws/5.0.1/auditmanagerframework"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAuditmanagerFramework creates a new instance of [AuditmanagerFramework].
func NewAuditmanagerFramework(name string, args AuditmanagerFrameworkArgs) *AuditmanagerFramework {
	return &AuditmanagerFramework{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AuditmanagerFramework)(nil)

// AuditmanagerFramework represents the Terraform resource aws_auditmanager_framework.
type AuditmanagerFramework struct {
	Name      string
	Args      AuditmanagerFrameworkArgs
	state     *auditmanagerFrameworkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AuditmanagerFramework].
func (af *AuditmanagerFramework) Type() string {
	return "aws_auditmanager_framework"
}

// LocalName returns the local name for [AuditmanagerFramework].
func (af *AuditmanagerFramework) LocalName() string {
	return af.Name
}

// Configuration returns the configuration (args) for [AuditmanagerFramework].
func (af *AuditmanagerFramework) Configuration() interface{} {
	return af.Args
}

// DependOn is used for other resources to depend on [AuditmanagerFramework].
func (af *AuditmanagerFramework) DependOn() terra.Reference {
	return terra.ReferenceResource(af)
}

// Dependencies returns the list of resources [AuditmanagerFramework] depends_on.
func (af *AuditmanagerFramework) Dependencies() terra.Dependencies {
	return af.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AuditmanagerFramework].
func (af *AuditmanagerFramework) LifecycleManagement() *terra.Lifecycle {
	return af.Lifecycle
}

// Attributes returns the attributes for [AuditmanagerFramework].
func (af *AuditmanagerFramework) Attributes() auditmanagerFrameworkAttributes {
	return auditmanagerFrameworkAttributes{ref: terra.ReferenceResource(af)}
}

// ImportState imports the given attribute values into [AuditmanagerFramework]'s state.
func (af *AuditmanagerFramework) ImportState(av io.Reader) error {
	af.state = &auditmanagerFrameworkState{}
	if err := json.NewDecoder(av).Decode(af.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", af.Type(), af.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AuditmanagerFramework] has state.
func (af *AuditmanagerFramework) State() (*auditmanagerFrameworkState, bool) {
	return af.state, af.state != nil
}

// StateMust returns the state for [AuditmanagerFramework]. Panics if the state is nil.
func (af *AuditmanagerFramework) StateMust() *auditmanagerFrameworkState {
	if af.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", af.Type(), af.LocalName()))
	}
	return af.state
}

// AuditmanagerFrameworkArgs contains the configurations for aws_auditmanager_framework.
type AuditmanagerFrameworkArgs struct {
	// ComplianceType: string, optional
	ComplianceType terra.StringValue `hcl:"compliance_type,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ControlSets: min=0
	ControlSets []auditmanagerframework.ControlSets `hcl:"control_sets,block" validate:"min=0"`
}
type auditmanagerFrameworkAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_auditmanager_framework.
func (af auditmanagerFrameworkAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("arn"))
}

// ComplianceType returns a reference to field compliance_type of aws_auditmanager_framework.
func (af auditmanagerFrameworkAttributes) ComplianceType() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("compliance_type"))
}

// Description returns a reference to field description of aws_auditmanager_framework.
func (af auditmanagerFrameworkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("description"))
}

// FrameworkType returns a reference to field framework_type of aws_auditmanager_framework.
func (af auditmanagerFrameworkAttributes) FrameworkType() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("framework_type"))
}

// Id returns a reference to field id of aws_auditmanager_framework.
func (af auditmanagerFrameworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("id"))
}

// Name returns a reference to field name of aws_auditmanager_framework.
func (af auditmanagerFrameworkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(af.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_auditmanager_framework.
func (af auditmanagerFrameworkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](af.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_auditmanager_framework.
func (af auditmanagerFrameworkAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](af.ref.Append("tags_all"))
}

func (af auditmanagerFrameworkAttributes) ControlSets() terra.SetValue[auditmanagerframework.ControlSetsAttributes] {
	return terra.ReferenceAsSet[auditmanagerframework.ControlSetsAttributes](af.ref.Append("control_sets"))
}

type auditmanagerFrameworkState struct {
	Arn            string                                   `json:"arn"`
	ComplianceType string                                   `json:"compliance_type"`
	Description    string                                   `json:"description"`
	FrameworkType  string                                   `json:"framework_type"`
	Id             string                                   `json:"id"`
	Name           string                                   `json:"name"`
	Tags           map[string]string                        `json:"tags"`
	TagsAll        map[string]string                        `json:"tags_all"`
	ControlSets    []auditmanagerframework.ControlSetsState `json:"control_sets"`
}
