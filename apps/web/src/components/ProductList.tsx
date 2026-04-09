import { useEffect, useState } from "react";

interface Product {
  id: string;
  name: string;
  description: string;
  price: number;
  currency: string;
}

export function ProductList() {
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch("/api/v1/products")
      .then((res) => res.json())
      .then((data) => {
        setProducts(data);
        setLoading(false);
      })
      .catch(() => setLoading(false));
  }, []);

  if (loading) return <p>Loading products...</p>;

  return (
    <table style={{ width: "100%", borderCollapse: "collapse" }}>
      <thead>
        <tr>
          <th style={{ textAlign: "left", padding: "0.5rem" }}>Name</th>
          <th style={{ textAlign: "left", padding: "0.5rem" }}>Description</th>
          <th style={{ textAlign: "right", padding: "0.5rem" }}>Price</th>
        </tr>
      </thead>
      <tbody>
        {products.map((p) => (
          <tr key={p.id} style={{ borderTop: "1px solid #eee" }}>
            <td style={{ padding: "0.5rem" }}>{p.name}</td>
            <td style={{ padding: "0.5rem" }}>{p.description}</td>
            <td style={{ padding: "0.5rem", textAlign: "right" }}>
              ${(p.price / 100).toFixed(2)} {p.currency}
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
