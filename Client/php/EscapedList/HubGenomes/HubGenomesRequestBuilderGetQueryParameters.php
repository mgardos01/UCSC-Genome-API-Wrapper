<?php

namespace ApiSdk\EscapedList\HubGenomes;

class HubGenomesRequestBuilderGetQueryParameters 
{
    /**
     * @var string|null $hubUrl 
    */
    public ?string $hubUrl = null;
    
    /**
     * Instantiates a new HubGenomesRequestBuilderGetQueryParameters and sets the default values.
     * @param string|null $hubUrl 
    */
    public function __construct(?string $hubUrl = null) {
        $this->hubUrl = $hubUrl;
    }

}
