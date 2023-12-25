import { createContext, PropsWithChildren, useContext, useState } from "react";
import { User } from "@/api/types";
import { useParams } from "react-router";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { addUser, deleteUser, getUsers } from "@/api/users";

const GET_USERS_QUERY_KEY = "get-users"
const ADD_USER_MUTATION_KEY = "add-user"
const DELETE_USER_MUTATION_KEY = "delete-user"

type ContextExports = {
  users: User[]
  createUser: (user: Partial<User>) => void
  removeUser: (userId: number) => void
}

const UsersContext = createContext<ContextExports>({} as ContextExports)

export function useUsers() {
  return useContext(UsersContext)
}

export function UsersProvider({ children }: PropsWithChildren) {
  const appId = useParams().appId
  const client = useQueryClient()
  const [users, setUsers] = useState<User[]>([])

  useQuery<User[]>({
    queryKey: [GET_USERS_QUERY_KEY, appId],
    queryFn: () => getUsers(appId),
    onSuccess: users => {
      setUsers(users)
    }
  })

  const { mutate: createUser } = useMutation({
    mutationKey: [ADD_USER_MUTATION_KEY, appId],
    mutationFn: (user: Partial<User>) => addUser(appId, user),
    onSuccess: () => client.invalidateQueries([GET_USERS_QUERY_KEY, appId])
  })

  const { mutate: removeUser } = useMutation({
    mutationKey: [DELETE_USER_MUTATION_KEY, appId],
    mutationFn: (userId: number) => deleteUser(appId, userId),
    onSuccess: () => client.invalidateQueries([GET_USERS_QUERY_KEY, appId])
  })

  const value = {
    users,
    createUser,
    removeUser,
  }

  return (
    <UsersContext.Provider value={value}>
      {children}
    </UsersContext.Provider>
  )
}