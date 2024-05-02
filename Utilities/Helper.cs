using System.Collections.Specialized;
using System.Text.Json;

namespace GenomeDataApi.Utilities;

public static class HttpResponseMessageExtensions
{
    public static async Task<IResult> ParseResponseAsync(this HttpResponseMessage response)
    {
        var content = await response.Content.ReadAsStringAsync();
        var jo = JsonSerializer.Deserialize<object>(content);
        return Results.Json(jo);
    }
}

public static class NameValueCollectionExtensions
{
    public static void SetIfNotNull(this NameValueCollection collection, string key, object? value)
    {
        if (value != null)
        {
            collection[key] = value.ToString();
        }
    }

    public static void SetIfNotNull(this NameValueCollection collection, string key, bool? value)
    {
        if (value != null)
        {
            collection[key] = value switch
            {
                true => "1",
                false => "0",
            };
        }
    }
}
