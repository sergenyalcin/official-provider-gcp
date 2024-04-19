// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RegionHealthCheckGRPCHealthCheckInitParameters struct {

	// The gRPC service name for the health check.
	// The value of grpcServiceName has the following meanings by convention:
	GRPCServiceName *string `json:"grpcServiceName,omitempty" tf:"grpc_service_name,omitempty"`

	// The port number for the health check request.
	// Must be specified if portName and portSpecification are not set
	// or if port_specification is USE_FIXED_PORT. Valid values are 1 through 65535.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`
}

type RegionHealthCheckGRPCHealthCheckObservation struct {

	// The gRPC service name for the health check.
	// The value of grpcServiceName has the following meanings by convention:
	GRPCServiceName *string `json:"grpcServiceName,omitempty" tf:"grpc_service_name,omitempty"`

	// The port number for the health check request.
	// Must be specified if portName and portSpecification are not set
	// or if port_specification is USE_FIXED_PORT. Valid values are 1 through 65535.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`
}

type RegionHealthCheckGRPCHealthCheckParameters struct {

	// The gRPC service name for the health check.
	// The value of grpcServiceName has the following meanings by convention:
	// +kubebuilder:validation:Optional
	GRPCServiceName *string `json:"grpcServiceName,omitempty" tf:"grpc_service_name,omitempty"`

	// The port number for the health check request.
	// Must be specified if portName and portSpecification are not set
	// or if port_specification is USE_FIXED_PORT. Valid values are 1 through 65535.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`
}

