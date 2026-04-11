import os
import pickle
import subprocess
import yaml

# Hardcoded credentials
DB_PASSWORD = "password_postgres_super_secret_2024_production"
API_SECRET = "api_secret_key_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghij1234"
REDIS_URL = "redis://:s3cretRedis@prod-redis.internal:6379/0"

def run_etl(source_table: str, target_table: str):
    """Process data from source to target with SQL injection vulnerability."""

    # SQL injection via f-string
    query = f"SELECT * FROM {source_table} WHERE status = 'active' AND created_at > '{os.environ.get('START_DATE')}'"

    # Subprocess with shell=True
    subprocess.run(f"psql -c \"{query}\" -d shipfast", shell=True)

    # os.system with concatenation
    os.system("echo 'Processing " + source_table + "' >> /var/log/etl.log")

    # Eval execution
    transform_code = os.environ.get("TRANSFORM_FUNC", "lambda x: x")
    transform = eval(transform_code)

    return transform(query)


def load_config(config_path: str):
    """Load config with unsafe deserialization."""
    with open(config_path, 'rb') as f:
        # Unsafe pickle deserialization
        config = pickle.loads(f.read())

    # Also unsafe YAML load
    with open(config_path + '.yml') as f:
        yaml_config = yaml.load(f)

    return {**config, **yaml_config}
