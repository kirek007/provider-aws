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

type AdditionalAuthenticationProviderInitParameters struct {

	// Authentication type. Valid values: API_KEY, AWS_IAM, AMAZON_COGNITO_USER_POOLS, OPENID_CONNECT, AWS_LAMBDA
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// Nested argument containing Lambda authorizer configuration. Defined below.
	LambdaAuthorizerConfig []LambdaAuthorizerConfigInitParameters `json:"lambdaAuthorizerConfig,omitempty" tf:"lambda_authorizer_config,omitempty"`

	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenIDConnectConfig []OpenIDConnectConfigInitParameters `json:"openidConnectConfig,omitempty" tf:"openid_connect_config,omitempty"`

	// Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig []UserPoolConfigInitParameters `json:"userPoolConfig,omitempty" tf:"user_pool_config,omitempty"`
}

type AdditionalAuthenticationProviderObservation struct {

	// Authentication type. Valid values: API_KEY, AWS_IAM, AMAZON_COGNITO_USER_POOLS, OPENID_CONNECT, AWS_LAMBDA
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// Nested argument containing Lambda authorizer configuration. Defined below.
	LambdaAuthorizerConfig []LambdaAuthorizerConfigObservation `json:"lambdaAuthorizerConfig,omitempty" tf:"lambda_authorizer_config,omitempty"`

	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenIDConnectConfig []OpenIDConnectConfigObservation `json:"openidConnectConfig,omitempty" tf:"openid_connect_config,omitempty"`

	// Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig []UserPoolConfigObservation `json:"userPoolConfig,omitempty" tf:"user_pool_config,omitempty"`
}

type AdditionalAuthenticationProviderParameters struct {

	// Authentication type. Valid values: API_KEY, AWS_IAM, AMAZON_COGNITO_USER_POOLS, OPENID_CONNECT, AWS_LAMBDA
	// +kubebuilder:validation:Optional
	AuthenticationType *string `json:"authenticationType" tf:"authentication_type,omitempty"`

	// Nested argument containing Lambda authorizer configuration. Defined below.
	// +kubebuilder:validation:Optional
	LambdaAuthorizerConfig []LambdaAuthorizerConfigParameters `json:"lambdaAuthorizerConfig,omitempty" tf:"lambda_authorizer_config,omitempty"`

	// Nested argument containing OpenID Connect configuration. Defined below.
	// +kubebuilder:validation:Optional
	OpenIDConnectConfig []OpenIDConnectConfigParameters `json:"openidConnectConfig,omitempty" tf:"openid_connect_config,omitempty"`

	// Amazon Cognito User Pool configuration. Defined below.
	// +kubebuilder:validation:Optional
	UserPoolConfig []UserPoolConfigParameters `json:"userPoolConfig,omitempty" tf:"user_pool_config,omitempty"`
}

