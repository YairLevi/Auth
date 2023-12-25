import axios from "axios";
import { SocialState } from "@/api/types";

export type ProviderDTO = {
  provider: string
  clientId: string
  clientSecret: string
}

export async function apiEnableProvider(appId: string, provider: string) {
  const res = await axios.put(`/apps/${appId}/oauth/${provider}/enable`)
}

export async function apiDisableProvider(appId: string, provider: string) {
  const res = await axios.delete(`/apps/${appId}/oauth/${provider}/disable`)
}

export async function apiUpdateProviderCredentials(appId: string, credentials: ProviderDTO) {
  const res = await axios.put(`/apps/${appId}/oauth/${credentials.provider}/update`, {
    clientId: credentials.clientId,
    clientSecret: credentials.clientSecret,
  })
}

export async function getOAuthState(appId: string): Promise<SocialState> {
  const res = await axios.get(`/apps/${appId}/oauth/providers`)
  return res.data
}