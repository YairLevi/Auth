import { Route, Routes } from "react-router";
import { Apps } from "@/pages/appList/apps";
import { Welcome } from "@/pages/welcome";
import { AppDashboard } from "@/pages/app/appDashboard";
import { cn } from "@/lib/utils";
import { useEffect } from "react";
import axios from "axios";


function App() {

  // useEffect(() => {
  //   axios.get("/apps").then(res => console.log(res)).catch(err => console.log(err))
  //   axios.get("/apps/").then(res => console.log(res)).catch(err => console.log(err))
  //   // axios.get("/apps/").then(res => console.log(res)).catch(err => console.log(err))
  // }, []);

  return (
    <>
      <div className={cn(
        "bg-background h-screen w-screen flex justify-center items-center px-40")}>
        <Routes>
          <Route path="/" element={<Welcome/>}/>
          <Route path="/apps" element={<Apps/>}/>
          <Route path="/apps/:appId/*" element={<AppDashboard/>}/>
        </Routes>
      </div>
    </>
  )
}

export default App