type GraphQLAPIInitParameters struct {

	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProvider []AdditionalAuthenticationProviderInitParameters `json:"additionalAuthenticationProvider,omitempty" tf:"additional_authentication_provider,omitempty"`

	// Authentication type. Valid values: API_KEY, AWS_IAM, AMAZON_COGNITO_USER_POOLS, OPENID_CONNECT, AWS_LAMBDA
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// Nested argument containing Lambda authorizer configuration. Defined below.
	LambdaAuthorizerConfig []GraphQLAPILambdaAuthorizerConfigInitParameters `json:"lambdaAuthorizerConfig,omitempty" tf:"lambda_authorizer_config,omitempty"`

	// Nested argument containing logging configuration. Defined below.
	LogConfig []LogConfigInitParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// User-supplied name for the GraphqlApi.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenIDConnectConfig []GraphQLAPIOpenIDConnectConfigInitParameters `json:"openidConnectConfig,omitempty" tf:"openid_connect_config,omitempty"`

	// Schema definition, in GraphQL schema language format.
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig []GraphQLAPIUserPoolConfigInitParameters `json:"userPoolConfig,omitempty" tf:"user_pool_config,omitempty"`

	// Sets the value of the GraphQL API to public (GLOBAL) or private (PRIVATE). If no value is provided, the visibility will be set to GLOBAL by default. This value cannot be changed once the API has been created.
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`

	// Whether tracing with X-ray is enabled. Defaults to false.
	XrayEnabled *bool `json:"xrayEnabled,omitempty" tf:"xray_enabled,omitempty"`
}

type GraphQLAPILambdaAuthorizerConfigInitParameters struct {

	// Number of seconds a response should be cached for. The default is 5 minutes (300 seconds). The Lambda function can override this by returning a ttlOverride key in its response. A value of 0 disables caching of responses. Minimum value of 0. Maximum value of 3600.
	AuthorizerResultTTLInSeconds *int64 `json:"authorizerResultTtlInSeconds,omitempty" tf:"authorizer_result_ttl_in_seconds,omitempty"`

	// ARN of the Lambda function to be called for authorization. Note: This Lambda function must have a resource-based policy assigned to it, to allow lambda:InvokeFunction from service principal appsync.amazonaws.com.
	AuthorizerURI *string `json:"authorizerUri,omitempty" tf:"authorizer_uri,omitempty"`

	// Regular expression for validation of tokens before the Lambda function is called.
	IdentityValidationExpression *string `json:"identityValidationExpression,omitempty" tf:"identity_validation_expression,omitempty"`
}

type GraphQLAPILambdaAuthorizerConfigObservation struct {

	// Number of seconds a response should be cached for. The default is 5 minutes (300 seconds). The Lambda function can override this by returning a ttlOverride key in its response. A value of 0 disables caching of responses. Minimum value of 0. Maximum value of 3600.
	AuthorizerResultTTLInSeconds *int64 `json:"authorizerResultTtlInSeconds,omitempty" tf:"authorizer_result_ttl_in_seconds,omitempty"`

	// ARN of the Lambda function to be called for authorization. Note: This Lambda function must have a resource-based policy assigned to it, to allow lambda:InvokeFunction from service principal appsync.amazonaws.com.
	AuthorizerURI *string `json:"authorizerUri,omitempty" tf:"authorizer_uri,omitempty"`

	// Regular expression for validation of tokens before the Lambda function is called.
	IdentityValidationExpression *string `json:"identityValidationExpression,omitempty" tf:"identity_validation_expression,omitempty"`
}

type GraphQLAPILambdaAuthorizerConfigParameters struct {

	// Number of seconds a response should be cached for. The default is 5 minutes (300 seconds). The Lambda function can override this by returning a ttlOverride key in its response. A value of 0 disables caching of responses. Minimum value of 0. Maximum value of 3600.
	// +kubebuilder:validation:Optional
	AuthorizerResultTTLInSeconds *int64 `json:"authorizerResultTtlInSeconds,omitempty" tf:"authorizer_result_ttl_in_seconds,omitempty"`

	// ARN of the Lambda function to be called for authorization. Note: This Lambda function must have a resource-based policy assigned to it, to allow lambda:InvokeFunction from service principal appsync.amazonaws.com.
	// +kubebuilder:validation:Optional
	AuthorizerURI *string `json:"authorizerUri" tf:"authorizer_uri,omitempty"`

	// Regular expression for validation of tokens before the Lambda function is called.
	// +kubebuilder:validation:Optional
	IdentityValidationExpression *string `json:"identityValidationExpression,omitempty" tf:"identity_validation_expression,omitempty"`
}

type GraphQLAPIObservation struct {

	// One or more additional authentication providers for the GraphqlApi. Defined below.
	AdditionalAuthenticationProvider []AdditionalAuthenticationProviderObservation `json:"additionalAuthenticationProvider,omitempty" tf:"additional_authentication_provider,omitempty"`

	// ARN
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Authentication type. Valid values: API_KEY, AWS_IAM, AMAZON_COGNITO_USER_POOLS, OPENID_CONNECT, AWS_LAMBDA
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// API ID
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Nested argument containing Lambda authorizer configuration. Defined below.
	LambdaAuthorizerConfig []GraphQLAPILambdaAuthorizerConfigObservation `json:"lambdaAuthorizerConfig,omitempty" tf:"lambda_authorizer_config,omitempty"`

	// Nested argument containing logging configuration. Defined below.
	LogConfig []LogConfigObservation `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// User-supplied name for the GraphqlApi.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Nested argument containing OpenID Connect configuration. Defined below.
	OpenIDConnectConfig []GraphQLAPIOpenIDConnectConfigObservation `json:"openidConnectConfig,omitempty" tf:"openid_connect_config,omitempty"`

	// Schema definition, in GraphQL schema language format.
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Map of URIs associated with the APIE.g., uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql
	Uris map[string]*string `json:"uris,omitempty" tf:"uris,omitempty"`

	// Amazon Cognito User Pool configuration. Defined below.
	UserPoolConfig []GraphQLAPIUserPoolConfigObservation `json:"userPoolConfig,omitempty" tf:"user_pool_config,omitempty"`

	// Sets the value of the GraphQL API to public (GLOBAL) or private (PRIVATE). If no value is provided, the visibility will be set to GLOBAL by default. This value cannot be changed once the API has been created.
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`

	// Whether tracing with X-ray is enabled. Defaults to false.
	XrayEnabled *bool `json:"xrayEnabled,omitempty" tf:"xray_enabled,omitempty"`
}

type GraphQLAPIOpenIDConnectConfigInitParameters struct {

	// Number of milliseconds a token is valid after being authenticated.
	AuthTTL *int64 `json:"authTtl,omitempty" tf:"auth_ttl,omitempty"`

	// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Number of milliseconds a token is valid after being issued to a user.
	IatTTL *int64 `json:"iatTtl,omitempty" tf:"iat_ttl,omitempty"`

	// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`
}

