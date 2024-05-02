package list

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TracksRequestBuilder builds and executes requests for operations under \list\tracks
type TracksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
type TracksRequestBuilderGetQueryParameters struct {
    Genome *string `uriparametername:"genome"`
    HubUrl *string `uriparametername:"hubUrl"`
    TrackLeavesOnly *bool `uriparametername:"trackLeavesOnly"`
}
// TracksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TracksRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TracksRequestBuilderGetQueryParameters
}
// NewTracksRequestBuilderInternal instantiates a new TracksRequestBuilder and sets the default values.
func NewTracksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TracksRequestBuilder) {
    m := &TracksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/list/tracks?genome={genome}{&hubUrl*,trackLeavesOnly*}", pathParameters),
    }
    return m
}
// NewTracksRequestBuilder instantiates a new TracksRequestBuilder and sets the default values.
func NewTracksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TracksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTracksRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a []byte when successful
func (m *TracksRequestBuilder) Get(ctx context.Context, requestConfiguration *TracksRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *TracksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TracksRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TracksRequestBuilder when successful
func (m *TracksRequestBuilder) WithUrl(rawUrl string)(*TracksRequestBuilder) {
    return NewTracksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
