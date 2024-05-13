using System.Net;
using Explorer.Stakeholders.API.Dtos;
using gRPCAdapter.Services;
using Explorer.Stakeholders.API.Public;
using Grpc.Core;

namespace gRPCAdapter.Services;

public class AuthService: Auth.AuthBase
{
    private readonly IAuthenticationService _authenticationService;

    public AuthService(IAuthenticationService authenticationService)
    {
        _authenticationService = authenticationService;
    }

    public override Task<LoginResponse> Login(LoginRequest request, ServerCallContext context)
    {
        CredentialsDto mappedToCredentials = new CredentialsDto();
        mappedToCredentials.Username = request.Username;
        mappedToCredentials.Password = request.Password;
        var result = _authenticationService.Login(mappedToCredentials);

        return Task.FromResult(new LoginResponse
        {
            Id = result.Value.Id,
            AccessToken = result.Value.AccessToken
        });
    }
}