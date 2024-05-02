package list

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ListRequestBuilder builds and executes requests for operations under \list
type ListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Chromosomes the chromosomes property
// returns a *ChromosomesRequestBuilder when successful
func (m *ListRequestBuilder) Chromosomes()(*ChromosomesRequestBuilder) {
    return NewChromosomesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewListRequestBuilderInternal instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/list", pathParameters),
    }
    return m
}
// NewListRequestBuilder instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListRequestBuilderInternal(urlParams, requestAdapter)
}
// Files the files property
// returns a *FilesRequestBuilder when successful
func (m *ListRequestBuilder) Files()(*FilesRequestBuilder) {
    return NewFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HubGenomes the hubGenomes property
// returns a *HubGenomesRequestBuilder when successful
func (m *ListRequestBuilder) HubGenomes()(*HubGenomesRequestBuilder) {
    return NewHubGenomesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PublicHubs the publicHubs property
// returns a *PublicHubsRequestBuilder when successful
func (m *ListRequestBuilder) PublicHubs()(*PublicHubsRequestBuilder) {
    return NewPublicHubsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Schema the schema property
// returns a *SchemaRequestBuilder when successful
func (m *ListRequestBuilder) Schema()(*SchemaRequestBuilder) {
    return NewSchemaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tracks the tracks property
// returns a *TracksRequestBuilder when successful
func (m *ListRequestBuilder) Tracks()(*TracksRequestBuilder) {
    return NewTracksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UcscGenomes the ucscGenomes property
// returns a *UcscGenomesRequestBuilder when successful
func (m *ListRequestBuilder) UcscGenomes()(*UcscGenomesRequestBuilder) {
    return NewUcscGenomesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
