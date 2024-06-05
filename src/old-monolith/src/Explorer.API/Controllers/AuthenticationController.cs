using Explorer.Stakeholders.API.Dtos;
using Explorer.Stakeholders.API.Public;
using FluentResults;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers;

[Route("api/users")]
public class AuthenticationController : BaseApiController
{
    private readonly IAuthenticationService _authenticationService;

    private static readonly HttpClient AuthHttpClient = new()
    {
        BaseAddress = new Uri("http://ms-auth:8040"),
    };

    public AuthenticationController(IAuthenticationService authenticationService)
    {
        _authenticationService = authenticationService;
    }

    [HttpPost]
    public ActionResult<AuthenticationTokensDto> RegisterTourist([FromBody] AccountRegistrationDto account)
    {
        var result = _authenticationService.RegisterTourist(account);
        return CreateResponse(result);
    }

    [HttpPost("login")]
    public ActionResult<AuthenticationTokensDto> Login([FromBody] CredentialsDto credentials)
    {
        var response = AuthHttpClient.PostAsJsonAsync("/login", credentials).Result;
        var auth = response.Content.ReadFromJsonAsync<AuthenticationTokensDto>().Result;
        var result = Result.Ok<AuthenticationTokensDto>(auth);
        return CreateResponse(result);
    }



}