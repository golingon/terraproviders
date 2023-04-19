// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	apprunnervpcingressconnection "github.com/golingon/terraproviders/aws/4.63.0/apprunnervpcingressconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApprunnerVpcIngressConnection creates a new instance of [ApprunnerVpcIngressConnection].
func NewApprunnerVpcIngressConnection(name string, args ApprunnerVpcIngressConnectionArgs) *ApprunnerVpcIngressConnection {
	return &ApprunnerVpcIngressConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApprunnerVpcIngressConnection)(nil)

// ApprunnerVpcIngressConnection represents the Terraform resource aws_apprunner_vpc_ingress_connection.
type ApprunnerVpcIngressConnection struct {
	Name      string
	Args      ApprunnerVpcIngressConnectionArgs
	state     *apprunnerVpcIngressConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApprunnerVpcIngressConnection].
func (avic *ApprunnerVpcIngressConnection) Type() string {
	return "aws_apprunner_vpc_ingress_connection"
}

// LocalName returns the local name for [ApprunnerVpcIngressConnection].
func (avic *ApprunnerVpcIngressConnection) LocalName() string {
	return avic.Name
}

// Configuration returns the configuration (args) for [ApprunnerVpcIngressConnection].
func (avic *ApprunnerVpcIngressConnection) Configuration() interface{} {
	return avic.Args
}

// DependOn is used for other resources to depend on [ApprunnerVpcIngressConnection].
func (avic *ApprunnerVpcIngressConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(avic)
}

// Dependencies returns the list of resources [ApprunnerVpcIngressConnection] depends_on.
func (avic *ApprunnerVpcIngressConnection) Dependencies() terra.Dependencies {
	return avic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApprunnerVpcIngressConnection].
func (avic *ApprunnerVpcIngressConnection) LifecycleManagement() *terra.Lifecycle {
	return avic.Lifecycle
}

// Attributes returns the attributes for [ApprunnerVpcIngressConnection].
func (avic *ApprunnerVpcIngressConnection) Attributes() apprunnerVpcIngressConnectionAttributes {
	return apprunnerVpcIngressConnectionAttributes{ref: terra.ReferenceResource(avic)}
}

// ImportState imports the given attribute values into [ApprunnerVpcIngressConnection]'s state.
func (avic *ApprunnerVpcIngressConnection) ImportState(av io.Reader) error {
	avic.state = &apprunnerVpcIngressConnectionState{}
	if err := json.NewDecoder(av).Decode(avic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avic.Type(), avic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApprunnerVpcIngressConnection] has state.
func (avic *ApprunnerVpcIngressConnection) State() (*apprunnerVpcIngressConnectionState, bool) {
	return avic.state, avic.state != nil
}

// StateMust returns the state for [ApprunnerVpcIngressConnection]. Panics if the state is nil.
func (avic *ApprunnerVpcIngressConnection) StateMust() *apprunnerVpcIngressConnectionState {
	if avic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avic.Type(), avic.LocalName()))
	}
	return avic.state
}

// ApprunnerVpcIngressConnectionArgs contains the configurations for aws_apprunner_vpc_ingress_connection.
type ApprunnerVpcIngressConnectionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ServiceArn: string, required
	ServiceArn terra.StringValue `hcl:"service_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// IngressVpcConfiguration: required
	IngressVpcConfiguration *apprunnervpcingressconnection.IngressVpcConfiguration `hcl:"ingress_vpc_configuration,block" validate:"required"`
}
type apprunnerVpcIngressConnectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_apprunner_vpc_ingress_connection.
func (avic apprunnerVpcIngressConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(avic.ref.Append("arn"))
}

// DomainName returns a reference to field domain_name of aws_apprunner_vpc_ingress_connection.
func (avic apprunnerVpcIngressConnectionAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(avic.ref.Append("domain_name"))
}

// Id returns a reference to field id of aws_apprunner_vpc_ingress_connection.
func (avic apprunnerVpcIngressConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avic.ref.Append("id"))
}

// Name returns a reference to field name of aws_apprunner_vpc_ingress_connection.
func (avic apprunnerVpcIngressConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avic.ref.Append("name"))
}

// ServiceArn returns a reference to field service_arn of aws_apprunner_vpc_ingress_connection.
func (avic apprunnerVpcIngressConnectionAttributes) ServiceArn() terra.StringValue {
	return terra.ReferenceAsString(avic.ref.Append("service_arn"))
}

// Status returns a reference to field status of aws_apprunner_vpc_ingress_connection.
func (avic apprunnerVpcIngressConnectionAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(avic.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_apprunner_vpc_ingress_connection.
func (avic apprunnerVpcIngressConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avic.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_apprunner_vpc_ingress_connection.
func (avic apprunnerVpcIngressConnectionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avic.ref.Append("tags_all"))
}

func (avic apprunnerVpcIngressConnectionAttributes) IngressVpcConfiguration() terra.ListValue[apprunnervpcingressconnection.IngressVpcConfigurationAttributes] {
	return terra.ReferenceAsList[apprunnervpcingressconnection.IngressVpcConfigurationAttributes](avic.ref.Append("ingress_vpc_configuration"))
}

type apprunnerVpcIngressConnectionState struct {
	Arn                     string                                                       `json:"arn"`
	DomainName              string                                                       `json:"domain_name"`
	Id                      string                                                       `json:"id"`
	Name                    string                                                       `json:"name"`
	ServiceArn              string                                                       `json:"service_arn"`
	Status                  string                                                       `json:"status"`
	Tags                    map[string]string                                            `json:"tags"`
	TagsAll                 map[string]string                                            `json:"tags_all"`
	IngressVpcConfiguration []apprunnervpcingressconnection.IngressVpcConfigurationState `json:"ingress_vpc_configuration"`
}
