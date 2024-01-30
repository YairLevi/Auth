import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useParams } from "react-router";
import { useEffect, useRef, useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle
} from "@/components/ui/dialog";
import { useOrigins } from "@/pages/security/queries";

function AllowedOrigins() {
  const { appId } = useParams()
  const originRef = useRef<HTMLInputElement>()
  const [open, setOpen] = useState(false)
  const { origins, addOrigin } = useOrigins(appId)
  const [error, setError] = useState("")

  function onClickAdd() {
    setError("")
    const url = originRef.current.value
    if (!url.startsWith("http://") && !url.startsWith("https://")) {
      return setError("Origin must contain http:// or https://")
    }
    addOrigin(originRef.current.value)
    setOpen(false)
  }

  return (
    <div className="flex mb-20 items-top justify-between gap-10">
      <div className="flex flex-col gap-3">
        <h1 className="font-semibold">Allowed Origins:</h1>
        <p className="text-sm">List of trusted domains authorized to make requests to your service.</p>
      </div>
      <div className="w-60">
        <p className="text-sm font-semibold text-muted-foreground">
          Currently allowed origins:
        </p>
        {
          origins?.map(org => <p className="text-sm my-2 text-muted-foreground">{org["url"] || "<invalid url>"}</p>)
        }
        <Button
          className="mt-2 w-full"
          variant="outline"
          onClick={() => setOpen(true)}
        >
          Add Origin
        </Button>
      </div>

      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent className="w-fit">
          <DialogHeader>
            <DialogTitle>
              Allow New Origin
            </DialogTitle>
            <DialogDescription>
              Allow another origin to make auth requests to your server.
            </DialogDescription>
          </DialogHeader>
          <Input ref={originRef}/>
          {error && <p className="text-sm text-red-500">{error}</p>}
          <DialogFooter>
            <Button onClick={onClickAdd}>Add</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  )
}

export function Security() {
  return (
    <div className="mx-auto max-w-3xl">
      <div className="flex mb-20 justify-between gap-10">
        <div className="flex flex-col gap-3">
          <h1 className="font-semibold">Account Lockout Threshold</h1>
          <p className="text-sm">The number of allowed failed login attempts before the account is locked.</p>
        </div>
        <Input className="min-w-[15rem] max-w-[15rem]"/>
      </div>
      <div className="flex mb-20 justify-between gap-10">
        <div className="flex flex-col gap-3">
          <h1 className="font-semibold">Lockout Duration:</h1>
          <p className="text-sm">The amount of time the account remains locked after reaching the account lockout
            threshold.</p>
        </div>
        <div>
          <Input className="min-w-[15rem] max-w-[15rem]"/>
          <p className="mt-1 text-xs text-muted-foreground">Value in seconds. Defaults to 30.</p>
        </div>
      </div>
      <AllowedOrigins/>
    </div>
  )
}