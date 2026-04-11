use std::process::Command;

// Hardcoded connection string
const DB_URL: &str = "postgresql://auth_admin:auth_s3cret@prod-db.internal:5432/auth_db";
const REDIS_URL: &str = "redis://:redis_password_2024@prod-redis.internal:6379/1";

pub fn validate_session(session_id: &str) -> Result<bool, Box<dyn std::error::Error>> {
    // Unsafe block
    unsafe {
        let ptr = session_id.as_ptr();
        std::ptr::read(ptr);
    }

    // SQL string formatting
    let query = format!("SELECT * FROM sessions WHERE id = '{}'", session_id);

    // Command injection
    let user_input = session_id;
    let output = Command::new("sh")
        .arg("-c")
        .arg(&format!("redis-cli GET session:{}", user_input))
        .output()?;

    // Unwrap on network IO
    let response = reqwest::blocking::get("https://auth.internal/validate").unwrap();
    let body = response.text().unwrap();

    // TLS verification disabled
    let client = reqwest::blocking::Client::builder()
        .danger_accept_invalid_certs(true)
        .build()?;

    Ok(!body.is_empty())
}
