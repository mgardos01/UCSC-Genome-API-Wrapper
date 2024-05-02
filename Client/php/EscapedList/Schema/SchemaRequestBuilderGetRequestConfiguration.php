<?php

namespace ApiSdk\EscapedList\Schema;

use Microsoft\Kiota\Abstractions\BaseRequestConfiguration;
use Microsoft\Kiota\Abstractions\RequestOption;

/**
 * Configuration for the request such as headers, query parameters, and middleware options.
*/
class SchemaRequestBuilderGetRequestConfiguration extends BaseRequestConfiguration 
{
    /**
     * @var SchemaRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public ?SchemaRequestBuilderGetQueryParameters $queryParameters = null;
    
    /**
     * Instantiates a new SchemaRequestBuilderGetRequestConfiguration and sets the default values.
     * @param array<string, array<string>|string>|null $headers Request headers
     * @param array<RequestOption>|null $options Request options
     * @param SchemaRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public function __construct(?array $headers = null, ?array $options = null, ?SchemaRequestBuilderGetQueryParameters $queryParameters = null) {
        parent::__construct($headers ?? [], $options ?? []);
        $this->queryParameters = $queryParameters;
    }

    /**
     * Instantiates a new SchemaRequestBuilderGetQueryParameters.
     * @param string|null $genome 
     * @param string|null $hubUrl 
     * @param string|null $track 
     * @return SchemaRequestBuilderGetQueryParameters
    */
    public static function createQueryParameters(?string $genome = null, ?string $hubUrl = null, ?string $track = null): SchemaRequestBuilderGetQueryParameters {
        return new SchemaRequestBuilderGetQueryParameters($genome, $hubUrl, $track);
    }

}
