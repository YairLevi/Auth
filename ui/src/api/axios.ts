import axios, { Axios } from "axios";
import { parseISO } from "date-fns";

export const DEV_ADDR = 'http://localhost:9999'
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

export function createAxiosCaller(ax: Axios) {
  ax.interceptors.response.use((rep) => {
    handleDates(rep.data);
    return rep;
  });
  ax.interceptors.request.use((rep) => {
    handleDates(rep.data)
    return rep
  })
  return ax
}