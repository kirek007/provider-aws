/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AuthenticateCognitoInitParameters struct {

	// Query parameters to include in the redirect request to the authorization endpoint. Max: 10. Detailed below.
	AuthenticationRequestExtraParams map[string]*string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`

	// Behavior if the user is not authenticated. Valid values are deny, allow and authenticate.
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`

	// Set of user claims to be requested from the IdP.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Name of the cookie used to maintain session information.
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`

	// Maximum duration of the authentication session, in seconds.
	SessionTimeout *int64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// ARN of the Cognito user pool.
	UserPoolArn *string `json:"userPoolArn,omitempty" tf:"user_pool_arn,omitempty"`

	// ID of the Cognito user pool client.
	UserPoolClientID *string `json:"userPoolClientId,omitempty" tf:"user_pool_client_id,omitempty"`

	// Domain prefix or fully-qualified domain name of the Cognito user pool.
	UserPoolDomain *string `json:"userPoolDomain,omitempty" tf:"user_pool_domain,omitempty"`
}

type AuthenticateCognitoObservation struct {

	// Query parameters to include in the redirect request to the authorization endpoint. Max: 10. Detailed below.
	AuthenticationRequestExtraParams map[string]*string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`

	// Behavior if the user is not authenticated. Valid values are deny, allow and authenticate.
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`

	// Set of user claims to be requested from the IdP.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Name of the cookie used to maintain session information.
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`

	// Maximum duration of the authentication session, in seconds.
	SessionTimeout *int64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// ARN of the Cognito user pool.
	UserPoolArn *string `json:"userPoolArn,omitempty" tf:"user_pool_arn,omitempty"`

	// ID of the Cognito user pool client.
	UserPoolClientID *string `json:"userPoolClientId,omitempty" tf:"user_pool_client_id,omitempty"`

	// Domain prefix or fully-qualified domain name of the Cognito user pool.
	UserPoolDomain *string `json:"userPoolDomain,omitempty" tf:"user_pool_domain,omitempty"`
}

type AuthenticateCognitoParameters struct {

	// Query parameters to include in the redirect request to the authorization endpoint. Max: 10. Detailed below.
	// +kubebuilder:validation:Optional
	AuthenticationRequestExtraParams map[string]*string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`

	// Behavior if the user is not authenticated. Valid values are deny, allow and authenticate.
	// +kubebuilder:validation:Optional
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`

	// Set of user claims to be requested from the IdP.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Name of the cookie used to maintain session information.
	// +kubebuilder:validation:Optional
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`

	// Maximum duration of the authentication session, in seconds.
	// +kubebuilder:validation:Optional
	SessionTimeout *int64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// ARN of the Cognito user pool.
	// +kubebuilder:validation:Optional
	UserPoolArn *string `json:"userPoolArn" tf:"user_pool_arn,omitempty"`

	// ID of the Cognito user pool client.
	// +kubebuilder:validation:Optional
	UserPoolClientID *string `json:"userPoolClientId" tf:"user_pool_client_id,omitempty"`

	// Domain prefix or fully-qualified domain name of the Cognito user pool.
	// +kubebuilder:validation:Optional
	UserPoolDomain *string `json:"userPoolDomain" tf:"user_pool_domain,omitempty"`
}

type AuthenticateOidcInitParameters struct {

	// Query parameters to include in the redirect request to the authorization endpoint. Max: 10.
	AuthenticationRequestExtraParams map[string]*string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`

	// Authorization endpoint of the IdP.
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" tf:"authorization_endpoint,omitempty"`

	// OAuth 2.0 client identifier.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// OIDC issuer identifier of the IdP.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// Behavior if the user is not authenticated. Valid values: deny, allow and authenticate
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`

	// Set of user claims to be requested from the IdP.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Name of the cookie used to maintain session information.
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`

	// Maximum duration of the authentication session, in seconds.
	SessionTimeout *int64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// Token endpoint of the IdP.
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" tf:"token_endpoint,omitempty"`

	// User info endpoint of the IdP.
	UserInfoEndpoint *string `json:"userInfoEndpoint,omitempty" tf:"user_info_endpoint,omitempty"`
}

