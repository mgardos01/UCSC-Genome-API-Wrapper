package getdata

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TrackRequestBuilder builds and executes requests for operations under \getData\track
type TrackRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
type TrackRequestBuilderGetQueryParameters struct {
    Chrom *string `uriparametername:"chrom"`
    End *int64 `uriparametername:"end"`
    Genome *string `uriparametername:"genome"`
    HubUrl *string `uriparametername:"hubUrl"`
    JsonOutputArrays *bool `uriparametername:"jsonOutputArrays"`
    MaxItemsOutput *int32 `uriparametername:"maxItemsOutput"`
    Start *int64 `uriparametername:"start"`
    Track *string `uriparametername:"track"`
}
// TrackRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TrackRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TrackRequestBuilderGetQueryParameters
}
// NewTrackRequestBuilderInternal instantiates a new TrackRequestBuilder and sets the default values.
func NewTrackRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TrackRequestBuilder) {
    m := &TrackRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/getData/track?genome={genome}&track={track}{&chrom*,end*,hubUrl*,jsonOutputArrays*,maxItemsOutput*,start*}", pathParameters),
    }
    return m
}
// NewTrackRequestBuilder instantiates a new TrackRequestBuilder and sets the default values.
func NewTrackRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TrackRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTrackRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a []byte when successful
func (m *TrackRequestBuilder) Get(ctx context.Context, requestConfiguration *TrackRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// returns a *RequestInformation when successful
func (m *TrackRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TrackRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TrackRequestBuilder when successful
func (m *TrackRequestBuilder) WithUrl(rawUrl string)(*TrackRequestBuilder) {
    return NewTrackRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
