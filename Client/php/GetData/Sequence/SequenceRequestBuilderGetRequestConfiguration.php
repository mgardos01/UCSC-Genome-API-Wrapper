<?php

namespace ApiSdk\GetData\Sequence;

use Microsoft\Kiota\Abstractions\BaseRequestConfiguration;
use Microsoft\Kiota\Abstractions\RequestOption;

/**
 * Configuration for the request such as headers, query parameters, and middleware options.
*/
class SequenceRequestBuilderGetRequestConfiguration extends BaseRequestConfiguration 
{
    /**
     * @var SequenceRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public ?SequenceRequestBuilderGetQueryParameters $queryParameters = null;
    
    /**
     * Instantiates a new SequenceRequestBuilderGetRequestConfiguration and sets the default values.
     * @param array<string, array<string>|string>|null $headers Request headers
     * @param array<RequestOption>|null $options Request options
     * @param SequenceRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public function __construct(?array $headers = null, ?array $options = null, ?SequenceRequestBuilderGetQueryParameters $queryParameters = null) {
        parent::__construct($headers ?? [], $options ?? []);
        $this->queryParameters = $queryParameters;
    }

    /**
     * Instantiates a new SequenceRequestBuilderGetQueryParameters.
     * @param string|null $chrom 
     * @param int|null $end 
     * @param string|null $genome 
     * @param string|null $hubUrl 
     * @param int|null $start 
     * @return SequenceRequestBuilderGetQueryParameters
    */
    public static function createQueryParameters(?string $chrom = null, ?int $end = null, ?string $genome = null, ?string $hubUrl = null, ?int $start = null): SequenceRequestBuilderGetQueryParameters {
        return new SequenceRequestBuilderGetQueryParameters($chrom, $end, $genome, $hubUrl, $start);
    }

}
