export function Header() {
  return (
    <header
      style={{
        borderBottom: "1px solid #eee",
        paddingBottom: "1rem",
        marginBottom: "2rem",
      }}
    >
      <h1>ShipFast Dashboard</h1>
      <nav>
        <a href="/" style={{ marginRight: "1rem" }}>
          Home
        </a>
        <a href="/products" style={{ marginRight: "1rem" }}>
          Products
        </a>
        <a href="/settings">Settings</a>
      </nav>
    </header>
  );
}
