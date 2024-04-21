using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Stakeholders.API.Dtos;
using Explorer.Stakeholders.API.Public;
using Explorer.Stakeholders.Infrastructure.Authentication;
using FluentResults;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers;

[Authorize(Policy = "userPolicy")]
[Route("api/profile")]
public class ProfileController : BaseApiController
{
    private readonly IProfileService _profileService;
    private static HttpClient _followerHttpClient = new()
    {
        BaseAddress = new Uri("http://localhost:8000"),
    };
    
    public ProfileController(IProfileService profileService)
    {
        _profileService = profileService;
    }

    [HttpGet("{userId:int}")]
    public ActionResult<AccountRegistrationDto> GetStakeholderProfile(long userId)
    {
        var result = _profileService.GetProfile(userId);
        return CreateResponse(result);
    }
    
    [HttpGet("zelimdaumrem/{userId:int}")]
    public ActionResult<AccountRegistrationDto> GetPersonDto(long userId)
    {
        var result = _profileService.GetPersonDto(userId);
        return CreateResponse(result);
    }

    [HttpGet("not-followed")]
    public ActionResult<PagedResult<PersonDto>> GetNonFollowedProfiles([FromQuery] int page, [FromQuery] int pageSize)
    {
        var result = _profileService.GetUserNonFollowedProfiles(page, pageSize, User.PersonId());
        return CreateResponse(result);
    }

    [HttpGet("followers")]
    public ActionResult<List<PersonDto>> GetFollowers()
    {
        var response = _followerHttpClient.GetFromJsonAsync<List<int>>($"/api/v1/followers/{User.PersonId()}").Result;
        // var result = _profileService.GetFollowers(User.PersonId());
        var result = _profileService.GetPersonsFromIds(response);
        return CreateResponse(result);
        return CreateResponse(result);
    }

    [HttpGet("following")]
    public ActionResult<List<PersonDto>> GetFollowing()
    {
        var response = _followerHttpClient.GetFromJsonAsync<List<int>>($"/api/v1/followings/{User.PersonId()}").Result;
        // var result = _profileService.GetFollowing(User.PersonId())
        var result = _profileService.GetPersonsFromIds(response);
        return CreateResponse(result);
    }

    [HttpPut("{id:int}")]
    public ActionResult<PersonDto> Update(int id, [FromBody] PersonDto updatedPerson)
    {
        updatedPerson.Id = id;

        var result = _profileService.UpdateProfile(updatedPerson);
        return CreateResponse(result);
    }

    [HttpPut("follow")]
    public ActionResult<PagedResult<PersonDto>> Follow([FromBody] long followedId)
    {
        try
        {
            var response = _followerHttpClient.PostAsJsonAsync<object>($"/api/v1/follow/{User.PersonId()}/{followedId}", null).Result;
            // var result = _profileService.Follow(User.PersonId(), followedId);
            return CreateResponse(Result.Ok());
        }
        catch (ArgumentException e)
        {
            return CreateResponse(Result.Fail(FailureCode.InvalidArgument).WithError(e.Message));
        }
    }

    [HttpPut("unfollow")]
    public ActionResult<PagedResult<PersonDto>> Unfollow([FromBody] long unfollowedId)
    {
        try
        {
            var result = _profileService.Unfollow(User.PersonId(), unfollowedId);
            return CreateResponse(result);
        }
        catch (ArgumentException e)
        {
            return CreateResponse(Result.Fail(FailureCode.InvalidArgument).WithError(e.Message));
        }
    }

    [HttpGet("canCreateEncounters")]
    public ActionResult<bool> CanTouristCreateEncounters()
    {
        var result = _profileService.CanTouristCreateEncounters(ClaimsPrincipalExtensions.PersonId(User));
        return CreateResponse(result);
    }
    
    [HttpGet("followerRecommendation")]
    public ActionResult<List<PersonDto>> GetFollowerRecommendation()
    {
        var response = _followerHttpClient.GetFromJsonAsync<IEnumerable<int>>($"/api/v1/recommendations/{User.PersonId()}").Result;
        var result = _profileService.GetPersonsFromIds(response);
        return CreateResponse(result);
    }
}