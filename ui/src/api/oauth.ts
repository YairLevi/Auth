import axios from "axios";
import { SocialState } from "@/api/types";
import { createAxiosCaller, DEV_ADDR } from "@/api/axios";

const oauthCaller = createAxiosCaller(axios.create({
  baseURL: `${DEV_ADDR}/console/oauth`
}))

export type ProviderDTO = {
  provider: string
  clientId: string
  clientSecret: string
}

export function enableProvider(provider: string) {
  return oauthCaller.post(`/${provider}`)
}

export function disableProvider(provider: string) {
  return oauthCaller.delete(`/${provider}`)
}

export function updateProviderCredentials(credentials: ProviderDTO) {
  return oauthCaller.put(`/${credentials.provider}`, {
    clientId: credentials.clientId,
    clientSecret: credentials.clientSecret,
  })
}

export async function getOAuthState(): Promise<SocialState> {
  const res = await oauthCaller.get(`/providers`)
  return res.data
}