type GraphQLAPIOpenIDConnectConfigObservation struct {

	// Number of milliseconds a token is valid after being authenticated.
	AuthTTL *int64 `json:"authTtl,omitempty" tf:"auth_ttl,omitempty"`

	// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Number of milliseconds a token is valid after being issued to a user.
	IatTTL *int64 `json:"iatTtl,omitempty" tf:"iat_ttl,omitempty"`

	// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`
}

type GraphQLAPIOpenIDConnectConfigParameters struct {

	// Number of milliseconds a token is valid after being authenticated.
	// +kubebuilder:validation:Optional
	AuthTTL *int64 `json:"authTtl,omitempty" tf:"auth_ttl,omitempty"`

	// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Number of milliseconds a token is valid after being issued to a user.
	// +kubebuilder:validation:Optional
	IatTTL *int64 `json:"iatTtl,omitempty" tf:"iat_ttl,omitempty"`

	// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer" tf:"issuer,omitempty"`
}

type GraphQLAPIParameters struct {

	// One or more additional authentication providers for the GraphqlApi. Defined below.
	// +kubebuilder:validation:Optional
	AdditionalAuthenticationProvider []AdditionalAuthenticationProviderParameters `json:"additionalAuthenticationProvider,omitempty" tf:"additional_authentication_provider,omitempty"`

	// Authentication type. Valid values: API_KEY, AWS_IAM, AMAZON_COGNITO_USER_POOLS, OPENID_CONNECT, AWS_LAMBDA
	// +kubebuilder:validation:Optional
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// Nested argument containing Lambda authorizer configuration. Defined below.
	// +kubebuilder:validation:Optional
	LambdaAuthorizerConfig []GraphQLAPILambdaAuthorizerConfigParameters `json:"lambdaAuthorizerConfig,omitempty" tf:"lambda_authorizer_config,omitempty"`

	// Nested argument containing logging configuration. Defined below.
	// +kubebuilder:validation:Optional
	LogConfig []LogConfigParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// User-supplied name for the GraphqlApi.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Nested argument containing OpenID Connect configuration. Defined below.
	// +kubebuilder:validation:Optional
	OpenIDConnectConfig []GraphQLAPIOpenIDConnectConfigParameters `json:"openidConnectConfig,omitempty" tf:"openid_connect_config,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Schema definition, in GraphQL schema language format.
	// +kubebuilder:validation:Optional
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Amazon Cognito User Pool configuration. Defined below.
	// +kubebuilder:validation:Optional
	UserPoolConfig []GraphQLAPIUserPoolConfigParameters `json:"userPoolConfig,omitempty" tf:"user_pool_config,omitempty"`

	// Sets the value of the GraphQL API to public (GLOBAL) or private (PRIVATE). If no value is provided, the visibility will be set to GLOBAL by default. This value cannot be changed once the API has been created.
	// +kubebuilder:validation:Optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`

	// Whether tracing with X-ray is enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	XrayEnabled *bool `json:"xrayEnabled,omitempty" tf:"xray_enabled,omitempty"`
}

