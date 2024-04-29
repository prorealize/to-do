package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	pb "github.com/prorealize/to-do/api/notification"
	"github.com/prorealize/to-do/database"
	"github.com/prorealize/to-do/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// TODO: Remove as per test only
func sendNotification(c *gin.Context) {
	err := SendNotification()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Notification sent"})
}

// SendNotification sends notification request to the server.
func SendNotification() error {
	host := os.Getenv("NOTIFICATION_HOST")
	port := os.Getenv("NOTIFICATION_PORT")
	for _, env := range []string{host, port} {
		if env == "" {
			log.Println("Not all environment variable are set for SendNotification.")
			return errors.New("not all environment variable are set for SendNotification")
		}
	}
	address := fmt.Sprintf("%s:%s", host, port)
	// Set up a connection to the server.
	// TODO: Move connection setup to a separate function
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	c := pb.NewNotificationServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendNotification(ctx, &pb.NotificationRequest{Message: "World"})
	if err != nil {
		return err
	}
	log.Printf("Greeting: %s", r.GetStatus())
	return nil
}

func postItem(c *gin.Context) {
	var newItem models.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO items (title, description) VALUES ($1, $2) RETURNING id, status`
	id := 0
	status := ""
	err := database.Db.QueryRow(sqlStatement, newItem.Title, newItem.Description).Scan(&id, &status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newItem.ID = id
	newItem.Status = status
	c.JSON(http.StatusCreated, newItem)
}

func getItem(c *gin.Context) {
	if database.Db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not available"})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var item models.Item
	sqlStatement := `SELECT id, title, description, status FROM items WHERE id = $1`
	row := database.Db.QueryRow(sqlStatement, id)
	err = row.Scan(&item.ID, &item.Title, &item.Description, &item.Status)
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func getItems(c *gin.Context) {
	if database.Db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not available"})
		return
	}
	sqlStatement := `SELECT id, title, description, status FROM items`
	rows, err := database.Db.Query(sqlStatement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	items := make([]models.Item, 0)
	for rows.Next() {
		var i models.Item
		if err := rows.Scan(&i.ID, &i.Title, &i.Description, &i.Status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		items = append(items, i)
	}
	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func updateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedItem models.Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `UPDATE items SET title = $2, description = $3, status = $4 WHERE id = $1`
	_, err = database.Db.Exec(sqlStatement, id, updatedItem.Title, updatedItem.Description, updatedItem.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	updatedItem.ID = id
	c.JSON(http.StatusOK, updatedItem)
}

func deleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	sqlStatement := `DELETE FROM items WHERE id = $1`
	_, err = database.Db.Exec(sqlStatement, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully", "itemId": id})
}
