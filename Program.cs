using Microsoft.AspNetCore.Mvc;
using GenomeDataApi.Endpoints;
using System.Text.Json.Serialization;
using System.Net.Http.Headers;
using Microsoft.OpenApi.Models;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen(options => 
{
    options.SwaggerDoc("v1", new OpenApiInfo 
    { 
        Title = "Genome Data API", 
        Version = "v1",
        Description = "API for accessing genomic data from UCSC Genome Browser."
    });

    // Define the server and the base path for the API
    options.AddServer(new OpenApiServer 
    { 
        Url = "https://api.genome.ucsc.edu/list/publicHubs",
        Description = "Primary API Server"
    });
});

builder.Services.AddHttpClient("genomeClient", client => 
{ 
    client.BaseAddress = new Uri("https://api.genome.ucsc.edu/"); 
    client.DefaultRequestHeaders.Accept.Clear();
    client.DefaultRequestHeaders.Accept.Add(new MediaTypeWithQualityHeaderValue("application/json"));
    // client.BaseAddress = new Uri("https://pokeapi.co/api/v2/"); 
});

builder.Services.ConfigureHttpJsonOptions(options =>
{
    options.SerializerOptions.Converters.Add(new JsonStringEnumConverter());
});

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();


SearchEndpoints.MapEndpoints(app);
GetDataEndpoints.MapEndpoints(app);
ListEndpoints.MapEndpoints(app);
// TestEndpoints.MapEndpoints(app);

app.Run();
