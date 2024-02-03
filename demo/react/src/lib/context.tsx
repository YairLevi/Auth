import { createContext, PropsWithChildren, useContext, useEffect, useState } from "react";
import { Exports, User } from "./types"
import axios from "axios"
import { setupAxiosAndGetEndpoints } from "./endpoints";

const AuthContext = createContext<Exports>({} as Exports)

export function useAuth() {
  return useContext(AuthContext)
}

type AuthContextPros = {
  appId: string
  serviceURL: string
}

export function AuthProvider({ children, serviceURL, appId }: PropsWithChildren & AuthContextPros) {
  const [user, setUser] = useState<User>()
  const isSignedIn = !!user
  const endpoints = setupAxiosAndGetEndpoints(serviceURL, appId)

  axios.defaults.headers["X-App-ID"] = appId

  useEffect(() => {
    (function () {
      axios.get(endpoints.loginCookie, {
        withCredentials: true,
      })
        .then(res => setUser(res.data))
        .catch(err => console.log(err))
    })()
  }, [])

  function login(email: string, password: string) {
    axios.post(endpoints.loginEmailPassword, {
      email, password, appId
    }, {
      withCredentials: true,
    })
      .then(res => {
        console.log(res.data)
        setUser(res.data)
      })
      .catch(err => console.log(err))
  }

  function logout() {
    axios.post(endpoints.logout, { appId, }, { withCredentials: true, })
      .then(_ => setUser(null))
      .catch(err => console.log(err))
  }

  function loginWithGoogle() {
    window.location.href = `http://localhost:9999/api/google/${appId}/login`
  }

  function loginWithGithub() {
    window.location.href = `http://localhost:9999/api/github/${appId}/login`
  }


  const value = {
    user, login, logout, isSignedIn, loginWithGoogle, loginWithGithub
  }

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  )
}