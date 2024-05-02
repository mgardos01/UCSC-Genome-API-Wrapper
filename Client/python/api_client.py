from __future__ import annotations
from kiota_abstractions.api_client_builder import enable_backing_store_for_serialization_writer_factory, register_default_deserializer, register_default_serializer
from kiota_abstractions.base_request_builder import BaseRequestBuilder
from kiota_abstractions.get_path_parameters import get_path_parameters
from kiota_abstractions.request_adapter import RequestAdapter
from kiota_abstractions.serialization import ParseNodeFactoryRegistry, SerializationWriterFactoryRegistry
from kiota_serialization_form.form_parse_node_factory import FormParseNodeFactory
from kiota_serialization_form.form_serialization_writer_factory import FormSerializationWriterFactory
from kiota_serialization_json.json_parse_node_factory import JsonParseNodeFactory
from kiota_serialization_json.json_serialization_writer_factory import JsonSerializationWriterFactory
from kiota_serialization_multipart.multipart_serialization_writer_factory import MultipartSerializationWriterFactory
from kiota_serialization_text.text_parse_node_factory import TextParseNodeFactory
from kiota_serialization_text.text_serialization_writer_factory import TextSerializationWriterFactory
from typing import Any, Callable, Dict, List, Optional, TYPE_CHECKING, Union

if TYPE_CHECKING:
    from .get_data.get_data_request_builder import GetDataRequestBuilder
    from .list_.list_request_builder import ListRequestBuilder
    from .search.search_request_builder import SearchRequestBuilder

class ApiClient(BaseRequestBuilder):
    """
    The main entry point of the SDK, exposes the configuration and the fluent API.
    """
    def __init__(self,request_adapter: RequestAdapter) -> None:
        """
        Instantiates a new ApiClient and sets the default values.
        param request_adapter: The request adapter to use to execute the requests.
        Returns: None
        """
        if not request_adapter:
            raise TypeError("request_adapter cannot be null.")
        super().__init__(request_adapter, "{+baseurl}", None)
        register_default_serializer(JsonSerializationWriterFactory)
        register_default_serializer(TextSerializationWriterFactory)
        register_default_serializer(FormSerializationWriterFactory)
        register_default_serializer(MultipartSerializationWriterFactory)
        register_default_deserializer(JsonParseNodeFactory)
        register_default_deserializer(TextParseNodeFactory)
        register_default_deserializer(FormParseNodeFactory)
        if not self.request_adapter.base_url:
            self.request_adapter.base_url = "https://api.genome.ucsc.edu/list/publicHubs"
        self.path_parameters["base_url"] = self.request_adapter.base_url
    
    @property
    def get_data(self) -> GetDataRequestBuilder:
        """
        The getData property
        """
        from .get_data.get_data_request_builder import GetDataRequestBuilder

        return GetDataRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def list_(self) -> ListRequestBuilder:
        """
        The list property
        """
        from .list_.list_request_builder import ListRequestBuilder

        return ListRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def search(self) -> SearchRequestBuilder:
        """
        The search property
        """
        from .search.search_request_builder import SearchRequestBuilder

        return SearchRequestBuilder(self.request_adapter, self.path_parameters)
    

