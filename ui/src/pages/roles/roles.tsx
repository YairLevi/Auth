import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { useRoles } from "@/pages/roles/queries";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { MoreHorizontal } from "lucide-react";
import * as React from "react";

export function Roles() {
  const { getRoles, deleteRole } = useRoles()


  return (
    <div className="mx-auto max-w-3xl">
      <header className="mb-14">
        <h1 className="text-xl font-semibold">Roles</h1>
        <p>Define and assign roles for users, to allow protection of resources.</p>
      </header>
      <p className="text-muted-foreground text-sm text-center mb-3">A list of your existing roles.</p>
      <Table className="rounded-md border border-separate">
        <TableHeader>
          <TableRow>
            <TableHead>Role</TableHead>
            <TableHead>Created At</TableHead>
            <TableHead>Assigned Users</TableHead>
            <TableHead className="w-[2rem]"></TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {!getRoles.isLoading && getRoles.data?.map(role => (
            <TableRow key={role.id}>
              <TableCell>{role.name}</TableCell>
              <TableCell>{role.createdAt.toLocaleDateString()}</TableCell>
              <TableCell>{role.UserRoles.length}</TableCell>
              <TableCell>
                <DropdownMenu>
                  <DropdownMenuTrigger asChild>
                    <Button variant="ghost" className="h-8 w-8 p-0">
                      <span className="sr-only">Open menu</span>
                      <MoreHorizontal className="h-4 w-4"/>
                    </Button>
                  </DropdownMenuTrigger>
                  <DropdownMenuContent align="end">
                    <DropdownMenuItem
                      className="text-red-700 hover:!text-red-700"
                      onClick={() => deleteRole.mutate(role.name)}
                    >
                      Delete Role
                    </DropdownMenuItem>
                  </DropdownMenuContent>
                </DropdownMenu>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  )
}