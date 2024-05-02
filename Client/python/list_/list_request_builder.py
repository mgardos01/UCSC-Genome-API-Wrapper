from __future__ import annotations
from kiota_abstractions.base_request_builder import BaseRequestBuilder
from kiota_abstractions.get_path_parameters import get_path_parameters
from kiota_abstractions.request_adapter import RequestAdapter
from typing import Any, Callable, Dict, List, Optional, TYPE_CHECKING, Union

if TYPE_CHECKING:
    from .chromosomes.chromosomes_request_builder import ChromosomesRequestBuilder
    from .files.files_request_builder import FilesRequestBuilder
    from .hub_genomes.hub_genomes_request_builder import HubGenomesRequestBuilder
    from .public_hubs.public_hubs_request_builder import PublicHubsRequestBuilder
    from .schema.schema_request_builder import SchemaRequestBuilder
    from .tracks.tracks_request_builder import TracksRequestBuilder
    from .ucsc_genomes.ucsc_genomes_request_builder import UcscGenomesRequestBuilder

class ListRequestBuilder(BaseRequestBuilder):
    """
    Builds and executes requests for operations under /list
    """
    def __init__(self,request_adapter: RequestAdapter, path_parameters: Union[str, Dict[str, Any]]) -> None:
        """
        Instantiates a new ListRequestBuilder and sets the default values.
        param path_parameters: The raw url or the url-template parameters for the request.
        param request_adapter: The request adapter to use to execute the requests.
        Returns: None
        """
        super().__init__(request_adapter, "{+baseurl}/list", path_parameters)
    
    @property
    def chromosomes(self) -> ChromosomesRequestBuilder:
        """
        The chromosomes property
        """
        from .chromosomes.chromosomes_request_builder import ChromosomesRequestBuilder

        return ChromosomesRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def files(self) -> FilesRequestBuilder:
        """
        The files property
        """
        from .files.files_request_builder import FilesRequestBuilder

        return FilesRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def hub_genomes(self) -> HubGenomesRequestBuilder:
        """
        The hubGenomes property
        """
        from .hub_genomes.hub_genomes_request_builder import HubGenomesRequestBuilder

        return HubGenomesRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def public_hubs(self) -> PublicHubsRequestBuilder:
        """
        The publicHubs property
        """
        from .public_hubs.public_hubs_request_builder import PublicHubsRequestBuilder

        return PublicHubsRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def schema(self) -> SchemaRequestBuilder:
        """
        The schema property
        """
        from .schema.schema_request_builder import SchemaRequestBuilder

        return SchemaRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def tracks(self) -> TracksRequestBuilder:
        """
        The tracks property
        """
        from .tracks.tracks_request_builder import TracksRequestBuilder

        return TracksRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def ucsc_genomes(self) -> UcscGenomesRequestBuilder:
        """
        The ucscGenomes property
        """
        from .ucsc_genomes.ucsc_genomes_request_builder import UcscGenomesRequestBuilder

        return UcscGenomesRequestBuilder(self.request_adapter, self.path_parameters)
    

