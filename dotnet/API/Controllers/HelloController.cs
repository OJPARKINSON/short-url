using System.Diagnostics;
using Microsoft.AspNetCore.Mvc;
using API.Models;
using Microsoft.AspNetCore.Http.HttpResults;

namespace API.Controllers;

[Route("api/hello")]
[ApiController]
public class HelloController : ControllerBase
{

    [HttpGet]
    public ActionResult GetHello()
    {
        try
        {
            return Ok("Hello world!");
        }
        catch (Exception ex)
        {
            Console.WriteLine(ex);
            return Forbid();
        }

    }
}
