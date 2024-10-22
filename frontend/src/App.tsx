import { useEffect } from "react";
import { createHashRouter, RouterProvider } from "react-router-dom";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import myFetch from "./lib/myFetch";
import GroupList from "./component/GroupList";
import ChatWindow from "./component/ChatWindow";

function App() {
  const token = localStorage.getItem("token");
  const router = createHashRouter([
    {
      path: "/",
      element: (
        <>
          <GroupList groups={[]} onGroupChange={() => {}} />
          <ChatWindow currentGroup="general" />
        </>
      )
    }
  ]);

  const initializeGoogle = async () => {
    const clientId = import.meta.env.VITE_GOOGLE_AUTH_CLIENT_ID;
    if (!clientId) {
      console.error("Google client ID not found!");
      return;
    }
    window.google.accounts.id.initialize({
      client_id: clientId,
      callback: onLogin,
      cancel_on_tap_outside: true,
      context: "use"
    });
    const googleButton = document.getElementById("googleButton");
    if (googleButton) {
      window.google.accounts.id.renderButton(googleButton, {
        theme: "outline",
        size: "large",
        type: "standard"
      });
    } else {
      console.error("Google button element not found!");
    }
  };
  const onLogin = async (response: Google.CredentialResponse) => {
    const url = "/auth/google";
    const data: Google.CredentialResponse = { credential: response.credential };
    const result = await myFetch<{
      email: string;
      name: string;
      token: string;
    }>(url, "POST", data);
    if (result.ok) {
      localStorage.setItem("token", result.data.token);
      localStorage.setItem("email", result.data.email);
      localStorage.setItem("name", result.data.name);
    }
  };

  const checkLogin = async () => {
    if (token) {
      onLogin({ credential: token });
    } else {
      initializeGoogle();
    }
  };

  useEffect(() => {
    checkLogin();
  }, []);

  return token ? (
    <RouterProvider router={router} />
  ) : (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <div id="googleButton" style={{ display: "flex", justifyContent: "center" }}></div>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">Click on the Vite and React logos to learn more</p>
    </>
  );
}

export default App;
