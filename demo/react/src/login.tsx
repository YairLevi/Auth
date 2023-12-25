import { useRef } from "react";
import { useAuth } from "./lib/context";
import { GoogleLogin, GoogleLoginResponse, GoogleLogout } from 'react-google-login'
import google from './client_secret.json'
import axios from "axios";

export function Login() {
  const { login, loginWithGoogle, loginWithGithub } = useAuth()
  const emailRef = useRef<HTMLInputElement>()
  const passwordRef = useRef<HTMLInputElement>()
  return (
    <div style={{ display: "flex", flexDirection: "column", gap: "0.5rem" }}>
      <input ref={emailRef} placeholder="email" style={{
        borderRadius: "0.25rem",
        border: "none",
        padding: "0.3rem 0.5rem"
      }}/>
      <input ref={passwordRef} placeholder="password" style={{
        borderRadius: "0.25rem",
        border: "none",
        padding: "0.3rem 0.5rem"
      }}/>
      <button onClick={() => login(emailRef.current.value, passwordRef.current.value)}>Login</button>
      <button onClick={() => loginWithGoogle()}>Sign in google</button>
      <button onClick={() => loginWithGithub()}>Sign in Github</button>
    </div>
  )
}