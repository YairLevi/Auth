import './App.css'
import { Register } from "./register";
import { Login } from "./login";
import { useAuth } from "./lib/context";
import { gapi, loadAuth2 } from 'gapi-script'
import a from './client_secret.json'
import { useEffect } from "react";

function App() {
  const { user, logout, isSignedIn } = useAuth()

  return (
    <div>
      {
        isSignedIn
          ? <div>
            This is a dashboard:
            {JSON.stringify(user)}
            <button onClick={() => logout()}>Logout</button>
          </div>
          : <Login/>
      }
    </div>
  )
}

export default App
