<?php

namespace ApiSdk\GetData\Track;

use Microsoft\Kiota\Abstractions\BaseRequestConfiguration;
use Microsoft\Kiota\Abstractions\RequestOption;

/**
 * Configuration for the request such as headers, query parameters, and middleware options.
*/
class TrackRequestBuilderGetRequestConfiguration extends BaseRequestConfiguration 
{
    /**
     * @var TrackRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public ?TrackRequestBuilderGetQueryParameters $queryParameters = null;
    
    /**
     * Instantiates a new TrackRequestBuilderGetRequestConfiguration and sets the default values.
     * @param array<string, array<string>|string>|null $headers Request headers
     * @param array<RequestOption>|null $options Request options
     * @param TrackRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public function __construct(?array $headers = null, ?array $options = null, ?TrackRequestBuilderGetQueryParameters $queryParameters = null) {
        parent::__construct($headers ?? [], $options ?? []);
        $this->queryParameters = $queryParameters;
    }

    /**
     * Instantiates a new TrackRequestBuilderGetQueryParameters.
     * @param string|null $chrom 
     * @param int|null $end 
     * @param string|null $genome 
     * @param string|null $hubUrl 
     * @param bool|null $jsonOutputArrays 
     * @param int|null $maxItemsOutput 
     * @param int|null $start 
     * @param string|null $track 
     * @return TrackRequestBuilderGetQueryParameters
    */
    public static function createQueryParameters(?string $chrom = null, ?int $end = null, ?string $genome = null, ?string $hubUrl = null, ?bool $jsonOutputArrays = null, ?int $maxItemsOutput = null, ?int $start = null, ?string $track = null): TrackRequestBuilderGetQueryParameters {
        return new TrackRequestBuilderGetQueryParameters($chrom, $end, $genome, $hubUrl, $jsonOutputArrays, $maxItemsOutput, $start, $track);
    }

}
