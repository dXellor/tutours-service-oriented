using System.Net;
using Explorer.Stakeholders.API.Dtos;
using gRPCAdapter.Services;
using Explorer.Stakeholders.API.Public;
using Grpc.Core;

namespace gRPCAdapter.Services;

public class AuthService: Monolith.MonolithBase
{
    private readonly IAuthenticationService _authenticationService;
    
    public AuthService(IAuthenticationService authenticationService)
    {
        _authenticationService = authenticationService;
    }

    public override Task<LoginResponse> Login(LoginRequest request, ServerCallContext context)
    {
        var mappedToCredentials = new CredentialsDto
        {
            Username = request.Username,
            Password = request.Password
        };
        var result = _authenticationService.Login(mappedToCredentials);
    
        return Task.FromResult(new LoginResponse
        {
            Id = result.Value.Id,
            AccessToken = result.Value.AccessToken
        });
    }
}