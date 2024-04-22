import logging
from neo4j import GraphDatabase
from neo4j import Record

class FollowerService():

    def __init__(self, db_name, db_uri, username, password) -> None:
        self.db_name = db_name
        self.driver = GraphDatabase.driver(db_uri, auth=(username, password))
        self.driver.verify_connectivity()
        self.__init_database()

    def __del__(self) -> None:
        self.driver.close()

    def __init_database(self):
        try:
            self.driver.execute_query("CREATE CONSTRAINT FOR (user:USER) REQUIRE user.id IS UNIQUE;")
        except:
            logging.info("Database is already setup")

    def __insertIfNotExist(self, userId: int) -> None:
        records, _, _ = self.driver.execute_query(
            "MATCH (u:USER {id: $userId}) RETURN u.id",
            userId=userId,
            database_= self.db_name
        )

        if len(records) == 0:
            self.driver.execute_query("CREATE (u:USER {id: $userId})", userId=userId)

    def follow(self, signedInUserId: int, userToBeFollowedId: int ) -> None:
        self.__insertIfNotExist(signedInUserId)
        self.__insertIfNotExist(userToBeFollowedId)

        self.driver.execute_query(
            "MATCH (signedInUser:USER {id: $signedInUserId}), (userToBeFollowed:USER{id: $userToBeFollowedId}) CREATE (signedInUser)-[:FOLLOWS]->(userToBeFollowed)",
            database_= self.db_name,
            signedInUserId=signedInUserId,
            userToBeFollowedId=userToBeFollowedId,
        )

    def get_followers(self, signedInUserId: int):
        records, summary, keys = self.driver.execute_query(
            "MATCH (u:USER {id: $userId})<-[:FOLLOWS]-(u2:USER) RETURN collect(u2.id)",
            database_= self.db_name,
            userId=signedInUserId
        )

        return records[0][0]

    def get_followings(self, signedInUserId: int):
        records, summary, keys = self.driver.execute_query(
            "MATCH (u:USER {id: $userId})-[:FOLLOWS]->(u2:USER) RETURN collect(u2.id)",
            database_= self.db_name,
            userId=signedInUserId
        )

        return records[0][0]
    
    def get_recommendations(self, signedInUserId: int):
        records, summary, keys = self.driver.execute_query(
            """MATCH (u1:USER {id: 1})-[:FOLLOWS]->(u2:USER)-[:FOLLOWS]->(u3:USER) 
            WHERE u1 <> u3 
            AND NOT (u1)-[:FOLLOWS]->(u3) 
            RETURN collect(u3.id)""",
            database_= self.db_name,
            userId=signedInUserId
        )

        return records[0][0]

        
