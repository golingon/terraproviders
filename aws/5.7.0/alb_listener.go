// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	alblistener "github.com/golingon/terraproviders/aws/5.7.0/alblistener"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAlbListener creates a new instance of [AlbListener].
func NewAlbListener(name string, args AlbListenerArgs) *AlbListener {
	return &AlbListener{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AlbListener)(nil)

// AlbListener represents the Terraform resource aws_alb_listener.
type AlbListener struct {
	Name      string
	Args      AlbListenerArgs
	state     *albListenerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AlbListener].
func (al *AlbListener) Type() string {
	return "aws_alb_listener"
}

// LocalName returns the local name for [AlbListener].
func (al *AlbListener) LocalName() string {
	return al.Name
}

// Configuration returns the configuration (args) for [AlbListener].
func (al *AlbListener) Configuration() interface{} {
	return al.Args
}

// DependOn is used for other resources to depend on [AlbListener].
func (al *AlbListener) DependOn() terra.Reference {
	return terra.ReferenceResource(al)
}

// Dependencies returns the list of resources [AlbListener] depends_on.
func (al *AlbListener) Dependencies() terra.Dependencies {
	return al.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AlbListener].
func (al *AlbListener) LifecycleManagement() *terra.Lifecycle {
	return al.Lifecycle
}

// Attributes returns the attributes for [AlbListener].
func (al *AlbListener) Attributes() albListenerAttributes {
	return albListenerAttributes{ref: terra.ReferenceResource(al)}
}

// ImportState imports the given attribute values into [AlbListener]'s state.
func (al *AlbListener) ImportState(av io.Reader) error {
	al.state = &albListenerState{}
	if err := json.NewDecoder(av).Decode(al.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", al.Type(), al.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AlbListener] has state.
func (al *AlbListener) State() (*albListenerState, bool) {
	return al.state, al.state != nil
}

// StateMust returns the state for [AlbListener]. Panics if the state is nil.
func (al *AlbListener) StateMust() *albListenerState {
	if al.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", al.Type(), al.LocalName()))
	}
	return al.state
}

// AlbListenerArgs contains the configurations for aws_alb_listener.
type AlbListenerArgs struct {
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
	DefaultAction []alblistener.DefaultAction `hcl:"default_action,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *alblistener.Timeouts `hcl:"timeouts,block"`
}
type albListenerAttributes struct {
	ref terra.Reference
}

// AlpnPolicy returns a reference to field alpn_policy of aws_alb_listener.
func (al albListenerAttributes) AlpnPolicy() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("alpn_policy"))
}

// Arn returns a reference to field arn of aws_alb_listener.
func (al albListenerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("arn"))
}

// CertificateArn returns a reference to field certificate_arn of aws_alb_listener.
func (al albListenerAttributes) CertificateArn() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("certificate_arn"))
}

// Id returns a reference to field id of aws_alb_listener.
func (al albListenerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("id"))
}

// LoadBalancerArn returns a reference to field load_balancer_arn of aws_alb_listener.
func (al albListenerAttributes) LoadBalancerArn() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("load_balancer_arn"))
}

// Port returns a reference to field port of aws_alb_listener.
func (al albListenerAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(al.ref.Append("port"))
}

// Protocol returns a reference to field protocol of aws_alb_listener.
func (al albListenerAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("protocol"))
}

// SslPolicy returns a reference to field ssl_policy of aws_alb_listener.
func (al albListenerAttributes) SslPolicy() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("ssl_policy"))
}

// Tags returns a reference to field tags of aws_alb_listener.
func (al albListenerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](al.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_alb_listener.
func (al albListenerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](al.ref.Append("tags_all"))
}

func (al albListenerAttributes) DefaultAction() terra.ListValue[alblistener.DefaultActionAttributes] {
	return terra.ReferenceAsList[alblistener.DefaultActionAttributes](al.ref.Append("default_action"))
}

func (al albListenerAttributes) Timeouts() alblistener.TimeoutsAttributes {
	return terra.ReferenceAsSingle[alblistener.TimeoutsAttributes](al.ref.Append("timeouts"))
}

type albListenerState struct {
	AlpnPolicy      string                           `json:"alpn_policy"`
	Arn             string                           `json:"arn"`
	CertificateArn  string                           `json:"certificate_arn"`
	Id              string                           `json:"id"`
	LoadBalancerArn string                           `json:"load_balancer_arn"`
	Port            float64                          `json:"port"`
	Protocol        string                           `json:"protocol"`
	SslPolicy       string                           `json:"ssl_policy"`
	Tags            map[string]string                `json:"tags"`
	TagsAll         map[string]string                `json:"tags_all"`
	DefaultAction   []alblistener.DefaultActionState `json:"default_action"`
	Timeouts        *alblistener.TimeoutsState       `json:"timeouts"`
}
