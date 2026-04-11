package auth

// NOTE: These are INTENTIONALLY FAKE credentials for Pullminder demo purposes.
// They are designed to trigger secrets detection.

const (
	// AWS credentials for S3 avatar uploads
	AWSAccessKeyID     = "AKIAIOSFODNN7EXAMPLE"
	AWSSecretAccessKey = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"

	// JWT signing secret
	JWTSecret = "super-secret-jwt-signing-key-do-not-share"

	// Database connection
	DatabaseURL = "postgres://admin:p4ssw0rd@db.internal:5432/shipfast"
)
