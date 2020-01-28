// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/external/smi/traffictarget/v1alpha1/traffic-target.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for TrafficTarget
func (this *TrafficTarget) MarshalJSON() ([]byte, error) {
	str, err := TrafficTargetMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficTarget
func (this *TrafficTarget) UnmarshalJSON(b []byte) error {
	return TrafficTargetUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficTargetSpec
func (this *TrafficTargetSpec) MarshalJSON() ([]byte, error) {
	str, err := TrafficTargetMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficTargetSpec
func (this *TrafficTargetSpec) UnmarshalJSON(b []byte) error {
	return TrafficTargetUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for IdentityBindingSubject
func (this *IdentityBindingSubject) MarshalJSON() ([]byte, error) {
	str, err := TrafficTargetMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for IdentityBindingSubject
func (this *IdentityBindingSubject) UnmarshalJSON(b []byte) error {
	return TrafficTargetUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	TrafficTargetMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	TrafficTargetUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)