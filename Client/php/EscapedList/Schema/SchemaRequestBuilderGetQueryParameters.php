<?php

namespace ApiSdk\EscapedList\Schema;

class SchemaRequestBuilderGetQueryParameters 
{
    /**
     * @var string|null $genome 
    */
    public ?string $genome = null;
    
    /**
     * @var string|null $hubUrl 
    */
    public ?string $hubUrl = null;
    
    /**
     * @var string|null $track 
    */
    public ?string $track = null;
    
    /**
     * Instantiates a new SchemaRequestBuilderGetQueryParameters and sets the default values.
     * @param string|null $genome 
     * @param string|null $hubUrl 
     * @param string|null $track 
    */
    public function __construct(?string $genome = null, ?string $hubUrl = null, ?string $track = null) {
        $this->genome = $genome;
        $this->hubUrl = $hubUrl;
        $this->track = $track;
    }

}