type GraphQLAPIUserPoolConfigInitParameters struct {

	// Regular expression for validating the incoming Amazon Cognito User Pool app client ID.
	AppIDClientRegex *string `json:"appIdClientRegex,omitempty" tf:"app_id_client_regex,omitempty"`

	// AWS region in which the user pool was created.
	AwsRegion *string `json:"awsRegion,omitempty" tf:"aws_region,omitempty"`

	// Action that you want your GraphQL API to take when a request that uses Amazon Cognito User Pool authentication doesn't match the Amazon Cognito User Pool configuration. Valid: ALLOW and DENY
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`
}

type GraphQLAPIUserPoolConfigObservation struct {

	// Regular expression for validating the incoming Amazon Cognito User Pool app client ID.
	AppIDClientRegex *string `json:"appIdClientRegex,omitempty" tf:"app_id_client_regex,omitempty"`

	// AWS region in which the user pool was created.
	AwsRegion *string `json:"awsRegion,omitempty" tf:"aws_region,omitempty"`

	// Action that you want your GraphQL API to take when a request that uses Amazon Cognito User Pool authentication doesn't match the Amazon Cognito User Pool configuration. Valid: ALLOW and DENY
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// User pool ID.
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`
}

type GraphQLAPIUserPoolConfigParameters struct {

	// Regular expression for validating the incoming Amazon Cognito User Pool app client ID.
	// +kubebuilder:validation:Optional
	AppIDClientRegex *string `json:"appIdClientRegex,omitempty" tf:"app_id_client_regex,omitempty"`

	// AWS region in which the user pool was created.
	// +kubebuilder:validation:Optional
	AwsRegion *string `json:"awsRegion,omitempty" tf:"aws_region,omitempty"`

	// Action that you want your GraphQL API to take when a request that uses Amazon Cognito User Pool authentication doesn't match the Amazon Cognito User Pool configuration. Valid: ALLOW and DENY
	// +kubebuilder:validation:Optional
	DefaultAction *string `json:"defaultAction" tf:"default_action,omitempty"`

	// User pool ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPool
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`

	// Reference to a UserPool in cognitoidp to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDRef *v1.Reference `json:"userPoolIdRef,omitempty" tf:"-"`

	// Selector for a UserPool in cognitoidp to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDSelector *v1.Selector `json:"userPoolIdSelector,omitempty" tf:"-"`
}

type LambdaAuthorizerConfigInitParameters struct {

	// Number of seconds a response should be cached for. The default is 5 minutes (300 seconds). The Lambda function can override this by returning a ttlOverride key in its response. A value of 0 disables caching of responses. Minimum value of 0. Maximum value of 3600.
	AuthorizerResultTTLInSeconds *int64 `json:"authorizerResultTtlInSeconds,omitempty" tf:"authorizer_result_ttl_in_seconds,omitempty"`

	// ARN of the Lambda function to be called for authorization. Note: This Lambda function must have a resource-based policy assigned to it, to allow lambda:InvokeFunction from service principal appsync.amazonaws.com.
	AuthorizerURI *string `json:"authorizerUri,omitempty" tf:"authorizer_uri,omitempty"`

	// Regular expression for validation of tokens before the Lambda function is called.
	IdentityValidationExpression *string `json:"identityValidationExpression,omitempty" tf:"identity_validation_expression,omitempty"`
}

type LambdaAuthorizerConfigObservation struct {

	// Number of seconds a response should be cached for. The default is 5 minutes (300 seconds). The Lambda function can override this by returning a ttlOverride key in its response. A value of 0 disables caching of responses. Minimum value of 0. Maximum value of 3600.
	AuthorizerResultTTLInSeconds *int64 `json:"authorizerResultTtlInSeconds,omitempty" tf:"authorizer_result_ttl_in_seconds,omitempty"`

	// ARN of the Lambda function to be called for authorization. Note: This Lambda function must have a resource-based policy assigned to it, to allow lambda:InvokeFunction from service principal appsync.amazonaws.com.
	AuthorizerURI *string `json:"authorizerUri,omitempty" tf:"authorizer_uri,omitempty"`

	// Regular expression for validation of tokens before the Lambda function is called.
	IdentityValidationExpression *string `json:"identityValidationExpression,omitempty" tf:"identity_validation_expression,omitempty"`
}

type LambdaAuthorizerConfigParameters struct {

	// Number of seconds a response should be cached for. The default is 5 minutes (300 seconds). The Lambda function can override this by returning a ttlOverride key in its response. A value of 0 disables caching of responses. Minimum value of 0. Maximum value of 3600.
	// +kubebuilder:validation:Optional
	AuthorizerResultTTLInSeconds *int64 `json:"authorizerResultTtlInSeconds,omitempty" tf:"authorizer_result_ttl_in_seconds,omitempty"`

	// ARN of the Lambda function to be called for authorization. Note: This Lambda function must have a resource-based policy assigned to it, to allow lambda:InvokeFunction from service principal appsync.amazonaws.com.
	// +kubebuilder:validation:Optional
	AuthorizerURI *string `json:"authorizerUri" tf:"authorizer_uri,omitempty"`

	// Regular expression for validation of tokens before the Lambda function is called.
	// +kubebuilder:validation:Optional
	IdentityValidationExpression *string `json:"identityValidationExpression,omitempty" tf:"identity_validation_expression,omitempty"`
}

type LogConfigInitParameters struct {

	// Set to TRUE to exclude sections that contain information such as headers, context, and evaluated mapping templates, regardless of logging  level. Valid values: true, false. Default value: false
	ExcludeVerboseContent *bool `json:"excludeVerboseContent,omitempty" tf:"exclude_verbose_content,omitempty"`

	// Field logging level. Valid values: ALL, ERROR, NONE.
	FieldLogLevel *string `json:"fieldLogLevel,omitempty" tf:"field_log_level,omitempty"`
}

type LogConfigObservation struct {

	// Amazon Resource Name of the service role that AWS AppSync will assume to publish to Amazon CloudWatch logs in your account.
	CloudwatchLogsRoleArn *string `json:"cloudwatchLogsRoleArn,omitempty" tf:"cloudwatch_logs_role_arn,omitempty"`

	// Set to TRUE to exclude sections that contain information such as headers, context, and evaluated mapping templates, regardless of logging  level. Valid values: true, false. Default value: false
	ExcludeVerboseContent *bool `json:"excludeVerboseContent,omitempty" tf:"exclude_verbose_content,omitempty"`

	// Field logging level. Valid values: ALL, ERROR, NONE.
	FieldLogLevel *string `json:"fieldLogLevel,omitempty" tf:"field_log_level,omitempty"`
}

type LogConfigParameters struct {

	// Amazon Resource Name of the service role that AWS AppSync will assume to publish to Amazon CloudWatch logs in your account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	CloudwatchLogsRoleArn *string `json:"cloudwatchLogsRoleArn,omitempty" tf:"cloudwatch_logs_role_arn,omitempty"`

	// Reference to a Role in iam to populate cloudwatchLogsRoleArn.
	// +kubebuilder:validation:Optional
	CloudwatchLogsRoleArnRef *v1.Reference `json:"cloudwatchLogsRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate cloudwatchLogsRoleArn.
	// +kubebuilder:validation:Optional
	CloudwatchLogsRoleArnSelector *v1.Selector `json:"cloudwatchLogsRoleArnSelector,omitempty" tf:"-"`

	// Set to TRUE to exclude sections that contain information such as headers, context, and evaluated mapping templates, regardless of logging  level. Valid values: true, false. Default value: false
	// +kubebuilder:validation:Optional
	ExcludeVerboseContent *bool `json:"excludeVerboseContent,omitempty" tf:"exclude_verbose_content,omitempty"`

	// Field logging level. Valid values: ALL, ERROR, NONE.
	// +kubebuilder:validation:Optional
	FieldLogLevel *string `json:"fieldLogLevel" tf:"field_log_level,omitempty"`
}

