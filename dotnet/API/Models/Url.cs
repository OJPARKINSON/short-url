using System;

namespace API.Models;

public class Url
{
    public string OriginalUrl { get; set; } = null!;

    public string ShortUrl { get; set; } = null!;

    public DateTime CreatedAt { get; set; }
    public DateTime UpdatedAt { get; set; }
}
