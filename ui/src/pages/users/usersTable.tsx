import * as React from "react"
import {
  ColumnDef,
  ColumnFiltersState,
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  SortingState,
  useReactTable,
  VisibilityState,
} from "@tanstack/react-table"
import { ArrowUpDown, ChevronDown, MoreHorizontal } from "lucide-react"

import { Button } from "@/components/ui/button"
import { Checkbox } from "@/components/ui/checkbox"
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Input } from "@/components/ui/input"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow, } from "@/components/ui/table"
import { type User } from '@/api/types'
import { deleteUser } from "@/api/users";
import { useUsers } from "@/pages/users/index";
import { useSearchParams } from "react-router-dom"
import { useEffect } from "react";


export const columns: ColumnDef<User>[] = [
  {
    id: "select",
    header: ({ table }) => (
      <Checkbox
        checked={
          table.getIsAllPageRowsSelected() ||
          (table.getIsSomePageRowsSelected() && "indeterminate") ||
          false
        }
        onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
      />
    ),
    cell: ({ row }) => (
      <Checkbox
        className="mr-5"
        checked={row.getIsSelected()}
        onCheckedChange={(value) => row.toggleSelected(!!value)}
        aria-label="Select row"
      />
    ),
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "email",
    header: ({ column }) => {
      return (
        <span
          className="min-w-[7rem] max-w-[10rem] flex cursor-pointer select-none hover:text-black"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          Email
          <ArrowUpDown className="ml-2 h-4 w-4"/>
        </span>
      )
    },
    cell: ({ row }) => (
      <div className="font-semibold">{row.getValue("email")}</div>
    ),
  },
  {
    accessorKey: "firstName",
    header: ({ column }) => {
      return (
        <span
          className="min-w-[7rem] max-w-[10rem] flex cursor-pointer select-none hover:text-black"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          First Name
          <ArrowUpDown className="ml-2 h-4 w-4"/>
        </span>
      )
    },
    cell: ({ row }) => <p>{row.getValue("firstName")}</p>
  },
  {
    accessorKey: "lastName",
    header: ({ column }) => {
      return (
        <span
          className="min-w-[7rem] max-w-[10rem] flex cursor-pointer select-none hover:text-black"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          Last Name
          <ArrowUpDown className="ml-2 h-4 w-4"/>
        </span>
      )
    },
    cell: ({ row }) => <p>{row.getValue("lastName")}</p>
  },
  {
    accessorKey: "createdAt",
    header: ({ column }) => {
      return (
        <span
          className="min-w-[7rem] max-w-[10rem] flex cursor-pointer select-none hover:text-black"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          Created At
          <ArrowUpDown className="ml-2 h-4 w-4"/>
        </span>
      )
    },
    cell: ({ row }) => <p>{row.getValue<Date>("createdAt").toLocaleDateString()}</p>
  },

  {
    accessorKey: "passwordHash",
    header: ({ column }) => {
      return (
        <span
          className="min-w-[7rem] max-w-[10rem] flex cursor-pointer select-none hover:text-black"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          Password Hash
          <ArrowUpDown className="ml-2 h-4 w-4"/>
        </span>
      )
    },
    cell: ({ row }) => <div>{row.getValue("passwordHash")}</div>,
  },
  {
    accessorKey: "id",
    header: () => <></>,
    cell: () => <></>
  }
]

type UserTableProps = {
  appId: number
}

export function UserTable({ appId }: UserTableProps) {
  const [searchParams, setSearchParams] = useSearchParams()
  useEffect(() => {
    setSearchParams({
      email: "",
      firstName: "",
      lastName: "",
      createdAt: "",
    })
  }, []);
  const [sorting, setSorting] = React.useState<SortingState>([])
  const [columnFilters, setColumnFilters] = React.useState<ColumnFiltersState>([])
  const [columnVisibility, setColumnVisibility] = React.useState<VisibilityState>({})
  const [rowSelection, setRowSelection] = React.useState({})
  const { users, removeUser } = useUsers()

  const table = useReactTable({
    data: users,
    columns,
    onSortingChange: setSorting,
    onColumnFiltersChange: setColumnFilters,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    onColumnVisibilityChange: setColumnVisibility,
    onRowSelectionChange: setRowSelection,
    state: {
      sorting,
      columnFilters,
      columnVisibility,
      rowSelection,
    },
  })

  return (
    <div className="w-full">
      <div className="flex items-center py-4 gap-4">
        <Input
          placeholder="Filter emails..."
          value={(table.getColumn("email")?.getFilterValue() as string) ?? ""}
          onChange={(event) =>
            table.getColumn("email")?.setFilterValue(event.target.value)
          }
          className="max-w-sm"
        />
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline" className="ml-auto">
              Columns <ChevronDown className="ml-2 h-4 w-4"/>
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            {table
              .getAllColumns()
              .filter((column) => column.getCanHide())
              .filter((column) => column.id != "id")
              .map((column) => {
                return (
                  <DropdownMenuCheckboxItem
                    key={column.id}
                    className="capitalize"
                    checked={searchParams.has(column.id)}
                    onCheckedChange={(value) => {
                      column.toggleVisibility(!!value)
                      setSearchParams(prev => {
                        const newParams = prev
                        if (!value) {
                          newParams.delete(column.id)
                        } else {
                          newParams.set(column.id, "")
                        }
                        return newParams
                      })
                    }}
                  >
                    {column.id}
                  </DropdownMenuCheckboxItem>
                )
              })}
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow className="hover:bg-background" key={headerGroup.id}>
                {headerGroup.headers
                  .filter(header => searchParams.has(header.id) || header.id == "select")
                  .map((header) => {
                  return (
                    <TableHead key={header.id}>
                      {header.isPlaceholder
                        ? null
                        : flexRender(
                          header.column.columnDef.header,
                          header.getContext()
                        )}
                    </TableHead>
                  )
                })}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row) => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && "selected"}
                >
                  {row.getVisibleCells()
                    .filter(cell => searchParams.has(cell.column.id) || cell.column.id == "select")
                    .map((cell) => (
                    <TableCell key={cell.id}>
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext()
                      )}
                    </TableCell>
                  ))}
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
                          onClick={() => removeUser(row.getValue("id"))}
                        >
                          Delete User
                        </DropdownMenuItem>
                      </DropdownMenuContent>
                    </DropdownMenu>
                  </TableCell>
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell
                  colSpan={columns.length}
                  className="h-24 text-center"
                >
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
      <div className="flex items-center justify-end space-x-2 py-4">
        <div className="flex-1 text-sm text-muted-foreground">
          {table.getFilteredSelectedRowModel().rows.length} of{" "}
          {table.getFilteredRowModel().rows.length} row(s) selected.
        </div>
        <div className="space-x-2">
          <Button
            variant="outline"
            size="sm"
            onClick={() => table.previousPage()}
            disabled={!table.getCanPreviousPage()}
          >
            Previous
          </Button>
          <Button
            variant="outline"
            size="sm"
            onClick={() => table.nextPage()}
            disabled={!table.getCanNextPage()}
          >
            Next
          </Button>
        </div>
      </div>
    </div>
  )
}
