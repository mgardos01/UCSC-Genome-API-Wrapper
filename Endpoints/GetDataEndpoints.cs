using System.Web;
using GenomeDataApi.Models;
using GenomeDataApi.Utilities;
using Microsoft.AspNetCore.Mvc;

namespace GenomeDataApi.Endpoints;

public static class GetDataEndpoints
{
    public static void MapEndpoints(WebApplication app)
    {
        var getData = app.MapGroup("/getData")
            .WithTags("getData");

        // Define mappings inside here
        getData.MapGet("/sequence", async (
            IHttpClientFactory httpClientFactory,
            [FromQuery(Name = "genome")] string name,
            [FromQuery(Name = "hubUrl")] string? url,
            [FromQuery(Name = "chrom")] string chrN,
            [FromQuery] long? start, //TODO add documentation that this and "end" always need to be together 
            [FromQuery] long? end    //TODO add documentation that this and "start always need to always be together 
            ) =>
        {
            var client = httpClientFactory.CreateClient("genomeClient");

            // Create the query string directly
            var query = HttpUtility.ParseQueryString(string.Empty);
            query.SetIfNotNull("genome", name);
            query.SetIfNotNull("hubUrl", url);
            query.SetIfNotNull("chrom", chrN);
            query.SetIfNotNull("start", start);
            query.SetIfNotNull("end", end);
            var response = await client.GetAsync($"/getData/sequence?{query}");
            return await response.ParseResponseAsync();
        });

        getData.MapGet("/track", async (IHttpClientFactory httpClientFactory,
            [FromQuery(Name = "genome")] string name,
            [FromQuery(Name = "track")] string trackName,
            [FromQuery(Name = "hubUrl")] string? url,
            [FromQuery(Name = "chrom")] string? chrN,
            [FromQuery] long? start, //TODO add documentation that this and "end" always need to be together 
            [FromQuery] long? end,    //TODO add documentation that this and "start always need to always be together 
            [FromQuery] bool? jsonOutputArrays,
            [FromQuery] int? maxItemsOutput = 1000) =>
        {
            var client = httpClientFactory.CreateClient("genomeClient");

            // Create the query string directly
            var query = HttpUtility.ParseQueryString(string.Empty);

            query.SetIfNotNull("genome", name);
            query.SetIfNotNull("track", trackName);
            query.SetIfNotNull("hubUrl", url);
            query.SetIfNotNull("chrom", chrN);
            query.SetIfNotNull("start", start);
            query.SetIfNotNull("end", end);
            query.SetIfNotNull("jsonOutputArrays", jsonOutputArrays);
            query.SetIfNotNull("maxItemsOutput", maxItemsOutput);

            var response = await client.GetAsync($"/getData/track?{query}");
            return await response.ParseResponseAsync();
        });
    }
}