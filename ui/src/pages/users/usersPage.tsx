import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
import { UserTable } from "@/pages/users/usersTable";
import { useState } from "react";
import { useParams } from "react-router";
import { AddUserDialog } from "@/pages/users/dialogs/addUserDialog";

export function Users() {
  const { appId: id } = useParams()
  const appId = Number(id)
  const [openAddUser, setOpenAddUser] = useState(false)

  return (
    <div>
      <Button onClick={() => setOpenAddUser(true)}>
        Add User<Plus className="ml-2 h-5 w-5"/>
      </Button>
      <UserTable appId={appId}/>

      <AddUserDialog open={openAddUser} setOpen={setOpenAddUser}/>
    </div>
  )
}