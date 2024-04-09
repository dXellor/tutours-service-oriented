from contextlib import asynccontextmanager
from fastapi import FastAPI
import logging

from api import router

logging.basicConfig(format='%(levelname)s - %(asctime)s - %(message)s')

@asynccontextmanager
async def lifespan(app: FastAPI):
    logging.info("Follower service started")
    yield
    logging.info("Follower service stopped")

app = FastAPI(
    title="Follower microservice",
    description="Microserice for user connections tracking",
    openapi_url="/api/v1/followers/openapi.json",
    docs_url="/api/v1/followers/docs",
    version="v1",
    lifespan=lifespan)

app.include_router(router, prefix='/api/v1/followers', tags=['followers'])