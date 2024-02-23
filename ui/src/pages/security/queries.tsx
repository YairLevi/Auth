import { useMutation, useQuery, useQueryClient } from "react-query";
import { SecurityConfig } from "@/api/types";
import * as api from "@/api/security"

const keys = {
  getSecurityConfig: "get-security-config",
  setLockoutDuration: "set-lockout-duration",
  setLockoutThreshold: "set-lockout-threshold",
  setSessionKey: "set-session-key",
  addOrigin: "add-origin",
  removeOrigin: "remove-origin",
}

export function useSecurityConfig() {
  const client = useQueryClient()

  const { data: securityConfig } = useQuery<SecurityConfig>({
    queryKey: [keys.getSecurityConfig],
    queryFn: () => api.getSecurityConfig(),
  })

  const { mutate: setLockoutDuration } = useMutation({
    mutationKey: [keys.setLockoutDuration],
    mutationFn: api.setLockoutDuration,
    onSuccess: () => client.invalidateQueries([keys.getSecurityConfig])
  })

  const { mutate: setLockoutThreshold } = useMutation({
    mutationKey: [keys.setLockoutThreshold],
    mutationFn: api.setLockoutThreshold,
    onSuccess: () => client.invalidateQueries([keys.getSecurityConfig])
  })

  const { mutate: setSessionKey } = useMutation({
    mutationKey: [keys.setSessionKey],
    mutationFn: api.setSessionKey,
    onSuccess: () => client.invalidateQueries([keys.getSecurityConfig])
  })

  const { mutate: addOrigin } = useMutation({
    mutationKey: [keys.addOrigin],
    mutationFn: api.addOrigin,
    onSuccess: () => client.invalidateQueries([keys.getSecurityConfig])
  })

  const { mutate: removeOrigin } = useMutation({
    mutationKey: [keys.removeOrigin],
    mutationFn: api.removeOrigin,
    onSuccess: () => client.invalidateQueries([keys.getSecurityConfig])
  })

  return {
    setLockoutDuration,
    setLockoutThreshold,
    setSessionKey,
    addOrigin,
    securityConfig,
    removeOrigin
  }
}