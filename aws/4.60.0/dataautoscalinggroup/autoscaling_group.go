// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataautoscalinggroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type LaunchTemplate struct{}

type LaunchTemplateAttributes struct {
	ref terra.Reference
}

func (lt LaunchTemplateAttributes) InternalRef() (terra.Reference, error) {
	return lt.ref, nil
}

func (lt LaunchTemplateAttributes) InternalWithRef(ref terra.Reference) LaunchTemplateAttributes {
	return LaunchTemplateAttributes{ref: ref}
}

func (lt LaunchTemplateAttributes) InternalTokens() hclwrite.Tokens {
	return lt.ref.InternalTokens()
}

func (lt LaunchTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("id"))
}

func (lt LaunchTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("name"))
}

func (lt LaunchTemplateAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("version"))
}

type LaunchTemplateState struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}
