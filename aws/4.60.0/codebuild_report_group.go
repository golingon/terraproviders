// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	codebuildreportgroup "github.com/golingon/terraproviders/aws/4.60.0/codebuildreportgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewCodebuildReportGroup(name string, args CodebuildReportGroupArgs) *CodebuildReportGroup {
	return &CodebuildReportGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodebuildReportGroup)(nil)

type CodebuildReportGroup struct {
	Name  string
	Args  CodebuildReportGroupArgs
	state *codebuildReportGroupState
}

func (crg *CodebuildReportGroup) Type() string {
	return "aws_codebuild_report_group"
}

func (crg *CodebuildReportGroup) LocalName() string {
	return crg.Name
}

func (crg *CodebuildReportGroup) Configuration() interface{} {
	return crg.Args
}

func (crg *CodebuildReportGroup) Attributes() codebuildReportGroupAttributes {
	return codebuildReportGroupAttributes{ref: terra.ReferenceResource(crg)}
}

func (crg *CodebuildReportGroup) ImportState(av io.Reader) error {
	crg.state = &codebuildReportGroupState{}
	if err := json.NewDecoder(av).Decode(crg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crg.Type(), crg.LocalName(), err)
	}
	return nil
}

func (crg *CodebuildReportGroup) State() (*codebuildReportGroupState, bool) {
	return crg.state, crg.state != nil
}

func (crg *CodebuildReportGroup) StateMust() *codebuildReportGroupState {
	if crg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crg.Type(), crg.LocalName()))
	}
	return crg.state
}

func (crg *CodebuildReportGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(crg)
}

type CodebuildReportGroupArgs struct {
	// DeleteReports: bool, optional
	DeleteReports terra.BoolValue `hcl:"delete_reports,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// ExportConfig: required
	ExportConfig *codebuildreportgroup.ExportConfig `hcl:"export_config,block" validate:"required"`
	// DependsOn contains resources that CodebuildReportGroup depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type codebuildReportGroupAttributes struct {
	ref terra.Reference
}

func (crg codebuildReportGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(crg.ref.Append("arn"))
}

func (crg codebuildReportGroupAttributes) Created() terra.StringValue {
	return terra.ReferenceString(crg.ref.Append("created"))
}

func (crg codebuildReportGroupAttributes) DeleteReports() terra.BoolValue {
	return terra.ReferenceBool(crg.ref.Append("delete_reports"))
}

func (crg codebuildReportGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(crg.ref.Append("id"))
}

func (crg codebuildReportGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(crg.ref.Append("name"))
}

func (crg codebuildReportGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](crg.ref.Append("tags"))
}

func (crg codebuildReportGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](crg.ref.Append("tags_all"))
}

func (crg codebuildReportGroupAttributes) Type() terra.StringValue {
	return terra.ReferenceString(crg.ref.Append("type"))
}

func (crg codebuildReportGroupAttributes) ExportConfig() terra.ListValue[codebuildreportgroup.ExportConfigAttributes] {
	return terra.ReferenceList[codebuildreportgroup.ExportConfigAttributes](crg.ref.Append("export_config"))
}

type codebuildReportGroupState struct {
	Arn           string                                   `json:"arn"`
	Created       string                                   `json:"created"`
	DeleteReports bool                                     `json:"delete_reports"`
	Id            string                                   `json:"id"`
	Name          string                                   `json:"name"`
	Tags          map[string]string                        `json:"tags"`
	TagsAll       map[string]string                        `json:"tags_all"`
	Type          string                                   `json:"type"`
	ExportConfig  []codebuildreportgroup.ExportConfigState `json:"export_config"`
}
