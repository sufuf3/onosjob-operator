// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowsDevice) DeepCopyInto(out *FlowsDevice) {
	*out = *in
	if in.Instructions != nil {
		in, out := &in.Instructions, &out.Instructions
		*out = make([]FlowsDeviceInstructions, len(*in))
		copy(*out, *in)
	}
	if in.Criteria != nil {
		in, out := &in.Criteria, &out.Criteria
		*out = make([]FlowsDeviceCriteria, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowsDevice.
func (in *FlowsDevice) DeepCopy() *FlowsDevice {
	if in == nil {
		return nil
	}
	out := new(FlowsDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowsDeviceCriteria) DeepCopyInto(out *FlowsDeviceCriteria) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowsDeviceCriteria.
func (in *FlowsDeviceCriteria) DeepCopy() *FlowsDeviceCriteria {
	if in == nil {
		return nil
	}
	out := new(FlowsDeviceCriteria)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowsDeviceInstructions) DeepCopyInto(out *FlowsDeviceInstructions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowsDeviceInstructions.
func (in *FlowsDeviceInstructions) DeepCopy() *FlowsDeviceInstructions {
	if in == nil {
		return nil
	}
	out := new(FlowsDeviceInstructions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostLocations) DeepCopyInto(out *HostLocations) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostLocations.
func (in *HostLocations) DeepCopy() *HostLocations {
	if in == nil {
		return nil
	}
	out := new(HostLocations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hosts) DeepCopyInto(out *Hosts) {
	*out = *in
	if in.IpAddresses != nil {
		in, out := &in.IpAddresses, &out.IpAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]HostLocations, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hosts.
func (in *Hosts) DeepCopy() *Hosts {
	if in == nil {
		return nil
	}
	out := new(Hosts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ONOSJob) DeepCopyInto(out *ONOSJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ONOSJob.
func (in *ONOSJob) DeepCopy() *ONOSJob {
	if in == nil {
		return nil
	}
	out := new(ONOSJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ONOSJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ONOSJobList) DeepCopyInto(out *ONOSJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ONOSJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ONOSJobList.
func (in *ONOSJobList) DeepCopy() *ONOSJobList {
	if in == nil {
		return nil
	}
	out := new(ONOSJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ONOSJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ONOSJobSpec) DeepCopyInto(out *ONOSJobSpec) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]Hosts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FlowsDevice != nil {
		in, out := &in.FlowsDevice, &out.FlowsDevice
		*out = make([]FlowsDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ONOSJobSpec.
func (in *ONOSJobSpec) DeepCopy() *ONOSJobSpec {
	if in == nil {
		return nil
	}
	out := new(ONOSJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ONOSJobStatus) DeepCopyInto(out *ONOSJobStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ONOSJobStatus.
func (in *ONOSJobStatus) DeepCopy() *ONOSJobStatus {
	if in == nil {
		return nil
	}
	out := new(ONOSJobStatus)
	in.DeepCopyInto(out)
	return out
}
