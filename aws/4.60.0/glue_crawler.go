// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	gluecrawler "github.com/golingon/terraproviders/aws/4.60.0/gluecrawler"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewGlueCrawler(name string, args GlueCrawlerArgs) *GlueCrawler {
	return &GlueCrawler{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlueCrawler)(nil)

type GlueCrawler struct {
	Name  string
	Args  GlueCrawlerArgs
	state *glueCrawlerState
}

func (gc *GlueCrawler) Type() string {
	return "aws_glue_crawler"
}

func (gc *GlueCrawler) LocalName() string {
	return gc.Name
}

func (gc *GlueCrawler) Configuration() interface{} {
	return gc.Args
}

func (gc *GlueCrawler) Attributes() glueCrawlerAttributes {
	return glueCrawlerAttributes{ref: terra.ReferenceResource(gc)}
}

func (gc *GlueCrawler) ImportState(av io.Reader) error {
	gc.state = &glueCrawlerState{}
	if err := json.NewDecoder(av).Decode(gc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gc.Type(), gc.LocalName(), err)
	}
	return nil
}

func (gc *GlueCrawler) State() (*glueCrawlerState, bool) {
	return gc.state, gc.state != nil
}

func (gc *GlueCrawler) StateMust() *glueCrawlerState {
	if gc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gc.Type(), gc.LocalName()))
	}
	return gc.state
}

func (gc *GlueCrawler) DependOn() terra.Reference {
	return terra.ReferenceResource(gc)
}

