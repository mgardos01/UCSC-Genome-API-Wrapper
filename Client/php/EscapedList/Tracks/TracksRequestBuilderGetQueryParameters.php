<?php

namespace ApiSdk\EscapedList\Tracks;

class TracksRequestBuilderGetQueryParameters 
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
     * @var bool|null $trackLeavesOnly 
    */
    public ?bool $trackLeavesOnly = null;
    
    /**
     * Instantiates a new TracksRequestBuilderGetQueryParameters and sets the default values.
     * @param string|null $genome 
     * @param string|null $hubUrl 
     * @param bool|null $trackLeavesOnly 
    */
    public function __construct(?string $genome = null, ?string $hubUrl = null, ?bool $trackLeavesOnly = null) {
        $this->genome = $genome;
        $this->hubUrl = $hubUrl;
        $this->trackLeavesOnly = $trackLeavesOnly;
    }

}
