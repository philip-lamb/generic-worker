// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package tcworkermanager

import (
	"encoding/json"

	tcclient "github.com/taskcluster/taskcluster-client-go"
)

type (
	// Proof that this call is coming from the worker identified by the other fields.
	// The form of this proof varies depending on the provider type.
	AwsProviderType struct {

		// Instance identity document that is obtained by
		// curl http://169.254.169.254/latest/dynamic/instance-identity/document on the instance
		Document json.RawMessage `json:"document"`

		// The signature for instance identity document. Can be obtained by
		// curl http://169.254.169.254/latest/dynamic/instance-identity/signature on the instance
		Signature string `json:"signature"`
	}

	// The credentials the worker
	// will need to perform its work.  Specifically, credentials with scopes
	// * `assume:worker-pool:<workerPoolId>`
	// * `assume:worker-id:<workerGroup>/<workerId>`
	// * `queue:worker-id:<workerGroup>/<workerId>`
	// * `secrets:get:worker-pool:<workerPoolId>`
	// * `queue:claim-work:<workerPoolId>`
	Credentials struct {
		AccessToken string `json:"accessToken"`

		// Note that a certificate may not be provided, if the credentials are not temporary.
		Certificate string `json:"certificate,omitempty"`

		ClientID string `json:"clientId"`
	}

	// Proof that this call is coming from the worker identified by the other fields.
	// The form of this proof varies depending on the provider type.
	GoogleProviderType struct {

		// A JWT token as defined in [this google documentation](https://cloud.google.com/compute/docs/instances/verifying-instance-identity)
		Token string `json:"token"`
	}

	// A list of providers
	ProviderList struct {

		// Opaque `continuationToken` to be given as query-string option to get the
		// next set of workers in the worker-manager.
		// This property is only present if another request is necessary to fetch all
		// results. In practice the next request with a `continuationToken` may not
		// return additional results, but it can. Thus, you can only be sure to have
		// all the results if you've called `listWorkerPools` with `continuationToken`
		// until you get a result without a `continuationToken`.
		ContinuationToken string `json:"continuationToken,omitempty"`

		// List of all providers
		Providers []Var `json:"providers"`
	}

	// Request body to `registerWorker`.
	RegisterWorkerRequest struct {

		// The provider that had started the worker and responsible for managing it.
		// Can be different from the provider that's currently in the worker pool config.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		ProviderID string `json:"providerId"`

		// Worker group to which this worker belongs
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerGroup string `json:"workerGroup"`

		// Worker ID
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerID string `json:"workerId"`

		// Proof that this call is coming from the worker identified by the other fields.
		// The form of this proof varies depending on the provider type.
		//
		// One of:
		//   * GoogleProviderType
		//   * StaticProviderType1
		//   * AwsProviderType
		WorkerIdentityProof json.RawMessage `json:"workerIdentityProof"`

		// The ID of this worker pool (of the form `providerId/workerType` for compatibility)
		//
		// Syntax:     ^[a-zA-Z0-9-_]{1,38}/[a-z]([-a-z0-9]{0,36}[a-z0-9])?$
		WorkerPoolID string `json:"workerPoolId"`
	}

	// Response body to `registerWorker`.
	RegisterWorkerResponse struct {

		// The credentials the worker
		// will need to perform its work.  Specifically, credentials with scopes
		// * `assume:worker-pool:<workerPoolId>`
		// * `assume:worker-id:<workerGroup>/<workerId>`
		// * `queue:worker-id:<workerGroup>/<workerId>`
		// * `secrets:get:worker-pool:<workerPoolId>`
		// * `queue:claim-work:<workerPoolId>`
		Credentials Credentials `json:"credentials"`

		// Time at which the included credentials will expire.  Workers must either
		// re-register (for static workers) or terminate (for dynamically
		// provisioned workers) before this time.
		Expires tcclient.Time `json:"expires"`
	}

	// Provider-specific information
	StaticProviderType struct {

		// A secret value shared with the worker.  This value must be passed in the `workerIdentityProof` of the `registerWorker` method.
		// The ideal way to generate a secret of this form is `slugid() + slugid()`.
		//
		// Secrets are traded for Taskcluster credentials, and should be treated with similar care.
		// Each worker should have a distinct secret.
		//
		// Syntax:     ^[a-zA-Z0-9_-]{44}$
		StaticSecret string `json:"staticSecret"`
	}

	// Proof that this call is coming from the worker identified by the other fields.
	// The form of this proof varies depending on the provider type.
	StaticProviderType1 struct {

		// The secret value that was configured when the worker was created (in `createWorker`).
		//
		// Syntax:     ^[a-zA-Z0-9_-]{44}$
		StaticSecret string `json:"staticSecret"`
	}

	Var struct {

		// The id of this provider
		ProviderID string `json:"providerId"`

		// The provider implementation underlying this provider
		ProviderType string `json:"providerType"`
	}

	// Request to create a worker
	WorkerCreationRequest struct {

		// Date and time when this worker will be deleted from the DB
		Expires tcclient.Time `json:"expires"`

		// Provider-specific information
		//
		// One of:
		//   * StaticProviderType
		ProviderInfo json.RawMessage `json:"providerInfo,omitempty"`
	}

	// A report of an error from a worker.  This will be recorded with kind
	// `worker-error`.
	//
	// The worker's `workerGroup` and `workerId` will be added to `extra`.
	WorkerErrorReport struct {

		// A longer description of what occured in the error.
		//
		// Max length: 10240
		Description string `json:"description"`

		// Any extra structured information about this error
		//
		// Additional properties allowed
		Extra json.RawMessage `json:"extra"`

		// A general machine-readable way to identify this sort of error.
		//
		// Syntax:     [-a-z0-9]+
		// Max length: 128
		Kind string `json:"kind"`

		// A human-readable version of `kind`.
		//
		// Max length: 128
		Title string `json:"title"`

		// Worker group to which this worker belongs
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerGroup string `json:"workerGroup"`

		// Worker ID
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerID string `json:"workerId"`
	}

	// A complete worker definition.
	WorkerFullDefinition struct {

		// Date and time when this worker was created
		Created tcclient.Time `json:"created"`

		// Date and time when this worker will be deleted from the DB
		Expires tcclient.Time `json:"expires"`

		// The provider that had started the worker and responsible for managing it.
		// Can be different from the provider that's currently in the worker pool config.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		ProviderID string `json:"providerId"`

		// A string specifying the state this worker is in so far as worker-manager knows.
		//
		// Possible values:
		//   * "requested"
		//   * "running"
		//   * "stopped"
		State string `json:"state"`

		// Worker group to which this worker belongs
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerGroup string `json:"workerGroup"`

		// Worker ID
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		WorkerID string `json:"workerId"`

		// The ID of this worker pool (of the form `providerId/workerType` for compatibility)
		//
		// Syntax:     ^[a-zA-Z0-9-_]{1,38}/[a-z]([-a-z0-9]{0,36}[a-z0-9])?$
		WorkerPoolID string `json:"workerPoolId"`
	}

	// A list of workers in a given worker pool
	WorkerListInAGivenWorkerPool struct {

		// Opaque `continuationToken` to be given as query-string option to get the
		// next set of workers in the worker-manager.
		// This property is only present if another request is necessary to fetch all
		// results. In practice the next request with a `continuationToken` may not
		// return additional results, but it can. Thus, you can only be sure to have
		// all the results if you've called `listWorkerPools` with `continuationToken`
		// until you get a result without a `continuationToken`.
		ContinuationToken string `json:"continuationToken,omitempty"`

		// List of all workers in a given worker pool
		Workers []WorkerFullDefinition `json:"workers"`
	}

	// Fields that are defined by a user for a worker pool.
	// Used to create worker-pool definitions. There is a larger
	// set of fields for viewing since some parts are generated
	// by the service.
	WorkerPoolDefinition struct {

		// Additional properties allowed
		Config json.RawMessage `json:"config"`

		// A description of this worker pool.
		//
		// Max length: 10240
		Description string `json:"description"`

		// If true, the owner should be emailed on provisioning errors
		EmailOnError bool `json:"emailOnError"`

		// An email address to notify when there are provisioning errors for this
		// worker pool.
		Owner string `json:"owner"`

		// The provider responsible for managing this worker pool.
		//
		// If this value is `"null-provider"`, then the worker pool is pending deletion
		// once all existing workers have terminated.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		ProviderID string `json:"providerId"`
	}

	// Fields that are defined by a user for a worker pool.
	// Used to modify worker-pool definitions.
	//
	// The `workerPoolId`, `created`, and `lastModified` fields are optional and
	// allowed only to ease the common practice of getting a worker pool definition
	// with `workerPool(..)`, modifying it, and writing it back with
	// `updateWorkerPool(..).  `workerPoolId` must be correct if
	// supplied, and the values of `created` and `lastModified` are ignored.
	WorkerPoolDefinition1 struct {

		// Additional properties allowed
		Config json.RawMessage `json:"config"`

		// Ignored on update
		Created tcclient.Time `json:"created,omitempty"`

		// A description of this worker pool.
		//
		// Max length: 10240
		Description string `json:"description"`

		// If true, the owner should be emailed on provisioning errors
		EmailOnError bool `json:"emailOnError"`

		// Ignored on update
		LastModified tcclient.Time `json:"lastModified,omitempty"`

		// An email address to notify when there are provisioning errors for this
		// worker pool.
		Owner string `json:"owner"`

		// The provider responsible for managing this worker pool.
		//
		// If this value is `"null-provider"`, then the worker pool is pending deletion
		// once all existing workers have terminated.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		ProviderID string `json:"providerId"`

		// Syntax:     ^[a-zA-Z0-9-_]{1,38}/[a-z]([-a-z0-9]{0,36}[a-z0-9])?$
		WorkerPoolID string `json:"workerPoolId,omitempty"`
	}

	// A complete worker pool error definition.
	WorkerPoolError struct {

		// A longer description of what occured in the error.
		//
		// Max length: 10240
		Description string `json:"description"`

		// An arbitary unique identifier for this error
		//
		// Syntax:     ^[A-Za-z0-9_-]{8}[Q-T][A-Za-z0-9_-][CGKOSWaeimquy26-][A-Za-z0-9_-]{10}[AQgw]$
		ErrorID string `json:"errorId"`

		// Any extra structured information about this error
		//
		// Additional properties allowed
		Extra json.RawMessage `json:"extra"`

		// A general machine-readable way to identify this sort of error.
		//
		// Syntax:     [-a-z0-9]+
		// Max length: 128
		Kind string `json:"kind"`

		// Date and time when this error was reported
		Reported tcclient.Time `json:"reported"`

		// A human-readable version of `kind`.
		//
		// Max length: 128
		Title string `json:"title"`

		// The ID of this worker pool (of the form `providerId/workerType` for compatibility)
		//
		// Syntax:     ^[a-zA-Z0-9-_]{1,38}/[a-z]([-a-z0-9]{0,36}[a-z0-9])?$
		WorkerPoolID string `json:"workerPoolId"`
	}

	// A list of worker pool errors
	WorkerPoolErrorList struct {

		// Opaque `continuationToken` to be given as query-string option to get the
		// next set of worker-types in the worker-manager.
		// This property is only present if another request is necessary to fetch all
		// results. In practice the next request with a `continuationToken` may not
		// return additional results, but it can. Thus, you can only be sure to have
		// all the results if you've called `listWorkerPools` with `continuationToken`
		// until you get a result without a `continuationToken`.
		ContinuationToken string `json:"continuationToken,omitempty"`

		// List of worker pool errors
		WorkerPoolErrors []WorkerPoolError `json:"workerPoolErrors"`
	}

	// A complete worker pool definition.
	WorkerPoolFullDefinition struct {

		// Additional properties allowed
		Config json.RawMessage `json:"config"`

		// Date and time when this worker pool was created
		Created tcclient.Time `json:"created"`

		// A description of this worker pool.
		//
		// Max length: 10240
		Description string `json:"description"`

		// If true, the owner should be emailed on provisioning errors
		EmailOnError bool `json:"emailOnError"`

		// Date and time when this worker pool was last updated
		LastModified tcclient.Time `json:"lastModified"`

		// An email address to notify when there are provisioning errors for this
		// worker pool.
		Owner string `json:"owner"`

		// The provider responsible for managing this worker pool.
		//
		// If this value is `"null-provider"`, then the worker pool is pending deletion
		// once all existing workers have terminated.
		//
		// Syntax:     ^([a-zA-Z0-9-_]*)$
		// Min length: 1
		// Max length: 38
		ProviderID string `json:"providerId"`

		// The ID of this worker pool (of the form `providerId/workerType` for compatibility)
		//
		// Syntax:     ^[a-zA-Z0-9-_]{1,38}/[a-z]([-a-z0-9]{0,36}[a-z0-9])?$
		WorkerPoolID string `json:"workerPoolId,omitempty"`
	}

	// A list of worker pools
	WorkerPoolList struct {

		// Opaque `continuationToken` to be given as query-string option to get the
		// next set of worker-types in the worker-manager.
		// This property is only present if another request is necessary to fetch all
		// results. In practice the next request with a `continuationToken` may not
		// return additional results, but it can. Thus, you can only be sure to have
		// all the results if you've called `listWorkerPools` with `continuationToken`
		// until you get a result without a `continuationToken`.
		ContinuationToken string `json:"continuationToken,omitempty"`

		// List of all worker pools
		WorkerPools []WorkerPoolFullDefinition `json:"workerPools"`
	}
)
