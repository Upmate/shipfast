import { Header } from "./components/Header";
import { ProductList } from "./components/ProductList";

export function App() {
  return (
    <div style={{ maxWidth: 960, margin: "0 auto", padding: "1rem" }}>
      <Header />
      <main>
        <h2>Products</h2>
        <ProductList />
      </main>
    </div>
  );
}
