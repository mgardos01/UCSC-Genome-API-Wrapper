require 'microsoft_kiota_abstractions'
require_relative '../api_sdk'
require_relative './chromosomes/chromosomes_request_builder'
require_relative './files/files_request_builder'
require_relative './hub_genomes/hub_genomes_request_builder'
require_relative './list'
require_relative './public_hubs/public_hubs_request_builder'
require_relative './schema/schema_request_builder'
require_relative './tracks/tracks_request_builder'
require_relative './ucsc_genomes/ucsc_genomes_request_builder'

module ApiSdk
    module List
        ## 
        # Builds and executes requests for operations under \list
        class ListRequestBuilder < MicrosoftKiotaAbstractions::BaseRequestBuilder
            
            ## 
            # The chromosomes property
            def chromosomes()
                return ApiSdk::List::Chromosomes::ChromosomesRequestBuilder.new(@path_parameters, @request_adapter)
            end
            ## 
            # The files property
            def files()
                return ApiSdk::List::Files::FilesRequestBuilder.new(@path_parameters, @request_adapter)
            end
            ## 
            # The hubGenomes property
            def hub_genomes()
                return ApiSdk::List::HubGenomes::HubGenomesRequestBuilder.new(@path_parameters, @request_adapter)
            end
            ## 
            # The publicHubs property
            def public_hubs()
                return ApiSdk::List::PublicHubs::PublicHubsRequestBuilder.new(@path_parameters, @request_adapter)
            end
            ## 
            # The schema property
            def schema()
                return ApiSdk::List::Schema::SchemaRequestBuilder.new(@path_parameters, @request_adapter)
            end
            ## 
            # The tracks property
            def tracks()
                return ApiSdk::List::Tracks::TracksRequestBuilder.new(@path_parameters, @request_adapter)
            end
            ## 
            # The ucscGenomes property
            def ucsc_genomes()
                return ApiSdk::List::UcscGenomes::UcscGenomesRequestBuilder.new(@path_parameters, @request_adapter)
            end
            ## 
            ## Instantiates a new ListRequestBuilder and sets the default values.
            ## @param path_parameters Path parameters for the request
            ## @param request_adapter The request adapter to use to execute the requests.
            ## @return a void
            ## 
            def initialize(path_parameters, request_adapter)
                super(path_parameters, request_adapter, "{+baseurl}/list")
            end
        end
    end
end
