require 'microsoft_kiota_abstractions'
require_relative '../api_sdk'
require_relative './get_data'
require_relative './sequence/sequence_request_builder'
require_relative './track/track_request_builder'

module ApiSdk
    module GetData
        ## 
        # Builds and executes requests for operations under \getData
        class GetDataRequestBuilder < MicrosoftKiotaAbstractions::BaseRequestBuilder
            
            ## 
            # The sequence property
            def sequence()
                return ApiSdk::GetData::Sequence::SequenceRequestBuilder.new(@path_parameters, @request_adapter)
            end
            ## 
            # The track property
            def track()
                return ApiSdk::GetData::Track::TrackRequestBuilder.new(@path_parameters, @request_adapter)
            end
            ## 
            ## Instantiates a new GetDataRequestBuilder and sets the default values.
            ## @param path_parameters Path parameters for the request
            ## @param request_adapter The request adapter to use to execute the requests.
            ## @return a void
            ## 
            def initialize(path_parameters, request_adapter)
                super(path_parameters, request_adapter, "{+baseurl}/getData")
            end
        end
    end
end
