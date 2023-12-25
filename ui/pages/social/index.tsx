import { Switch } from "@/components/ui/switch";
import { useRef, useState } from "react";
import { providerIcons } from "@/pages/social/providerIcons";
import { AlertCircle, ExternalLink, SlidersHorizontal } from "lucide-react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useSocial } from "@/pages/social/queries";
import { HoverCard, HoverCardContent, HoverCardTrigger } from "@/components/ui/hover-card";

export function SocialConnections() {
  const [open, setOpen] = useState(false)
  const { oauthState, enableProvider, updateProvider, disableProvider } = useSocial()
  const [selectedProvider, setSelectedProvider] = useState("")
  const [clientID, setClientID] = useState("")
  const [clientSecret, setClientSecret] = useState("")

  function openProviderCredentials(provider: string) {
    setSelectedProvider(provider)
    setClientID(oauthState[provider].clientID)
    setClientSecret(oauthState[provider].clientSecret)
    setOpen(true)
  }

  function saveUpdatedCredentials() {
    updateProvider({
      provider: selectedProvider,
      clientId: clientID,
      clientSecret: clientSecret,
    })
    setOpen(false)
  }

  return (
    <div className="h-full max-w-[40rem] mx-auto">
      <h3 className="scroll-m-20 text-2xl font-semibold tracking-tight">
        Social Connections (OAuth2)
      </h3>
      <p className="my-2 text-muted-foreground">
        Use famous auth providers.
      </p>
      <div className="[&:not(:last-child)]:border-b">
        {
          Object.keys(oauthState).map(provider => {
            const emptyCredentials = oauthState[provider].clientID == "" || oauthState[provider].clientSecret == ""

            return (
              <div key={provider} className="w-full flex justify-between items-center py-4">
                <div className="w-full flex gap-5">
                  <img src={providerIcons[provider]} alt="google" className="w-6 h-6"/>
                  <p className="first-letter:capitalize">{provider}</p>
                </div>
                {
                  oauthState[provider].enabled &&
                    <>
                      {
                        emptyCredentials &&
                          <HoverCard openDelay={0} closeDelay={0}>
                              <HoverCardTrigger>
                                  <AlertCircle size={20} className="text-red-500 mr-3"/>
                              </HoverCardTrigger>
                              <HoverCardContent>
                                  <h6 className="font-semibold mb-1">Empty Credentials</h6>
                                  <p className="text-sm">
                                      A <strong>Client ID</strong> and <strong>Client Secret</strong> must be set,
                                      otherwise your users will not be able to use this provider.
                                  </p>
                                  <a
                                      className="flex gap-1 text-sm text-blue-600 hover:underline hover:cursor-pointer items-center"
                                  >
                                      Learn More <ExternalLink size={12}/>
                                  </a>
                              </HoverCardContent>
                          </HoverCard>
                      }
                        <SlidersHorizontal
                            className="text-gray-600 hover:text-black transition-colors mr-3"
                            size={20}
                            onClick={() => openProviderCredentials(provider)}
                        />
                    </>
                }
                <Switch
                  checked={oauthState[provider].enabled}
                  onCheckedChange={(checked) => checked ? enableProvider(provider) : disableProvider(provider)}
                />
              </div>
            )
          })
        }
      </div>

      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent onOpenAutoFocus={e => e.preventDefault()}>
          <DialogHeader>
            <DialogTitle>
              Credentials
            </DialogTitle>
            <DialogDescription>
              Add your custom credentials for Google.
            </DialogDescription>
          </DialogHeader>
          <p className="-mb-2 text-sm font-semibold">Client ID:</p>
          <Input
            value={clientID}
            onChange={e => setClientID(e.target.value)}
          />
          <p className="-mb-2 text-sm font-semibold">Client Secret:</p>
          <Input
            value={clientSecret}
            onChange={e => setClientSecret(e.target.value)}
          />
          <DialogFooter>
            <Button variant="ghost" onClick={() => setOpen(false)}>Cancel</Button>
            <Button onClick={() => saveUpdatedCredentials()}>
              Save
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  )
}