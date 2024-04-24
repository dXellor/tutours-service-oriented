using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Encounters.API.Dtos;
using Explorer.Encounters.API.Public;
using Explorer.Encounters.Core.Domain;
using Explorer.Stakeholders.Infrastructure.Authentication;
using Explorer.Tours.API.Dtos;
using FluentResults;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using System.Diagnostics.Metrics;

namespace Explorer.API.Controllers
{
    [Authorize(Policy = "userPolicy")]
    [Route("api/encounter")]
    public class EncounterController : BaseApiController
    {
        private readonly IEncounterService _encounterService;
        private static HttpClient _encounterHttpClient = new()
        {
            BaseAddress = new Uri("http://ms-encounters:7007"),
        };

        public EncounterController(IEncounterService encounterService)
        {
            _encounterService = encounterService;
        }


        [HttpGet]
        public ActionResult<PagedResult<EncounterDto>> GetApproved([FromQuery] int page, [FromQuery] int pageSize) // treating like getALl for now
        {
            //var result = _encounterService.GetApproved(page, pageSize);

            // need some better protection when service is down
            // conversion bad, I can't
            var encounterDto = _encounterHttpClient.GetFromJsonAsync<List<EncounterDto>>("/all").Result;
            var pagedResult = new PagedResult<EncounterDto>(encounterDto, encounterDto.Count);
            var result = Result.Ok<PagedResult<EncounterDto>>(pagedResult);

            return CreateResponse(result);
        }

        [HttpPost]
        public ActionResult<EncounterDto> Create([FromBody] EncounterDto encounter)
        {
            encounter.UserId = User.PersonId();
            //var result = _encounterService.Create(encounter);

            var response = _encounterHttpClient.PostAsJsonAsync<EncounterDto>("/", encounter).Result;
            var encounterDto = response.Content.ReadFromJsonAsync<EncounterDto>().Result;
            var result = Result.Ok<EncounterDto>(encounterDto);

            return CreateResponse(result);
        }

        [HttpPut("{id:int}")]
        public ActionResult<EncounterDto> Update(int id, [FromBody] EncounterDto encounter)
        {
            //var result = _encounterService.Update(encounter);

            var response = _encounterHttpClient.PutAsJsonAsync<EncounterDto>($"/{id}", encounter).Result;
            var encounterDto = response.Content.ReadFromJsonAsync<EncounterDto>().Result;
            var result = Result.Ok<EncounterDto>(encounterDto);

            return CreateResponse(result);
        }

        [HttpDelete("{id:int}")]
        public ActionResult Delete(int id)
        {
            //var result = _encounterService.Delete(id);

            var response = _encounterHttpClient.DeleteAsync($"/{id}").Result;
            var result = Result.Ok();

            return CreateResponse(result);
        }

        [HttpGet("status")]
        public ActionResult<PagedResult<EncounterDto>> GetApprovedByStatus([FromQuery] int page, [FromQuery] int pageSize, [FromQuery] string status)
        {
            //var result = _encounterService.GetApprovedByStatus(page, pageSize, status);

            var encounterDto = _encounterHttpClient.GetFromJsonAsync<List<EncounterDto>>($"/status?status={status}").Result;
            var pagedResult = new PagedResult<EncounterDto>(encounterDto, encounterDto.Count);
            var result = Result.Ok<PagedResult<EncounterDto>>(pagedResult);

            return CreateResponse(result);
        }        
        
        [HttpGet("nearbyHidden")]
        public ActionResult<PagedResult<EncounterDto>> GetNearbyHidden([FromQuery] int page, [FromQuery] int pageSize)
        {
            var userId = ClaimsPrincipalExtensions.PersonId(User);
            var result = _encounterService.GetNearbyHidden(page, pageSize, userId);
            return CreateResponse(result);
        }

        [HttpGet("nearby")]
        public ActionResult<PagedResult<EncounterDto>> GetNearby([FromQuery] int page, [FromQuery] int pageSize)
        {
            var userId = ClaimsPrincipalExtensions.PersonId(User);
            var result = _encounterService.GetNearby(page, pageSize, userId);
            return CreateResponse(result);
        }

        [HttpGet("byUser")]
        public ActionResult<PagedResult<EncounterDto>> GetByUser([FromQuery] int page, [FromQuery] int pageSize)
        {
            //var result = _encounterService.GetByUser(page, pageSize, ClaimsPrincipalExtensions.PersonId(User));

            var userId = ClaimsPrincipalExtensions.PersonId(User);
            var encounterDto = _encounterHttpClient.GetFromJsonAsync<List<EncounterDto>>($"/byUser/{userId}").Result;
            var pagedResult = new PagedResult<EncounterDto>(encounterDto, encounterDto.Count);
            var result = Result.Ok<PagedResult<EncounterDto>>(pagedResult);

            return CreateResponse(result);
        }

        [HttpPut("approve")]
        [Authorize(Roles = "administrator")]
        public ActionResult<EncounterDto> Approve(EncounterDto encounter)
        {
            //var result = _encounterService.Approve(encounter);
            var response = _encounterHttpClient.PutAsJsonAsync<EncounterDto>("/approve", encounter).Result;
            var encounterDto = response.Content.ReadFromJsonAsync<EncounterDto>().Result;
            var result = Result.Ok<EncounterDto>(encounterDto);

            return CreateResponse(result);
        }

        [HttpPut("decline")]
        [Authorize(Roles = "administrator")]
        public ActionResult<EncounterDto> Decline(EncounterDto encounter)
        {
            //var result = _encounterService.Decline(encounter);
            var response = _encounterHttpClient.PutAsJsonAsync<EncounterDto>("/decline", encounter).Result;
            var encounterDto = response.Content.ReadFromJsonAsync<EncounterDto>().Result;
            var result = Result.Ok<EncounterDto>(encounterDto);

            return CreateResponse(result);
        }

        [HttpGet("touristCreatedEncouters")]
        public ActionResult<PagedResult<EncounterDto>> GetTouristCreatedEncounters([FromQuery] int page, [FromQuery] int pageSize)
        {
            //var result = _encounterService.GetTouristCreatedEncounters(page, pageSize);

            var encounterDto = _encounterHttpClient.GetFromJsonAsync<List<EncounterDto>>("/touristCreatedEncounters").Result;
            var pagedResult = new PagedResult<EncounterDto>(encounterDto, encounterDto.Count);
            var result = Result.Ok<PagedResult<EncounterDto>>(pagedResult);

            return CreateResponse(result);
        }
        
    }
}
