import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import * as React from "react";
import { useEffect, useRef, useState } from "react";
import { ChevronDown, Trash, X } from 'lucide-react'
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
import { cn } from "@/lib/utils";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Label } from "@/components/ui/label";


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

function EmailFilters() {
  const { securityConfig, addEmailFilter, removeEmailFilter } = useSecurityConfig()
  const [open, setOpen] = useState(false)
  const [openAdd, setOpenAdd] = useState(false)
  const emailRef = useRef<HTMLInputElement>(null)
  const [isWhitelist, setIsWhitelist] = useState(false)

  function add() {
    const email = emailRef.current.value
    addEmailFilter({ email, isWhitelist })
    setOpenAdd(false)
  }
  const emailFilters = securityConfig.emailFilters || []

  return (
    <div className="overflow-hidden mb-20">
      <div className="flex justify-between items-center gap-3">
        <div>
          <h1 className="font-semibold">Allow or Deny Email Patterns</h1>
          <p className="text-sm">Email filters to allow only specific addresses, or deny some.</p>
          <p className="text-xs text-muted-foreground"> Notice that, if you provide at least one whitelisted address, the service will deny all other emails that do not comply with the allowed patterns.</p>
        </div>
        <div className="flex gap-4 items-center">
          <Button onClick={() => setOpenAdd(true)}>
            Add Email
          </Button>
          <ChevronDown
            className={cn(
              "mr-4 transition-all duration-200",
              open ? "rotate-180" : ""
            )}
            size={28}
            onClick={() => setOpen(!open)}
          />
        </div>
      </div>
      <div className={cn("mt-5 transition-transform duration-300", open ? "h-fit" : "h-0")}>
        <Table className="rounded-md border border-separate">
          <TableHeader className="sticky top-0">
            <TableRow>
              <TableHead>Pattern</TableHead>
              <TableHead>Allow Status</TableHead>
              <TableHead>Created At</TableHead>
              <TableHead className="w-[2rem] max-w-[2rem]"></TableHead>
            </TableRow>
          </TableHeader>
          <TableBody className="overflow-auto">
            {!!emailFilters && emailFilters.map(email => (
              <TableRow key={email.id}>
                <TableCell>{email.email}</TableCell>
                <TableCell>{email.isWhitelist ? 'whitelist' : 'denied'}</TableCell>
                <TableCell>{email.createdAt.toLocaleDateString()}</TableCell>
                <TableCell>
                  <Trash size={24}
                         className="p-1 hover:bg-gray-100 rounded-md"
                         onClick={() => removeEmailFilter(email.id)}/>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>

      <Dialog open={openAdd} onOpenChange={setOpenAdd}>
        <DialogContent className="w-fit">
          <DialogHeader>
            <DialogTitle>
              Allow New Origin
            </DialogTitle>
            <DialogDescription>
              Allow another origin to make auth requests to your server.
            </DialogDescription>
          </DialogHeader>
          <Input ref={emailRef}/>
          <div className="flex gap-2 items-center w-fit">
            <Input type='radio' radioGroup="emails" id="whitelist" checked={isWhitelist} onChange={e => setIsWhitelist(true)}/>
            <Label htmlFor="whitelist" className="mr-4">Whitelist</Label>
            <Input type='radio' radioGroup="emails" id="blacklist" checked={!isWhitelist} onChange={e => setIsWhitelist(false)}/>
            <Label htmlFor="blacklist">Blacklist</Label>
          </div>
          <DialogFooter>
            <Button onClick={add}>Add</Button>
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
      <EmailFilters/>
      <Separator className="w-full mb-14"/>
      <TokenCustomization/>
    </div>
  )
}