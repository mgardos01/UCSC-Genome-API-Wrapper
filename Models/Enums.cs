using System.ComponentModel.DataAnnotations;
using System.Text.Json.Serialization;
using Microsoft.AspNetCore.Mvc;

namespace GenomeDataApi.Models;

[JsonConverter(typeof(JsonStringEnumConverter))]
public enum Category
{
    helpDocs,
    publicHubs,
    trackDb
}

[JsonConverter(typeof(JsonStringEnumConverter))]
public enum Format
{
    text
}
