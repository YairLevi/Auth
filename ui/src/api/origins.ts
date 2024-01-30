import axios from "axios";

export async function apiAddOriginToApp(origin: string, appId: string) {
  let originToAdd = origin
  if (origin.endsWith("/")) {
    originToAdd = originToAdd.slice(0, -1)
  }

  try {
    const res = await axios.post(`/apps/${appId}/origins/`, { origin: originToAdd })
  } catch (e) {

  }
}

export async function apiGetOrigins(appId: string): Promise<string[]> {
  try {
    const res = await axios.get(`/apps/${appId}/origins/`)
    return res.data
  } catch (e) {
    return []
  }
}