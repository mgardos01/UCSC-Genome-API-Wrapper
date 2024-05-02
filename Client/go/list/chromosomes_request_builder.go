package list

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ChromosomesRequestBuilder builds and executes requests for operations under \list\chromosomes
type ChromosomesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
type ChromosomesRequestBuilderGetQueryParameters struct {
    Genome *string `uriparametername:"genome"`
    HubUrl *string `uriparametername:"hubUrl"`
    Track *string `uriparametername:"track"`
}
// ChromosomesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChromosomesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChromosomesRequestBuilderGetQueryParameters
}
// NewChromosomesRequestBuilderInternal instantiates a new ChromosomesRequestBuilder and sets the default values.
func NewChromosomesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChromosomesRequestBuilder) {
    m := &ChromosomesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/list/chromosomes?genome={genome}{&hubUrl*,track*}", pathParameters),
    }
    return m
}
// NewChromosomesRequestBuilder instantiates a new ChromosomesRequestBuilder and sets the default values.
func NewChromosomesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChromosomesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChromosomesRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a []byte when successful
func (m *ChromosomesRequestBuilder) Get(ctx context.Context, requestConfiguration *ChromosomesRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *ChromosomesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ChromosomesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ChromosomesRequestBuilder when successful
func (m *ChromosomesRequestBuilder) WithUrl(rawUrl string)(*ChromosomesRequestBuilder) {
    return NewChromosomesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
