package components

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-proserve-lib-go/cdklabscdkproservelib/components/internal"
)

// Creates an AWS Web Application Firewall (WAF) that can be associated with resources such as an Application Load Balancer.
//
// It allows configuring AWS managed rule groups, logging, and visibility
// settings. The construct simplifies the creation of a WAF by providing
// available AWS managed rule groups that can be utilized.
//
// Currently, the only resource that is supported to associate the WAF with is
// an ALB.
//
// Example:
//   import { ApplicationLoadBalancer } from 'aws-cdk-lib/aws-elasticloadbalancingv2';
//   import { WebApplicationFirewall } from '@cdklabs/cdk-proserve-lib/constructs';
//
//   const alb = new ApplicationLoadBalancer(this, 'Alb', { vpc });
//   const waf = new WebApplicationFirewall(this, 'WAF', {
//     awsManagedRuleGroups: [
//       WebApplicationFirewall.AwsManagedRuleGroup.COMMON_RULE_SET,
//       WebApplicationFirewall.AwsManagedRuleGroup.LINUX_RULE_SET
//     ]
//   });
//   waf.associate(alb);  // Associate the WAF with the ALB
//
// Experimental.
type WebApplicationFirewall interface {
	constructs.Construct
	// Optional CloudWatch log group for WAF logging.
	//
	// This is available if you
	// have configured `logging` on the construct.
	// Experimental.
	LogGroup() awslogs.LogGroup
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The WAF Web ACL (Access Control List) resource.
	// Experimental.
	WebAcl() awswafv2.CfnWebACL
	// Associates the Web Application Firewall to an applicable resource.
	// Experimental.
	Associate(resource awselasticloadbalancingv2.ApplicationLoadBalancer)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WebApplicationFirewall
type jsiiProxy_WebApplicationFirewall struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_WebApplicationFirewall) LogGroup() awslogs.LogGroup {
	var returns awslogs.LogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewall) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewall) WebAcl() awswafv2.CfnWebACL {
	var returns awswafv2.CfnWebACL
	_jsii_.Get(
		j,
		"webAcl",
		&returns,
	)
	return returns
}


// Creates an AWS Web Application Firewall (WAF) that can be associated with resources such as an Application Load Balancer.
// Experimental.
func NewWebApplicationFirewall(scope constructs.Construct, id *string, props *WebApplicationFirewallProps) WebApplicationFirewall {
	_init_.Initialize()

	if err := validateNewWebApplicationFirewallParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebApplicationFirewall{}

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.WebApplicationFirewall",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an AWS Web Application Firewall (WAF) that can be associated with resources such as an Application Load Balancer.
// Experimental.
func NewWebApplicationFirewall_Override(w WebApplicationFirewall, scope constructs.Construct, id *string, props *WebApplicationFirewallProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-proserve-lib.constructs.WebApplicationFirewall",
		[]interface{}{scope, id, props},
		w,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func WebApplicationFirewall_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWebApplicationFirewall_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-proserve-lib.constructs.WebApplicationFirewall",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewall) Associate(resource awselasticloadbalancingv2.ApplicationLoadBalancer) {
	if err := w.validateAssociateParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"associate",
		[]interface{}{resource},
	)
}

func (w *jsiiProxy_WebApplicationFirewall) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

