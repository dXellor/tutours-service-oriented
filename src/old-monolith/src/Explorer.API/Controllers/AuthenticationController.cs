using Explorer.Stakeholders.API.Dtos;
using Explorer.Stakeholders.API.Public;
using FluentResults;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers;

[Route("api/users")]
public class AuthenticationController : BaseApiController
{
    private readonly IAuthenticationService _authenticationService;

    private static HttpClient _HttpClient = new()
    {
        BaseAddress = new Uri("http://ms-auth:7007"),
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
        var response = _HttpClient.PostAsJsonAsync<CredentialsDto>("/auth/login", credentials).Result;
        var credentialsDto = response.Content.ReadFromJsonAsync<CredentialsDto>().Result;
        var result = Result.Ok<CredentialsDto>(credentialsDto);
        return CreateResponse(result);
    }



}