<?php

namespace ApiSdk\Search;

use ApiSdk\Models\Category;

class SearchRequestBuilderGetQueryParameters 
{
    /**
     * @var Category|null $categories 
    */
    public ?Category $categories = null;
    
    /**
     * @var string|null $genome 
    */
    public ?string $genome = null;
    
    /**
     * @var string|null $search 
    */
    public ?string $search = null;
    
    /**
     * Instantiates a new SearchRequestBuilderGetQueryParameters and sets the default values.
     * @param Category|null $categories 
     * @param string|null $genome 
     * @param string|null $search 
    */
    public function __construct(?Category $categories = null, ?string $genome = null, ?string $search = null) {
        $this->categories = $categories;
        $this->genome = $genome;
        $this->search = $search;
    }

}
