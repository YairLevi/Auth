import { useMutation, useQuery, useQueryClient } from "react-query";
import * as api from "@/api/oauth";
import { SocialState } from "@/api/types"

const keys = {
  getOAuthState: "get-oauth-state",
  enableProvider: "enable-provider",
  disableProvider: "disable-provider",
  updateProvider: "update-provider",
}

export function useSocial(): {
  oauthState: SocialState
  enableProvider: (providerName: string) => void
  disableProvider: (providerName: string) => void
  updateProvider: (dto: api.ProviderDTO) => void
} {
  const client = useQueryClient()

  const { data: oauthState } = useQuery<SocialState>({
    queryKey: [keys.getOAuthState],
    queryFn: () => api.getOAuthState(),
    initialData: {}
  })

  const { mutate: enableProvider } = useMutation({
    mutationKey: [keys.enableProvider],
    mutationFn: (providerName: string) => api.enableProvider(providerName),
    onSuccess: () => client.invalidateQueries([keys.getOAuthState]),
  })

  const { mutate: disableProvider } = useMutation({
    mutationKey: [keys.disableProvider],
    mutationFn: (providerName: string) => api.disableProvider(providerName),
    onSuccess: () => client.invalidateQueries([keys.getOAuthState]),
  })

  const { mutate: updateProvider } = useMutation({
    mutationKey: [keys.updateProvider],
    mutationFn: (dto: api.ProviderDTO) => api.updateProviderCredentials(dto),
    onSuccess: () => client.invalidateQueries([keys.getOAuthState]),
    onError: () => {}
  })

  return {
    enableProvider,
    disableProvider,
    updateProvider,
    oauthState,
  }
}