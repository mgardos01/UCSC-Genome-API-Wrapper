package list

import (
    "context"
    i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9 "apisdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// FilesRequestBuilder builds and executes requests for operations under \list\files
type FilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
type FilesRequestBuilderGetQueryParameters struct {
    // Deprecated: This property is deprecated, use FormatAsFormat instead
    Format *string `uriparametername:"format"`
    FormatAsFormat *i6de1841a269ca673200fabfa47c4e18acac3830338ac398493ce16e418186aa9.Format `uriparametername:"format"`
    Genome *string `uriparametername:"genome"`
    MaxItemsOutput *int32 `uriparametername:"maxItemsOutput"`
}
// FilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilesRequestBuilderGetQueryParameters
}
// NewFilesRequestBuilderInternal instantiates a new FilesRequestBuilder and sets the default values.
func NewFilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilesRequestBuilder) {
    m := &FilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/list/files?genome={genome}{&format*,maxItemsOutput*}", pathParameters),
    }
    return m
}
// NewFilesRequestBuilder instantiates a new FilesRequestBuilder and sets the default values.
func NewFilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilesRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a []byte when successful
func (m *FilesRequestBuilder) Get(ctx context.Context, requestConfiguration *FilesRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *FilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilesRequestBuilder when successful
func (m *FilesRequestBuilder) WithUrl(rawUrl string)(*FilesRequestBuilder) {
    return NewFilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
