import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import { AuthProvider } from "./lib/context";
import App from "./App";

const APP_ID_TEST = "22f8fa89-acdb-4379-b924-1b0c7fbd5931"
const SERVER_ADDRESS = "http://localhost:9999"

ReactDOM.createRoot(document.getElementById('root')!).render(
  <AuthProvider appId={APP_ID_TEST} serviceURL={SERVER_ADDRESS}>
    <App/>
  </AuthProvider>,
)
