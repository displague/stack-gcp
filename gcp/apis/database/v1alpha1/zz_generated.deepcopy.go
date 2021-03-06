// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// autogenerated by controller-gen object, do not modify manually

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstance) DeepCopyInto(out *CloudsqlInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstance.
func (in *CloudsqlInstance) DeepCopy() *CloudsqlInstance {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudsqlInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstanceClass) DeepCopyInto(out *CloudsqlInstanceClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.SpecTemplate.DeepCopyInto(&out.SpecTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstanceClass.
func (in *CloudsqlInstanceClass) DeepCopy() *CloudsqlInstanceClass {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstanceClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudsqlInstanceClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstanceClassList) DeepCopyInto(out *CloudsqlInstanceClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudsqlInstanceClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstanceClassList.
func (in *CloudsqlInstanceClassList) DeepCopy() *CloudsqlInstanceClassList {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstanceClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudsqlInstanceClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstanceClassSpecTemplate) DeepCopyInto(out *CloudsqlInstanceClassSpecTemplate) {
	*out = *in
	in.ResourceClassSpecTemplate.DeepCopyInto(&out.ResourceClassSpecTemplate)
	in.CloudsqlInstanceParameters.DeepCopyInto(&out.CloudsqlInstanceParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstanceClassSpecTemplate.
func (in *CloudsqlInstanceClassSpecTemplate) DeepCopy() *CloudsqlInstanceClassSpecTemplate {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstanceClassSpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstanceList) DeepCopyInto(out *CloudsqlInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudsqlInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstanceList.
func (in *CloudsqlInstanceList) DeepCopy() *CloudsqlInstanceList {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudsqlInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstanceParameters) DeepCopyInto(out *CloudsqlInstanceParameters) {
	*out = *in
	if in.AuthorizedNetworks != nil {
		in, out := &in.AuthorizedNetworks, &out.AuthorizedNetworks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstanceParameters.
func (in *CloudsqlInstanceParameters) DeepCopy() *CloudsqlInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstanceSpec) DeepCopyInto(out *CloudsqlInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.CloudsqlInstanceParameters.DeepCopyInto(&out.CloudsqlInstanceParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstanceSpec.
func (in *CloudsqlInstanceSpec) DeepCopy() *CloudsqlInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstanceStatus) DeepCopyInto(out *CloudsqlInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstanceStatus.
func (in *CloudsqlInstanceStatus) DeepCopy() *CloudsqlInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstanceStatus)
	in.DeepCopyInto(out)
	return out
}
