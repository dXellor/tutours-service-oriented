from fastapi import APIRouter

router = APIRouter()

@router.get('/', status_code=200)
async def test_endpoint():
    reponse = {
        "message": 'test'
    }

    return reponse