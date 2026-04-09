def clean_data(records: list[dict]) -> list[dict]:
    """Clean and transform raw records."""
    cleaned = []
    for record in records:
        cleaned.append(
            {
                "name": record["name"].strip(),
                "price": float(record["price"]),
                "active": record["active"].lower() == "true",
            }
        )
    return cleaned
