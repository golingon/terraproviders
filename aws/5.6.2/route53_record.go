// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	route53record "github.com/golingon/terraproviders/aws/5.6.2/route53record"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRoute53Record creates a new instance of [Route53Record].
func NewRoute53Record(name string, args Route53RecordArgs) *Route53Record {
	return &Route53Record{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53Record)(nil)

// Route53Record represents the Terraform resource aws_route53_record.
type Route53Record struct {
	Name      string
	Args      Route53RecordArgs
	state     *route53RecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53Record].
func (rr *Route53Record) Type() string {
	return "aws_route53_record"
}

// LocalName returns the local name for [Route53Record].
func (rr *Route53Record) LocalName() string {
	return rr.Name
}

// Configuration returns the configuration (args) for [Route53Record].
func (rr *Route53Record) Configuration() interface{} {
	return rr.Args
}

// DependOn is used for other resources to depend on [Route53Record].
func (rr *Route53Record) DependOn() terra.Reference {
	return terra.ReferenceResource(rr)
}

// Dependencies returns the list of resources [Route53Record] depends_on.
func (rr *Route53Record) Dependencies() terra.Dependencies {
	return rr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53Record].
func (rr *Route53Record) LifecycleManagement() *terra.Lifecycle {
	return rr.Lifecycle
}

// Attributes returns the attributes for [Route53Record].
func (rr *Route53Record) Attributes() route53RecordAttributes {
	return route53RecordAttributes{ref: terra.ReferenceResource(rr)}
}

// ImportState imports the given attribute values into [Route53Record]'s state.
func (rr *Route53Record) ImportState(av io.Reader) error {
	rr.state = &route53RecordState{}
	if err := json.NewDecoder(av).Decode(rr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rr.Type(), rr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53Record] has state.
func (rr *Route53Record) State() (*route53RecordState, bool) {
	return rr.state, rr.state != nil
}

// StateMust returns the state for [Route53Record]. Panics if the state is nil.
func (rr *Route53Record) StateMust() *route53RecordState {
	if rr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rr.Type(), rr.LocalName()))
	}
	return rr.state
}

