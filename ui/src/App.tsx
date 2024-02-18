import { Route, Routes } from "react-router";
import { Welcome } from "@/pages/welcome";
import { cn } from "@/lib/utils";
import { Dashboard } from "@/pages/dashboard";


function App() {

  return (
    <>
      <div className={cn(
        "bg-background h-screen w-screen flex justify-center items-center px-40")}>
        <Routes>
          <Route path="/" element={<Welcome/>}/>
          <Route path="/console/*" element={<Dashboard/>}/>
        </Routes>
      </div>
    </>
  )
}

export default App