type OpenIDConnectConfigInitParameters struct {

	// Number of milliseconds a token is valid after being authenticated.
	AuthTTL *int64 `json:"authTtl,omitempty" tf:"auth_ttl,omitempty"`

	// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Number of milliseconds a token is valid after being issued to a user.
	IatTTL *int64 `json:"iatTtl,omitempty" tf:"iat_ttl,omitempty"`

	// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`
}

type OpenIDConnectConfigObservation struct {

	// Number of milliseconds a token is valid after being authenticated.
	AuthTTL *int64 `json:"authTtl,omitempty" tf:"auth_ttl,omitempty"`

	// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Number of milliseconds a token is valid after being issued to a user.
	IatTTL *int64 `json:"iatTtl,omitempty" tf:"iat_ttl,omitempty"`

	// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`
}

type OpenIDConnectConfigParameters struct {

	// Number of milliseconds a token is valid after being authenticated.
	// +kubebuilder:validation:Optional
	AuthTTL *int64 `json:"authTtl,omitempty" tf:"auth_ttl,omitempty"`

	// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Number of milliseconds a token is valid after being issued to a user.
	// +kubebuilder:validation:Optional
	IatTTL *int64 `json:"iatTtl,omitempty" tf:"iat_ttl,omitempty"`

	// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer" tf:"issuer,omitempty"`
}

