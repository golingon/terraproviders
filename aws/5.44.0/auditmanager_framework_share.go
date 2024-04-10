// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewAuditmanagerFrameworkShare creates a new instance of [AuditmanagerFrameworkShare].
func NewAuditmanagerFrameworkShare(name string, args AuditmanagerFrameworkShareArgs) *AuditmanagerFrameworkShare {
	return &AuditmanagerFrameworkShare{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AuditmanagerFrameworkShare)(nil)

// AuditmanagerFrameworkShare represents the Terraform resource aws_auditmanager_framework_share.
type AuditmanagerFrameworkShare struct {
	Name      string
	Args      AuditmanagerFrameworkShareArgs
	state     *auditmanagerFrameworkShareState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AuditmanagerFrameworkShare].
func (afs *AuditmanagerFrameworkShare) Type() string {
	return "aws_auditmanager_framework_share"
}

// LocalName returns the local name for [AuditmanagerFrameworkShare].
func (afs *AuditmanagerFrameworkShare) LocalName() string {
	return afs.Name
}

// Configuration returns the configuration (args) for [AuditmanagerFrameworkShare].
func (afs *AuditmanagerFrameworkShare) Configuration() interface{} {
	return afs.Args
}

// DependOn is used for other resources to depend on [AuditmanagerFrameworkShare].
func (afs *AuditmanagerFrameworkShare) DependOn() terra.Reference {
	return terra.ReferenceResource(afs)
}

// Dependencies returns the list of resources [AuditmanagerFrameworkShare] depends_on.
func (afs *AuditmanagerFrameworkShare) Dependencies() terra.Dependencies {
	return afs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AuditmanagerFrameworkShare].
func (afs *AuditmanagerFrameworkShare) LifecycleManagement() *terra.Lifecycle {
	return afs.Lifecycle
}

// Attributes returns the attributes for [AuditmanagerFrameworkShare].
func (afs *AuditmanagerFrameworkShare) Attributes() auditmanagerFrameworkShareAttributes {
	return auditmanagerFrameworkShareAttributes{ref: terra.ReferenceResource(afs)}
}

// ImportState imports the given attribute values into [AuditmanagerFrameworkShare]'s state.
func (afs *AuditmanagerFrameworkShare) ImportState(av io.Reader) error {
	afs.state = &auditmanagerFrameworkShareState{}
	if err := json.NewDecoder(av).Decode(afs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", afs.Type(), afs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AuditmanagerFrameworkShare] has state.
func (afs *AuditmanagerFrameworkShare) State() (*auditmanagerFrameworkShareState, bool) {
	return afs.state, afs.state != nil
}

// StateMust returns the state for [AuditmanagerFrameworkShare]. Panics if the state is nil.
func (afs *AuditmanagerFrameworkShare) StateMust() *auditmanagerFrameworkShareState {
	if afs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", afs.Type(), afs.LocalName()))
	}
	return afs.state
}

// AuditmanagerFrameworkShareArgs contains the configurations for aws_auditmanager_framework_share.
type AuditmanagerFrameworkShareArgs struct {
	// Comment: string, optional
	Comment terra.StringValue `hcl:"comment,attr"`
	// DestinationAccount: string, required
	DestinationAccount terra.StringValue `hcl:"destination_account,attr" validate:"required"`
	// DestinationRegion: string, required
	DestinationRegion terra.StringValue `hcl:"destination_region,attr" validate:"required"`
	// FrameworkId: string, required
	FrameworkId terra.StringValue `hcl:"framework_id,attr" validate:"required"`
}
type auditmanagerFrameworkShareAttributes struct {
	ref terra.Reference
}

// Comment returns a reference to field comment of aws_auditmanager_framework_share.
func (afs auditmanagerFrameworkShareAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(afs.ref.Append("comment"))
}

// DestinationAccount returns a reference to field destination_account of aws_auditmanager_framework_share.
func (afs auditmanagerFrameworkShareAttributes) DestinationAccount() terra.StringValue {
	return terra.ReferenceAsString(afs.ref.Append("destination_account"))
}

// DestinationRegion returns a reference to field destination_region of aws_auditmanager_framework_share.
func (afs auditmanagerFrameworkShareAttributes) DestinationRegion() terra.StringValue {
	return terra.ReferenceAsString(afs.ref.Append("destination_region"))
}

// FrameworkId returns a reference to field framework_id of aws_auditmanager_framework_share.
func (afs auditmanagerFrameworkShareAttributes) FrameworkId() terra.StringValue {
	return terra.ReferenceAsString(afs.ref.Append("framework_id"))
}

// Id returns a reference to field id of aws_auditmanager_framework_share.
func (afs auditmanagerFrameworkShareAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(afs.ref.Append("id"))
}

// Status returns a reference to field status of aws_auditmanager_framework_share.
func (afs auditmanagerFrameworkShareAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(afs.ref.Append("status"))
}

type auditmanagerFrameworkShareState struct {
	Comment            string `json:"comment"`
	DestinationAccount string `json:"destination_account"`
	DestinationRegion  string `json:"destination_region"`
	FrameworkId        string `json:"framework_id"`
	Id                 string `json:"id"`
	Status             string `json:"status"`
}
