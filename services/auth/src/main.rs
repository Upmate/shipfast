use actix_web::{web, App, HttpResponse, HttpServer};
use serde::{Deserialize, Serialize};

#[derive(Serialize)]
struct HealthResponse {
    status: String,
    service: String,
}

#[derive(Deserialize)]
struct TokenRequest {
    username: String,
    password: String,
}

#[derive(Serialize)]
struct TokenResponse {
    token: String,
    expires_in: u64,
}

async fn health() -> HttpResponse {
    HttpResponse::Ok().json(HealthResponse {
        status: "ok".to_string(),
        service: "shipfast-auth".to_string(),
    })
}

async fn validate_token(req: web::Json<TokenRequest>) -> HttpResponse {
    if req.username.is_empty() || req.password.is_empty() {
        return HttpResponse::BadRequest().json(serde_json::json!({
            "error": "username and password required"
        }));
    }

    HttpResponse::Ok().json(TokenResponse {
        token: "validated".to_string(),
        expires_in: 3600,
    })
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    println!("Auth service starting on :8090");
    HttpServer::new(|| {
        App::new()
            .route("/health", web::get().to(health))
            .route("/validate", web::post().to(validate_token))
    })
    .bind("0.0.0.0:8090")?
    .run()
    .await
}
