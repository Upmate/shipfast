export function LoginForm() {
  return (
    <form
      style={{
        maxWidth: 400,
        margin: "2rem auto",
        padding: "2rem",
        border: "1px solid #e0e0e0",
        borderRadius: 8,
      }}
      onSubmit={(e) => e.preventDefault()}
    >
      <h2 style={{ marginTop: 0 }}>Sign In</h2>
      <div style={{ marginBottom: "1rem" }}>
        <label htmlFor="email" style={{ display: "block", marginBottom: 4 }}>
          Email
        </label>
        <input
          id="email"
          type="email"
          style={{
            width: "100%",
            padding: "0.5rem",
            borderRadius: 4,
            border: "1px solid #ccc",
          }}
        />
      </div>
      <div style={{ marginBottom: "1rem" }}>
        <label htmlFor="password" style={{ display: "block", marginBottom: 4 }}>
          Password
        </label>
        <input
          id="password"
          type="password"
          style={{
            width: "100%",
            padding: "0.5rem",
            borderRadius: 4,
            border: "1px solid #ccc",
          }}
        />
      </div>
      <button
        type="submit"
        style={{
          width: "100%",
          padding: "0.75rem",
          backgroundColor: "#0066ff",
          color: "white",
          border: "none",
          borderRadius: 4,
          cursor: "pointer",
          fontSize: "1rem",
        }}
      >
        Sign In
      </button>
    </form>
  );
}
