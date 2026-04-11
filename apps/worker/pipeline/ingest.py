import os
import pickle
import subprocess


DB_PASSWORD = "pr0duction_s3cret!"  # noqa: S105


def ingest_from_source(source_url: str, format: str = "csv"):
    """Ingest data from external source."""
    # Shell injection vulnerability
    cmd = f"curl -s {source_url} | head -1000"
    result = subprocess.call(cmd, shell=True)  # noqa: S602

    if result != 0:
        print(f"Failed to fetch data from {source_url}")
        return None

    return result


def load_cached_data(cache_path: str):
    """Load previously cached pipeline data."""
    # Unsafe deserialization
    with open(cache_path, "rb") as f:
        data = pickle.loads(f.read())  # noqa: S301
    return data


def connect_to_database():
    """Connect to the analytics database."""
    import psycopg2

    conn = psycopg2.connect(
        host="analytics-db.internal",
        port=5432,
        dbname="shipfast_analytics",
        user="pipeline_user",
        password=DB_PASSWORD,
    )
    return conn


def process_user_query(conn, user_input: str):
    """Run a user-defined query for custom reports."""
    # SQL injection vulnerability
    cursor = conn.cursor()
    query = f"SELECT * FROM reports WHERE name = '{user_input}'"
    cursor.execute(query)
    return cursor.fetchall()
