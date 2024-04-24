using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Encounters.API.Dtos;
using Explorer.Encounters.API.Public;
using Explorer.Encounters.Core.UseCases;
using Explorer.Stakeholders.Infrastructure.Authentication;
using FluentResults;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers.Tourist.Statistics
{
    [Authorize(Policy = "touristPolicy")]
    [Route("api/tourist/encounterStatistics")]
    public class EncounterStatisticsController : BaseApiController
    {
        private readonly IStatisticsService _statisticsService;
        private static HttpClient _encounterHttpClient = new()
        {
            BaseAddress = new Uri("http://ms-encounters:7007"),
        };
        
        public EncounterStatisticsController(IStatisticsService statisticsService)
        {
           _statisticsService = statisticsService;
        }

        [HttpGet("completions")]
        public ActionResult<EncounterStatsDto> GetByUser()
        {
            var userId = ClaimsPrincipalExtensions.PersonId(User);
            var response = _encounterHttpClient.GetFromJsonAsync<EncounterStatsDto>($"/completions/{userId}").Result;
            var result = Result.Ok<EncounterStatsDto>(response);
            return CreateResponse(result);
        }

        [HttpGet("yearCompletions")]
        public ActionResult<EncounterYearStatsDto> GetByUserAndYear([FromQuery] int year)
        {
            var userId = ClaimsPrincipalExtensions.PersonId(User);
            // var result = _statisticsService.GetEncounterYearStatsByUser(userId, year);
            var response = _encounterHttpClient.GetFromJsonAsync<EncounterYearStatsDto>($"/yearCompletions/{userId}?year={year}").Result;
            var result = Result.Ok<EncounterYearStatsDto>(response);
            return CreateResponse(result);
        }

    }
}
