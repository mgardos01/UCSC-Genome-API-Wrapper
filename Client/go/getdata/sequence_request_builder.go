package getdata

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SequenceRequestBuilder builds and executes requests for operations under \getData\sequence
type SequenceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
type SequenceRequestBuilderGetQueryParameters struct {
    Chrom *string `uriparametername:"chrom"`
    End *int64 `uriparametername:"end"`
    Genome *string `uriparametername:"genome"`
    HubUrl *string `uriparametername:"hubUrl"`
    Start *int64 `uriparametername:"start"`
}
// SequenceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SequenceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SequenceRequestBuilderGetQueryParameters
}
// NewSequenceRequestBuilderInternal instantiates a new SequenceRequestBuilder and sets the default values.
func NewSequenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SequenceRequestBuilder) {
    m := &SequenceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/getData/sequence?chrom={chrom}&genome={genome}{&end*,hubUrl*,start*}", pathParameters),
    }
    return m
}
// NewSequenceRequestBuilder instantiates a new SequenceRequestBuilder and sets the default values.
func NewSequenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SequenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSequenceRequestBuilderInternal(urlParams, requestAdapter)
}
// returns a []byte when successful
func (m *SequenceRequestBuilder) Get(ctx context.Context, requestConfiguration *SequenceRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *SequenceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SequenceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SequenceRequestBuilder when successful
func (m *SequenceRequestBuilder) WithUrl(rawUrl string)(*SequenceRequestBuilder) {
    return NewSequenceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
