using System.Text.Json.Serialization;
using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Stakeholders.Infrastructure.Authentication;
using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.Public.TourAuthoring;
using FluentResults;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers.Author;

[Authorize(Policy = "userPolicy")]
[Route("api/author/tours/")]
public class TourManagementController : BaseApiController
{
    private readonly ITourService _tourService;
    private static HttpClient _tourHttpClient = new()
    {
        BaseAddress = new Uri("http://localhost:7007"),
    };

    public TourManagementController(ITourService tourService)
    {
        _tourService = tourService;
    }

    [HttpGet]
    [Authorize(Roles = "author, tourist")]
    public ActionResult<PagedResult<TourDto>> GetAll([FromQuery] int page, [FromQuery] int pageSize)
    {
        var result = _tourService.GetPaged(page, pageSize);
        return CreateResponse(result);
    }

    [AllowAnonymous]
    [HttpGet("{tourId:int}")]
    [Authorize(Roles = "author")]
    public ActionResult<TourDto> GetById([FromRoute] int tourId)
    {
        var result = _tourService.Get(tourId);
        return CreateResponse(result);
    }

    [HttpPost]
    [Authorize(Roles = "author")]
    public ActionResult<TourDto> Create([FromBody] TourDto tour)
    {
        tour.UserId = ClaimsPrincipalExtensions.PersonId(User);
        var response = _tourHttpClient.PostAsJsonAsync<TourDto>("/", tour).Result;
        var tourDto = response.Content.ReadFromJsonAsync<TourDto>().Result;
        var result = Result.Ok<TourDto>(tourDto);

        return CreateResponse(result);
    }

    [HttpPut("{id:int}")]
    [Authorize(Roles = "author,tourist")]
    public ActionResult<TourDto> Update(int id, [FromBody] TourDto tour)
    {
        var response = _tourHttpClient.PutAsJsonAsync<TourDto>($"/{id}", tour).Result;
        var tourDto = response.Content.ReadFromJsonAsync<TourDto>().Result;
        var result = Result.Ok<TourDto>(tourDto);

        return CreateResponse(result);
    }
    
    [HttpDelete("{id:int}")]
    [Authorize(Roles = "author")]
    public ActionResult Delete(int id)
    {
        var response = _tourHttpClient.DeleteAsync($"/{id}").Result;
        var result = Result.Ok();

        return CreateResponse(result);
    }

    [HttpGet("author")]
    [Authorize(Roles = "author")]
    public ActionResult<PagedResult<TourDto>> GetByAuthor([FromQuery] int page, [FromQuery] int pageSize)
    {
        var authorId = ClaimsPrincipalExtensions.PersonId(User);
        var response = _tourHttpClient.GetFromJsonAsync<TourDto[]>($"/{authorId}").Result;
        var result = Result.Ok<TourDto[]>(response);
        return CreateResponse(result);
    }

    [AllowAnonymous]
    [HttpPut("disable/{id:int}")]
    public ActionResult<TourDto> Disable([FromBody] TourDto tour)
    {
        if (User.IsInRole("administrator"))
        {
            tour.UserId = User.PersonId();
            var result = _tourService.Update(tour);
            return CreateResponse(result);
        }
        return null;
    }
    
    [HttpPost("custom")]
    public ActionResult<TourDto> CreateCustomTour([FromBody] TourDto tourDto)
    {   
        tourDto.UserId = ClaimsPrincipalExtensions.PersonId(User);
        var result = _tourService.CreateCustom(tourDto);
        return CreateResponse(result);
    }

    [HttpPost("campaign")]
    public ActionResult<TourDto> CreateCampaignTour([FromBody] TourDto tourDto) 
    {
        tourDto.UserId = ClaimsPrincipalExtensions.PersonId(User);
        var result = _tourService.CreateCampaign(tourDto);
        return CreateResponse(result);
    }
}