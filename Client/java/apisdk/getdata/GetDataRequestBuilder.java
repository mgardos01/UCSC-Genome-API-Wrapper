package apisdk.getdata;

import apisdk.getdata.sequence.SequenceRequestBuilder;
import apisdk.getdata.track.TrackRequestBuilder;
import com.microsoft.kiota.BaseRequestBuilder;
import com.microsoft.kiota.RequestAdapter;
import java.util.HashMap;
import java.util.Objects;
/**
 * Builds and executes requests for operations under /getData
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public class GetDataRequestBuilder extends BaseRequestBuilder {
    /**
     * The sequence property
     * @return a {@link SequenceRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public SequenceRequestBuilder sequence() {
        return new SequenceRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * The track property
     * @return a {@link TrackRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public TrackRequestBuilder track() {
        return new TrackRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * Instantiates a new {@link GetDataRequestBuilder} and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public GetDataRequestBuilder(@jakarta.annotation.Nonnull final HashMap<String, Object> pathParameters, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/getData", pathParameters);
    }
    /**
     * Instantiates a new {@link GetDataRequestBuilder} and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public GetDataRequestBuilder(@jakarta.annotation.Nonnull final String rawUrl, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/getData", rawUrl);
    }
}
