/*
 * Copyright 2016 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package internal contains gRPC-internal code, to avoid polluting
// the godoc of the top-level grpc package.  It must not import any grpc
// symbols to avoid circular dependencies.
package internal

import (
	"context"
	"time"

	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/serviceconfig"
)

var (
	// WithHealthCheckFunc is set by dialoptions.go
	WithHealthCheckFunc interface{} // func (HealthChecker) DialOption
	// HealthCheckFunc is used to provide client-side LB channel health checking
	HealthCheckFunc HealthChecker
	// BalancerUnregister is exported by package balancer to unregister a balancer.
	BalancerUnregister func(name string)
	// KeepaliveMinPingTime is the minimum ping interval.  This must be 10s by
	// default, but tests may wish to set it lower for convenience.
	KeepaliveMinPingTime = 10 * time.Second
	// ParseServiceConfig parses a JSON representation of the service config.
	ParseServiceConfig interface{} // func(string) *serviceconfig.ParseResult
	// EqualServiceConfigForTesting is for testing service config generation and
	// parsing. Both a and b should be returned by ParseServiceConfig.
	// This function compares the config without rawJSON stripped, in case the
	// there's difference in white space.
	EqualServiceConfigForTesting func(a, b serviceconfig.Config) bool
	// GetCertificateProviderBuilder returns the registered builder for the
	// given name. This is set by package certprovider for use from xDS
	// bootstrap code while parsing certificate provider configs in the
	// bootstrap file.
	GetCertificateProviderBuilder interface{} // func(string) certprovider.Builder
	// GetXDSHandshakeInfoForTesting returns a pointer to the xds.HandshakeInfo
	// stored in the passed in attributes. This is set by
	// credentials/xds/xds.go.
	GetXDSHandshakeInfoForTesting interface{} // func (*attributes.Attributes) *xds.HandshakeInfo
	// GetServerCredentials returns the transport credentials configured on a
	// gRPC server. An xDS-enabled server needs to know what type of credentials
	// is configured on the underlying gRPC server. This is set by server.go.
	GetServerCredentials interface{} // func (*grpc.Server) credentials.TransportCredentials
	// DrainServerTransports initiates a graceful close of existing connections
	// on a gRPC server accepted on the provided listener address. An
	// xDS-enabled server invokes this method on a grpc.Server when a particular
	// listener moves to "not-serving" mode.
	DrainServerTransports interface{} // func(*grpc.Server, string)
	// AddExtraServerOptions adds an array of ServerOption that will be
	// effective globally for newly created servers. The priority will be: 1.
	// user-provided; 2. this method; 3. default values.
	AddExtraServerOptions interface{} // func(opt ...ServerOption)
	// ClearExtraServerOptions clears the array of extra ServerOption. This
	// method is useful in testing and benchmarking.
	ClearExtraServerOptions func()
	// AddExtraDialOptions adds an array of DialOption that will be effective
	// globally for newly created client channels. The priority will be: 1.
	// user-provided; 2. this method; 3. default values.
	AddExtraDialOptions interface{} // func(opt ...DialOption)
	// ClearExtraDialOptions clears the array of extra DialOption. This
	// method is useful in testing and benchmarking.
	ClearExtraDialOptions func()

	// NewXDSResolverWithConfigForTesting creates a new xds resolver builder using
	// the provided xds bootstrap config instead of the global configuration from
	// the supported environment variables.  The resolver.Builder is meant to be
	// used in conjunction with the grpc.WithResolvers DialOption.
	//
	// Testing Only
	//
	// This function should ONLY be used for testing and may not work with some
	// other features, including the CSDS service.
	NewXDSResolverWithConfigForTesting interface{} // func([]byte) (resolver.Builder, error)

	// RegisterRLSClusterSpecifierPluginForTesting registers the RLS Cluster
	// Specifier Plugin for testing purposes, regardless of the XDSRLS environment
	// variable.
	//
	// TODO: Remove this function once the RLS env var is removed.
	RegisterRLSClusterSpecifierPluginForTesting func()

	// UnregisterRLSClusterSpecifierPluginForTesting unregisters the RLS Cluster
	// Specifier Plugin for testing purposes. This is needed because there is no way
	// to unregister the RLS Cluster Specifier Plugin after registering it solely
	// for testing purposes using RegisterRLSClusterSpecifierPluginForTesting().
	//
	// TODO: Remove this function once the RLS env var is removed.
	UnregisterRLSClusterSpecifierPluginForTesting func()

	// RegisterRBACHTTPFilterForTesting registers the RBAC HTTP Filter for testing
	// purposes, regardless of the RBAC environment variable.
	//
	// TODO: Remove this function once the RBAC env var is removed.
	RegisterRBACHTTPFilterForTesting func()

	// UnregisterRBACHTTPFilterForTesting unregisters the RBAC HTTP Filter for
	// testing purposes. This is needed because there is no way to unregister the
	// HTTP Filter after registering it solely for testing purposes using
	// RegisterRBACHTTPFilterForTesting().
	//
	// TODO: Remove this function once the RBAC env var is removed.
	UnregisterRBACHTTPFilterForTesting func()
)

// HealthChecker defines the signature of the client-side LB channel health checking function.
//
// The implementation is expected to create a health checking RPC stream by
// calling newStream(), watch for the health status of serviceName, and report
// it's health back by calling setConnectivityState().
//
// The health checking protocol is defined at:
// https://github.com/grpc/grpc/blob/master/doc/health-checking.md
type HealthChecker func(ctx context.Context, newStream func(string) (interface{}, error), setConnectivityState func(connectivity.State, error), serviceName string) error

const (
	// CredsBundleModeFallback switches GoogleDefaultCreds to fallback mode.
	CredsBundleModeFallback = "fallback"
	// CredsBundleModeBalancer switches GoogleDefaultCreds to grpclb balancer
	// mode.
	CredsBundleModeBalancer = "balancer"
	// CredsBundleModeBackendFromBalancer switches GoogleDefaultCreds to mode
	// that supports backend returned by grpclb balancer.
	CredsBundleModeBackendFromBalancer = "backend-from-balancer"
)

// RLSLoadBalancingPolicyName is the name of the RLS LB policy.
//
// It currently has an experimental suffix which would be removed once
// end-to-end testing of the policy is completed.
const RLSLoadBalancingPolicyName = "rls_experimental"
