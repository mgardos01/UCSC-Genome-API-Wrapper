require 'microsoft_kiota_abstractions'
require_relative '../../api_sdk'
require_relative '../get_data'
require_relative './track'

module ApiSdk
    module GetData
        module Track
            ## 
            # Builds and executes requests for operations under \getData\track
            class TrackRequestBuilder < MicrosoftKiotaAbstractions::BaseRequestBuilder
                
                ## 
                ## Instantiates a new TrackRequestBuilder and sets the default values.
                ## @param path_parameters Path parameters for the request
                ## @param request_adapter The request adapter to use to execute the requests.
                ## @return a void
                ## 
                def initialize(path_parameters, request_adapter)
                    super(path_parameters, request_adapter, "{+baseurl}/getData/track?genome={genome}&track={track}{&chrom*,end*,hubUrl*,jsonOutputArrays*,maxItemsOutput*,start*}")
                end
                ## 
                ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                ## @return a Fiber of binary
                ## 
                def get(request_configuration=nil)
                    request_info = self.to_get_request_information(
                        request_configuration
                    )
                    return @request_adapter.send_async(request_info, Binary, nil)
                end
                ## 
                ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                ## @return a request_information
                ## 
                def to_get_request_information(request_configuration=nil)
                    request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
                    unless request_configuration.nil?
                        request_info.add_headers_from_raw_object(request_configuration.headers)
                        request_info.set_query_string_parameters_from_raw_object(request_configuration.query_parameters)
                        request_info.add_request_options(request_configuration.options)
                    end
                    request_info.url_template = @url_template
                    request_info.path_parameters = @path_parameters
                    request_info.http_method = :GET
                    return request_info
                end
                ## 
                ## Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
                ## @param raw_url The raw URL to use for the request builder.
                ## @return a track_request_builder
                ## 
                def with_url(raw_url)
                    raise StandardError, 'raw_url cannot be null' if raw_url.nil?
                    return TrackRequestBuilder.new(raw_url, @request_adapter)
                end

                class TrackRequestBuilderGetQueryParameters
                    
                    attr_accessor :chrom
                    attr_accessor :end_escaped
                    attr_accessor :genome
                    attr_accessor :hub_url
                    attr_accessor :json_output_arrays
                    attr_accessor :max_items_output
                    attr_accessor :start
                    attr_accessor :track
                    ## 
                    ## Maps the query parameters names to their encoded names for the URI template parsing.
                    ## @param original_name The original query parameter name in the class.
                    ## @return a string
                    ## 
                    def get_query_parameter(original_name)
                        raise StandardError, 'original_name cannot be null' if original_name.nil?
                        case original_name
                            when "hub_url"
                                return "hubUrl"
                            when "json_output_arrays"
                                return "jsonOutputArrays"
                            when "max_items_output"
                                return "maxItemsOutput"
                            else
                                return original_name
                        end
                    end
                end
            end
        end
    end
end
