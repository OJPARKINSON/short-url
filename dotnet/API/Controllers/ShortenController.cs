using Microsoft.AspNetCore.Mvc;

namespace API.Controllers;

[Route("api/shorten")]
[ApiController]
public class ShortenController : ControllerBase
{
    [HttpGet]
    public ActionResult GetShortenUrl()
    {
        return Ok();
    }
}
