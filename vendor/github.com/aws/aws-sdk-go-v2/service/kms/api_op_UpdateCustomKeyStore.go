// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the properties of a custom key store. You can use this operation to
// change the properties of an CloudHSM key store or an external key store. Use the
// required CustomKeyStoreId parameter to identify the custom key store. Use the
// remaining optional parameters to change its properties. This operation does not
// return any property values. To verify the updated property values, use the
// DescribeCustomKeyStores operation. This operation is part of the custom key
// stores (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// feature in KMS, which combines the convenience and extensive integration of KMS
// with the isolation and control of a key store that you own and manage. When
// updating the properties of an external key store, verify that the updated
// settings connect your key store, via the external key store proxy, to the same
// external key manager as the previous settings, or to a backup or snapshot of the
// external key manager with the same cryptographic keys. If the updated connection
// settings fail, you can fix them and retry, although an extended delay might
// disrupt Amazon Web Services services. However, if KMS permanently loses its
// access to cryptographic keys, ciphertext encrypted under those keys is
// unrecoverable. For external key stores: Some external key managers provide a
// simpler method for updating an external key store. For details, see your
// external key manager documentation. When updating an external key store in the
// KMS console, you can upload a JSON-based proxy configuration file with the
// desired values. You cannot upload the proxy configuration file to the
// UpdateCustomKeyStore operation. However, you can use the file to help you
// determine the correct values for the UpdateCustomKeyStore parameters. For an
// CloudHSM key store, you can use this operation to change the custom key store
// friendly name ( NewCustomKeyStoreName ), to tell KMS about a change to the
// kmsuser crypto user password ( KeyStorePassword ), or to associate the custom
// key store with a different, but related, CloudHSM cluster ( CloudHsmClusterId ).
// To update any property of an CloudHSM key store, the ConnectionState of the
// CloudHSM key store must be DISCONNECTED . For an external key store, you can use
// this operation to change the custom key store friendly name (
// NewCustomKeyStoreName ), or to tell KMS about a change to the external key store
// proxy authentication credentials ( XksProxyAuthenticationCredential ),
// connection method ( XksProxyConnectivity ), external proxy endpoint (
// XksProxyUriEndpoint ) and path ( XksProxyUriPath ). For external key stores with
// an XksProxyConnectivity of VPC_ENDPOINT_SERVICE , you can also update the Amazon
// VPC endpoint service name ( XksProxyVpcEndpointServiceName ). To update most
// properties of an external key store, the ConnectionState of the external key
// store must be DISCONNECTED . However, you can update the CustomKeyStoreName ,
// XksProxyAuthenticationCredential , and XksProxyUriPath of an external key store
// when it is in the CONNECTED or DISCONNECTED state. If your update requires a
// DISCONNECTED state, before using UpdateCustomKeyStore , use the
// DisconnectCustomKeyStore operation to disconnect the custom key store. After the
// UpdateCustomKeyStore operation completes, use the ConnectCustomKeyStore to
// reconnect the custom key store. To find the ConnectionState of the custom key
// store, use the DescribeCustomKeyStores operation. Before updating the custom
// key store, verify that the new values allow KMS to connect the custom key store
// to its backing key store. For example, before you change the XksProxyUriPath
// value, verify that the external key store proxy is reachable at the new path. If
// the operation succeeds, it returns a JSON object with no properties.
// Cross-account use: No. You cannot perform this operation on a custom key store
// in a different Amazon Web Services account. Required permissions:
// kms:UpdateCustomKeyStore (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (IAM policy) Related operations:
//   - ConnectCustomKeyStore
//   - CreateCustomKeyStore
//   - DeleteCustomKeyStore
//   - DescribeCustomKeyStores
//   - DisconnectCustomKeyStore
func (c *Client) UpdateCustomKeyStore(ctx context.Context, params *UpdateCustomKeyStoreInput, optFns ...func(*Options)) (*UpdateCustomKeyStoreOutput, error) {
	if params == nil {
		params = &UpdateCustomKeyStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCustomKeyStore", params, optFns, c.addOperationUpdateCustomKeyStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCustomKeyStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCustomKeyStoreInput struct {

	// Identifies the custom key store that you want to update. Enter the ID of the
	// custom key store. To find the ID of a custom key store, use the
	// DescribeCustomKeyStores operation.
	//
	// This member is required.
	CustomKeyStoreId *string

	// Associates the custom key store with a related CloudHSM cluster. This parameter
	// is valid only for custom key stores with a CustomKeyStoreType of AWS_CLOUDHSM .
	// Enter the cluster ID of the cluster that you used to create the custom key store
	// or a cluster that shares a backup history and has the same cluster certificate
	// as the original cluster. You cannot use this parameter to associate a custom key
	// store with an unrelated cluster. In addition, the replacement cluster must
	// fulfill the requirements (https://docs.aws.amazon.com/kms/latest/developerguide/create-keystore.html#before-keystore)
	// for a cluster associated with a custom key store. To view the cluster
	// certificate of a cluster, use the DescribeClusters (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_DescribeClusters.html)
	// operation. To change this value, the CloudHSM key store must be disconnected.
	CloudHsmClusterId *string

	// Enter the current password of the kmsuser crypto user (CU) in the CloudHSM
	// cluster that is associated with the custom key store. This parameter is valid
	// only for custom key stores with a CustomKeyStoreType of AWS_CLOUDHSM . This
	// parameter tells KMS the current password of the kmsuser crypto user (CU). It
	// does not set or change the password of any users in the CloudHSM cluster. To
	// change this value, the CloudHSM key store must be disconnected.
	KeyStorePassword *string

	// Changes the friendly name of the custom key store to the value that you
	// specify. The custom key store name must be unique in the Amazon Web Services
	// account. Do not include confidential or sensitive information in this field.
	// This field may be displayed in plaintext in CloudTrail logs and other output. To
	// change this value, an CloudHSM key store must be disconnected. An external key
	// store can be connected or disconnected.
	NewCustomKeyStoreName *string

	// Changes the credentials that KMS uses to sign requests to the external key
	// store proxy (XKS proxy). This parameter is valid only for custom key stores with
	// a CustomKeyStoreType of EXTERNAL_KEY_STORE . You must specify both the
	// AccessKeyId and SecretAccessKey value in the authentication credential, even if
	// you are only updating one value. This parameter doesn't establish or change your
	// authentication credentials on the proxy. It just tells KMS the credential that
	// you established with your external key store proxy. For example, if you rotate
	// the credential on your external key store proxy, you can use this parameter to
	// update the credential in KMS. You can change this value when the external key
	// store is connected or disconnected.
	XksProxyAuthenticationCredential *types.XksProxyAuthenticationCredentialType

	// Changes the connectivity setting for the external key store. To indicate that
	// the external key store proxy uses a Amazon VPC endpoint service to communicate
	// with KMS, specify VPC_ENDPOINT_SERVICE . Otherwise, specify PUBLIC_ENDPOINT . If
	// you change the XksProxyConnectivity to VPC_ENDPOINT_SERVICE , you must also
	// change the XksProxyUriEndpoint and add an XksProxyVpcEndpointServiceName value.
	// If you change the XksProxyConnectivity to PUBLIC_ENDPOINT , you must also change
	// the XksProxyUriEndpoint and specify a null or empty string for the
	// XksProxyVpcEndpointServiceName value. To change this value, the external key
	// store must be disconnected.
	XksProxyConnectivity types.XksProxyConnectivityType

	// Changes the URI endpoint that KMS uses to connect to your external key store
	// proxy (XKS proxy). This parameter is valid only for custom key stores with a
	// CustomKeyStoreType of EXTERNAL_KEY_STORE . For external key stores with an
	// XksProxyConnectivity value of PUBLIC_ENDPOINT , the protocol must be HTTPS. For
	// external key stores with an XksProxyConnectivity value of VPC_ENDPOINT_SERVICE ,
	// specify https:// followed by the private DNS name associated with the VPC
	// endpoint service. Each external key store must use a different private DNS name.
	// The combined XksProxyUriEndpoint and XksProxyUriPath values must be unique in
	// the Amazon Web Services account and Region. To change this value, the external
	// key store must be disconnected.
	XksProxyUriEndpoint *string

	// Changes the base path to the proxy APIs for this external key store. To find
	// this value, see the documentation for your external key manager and external key
	// store proxy (XKS proxy). This parameter is valid only for custom key stores with
	// a CustomKeyStoreType of EXTERNAL_KEY_STORE . The value must start with / and
	// must end with /kms/xks/v1 , where v1 represents the version of the KMS external
	// key store proxy API. You can include an optional prefix between the required
	// elements such as /example/kms/xks/v1 . The combined XksProxyUriEndpoint and
	// XksProxyUriPath values must be unique in the Amazon Web Services account and
	// Region. You can change this value when the external key store is connected or
	// disconnected.
	XksProxyUriPath *string

	// Changes the name that KMS uses to identify the Amazon VPC endpoint service for
	// your external key store proxy (XKS proxy). This parameter is valid when the
	// CustomKeyStoreType is EXTERNAL_KEY_STORE and the XksProxyConnectivity is
	// VPC_ENDPOINT_SERVICE . To change this value, the external key store must be
	// disconnected.
	XksProxyVpcEndpointServiceName *string

	noSmithyDocumentSerde
}

type UpdateCustomKeyStoreOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCustomKeyStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCustomKeyStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCustomKeyStore{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCustomKeyStore"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateCustomKeyStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCustomKeyStore(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateCustomKeyStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCustomKeyStore",
	}
}
