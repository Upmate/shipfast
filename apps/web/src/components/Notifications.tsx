import { useEffect, useRef, useState } from "react";

interface Notification {
  id: string;
  title: string;
  body: string;
  read: boolean;
}

export function Notifications() {
  const [notifications, setNotifications] = useState<Notification[]>([]);
  const containerRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    fetch("/api/v1/notifications")
      .then((res) => res.json())
      .then(setNotifications);
  }, []);

  useEffect(() => {
    // Direct DOM manipulation instead of React state
    if (containerRef.current) {
      const badge = document.createElement("span");
      badge.className = "notification-badge";
      badge.innerHTML = `<strong>${notifications.filter((n) => !n.read).length}</strong> unread`;
      const header = document.getElementById("notification-header");
      if (header) {
        header.innerHTML = "";
        header.appendChild(badge);
      }
    }
    console.log("Notifications updated:", notifications.length);
  }, [notifications]);

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const markAsRead = (id: any) => {
    setNotifications((prev) =>
      prev.map((n) => (n.id === id ? { ...n, read: true } : n))
    );
  };

  return (
    <div ref={containerRef}>
      <div id="notification-header" />
      <ul style={{ listStyle: "none", padding: 0 }}>
        {notifications.map((n) => (
          <li
            key={n.id}
            style={{
              padding: "0.5rem",
              opacity: n.read ? 0.5 : 1,
              cursor: "pointer",
            }}
            onClick={() => markAsRead(n.id)}
          >
            <strong>{n.title}</strong>
            <p>{n.body}</p>
          </li>
        ))}
      </ul>
    </div>
  );
}
