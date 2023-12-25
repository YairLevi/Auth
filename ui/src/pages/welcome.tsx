import { useRef, useState } from "react";
import { useNavigate } from "react-router";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import axios from "axios";

export function Welcome() {
  const [error, setError] = useState('')
  // TODO: actually use the address that the user entered, and not a preconfigured one.
  const addressRef = useRef<HTMLInputElement>(null)
  const navigate = useNavigate()

  async function onClick() {
    let res
    try {
      res = await axios.get('/test')
      console.log(res.status)
      if (res.status == axios.HttpStatusCode.Ok) {
        navigate("/apps")
      } else {
        setError("Couldn't connect. Check your service's status.")
      }
    } catch (e) {
      return setError("Couldn't connect. Check your service's status.")
    }
  }

  return (
    <div className="w-96 flex flex-col gap-2 p-4">
      <Label htmlFor="address">Enter Address:</Label>
      <Input placeholder="example.com" id="address" ref={addressRef}/>
      {error && <p className="text-red-700 text-[0.9rem] font-semibold">{error}</p>}
      <Button className="mt-5" onClick={onClick}>Connect</Button>
    </div>
  )
}