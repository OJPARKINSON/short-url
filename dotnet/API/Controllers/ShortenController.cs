using API.Services;
using Microsoft.AspNetCore.Mvc;

namespace API.Controllers;

public class CreateBody
{
    public string url { get; set; } = null!;
}

[Route("api/shorten")]
[ApiController]
public class ShortenController : ControllerBase
{
    private readonly IUrlService _urlService;

    public ShortenController(IUrlService urlService)
    {
        _urlService = urlService;
    }

    [HttpPost]
    public async Task<ActionResult> CreateShortenUrl([FromBody] CreateBody body)
    {
        var result = await _urlService.CreateShortUrlAsync(body.url);

        return Ok(result);
    }

    [HttpGet]
    [Route("{shortcode}")]
    public async Task<ActionResult> GetShortenUrl(string shortcode)
    {
        var url = await _urlService.GetOriginalUrlAsync(shortcode);

        return Ok(url);
    }
}
