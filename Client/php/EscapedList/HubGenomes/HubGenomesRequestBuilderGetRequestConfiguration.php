<?php

namespace ApiSdk\EscapedList\HubGenomes;

use Microsoft\Kiota\Abstractions\BaseRequestConfiguration;
use Microsoft\Kiota\Abstractions\RequestOption;

/**
 * Configuration for the request such as headers, query parameters, and middleware options.
*/
class HubGenomesRequestBuilderGetRequestConfiguration extends BaseRequestConfiguration 
{
    /**
     * @var HubGenomesRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public ?HubGenomesRequestBuilderGetQueryParameters $queryParameters = null;
    
    /**
     * Instantiates a new HubGenomesRequestBuilderGetRequestConfiguration and sets the default values.
     * @param array<string, array<string>|string>|null $headers Request headers
     * @param array<RequestOption>|null $options Request options
     * @param HubGenomesRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public function __construct(?array $headers = null, ?array $options = null, ?HubGenomesRequestBuilderGetQueryParameters $queryParameters = null) {
        parent::__construct($headers ?? [], $options ?? []);
        $this->queryParameters = $queryParameters;
    }

    /**
     * Instantiates a new HubGenomesRequestBuilderGetQueryParameters.
     * @param string|null $hubUrl 
     * @return HubGenomesRequestBuilderGetQueryParameters
    */
    public static function createQueryParameters(?string $hubUrl = null): HubGenomesRequestBuilderGetQueryParameters {
        return new HubGenomesRequestBuilderGetQueryParameters($hubUrl);
    }

}
