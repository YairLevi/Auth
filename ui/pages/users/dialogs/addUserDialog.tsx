import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle
} from "@/components/ui/dialog";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { ChangeEvent, useState } from "react";
import { User } from "@/api/types";
import { useUsers } from "@/pages/users";

type AddUserDialogProps = {
  open: boolean
  setOpen: (flag: boolean) => void
}

export function AddUserDialog({ open, setOpen }: AddUserDialogProps) {
  const [info, setInfo] = useState<Partial<User>>({})
  const { createUser } = useUsers()

  function updateField(key: string, e: ChangeEvent<HTMLInputElement>) {
    setInfo(prev => {
      return {
        ...prev,
        [key]: e.target.value
      }
    })
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Add User</DialogTitle>
          <DialogDescription>Create a new user entry.</DialogDescription>
        </DialogHeader>
        <div className="flex gap-10">
          <div className="w-full">
            <Label>First Name</Label>
            <Input onChange={e => updateField("firstName", e)}/>
          </div>
          <div className="w-full">
            <Label>Last Name</Label>
            <Input onChange={e => updateField("lastName", e)}/>
          </div>
        </div>
        <div>
          <Label>Email</Label>
          <Input onChange={e => updateField("email", e)}/>
        </div>
        <div>
          <Label>Password</Label>
          <Input onChange={e => updateField("passwordHash", e)}/>
        </div>
        <div>
          <Label>Phone Number</Label>
          <Input onChange={e => updateField("phoneNumber", e)}/>
        </div>
        <div>
          <Label>Birthday</Label>
          <Input type="date" onChange={e => updateField("birthday", e)}/>
        </div>
        <DialogFooter>
          <Button variant="ghost" onClick={() => setOpen(false)}>Cancel</Button>
          <Button onClick={() => {
            createUser(info)
            setOpen(false)
          }}>
            Add
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}