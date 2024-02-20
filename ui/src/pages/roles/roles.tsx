import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { useRoles } from "@/pages/roles/queries";

export function Roles() {
  const { getRoles } = useRoles()

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
          </TableRow>
        </TableHeader>
        <TableBody>
          {!getRoles.isLoading && getRoles.data?.map(role => (
            <TableRow key={role.id}>
              <TableCell>{role.name}</TableCell>
              <TableCell>{role.createdAt.toLocaleDateString()}</TableCell>
              <TableCell>{role.UserRoles.length}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  )
}