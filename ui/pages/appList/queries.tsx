import { useState } from "react";
import { App } from "@/api/types";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { getApps, addApp } from "@/api/apps";

const GET_APPS_QUERY_KEY = "get-appList"
const CREATE_APP_MUTATION_KEY = "create-app"

export function useApps() {
  const [apps, setApps] = useState<App[]>([])
  const client = useQueryClient()

  useQuery<App[]>({
    queryKey: [GET_APPS_QUERY_KEY],
    queryFn: () => getApps(),
    onSuccess: data => setApps(data)
  })

  const createAppMutation = useMutation({
    mutationKey: [CREATE_APP_MUTATION_KEY],
    mutationFn: addApp,
    onSuccess: () => client.invalidateQueries([GET_APPS_QUERY_KEY])
  })
  function createApp(name: string) {
    createAppMutation.mutate(name)
  }

  return { apps, createApp }
}