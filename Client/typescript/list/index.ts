/* tslint:disable */
/* eslint-disable */
// Generated by Microsoft Kiota
// @ts-ignore
import { ChromosomesRequestBuilderRequestsMetadata, type ChromosomesRequestBuilder } from './chromosomes/';
// @ts-ignore
import { FilesRequestBuilderRequestsMetadata, type FilesRequestBuilder } from './files/';
// @ts-ignore
import { HubGenomesRequestBuilderRequestsMetadata, type HubGenomesRequestBuilder } from './hubGenomes/';
// @ts-ignore
import { PublicHubsRequestBuilderRequestsMetadata, type PublicHubsRequestBuilder } from './publicHubs/';
// @ts-ignore
import { SchemaRequestBuilderRequestsMetadata, type SchemaRequestBuilder } from './schema/';
// @ts-ignore
import { TracksRequestBuilderRequestsMetadata, type TracksRequestBuilder } from './tracks/';
// @ts-ignore
import { type UcscGenomesRequestBuilder, UcscGenomesRequestBuilderRequestsMetadata } from './ucscGenomes/';
// @ts-ignore
import { type BaseRequestBuilder, type KeysToExcludeForNavigationMetadata, type NavigationMetadata } from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /list
 */
export interface ListRequestBuilder extends BaseRequestBuilder<ListRequestBuilder> {
    /**
     * The chromosomes property
     */
    get chromosomes(): ChromosomesRequestBuilder;
    /**
     * The files property
     */
    get files(): FilesRequestBuilder;
    /**
     * The hubGenomes property
     */
    get hubGenomes(): HubGenomesRequestBuilder;
    /**
     * The publicHubs property
     */
    get publicHubs(): PublicHubsRequestBuilder;
    /**
     * The schema property
     */
    get schema(): SchemaRequestBuilder;
    /**
     * The tracks property
     */
    get tracks(): TracksRequestBuilder;
    /**
     * The ucscGenomes property
     */
    get ucscGenomes(): UcscGenomesRequestBuilder;
}
/**
 * Uri template for the request builder.
 */
export const ListRequestBuilderUriTemplate = "{+baseurl}/list";
/**
 * Metadata for all the navigation properties in the request builder.
 */
export const ListRequestBuilderNavigationMetadata: Record<Exclude<keyof ListRequestBuilder, KeysToExcludeForNavigationMetadata>, NavigationMetadata> = {
    chromosomes: {
        requestsMetadata: ChromosomesRequestBuilderRequestsMetadata,
    },
    files: {
        requestsMetadata: FilesRequestBuilderRequestsMetadata,
    },
    hubGenomes: {
        requestsMetadata: HubGenomesRequestBuilderRequestsMetadata,
    },
    publicHubs: {
        requestsMetadata: PublicHubsRequestBuilderRequestsMetadata,
    },
    schema: {
        requestsMetadata: SchemaRequestBuilderRequestsMetadata,
    },
    tracks: {
        requestsMetadata: TracksRequestBuilderRequestsMetadata,
    },
    ucscGenomes: {
        requestsMetadata: UcscGenomesRequestBuilderRequestsMetadata,
    },
};
/* tslint:enable */
/* eslint-enable */
