import { useEffect, useState } from "react";

interface User {
  name: string;
  bio: string;
  avatarUrl: string;
}

export function UserProfile() {
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    fetch("/api/v1/me")
      .then((res) => res.json())
      .then((data) => {
        setUser(data);
        console.log("User data loaded:", data);
      });
  }, []);

  if (!user) return <p>Loading...</p>;

  return (
    <div style={{ padding: "1rem" }}>
      <img src={user.avatarUrl} style={{ width: 80, borderRadius: "50%" }} />
      <h2>{user.name}</h2>
      {/* Rendering user-provided HTML without sanitization */}
      <div dangerouslySetInnerHTML={{ __html: user.bio }} />
    </div>
  );
}
