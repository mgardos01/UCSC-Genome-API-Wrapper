/* tslint:disable */
/* eslint-disable */
// Generated by Microsoft Kiota
// @ts-ignore
import { type BaseRequestBuilder, type Parsable, type ParsableFactory, type RequestConfiguration, type RequestInformation, type RequestsMetadata } from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /list/tracks
 */
export interface TracksRequestBuilder extends BaseRequestBuilder<TracksRequestBuilder> {
    /**
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns {Promise<ArrayBuffer>}
     */
     get(requestConfiguration?: RequestConfiguration<TracksRequestBuilderGetQueryParameters> | undefined) : Promise<ArrayBuffer | undefined>;
    /**
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns {RequestInformation}
     */
     toGetRequestInformation(requestConfiguration?: RequestConfiguration<TracksRequestBuilderGetQueryParameters> | undefined) : RequestInformation;
}
export interface TracksRequestBuilderGetQueryParameters {
    genome?: string;
    hubUrl?: string;
    trackLeavesOnly?: boolean;
}
/**
 * Uri template for the request builder.
 */
export const TracksRequestBuilderUriTemplate = "{+baseurl}/list/tracks?genome={genome}{&hubUrl*,trackLeavesOnly*}";
/**
 * Metadata for all the requests in the request builder.
 */
export const TracksRequestBuilderRequestsMetadata: RequestsMetadata = {
    get: {
        uriTemplate: TracksRequestBuilderUriTemplate,
        adapterMethodName: "sendPrimitive",
        responseBodyFactory:  "ArrayBuffer",
    },
};
/* tslint:enable */
/* eslint-enable */