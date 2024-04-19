//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionInitParameters) DeepCopyInto(out *ConditionInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionInitParameters.
func (in *ConditionInitParameters) DeepCopy() *ConditionInitParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionObservation) DeepCopyInto(out *ConditionObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionObservation.
func (in *ConditionObservation) DeepCopy() *ConditionObservation {
	if in == nil {
		return nil
	}
	out := new(ConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionParameters) DeepCopyInto(out *ConditionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionParameters.
func (in *ConditionParameters) DeepCopy() *ConditionParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FolderIAMMember) DeepCopyInto(out *FolderIAMMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FolderIAMMember.
func (in *FolderIAMMember) DeepCopy() *FolderIAMMember {
	if in == nil {
		return nil
	}
	out := new(FolderIAMMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FolderIAMMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FolderIAMMemberInitParameters) DeepCopyInto(out *FolderIAMMemberInitParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(ConditionInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.FolderRef != nil {
		in, out := &in.FolderRef, &out.FolderRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FolderSelector != nil {
		in, out := &in.FolderSelector, &out.FolderSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FolderIAMMemberInitParameters.
func (in *FolderIAMMemberInitParameters) DeepCopy() *FolderIAMMemberInitParameters {
	if in == nil {
		return nil
	}
	out := new(FolderIAMMemberInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FolderIAMMemberList) DeepCopyInto(out *FolderIAMMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FolderIAMMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FolderIAMMemberList.
func (in *FolderIAMMemberList) DeepCopy() *FolderIAMMemberList {
	if in == nil {
		return nil
	}
	out := new(FolderIAMMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FolderIAMMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FolderIAMMemberObservation) DeepCopyInto(out *FolderIAMMemberObservation) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(ConditionObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FolderIAMMemberObservation.
func (in *FolderIAMMemberObservation) DeepCopy() *FolderIAMMemberObservation {
	if in == nil {
		return nil
	}
	out := new(FolderIAMMemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FolderIAMMemberParameters) DeepCopyInto(out *FolderIAMMemberParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(ConditionParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.FolderRef != nil {
		in, out := &in.FolderRef, &out.FolderRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FolderSelector != nil {
		in, out := &in.FolderSelector, &out.FolderSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FolderIAMMemberParameters.
func (in *FolderIAMMemberParameters) DeepCopy() *FolderIAMMemberParameters {
	if in == nil {
		return nil
	}
	out := new(FolderIAMMemberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FolderIAMMemberSpec) DeepCopyInto(out *FolderIAMMemberSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FolderIAMMemberSpec.
func (in *FolderIAMMemberSpec) DeepCopy() *FolderIAMMemberSpec {
	if in == nil {
		return nil
	}
	out := new(FolderIAMMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FolderIAMMemberStatus) DeepCopyInto(out *FolderIAMMemberStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FolderIAMMemberStatus.
func (in *FolderIAMMemberStatus) DeepCopy() *FolderIAMMemberStatus {
	if in == nil {
		return nil
	}
	out := new(FolderIAMMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationIAMMember) DeepCopyInto(out *OrganizationIAMMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationIAMMember.
func (in *OrganizationIAMMember) DeepCopy() *OrganizationIAMMember {
	if in == nil {
		return nil
	}
	out := new(OrganizationIAMMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationIAMMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationIAMMemberConditionInitParameters) DeepCopyInto(out *OrganizationIAMMemberConditionInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationIAMMemberConditionInitParameters.
func (in *OrganizationIAMMemberConditionInitParameters) DeepCopy() *OrganizationIAMMemberConditionInitParameters {
	if in == nil {
		return nil
	}
	out := new(OrganizationIAMMemberConditionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationIAMMemberConditionObservation) DeepCopyInto(out *OrganizationIAMMemberConditionObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationIAMMemberConditionObservation.
func (in *OrganizationIAMMemberConditionObservation) DeepCopy() *OrganizationIAMMemberConditionObservation {
	if in == nil {
		return nil
	}
	out := new(OrganizationIAMMemberConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationIAMMemberConditionParameters) DeepCopyInto(out *OrganizationIAMMemberConditionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationIAMMemberConditionParameters.
func (in *OrganizationIAMMemberConditionParameters) DeepCopy() *OrganizationIAMMemberConditionParameters {
	if in == nil {
		return nil
	}
	out := new(OrganizationIAMMemberConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationIAMMemberInitParameters) DeepCopyInto(out *OrganizationIAMMemberInitParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(OrganizationIAMMemberConditionInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationIAMMemberInitParameters.
func (in *OrganizationIAMMemberInitParameters) DeepCopy() *OrganizationIAMMemberInitParameters {
	if in == nil {
		return nil
	}
	out := new(OrganizationIAMMemberInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationIAMMemberList) DeepCopyInto(out *OrganizationIAMMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OrganizationIAMMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationIAMMemberList.
func (in *OrganizationIAMMemberList) DeepCopy() *OrganizationIAMMemberList {
	if in == nil {
		return nil
	}
	out := new(OrganizationIAMMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationIAMMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationIAMMemberObservation) DeepCopyInto(out *OrganizationIAMMemberObservation) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(OrganizationIAMMemberConditionObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationIAMMemberObservation.
func (in *OrganizationIAMMemberObservation) DeepCopy() *OrganizationIAMMemberObservation {
	if in == nil {
		return nil
	}
	out := new(OrganizationIAMMemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationIAMMemberParameters) DeepCopyInto(out *OrganizationIAMMemberParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(OrganizationIAMMemberConditionParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationIAMMemberParameters.
func (in *OrganizationIAMMemberParameters) DeepCopy() *OrganizationIAMMemberParameters {
	if in == nil {
		return nil
	}
	out := new(OrganizationIAMMemberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationIAMMemberSpec) DeepCopyInto(out *OrganizationIAMMemberSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationIAMMemberSpec.
func (in *OrganizationIAMMemberSpec) DeepCopy() *OrganizationIAMMemberSpec {
	if in == nil {
		return nil
	}
	out := new(OrganizationIAMMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationIAMMemberStatus) DeepCopyInto(out *OrganizationIAMMemberStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationIAMMemberStatus.
func (in *OrganizationIAMMemberStatus) DeepCopy() *OrganizationIAMMemberStatus {
	if in == nil {
		return nil
	}
	out := new(OrganizationIAMMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectIAMMember) DeepCopyInto(out *ProjectIAMMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectIAMMember.
func (in *ProjectIAMMember) DeepCopy() *ProjectIAMMember {
	if in == nil {
		return nil
	}
	out := new(ProjectIAMMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectIAMMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectIAMMemberConditionInitParameters) DeepCopyInto(out *ProjectIAMMemberConditionInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectIAMMemberConditionInitParameters.
func (in *ProjectIAMMemberConditionInitParameters) DeepCopy() *ProjectIAMMemberConditionInitParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectIAMMemberConditionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectIAMMemberConditionObservation) DeepCopyInto(out *ProjectIAMMemberConditionObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectIAMMemberConditionObservation.
func (in *ProjectIAMMemberConditionObservation) DeepCopy() *ProjectIAMMemberConditionObservation {
	if in == nil {
		return nil
	}
	out := new(ProjectIAMMemberConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectIAMMemberConditionParameters) DeepCopyInto(out *ProjectIAMMemberConditionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectIAMMemberConditionParameters.
func (in *ProjectIAMMemberConditionParameters) DeepCopy() *ProjectIAMMemberConditionParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectIAMMemberConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectIAMMemberInitParameters) DeepCopyInto(out *ProjectIAMMemberInitParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(ProjectIAMMemberConditionInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ProjectSelector != nil {
		in, out := &in.ProjectSelector, &out.ProjectSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectIAMMemberInitParameters.
func (in *ProjectIAMMemberInitParameters) DeepCopy() *ProjectIAMMemberInitParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectIAMMemberInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectIAMMemberList) DeepCopyInto(out *ProjectIAMMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProjectIAMMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectIAMMemberList.
func (in *ProjectIAMMemberList) DeepCopy() *ProjectIAMMemberList {
	if in == nil {
		return nil
	}
	out := new(ProjectIAMMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectIAMMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectIAMMemberObservation) DeepCopyInto(out *ProjectIAMMemberObservation) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(ProjectIAMMemberConditionObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectIAMMemberObservation.
func (in *ProjectIAMMemberObservation) DeepCopy() *ProjectIAMMemberObservation {
	if in == nil {
		return nil
	}
	out := new(ProjectIAMMemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectIAMMemberParameters) DeepCopyInto(out *ProjectIAMMemberParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(ProjectIAMMemberConditionParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ProjectSelector != nil {
		in, out := &in.ProjectSelector, &out.ProjectSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectIAMMemberParameters.
func (in *ProjectIAMMemberParameters) DeepCopy() *ProjectIAMMemberParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectIAMMemberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectIAMMemberSpec) DeepCopyInto(out *ProjectIAMMemberSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectIAMMemberSpec.
func (in *ProjectIAMMemberSpec) DeepCopy() *ProjectIAMMemberSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectIAMMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectIAMMemberStatus) DeepCopyInto(out *ProjectIAMMemberStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectIAMMemberStatus.
func (in *ProjectIAMMemberStatus) DeepCopy() *ProjectIAMMemberStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectIAMMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountIAMMember) DeepCopyInto(out *ServiceAccountIAMMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountIAMMember.
func (in *ServiceAccountIAMMember) DeepCopy() *ServiceAccountIAMMember {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountIAMMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceAccountIAMMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountIAMMemberConditionInitParameters) DeepCopyInto(out *ServiceAccountIAMMemberConditionInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountIAMMemberConditionInitParameters.
func (in *ServiceAccountIAMMemberConditionInitParameters) DeepCopy() *ServiceAccountIAMMemberConditionInitParameters {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountIAMMemberConditionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountIAMMemberConditionObservation) DeepCopyInto(out *ServiceAccountIAMMemberConditionObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountIAMMemberConditionObservation.
func (in *ServiceAccountIAMMemberConditionObservation) DeepCopy() *ServiceAccountIAMMemberConditionObservation {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountIAMMemberConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountIAMMemberConditionParameters) DeepCopyInto(out *ServiceAccountIAMMemberConditionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountIAMMemberConditionParameters.
func (in *ServiceAccountIAMMemberConditionParameters) DeepCopy() *ServiceAccountIAMMemberConditionParameters {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountIAMMemberConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountIAMMemberInitParameters) DeepCopyInto(out *ServiceAccountIAMMemberInitParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(ServiceAccountIAMMemberConditionInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountID != nil {
		in, out := &in.ServiceAccountID, &out.ServiceAccountID
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountIDRef != nil {
		in, out := &in.ServiceAccountIDRef, &out.ServiceAccountIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccountIDSelector != nil {
		in, out := &in.ServiceAccountIDSelector, &out.ServiceAccountIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountIAMMemberInitParameters.
func (in *ServiceAccountIAMMemberInitParameters) DeepCopy() *ServiceAccountIAMMemberInitParameters {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountIAMMemberInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountIAMMemberList) DeepCopyInto(out *ServiceAccountIAMMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceAccountIAMMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountIAMMemberList.
func (in *ServiceAccountIAMMemberList) DeepCopy() *ServiceAccountIAMMemberList {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountIAMMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceAccountIAMMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountIAMMemberObservation) DeepCopyInto(out *ServiceAccountIAMMemberObservation) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(ServiceAccountIAMMemberConditionObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountID != nil {
		in, out := &in.ServiceAccountID, &out.ServiceAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountIAMMemberObservation.
func (in *ServiceAccountIAMMemberObservation) DeepCopy() *ServiceAccountIAMMemberObservation {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountIAMMemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountIAMMemberParameters) DeepCopyInto(out *ServiceAccountIAMMemberParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(ServiceAccountIAMMemberConditionParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountID != nil {
		in, out := &in.ServiceAccountID, &out.ServiceAccountID
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountIDRef != nil {
		in, out := &in.ServiceAccountIDRef, &out.ServiceAccountIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccountIDSelector != nil {
		in, out := &in.ServiceAccountIDSelector, &out.ServiceAccountIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountIAMMemberParameters.
func (in *ServiceAccountIAMMemberParameters) DeepCopy() *ServiceAccountIAMMemberParameters {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountIAMMemberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountIAMMemberSpec) DeepCopyInto(out *ServiceAccountIAMMemberSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountIAMMemberSpec.
func (in *ServiceAccountIAMMemberSpec) DeepCopy() *ServiceAccountIAMMemberSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountIAMMemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountIAMMemberStatus) DeepCopyInto(out *ServiceAccountIAMMemberStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountIAMMemberStatus.
func (in *ServiceAccountIAMMemberStatus) DeepCopy() *ServiceAccountIAMMemberStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountIAMMemberStatus)
	in.DeepCopyInto(out)
	return out
}
