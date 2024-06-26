/* tslint:disable */
/* eslint-disable */
// Generated by Microsoft Kiota
// @ts-ignore
import { Format } from '../../models/';
// @ts-ignore
import { type BaseRequestBuilder, type Parsable, type ParsableFactory, type RequestConfiguration, type RequestInformation, type RequestsMetadata } from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /list/files
 */
export interface FilesRequestBuilder extends BaseRequestBuilder<FilesRequestBuilder> {
    /**
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns {Promise<ArrayBuffer>}
     */
     get(requestConfiguration?: RequestConfiguration<FilesRequestBuilderGetQueryParameters> | undefined) : Promise<ArrayBuffer | undefined>;
    /**
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns {RequestInformation}
     */
     toGetRequestInformation(requestConfiguration?: RequestConfiguration<FilesRequestBuilderGetQueryParameters> | undefined) : RequestInformation;
}
export interface FilesRequestBuilderGetQueryParameters {
    format?: Format;
    genome?: string;
    maxItemsOutput?: number;
}
/**
 * Uri template for the request builder.
 */
export const FilesRequestBuilderUriTemplate = "{+baseurl}/list/files?genome={genome}{&format*,maxItemsOutput*}";
/**
 * Metadata for all the requests in the request builder.
 */
export const FilesRequestBuilderRequestsMetadata: RequestsMetadata = {
    get: {
        uriTemplate: FilesRequestBuilderUriTemplate,
        adapterMethodName: "sendPrimitive",
        responseBodyFactory:  "ArrayBuffer",
    },
};
/* tslint:enable */
/* eslint-enable */
