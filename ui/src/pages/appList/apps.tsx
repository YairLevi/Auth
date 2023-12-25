import { Plus } from "lucide-react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { useRef, useState } from "react";
import { Input } from "@/components/ui/input";
import { useApps } from "@/pages/appList/queries";
import { useNavigate } from "react-router";

type AppCardProps = {
  id: number
  name: string
  createdAt: Date
}

function AppCard({ name, createdAt, id }: AppCardProps) {
  const navigate = useNavigate()
  return (
    <div
      onClick={() => navigate(`/apps/${id}/users`)}
      className="rounded-lg border w-56 h-80 flex flex-col px-8 gap-4 py-8 drop-shadow bg-background hover:drop-shadow-lg transition ease-in-out duration-200 cursor-pointer">
      <h1 className="scroll-m-20 text-xl font-semibold tracking-tight">{name}</h1>
      <div>
        <p className="text-gray-500 text-sm">Created At:</p>
        <p className="text-gray-600 font-medium">{createdAt.toLocaleDateString()}</p>
      </div>
    </div>
  )
}

export function Apps() {
  const [open, setOpen] = useState(false)
  const nameRef = useRef<HTMLInputElement>()
  const { apps, createApp } = useApps()

  return (
    <div className="flex gap-5 flex-wrap justify-center items-center py-20">
      <div
        onClick={() => setOpen(true)}
        className="rounded-lg border w-56 h-80 flex flex-col justify-center items-center px-8 gap-4 py-8 drop-shadow bg-background hover:drop-shadow-lg transition ease-in-out duration-200 cursor-pointer"
      >
        <Plus className="text-gray-400 mb-5" size={30}/>
        <p className="text-sm text-muted-foreground">Create New App</p>
      </div>
      {
        apps.map(app => (
          <AppCard
            key={app.id}
            id={app.id}
            name={app.name}
            createdAt={app.createdAt}
          />
        ))
      }

      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>
              Create New App
            </DialogTitle>
            <DialogDescription>
              Give a name for your new app.
            </DialogDescription>
          </DialogHeader>
          <Input placeholder="e.g. My Test App" ref={nameRef}/>
          <DialogFooter>
            <Button onClick={() => {
              createApp(nameRef.current.value)
              setOpen(false)
            }}>
              Create
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  )
}