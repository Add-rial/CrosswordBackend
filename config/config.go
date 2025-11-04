package config

import (
	"log"
	"os"

	"CrosswordBackend/model"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var(
	DB *gorm.DB
	JwtKey []byte
	AdminKey string
	GoogleClientID string
)

func InitDB(){
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Postgres:", err)
	}

	err = DB.AutoMigrate(&model.User{}, &model.CrosswordAnswer{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate database schema: %v", err)
	}

	log.Println("Database connection and migration successful.")
}

func InitEnv() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment variables")
    }
	JwtKey = []byte(os.Getenv("JWT_SECRET"))
	AdminKey = os.Getenv("ADMIN_KEY")
	GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
}

func GetSchemaConfig() * genai.GenerateContentConfig{
	t := true
	return &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"rows": {Type: genai.TypeInteger},
				"columns": {Type: genai.TypeInteger},
				"grid": {
					Type: genai.TypeArray,
					Items: &genai.Schema{
						Type: genai.TypeArray,
						Items: &genai.Schema{
							Type: genai.TypeObject,
							Properties: map[string]*genai.Schema{
								"isBlank": {Type: genai.TypeBoolean},
								"numberAssociated": {
									Type: genai.TypeInteger,
									Nullable: &t,
								},
							},
						},
					},
				},
			},
		},
	}
}

func GetPrompt() string{
	return `Analyze the attached image of a crossword puzzle.

				Your task is to extract the complete structure of the grid and the list of clues.
				The output MUST be a single JSON object.

				1.  **Rows:** Determine the maximum row size needed to represent the entire puzzle.
				2.  **Columns:** Determine the maximum column size needed to represent the entire puzzle.
				3.  **Grid:** Generate a 2D array of 'Cell' objects based on the grid layout.

				**Cell Structure Rules:**
				* **isBlank:** Set to 'true' if the cell is a black square or separator (cannot accept a letter). Otherwise, set to 'false'.
				* **numberAssociated:** Set to the integer number (e.g., 1, 2, 3) if the cell has a clue number in the top-left corner. If no number is present (most letter cells), set this to 'null'.

				**Target JSON Schema:**

				{
				"rows": Integer
				"columns": Integer 
				"grid": [
					[ /* Row 0 of Cell objects */ ],
					// ...
				],
				}

				Ensure the JSON is perfectly valid and ready for immediate parsing.`
}