using System.Net.Mime;
using System.Runtime.InteropServices.JavaScript;
using Microsoft.AspNetCore.Mvc;

namespace short_url.Controllers
{
    [ApiController]
    [Route("api/shorten")]
    [Produces("application/json")]
    public class ShortenController : ControllerBase
    {
        private readonly ILogger<ShortenController> _logger;

        public ShortenController(ILogger<ShortenController> logger)
        {
            _logger = logger;
        }
        
        [HttpGet]
        public ActionResult<urlStore> GetProducts()
        {
            urlStore urlStore = new urlStore();
            return Ok(urlStore);
        }

        [HttpPost]
        [Consumes(MediaTypeNames.Application.Json)]
        [ProducesResponseType<urlStore>(StatusCodes.Status201Created)]
        [ProducesResponseType(StatusCodes.Status400BadRequest)]
        [ProducesResponseType(StatusCodes.Status409Conflict)]
        public ActionResult<urlStore> ShortenUrl(HttpRequest req)
        {
            urlStore urlStore = new urlStore();
            urlStore.url = req.Form["url"];
            return Created("test",urlStore);
        }
        
    }
}