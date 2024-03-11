import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import { AuthProvider } from 'auth-ts-sdk'
import App from "./App";

const APP_ID_TEST = "22f8fa89-acdb-4379-b924-1b0c7fbd5931"
const SERVER_ADDRESS = "http://localhost:9999"

ReactDOM.createRoot(document.getElementById('root')!).render(
  <AuthProvider password={"1234"} serviceURL={SERVER_ADDRESS}>
    <App/>
  </AuthProvider>,
)
