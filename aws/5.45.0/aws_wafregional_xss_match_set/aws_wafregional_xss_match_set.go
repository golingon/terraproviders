// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_wafregional_xss_match_set

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

// Resource represents the Terraform resource aws_wafregional_xss_match_set.
type Resource struct {
	Name      string
	Args      Args
	state     *awsWafregionalXssMatchSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (awxms *Resource) Type() string {
	return "aws_wafregional_xss_match_set"
}

// LocalName returns the local name for [Resource].
func (awxms *Resource) LocalName() string {
	return awxms.Name
}

// Configuration returns the configuration (args) for [Resource].
func (awxms *Resource) Configuration() interface{} {
	return awxms.Args
}

// DependOn is used for other resources to depend on [Resource].
func (awxms *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(awxms)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (awxms *Resource) Dependencies() terra.Dependencies {
	return awxms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (awxms *Resource) LifecycleManagement() *terra.Lifecycle {
	return awxms.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (awxms *Resource) Attributes() awsWafregionalXssMatchSetAttributes {
	return awsWafregionalXssMatchSetAttributes{ref: terra.ReferenceResource(awxms)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (awxms *Resource) ImportState(state io.Reader) error {
	awxms.state = &awsWafregionalXssMatchSetState{}
	if err := json.NewDecoder(state).Decode(awxms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", awxms.Type(), awxms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (awxms *Resource) State() (*awsWafregionalXssMatchSetState, bool) {
	return awxms.state, awxms.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (awxms *Resource) StateMust() *awsWafregionalXssMatchSetState {
	if awxms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", awxms.Type(), awxms.LocalName()))
	}
	return awxms.state
}

// Args contains the configurations for aws_wafregional_xss_match_set.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// XssMatchTuple: min=0
	XssMatchTuple []XssMatchTuple `hcl:"xss_match_tuple,block" validate:"min=0"`
}

type awsWafregionalXssMatchSetAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_wafregional_xss_match_set.
func (awxms awsWafregionalXssMatchSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(awxms.ref.Append("id"))
}

// Name returns a reference to field name of aws_wafregional_xss_match_set.
func (awxms awsWafregionalXssMatchSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(awxms.ref.Append("name"))
}

func (awxms awsWafregionalXssMatchSetAttributes) XssMatchTuple() terra.SetValue[XssMatchTupleAttributes] {
	return terra.ReferenceAsSet[XssMatchTupleAttributes](awxms.ref.Append("xss_match_tuple"))
}

type awsWafregionalXssMatchSetState struct {
	Id            string               `json:"id"`
	Name          string               `json:"name"`
	XssMatchTuple []XssMatchTupleState `json:"xss_match_tuple"`
}
