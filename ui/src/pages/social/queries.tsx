import { useMutation, useQuery, useQueryClient } from "react-query";
import { useParams } from "react-router";
import { useState } from "react";
import {
  apiEnableProvider,
  apiDisableProvider,
  getOAuthState,
  ProviderDTO,
  apiUpdateProviderCredentials
} from "@/api/oauth";
import { SocialState } from "@/api/types"

const GET_OAUTH_STATE = "get-oauth-state"
const ENABLE_PROVIDER = "enable-provider"
const DISABLE_PROVIDER = "disable-provider"
const UPDATE_PROVIDER = "update-provider"

export function useSocial(): {
  oauthState: SocialState
  enableProvider: (providerName: string) => void
  disableProvider: (providerName: string) => void
  updateProvider: (dto: ProviderDTO) => void
} {
  const { appId } = useParams()
  const [oauthState, setOAuthState] = useState<SocialState>({})
  const client = useQueryClient()

  useQuery<SocialState>({
    queryKey: [GET_OAUTH_STATE, appId],
    queryFn: () => getOAuthState(appId),
    onSuccess: data => setOAuthState(data)
  })

  const { mutate: enableProvider } = useMutation({
    mutationKey: [ENABLE_PROVIDER, appId],
    mutationFn: (providerName: string) => apiEnableProvider(appId, providerName),
    onSuccess: () => {
      client.invalidateQueries([GET_OAUTH_STATE, appId])
    },
    onError: err => {
      console.log(err)
    },
  })

  const { mutate: disableProvider } = useMutation({
    mutationKey: [DISABLE_PROVIDER, appId],
    mutationFn: (providerName: string) => apiDisableProvider(appId, providerName),
    onSuccess: () => {
      client.invalidateQueries([GET_OAUTH_STATE, appId])
    },
    onError: err => {
      console.log(err)
    },
  })

  const { mutate: updateProvider } = useMutation({
    mutationKey: [UPDATE_PROVIDER, appId],
    mutationFn: (dto: ProviderDTO) => apiUpdateProviderCredentials(appId, dto),
    onSuccess: () => client.invalidateQueries([GET_OAUTH_STATE, appId]),
    onError: () => {}
  })

  return {
    enableProvider,
    disableProvider,
    updateProvider,
    oauthState,
  }
}