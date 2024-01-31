import { Route, Routes, useMatch, useNavigate } from "react-router"
import { Users, UsersProvider } from "@/pages/users";
import { SocialConnections } from "@/pages/social";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { DoorOpen, KeyRound, LockIcon, LucideIcon, UsersIcon } from "lucide-react";
import { Security } from "@/pages/security";
import { ReactNode } from "react";

type SidebarItemProps = {
  text: string
  path: string
  Icon?: LucideIcon
}

type Route = {
  name: string
  icon: LucideIcon,
  path: string,
  element: ReactNode
}

const routes: Route[] = [
  {
    name: "Users",
    icon: UsersIcon,
    path: "users",
    element: <>
      <UsersProvider>
        <Users/>
      </UsersProvider>
    </>
  },
  {
    name: "Social Connections",
    icon: KeyRound,
    path: "social",
    element: <SocialConnections/>
  },
  {
    name: "Security",
    icon: LockIcon,
    path: "security",
    element: <Security/>
  },
]

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
      {
        routes.map(r => (
          <SidebarItem key={`sidebar-item${r.path}`} text={r.name} path={r.path} Icon={r.icon}/>
        ))
      }
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
          {
            routes.map(r => (
              <Route key={`route${r}`} path={`/${r.path}`} element={r.element}/>
            ))
          }
        </Routes>
      </div>
    </div>
  )
}