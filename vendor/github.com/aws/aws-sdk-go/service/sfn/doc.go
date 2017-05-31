// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sfn provides the client and types for making API
// requests to AWS Step Functions.
//
// AWS Step Functions is a web service that enables you to coordinate the components
// of distributed applications and microservices using visual workflows. You
// build applications from individual components that each perform a discrete
// function, or task, allowing you to scale and change applications quickly.
// Step Functions provides a graphical console to visualize the components of
// your application as a series of steps. It automatically triggers and tracks
// each step, and retries when there are errors, so your application executes
// in order and as expected, every time. Step Functions logs the state of each
// step, so when things do go wrong, you can diagnose and debug problems quickly.
//
// Step Functions manages the operations and underlying infrastructure for you
// to ensure your application is available at any scale. You can run tasks on
// the AWS cloud, on your own servers, or an any system that has access to AWS.
// Step Functions can be accessed and used with the Step Functions console,
// the AWS SDKs (included with your Beta release invitation email), or an HTTP
// API (the subject of this document).
//
// See https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23 for more information on this service.
//
// See sfn package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/sfn/
//
// Using the Client
//
// To use the client for AWS Step Functions you will first need
// to create a new instance of it.
//
// When creating a client for an AWS service you'll first need to have a Session
// already created. The Session provides configuration that can be shared
// between multiple service clients. Additional configuration can be applied to
// the Session and service's client when they are constructed. The aws package's
// Config type contains several fields such as Region for the AWS Region the
// client should make API requests too. The optional Config value can be provided
// as the variadic argument for Sessions and client creation.
//
// Once the service's client is created you can use it to make API requests the
// AWS service. These clients are safe to use concurrently.
//
//   // Create a session to share configuration, and load external configuration.
//   sess := session.Must(session.NewSession())
//
//   // Create the service's client with the session.
//   svc := sfn.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Step Functions client SFN for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/sfn/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.CreateActivity(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("CreateActivity result:")
//   fmt.Println(result)
//
// Using the Client with Context
//
// The service's client also provides methods to make API requests with a Context
// value. This allows you to control the timeout, and cancellation of pending
// requests. These methods also take request Option as variadic parameter to apply
// additional configuration to the API request.
//
//   ctx := context.Background()
//
//   result, err := svc.CreateActivityWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package sfn
