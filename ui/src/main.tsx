import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import "./index.css"
import { BrowserRouter } from "react-router-dom";
import axios from "axios";
import { parseISO } from 'date-fns';
import { QueryClient, QueryClientProvider } from "react-query";

const DEV_ADDR = 'http://localhost:9999'
axios.defaults.baseURL = DEV_ADDR
const ISODateFormat = /^\d{4}-\d{2}-\d{2}(?:T\d{2}:\d{2}:\d{2}(?:\.\d*)?(?:[-+]\d{2}:?\d{2}|Z)?)?$/;


const isIsoDateString = (value: unknown): value is string => {
  return typeof value === "string" && ISODateFormat.test(value);
};

const handleDates = (data: unknown) => {
  if (isIsoDateString(data)) return parseISO(data);
  if (data === null || data === undefined || typeof data !== "object") return data;

  for (const [key, val] of Object.entries(data)) {
    if (isIsoDateString(val)) data[key] = parseISO(val);
    else if (typeof val === "object") handleDates(val);
  }

  return data
};

axios.interceptors.response.use((rep) => {
  handleDates(rep.data);
  return rep;
});
axios.interceptors.request.use((rep) => {
  handleDates(rep.data)
  return rep
})

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
