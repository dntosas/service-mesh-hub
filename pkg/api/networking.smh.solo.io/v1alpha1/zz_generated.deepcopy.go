// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for networking.smh.solo.io/v1alpha1 resources

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for TrafficPolicy

func (in *TrafficPolicy) DeepCopyInto(out *TrafficPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

func (in *TrafficPolicy) DeepCopy() *TrafficPolicy {
	if in == nil {
		return nil
	}
	out := new(TrafficPolicy)
	in.DeepCopyInto(out)
	return out
}

func (in *TrafficPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *TrafficPolicyList) DeepCopyInto(out *TrafficPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrafficPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *TrafficPolicyList) DeepCopy() *TrafficPolicyList {
	if in == nil {
		return nil
	}
	out := new(TrafficPolicyList)
	in.DeepCopyInto(out)
	return out
}

func (in *TrafficPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for AccessControlPolicy

func (in *AccessControlPolicy) DeepCopyInto(out *AccessControlPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

func (in *AccessControlPolicy) DeepCopy() *AccessControlPolicy {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicy)
	in.DeepCopyInto(out)
	return out
}

func (in *AccessControlPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *AccessControlPolicyList) DeepCopyInto(out *AccessControlPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessControlPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *AccessControlPolicyList) DeepCopy() *AccessControlPolicyList {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyList)
	in.DeepCopyInto(out)
	return out
}

func (in *AccessControlPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for VirtualMesh

func (in *VirtualMesh) DeepCopyInto(out *VirtualMesh) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

func (in *VirtualMesh) DeepCopy() *VirtualMesh {
	if in == nil {
		return nil
	}
	out := new(VirtualMesh)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMesh) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VirtualMeshList) DeepCopyInto(out *VirtualMeshList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMesh, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *VirtualMeshList) DeepCopy() *VirtualMeshList {
	if in == nil {
		return nil
	}
	out := new(VirtualMeshList)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMeshList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for FailoverService

func (in *FailoverService) DeepCopyInto(out *FailoverService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

func (in *FailoverService) DeepCopy() *FailoverService {
	if in == nil {
		return nil
	}
	out := new(FailoverService)
	in.DeepCopyInto(out)
	return out
}

func (in *FailoverService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *FailoverServiceList) DeepCopyInto(out *FailoverServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FailoverService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *FailoverServiceList) DeepCopy() *FailoverServiceList {
	if in == nil {
		return nil
	}
	out := new(FailoverServiceList)
	in.DeepCopyInto(out)
	return out
}

func (in *FailoverServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
