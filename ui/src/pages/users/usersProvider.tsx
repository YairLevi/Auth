import { createContext, PropsWithChildren, useContext } from "react";
import { User } from "@/api/types";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { addUser, deleteUser, getUsers } from "@/api/users";

const keys = {
  getUsers: "get-users",
  addUser: "add-user",
  deleteUser: "delete-user"
}

type ContextExports = {
  users: User[]
  createUser: (user: Partial<User>) => void
  removeUser: (userId: string) => void
}

const UsersContext = createContext<ContextExports>({} as ContextExports)

export function useUsers() {
  return useContext(UsersContext)
}

export function UsersProvider({ children }: PropsWithChildren) {
  const client = useQueryClient()

  const { data: users } = useQuery<User[]>({
    queryKey: [keys.getUsers],
    queryFn: () => getUsers(),
    initialData: []
  })

  const { mutate: createUser } = useMutation({
    mutationKey: [keys.addUser],
    mutationFn: (user: Partial<User>) => addUser(user),
    onSuccess: () => client.invalidateQueries([keys.getUsers])
  })

  const { mutate: removeUser } = useMutation({
    mutationKey: [keys.deleteUser],
    mutationFn: (userId: string) => deleteUser(userId),
    onSuccess: () => client.invalidateQueries([keys.getUsers])
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