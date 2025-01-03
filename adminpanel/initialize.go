package adminpanel

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/GoAdminGroup/go-admin/adapter"
	ginadapter "github.com/GoAdminGroup/go-admin/adapter/gin"
	"github.com/GoAdminGroup/go-admin/engine"
	gaConfig "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	_ "github.com/GoAdminGroup/themes/adminlte"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var AdminEngine *engine.Engine

func InitializeGoAdmin(r *gin.Engine) error {
	// PostgreSQL DSN
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=football_chat sslmode=disable"

	// Initialize database connection for manual queries
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	defer db.Close()

	// GoAdmin config
	cfg := &gaConfig.Config{
		Debug:     true,
		Theme:     "adminlte",
		UrlPrefix: "admin",
		IndexUrl:  "/admin/dashboard",
		Language:  "en",
		Store: gaConfig.Store{
			Path:   "./sessions",
			Prefix: "session",
		},
		Databases: gaConfig.DatabaseList{
			"default": {
				Driver:       "postgresql",
				Dsn:          dsn,
				MaxIdleConns: 50,
				MaxOpenConns: 150,
			},
		},
	}

	AdminEngine = engine.Default()

	// Gin adapter and admin plugin
	ginAdapter := ginadapter.Gin{BaseAdapter: adapter.BaseAdapter{}}
	adminPlugin := admin.NewAdmin()

	// Add config & plugins
	if err := AdminEngine.
		AddConfig(cfg).
		AddAdapter(&ginAdapter).
		AddPlugins(adminPlugin).
		Use(r); err != nil {
		return err
	}

	log.Println("GoAdmin initialized successfully at /admin")

	// Create or update the default admin user
	createOrUpdateDefaultAdmin(db, "myadmin", "my-secret-password")

	return nil
}

// createOrUpdateDefaultAdmin inserts or updates a user in the `goadmin_users` table.
func createOrUpdateDefaultAdmin(db *sql.DB, username, plainPassword string) {
	// Hash the plain password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return
	}

	// SQL query to insert or update the admin user
	query := `
        INSERT INTO goadmin_users (username, password, name, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (username) DO UPDATE 
          SET password = EXCLUDED.password,
              updated_at = EXCLUDED.updated_at
    `

	// Execute the query
	_, execErr := db.Exec(query, username, string(hashedPassword), "My Admin", time.Now(), time.Now())
	if execErr != nil {
		log.Println("Error creating/updating default admin user:", execErr)
	} else {
		log.Printf("Default admin user created or updated. username=%s\n", username)
	}
}
