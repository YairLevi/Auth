import axios from "axios";
import { App } from "@/api/types";

export async function getApps(): Promise<App[]> {
  try {
    const res = await axios.get<App[]>('/apps/')
    return res.data
  } catch (e) {
    return []
  }
}

export async function addApp(name: string): Promise<boolean> {
  try {
    const res = await axios.post('/apps/', { name })
    return res.status == axios.HttpStatusCode.Created
  } catch (e) {
    return false
  }
}

export async function deleteApp(appId: string) {
  try {
    const res = await axios.delete(`/apps/${appId}`)
    return res.status == axios.HttpStatusCode.Ok
  } catch (e) {
    return false
  }
}