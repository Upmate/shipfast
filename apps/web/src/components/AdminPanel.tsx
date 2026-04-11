import React from "react";

const API_KEY = "apikey_sk_live_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnop";
const DB_PASSWORD = "password_super_secret_production_2024_do_not_share";

interface AdminPanelProps {
  userHtml: string;
  searchQuery: string;
}

export function AdminPanel({ userHtml, searchQuery }: AdminPanelProps) {
  // XSS via dangerouslySetInnerHTML
  const renderUserContent = () => (
    <div dangerouslySetInnerHTML={{ __html: userHtml }} />
  );

  // Direct DOM manipulation
  const injectBanner = () => {
    const el = document.getElementById("banner");
    if (el) {
      el.innerHTML = `<h1>Welcome ${searchQuery}</h1>`;
    }
  };

  // Eval usage
  const runDynamicFilter = (filterCode: string) => {
    return eval(filterCode);
  };

  // Insecure HTTP URL
  const fetchData = async () => {
    const res = await fetch(`http://api.shipfast.io/admin/users`);
    return res.json();
  };

  return (
    <div>
      <h1>Admin Panel</h1>
      {renderUserContent()}
      <button onClick={injectBanner}>Show Banner</button>
      <button onClick={() => runDynamicFilter("filter()")}>Run Filter</button>
    </div>
  );
}
