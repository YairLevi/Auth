import { useRef } from "react";
import axios from "axios";

export function Register() {
  const emailRef = useRef<HTMLInputElement>()
  const passwordRef = useRef<HTMLInputElement>()

  function onClick() {
    const email = emailRef.current.value
    const password = passwordRef.current.value

    axios.post("http://localhost:9999/api/register", {
      email, password, appId: 1
    })
      .then(res => console.log(res))
      .catch(err => console.log(err))
  }

  return (
    <div>
      <input ref={emailRef} placeholder="email"/>
      <input ref={passwordRef} placeholder="password"/>
      <button onClick={onClick}>Register</button>
    </div>
  )
}