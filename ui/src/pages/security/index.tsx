import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useEffect, useRef, useState } from "react";
import { X } from 'lucide-react'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle
} from "@/components/ui/dialog";
import { Separator } from "@/components/ui/separator";
import { useSecurityConfig } from "@/pages/security/queries";
import { Model } from "@/api/types";


function LockoutDuration() {
  const { securityConfig, setLockoutDuration } = useSecurityConfig()
  const [edit, setEdit] = useState(false)
  const ref = useRef<HTMLInputElement>(null)

  useEffect(() => {
    if (ref.current) {
      ref.current.value = `${securityConfig.lockoutDuration}`
    }
  }, []);

  function save() {
    setLockoutDuration(parseInt(ref.current.value))
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
  const { securityConfig, setLockoutThreshold } = useSecurityConfig()
  const [edit, setEdit] = useState(false)
  const ref = useRef<HTMLInputElement>(null)

  useEffect(() => {
    if (ref.current) {
      ref.current.value = `${securityConfig.lockoutThreshold}`
    }
  }, []);

  function save() {
    setLockoutThreshold(parseInt(ref.current.value))
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

function OriginTab({ origin }: { origin: Model & { url: string } }) {
  const { removeOrigin } = useSecurityConfig()
  return (
    <div className="flex justify-between items-center py-1">
      <p key={origin.id} className="text-sm font-semibold my-2 text-muted-foreground">{origin.url}</p>
      <X
        size={26}
        className="p-1 mr-4 hover:bg-gray-100 transition-colors rounded-md"
        onClick={() => removeOrigin(origin.id)}
      />
    </div>
  )
}

function AllowedOrigins() {
  const originRef = useRef<HTMLInputElement>()
  const [open, setOpen] = useState(false)

  const { securityConfig, addOrigin } = useSecurityConfig()

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
        <div className="h-28 overflow-auto">
          {
            securityConfig.allowedOrigins?.map(org => (
              <>
                <OriginTab origin={org} key={org.id}/>
              </>
            ))
          }
          {
            securityConfig.allowedOrigins?.length == 0 &&
              <p className="text-sm my-2 text-muted-foreground py-5 text-center">No allowed origins are set.</p>
          }
        </div>
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

  const { setSessionKey } = useSecurityConfig()

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
  const { securityConfig } = useSecurityConfig()

  return !!securityConfig && (
    <div className="mx-auto max-w-3xl">
      <LockoutThreshold/>
      <LockoutDuration/>
      <AllowedOrigins/>
      <Separator className="w-full mb-14"/>
      <TokenCustomization/>
    </div>
  )
}