<?php

namespace ApiSdk\EscapedList\Files;

use ApiSdk\Models\Format;

class FilesRequestBuilderGetQueryParameters 
{
    /**
     * @var Format|null $format 
    */
    public ?Format $format = null;
    
    /**
     * @var string|null $genome 
    */
    public ?string $genome = null;
    
    /**
     * @var int|null $maxItemsOutput 
    */
    public ?int $maxItemsOutput = null;
    
    /**
     * Instantiates a new FilesRequestBuilderGetQueryParameters and sets the default values.
     * @param Format|null $format 
     * @param string|null $genome 
     * @param int|null $maxItemsOutput 
    */
    public function __construct(?Format $format = null, ?string $genome = null, ?int $maxItemsOutput = null) {
        $this->format = $format;
        $this->genome = $genome;
        $this->maxItemsOutput = $maxItemsOutput;
    }

}
