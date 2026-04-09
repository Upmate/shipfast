import json
import os
from datetime import datetime


def export_to_s3(data: list[dict], bucket: str, prefix: str) -> str:
    """Export processed data to S3."""
    timestamp = datetime.now().strftime("%Y%m%d_%H%M%S")
    key = f"{prefix}/{timestamp}.json"

    # In production, this would use boto3
    output = json.dumps(data, indent=2, default=str)
    print(f"Exporting {len(data)} records to s3://{bucket}/{key}")

    return f"s3://{bucket}/{key}"


def export_to_csv(data: list[dict], output_dir: str) -> str:
    """Export data as CSV file."""
    if not data:
        return ""

    timestamp = datetime.now().strftime("%Y%m%d_%H%M%S")
    filepath = os.path.join(output_dir, f"export_{timestamp}.csv")

    headers = data[0].keys()
    lines = [",".join(headers)]
    for row in data:
        lines.append(",".join(str(row.get(h, "")) for h in headers))

    with open(filepath, "w") as f:
        f.write("\n".join(lines))

    print(f"Exported {len(data)} records to {filepath}")
    return filepath
