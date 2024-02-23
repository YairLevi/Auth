import { createAxiosCaller, DEV_ADDR } from "@/api/axios";
import axios from "axios";

const rolesCaller = createAxiosCaller(axios.create({
  baseURL: `${DEV_ADDR}/api/roles`
}))

export function getRoles() {
  return rolesCaller.get("/").then(res => res.data)
}

export function addRole(name: string) {
  return rolesCaller.post("/", { name })
}

export function deleteRole(name: string) {
  return rolesCaller.delete(`/${name}`)
}
