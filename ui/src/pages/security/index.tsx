import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useEffect, useRef, useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle
} from "@/components/ui/dialog";
import { Separator } from "@/components/ui/separator";
import {
  useAddOrigin,
  useSecuritySettings,
  useSetLockoutDuration,
  useSetLockoutThreshold,
  useSetSessionKey
} from "@/pages/security/api";


function LockoutDuration() {
  const { data: settings } = useSecuritySettings()
  const { mutate: setDuration } = useSetLockoutDuration()
  const [edit, setEdit] = useState(false)
  const ref = useRef<HTMLInputElement>(null)

  useEffect(() => {
    if (ref.current) {
      ref.current.value = `${settings.lockoutDuration}`
    }
  }, []);

  function save() {
    setDuration(parseInt(ref.current.value))
    setEdit(false)
  }

  return (
    <div className="flex mb-20 justify-between gap-10">
      <div className="flex flex-col gap-3">
        <h1 className="font-semibold">Lockout Duration</h1>
        <p className="text-sm">The amount of time the account remains locked after reaching the account lockout
          threshold.</p>
      </div>
      <div>
        <div className="flex gap-2">
          <Input
            ref={ref}
            disabled={!edit}
            className="min-w-[11rem] max-w-[11rem]"
          />
          <Button
            className="min-w-[3.5rem] max-w-[3rem]"
            onClick={edit ? save : () => setEdit(true)}
          >
            {edit ? "Save" : "Edit"}
          </Button>
        </div>
        <p className="mt-1 text-xs text-muted-foreground">Value in seconds. Defaults to 30.</p>
      </div>
    </div>
  )
}

function LockoutThreshold() {
  const { data: settings } = useSecuritySettings()
  const { mutate: setThreshold } = useSetLockoutThreshold()
  const [edit, setEdit] = useState(false)
  const ref = useRef<HTMLInputElement>(null)

  useEffect(() => {
    if (ref.current) {
      ref.current.value = `${settings.lockoutThreshold}`
    }
  }, []);

  function save() {
    setThreshold(parseInt(ref.current.value))
    setEdit(false)
  }

  return (
    <div className="flex mb-20 justify-between gap-10">
      <div className="flex flex-col gap-3">
        <h1 className="font-semibold">Account Lockout Threshold</h1>
        <p className="text-sm">The number of allowed failed login attempts before the account is locked.</p>
      </div>
      <div className="flex gap-2">
        <Input
          ref={ref}
          disabled={!edit}
          className="min-w-[11rem] max-w-[11rem]"
        />
        <Button
          className="min-w-[3.5rem] max-w-[3rem]"
          onClick={edit ? save : () => setEdit(true)}
        >
          {edit ? "Save" : "Edit"}
        </Button>
      </div>
    </div>
  )
}

function AllowedOrigins() {
  const originRef = useRef<HTMLInputElement>()
  const [open, setOpen] = useState(false)

  const { data: settings } = useSecuritySettings()
  const { mutate: addOrigin } = useAddOrigin()

  function onClickAdd() {
    addOrigin(originRef.current.value)
    setOpen(false)
  }

  return (
    <div className="flex mb-20 items-top justify-between gap-10">
      <div className="flex flex-col gap-3">
        <h1 className="font-semibold">Allowed Origins</h1>
        <p className="text-sm">List of trusted domains authorized to make requests to your service.</p>
      </div>
      <div className="w-60">
        <p className="text-sm font-semibold text-muted-foreground">
          Currently allowed origins:
        </p>
        {
          settings.allowedOrigins?.map(org => <p key={org.id} className="text-sm my-2 text-muted-foreground">{org.url}</p>)
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
          <DialogFooter>
            <Button onClick={onClickAdd}>Add</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  )
}

function TokenCustomization() {
  const [open, setOpen] = useState(false)
  const keyRef = useRef<HTMLInputElement>(null)

  const { mutate: setSessionKey } = useSetSessionKey()

  function onclick() {
    setSessionKey(keyRef.current.value)
    setOpen(false)
  }

  return (
    <div className="flex mb-20 justify-between gap-10">
      <div className="flex flex-col gap-3">
        <h1 className="font-semibold">Token</h1>
        <p className="text-sm">
          Choose a signing key for session tokens.
          <br/>For security reasons, the current one cannot be seen.
        </p>
      </div>
      <div className="w-60">
        <Button
          variant="outline"
          className="w-full"
          onClick={() => setOpen(true)}
        >
          Change Key
        </Button>
      </div>

      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent className="w-fit">
          <DialogHeader>
            <DialogTitle>
              New Token Key
            </DialogTitle>
            <DialogDescription>
              Put your new session token key here.
              It is recommended to have a relatively random key.
            </DialogDescription>
          </DialogHeader>
          <Input ref={keyRef}/>
          <DialogFooter>
            <Button onClick={onclick}>Set New Key</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  )
}

export function Security() {
  const { data: settings } = useSecuritySettings()

  return settings != null && (
    <div className="mx-auto max-w-3xl">
      <LockoutThreshold/>
      <LockoutDuration/>
      <AllowedOrigins/>
      <Separator className="w-full mb-14"/>
      <TokenCustomization/>
    </div>
  )
}