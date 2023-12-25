import { User } from "@/api/types";
import axios from "axios";

export async function getUsers(appId: string): Promise<User[]> {
  try {
    const res = await axios.get(`/apps/${appId}/users/`)
    return res.data
  } catch (e) {
    return []
  }
}

export async function addUser(appId: string, user: Partial<User>) {
  try {
    const res = await axios.post(`/apps/${appId}/users/`, user)
  } catch (e) {

  }
}

export async function deleteUser(appId: string, userId: number) {
  try {
    const res = await axios.delete(`/apps/${appId}/users/${userId}`)
    console.log(res)
  } catch (e) {
    console.log(e)
  }
}