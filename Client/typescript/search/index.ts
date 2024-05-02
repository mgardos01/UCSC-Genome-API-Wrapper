/* tslint:disable */
/* eslint-disable */
// Generated by Microsoft Kiota
// @ts-ignore
import { Category } from '../models/';
// @ts-ignore
import { type BaseRequestBuilder, type Parsable, type ParsableFactory, type RequestConfiguration, type RequestInformation, type RequestsMetadata } from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /search
 */
export interface SearchRequestBuilder extends BaseRequestBuilder<SearchRequestBuilder> {
    /**
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns {Promise<ArrayBuffer>}
     */
     get(requestConfiguration?: RequestConfiguration<SearchRequestBuilderGetQueryParameters> | undefined) : Promise<ArrayBuffer | undefined>;
    /**
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns {RequestInformation}
     */
     toGetRequestInformation(requestConfiguration?: RequestConfiguration<SearchRequestBuilderGetQueryParameters> | undefined) : RequestInformation;
}
export interface SearchRequestBuilderGetQueryParameters {
    categories?: Category;
    genome?: string;
    search?: string;
}
/**
 * Uri template for the request builder.
 */
export const SearchRequestBuilderUriTemplate = "{+baseurl}/search?genome={genome}&search={search}{&categories*}";
/**
 * Metadata for all the requests in the request builder.
 */
export const SearchRequestBuilderRequestsMetadata: RequestsMetadata = {
    get: {
        uriTemplate: SearchRequestBuilderUriTemplate,
        adapterMethodName: "sendPrimitive",
        responseBodyFactory:  "ArrayBuffer",
    },
};
/* tslint:enable */
/* eslint-enable */