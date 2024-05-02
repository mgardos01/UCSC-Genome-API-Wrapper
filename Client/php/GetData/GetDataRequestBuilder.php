<?php

namespace ApiSdk\GetData;

use ApiSdk\GetData\Sequence\SequenceRequestBuilder;
use ApiSdk\GetData\Track\TrackRequestBuilder;
use Microsoft\Kiota\Abstractions\BaseRequestBuilder;
use Microsoft\Kiota\Abstractions\RequestAdapter;

/**
 * Builds and executes requests for operations under /getData
*/
class GetDataRequestBuilder extends BaseRequestBuilder 
{
    /**
     * The sequence property
    */
    public function sequence(): SequenceRequestBuilder {
        return new SequenceRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /**
     * The track property
    */
    public function track(): TrackRequestBuilder {
        return new TrackRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /**
     * Instantiates a new GetDataRequestBuilder and sets the default values.
     * @param array<string, mixed>|string $pathParametersOrRawUrl Path parameters for the request or a String representing the raw URL.
     * @param RequestAdapter $requestAdapter The request adapter to use to execute the requests.
    */
    public function __construct($pathParametersOrRawUrl, RequestAdapter $requestAdapter) {
        parent::__construct($requestAdapter, [], '{+baseurl}/getData');
        if (is_array($pathParametersOrRawUrl)) {
            $this->pathParameters = $pathParametersOrRawUrl;
        } else {
            $this->pathParameters = ['request-raw-url' => $pathParametersOrRawUrl];
        }
    }

}
