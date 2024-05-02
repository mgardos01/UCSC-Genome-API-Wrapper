<?php

namespace ApiSdk\GetData\Sequence;

class SequenceRequestBuilderGetQueryParameters 
{
    /**
     * @var string|null $chrom 
    */
    public ?string $chrom = null;
    
    /**
     * @var int|null $end 
    */
    public ?int $end = null;
    
    /**
     * @var string|null $genome 
    */
    public ?string $genome = null;
    
    /**
     * @var string|null $hubUrl 
    */
    public ?string $hubUrl = null;
    
    /**
     * @var int|null $start 
    */
    public ?int $start = null;
    
    /**
     * Instantiates a new SequenceRequestBuilderGetQueryParameters and sets the default values.
     * @param string|null $chrom 
     * @param int|null $end 
     * @param string|null $genome 
     * @param string|null $hubUrl 
     * @param int|null $start 
    */
    public function __construct(?string $chrom = null, ?int $end = null, ?string $genome = null, ?string $hubUrl = null, ?int $start = null) {
        $this->chrom = $chrom;
        $this->end = $end;
        $this->genome = $genome;
        $this->hubUrl = $hubUrl;
        $this->start = $start;
    }

}
