using System.Diagnostics;
using System.Text.Json;
using System.Web;
using GenomeDataApi.Models;
using GenomeDataApi.Utilities;
using Microsoft.AspNetCore.Http.HttpResults;
using Microsoft.AspNetCore.Mvc;

namespace GenomeDataApi.Endpoints;

public static class SearchEndpoints
{
    public static void MapEndpoints(WebApplication app)
    {
        var search = app.MapGroup("/search")
            .WithTags("search");

        // Define mappings inside here
        search.MapGet("/", async (
            [FromQuery(Name = "search")] string term, // TODO Make serializable else later
            [FromQuery(Name = "genome")] string name, // TODO Make serializable else later
            [FromQuery(Name = "categories")] Category? category,
            IHttpClientFactory httpClientFactory) =>
        {
            var client = httpClientFactory.CreateClient("genomeClient");

            // Create the query string directly
            var query = HttpUtility.ParseQueryString(string.Empty);

            query.SetIfNotNull("search", term);
            query.SetIfNotNull("genome", name);
            query.SetIfNotNull("categories", category);

            var response = await client.GetAsync($"search?{query}");
            
            // var content = await response.Content.ReadAsStringAsync();
            // var jo = JsonSerializer.Deserialize<object>(content);
            // return Results.Json(jo);
            return await response.ParseResponseAsync();
        });

    }
}
