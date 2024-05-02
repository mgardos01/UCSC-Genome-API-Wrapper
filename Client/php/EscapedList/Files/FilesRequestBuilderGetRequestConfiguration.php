<?php

namespace ApiSdk\EscapedList\Files;

use ApiSdk\Models\Format;
use Microsoft\Kiota\Abstractions\BaseRequestConfiguration;
use Microsoft\Kiota\Abstractions\RequestOption;

/**
 * Configuration for the request such as headers, query parameters, and middleware options.
*/
class FilesRequestBuilderGetRequestConfiguration extends BaseRequestConfiguration 
{
    /**
     * @var FilesRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public ?FilesRequestBuilderGetQueryParameters $queryParameters = null;
    
    /**
     * Instantiates a new FilesRequestBuilderGetRequestConfiguration and sets the default values.
     * @param array<string, array<string>|string>|null $headers Request headers
     * @param array<RequestOption>|null $options Request options
     * @param FilesRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public function __construct(?array $headers = null, ?array $options = null, ?FilesRequestBuilderGetQueryParameters $queryParameters = null) {
        parent::__construct($headers ?? [], $options ?? []);
        $this->queryParameters = $queryParameters;
    }

    /**
     * Instantiates a new FilesRequestBuilderGetQueryParameters.
     * @param Format|null $format 
     * @param string|null $genome 
     * @param int|null $maxItemsOutput 
     * @return FilesRequestBuilderGetQueryParameters
    */
    public static function createQueryParameters(?Format $format = null, ?string $genome = null, ?int $maxItemsOutput = null): FilesRequestBuilderGetQueryParameters {
        return new FilesRequestBuilderGetQueryParameters($format, $genome, $maxItemsOutput);
    }

}
