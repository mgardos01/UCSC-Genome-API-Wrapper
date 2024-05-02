using System.Text.Json;
using System.Web;
using GenomeDataApi.Models;
using GenomeDataApi.Utilities;
using Microsoft.AspNetCore.Mvc;

namespace GenomeDataApi.Endpoints;

public static class ListEndpoints
{

    public static void MapEndpoints(WebApplication app)
    {
        var list = app.MapGroup("/list")
            .WithTags("list");

        // Define mappings inside here
        list.MapGet("/publicHubs", async (IHttpClientFactory httpClientFactory) =>
        {
            var client = httpClientFactory.CreateClient("genomeClient");
            var url = $"{client.BaseAddress}publicHubs";
            var response = await client.GetAsync("/list/publicHubs");
            return await response.ParseResponseAsync();
        });

        list.MapGet("/ucscGenomes", async (IHttpClientFactory httpClientFactory) =>
        {
            var client = httpClientFactory.CreateClient("genomeClient");
            var response = await client.GetAsync("/list/ucscGenomes");
            return await response.ParseResponseAsync();
        });

        list.MapGet("/hubGenomes", async (
            IHttpClientFactory httpClientFactory,
            [FromQuery(Name = "hubUrl")] string url // TODO Make serializable else later
            ) =>
        {
            var client = httpClientFactory.CreateClient("genomeClient");
            // Create the query string directly
            var query = HttpUtility.ParseQueryString(string.Empty);
            query.SetIfNotNull("hubUrl", url);
            var response = await client.GetAsync($"/list/hubGenomes?{query}");
            return await response.ParseResponseAsync();
        });

        list.MapGet("/files", async (
            IHttpClientFactory httpClientFactory,
            [FromQuery(Name = "genome")] string name, // TODO Make serializable else later
            [FromQuery] Format? format, // TODO Make serializable else later
            [FromQuery] int? maxItemsOutput = 1000 // MAXIMUM: 1_000_000
            ) =>
        {
            var client = httpClientFactory.CreateClient("genomeClient");

            // Create the query string directly
            var query = HttpUtility.ParseQueryString(string.Empty);
            query.SetIfNotNull("genome", name);
            query.SetIfNotNull("format", format);
            query.SetIfNotNull("maxItemsOutput", maxItemsOutput);
            var response = await client.GetAsync($"/list/files?{query}");
            return await response.ParseResponseAsync();
        });

        list.MapGet("/tracks", async (
            IHttpClientFactory httpClientFactory,
            [FromQuery(Name = "genome")] string name,
            [FromQuery(Name = "hubUrl")] string? url,
            [FromQuery] bool? trackLeavesOnly
            ) =>
        {
            var client = httpClientFactory.CreateClient("genomeClient");

            // Create the query string directly
            var query = HttpUtility.ParseQueryString(string.Empty);
            query.SetIfNotNull("genome", name);
            query.SetIfNotNull("hubUrl", url);
            query.SetIfNotNull("trackLeavesOnly", trackLeavesOnly);
            var response = await client.GetAsync($"/list/tracks?{query}");
            return await response.ParseResponseAsync();
        });

        list.MapGet("/chromosomes", async (
            IHttpClientFactory httpClientFactory,
            [FromQuery(Name = "genome")] string name,
            [FromQuery(Name = "hubUrl")] string? url,
            [FromQuery(Name = "track")] string? trackName
            ) =>
        {
            var client = httpClientFactory.CreateClient("genomeClient");

            // Create the query string directly
            var query = HttpUtility.ParseQueryString(string.Empty);
            query.SetIfNotNull("genome", name);
            query.SetIfNotNull("hubUrl", url);
            query.SetIfNotNull("track", trackName);

            var response = await client.GetAsync($"/list/chromosomes?{query}");
            return await response.ParseResponseAsync();
        });

        list.MapGet("/schema", async (
            IHttpClientFactory httpClientFactory,
            [FromQuery(Name = "genome")] string name,
            [FromQuery(Name = "hubUrl")] string? url,
            [FromQuery(Name = "track")] string trackName
            ) =>
        {
            var client = httpClientFactory.CreateClient("genomeClient");

            // Create the query string directly
            var query = HttpUtility.ParseQueryString(string.Empty);

            query.SetIfNotNull("genome", name);
            query.SetIfNotNull("hubUrl", url);
            query.SetIfNotNull("track", trackName);

            var response = await client.GetAsync($"/list/schema?{query}");
            return await response.ParseResponseAsync();
        });
    }
}