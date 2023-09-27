package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	cus "github.com/Subasri-V/user-role.git/proto"
	"github.com/Subasri-V/user-role.git/services"
)

func main() {
	r := gin.Default()
	conn, err := grpc.Dial("localhost:2000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := cus.NewUserRoleServiceClient(conn)
	r.POST("/adduser", func(c *gin.Context) {
		var request cus.UserRequest

		// Parse incoming JSON
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		pass, err2 := services.HashPassword(request.Password)
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
			return
		}
		request.Password = pass

		// Call the gRPC service
		response, err := client.AddUser(c.Request.Context(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"value": response})
	})

	r.POST("/enable/:name", func(c *gin.Context) {
		name := c.Param("name")
		response, err := client.EnableUser(context.Background(), &cus.Name{
			Name: name,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"value": response})

	})

	r.POST("/disable/:name", func(c *gin.Context) {
		name := c.Param("name")
		response, err := client.DisableUser(context.Background(), &cus.Name{
			Name: name,
		})
		fmt.Println("client",name)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Disable unsuccessful")
		} else {
			c.JSON(http.StatusAccepted, response)
		}
	})

	r.POST("/updaterole/:name",func (c *gin.Context)  {
		name:=c.Param("name")
		var request cus.UpdateRoleRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response,err:=client.UpdateRole(context.Background(),&cus.UpdateRoleRequest{
			Name: name,
			Role: request.Role,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"value": response})		
	})

	r.POST("/associaterole/:name",func (c *gin.Context)  {
		name:=c.Param("name")
		var request cus.AssociateRoleRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response,err:=client.AssociateRole(context.Background(),&cus.AssociateRoleRequest{
			Name: name,
			Role: request.Role,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"value": response})
	})

	r.POST("/remove/:name",func (c *gin.Context)  {
		name:=c.Param("name")
		var request cus.AssociateRoleRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response,err:=client.Remove(context.Background(),&cus.AssociateRoleRequest{
			Name: name,
			Role: request.Role,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"value": response})
	})

	r.POST("/appendarray/:name",func (c *gin.Context)  {
		name:=c.Param("name")
		var request cus.UpdateRoleRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response,err:=client.AppendArray(context.Background(),&cus.UpdateRoleRequest{
			Name: name,
			Role: request.Role,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"value": response})
	})

	r.POST("listrole/:role",func (c *gin.Context)  {
		role:=c.Param("role")
		response,err:=client.ListFeatures(context.Background(),&cus.Role{
			Role: role,
		})
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"value": response})
	})
	r.Run(":8080")
}