type UserPoolConfigInitParameters struct {

	// Regular expression for validating the incoming Amazon Cognito User Pool app client ID.
	AppIDClientRegex *string `json:"appIdClientRegex,omitempty" tf:"app_id_client_regex,omitempty"`

	// AWS region in which the user pool was created.
	AwsRegion *string `json:"awsRegion,omitempty" tf:"aws_region,omitempty"`

	// User pool ID.
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`
}

type UserPoolConfigObservation struct {

	// Regular expression for validating the incoming Amazon Cognito User Pool app client ID.
	AppIDClientRegex *string `json:"appIdClientRegex,omitempty" tf:"app_id_client_regex,omitempty"`

	// AWS region in which the user pool was created.
	AwsRegion *string `json:"awsRegion,omitempty" tf:"aws_region,omitempty"`

	// User pool ID.
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`
}

type UserPoolConfigParameters struct {

	// Regular expression for validating the incoming Amazon Cognito User Pool app client ID.
	// +kubebuilder:validation:Optional
	AppIDClientRegex *string `json:"appIdClientRegex,omitempty" tf:"app_id_client_regex,omitempty"`

	// AWS region in which the user pool was created.
	// +kubebuilder:validation:Optional
	AwsRegion *string `json:"awsRegion,omitempty" tf:"aws_region,omitempty"`

	// User pool ID.
	// +kubebuilder:validation:Optional
	UserPoolID *string `json:"userPoolId" tf:"user_pool_id,omitempty"`
}

// GraphQLAPISpec defines the desired state of GraphQLAPI
type GraphQLAPISpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GraphQLAPIParameters `json:"forProvider"`
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
	InitProvider GraphQLAPIInitParameters `json:"initProvider,omitempty"`
}

// GraphQLAPIStatus defines the observed state of GraphQLAPI.
type GraphQLAPIStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GraphQLAPIObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GraphQLAPI is the Schema for the GraphQLAPIs API. Provides an AppSync GraphQL API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GraphQLAPI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authenticationType) || has(self.initProvider.authenticationType)",message="authenticationType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   GraphQLAPISpec   `json:"spec"`
	Status GraphQLAPIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GraphQLAPIList contains a list of GraphQLAPIs
type GraphQLAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GraphQLAPI `json:"items"`
}

// Repository type metadata.
var (
	GraphQLAPI_Kind             = "GraphQLAPI"
	GraphQLAPI_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GraphQLAPI_Kind}.String()
	GraphQLAPI_KindAPIVersion   = GraphQLAPI_Kind + "." + CRDGroupVersion.String()
	GraphQLAPI_GroupVersionKind = CRDGroupVersion.WithKind(GraphQLAPI_Kind)
)

func init() {
	SchemeBuilder.Register(&GraphQLAPI{}, &GraphQLAPIList{})
}
