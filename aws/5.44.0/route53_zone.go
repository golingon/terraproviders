// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	route53zone "github.com/golingon/terraproviders/aws/5.44.0/route53zone"
	"io"
)

// NewRoute53Zone creates a new instance of [Route53Zone].
func NewRoute53Zone(name string, args Route53ZoneArgs) *Route53Zone {
	return &Route53Zone{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53Zone)(nil)

// Route53Zone represents the Terraform resource aws_route53_zone.
type Route53Zone struct {
	Name      string
	Args      Route53ZoneArgs
	state     *route53ZoneState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53Zone].
func (rz *Route53Zone) Type() string {
	return "aws_route53_zone"
}

// LocalName returns the local name for [Route53Zone].
func (rz *Route53Zone) LocalName() string {
	return rz.Name
}

// Configuration returns the configuration (args) for [Route53Zone].
func (rz *Route53Zone) Configuration() interface{} {
	return rz.Args
}

// DependOn is used for other resources to depend on [Route53Zone].
func (rz *Route53Zone) DependOn() terra.Reference {
	return terra.ReferenceResource(rz)
}

// Dependencies returns the list of resources [Route53Zone] depends_on.
func (rz *Route53Zone) Dependencies() terra.Dependencies {
	return rz.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53Zone].
func (rz *Route53Zone) LifecycleManagement() *terra.Lifecycle {
	return rz.Lifecycle
}

// Attributes returns the attributes for [Route53Zone].
func (rz *Route53Zone) Attributes() route53ZoneAttributes {
	return route53ZoneAttributes{ref: terra.ReferenceResource(rz)}
}

// ImportState imports the given attribute values into [Route53Zone]'s state.
func (rz *Route53Zone) ImportState(av io.Reader) error {
	rz.state = &route53ZoneState{}
	if err := json.NewDecoder(av).Decode(rz.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rz.Type(), rz.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53Zone] has state.
func (rz *Route53Zone) State() (*route53ZoneState, bool) {
	return rz.state, rz.state != nil
}

// StateMust returns the state for [Route53Zone]. Panics if the state is nil.
func (rz *Route53Zone) StateMust() *route53ZoneState {
	if rz.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rz.Type(), rz.LocalName()))
	}
	return rz.state
}

// Route53ZoneArgs contains the configurations for aws_route53_zone.
type Route53ZoneArgs struct {
	// Comment: string, optional
	Comment terra.StringValue `hcl:"comment,attr"`
	// DelegationSetId: string, optional
	DelegationSetId terra.StringValue `hcl:"delegation_set_id,attr"`
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Vpc: min=0
	Vpc []route53zone.Vpc `hcl:"vpc,block" validate:"min=0"`
}
type route53ZoneAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_route53_zone.
func (rz route53ZoneAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rz.ref.Append("arn"))
}

// Comment returns a reference to field comment of aws_route53_zone.
func (rz route53ZoneAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(rz.ref.Append("comment"))
}

// DelegationSetId returns a reference to field delegation_set_id of aws_route53_zone.
func (rz route53ZoneAttributes) DelegationSetId() terra.StringValue {
	return terra.ReferenceAsString(rz.ref.Append("delegation_set_id"))
}

// ForceDestroy returns a reference to field force_destroy of aws_route53_zone.
func (rz route53ZoneAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(rz.ref.Append("force_destroy"))
}

// Id returns a reference to field id of aws_route53_zone.
func (rz route53ZoneAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rz.ref.Append("id"))
}

// Name returns a reference to field name of aws_route53_zone.
func (rz route53ZoneAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rz.ref.Append("name"))
}

// NameServers returns a reference to field name_servers of aws_route53_zone.
func (rz route53ZoneAttributes) NameServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rz.ref.Append("name_servers"))
}

// PrimaryNameServer returns a reference to field primary_name_server of aws_route53_zone.
func (rz route53ZoneAttributes) PrimaryNameServer() terra.StringValue {
	return terra.ReferenceAsString(rz.ref.Append("primary_name_server"))
}

// Tags returns a reference to field tags of aws_route53_zone.
func (rz route53ZoneAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rz.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_route53_zone.
func (rz route53ZoneAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rz.ref.Append("tags_all"))
}

// ZoneId returns a reference to field zone_id of aws_route53_zone.
func (rz route53ZoneAttributes) ZoneId() terra.StringValue {
	return terra.ReferenceAsString(rz.ref.Append("zone_id"))
}

func (rz route53ZoneAttributes) Vpc() terra.SetValue[route53zone.VpcAttributes] {
	return terra.ReferenceAsSet[route53zone.VpcAttributes](rz.ref.Append("vpc"))
}

type route53ZoneState struct {
	Arn               string                 `json:"arn"`
	Comment           string                 `json:"comment"`
	DelegationSetId   string                 `json:"delegation_set_id"`
	ForceDestroy      bool                   `json:"force_destroy"`
	Id                string                 `json:"id"`
	Name              string                 `json:"name"`
	NameServers       []string               `json:"name_servers"`
	PrimaryNameServer string                 `json:"primary_name_server"`
	Tags              map[string]string      `json:"tags"`
	TagsAll           map[string]string      `json:"tags_all"`
	ZoneId            string                 `json:"zone_id"`
	Vpc               []route53zone.VpcState `json:"vpc"`
}
