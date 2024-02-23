import { useMutation, useQuery, useQueryClient } from "react-query";
import * as api from "@/api/roles";
import { Role } from "@/api/types";

const keys = {
  getRoles: "get-roles",
  addRole: "add-role",
  deleteRole: 'delete-role',
  getUserRoles: 'get-user-roles',
  assignRoleToUser: "assign-role-to-user",
  revokeRoleFromUser: 'revoke-role-from-user'
}

export function useRoles() {
  const client = useQueryClient()
  const getRoles = useQuery<Role[]>(keys.getRoles, api.getRoles, {
    initialData: []
  })
  const addRole = useMutation(keys.addRole, api.addRole, {
    onSuccess: () => client.invalidateQueries(keys.getRoles)
  })
  const deleteRole = useMutation(keys.deleteRole, api.deleteRole, {
    onSuccess: () => client.invalidateQueries(keys.getRoles)
  })

  return { getRoles, addRole, deleteRole }
}