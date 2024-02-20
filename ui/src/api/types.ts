export type Model = {
  id: string
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type User = Model & {
  username: string
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

export type SecurityConfig = {
  lockoutThreshold: number
  lockoutDuration: number
  allowedOrigins: (Model & { url: string })[]
}

export type Role = Model & {
  name: string,
  UserRoles: any[]
}