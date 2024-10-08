package main

import (
	"context"
	"github.com/gin-gonic/gin"
)

func RFunc501() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(501, gin.H{
			"message": "Not Implemented",
		})
	}
}

func RHandlerPing() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func RHandlerCompetitionCreate(comp *ICompetition) func(c *gin.Context) {
	return func(c *gin.Context) {
		print("Create competition: ")

		var HCreateCompetitionBody HCompetitionCreateBody
		err := c.ShouldBindBodyWithJSON(&HCreateCompetitionBody)
		if err != nil {
			print(err.Error())
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		insertedCompetition, err := comp.Create(context.Background(), HCreateCompetitionBody)
		if err != nil {
			print(err.Error())
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, insertedCompetition)
	}
}

func RHandlerCompetitionList(comp *ICompetition) func(c *gin.Context) {
	return func(c *gin.Context) {
		print("List")
		list, err := comp.List(context.Background())
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		//TODO: Is it necessary or do I just return an empty list?
		if list.Competitions == nil {
			c.JSON(204, gin.H{})
			return
		}
		c.JSON(200, list)
	}
}

func RHandlerCompetitionDayCreate(compDay *ICompetitionDay) func(c *gin.Context) {
	return func(c *gin.Context) {
		print("Create competition day: ")

		var HCreateCompetitionDayBody HCompetitionDayCreateBody
		err := c.ShouldBindBodyWithJSON(&HCreateCompetitionDayBody)
		if err != nil {
			print(err.Error())
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		insertedCompetitionDay, err := compDay.Create(context.Background(), HCreateCompetitionDayBody)
		if err != nil {
			print(err.Error())
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, insertedCompetitionDay)
	}
}

func RHandlerCompetitionDayList(compDay *ICompetitionDay) func(c *gin.Context) {
	return func(c *gin.Context) {
		print("List")
		list, err := compDay.List(context.Background())
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, list)
	}
}

func RHandlerParentOrgCreate(parentOrg *IParentOrg) func(c *gin.Context) {
	return func(c *gin.Context) {
		print("Create parent org: ")

		var HCreateParentOrgBody HParentOrgCreateBody
		err := c.ShouldBindBodyWithJSON(&HCreateParentOrgBody)
		if err != nil {
			print(err.Error())
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		insertedParentOrg, err := parentOrg.Create(context.Background(), HCreateParentOrgBody)
		if err != nil {
			print(err.Error())
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, insertedParentOrg)
	}
}

func RHandlerParentOrgList(parentOrg *IParentOrg) func(c *gin.Context) {
	return func(c *gin.Context) {
		print("List")
		list, err := parentOrg.List(context.Background())
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, list)
	}
}

func RHandlerTeamCreate(team *ITeam) func(c *gin.Context) {
	return func(c *gin.Context) {
		print("Create team: ")

		var HCreateTeamBody HTeamCreateBody
		err := c.ShouldBindBodyWithJSON(&HCreateTeamBody)
		if err != nil {
			print(err.Error())
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		insertedTeam, err := team.Create(context.Background(), HCreateTeamBody)
		if err != nil {
			print(err.Error())
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, insertedTeam)
	}
}

func RHandlerTeamList(team *ITeam) func(c *gin.Context) {
	return func(c *gin.Context) {
		print("List")
		list, err := team.List(context.Background())
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, list)
	}
}

func RHandlerRefCreate(refs *IRefs) func(c *gin.Context) {
	return func(c *gin.Context) {
		print("Create ref: ")

		var HCreateRefBody HRefCreateBody
		err := c.ShouldBindBodyWithJSON(&HCreateRefBody)
		if err != nil {
			print(err.Error())
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		insertedRef, err := refs.Create(context.Background(), HCreateRefBody)
		if err != nil {
			print(err.Error())
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, insertedRef)
	}
}

func RHandlerRefList(refs *IRefs) func(c *gin.Context) {
	return func(c *gin.Context) {
		print("List")
		list, err := refs.List(context.Background())
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, list)
	}
}
