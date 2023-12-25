import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

export function Security() {
  return (
    <div className="mx-auto max-w-3xl">
      <div className="flex mb-20 items-center justify-between gap-10">
        <div className="flex flex-col gap-3">
          <h1 className="font-semibold">Account Lockout Threshold</h1>
          <p className="text-sm">The number of allowed failed login attempts before the account is locked.</p>
        </div>
        <Input className="min-w-[15rem] max-w-[15rem]"/>
      </div>
      <div className="flex mb-20 items-center justify-between gap-10">
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
      <div className="flex mb-20 items-center justify-between gap-10">
        <div className="flex flex-col gap-3">
          <h1 className="font-semibold">Allowed Origins:</h1>
          <p className="text-sm">List of trusted domains authorized to make requests to your service.</p>
        </div>
        <div>
          <Input className="min-w-[15rem] max-w-[15rem]"></Input>
          <Button>Add</Button>
        </div>
      </div>
    </div>
  )
}