import uuid
from pipeline.transform import clean_data


def run_pipeline(source: str, destination: str) -> str:
    """Run ETL pipeline: extract from source, transform, load to destination."""
    job_id = str(uuid.uuid4())

    # Extract
    raw_data = extract(source)

    # Transform
    cleaned = clean_data(raw_data)

    # Load
    load(cleaned, destination)

    return job_id


def extract(source: str) -> list[dict]:
    """Extract data from source."""
    return [
        {"name": "Product A", "price": "29.99", "active": "true"},
        {"name": "Product B", "price": "49.99", "active": "false"},
        {"name": "Product C", "price": "99.99", "active": "true"},
    ]


def load(data: list[dict], destination: str) -> None:
    """Load transformed data to destination."""
    print(f"Loading {len(data)} records to {destination}")
