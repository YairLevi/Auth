import { useMutation, useQuery, useQueryClient } from "react-query";
import { apiGetOrigins, apiAddOriginToApp } from "@/api/origins";

const GET_ORIGINS = "get-origins"
const ADD_ORIGIN = "add-origin"

export function useOrigins(appId: string) {
  const client = useQueryClient()

  const { data: origins } = useQuery<string[]>({
    queryKey: [GET_ORIGINS, appId],
    queryFn: () => apiGetOrigins(appId),
    initialData: [],
  })

  const { mutate: addOrigin } = useMutation({
    mutationKey: [ADD_ORIGIN, appId],
    mutationFn: (origin: string) => apiAddOriginToApp(origin, appId),
    onSuccess: () => client.invalidateQueries([GET_ORIGINS, appId])
  })

  return { origins, addOrigin }
}