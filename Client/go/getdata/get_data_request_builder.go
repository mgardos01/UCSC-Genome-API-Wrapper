package getdata

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetDataRequestBuilder builds and executes requests for operations under \getData
type GetDataRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewGetDataRequestBuilderInternal instantiates a new GetDataRequestBuilder and sets the default values.
func NewGetDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetDataRequestBuilder) {
    m := &GetDataRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/getData", pathParameters),
    }
    return m
}
// NewGetDataRequestBuilder instantiates a new GetDataRequestBuilder and sets the default values.
func NewGetDataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetDataRequestBuilderInternal(urlParams, requestAdapter)
}
// Sequence the sequence property
// returns a *SequenceRequestBuilder when successful
func (m *GetDataRequestBuilder) Sequence()(*SequenceRequestBuilder) {
    return NewSequenceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Track the track property
// returns a *TrackRequestBuilder when successful
func (m *GetDataRequestBuilder) Track()(*TrackRequestBuilder) {
    return NewTrackRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
