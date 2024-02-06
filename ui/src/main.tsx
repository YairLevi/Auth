import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import "./index.css"
import { BrowserRouter } from "react-router-dom";
import "@/api/axios";
import { QueryClient, QueryClientProvider } from "react-query";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      refetchInterval: 1000 * 10
    }
  }
})

ReactDOM.createRoot(document.getElementById('root')!).render(
  <QueryClientProvider client={queryClient}>
    <BrowserRouter>
      <App/>
    </BrowserRouter>
  </QueryClientProvider>,
)
