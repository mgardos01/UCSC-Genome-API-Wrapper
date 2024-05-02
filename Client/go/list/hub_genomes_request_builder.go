package list

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// HubGenomesRequestBuilder builds and executes requests for operations under \list\hubGenomes
type HubGenomesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
type HubGenomesRequestBuilderGetQueryParameters struct {
    HubUrl *string `uriparametername:"hubUrl"`
}
// HubGenomesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HubGenomesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *HubGenomesRequestBuilderGetQueryParameters
}
// NewHubGenomesRequestBuilderInternal instantiates a new HubGenomesRequestBuilder and sets the default values.
func NewHubGenomesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HubGenomesRequestBuilder) {
    m := &HubGenomesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/list/hubGenomes?hubUrl={hubUrl}", pathParameters),
    }
    return m
}
// NewHubGenomesRequestBuilder instantiates a new HubGenomesRequestBuilder and sets the default values.
func NewHubGenomesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HubGenomesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHubGenomesRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a []byte when successful
func (m *HubGenomesRequestBuilder) Get(ctx context.Context, requestConfiguration *HubGenomesRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *HubGenomesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HubGenomesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *HubGenomesRequestBuilder when successful
func (m *HubGenomesRequestBuilder) WithUrl(rawUrl string)(*HubGenomesRequestBuilder) {
    return NewHubGenomesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
