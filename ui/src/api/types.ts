export type Model = {
  id: number
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type App = Model & {
  name: string
}

export type User = Model & {
  firstName: string
  lastName: string
  email: string
  passwordHash: string
  phoneNumber: string
  lastLogin: Date
  birthday: Date
}

export type SocialState = {
  [provider: string]: {
    enabled: boolean
    clientID: string
    clientSecret: string
  }
}