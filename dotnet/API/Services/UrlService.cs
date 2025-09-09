using System;
using System.CodeDom.Compiler;
using API.Models;
using Microsoft.AspNetCore.Http.HttpResults;
using MongoDB.Bson;
using MongoDB.Driver;

namespace API.Services;

public interface IUrlService
{
    Task<ShortUrl> CreateShortUrlAsync(string OriginalUrl);
    Task<string> GetOriginalUrlAsync(string ShortCode);
}

public class UrlService : IUrlService
{
    private readonly IMongoCollection<ShortUrl> _collection;

    public UrlService(IMongoDatabase database)
    {
        _collection = database.GetCollection<ShortUrl>("urls");
    }

    public async Task<ShortUrl> CreateShortUrlAsync(string OriginalUrl)
    {
        // todo at url validation

        var shortCode = GenerateShortCode();

        var shortUrl = new ShortUrl
        {
            OriginalUrl = OriginalUrl,
            ShortCode = shortCode,
            CreatedAt = DateTime.UtcNow,
            UpdatedAt = DateTime.UtcNow,
            AccessCount = 0,
        };


        await _collection.InsertOneAsync(shortUrl);

        return shortUrl;
    }
    public async Task<String> GetOriginalUrlAsync(string shortCode)
    {
        var url = await _collection.Find(u => u.ShortCode == shortCode).FirstOrDefaultAsync();

        Console.WriteLine(url);

        if (url == null)
        {
            // Not found exception
            throw new Exception("Short URL not found");
        }

        return url.OriginalUrl;
    }

    private static Random random = new Random();

    private string GenerateShortCode()
    {
        const int SortCodeLength = 6;
        const string chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
        return new string(Enumerable.Repeat(chars, SortCodeLength)
            .Select(s => s[random.Next(s.Length)]).ToArray());
    }
}
