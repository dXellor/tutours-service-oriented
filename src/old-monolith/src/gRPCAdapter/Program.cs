using Explorer.Blog.Infrastructure;
using Explorer.Encounters.Infrastructure;
using Explorer.Payments.Infrastructure;
using Explorer.Stakeholders.API.Public;
using Explorer.Stakeholders.Core.Domain.RepositoryInterfaces;
using Explorer.Stakeholders.Core.UseCases;
using Explorer.Stakeholders.Infrastructure;
using Explorer.Stakeholders.Infrastructure.Database.Repositories;
using Explorer.Tours.Infrastructure;
using gRPCAdapter.Services;

var builder = WebApplication.CreateBuilder(args);

// Additional configuration is required to successfully run gRPC on macOS.
// For instructions on how to configure Kestrel and gRPC clients on macOS, visit https://go.microsoft.com/fwlink/?linkid=2099682

// Add services to the container.
builder.Services.AddGrpc();
builder.Services.ConfigureStakeholdersModule();
builder.Services.ConfigureToursModule();
builder.Services.ConfigureBlogModule();
builder.Services.ConfigureEncountersModule();
builder.Services.ConfigurePaymentsModule();

var app = builder.Build();

// Configure the HTTP request pipeline.
app.MapGrpcService<AuthService>();
app.Run();