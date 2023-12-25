import { Route, Routes, useMatch, useNavigate } from "react-router"
import { Users, UsersProvider } from "@/pages/users";
import { SocialConnections } from "@/pages/social";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { DoorOpen, KeyRound, LockIcon, LucideIcon, SettingsIcon, SquareAsteriskIcon, UsersIcon } from "lucide-react";
import { Settings } from "@/pages/settings";
import { Security } from "@/pages/security";

type SidebarItemProps = {
  text: string
  path: string
  Icon?: LucideIcon
}

function SidebarItem({ text, path, Icon }: SidebarItemProps) {
  const match = useMatch('/apps/:appId/:lastPart')
  const isOnItemPage = match.params.lastPart == path
  const navigate = useNavigate()

  return (
    <Button
      variant="ghost"
      className={cn("flex justify-start gap-3", isOnItemPage && "bg-gray-200")}
      onClick={() => navigate(`./${path}`, { relative: "path" })}
    >
      <Icon size={20}/>{text}
    </Button>

  )
}

export function Sidebar() {
  const navigate = useNavigate()

  return (
    <div className="h-full min-w-[15rem] max-w-[15rem] w-full border-r mr-24 flex flex-col px-4 gap-1">
      <SidebarItem text="Users" path="users" Icon={UsersIcon}/>
      <SidebarItem text="Social Connections" path="social" Icon={KeyRound}/>
      <SidebarItem text="Settings" path="settings" Icon={SettingsIcon}/>
      <SidebarItem text="Security" path="security" Icon={LockIcon}/>
      <Button
        onClick={() => navigate("/apps")}
        className="flex gap-3 justify-start mt-auto"
      >
        <DoorOpen size={20}/>
        Exit App
      </Button>
    </div>
  )
}

export function AppDashboard() {
  return (
    <div className="overflow-hidden flex h-5/6 w-full">
      <Sidebar/>
      <div className="w-full overflow-auto">
        <Routes>
          <Route path="/users" element={
            <UsersProvider>
              <Users/>
            </UsersProvider>
          }/>
          <Route path="/settings" element={<Settings/>}/>
          <Route path="/social" element={<SocialConnections/>}/>
          <Route path="/security" element={<Security/>}/>
        </Routes>
      </div>
    </div>
  )
}