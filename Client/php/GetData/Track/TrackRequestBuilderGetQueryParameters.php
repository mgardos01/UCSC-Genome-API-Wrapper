<?php

namespace ApiSdk\GetData\Track;

class TrackRequestBuilderGetQueryParameters 
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
     * @var bool|null $jsonOutputArrays 
    */
    public ?bool $jsonOutputArrays = null;
    
    /**
     * @var int|null $maxItemsOutput 
    */
    public ?int $maxItemsOutput = null;
    
    /**
     * @var int|null $start 
    */
    public ?int $start = null;
    
    /**
     * @var string|null $track 
    */
    public ?string $track = null;
    
    /**
     * Instantiates a new TrackRequestBuilderGetQueryParameters and sets the default values.
     * @param string|null $chrom 
     * @param int|null $end 
     * @param string|null $genome 
     * @param string|null $hubUrl 
     * @param bool|null $jsonOutputArrays 
     * @param int|null $maxItemsOutput 
     * @param int|null $start 
     * @param string|null $track 
    */
    public function __construct(?string $chrom = null, ?int $end = null, ?string $genome = null, ?string $hubUrl = null, ?bool $jsonOutputArrays = null, ?int $maxItemsOutput = null, ?int $start = null, ?string $track = null) {
        $this->chrom = $chrom;
        $this->end = $end;
        $this->genome = $genome;
        $this->hubUrl = $hubUrl;
        $this->jsonOutputArrays = $jsonOutputArrays;
        $this->maxItemsOutput = $maxItemsOutput;
        $this->start = $start;
        $this->track = $track;
    }

}
