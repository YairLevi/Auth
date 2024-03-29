import { createAxiosCaller, DEV_ADDR } from "@/api/axios";
import axios from "axios";

const securityCaller = createAxiosCaller(axios.create({
  baseURL: `${DEV_ADDR}/console/security`
}))

export function getSecurityConfig() {
  return securityCaller.get("/").then(res => res.data)
}

export function setLockoutThreshold(threshold: number) {
  return securityCaller.put("/lockout/threshold", { threshold })
}

export function setLockoutDuration(duration: number) {
  return securityCaller.put("/lockout/duration", { duration })
}

export function setSessionKey(sessionKey: string) {
  return securityCaller.put("/session", { sessionKey })
}

export function addOrigin(origin: string) {
  return securityCaller.put("/origins", { origin })
}

export function removeOrigin(originId: string) {
  return securityCaller.delete(`/origins/${originId}`)
}

export function addEmailFilter(args: { email: string, isWhitelist: boolean }) {
  return securityCaller.post("/emails", args)
}

export function removeEmailFilter(emailId: string) {
  return securityCaller.delete(`/emails/${emailId}`)
}