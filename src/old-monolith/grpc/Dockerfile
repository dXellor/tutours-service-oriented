# https://hub.docker.com/_/microsoft-dotnet
FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS base
WORKDIR /app

FROM mcr.microsoft.com/dotnet/sdk:7.0 as build
WORKDIR /src
COPY . .
WORKDIR /src/src
RUN dotnet restore gRPCAdapter/gRPCAdapter.csproj
RUN dotnet build gRPCAdapter/gRPCAdapter.csproj -c Release

FROM build as publish
RUN dotnet publish gRPCAdapter/gRPCAdapter.csproj -c Release -o /app/publish

ENV ASPNETCORE_URLS=http://0.0.0.0:5172
FROM base AS final
COPY --from=publish /app .
WORKDIR /app/publish

EXPOSE 5172
CMD ["dotnet", "gRPCAdapter.dll"]