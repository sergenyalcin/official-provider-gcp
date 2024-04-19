//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationInitParameters) DeepCopyInto(out *AuthorizationInitParameters) {
	*out = *in
	if in.AdminGroups != nil {
		in, out := &in.AdminGroups, &out.AdminGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AdminUsers != nil {
		in, out := &in.AdminUsers, &out.AdminUsers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationInitParameters.
func (in *AuthorizationInitParameters) DeepCopy() *AuthorizationInitParameters {
	if in == nil {
		return nil
	}
	out := new(AuthorizationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationObservation) DeepCopyInto(out *AuthorizationObservation) {
	*out = *in
	if in.AdminGroups != nil {
		in, out := &in.AdminGroups, &out.AdminGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AdminUsers != nil {
		in, out := &in.AdminUsers, &out.AdminUsers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationObservation.
func (in *AuthorizationObservation) DeepCopy() *AuthorizationObservation {
	if in == nil {
		return nil
	}
	out := new(AuthorizationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationParameters) DeepCopyInto(out *AuthorizationParameters) {
	*out = *in
	if in.AdminGroups != nil {
		in, out := &in.AdminGroups, &out.AdminGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AdminUsers != nil {
		in, out := &in.AdminUsers, &out.AdminUsers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationParameters.
func (in *AuthorizationParameters) DeepCopy() *AuthorizationParameters {
	if in == nil {
		return nil
	}
	out := new(AuthorizationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BinaryAuthorizationInitParameters) DeepCopyInto(out *BinaryAuthorizationInitParameters) {
	*out = *in
	if in.EvaluationMode != nil {
		in, out := &in.EvaluationMode, &out.EvaluationMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BinaryAuthorizationInitParameters.
func (in *BinaryAuthorizationInitParameters) DeepCopy() *BinaryAuthorizationInitParameters {
	if in == nil {
		return nil
	}
	out := new(BinaryAuthorizationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BinaryAuthorizationObservation) DeepCopyInto(out *BinaryAuthorizationObservation) {
	*out = *in
	if in.EvaluationMode != nil {
		in, out := &in.EvaluationMode, &out.EvaluationMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BinaryAuthorizationObservation.
func (in *BinaryAuthorizationObservation) DeepCopy() *BinaryAuthorizationObservation {
	if in == nil {
		return nil
	}
	out := new(BinaryAuthorizationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BinaryAuthorizationParameters) DeepCopyInto(out *BinaryAuthorizationParameters) {
	*out = *in
	if in.EvaluationMode != nil {
		in, out := &in.EvaluationMode, &out.EvaluationMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BinaryAuthorizationParameters.
func (in *BinaryAuthorizationParameters) DeepCopy() *BinaryAuthorizationParameters {
	if in == nil {
		return nil
	}
	out := new(BinaryAuthorizationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInitParameters) DeepCopyInto(out *ClusterInitParameters) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = new(AuthorizationInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.BinaryAuthorization != nil {
		in, out := &in.BinaryAuthorization, &out.BinaryAuthorization
		*out = new(BinaryAuthorizationInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.DeletionPolicy != nil {
		in, out := &in.DeletionPolicy, &out.DeletionPolicy
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Distribution != nil {
		in, out := &in.Distribution, &out.Distribution
		*out = new(string)
		**out = **in
	}
	if in.Fleet != nil {
		in, out := &in.Fleet, &out.Fleet
		*out = new(FleetInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.LoggingConfig != nil {
		in, out := &in.LoggingConfig, &out.LoggingConfig
		*out = new(LoggingConfigInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.MonitoringConfig != nil {
		in, out := &in.MonitoringConfig, &out.MonitoringConfig
		*out = new(MonitoringConfigInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.OidcConfig != nil {
		in, out := &in.OidcConfig, &out.OidcConfig
		*out = new(OidcConfigInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.PlatformVersion != nil {
		in, out := &in.PlatformVersion, &out.PlatformVersion
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ProxyConfig != nil {
		in, out := &in.ProxyConfig, &out.ProxyConfig
		*out = new(ProxyConfigInitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInitParameters.
func (in *ClusterInitParameters) DeepCopy() *ClusterInitParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObservation) DeepCopyInto(out *ClusterObservation) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = new(AuthorizationObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.BinaryAuthorization != nil {
		in, out := &in.BinaryAuthorization, &out.BinaryAuthorization
		*out = new(BinaryAuthorizationObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterRegion != nil {
		in, out := &in.ClusterRegion, &out.ClusterRegion
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.DeletionPolicy != nil {
		in, out := &in.DeletionPolicy, &out.DeletionPolicy
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Distribution != nil {
		in, out := &in.Distribution, &out.Distribution
		*out = new(string)
		**out = **in
	}
	if in.EffectiveAnnotations != nil {
		in, out := &in.EffectiveAnnotations, &out.EffectiveAnnotations
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]ErrorsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Fleet != nil {
		in, out := &in.Fleet, &out.Fleet
		*out = new(FleetObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.LoggingConfig != nil {
		in, out := &in.LoggingConfig, &out.LoggingConfig
		*out = new(LoggingConfigObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.MonitoringConfig != nil {
		in, out := &in.MonitoringConfig, &out.MonitoringConfig
		*out = new(MonitoringConfigObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.OidcConfig != nil {
		in, out := &in.OidcConfig, &out.OidcConfig
		*out = new(OidcConfigObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.PlatformVersion != nil {
		in, out := &in.PlatformVersion, &out.PlatformVersion
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ProxyConfig != nil {
		in, out := &in.ProxyConfig, &out.ProxyConfig
		*out = new(ProxyConfigObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Reconciling != nil {
		in, out := &in.Reconciling, &out.Reconciling
		*out = new(bool)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.WorkloadIdentityConfig != nil {
		in, out := &in.WorkloadIdentityConfig, &out.WorkloadIdentityConfig
		*out = make([]WorkloadIdentityConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObservation.
func (in *ClusterObservation) DeepCopy() *ClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameters) DeepCopyInto(out *ClusterParameters) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = new(AuthorizationParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.BinaryAuthorization != nil {
		in, out := &in.BinaryAuthorization, &out.BinaryAuthorization
		*out = new(BinaryAuthorizationParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.DeletionPolicy != nil {
		in, out := &in.DeletionPolicy, &out.DeletionPolicy
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Distribution != nil {
		in, out := &in.Distribution, &out.Distribution
		*out = new(string)
		**out = **in
	}
	if in.Fleet != nil {
		in, out := &in.Fleet, &out.Fleet
		*out = new(FleetParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.LoggingConfig != nil {
		in, out := &in.LoggingConfig, &out.LoggingConfig
		*out = new(LoggingConfigParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.MonitoringConfig != nil {
		in, out := &in.MonitoringConfig, &out.MonitoringConfig
		*out = new(MonitoringConfigParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.OidcConfig != nil {
		in, out := &in.OidcConfig, &out.OidcConfig
		*out = new(OidcConfigParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.PlatformVersion != nil {
		in, out := &in.PlatformVersion, &out.PlatformVersion
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ProxyConfig != nil {
		in, out := &in.ProxyConfig, &out.ProxyConfig
		*out = new(ProxyConfigParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameters.
func (in *ClusterParameters) DeepCopy() *ClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentConfigInitParameters) DeepCopyInto(out *ComponentConfigInitParameters) {
	*out = *in
	if in.EnableComponents != nil {
		in, out := &in.EnableComponents, &out.EnableComponents
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentConfigInitParameters.
func (in *ComponentConfigInitParameters) DeepCopy() *ComponentConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(ComponentConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentConfigObservation) DeepCopyInto(out *ComponentConfigObservation) {
	*out = *in
	if in.EnableComponents != nil {
		in, out := &in.EnableComponents, &out.EnableComponents
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentConfigObservation.
func (in *ComponentConfigObservation) DeepCopy() *ComponentConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ComponentConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentConfigParameters) DeepCopyInto(out *ComponentConfigParameters) {
	*out = *in
	if in.EnableComponents != nil {
		in, out := &in.EnableComponents, &out.EnableComponents
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentConfigParameters.
func (in *ComponentConfigParameters) DeepCopy() *ComponentConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ComponentConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorsInitParameters) DeepCopyInto(out *ErrorsInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorsInitParameters.
func (in *ErrorsInitParameters) DeepCopy() *ErrorsInitParameters {
	if in == nil {
		return nil
	}
	out := new(ErrorsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorsObservation) DeepCopyInto(out *ErrorsObservation) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorsObservation.
func (in *ErrorsObservation) DeepCopy() *ErrorsObservation {
	if in == nil {
		return nil
	}
	out := new(ErrorsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorsParameters) DeepCopyInto(out *ErrorsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorsParameters.
func (in *ErrorsParameters) DeepCopy() *ErrorsParameters {
	if in == nil {
		return nil
	}
	out := new(ErrorsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetInitParameters) DeepCopyInto(out *FleetInitParameters) {
	*out = *in
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetInitParameters.
func (in *FleetInitParameters) DeepCopy() *FleetInitParameters {
	if in == nil {
		return nil
	}
	out := new(FleetInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetObservation) DeepCopyInto(out *FleetObservation) {
	*out = *in
	if in.Membership != nil {
		in, out := &in.Membership, &out.Membership
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetObservation.
func (in *FleetObservation) DeepCopy() *FleetObservation {
	if in == nil {
		return nil
	}
	out := new(FleetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetParameters) DeepCopyInto(out *FleetParameters) {
	*out = *in
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetParameters.
func (in *FleetParameters) DeepCopy() *FleetParameters {
	if in == nil {
		return nil
	}
	out := new(FleetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesSecretInitParameters) DeepCopyInto(out *KubernetesSecretInitParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesSecretInitParameters.
func (in *KubernetesSecretInitParameters) DeepCopy() *KubernetesSecretInitParameters {
	if in == nil {
		return nil
	}
	out := new(KubernetesSecretInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesSecretObservation) DeepCopyInto(out *KubernetesSecretObservation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesSecretObservation.
func (in *KubernetesSecretObservation) DeepCopy() *KubernetesSecretObservation {
	if in == nil {
		return nil
	}
	out := new(KubernetesSecretObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesSecretParameters) DeepCopyInto(out *KubernetesSecretParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesSecretParameters.
func (in *KubernetesSecretParameters) DeepCopy() *KubernetesSecretParameters {
	if in == nil {
		return nil
	}
	out := new(KubernetesSecretParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfigInitParameters) DeepCopyInto(out *LoggingConfigInitParameters) {
	*out = *in
	if in.ComponentConfig != nil {
		in, out := &in.ComponentConfig, &out.ComponentConfig
		*out = new(ComponentConfigInitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfigInitParameters.
func (in *LoggingConfigInitParameters) DeepCopy() *LoggingConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(LoggingConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfigObservation) DeepCopyInto(out *LoggingConfigObservation) {
	*out = *in
	if in.ComponentConfig != nil {
		in, out := &in.ComponentConfig, &out.ComponentConfig
		*out = new(ComponentConfigObservation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfigObservation.
func (in *LoggingConfigObservation) DeepCopy() *LoggingConfigObservation {
	if in == nil {
		return nil
	}
	out := new(LoggingConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfigParameters) DeepCopyInto(out *LoggingConfigParameters) {
	*out = *in
	if in.ComponentConfig != nil {
		in, out := &in.ComponentConfig, &out.ComponentConfig
		*out = new(ComponentConfigParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfigParameters.
func (in *LoggingConfigParameters) DeepCopy() *LoggingConfigParameters {
	if in == nil {
		return nil
	}
	out := new(LoggingConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPrometheusConfigInitParameters) DeepCopyInto(out *ManagedPrometheusConfigInitParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPrometheusConfigInitParameters.
func (in *ManagedPrometheusConfigInitParameters) DeepCopy() *ManagedPrometheusConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(ManagedPrometheusConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPrometheusConfigObservation) DeepCopyInto(out *ManagedPrometheusConfigObservation) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPrometheusConfigObservation.
func (in *ManagedPrometheusConfigObservation) DeepCopy() *ManagedPrometheusConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ManagedPrometheusConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedPrometheusConfigParameters) DeepCopyInto(out *ManagedPrometheusConfigParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedPrometheusConfigParameters.
func (in *ManagedPrometheusConfigParameters) DeepCopy() *ManagedPrometheusConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ManagedPrometheusConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConfigInitParameters) DeepCopyInto(out *MonitoringConfigInitParameters) {
	*out = *in
	if in.ManagedPrometheusConfig != nil {
		in, out := &in.ManagedPrometheusConfig, &out.ManagedPrometheusConfig
		*out = new(ManagedPrometheusConfigInitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConfigInitParameters.
func (in *MonitoringConfigInitParameters) DeepCopy() *MonitoringConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(MonitoringConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConfigObservation) DeepCopyInto(out *MonitoringConfigObservation) {
	*out = *in
	if in.ManagedPrometheusConfig != nil {
		in, out := &in.ManagedPrometheusConfig, &out.ManagedPrometheusConfig
		*out = new(ManagedPrometheusConfigObservation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConfigObservation.
func (in *MonitoringConfigObservation) DeepCopy() *MonitoringConfigObservation {
	if in == nil {
		return nil
	}
	out := new(MonitoringConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringConfigParameters) DeepCopyInto(out *MonitoringConfigParameters) {
	*out = *in
	if in.ManagedPrometheusConfig != nil {
		in, out := &in.ManagedPrometheusConfig, &out.ManagedPrometheusConfig
		*out = new(ManagedPrometheusConfigParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringConfigParameters.
func (in *MonitoringConfigParameters) DeepCopy() *MonitoringConfigParameters {
	if in == nil {
		return nil
	}
	out := new(MonitoringConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OidcConfigInitParameters) DeepCopyInto(out *OidcConfigInitParameters) {
	*out = *in
	if in.IssuerURL != nil {
		in, out := &in.IssuerURL, &out.IssuerURL
		*out = new(string)
		**out = **in
	}
	if in.Jwks != nil {
		in, out := &in.Jwks, &out.Jwks
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OidcConfigInitParameters.
func (in *OidcConfigInitParameters) DeepCopy() *OidcConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(OidcConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OidcConfigObservation) DeepCopyInto(out *OidcConfigObservation) {
	*out = *in
	if in.IssuerURL != nil {
		in, out := &in.IssuerURL, &out.IssuerURL
		*out = new(string)
		**out = **in
	}
	if in.Jwks != nil {
		in, out := &in.Jwks, &out.Jwks
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OidcConfigObservation.
func (in *OidcConfigObservation) DeepCopy() *OidcConfigObservation {
	if in == nil {
		return nil
	}
	out := new(OidcConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OidcConfigParameters) DeepCopyInto(out *OidcConfigParameters) {
	*out = *in
	if in.IssuerURL != nil {
		in, out := &in.IssuerURL, &out.IssuerURL
		*out = new(string)
		**out = **in
	}
	if in.Jwks != nil {
		in, out := &in.Jwks, &out.Jwks
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OidcConfigParameters.
func (in *OidcConfigParameters) DeepCopy() *OidcConfigParameters {
	if in == nil {
		return nil
	}
	out := new(OidcConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyConfigInitParameters) DeepCopyInto(out *ProxyConfigInitParameters) {
	*out = *in
	if in.KubernetesSecret != nil {
		in, out := &in.KubernetesSecret, &out.KubernetesSecret
		*out = new(KubernetesSecretInitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyConfigInitParameters.
func (in *ProxyConfigInitParameters) DeepCopy() *ProxyConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(ProxyConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyConfigObservation) DeepCopyInto(out *ProxyConfigObservation) {
	*out = *in
	if in.KubernetesSecret != nil {
		in, out := &in.KubernetesSecret, &out.KubernetesSecret
		*out = new(KubernetesSecretObservation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyConfigObservation.
func (in *ProxyConfigObservation) DeepCopy() *ProxyConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ProxyConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyConfigParameters) DeepCopyInto(out *ProxyConfigParameters) {
	*out = *in
	if in.KubernetesSecret != nil {
		in, out := &in.KubernetesSecret, &out.KubernetesSecret
		*out = new(KubernetesSecretParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyConfigParameters.
func (in *ProxyConfigParameters) DeepCopy() *ProxyConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ProxyConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityConfigInitParameters) DeepCopyInto(out *WorkloadIdentityConfigInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityConfigInitParameters.
func (in *WorkloadIdentityConfigInitParameters) DeepCopy() *WorkloadIdentityConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityConfigObservation) DeepCopyInto(out *WorkloadIdentityConfigObservation) {
	*out = *in
	if in.IdentityProvider != nil {
		in, out := &in.IdentityProvider, &out.IdentityProvider
		*out = new(string)
		**out = **in
	}
	if in.IssuerURI != nil {
		in, out := &in.IssuerURI, &out.IssuerURI
		*out = new(string)
		**out = **in
	}
	if in.WorkloadPool != nil {
		in, out := &in.WorkloadPool, &out.WorkloadPool
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityConfigObservation.
func (in *WorkloadIdentityConfigObservation) DeepCopy() *WorkloadIdentityConfigObservation {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadIdentityConfigParameters) DeepCopyInto(out *WorkloadIdentityConfigParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadIdentityConfigParameters.
func (in *WorkloadIdentityConfigParameters) DeepCopy() *WorkloadIdentityConfigParameters {
	if in == nil {
		return nil
	}
	out := new(WorkloadIdentityConfigParameters)
	in.DeepCopyInto(out)
	return out
}