type RegionHealthCheckHTTPHealthCheckInitParameters struct {

	// The value of the host header in the HTTP health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP health check request.
	// The default value is 80.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTP health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckHTTPHealthCheckObservation struct {

	// The value of the host header in the HTTP health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP health check request.
	// The default value is 80.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTP health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckHTTPHealthCheckParameters struct {

	// The value of the host header in the HTTP health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP health check request.
	// The default value is 80.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTP health check request.
	// The default value is /.
	// +kubebuilder:validation:Optional
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckHTTPSHealthCheckInitParameters struct {

	// The value of the host header in the HTTPS health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTPS health check request.
	// The default value is 443.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTPS health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckHTTPSHealthCheckObservation struct {

	// The value of the host header in the HTTPS health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTPS health check request.
	// The default value is 443.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTPS health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckHTTPSHealthCheckParameters struct {

	// The value of the host header in the HTTPS health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTPS health check request.
	// The default value is 443.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTPS health check request.
	// The default value is /.
	// +kubebuilder:validation:Optional
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckHttp2HealthCheckInitParameters struct {

	// The value of the host header in the HTTP2 health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP2 health check request.
	// The default value is 443.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTP2 health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckHttp2HealthCheckObservation struct {

	// The value of the host header in the HTTP2 health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP2 health check request.
	// The default value is 443.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTP2 health check request.
	// The default value is /.
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckHttp2HealthCheckParameters struct {

	// The value of the host header in the HTTP2 health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The TCP port number for the HTTP2 health check request.
	// The default value is 443.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The request path of the HTTP2 health check request.
	// The default value is /.
	// +kubebuilder:validation:Optional
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckInitParameters struct {

	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec *float64 `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A nested object resource
	// Structure is documented below.
	GRPCHealthCheck []RegionHealthCheckGRPCHealthCheckInitParameters `json:"grpcHealthCheck,omitempty" tf:"grpc_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	HTTPHealthCheck []RegionHealthCheckHTTPHealthCheckInitParameters `json:"httpHealthCheck,omitempty" tf:"http_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	HTTPSHealthCheck []RegionHealthCheckHTTPSHealthCheckInitParameters `json:"httpsHealthCheck,omitempty" tf:"https_health_check,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// A nested object resource
	// Structure is documented below.
	Http2HealthCheck []RegionHealthCheckHttp2HealthCheckInitParameters `json:"http2HealthCheck,omitempty" tf:"http2_health_check,omitempty"`

	// Configure logging on this health check.
	// Structure is documented below.
	LogConfig []RegionHealthCheckLogConfigInitParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A nested object resource
	// Structure is documented below.
	SSLHealthCheck []RegionHealthCheckSSLHealthCheckInitParameters `json:"sslHealthCheck,omitempty" tf:"ssl_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	TCPHealthCheck []RegionHealthCheckTCPHealthCheckInitParameters `json:"tcpHealthCheck,omitempty" tf:"tcp_health_check,omitempty"`

	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec *float64 `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`

	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type RegionHealthCheckLogConfigInitParameters struct {

	// Indicates whether or not to export logs. This is false by default,
	// which means no health check logging will be done.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`
}

type RegionHealthCheckLogConfigObservation struct {

	// Indicates whether or not to export logs. This is false by default,
	// which means no health check logging will be done.
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`
}

type RegionHealthCheckLogConfigParameters struct {

	// Indicates whether or not to export logs. This is false by default,
	// which means no health check logging will be done.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`
}

type RegionHealthCheckObservation struct {

	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	CheckIntervalSec *float64 `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A nested object resource
	// Structure is documented below.
	GRPCHealthCheck []RegionHealthCheckGRPCHealthCheckObservation `json:"grpcHealthCheck,omitempty" tf:"grpc_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	HTTPHealthCheck []RegionHealthCheckHTTPHealthCheckObservation `json:"httpHealthCheck,omitempty" tf:"http_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	HTTPSHealthCheck []RegionHealthCheckHTTPSHealthCheckObservation `json:"httpsHealthCheck,omitempty" tf:"https_health_check,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// A nested object resource
	// Structure is documented below.
	Http2HealthCheck []RegionHealthCheckHttp2HealthCheckObservation `json:"http2HealthCheck,omitempty" tf:"http2_health_check,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/healthChecks/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configure logging on this health check.
	// Structure is documented below.
	LogConfig []RegionHealthCheckLogConfigObservation `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Region in which the created health check should reside.
	// If it is not provided, the provider region is used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A nested object resource
	// Structure is documented below.
	SSLHealthCheck []RegionHealthCheckSSLHealthCheckObservation `json:"sslHealthCheck,omitempty" tf:"ssl_health_check,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// A nested object resource
	// Structure is documented below.
	TCPHealthCheck []RegionHealthCheckTCPHealthCheckObservation `json:"tcpHealthCheck,omitempty" tf:"tcp_health_check,omitempty"`

	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	TimeoutSec *float64 `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`

	// The type of the health check. One of HTTP, HTTP2, HTTPS, TCP, or SSL.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type RegionHealthCheckParameters struct {

	// How often (in seconds) to send a health check. The default value is 5
	// seconds.
	// +kubebuilder:validation:Optional
	CheckIntervalSec *float64 `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GRPCHealthCheck []RegionHealthCheckGRPCHealthCheckParameters `json:"grpcHealthCheck,omitempty" tf:"grpc_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	HTTPHealthCheck []RegionHealthCheckHTTPHealthCheckParameters `json:"httpHealthCheck,omitempty" tf:"http_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	HTTPSHealthCheck []RegionHealthCheckHTTPSHealthCheckParameters `json:"httpsHealthCheck,omitempty" tf:"https_health_check,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	// +kubebuilder:validation:Optional
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Http2HealthCheck []RegionHealthCheckHttp2HealthCheckParameters `json:"http2HealthCheck,omitempty" tf:"http2_health_check,omitempty"`

	// Configure logging on this health check.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	LogConfig []RegionHealthCheckLogConfigParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Region in which the created health check should reside.
	// If it is not provided, the provider region is used.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SSLHealthCheck []RegionHealthCheckSSLHealthCheckParameters `json:"sslHealthCheck,omitempty" tf:"ssl_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TCPHealthCheck []RegionHealthCheckTCPHealthCheckParameters `json:"tcpHealthCheck,omitempty" tf:"tcp_health_check,omitempty"`

	// How long (in seconds) to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	// +kubebuilder:validation:Optional
	TimeoutSec *float64 `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`

	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	// +kubebuilder:validation:Optional
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type RegionHealthCheckSSLHealthCheckInitParameters struct {

	// The TCP port number for the SSL health check request.
	// The default value is 443.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The application data to send once the SSL connection has been
	// established (default value is empty). If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	Request *string `json:"request,omitempty" tf:"request,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckSSLHealthCheckObservation struct {

	// The TCP port number for the SSL health check request.
	// The default value is 443.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The application data to send once the SSL connection has been
	// established (default value is empty). If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	Request *string `json:"request,omitempty" tf:"request,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckSSLHealthCheckParameters struct {

	// The TCP port number for the SSL health check request.
	// The default value is 443.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The application data to send once the SSL connection has been
	// established (default value is empty). If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	// +kubebuilder:validation:Optional
	Request *string `json:"request,omitempty" tf:"request,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckTCPHealthCheckInitParameters struct {

	// The TCP port number for the TCP health check request.
	// The default value is 80.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The application data to send once the TCP connection has been
	// established (default value is empty). If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	Request *string `json:"request,omitempty" tf:"request,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckTCPHealthCheckObservation struct {

	// The TCP port number for the TCP health check request.
	// The default value is 80.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The application data to send once the TCP connection has been
	// established (default value is empty). If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	Request *string `json:"request,omitempty" tf:"request,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckTCPHealthCheckParameters struct {

	// The TCP port number for the TCP health check request.
	// The default value is 80.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend.
	// Default value is NONE.
	// Possible values are: NONE, PROXY_V1.
	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// The application data to send once the TCP connection has been
	// established (default value is empty). If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	// +kubebuilder:validation:Optional
	Request *string `json:"request,omitempty" tf:"request,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

// RegionHealthCheckSpec defines the desired state of RegionHealthCheck
type RegionHealthCheckSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionHealthCheckParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RegionHealthCheckInitParameters `json:"initProvider,omitempty"`
}

// RegionHealthCheckStatus defines the observed state of RegionHealthCheck.
type RegionHealthCheckStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionHealthCheckObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RegionHealthCheck is the Schema for the RegionHealthChecks API. Health Checks determine whether instances are responsive and able to do work.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RegionHealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegionHealthCheckSpec   `json:"spec"`
	Status            RegionHealthCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionHealthCheckList contains a list of RegionHealthChecks
type RegionHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionHealthCheck `json:"items"`
}

// Repository type metadata.
var (
	RegionHealthCheck_Kind             = "RegionHealthCheck"
	RegionHealthCheck_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegionHealthCheck_Kind}.String()
	RegionHealthCheck_KindAPIVersion   = RegionHealthCheck_Kind + "." + CRDGroupVersion.String()
	RegionHealthCheck_GroupVersionKind = CRDGroupVersion.WithKind(RegionHealthCheck_Kind)
)

func init() {
	SchemeBuilder.Register(&RegionHealthCheck{}, &RegionHealthCheckList{})
}
