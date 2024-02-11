import { createContext, PropsWithChildren, useContext, useEffect, useMemo, useState } from "react";
import { Exports, Provider, User } from "./types"
import axios, { Axios } from "axios"
import { setupAxiosAndGetEndpoints } from "./endpoints";

const AuthContext = createContext<Exports>({} as Exports)

export function useAuth() {
  return useContext(AuthContext)
}

type AuthContextPros = {
  appId: string
  serviceURL: string
  password: string
}

export function AuthProvider({ children, serviceURL, appId, password }: PropsWithChildren & AuthContextPros) {
  const [user, setUser] = useState<User>()
  const isSignedIn = !!user
  const endpoints = setupAxiosAndGetEndpoints(serviceURL, appId)

  axios.defaults.headers["X-App-ID"] = appId
  axios.defaults.headers["X-PW"] = password
  const instance = useMemo(() => axios.create({ withCredentials: true }), [])

  useEffect(() => {
    (function () {
      instance.get(endpoints.loginCookie)
        .then(res => setUser(res.data))
        .catch(err => console.log(err))
    })()
  }, [])

  function login(email: string, password: string) {
    axios.post(endpoints.loginEmailPassword, { email, password })
      .then(res => setUser(res.data))
      .catch(err => console.log(err))
  }

  function logout() {
    axios.post(endpoints.logout)
      .then(_ => setUser(null))
      .catch(err => console.log(err))
  }

  function loginWithProvider(provider: Provider) {
    window.location.href = `${serviceURL}/api/${provider}/login`
  }

  const value = {
    user, login, logout, isSignedIn, loginWithProvider
  }

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  )
}