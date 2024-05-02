package list

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UcscGenomesRequestBuilder builds and executes requests for operations under \list\ucscGenomes
type UcscGenomesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UcscGenomesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UcscGenomesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUcscGenomesRequestBuilderInternal instantiates a new UcscGenomesRequestBuilder and sets the default values.
func NewUcscGenomesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UcscGenomesRequestBuilder) {
    m := &UcscGenomesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/list/ucscGenomes", pathParameters),
    }
    return m
}
// NewUcscGenomesRequestBuilder instantiates a new UcscGenomesRequestBuilder and sets the default values.
func NewUcscGenomesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UcscGenomesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUcscGenomesRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a []byte when successful
func (m *UcscGenomesRequestBuilder) Get(ctx context.Context, requestConfiguration *UcscGenomesRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *UcscGenomesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UcscGenomesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UcscGenomesRequestBuilder when successful
func (m *UcscGenomesRequestBuilder) WithUrl(rawUrl string)(*UcscGenomesRequestBuilder) {
    return NewUcscGenomesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
