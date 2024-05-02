/* tslint:disable */
/* eslint-disable */
// Generated by Microsoft Kiota
// @ts-ignore
import { type BaseRequestBuilder, type Parsable, type ParsableFactory, type RequestConfiguration, type RequestInformation, type RequestsMetadata } from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /list/hubGenomes
 */
export interface HubGenomesRequestBuilder extends BaseRequestBuilder<HubGenomesRequestBuilder> {
    /**
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns {Promise<ArrayBuffer>}
     */
     get(requestConfiguration?: RequestConfiguration<HubGenomesRequestBuilderGetQueryParameters> | undefined) : Promise<ArrayBuffer | undefined>;
    /**
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns {RequestInformation}
     */
     toGetRequestInformation(requestConfiguration?: RequestConfiguration<HubGenomesRequestBuilderGetQueryParameters> | undefined) : RequestInformation;
}
export interface HubGenomesRequestBuilderGetQueryParameters {
    hubUrl?: string;
}
/**
 * Uri template for the request builder.
 */
export const HubGenomesRequestBuilderUriTemplate = "{+baseurl}/list/hubGenomes?hubUrl={hubUrl}";
/**
 * Metadata for all the requests in the request builder.
 */
export const HubGenomesRequestBuilderRequestsMetadata: RequestsMetadata = {
    get: {
        uriTemplate: HubGenomesRequestBuilderUriTemplate,
        adapterMethodName: "sendPrimitive",
        responseBodyFactory:  "ArrayBuffer",
    },
};
/* tslint:enable */
/* eslint-enable */