// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafv2webaclassociation "github.com/golingon/terraproviders/aws/5.8.0/wafv2webaclassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafv2WebAclAssociation creates a new instance of [Wafv2WebAclAssociation].
func NewWafv2WebAclAssociation(name string, args Wafv2WebAclAssociationArgs) *Wafv2WebAclAssociation {
	return &Wafv2WebAclAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Wafv2WebAclAssociation)(nil)

// Wafv2WebAclAssociation represents the Terraform resource aws_wafv2_web_acl_association.
type Wafv2WebAclAssociation struct {
	Name      string
	Args      Wafv2WebAclAssociationArgs
	state     *wafv2WebAclAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Wafv2WebAclAssociation].
func (wwaa *Wafv2WebAclAssociation) Type() string {
	return "aws_wafv2_web_acl_association"
}

// LocalName returns the local name for [Wafv2WebAclAssociation].
func (wwaa *Wafv2WebAclAssociation) LocalName() string {
	return wwaa.Name
}

// Configuration returns the configuration (args) for [Wafv2WebAclAssociation].
func (wwaa *Wafv2WebAclAssociation) Configuration() interface{} {
	return wwaa.Args
}

// DependOn is used for other resources to depend on [Wafv2WebAclAssociation].
func (wwaa *Wafv2WebAclAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(wwaa)
}

// Dependencies returns the list of resources [Wafv2WebAclAssociation] depends_on.
func (wwaa *Wafv2WebAclAssociation) Dependencies() terra.Dependencies {
	return wwaa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Wafv2WebAclAssociation].
func (wwaa *Wafv2WebAclAssociation) LifecycleManagement() *terra.Lifecycle {
	return wwaa.Lifecycle
}

// Attributes returns the attributes for [Wafv2WebAclAssociation].
func (wwaa *Wafv2WebAclAssociation) Attributes() wafv2WebAclAssociationAttributes {
	return wafv2WebAclAssociationAttributes{ref: terra.ReferenceResource(wwaa)}
}

// ImportState imports the given attribute values into [Wafv2WebAclAssociation]'s state.
func (wwaa *Wafv2WebAclAssociation) ImportState(av io.Reader) error {
	wwaa.state = &wafv2WebAclAssociationState{}
	if err := json.NewDecoder(av).Decode(wwaa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wwaa.Type(), wwaa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Wafv2WebAclAssociation] has state.
func (wwaa *Wafv2WebAclAssociation) State() (*wafv2WebAclAssociationState, bool) {
	return wwaa.state, wwaa.state != nil
}

// StateMust returns the state for [Wafv2WebAclAssociation]. Panics if the state is nil.
func (wwaa *Wafv2WebAclAssociation) StateMust() *wafv2WebAclAssociationState {
	if wwaa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wwaa.Type(), wwaa.LocalName()))
	}
	return wwaa.state
}

// Wafv2WebAclAssociationArgs contains the configurations for aws_wafv2_web_acl_association.
type Wafv2WebAclAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
	// WebAclArn: string, required
	WebAclArn terra.StringValue `hcl:"web_acl_arn,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *wafv2webaclassociation.Timeouts `hcl:"timeouts,block"`
}
type wafv2WebAclAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_wafv2_web_acl_association.
func (wwaa wafv2WebAclAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wwaa.ref.Append("id"))
}

// ResourceArn returns a reference to field resource_arn of aws_wafv2_web_acl_association.
func (wwaa wafv2WebAclAssociationAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(wwaa.ref.Append("resource_arn"))
}

// WebAclArn returns a reference to field web_acl_arn of aws_wafv2_web_acl_association.
func (wwaa wafv2WebAclAssociationAttributes) WebAclArn() terra.StringValue {
	return terra.ReferenceAsString(wwaa.ref.Append("web_acl_arn"))
}

func (wwaa wafv2WebAclAssociationAttributes) Timeouts() wafv2webaclassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[wafv2webaclassociation.TimeoutsAttributes](wwaa.ref.Append("timeouts"))
}

type wafv2WebAclAssociationState struct {
	Id          string                                `json:"id"`
	ResourceArn string                                `json:"resource_arn"`
	WebAclArn   string                                `json:"web_acl_arn"`
	Timeouts    *wafv2webaclassociation.TimeoutsState `json:"timeouts"`
}
