using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace GenomeDataApi.Endpoints;

public static class TestEndpoints
{    private record WeatherForecast(DateOnly Date, int TemperatureC, string? Summary, int? Toggle)
    {
        public int TemperatureF => 32 + (int)(TemperatureC / 0.5556);
    }

    private static readonly string[] summaries =
    [
    "Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"
    ];

    public static void MapEndpoints(WebApplication app)
    {
        var weatherforecast = app.MapGroup("/weatherforecast");
        weatherforecast.MapGet("/", ([FromQuery] int page) =>
        {
            var forecast = Enumerable.Range(1, 5).Select(index =>
                new WeatherForecast
                (
                    DateOnly.FromDateTime(DateTime.Now.AddDays(index)),
                    Random.Shared.Next(-20, 55),
                    summaries[Random.Shared.Next(summaries.Length)],
                    page
                ))
                .ToArray();
            return forecast;
        })
        .WithName("GetWeatherForecast")
        .WithOpenApi();

        app.MapGet("/pokemon", async (
            IHttpClientFactory httpClientFactory
        ) => {
            var client = httpClientFactory.CreateClient("genomeClient");
            var response = await client.GetAsync("pokemon/ditto");
            return Results.Json(response);
        })
        .WithOpenApi();
    }
}


