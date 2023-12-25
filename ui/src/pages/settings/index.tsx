import { EyeIcon } from "lucide-react";
import { useState } from "react";
import { useParams } from "react-router";

export function Settings() {
  const [visible, setVisible] = useState(false)
  const { appId } = useParams()
  const dummyJWTKey = "Hello world!"

  return (
    <div className="flex flex-col gap-10">
      <div className="flex justify-around items-center">
        <p className="font-semibold">
          Application ID:
        </p>
        <div className="bg-gray-100 rounded flex px-5 py-3 gap-5">
          <p className="text-gray-600 text-sm font-semibold tracking-tight">
            {visible ? appId : 'X'.repeat(50)}
          </p>
          <EyeIcon size={20} className="text-gray-600" onClick={() => setVisible(!visible)}/>
        </div>
      </div>
      <div className="flex justify-around items-center">
        <p className="font-semibold">
          JWT secret key:
        </p>
        <div className="bg-gray-100 rounded flex px-5 py-3 gap-5">
          <p className="text-gray-600 text-sm font-semibold tracking-tight">
            {visible ? appId : 'X'.repeat(50)}
          </p>
          <EyeIcon size={20} className="text-gray-600" onClick={() => setVisible(!visible)}/>
        </div>
      </div>
    </div>
  )
}