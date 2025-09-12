package patterns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/patterns/internal"
)

// Deploys a production-grade Keycloak service.
//
// This service deploys a containerized version of Keycloak using AWS Fargate to host and scale the application. The
// backend database is supported via Amazon Relational Database Service (RDS) and the application is fronted by a
// Network Load Balancer.
//
// The database will auto-scale based on CDK defaults or a consumer-specified scaling policy. The containers will not
// automatically scale unless a consumer-specified policy is applied.
//
// It is recommended to set the CDK feature flag `@aws-cdk/aws-rds:databaseProxyUniqueResourceName` in
// `cdk.json` to true. If not done, the database proxy name may conflict with other proxies in your account and
// will prevent you from being able to deploy more than one KeycloakService.
//
// At a minimum this pattern requires the consumer to build and provide their own Keycloak container image for
// deployment as well provide hostname configuration details. The Keycloak container image version MUST match the
// version specified for use here and must include the Amazon Aurora JDBC driver pre-installed. A minimum viable
// Dockerfile for that container image looks like:
//
// ```Dockerfile
// ARG VERSION=26.3.2
//
// FROM quay.io/keycloak/keycloak:${VERSION} AS builder
//
// # Optimizations (not necessary but speed up the container startup)
// ENV KC_DB=postgres
// ENV KC_DB_DRIVER=software.amazon.jdbc.Driver
//
// WORKDIR /opt/keycloak
//
// # TLS Configuration
// COPY --chmod=0666 certs/server.keystore conf/
//
// # Database Provider
// ADD --chmod=0666 https://github.com/aws/aws-advanced-jdbc-wrapper/releases/download/2.6.2/aws-advanced-jdbc-wrapper-2.6.2.jar providers/aws-advanced-jdbc-wrapper.jar
//
// RUN /opt/keycloak/bin/kc.sh build
//
// FROM quay.io/keycloak/keycloak:${VERSION}
// COPY --from=builder /opt/keycloak /opt/keycloak
//
// ENTRYPOINT [ "/opt/keycloak/bin/kc.sh" ]
// CMD [ "start" ]
// ```
// ---
// By default, the Keycloak service is deployed internally in isolated and/or private subnets but can be exposed by
// providing the fabric configuration option to expose the service with an internet-facing load balancer.
//
// Example:
//   import { join } from 'node:path';
//   import { KeycloakService } from '@cdklabs/cdk-proserve-lib/patterns';
//   import { App, Environment, RemovalPolicy, Stack } from 'aws-cdk-lib';
//   import { IpAddresses, SubnetType, Vpc } from 'aws-cdk-lib/aws-ec2';
//   import { AssetImage } from 'aws-cdk-lib/aws-ecs';
//   import { Platform } from 'aws-cdk-lib/aws-ecr-assets';
//
//   const dnsZoneName = 'example.com';
//   const network = Vpc.fromLookup(this, 'Network', {
//       vpcId: 'vpc-xxxx'
//   });
//
//   new KeycloakService(this, 'Keycloak', {
//       keycloak: {
//           image: AssetImage.fromAsset(join(__dirname, '..', 'src', 'keycloak'), {
//               platform: Platform.LINUX_AMD64
//           }),
//           configuration: {
//               hostnames: {
//                   default: `auth.${dnsZoneName}`,
//                   admin: `admin.auth.${dnsZoneName}`
//               },
//               loggingLevel: 'info'
//           },
//           version: KeycloakService.EngineVersion.V26_3_2
//       },
//       overrides: {
//           cluster: {
//               scaling: {
//                   minimum: 1,
//                   maximum: 2
//               }
//           }
//           fabric: {
//               dnsZoneName: dnsZoneName,
//               internetFacing: true
//           }
//       },
//       vpc: network
//   });
//
// Experimental.
type KeycloakService interface {
	constructs.Construct
	// Credentials for bootstrapping a local admin user in Keycloak.
	// Experimental.
	AdminUser() awssecretsmanager.ISecret
	// Endpoint for the Keycloak service.
	// Experimental.
	Endpoint() interface{}
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KeycloakService
type jsiiProxy_KeycloakService struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_KeycloakService) AdminUser() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"adminUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeycloakService) Endpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeycloakService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Create a new Keycloak service.
// Experimental.
func NewKeycloakService(scope constructs.Construct, id *string, props *KeycloakServiceProps) KeycloakService {
	_init_.Initialize()

	if err := validateNewKeycloakServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeycloakService{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new Keycloak service.
// Experimental.
func NewKeycloakService_Override(k KeycloakService, scope constructs.Construct, id *string, props *KeycloakServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService",
		[]interface{}{scope, id, props},
		k,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func KeycloakService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKeycloakService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-proserve-lib.patterns.KeycloakService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeycloakService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

