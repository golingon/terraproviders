// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewElasticacheUserGroup(name string, args ElasticacheUserGroupArgs) *ElasticacheUserGroup {
	return &ElasticacheUserGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElasticacheUserGroup)(nil)

type ElasticacheUserGroup struct {
	Name  string
	Args  ElasticacheUserGroupArgs
	state *elasticacheUserGroupState
}

func (eug *ElasticacheUserGroup) Type() string {
	return "aws_elasticache_user_group"
}

func (eug *ElasticacheUserGroup) LocalName() string {
	return eug.Name
}

func (eug *ElasticacheUserGroup) Configuration() interface{} {
	return eug.Args
}

func (eug *ElasticacheUserGroup) Attributes() elasticacheUserGroupAttributes {
	return elasticacheUserGroupAttributes{ref: terra.ReferenceResource(eug)}
}

func (eug *ElasticacheUserGroup) ImportState(av io.Reader) error {
	eug.state = &elasticacheUserGroupState{}
	if err := json.NewDecoder(av).Decode(eug.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", eug.Type(), eug.LocalName(), err)
	}
	return nil
}

func (eug *ElasticacheUserGroup) State() (*elasticacheUserGroupState, bool) {
	return eug.state, eug.state != nil
}

func (eug *ElasticacheUserGroup) StateMust() *elasticacheUserGroupState {
	if eug.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", eug.Type(), eug.LocalName()))
	}
	return eug.state
}

func (eug *ElasticacheUserGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(eug)
}

type ElasticacheUserGroupArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Engine: string, required
	Engine terra.StringValue `hcl:"engine,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// UserGroupId: string, required
	UserGroupId terra.StringValue `hcl:"user_group_id,attr" validate:"required"`
	// UserIds: set of string, optional
	UserIds terra.SetValue[terra.StringValue] `hcl:"user_ids,attr"`
	// DependsOn contains resources that ElasticacheUserGroup depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type elasticacheUserGroupAttributes struct {
	ref terra.Reference
}

func (eug elasticacheUserGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(eug.ref.Append("arn"))
}

func (eug elasticacheUserGroupAttributes) Engine() terra.StringValue {
	return terra.ReferenceString(eug.ref.Append("engine"))
}

func (eug elasticacheUserGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(eug.ref.Append("id"))
}

func (eug elasticacheUserGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](eug.ref.Append("tags"))
}

func (eug elasticacheUserGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](eug.ref.Append("tags_all"))
}

func (eug elasticacheUserGroupAttributes) UserGroupId() terra.StringValue {
	return terra.ReferenceString(eug.ref.Append("user_group_id"))
}

func (eug elasticacheUserGroupAttributes) UserIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](eug.ref.Append("user_ids"))
}

type elasticacheUserGroupState struct {
	Arn         string            `json:"arn"`
	Engine      string            `json:"engine"`
	Id          string            `json:"id"`
	Tags        map[string]string `json:"tags"`
	TagsAll     map[string]string `json:"tags_all"`
	UserGroupId string            `json:"user_group_id"`
	UserIds     []string          `json:"user_ids"`
}
