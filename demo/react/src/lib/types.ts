export type Model = {
  id: number
  createdAt: Date
  updatedAt: Date
  deletedAt: Date
}

export type User = Model & {
  id: string
  firstName: string
  lastName: string
  email: string
  passwordHash: string
  phoneNumber: string
  lastLogin: Date
  birthday: Date
}

export type Exports = {
  user: User,
  isSignedIn: boolean
  login: (email: string, password: string) => void
  logout: () => void
  loginWithGoogle: () => void
  loginWithGithub: () => void
}