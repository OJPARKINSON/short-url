using System;
using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace API.Models;

public class ShortUrl
{
    [BsonId]
    [BsonRepresentation(BsonType.ObjectId)]
    public string? Id { get; set; }

    [BsonElement("url")]
    public string OriginalUrl { get; set; } = null!;

    [BsonElement("shortcode")]
    public string ShortCode { get; set; } = null!;

    [BsonElement("createdat")]
    public DateTime CreatedAt { get; set; }

    [BsonElement("updatedat")]
    public DateTime UpdatedAt { get; set; }

    [BsonElement("accesscount")]
    public int AccessCount { get; set; }
}
