from dotenv import load_dotenv
from fastapi import APIRouter
import os

from service.follower_service import FollowerService

load_dotenv()
follower_service = FollowerService(os.getenv('DB_NAME'),
                                    os.getenv('DB_URI'),
                                    os.getenv('DB_USER'), 
                                    os.getenv('DB_PWD'))
router = APIRouter()

@router.get('/', status_code=200)
async def test_endpoint():
    reponse = {
        "message": 'test'
    }

    return reponse

@router.post('/follow/{userId1}/{userId2}')
async def follow(userId1: int, userId2: int):
    follower_service.follow(userId1, userId2)

@router.get('/followers/{userId}')
async def get_followings(userId: int):
    return follower_service.get_followers(userId)


@router.get('/followings/{userId}')
async def get_followings(userId: int):
    return follower_service.get_followings(userId)

@router.get('/recommendations/{userId}')
async def get_recommendations(userId: int):
    return follower_service.get_recommendations(userId)