// Route53RecordArgs contains the configurations for aws_route53_record.
type Route53RecordArgs struct {
	// AllowOverwrite: bool, optional
	AllowOverwrite terra.BoolValue `hcl:"allow_overwrite,attr"`
	// HealthCheckId: string, optional
	HealthCheckId terra.StringValue `hcl:"health_check_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MultivalueAnswerRoutingPolicy: bool, optional
	MultivalueAnswerRoutingPolicy terra.BoolValue `hcl:"multivalue_answer_routing_policy,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Records: set of string, optional
	Records terra.SetValue[terra.StringValue] `hcl:"records,attr"`
	// SetIdentifier: string, optional
	SetIdentifier terra.StringValue `hcl:"set_identifier,attr"`
	// Ttl: number, optional
	Ttl terra.NumberValue `hcl:"ttl,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// ZoneId: string, required
	ZoneId terra.StringValue `hcl:"zone_id,attr" validate:"required"`
	// Alias: optional
	Alias *route53record.Alias `hcl:"alias,block"`
	// CidrRoutingPolicy: optional
	CidrRoutingPolicy *route53record.CidrRoutingPolicy `hcl:"cidr_routing_policy,block"`
	// FailoverRoutingPolicy: optional
	FailoverRoutingPolicy *route53record.FailoverRoutingPolicy `hcl:"failover_routing_policy,block"`
	// GeolocationRoutingPolicy: optional
	GeolocationRoutingPolicy *route53record.GeolocationRoutingPolicy `hcl:"geolocation_routing_policy,block"`
	// LatencyRoutingPolicy: optional
	LatencyRoutingPolicy *route53record.LatencyRoutingPolicy `hcl:"latency_routing_policy,block"`
	// WeightedRoutingPolicy: optional
	WeightedRoutingPolicy *route53record.WeightedRoutingPolicy `hcl:"weighted_routing_policy,block"`
}
type route53RecordAttributes struct {
	ref terra.Reference
}

// AllowOverwrite returns a reference to field allow_overwrite of aws_route53_record.
func (rr route53RecordAttributes) AllowOverwrite() terra.BoolValue {
	return terra.ReferenceAsBool(rr.ref.Append("allow_overwrite"))
}

// Fqdn returns a reference to field fqdn of aws_route53_record.
func (rr route53RecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("fqdn"))
}

// HealthCheckId returns a reference to field health_check_id of aws_route53_record.
func (rr route53RecordAttributes) HealthCheckId() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("health_check_id"))
}

// Id returns a reference to field id of aws_route53_record.
func (rr route53RecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("id"))
}

// MultivalueAnswerRoutingPolicy returns a reference to field multivalue_answer_routing_policy of aws_route53_record.
func (rr route53RecordAttributes) MultivalueAnswerRoutingPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(rr.ref.Append("multivalue_answer_routing_policy"))
}

// Name returns a reference to field name of aws_route53_record.
func (rr route53RecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("name"))
}

// Records returns a reference to field records of aws_route53_record.
func (rr route53RecordAttributes) Records() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rr.ref.Append("records"))
}

// SetIdentifier returns a reference to field set_identifier of aws_route53_record.
func (rr route53RecordAttributes) SetIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("set_identifier"))
}

// Ttl returns a reference to field ttl of aws_route53_record.
func (rr route53RecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(rr.ref.Append("ttl"))
}

// Type returns a reference to field type of aws_route53_record.
func (rr route53RecordAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("type"))
}

// ZoneId returns a reference to field zone_id of aws_route53_record.
func (rr route53RecordAttributes) ZoneId() terra.StringValue {
	return terra.ReferenceAsString(rr.ref.Append("zone_id"))
}

func (rr route53RecordAttributes) Alias() terra.ListValue[route53record.AliasAttributes] {
	return terra.ReferenceAsList[route53record.AliasAttributes](rr.ref.Append("alias"))
}

func (rr route53RecordAttributes) CidrRoutingPolicy() terra.ListValue[route53record.CidrRoutingPolicyAttributes] {
	return terra.ReferenceAsList[route53record.CidrRoutingPolicyAttributes](rr.ref.Append("cidr_routing_policy"))
}

func (rr route53RecordAttributes) FailoverRoutingPolicy() terra.ListValue[route53record.FailoverRoutingPolicyAttributes] {
	return terra.ReferenceAsList[route53record.FailoverRoutingPolicyAttributes](rr.ref.Append("failover_routing_policy"))
}

func (rr route53RecordAttributes) GeolocationRoutingPolicy() terra.ListValue[route53record.GeolocationRoutingPolicyAttributes] {
	return terra.ReferenceAsList[route53record.GeolocationRoutingPolicyAttributes](rr.ref.Append("geolocation_routing_policy"))
}

func (rr route53RecordAttributes) LatencyRoutingPolicy() terra.ListValue[route53record.LatencyRoutingPolicyAttributes] {
	return terra.ReferenceAsList[route53record.LatencyRoutingPolicyAttributes](rr.ref.Append("latency_routing_policy"))
}

func (rr route53RecordAttributes) WeightedRoutingPolicy() terra.ListValue[route53record.WeightedRoutingPolicyAttributes] {
	return terra.ReferenceAsList[route53record.WeightedRoutingPolicyAttributes](rr.ref.Append("weighted_routing_policy"))
}

type route53RecordState struct {
	AllowOverwrite                bool                                          `json:"allow_overwrite"`
	Fqdn                          string                                        `json:"fqdn"`
	HealthCheckId                 string                                        `json:"health_check_id"`
	Id                            string                                        `json:"id"`
	MultivalueAnswerRoutingPolicy bool                                          `json:"multivalue_answer_routing_policy"`
	Name                          string                                        `json:"name"`
	Records                       []string                                      `json:"records"`
	SetIdentifier                 string                                        `json:"set_identifier"`
	Ttl                           float64                                       `json:"ttl"`
	Type                          string                                        `json:"type"`
	ZoneId                        string                                        `json:"zone_id"`
	Alias                         []route53record.AliasState                    `json:"alias"`
	CidrRoutingPolicy             []route53record.CidrRoutingPolicyState        `json:"cidr_routing_policy"`
	FailoverRoutingPolicy         []route53record.FailoverRoutingPolicyState    `json:"failover_routing_policy"`
	GeolocationRoutingPolicy      []route53record.GeolocationRoutingPolicyState `json:"geolocation_routing_policy"`
	LatencyRoutingPolicy          []route53record.LatencyRoutingPolicyState     `json:"latency_routing_policy"`
	WeightedRoutingPolicy         []route53record.WeightedRoutingPolicyState    `json:"weighted_routing_policy"`
}
