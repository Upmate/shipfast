from pipeline.transform import clean_data


def test_clean_data_converts_types():
    raw = [{"name": " Test Product ", "price": "19.99", "active": "true"}]
    result = clean_data(raw)
    assert len(result) == 1
    assert result[0]["name"] == "Test Product"
    assert result[0]["price"] == 19.99
    assert result[0]["active"] is True


def test_clean_data_inactive_product():
    raw = [{"name": "Inactive", "price": "0.00", "active": "false"}]
    result = clean_data(raw)
    assert result[0]["active"] is False
