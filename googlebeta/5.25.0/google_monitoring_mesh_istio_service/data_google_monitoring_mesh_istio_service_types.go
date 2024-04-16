// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_monitoring_mesh_istio_service

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTelemetryAttributes struct {
	ref terra.Reference
}

func (t DataTelemetryAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTelemetryAttributes) InternalWithRef(ref terra.Reference) DataTelemetryAttributes {
	return DataTelemetryAttributes{ref: ref}
}

func (t DataTelemetryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTelemetryAttributes) ResourceName() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("resource_name"))
}

type DataTelemetryState struct {
	ResourceName string `json:"resource_name"`
}
