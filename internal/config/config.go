package config

var (
	USER_COLLECTION         = "users"
	QUESTION_COLLECTION     = "questions"
	CSV_DIR                 = "C:\\go workspace\\CODESAGE\\csv"
	GPT_API_ENDPOINT        = "https://api.openai.com/v1/chat/completions"
	GPT_MODEL               = "gpt-4o"
	LEETCODE_API            = "https://Leetcode.com/graphql/"
	RECENT_SUBMISSION_LIMIT = 10
	DB_USER                 = "postgres"
	DB_PASSWORD             = "password"
	DB_NAME                 = "codesage"
	DB_HOST                 = "localhost"
	DB_PORT                 = "5432"
	LOG_FILE                = "/Users/pravin/Desktop/CODESAGE/logs.log"
	SECRET_KEY              = []byte("secret-key")
	PORT                    = ":8080"
)
