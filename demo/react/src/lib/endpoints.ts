import axios from "axios"
import { parseISO } from 'date-fns';


export function setupAxiosAndGetEndpoints(serviceURL: string, appID: string) {
  let baseURL = serviceURL
  if (!serviceURL.endsWith("/")) {
    baseURL += "/"
  }
  axios.defaults.baseURL = baseURL + "api"

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

  return {
    registerEmailPassword: "/register",
    loginEmailPassword: "/login",
    loginCookie: "/login",
    logout: "/logout",

    googleLogin: `/${appID}/google/login`,
    githubLogin: `/${appID}/github/login`,
  }
}