<?php

namespace ApiSdk\EscapedList\Tracks;

use Microsoft\Kiota\Abstractions\BaseRequestConfiguration;
use Microsoft\Kiota\Abstractions\RequestOption;

/**
 * Configuration for the request such as headers, query parameters, and middleware options.
*/
class TracksRequestBuilderGetRequestConfiguration extends BaseRequestConfiguration 
{
    /**
     * @var TracksRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public ?TracksRequestBuilderGetQueryParameters $queryParameters = null;
    
    /**
     * Instantiates a new TracksRequestBuilderGetRequestConfiguration and sets the default values.
     * @param array<string, array<string>|string>|null $headers Request headers
     * @param array<RequestOption>|null $options Request options
     * @param TracksRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public function __construct(?array $headers = null, ?array $options = null, ?TracksRequestBuilderGetQueryParameters $queryParameters = null) {
        parent::__construct($headers ?? [], $options ?? []);
        $this->queryParameters = $queryParameters;
    }

    /**
     * Instantiates a new TracksRequestBuilderGetQueryParameters.
     * @param string|null $genome 
     * @param string|null $hubUrl 
     * @param bool|null $trackLeavesOnly 
     * @return TracksRequestBuilderGetQueryParameters
    */
    public static function createQueryParameters(?string $genome = null, ?string $hubUrl = null, ?bool $trackLeavesOnly = null): TracksRequestBuilderGetQueryParameters {
        return new TracksRequestBuilderGetQueryParameters($genome, $hubUrl, $trackLeavesOnly);
    }

}
