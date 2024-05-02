<?php

namespace ApiSdk\EscapedList\Chromosomes;

use Microsoft\Kiota\Abstractions\BaseRequestConfiguration;
use Microsoft\Kiota\Abstractions\RequestOption;

/**
 * Configuration for the request such as headers, query parameters, and middleware options.
*/
class ChromosomesRequestBuilderGetRequestConfiguration extends BaseRequestConfiguration 
{
    /**
     * @var ChromosomesRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public ?ChromosomesRequestBuilderGetQueryParameters $queryParameters = null;
    
    /**
     * Instantiates a new ChromosomesRequestBuilderGetRequestConfiguration and sets the default values.
     * @param array<string, array<string>|string>|null $headers Request headers
     * @param array<RequestOption>|null $options Request options
     * @param ChromosomesRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public function __construct(?array $headers = null, ?array $options = null, ?ChromosomesRequestBuilderGetQueryParameters $queryParameters = null) {
        parent::__construct($headers ?? [], $options ?? []);
        $this->queryParameters = $queryParameters;
    }

    /**
     * Instantiates a new ChromosomesRequestBuilderGetQueryParameters.
     * @param string|null $genome 
     * @param string|null $hubUrl 
     * @param string|null $track 
     * @return ChromosomesRequestBuilderGetQueryParameters
    */
    public static function createQueryParameters(?string $genome = null, ?string $hubUrl = null, ?string $track = null): ChromosomesRequestBuilderGetQueryParameters {
        return new ChromosomesRequestBuilderGetQueryParameters($genome, $hubUrl, $track);
    }

}
