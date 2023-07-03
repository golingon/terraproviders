// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lblistener "github.com/golingon/terraproviders/aws/5.6.2/lblistener"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLbListener creates a new instance of [LbListener].
func NewLbListener(name string, args LbListenerArgs) *LbListener {
	return &LbListener{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LbListener)(nil)

// LbListener represents the Terraform resource aws_lb_listener.
type LbListener struct {
	Name      string
	Args      LbListenerArgs
	state     *lbListenerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LbListener].
func (ll *LbListener) Type() string {
	return "aws_lb_listener"
}

// LocalName returns the local name for [LbListener].
func (ll *LbListener) LocalName() string {
	return ll.Name
}

// Configuration returns the configuration (args) for [LbListener].
func (ll *LbListener) Configuration() interface{} {
	return ll.Args
}

// DependOn is used for other resources to depend on [LbListener].
func (ll *LbListener) DependOn() terra.Reference {
	return terra.ReferenceResource(ll)
}

// Dependencies returns the list of resources [LbListener] depends_on.
func (ll *LbListener) Dependencies() terra.Dependencies {
	return ll.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LbListener].
func (ll *LbListener) LifecycleManagement() *terra.Lifecycle {
	return ll.Lifecycle
}

// Attributes returns the attributes for [LbListener].
func (ll *LbListener) Attributes() lbListenerAttributes {
	return lbListenerAttributes{ref: terra.ReferenceResource(ll)}
}

// ImportState imports the given attribute values into [LbListener]'s state.
func (ll *LbListener) ImportState(av io.Reader) error {
	ll.state = &lbListenerState{}
	if err := json.NewDecoder(av).Decode(ll.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ll.Type(), ll.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LbListener] has state.
func (ll *LbListener) State() (*lbListenerState, bool) {
	return ll.state, ll.state != nil
}

// StateMust returns the state for [LbListener]. Panics if the state is nil.
func (ll *LbListener) StateMust() *lbListenerState {
	if ll.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ll.Type(), ll.LocalName()))
	}
	return ll.state
}

// LbListenerArgs contains the configurations for aws_lb_listener.
type LbListenerArgs struct {
	// AlpnPolicy: string, optional
	AlpnPolicy terra.StringValue `hcl:"alpn_policy,attr"`
	// CertificateArn: string, optional
	CertificateArn terra.StringValue `hcl:"certificate_arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LoadBalancerArn: string, required
	LoadBalancerArn terra.StringValue `hcl:"load_balancer_arn,attr" validate:"required"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// SslPolicy: string, optional
	SslPolicy terra.StringValue `hcl:"ssl_policy,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DefaultAction: min=1
	DefaultAction []lblistener.DefaultAction `hcl:"default_action,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *lblistener.Timeouts `hcl:"timeouts,block"`
}
type lbListenerAttributes struct {
	ref terra.Reference
}

// AlpnPolicy returns a reference to field alpn_policy of aws_lb_listener.
func (ll lbListenerAttributes) AlpnPolicy() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("alpn_policy"))
}

// Arn returns a reference to field arn of aws_lb_listener.
func (ll lbListenerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("arn"))
}

// CertificateArn returns a reference to field certificate_arn of aws_lb_listener.
func (ll lbListenerAttributes) CertificateArn() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("certificate_arn"))
}

// Id returns a reference to field id of aws_lb_listener.
func (ll lbListenerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("id"))
}

// LoadBalancerArn returns a reference to field load_balancer_arn of aws_lb_listener.
func (ll lbListenerAttributes) LoadBalancerArn() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("load_balancer_arn"))
}

// Port returns a reference to field port of aws_lb_listener.
func (ll lbListenerAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ll.ref.Append("port"))
}

// Protocol returns a reference to field protocol of aws_lb_listener.
func (ll lbListenerAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("protocol"))
}

// SslPolicy returns a reference to field ssl_policy of aws_lb_listener.
func (ll lbListenerAttributes) SslPolicy() terra.StringValue {
	return terra.ReferenceAsString(ll.ref.Append("ssl_policy"))
}

// Tags returns a reference to field tags of aws_lb_listener.
func (ll lbListenerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ll.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_lb_listener.
func (ll lbListenerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ll.ref.Append("tags_all"))
}

func (ll lbListenerAttributes) DefaultAction() terra.ListValue[lblistener.DefaultActionAttributes] {
	return terra.ReferenceAsList[lblistener.DefaultActionAttributes](ll.ref.Append("default_action"))
}

func (ll lbListenerAttributes) Timeouts() lblistener.TimeoutsAttributes {
	return terra.ReferenceAsSingle[lblistener.TimeoutsAttributes](ll.ref.Append("timeouts"))
}

type lbListenerState struct {
	AlpnPolicy      string                          `json:"alpn_policy"`
	Arn             string                          `json:"arn"`
	CertificateArn  string                          `json:"certificate_arn"`
	Id              string                          `json:"id"`
	LoadBalancerArn string                          `json:"load_balancer_arn"`
	Port            float64                         `json:"port"`
	Protocol        string                          `json:"protocol"`
	SslPolicy       string                          `json:"ssl_policy"`
	Tags            map[string]string               `json:"tags"`
	TagsAll         map[string]string               `json:"tags_all"`
	DefaultAction   []lblistener.DefaultActionState `json:"default_action"`
	Timeouts        *lblistener.TimeoutsState       `json:"timeouts"`
}
