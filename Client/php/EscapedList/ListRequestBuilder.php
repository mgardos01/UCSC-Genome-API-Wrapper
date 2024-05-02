<?php

namespace ApiSdk\EscapedList;

use ApiSdk\EscapedList\Chromosomes\ChromosomesRequestBuilder;
use ApiSdk\EscapedList\Files\FilesRequestBuilder;
use ApiSdk\EscapedList\HubGenomes\HubGenomesRequestBuilder;
use ApiSdk\EscapedList\PublicHubs\PublicHubsRequestBuilder;
use ApiSdk\EscapedList\Schema\SchemaRequestBuilder;
use ApiSdk\EscapedList\Tracks\TracksRequestBuilder;
use ApiSdk\EscapedList\UcscGenomes\UcscGenomesRequestBuilder;
use Microsoft\Kiota\Abstractions\BaseRequestBuilder;
use Microsoft\Kiota\Abstractions\RequestAdapter;

/**
 * Builds and executes requests for operations under /list
*/
class ListRequestBuilder extends BaseRequestBuilder 
{
    /**
     * The chromosomes property
    */
    public function chromosomes(): ChromosomesRequestBuilder {
        return new ChromosomesRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /**
     * The files property
    */
    public function files(): FilesRequestBuilder {
        return new FilesRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /**
     * The hubGenomes property
    */
    public function hubGenomes(): HubGenomesRequestBuilder {
        return new HubGenomesRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /**
     * The publicHubs property
    */
    public function publicHubs(): PublicHubsRequestBuilder {
        return new PublicHubsRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /**
     * The schema property
    */
    public function schema(): SchemaRequestBuilder {
        return new SchemaRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /**
     * The tracks property
    */
    public function tracks(): TracksRequestBuilder {
        return new TracksRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /**
     * The ucscGenomes property
    */
    public function ucscGenomes(): UcscGenomesRequestBuilder {
        return new UcscGenomesRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /**
     * Instantiates a new ListRequestBuilder and sets the default values.
     * @param array<string, mixed>|string $pathParametersOrRawUrl Path parameters for the request or a String representing the raw URL.
     * @param RequestAdapter $requestAdapter The request adapter to use to execute the requests.
    */
    public function __construct($pathParametersOrRawUrl, RequestAdapter $requestAdapter) {
        parent::__construct($requestAdapter, [], '{+baseurl}/list');
        if (is_array($pathParametersOrRawUrl)) {
            $this->pathParameters = $pathParametersOrRawUrl;
        } else {
            $this->pathParameters = ['request-raw-url' => $pathParametersOrRawUrl];
        }
    }

}