type GlueCrawlerArgs struct {
	// Classifiers: list of string, optional
	Classifiers terra.ListValue[terra.StringValue] `hcl:"classifiers,attr"`
	// Configuration: string, optional
	Configuration terra.StringValue `hcl:"configuration,attr"`
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Schedule: string, optional
	Schedule terra.StringValue `hcl:"schedule,attr"`
	// SecurityConfiguration: string, optional
	SecurityConfiguration terra.StringValue `hcl:"security_configuration,attr"`
	// TablePrefix: string, optional
	TablePrefix terra.StringValue `hcl:"table_prefix,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// CatalogTarget: min=0
	CatalogTarget []gluecrawler.CatalogTarget `hcl:"catalog_target,block" validate:"min=0"`
	// DeltaTarget: min=0
	DeltaTarget []gluecrawler.DeltaTarget `hcl:"delta_target,block" validate:"min=0"`
	// DynamodbTarget: min=0
	DynamodbTarget []gluecrawler.DynamodbTarget `hcl:"dynamodb_target,block" validate:"min=0"`
	// JdbcTarget: min=0
	JdbcTarget []gluecrawler.JdbcTarget `hcl:"jdbc_target,block" validate:"min=0"`
	// LakeFormationConfiguration: optional
	LakeFormationConfiguration *gluecrawler.LakeFormationConfiguration `hcl:"lake_formation_configuration,block"`
	// LineageConfiguration: optional
	LineageConfiguration *gluecrawler.LineageConfiguration `hcl:"lineage_configuration,block"`
	// MongodbTarget: min=0
	MongodbTarget []gluecrawler.MongodbTarget `hcl:"mongodb_target,block" validate:"min=0"`
	// RecrawlPolicy: optional
	RecrawlPolicy *gluecrawler.RecrawlPolicy `hcl:"recrawl_policy,block"`
	// S3Target: min=0
	S3Target []gluecrawler.S3Target `hcl:"s3_target,block" validate:"min=0"`
	// SchemaChangePolicy: optional
	SchemaChangePolicy *gluecrawler.SchemaChangePolicy `hcl:"schema_change_policy,block"`
	// DependsOn contains resources that GlueCrawler depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type glueCrawlerAttributes struct {
	ref terra.Reference
}

func (gc glueCrawlerAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(gc.ref.Append("arn"))
}

func (gc glueCrawlerAttributes) Classifiers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](gc.ref.Append("classifiers"))
}

func (gc glueCrawlerAttributes) Configuration() terra.StringValue {
	return terra.ReferenceString(gc.ref.Append("configuration"))
}

func (gc glueCrawlerAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceString(gc.ref.Append("database_name"))
}

func (gc glueCrawlerAttributes) Description() terra.StringValue {
	return terra.ReferenceString(gc.ref.Append("description"))
}

func (gc glueCrawlerAttributes) Id() terra.StringValue {
	return terra.ReferenceString(gc.ref.Append("id"))
}

func (gc glueCrawlerAttributes) Name() terra.StringValue {
	return terra.ReferenceString(gc.ref.Append("name"))
}

func (gc glueCrawlerAttributes) Role() terra.StringValue {
	return terra.ReferenceString(gc.ref.Append("role"))
}

func (gc glueCrawlerAttributes) Schedule() terra.StringValue {
	return terra.ReferenceString(gc.ref.Append("schedule"))
}

func (gc glueCrawlerAttributes) SecurityConfiguration() terra.StringValue {
	return terra.ReferenceString(gc.ref.Append("security_configuration"))
}

func (gc glueCrawlerAttributes) TablePrefix() terra.StringValue {
	return terra.ReferenceString(gc.ref.Append("table_prefix"))
}

func (gc glueCrawlerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](gc.ref.Append("tags"))
}

func (gc glueCrawlerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](gc.ref.Append("tags_all"))
}

func (gc glueCrawlerAttributes) CatalogTarget() terra.ListValue[gluecrawler.CatalogTargetAttributes] {
	return terra.ReferenceList[gluecrawler.CatalogTargetAttributes](gc.ref.Append("catalog_target"))
}

func (gc glueCrawlerAttributes) DeltaTarget() terra.ListValue[gluecrawler.DeltaTargetAttributes] {
	return terra.ReferenceList[gluecrawler.DeltaTargetAttributes](gc.ref.Append("delta_target"))
}

func (gc glueCrawlerAttributes) DynamodbTarget() terra.ListValue[gluecrawler.DynamodbTargetAttributes] {
	return terra.ReferenceList[gluecrawler.DynamodbTargetAttributes](gc.ref.Append("dynamodb_target"))
}

func (gc glueCrawlerAttributes) JdbcTarget() terra.ListValue[gluecrawler.JdbcTargetAttributes] {
	return terra.ReferenceList[gluecrawler.JdbcTargetAttributes](gc.ref.Append("jdbc_target"))
}

func (gc glueCrawlerAttributes) LakeFormationConfiguration() terra.ListValue[gluecrawler.LakeFormationConfigurationAttributes] {
	return terra.ReferenceList[gluecrawler.LakeFormationConfigurationAttributes](gc.ref.Append("lake_formation_configuration"))
}

func (gc glueCrawlerAttributes) LineageConfiguration() terra.ListValue[gluecrawler.LineageConfigurationAttributes] {
	return terra.ReferenceList[gluecrawler.LineageConfigurationAttributes](gc.ref.Append("lineage_configuration"))
}

func (gc glueCrawlerAttributes) MongodbTarget() terra.ListValue[gluecrawler.MongodbTargetAttributes] {
	return terra.ReferenceList[gluecrawler.MongodbTargetAttributes](gc.ref.Append("mongodb_target"))
}

func (gc glueCrawlerAttributes) RecrawlPolicy() terra.ListValue[gluecrawler.RecrawlPolicyAttributes] {
	return terra.ReferenceList[gluecrawler.RecrawlPolicyAttributes](gc.ref.Append("recrawl_policy"))
}

func (gc glueCrawlerAttributes) S3Target() terra.ListValue[gluecrawler.S3TargetAttributes] {
	return terra.ReferenceList[gluecrawler.S3TargetAttributes](gc.ref.Append("s3_target"))
}

func (gc glueCrawlerAttributes) SchemaChangePolicy() terra.ListValue[gluecrawler.SchemaChangePolicyAttributes] {
	return terra.ReferenceList[gluecrawler.SchemaChangePolicyAttributes](gc.ref.Append("schema_change_policy"))
}

type glueCrawlerState struct {
	Arn                        string                                        `json:"arn"`
	Classifiers                []string                                      `json:"classifiers"`
	Configuration              string                                        `json:"configuration"`
	DatabaseName               string                                        `json:"database_name"`
	Description                string                                        `json:"description"`
	Id                         string                                        `json:"id"`
	Name                       string                                        `json:"name"`
	Role                       string                                        `json:"role"`
	Schedule                   string                                        `json:"schedule"`
	SecurityConfiguration      string                                        `json:"security_configuration"`
	TablePrefix                string                                        `json:"table_prefix"`
	Tags                       map[string]string                             `json:"tags"`
	TagsAll                    map[string]string                             `json:"tags_all"`
	CatalogTarget              []gluecrawler.CatalogTargetState              `json:"catalog_target"`
	DeltaTarget                []gluecrawler.DeltaTargetState                `json:"delta_target"`
	DynamodbTarget             []gluecrawler.DynamodbTargetState             `json:"dynamodb_target"`
	JdbcTarget                 []gluecrawler.JdbcTargetState                 `json:"jdbc_target"`
	LakeFormationConfiguration []gluecrawler.LakeFormationConfigurationState `json:"lake_formation_configuration"`
	LineageConfiguration       []gluecrawler.LineageConfigurationState       `json:"lineage_configuration"`
	MongodbTarget              []gluecrawler.MongodbTargetState              `json:"mongodb_target"`
	RecrawlPolicy              []gluecrawler.RecrawlPolicyState              `json:"recrawl_policy"`
	S3Target                   []gluecrawler.S3TargetState                   `json:"s3_target"`
	SchemaChangePolicy         []gluecrawler.SchemaChangePolicyState         `json:"schema_change_policy"`
}