type AuthenticateOidcObservation struct {

	// Query parameters to include in the redirect request to the authorization endpoint. Max: 10.
	AuthenticationRequestExtraParams map[string]*string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`

	// Authorization endpoint of the IdP.
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" tf:"authorization_endpoint,omitempty"`

	// OAuth 2.0 client identifier.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// OIDC issuer identifier of the IdP.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// Behavior if the user is not authenticated. Valid values: deny, allow and authenticate
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`

	// Set of user claims to be requested from the IdP.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Name of the cookie used to maintain session information.
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`

	// Maximum duration of the authentication session, in seconds.
	SessionTimeout *int64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// Token endpoint of the IdP.
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" tf:"token_endpoint,omitempty"`

	// User info endpoint of the IdP.
	UserInfoEndpoint *string `json:"userInfoEndpoint,omitempty" tf:"user_info_endpoint,omitempty"`
}

type AuthenticateOidcParameters struct {

	// Query parameters to include in the redirect request to the authorization endpoint. Max: 10.
	// +kubebuilder:validation:Optional
	AuthenticationRequestExtraParams map[string]*string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`

	// Authorization endpoint of the IdP.
	// +kubebuilder:validation:Optional
	AuthorizationEndpoint *string `json:"authorizationEndpoint" tf:"authorization_endpoint,omitempty"`

	// OAuth 2.0 client identifier.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// OAuth 2.0 client secret.
	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// OIDC issuer identifier of the IdP.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer" tf:"issuer,omitempty"`

	// Behavior if the user is not authenticated. Valid values: deny, allow and authenticate
	// +kubebuilder:validation:Optional
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`

	// Set of user claims to be requested from the IdP.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Name of the cookie used to maintain session information.
	// +kubebuilder:validation:Optional
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`

	// Maximum duration of the authentication session, in seconds.
	// +kubebuilder:validation:Optional
	SessionTimeout *int64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// Token endpoint of the IdP.
	// +kubebuilder:validation:Optional
	TokenEndpoint *string `json:"tokenEndpoint" tf:"token_endpoint,omitempty"`

	// User info endpoint of the IdP.
	// +kubebuilder:validation:Optional
	UserInfoEndpoint *string `json:"userInfoEndpoint" tf:"user_info_endpoint,omitempty"`
}

