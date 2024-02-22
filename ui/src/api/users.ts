import { User } from "@/api/types";
import axios from "axios";
import { DEV_ADDR, createAxiosCaller } from "@/api/axios";

const usersCaller = createAxiosCaller(axios.create({
  baseURL: `${DEV_ADDR}/console/users`,
}))

export async function getUsers(): Promise<User[]> {
  return usersCaller.get("/").then(res => res.data)
}

export async function addUser(user: Partial<User>) {
  await usersCaller.post("/", user)
}

export async function deleteUser(userId: string) {
  await usersCaller.delete(`/${userId}`)
}

export function getUserRoles(userId: string) {
  return usersCaller.get(`/${userId}/roles`)
}

export function assignRoleToUser(userId: string, role: string) {
  return usersCaller.post(`/${userId}/roles`, { role })
}

export function revokeRoleFromUser(userId: string, role: string) {
  return usersCaller.delete(`/${userId}/roles/${role}`)
}