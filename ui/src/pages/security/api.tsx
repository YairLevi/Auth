import { Mutation, useMutation, useQuery, useQueryClient } from "react-query";
import { SecuritySettings } from "@/api/types";
import axios from "axios";
import { useParams } from "react-router";

const settingsKey = (appId: string) => ['get-security-settings', appId]

export function useSecuritySettings() {
  const { appId } = useParams()

  return useQuery<SecuritySettings>({
    queryKey: settingsKey(appId),
    queryFn: async () => {
      try {
        const res = await axios.get<SecuritySettings>(`/apps/${appId}/security/`)
        return res.data
      } catch (e) {
        return null
      }
    },
  })
}

export function useSetLockoutDuration() {
  const { appId } = useParams()
  const client = useQueryClient()

  return useMutation({
    mutationFn: (duration: number) => axios.post(`/apps/${appId}/security/lockout/duration`, { duration }),
    onSuccess: () => client.invalidateQueries(settingsKey(appId))
  })
}

export function useSetLockoutThreshold() {
  const { appId } = useParams()
  const client = useQueryClient()

  return useMutation({
    mutationFn: (threshold: number) => axios.post(`/apps/${appId}/security/lockout/threshold`, { threshold }),
    onSuccess: () => client.invalidateQueries(settingsKey(appId))
  })
}

export function useSetSessionKey() {
  const { appId } = useParams()
  const client = useQueryClient()

  return useMutation({
    mutationFn: (key: string) => axios.post(`/apps/${appId}/security/session`, { key }),
    onSuccess: () => client.invalidateQueries(settingsKey(appId))
  })
}

export function useAddOrigin() {
  const { appId } = useParams()
  const client = useQueryClient()

  return useMutation({
    mutationFn: (origin: string) => axios.post(`/apps/${appId}/security/origins`, { origin }),
    onSuccess: () => client.invalidateQueries(settingsKey(appId))
  })
}