type DefaultActionInitParameters struct {

	// Configuration block for using Amazon Cognito to authenticate users. Specify only when type is authenticate-cognito. Detailed below.
	AuthenticateCognito []AuthenticateCognitoInitParameters `json:"authenticateCognito,omitempty" tf:"authenticate_cognito,omitempty"`

	// Configuration block for an identity provider that is compliant with OpenID Connect (OIDC). Specify only when type is authenticate-oidc. Detailed below.
	AuthenticateOidc []AuthenticateOidcInitParameters `json:"authenticateOidc,omitempty" tf:"authenticate_oidc,omitempty"`

	// Information for creating an action that returns a custom HTTP response. Required if type is fixed-response.
	FixedResponse []FixedResponseInitParameters `json:"fixedResponse,omitempty" tf:"fixed_response,omitempty"`

	// Configuration block for creating an action that distributes requests among one or more target groups. Specify only if type is forward. If you specify both forward block and target_group_arn attribute, you can specify only one target group using forward and it must be the same target group specified in target_group_arn. Detailed below.
	Forward []ForwardInitParameters `json:"forward,omitempty" tf:"forward,omitempty"`

	// Order for the action. This value is required for rules with multiple actions. The action with the lowest value for order is performed first. Valid values are between 1 and 50000.
	Order *int64 `json:"order,omitempty" tf:"order,omitempty"`

	// Configuration block for creating a redirect action. Required if type is redirect. Detailed below.
	Redirect []RedirectInitParameters `json:"redirect,omitempty" tf:"redirect,omitempty"`

	// Type of routing action. Valid values are forward, redirect, fixed-response, authenticate-cognito and authenticate-oidc.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DefaultActionObservation struct {

	// Configuration block for using Amazon Cognito to authenticate users. Specify only when type is authenticate-cognito. Detailed below.
	AuthenticateCognito []AuthenticateCognitoObservation `json:"authenticateCognito,omitempty" tf:"authenticate_cognito,omitempty"`

	// Configuration block for an identity provider that is compliant with OpenID Connect (OIDC). Specify only when type is authenticate-oidc. Detailed below.
	AuthenticateOidc []AuthenticateOidcObservation `json:"authenticateOidc,omitempty" tf:"authenticate_oidc,omitempty"`

	// Information for creating an action that returns a custom HTTP response. Required if type is fixed-response.
	FixedResponse []FixedResponseObservation `json:"fixedResponse,omitempty" tf:"fixed_response,omitempty"`

	// Configuration block for creating an action that distributes requests among one or more target groups. Specify only if type is forward. If you specify both forward block and target_group_arn attribute, you can specify only one target group using forward and it must be the same target group specified in target_group_arn. Detailed below.
	Forward []ForwardObservation `json:"forward,omitempty" tf:"forward,omitempty"`

	// Order for the action. This value is required for rules with multiple actions. The action with the lowest value for order is performed first. Valid values are between 1 and 50000.
	Order *int64 `json:"order,omitempty" tf:"order,omitempty"`

	// Configuration block for creating a redirect action. Required if type is redirect. Detailed below.
	Redirect []RedirectObservation `json:"redirect,omitempty" tf:"redirect,omitempty"`

	// ARN of the Target Group to which to route traffic. Specify only if type is forward and you want to route to a single target group. To route to one or more target groups, use a forward block instead.
	TargetGroupArn *string `json:"targetGroupArn,omitempty" tf:"target_group_arn,omitempty"`

	// Type of routing action. Valid values are forward, redirect, fixed-response, authenticate-cognito and authenticate-oidc.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DefaultActionParameters struct {

	// Configuration block for using Amazon Cognito to authenticate users. Specify only when type is authenticate-cognito. Detailed below.
	// +kubebuilder:validation:Optional
	AuthenticateCognito []AuthenticateCognitoParameters `json:"authenticateCognito,omitempty" tf:"authenticate_cognito,omitempty"`

	// Configuration block for an identity provider that is compliant with OpenID Connect (OIDC). Specify only when type is authenticate-oidc. Detailed below.
	// +kubebuilder:validation:Optional
	AuthenticateOidc []AuthenticateOidcParameters `json:"authenticateOidc,omitempty" tf:"authenticate_oidc,omitempty"`

	// Information for creating an action that returns a custom HTTP response. Required if type is fixed-response.
	// +kubebuilder:validation:Optional
	FixedResponse []FixedResponseParameters `json:"fixedResponse,omitempty" tf:"fixed_response,omitempty"`

	// Configuration block for creating an action that distributes requests among one or more target groups. Specify only if type is forward. If you specify both forward block and target_group_arn attribute, you can specify only one target group using forward and it must be the same target group specified in target_group_arn. Detailed below.
	// +kubebuilder:validation:Optional
	Forward []ForwardParameters `json:"forward,omitempty" tf:"forward,omitempty"`

	// Order for the action. This value is required for rules with multiple actions. The action with the lowest value for order is performed first. Valid values are between 1 and 50000.
	// +kubebuilder:validation:Optional
	Order *int64 `json:"order,omitempty" tf:"order,omitempty"`

	// Configuration block for creating a redirect action. Required if type is redirect. Detailed below.
	// +kubebuilder:validation:Optional
	Redirect []RedirectParameters `json:"redirect,omitempty" tf:"redirect,omitempty"`

	// ARN of the Target Group to which to route traffic. Specify only if type is forward and you want to route to a single target group. To route to one or more target groups, use a forward block instead.
	// +crossplane:generate:reference:type=LBTargetGroup
	// +kubebuilder:validation:Optional
	TargetGroupArn *string `json:"targetGroupArn,omitempty" tf:"target_group_arn,omitempty"`

	// Reference to a LBTargetGroup to populate targetGroupArn.
	// +kubebuilder:validation:Optional
	TargetGroupArnRef *v1.Reference `json:"targetGroupArnRef,omitempty" tf:"-"`

	// Selector for a LBTargetGroup to populate targetGroupArn.
	// +kubebuilder:validation:Optional
	TargetGroupArnSelector *v1.Selector `json:"targetGroupArnSelector,omitempty" tf:"-"`

	// Type of routing action. Valid values are forward, redirect, fixed-response, authenticate-cognito and authenticate-oidc.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type FixedResponseInitParameters struct {

	// Content type. Valid values are text/plain, text/css, text/html, application/javascript and application/json.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Message body.
	MessageBody *string `json:"messageBody,omitempty" tf:"message_body,omitempty"`

	// HTTP response code. Valid values are 2XX, 4XX, or 5XX.
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type FixedResponseObservation struct {

	// Content type. Valid values are text/plain, text/css, text/html, application/javascript and application/json.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Message body.
	MessageBody *string `json:"messageBody,omitempty" tf:"message_body,omitempty"`

	// HTTP response code. Valid values are 2XX, 4XX, or 5XX.
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type FixedResponseParameters struct {

	// Content type. Valid values are text/plain, text/css, text/html, application/javascript and application/json.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// Message body.
	// +kubebuilder:validation:Optional
	MessageBody *string `json:"messageBody,omitempty" tf:"message_body,omitempty"`

	// HTTP response code. Valid values are 2XX, 4XX, or 5XX.
	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type ForwardInitParameters struct {

	// Configuration block for target group stickiness for the rule. Detailed below.
	Stickiness []StickinessInitParameters `json:"stickiness,omitempty" tf:"stickiness,omitempty"`

	// Set of 1-5 target group blocks. Detailed below.
	TargetGroup []TargetGroupInitParameters `json:"targetGroup,omitempty" tf:"target_group,omitempty"`
}

type ForwardObservation struct {

	// Configuration block for target group stickiness for the rule. Detailed below.
	Stickiness []StickinessObservation `json:"stickiness,omitempty" tf:"stickiness,omitempty"`

	// Set of 1-5 target group blocks. Detailed below.
	TargetGroup []TargetGroupObservation `json:"targetGroup,omitempty" tf:"target_group,omitempty"`
}

type ForwardParameters struct {

	// Configuration block for target group stickiness for the rule. Detailed below.
	// +kubebuilder:validation:Optional
	Stickiness []StickinessParameters `json:"stickiness,omitempty" tf:"stickiness,omitempty"`

	// Set of 1-5 target group blocks. Detailed below.
	// +kubebuilder:validation:Optional
	TargetGroup []TargetGroupParameters `json:"targetGroup" tf:"target_group,omitempty"`
}

type LBListenerInitParameters struct {

	// Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if protocol is TLS. Valid values are HTTP1Only, HTTP2Only, HTTP2Optional, HTTP2Preferred, and None.
	AlpnPolicy *string `json:"alpnPolicy,omitempty" tf:"alpn_policy,omitempty"`

	// ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the aws_lb_listener_certificate resource.
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Configuration block for default actions. Detailed below.
	DefaultAction []DefaultActionInitParameters `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// Port on which the load balancer is listening. Not valid for Gateway Load Balancers.
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are HTTP and HTTPS, with a default of HTTP. For Network Load Balancers, valid values are TCP, TLS, UDP, and TCP_UDP. Not valid to use UDP or TCP_UDP if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Name of the SSL Policy for the listener. Required if protocol is HTTPS or TLS.
	SSLPolicy *string `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LBListenerObservation struct {

	// Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if protocol is TLS. Valid values are HTTP1Only, HTTP2Only, HTTP2Optional, HTTP2Preferred, and None.
	AlpnPolicy *string `json:"alpnPolicy,omitempty" tf:"alpn_policy,omitempty"`

	// ARN of the listener (matches id).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the aws_lb_listener_certificate resource.
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Configuration block for default actions. Detailed below.
	DefaultAction []DefaultActionObservation `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// ARN of the listener (matches arn).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of the load balancer.
	LoadBalancerArn *string `json:"loadBalancerArn,omitempty" tf:"load_balancer_arn,omitempty"`

	// Port on which the load balancer is listening. Not valid for Gateway Load Balancers.
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are HTTP and HTTPS, with a default of HTTP. For Network Load Balancers, valid values are TCP, TLS, UDP, and TCP_UDP. Not valid to use UDP or TCP_UDP if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Name of the SSL Policy for the listener. Required if protocol is HTTPS or TLS.
	SSLPolicy *string `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LBListenerParameters struct {

	// Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if protocol is TLS. Valid values are HTTP1Only, HTTP2Only, HTTP2Optional, HTTP2Preferred, and None.
	// +kubebuilder:validation:Optional
	AlpnPolicy *string `json:"alpnPolicy,omitempty" tf:"alpn_policy,omitempty"`

	// ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the aws_lb_listener_certificate resource.
	// +kubebuilder:validation:Optional
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Configuration block for default actions. Detailed below.
	// +kubebuilder:validation:Optional
	DefaultAction []DefaultActionParameters `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// ARN of the load balancer.
	// +crossplane:generate:reference:type=LB
	// +kubebuilder:validation:Optional
	LoadBalancerArn *string `json:"loadBalancerArn,omitempty" tf:"load_balancer_arn,omitempty"`

	// Reference to a LB to populate loadBalancerArn.
	// +kubebuilder:validation:Optional
	LoadBalancerArnRef *v1.Reference `json:"loadBalancerArnRef,omitempty" tf:"-"`

	// Selector for a LB to populate loadBalancerArn.
	// +kubebuilder:validation:Optional
	LoadBalancerArnSelector *v1.Selector `json:"loadBalancerArnSelector,omitempty" tf:"-"`

	// Port on which the load balancer is listening. Not valid for Gateway Load Balancers.
	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are HTTP and HTTPS, with a default of HTTP. For Network Load Balancers, valid values are TCP, TLS, UDP, and TCP_UDP. Not valid to use UDP or TCP_UDP if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Name of the SSL Policy for the listener. Required if protocol is HTTPS or TLS.
	// +kubebuilder:validation:Optional
	SSLPolicy *string `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RedirectInitParameters struct {

	// Hostname. This component is not percent-encoded. The hostname can contain #{host}. Defaults to #{host}.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// Absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to /#{path}.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Port. Specify a value from 1 to 65535 or #{port}. Defaults to #{port}.
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol. Valid values are HTTP, HTTPS, or #{protocol}. Defaults to #{protocol}.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to #{query}.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// HTTP redirect code. The redirect is either permanent (HTTP_301) or temporary (HTTP_302).
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type RedirectObservation struct {

	// Hostname. This component is not percent-encoded. The hostname can contain #{host}. Defaults to #{host}.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// Absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to /#{path}.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Port. Specify a value from 1 to 65535 or #{port}. Defaults to #{port}.
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol. Valid values are HTTP, HTTPS, or #{protocol}. Defaults to #{protocol}.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to #{query}.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// HTTP redirect code. The redirect is either permanent (HTTP_301) or temporary (HTTP_302).
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type RedirectParameters struct {

	// Hostname. This component is not percent-encoded. The hostname can contain #{host}. Defaults to #{host}.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// Absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to /#{path}.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Port. Specify a value from 1 to 65535 or #{port}. Defaults to #{port}.
	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// Protocol. Valid values are HTTP, HTTPS, or #{protocol}. Defaults to #{protocol}.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to #{query}.
	// +kubebuilder:validation:Optional
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// HTTP redirect code. The redirect is either permanent (HTTP_301) or temporary (HTTP_302).
	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode" tf:"status_code,omitempty"`
}

type StickinessInitParameters struct {

	// Time period, in seconds, during which requests from a client should be routed to the same target group. The range is 1-604800 seconds (7 days).
	Duration *int64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// Whether target group stickiness is enabled. Default is false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type StickinessObservation struct {

	// Time period, in seconds, during which requests from a client should be routed to the same target group. The range is 1-604800 seconds (7 days).
	Duration *int64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// Whether target group stickiness is enabled. Default is false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type StickinessParameters struct {

	// Time period, in seconds, during which requests from a client should be routed to the same target group. The range is 1-604800 seconds (7 days).
	// +kubebuilder:validation:Optional
	Duration *int64 `json:"duration" tf:"duration,omitempty"`

	// Whether target group stickiness is enabled. Default is false.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type TargetGroupInitParameters struct {

	// Weight. The range is 0 to 999.
	Weight *int64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TargetGroupObservation struct {

	// ARN of the target group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Weight. The range is 0 to 999.
	Weight *int64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TargetGroupParameters struct {

	// ARN of the target group.
	// +crossplane:generate:reference:type=LBTargetGroup
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Reference to a LBTargetGroup to populate arn.
	// +kubebuilder:validation:Optional
	ArnRef *v1.Reference `json:"arnRef,omitempty" tf:"-"`

	// Selector for a LBTargetGroup to populate arn.
	// +kubebuilder:validation:Optional
	ArnSelector *v1.Selector `json:"arnSelector,omitempty" tf:"-"`

	// Weight. The range is 0 to 999.
	// +kubebuilder:validation:Optional
	Weight *int64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

// LBListenerSpec defines the desired state of LBListener
type LBListenerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBListenerParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider LBListenerInitParameters `json:"initProvider,omitempty"`
}

// LBListenerStatus defines the observed state of LBListener.
type LBListenerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBListenerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LBListener is the Schema for the LBListeners API. Provides a Load Balancer Listener resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LBListener struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultAction) || has(self.initProvider.defaultAction)",message="defaultAction is a required parameter"
	Spec   LBListenerSpec   `json:"spec"`
	Status LBListenerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBListenerList contains a list of LBListeners
type LBListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBListener `json:"items"`
}

// Repository type metadata.
var (
	LBListener_Kind             = "LBListener"
	LBListener_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBListener_Kind}.String()
	LBListener_KindAPIVersion   = LBListener_Kind + "." + CRDGroupVersion.String()
	LBListener_GroupVersionKind = CRDGroupVersion.WithKind(LBListener_Kind)
)

func init() {
	SchemeBuilder.Register(&LBListener{}, &LBListenerList{})
}
