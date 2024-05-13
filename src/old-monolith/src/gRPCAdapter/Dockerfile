FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build
ARG BUILD_CONFIGURATION=Release
WORKDIR /src
COPY ["gRPCAdapter/gRPCAdapter.csproj", "gRPCAdapter/"]
RUN dotnet restore "gRPCAdapter/gRPCAdapter.csproj"
COPY . .
WORKDIR "/src/gRPCAdapter"
RUN dotnet build "gRPCAdapter.csproj" -c $BUILD_CONFIGURATION -o /app/build

FROM build AS publish
ARG BUILD_CONFIGURATION=Release
RUN dotnet publish "gRPCAdapter.csproj" -c $BUILD_CONFIGURATION -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "gRPCAdapter.dll"]
