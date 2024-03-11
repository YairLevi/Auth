import './App.css'
import { Login } from "./login";
import { useAuth } from 'auth-ts-sdk'

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
