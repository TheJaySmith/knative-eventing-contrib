// +build !ignore_autogenerated

/*
Copyright 2020 The Knative Authors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "knative.dev/pkg/apis/duck/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusSource) DeepCopyInto(out *PrometheusSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusSource.
func (in *PrometheusSource) DeepCopy() *PrometheusSource {
	if in == nil {
		return nil
	}
	out := new(PrometheusSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrometheusSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusSourceList) DeepCopyInto(out *PrometheusSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrometheusSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusSourceList.
func (in *PrometheusSourceList) DeepCopy() *PrometheusSourceList {
	if in == nil {
		return nil
	}
	out := new(PrometheusSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrometheusSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusSourceSpec) DeepCopyInto(out *PrometheusSourceSpec) {
	*out = *in
	if in.Sink != nil {
		in, out := &in.Sink, &out.Sink
		*out = new(v1.Destination)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusSourceSpec.
func (in *PrometheusSourceSpec) DeepCopy() *PrometheusSourceSpec {
	if in == nil {
		return nil
	}
	out := new(PrometheusSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusSourceStatus) DeepCopyInto(out *PrometheusSourceStatus) {
	*out = *in
	in.SourceStatus.DeepCopyInto(&out.SourceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusSourceStatus.
func (in *PrometheusSourceStatus) DeepCopy() *PrometheusSourceStatus {
	if in == nil {
		return nil
	}
	out := new(PrometheusSourceStatus)
	in.DeepCopyInto(out)
	return out
}
