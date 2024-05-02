require 'microsoft_kiota_abstractions'
require 'microsoft_kiota_serialization_json'
require_relative './api_sdk'
require_relative './get_data/get_data_request_builder'
require_relative './list/list_request_builder'
require_relative './search/search_request_builder'

module ApiSdk
    ## 
    # The main entry point of the SDK, exposes the configuration and the fluent API.
    class ApiClient < MicrosoftKiotaAbstractions::BaseRequestBuilder
        
        ## 
        # The getData property
        def get_data()
            return ApiSdk::GetData::GetDataRequestBuilder.new(@path_parameters, @request_adapter)
        end
        ## 
        # The list property
        def list()
            return ApiSdk::List::ListRequestBuilder.new(@path_parameters, @request_adapter)
        end
        ## 
        # The search property
        def search()
            return ApiSdk::Search::SearchRequestBuilder.new(@path_parameters, @request_adapter)
        end
        ## 
        ## Instantiates a new ApiClient and sets the default values.
        ## @param request_adapter The request adapter to use to execute the requests.
        ## @return a void
        ## 
        def initialize(request_adapter)
            super(Hash.new, request_adapter, "{+baseurl}")
            MicrosoftKiotaAbstractions::ApiClientBuilder.register_default_serializer(MicrosoftKiotaSerializationJson::JsonSerializationWriterFactory)
            MicrosoftKiotaAbstractions::ApiClientBuilder.register_default_deserializer(MicrosoftKiotaSerializationJson::JsonParseNodeFactory)
            if @request_adapter.get_base_url.nil? || @request_adapter.get_base_url.empty?
                @request_adapter.set_base_url('https://api.genome.ucsc.edu/list/publicHubs')
            end
            @path_parameters['baseurl'] = @request_adapter.get_base_url
        end
    end
end
