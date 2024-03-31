using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Encounters.API.Dtos;
using Explorer.Encounters.Core.Domain;
using Explorer.Stakeholders.Infrastructure.Authentication;
using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.Public.TourAuthoring;
using Explorer.Tours.Core.Domain;
using FluentResults;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers.Author;

[Authorize(Policy = "personPolicy")]
[Route("api/author/keypoints")]
public class KeypointController : BaseApiController
{
    private readonly IKeypointService _keypointService;
    private static HttpClient _keypointHttpClient = new()
    {
        BaseAddress = new Uri("http://localhost:7007"),
    };

    public KeypointController(IKeypointService keypointService)
    {
        _keypointService = keypointService;
    }

    [HttpGet]
    public ActionResult<PagedResult<KeypointDto>> GetAll([FromQuery] int page, [FromQuery] int pageSize)
    {
        //var result = _keypointService.GetPaged(page, pageSize);

        var keypointDto = _keypointHttpClient.GetFromJsonAsync<List<KeypointDto>>("/all").Result;
        var pagedResult = new PagedResult<KeypointDto>(keypointDto, keypointDto.Count);
        var result = Result.Ok<PagedResult<KeypointDto>>(pagedResult);

        return CreateResponse(result);
    }

    [HttpGet("tour/{tourId:int}")]
    public ActionResult<PagedResult<KeypointDto>> GetByTour([FromQuery] int page, [FromQuery] int pageSize,
        [FromRoute] int tourId)
    {
        //var result = _keypointService.GetByTourId(page, pageSize, tourId);

        var response = _keypointHttpClient.GetFromJsonAsync<KeypointDto[]>($"/tour/{tourId}").Result;
        var result = Result.Ok<KeypointDto[]>(response);

        return CreateResponse(result);
    }

    [HttpPost]
    public ActionResult<KeypointDto> Create([FromBody] KeypointDto keypoint)
    {
        //var result = _keypointService.Create(keypoint);

        var response = _keypointHttpClient.PostAsJsonAsync<KeypointDto>("/", keypoint).Result;
        var keypointDto = response.Content.ReadFromJsonAsync<KeypointDto>().Result;
        var result = Result.Ok<KeypointDto>(keypointDto);

        return CreateResponse(result);
    }

    [HttpPost("/multiple")]
    public ActionResult<KeypointDto> CreateMultiple([FromBody] List<KeypointDto> keypoints)
    {
        var result = _keypointService.CreateMultiple(keypoints);
        return CreateResponse(result);
    }

    [HttpPut("{id:int}")]
    public ActionResult<KeypointDto> Update([FromBody] KeypointDto keypoint)
    {
        //var result = _keypointService.Update(keypoint);

        var response = _keypointHttpClient.PutAsJsonAsync<KeypointDto>("/", keypoint).Result;
        var keypointDto = response.Content.ReadFromJsonAsync<KeypointDto>().Result;
        var result = Result.Ok<KeypointDto>(keypointDto);

        return CreateResponse(result);
    }

    [HttpDelete("{id:int}")]
    public ActionResult Delete(int id)
    {
        //var result = _keypointService.Delete(id);

        var response = _keypointHttpClient.DeleteAsync($"/{id}").Result;
        var result = Result.Ok();

        return CreateResponse(result);
    }
}