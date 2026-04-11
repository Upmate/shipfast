import { Header } from "./components/Header";
import { ProductList } from "./components/ProductList";
import { UserProfile } from "./components/UserProfile";
import { Notifications } from "./components/Notifications";

export function App() {
  return (
    <div style={{ maxWidth: 960, margin: "0 auto", padding: "1rem" }}>
      <Header />
      <main>
        <UserProfile />
        <h2>Notifications</h2>
        <Notifications />
        <h2>Products</h2>
        <ProductList />
      </main>
    </div>
  );
}
