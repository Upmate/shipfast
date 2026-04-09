from fastapi import FastAPI

app = FastAPI(title="ShipFast Worker", version="1.0.0")


@app.get("/health")
def health():
    return {"status": "ok", "service": "shipfast-worker"}


@app.post("/jobs/etl")
def trigger_etl(source: str, destination: str):
    from pipeline.etl import run_pipeline

    result = run_pipeline(source, destination)
    return {"status": "queued", "job_id": result}
