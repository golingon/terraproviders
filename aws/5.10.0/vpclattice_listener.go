// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpclatticelistener "github.com/golingon/terraproviders/aws/5.10.0/vpclatticelistener"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpclatticeListener creates a new instance of [VpclatticeListener].
func NewVpclatticeListener(name string, args VpclatticeListenerArgs) *VpclatticeListener {
	return &VpclatticeListener{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpclatticeListener)(nil)

// VpclatticeListener represents the Terraform resource aws_vpclattice_listener.
type VpclatticeListener struct {
	Name      string
	Args      VpclatticeListenerArgs
	state     *vpclatticeListenerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpclatticeListener].
func (vl *VpclatticeListener) Type() string {
	return "aws_vpclattice_listener"
}

// LocalName returns the local name for [VpclatticeListener].
func (vl *VpclatticeListener) LocalName() string {
	return vl.Name
}

// Configuration returns the configuration (args) for [VpclatticeListener].
func (vl *VpclatticeListener) Configuration() interface{} {
	return vl.Args
}

// DependOn is used for other resources to depend on [VpclatticeListener].
func (vl *VpclatticeListener) DependOn() terra.Reference {
	return terra.ReferenceResource(vl)
}

// Dependencies returns the list of resources [VpclatticeListener] depends_on.
func (vl *VpclatticeListener) Dependencies() terra.Dependencies {
	return vl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpclatticeListener].
func (vl *VpclatticeListener) LifecycleManagement() *terra.Lifecycle {
	return vl.Lifecycle
}

// Attributes returns the attributes for [VpclatticeListener].
func (vl *VpclatticeListener) Attributes() vpclatticeListenerAttributes {
	return vpclatticeListenerAttributes{ref: terra.ReferenceResource(vl)}
}

// ImportState imports the given attribute values into [VpclatticeListener]'s state.
func (vl *VpclatticeListener) ImportState(av io.Reader) error {
	vl.state = &vpclatticeListenerState{}
	if err := json.NewDecoder(av).Decode(vl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vl.Type(), vl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpclatticeListener] has state.
func (vl *VpclatticeListener) State() (*vpclatticeListenerState, bool) {
	return vl.state, vl.state != nil
}

// StateMust returns the state for [VpclatticeListener]. Panics if the state is nil.
func (vl *VpclatticeListener) StateMust() *vpclatticeListenerState {
	if vl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vl.Type(), vl.LocalName()))
	}
	return vl.state
}

// VpclatticeListenerArgs contains the configurations for aws_vpclattice_listener.
type VpclatticeListenerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
	// ServiceArn: string, optional
	ServiceArn terra.StringValue `hcl:"service_arn,attr"`
	// ServiceIdentifier: string, optional
	ServiceIdentifier terra.StringValue `hcl:"service_identifier,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DefaultAction: required
	DefaultAction *vpclatticelistener.DefaultAction `hcl:"default_action,block" validate:"required"`
	// Timeouts: optional
	Timeouts *vpclatticelistener.Timeouts `hcl:"timeouts,block"`
}
type vpclatticeListenerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpclattice_listener.
func (vl vpclatticeListenerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("arn"))
}

// CreatedAt returns a reference to field created_at of aws_vpclattice_listener.
func (vl vpclatticeListenerAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("created_at"))
}

// Id returns a reference to field id of aws_vpclattice_listener.
func (vl vpclatticeListenerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("id"))
}

// LastUpdatedAt returns a reference to field last_updated_at of aws_vpclattice_listener.
func (vl vpclatticeListenerAttributes) LastUpdatedAt() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("last_updated_at"))
}

// ListenerId returns a reference to field listener_id of aws_vpclattice_listener.
func (vl vpclatticeListenerAttributes) ListenerId() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("listener_id"))
}

// Name returns a reference to field name of aws_vpclattice_listener.
func (vl vpclatticeListenerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("name"))
}

// Port returns a reference to field port of aws_vpclattice_listener.
func (vl vpclatticeListenerAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(vl.ref.Append("port"))
}

// Protocol returns a reference to field protocol of aws_vpclattice_listener.
func (vl vpclatticeListenerAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("protocol"))
}

// ServiceArn returns a reference to field service_arn of aws_vpclattice_listener.
func (vl vpclatticeListenerAttributes) ServiceArn() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("service_arn"))
}

// ServiceIdentifier returns a reference to field service_identifier of aws_vpclattice_listener.
func (vl vpclatticeListenerAttributes) ServiceIdentifier() terra.StringValue {
	return terra.ReferenceAsString(vl.ref.Append("service_identifier"))
}

// Tags returns a reference to field tags of aws_vpclattice_listener.
func (vl vpclatticeListenerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vl.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpclattice_listener.
func (vl vpclatticeListenerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vl.ref.Append("tags_all"))
}

func (vl vpclatticeListenerAttributes) DefaultAction() terra.ListValue[vpclatticelistener.DefaultActionAttributes] {
	return terra.ReferenceAsList[vpclatticelistener.DefaultActionAttributes](vl.ref.Append("default_action"))
}

func (vl vpclatticeListenerAttributes) Timeouts() vpclatticelistener.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpclatticelistener.TimeoutsAttributes](vl.ref.Append("timeouts"))
}

type vpclatticeListenerState struct {
	Arn               string                                  `json:"arn"`
	CreatedAt         string                                  `json:"created_at"`
	Id                string                                  `json:"id"`
	LastUpdatedAt     string                                  `json:"last_updated_at"`
	ListenerId        string                                  `json:"listener_id"`
	Name              string                                  `json:"name"`
	Port              float64                                 `json:"port"`
	Protocol          string                                  `json:"protocol"`
	ServiceArn        string                                  `json:"service_arn"`
	ServiceIdentifier string                                  `json:"service_identifier"`
	Tags              map[string]string                       `json:"tags"`
	TagsAll           map[string]string                       `json:"tags_all"`
	DefaultAction     []vpclatticelistener.DefaultActionState `json:"default_action"`
	Timeouts          *vpclatticelistener.TimeoutsState       `json:"timeouts"`
}
