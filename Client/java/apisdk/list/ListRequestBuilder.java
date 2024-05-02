package apisdk.list;

import apisdk.list.chromosomes.ChromosomesRequestBuilder;
import apisdk.list.files.FilesRequestBuilder;
import apisdk.list.hubgenomes.HubGenomesRequestBuilder;
import apisdk.list.publichubs.PublicHubsRequestBuilder;
import apisdk.list.schema.SchemaRequestBuilder;
import apisdk.list.tracks.TracksRequestBuilder;
import apisdk.list.ucscgenomes.UcscGenomesRequestBuilder;
import com.microsoft.kiota.BaseRequestBuilder;
import com.microsoft.kiota.RequestAdapter;
import java.util.HashMap;
import java.util.Objects;
/**
 * Builds and executes requests for operations under /list
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public class ListRequestBuilder extends BaseRequestBuilder {
    /**
     * The chromosomes property
     * @return a {@link ChromosomesRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public ChromosomesRequestBuilder chromosomes() {
        return new ChromosomesRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * The files property
     * @return a {@link FilesRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public FilesRequestBuilder files() {
        return new FilesRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * The hubGenomes property
     * @return a {@link HubGenomesRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public HubGenomesRequestBuilder hubGenomes() {
        return new HubGenomesRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * The publicHubs property
     * @return a {@link PublicHubsRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public PublicHubsRequestBuilder publicHubs() {
        return new PublicHubsRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * The schema property
     * @return a {@link SchemaRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public SchemaRequestBuilder schema() {
        return new SchemaRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * The tracks property
     * @return a {@link TracksRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public TracksRequestBuilder tracks() {
        return new TracksRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * The ucscGenomes property
     * @return a {@link UcscGenomesRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public UcscGenomesRequestBuilder ucscGenomes() {
        return new UcscGenomesRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * Instantiates a new {@link ListRequestBuilder} and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public ListRequestBuilder(@jakarta.annotation.Nonnull final HashMap<String, Object> pathParameters, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/list", pathParameters);
    }
    /**
     * Instantiates a new {@link ListRequestBuilder} and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public ListRequestBuilder(@jakarta.annotation.Nonnull final String rawUrl, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/list", rawUrl);
    }
}
