package commands

import (
	"fmt"
	"log"

	"chat-analytics-db-migration/configs"
	"chat-analytics-db-migration/database"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// RunServer creates and returns the run server command
// This command starts a simple HTTP server for testing database connectivity
func RunServer() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Run test HTTP server",
		Long:  "Start a simple HTTP server to test database connectivity and configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Test database connection before starting server
			fmt.Println("[Server] Testing database connection...")
			_, sqlConnection := database.Connection()
			defer sqlConnection.Close()

			// Ping database to ensure connectivity
			if err := sqlConnection.Ping(); err != nil {
				return fmt.Errorf("database connection failed: %w", err)
			}
			fmt.Println("[Server] Database connection successful!")

			// Create Gin router based on environment
			if configs.App.Env == "production" {
				gin.SetMode(gin.ReleaseMode)
			}
			router := gin.Default()

			// Add basic health check endpoint
			router.GET("/health", func(c *gin.Context) {
				// Check if database is still reachable
				if err := sqlConnection.Ping(); err != nil {
					c.JSON(500, gin.H{
						"status": "error",
						"error":  "database unreachable",
					})
					return
				}

				// Return healthy status with basic info
				c.JSON(200, gin.H{
					"status":      "healthy",
					"environment": configs.App.Env,
					"database":    configs.DB.Database,
					"host":        configs.DB.Host + ":" + configs.DB.Port,
				})
			})

			// Add database stats endpoint
			router.GET("/db-stats", func(c *gin.Context) {
				stats := sqlConnection.Stats()
				c.JSON(200, gin.H{
					"max_open_connections": stats.MaxOpenConnections,
					"open_connections":     stats.OpenConnections,
					"in_use":               stats.InUse,
					"idle":                 stats.Idle,
					"wait_count":           stats.WaitCount,
					"wait_duration_ms":     stats.WaitDuration.Milliseconds(),
					"max_idle_closed":      stats.MaxIdleClosed,
					"max_idle_time_closed": stats.MaxIdleTimeClosed,
					"max_lifetime_closed":  stats.MaxLifetimeClosed,
				})
			})

			// Start server
			serverAddr := ":" + configs.App.Port
			fmt.Printf("[Server] Starting server on http://localhost%s\n", serverAddr)
			fmt.Printf("[Server] Health check: http://localhost%s/health\n", serverAddr)
			fmt.Printf("[Server] Database stats: http://localhost%s/db-stats\n", serverAddr)

			log.Fatal(router.Run(serverAddr))
			return nil
		},
	}
}
