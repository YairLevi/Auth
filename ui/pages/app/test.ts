import { User } from "@/api/types";
import { useState } from "react";

const mockUsers: User[] = [
  {
    "id": 1,
    "createdAt": new Date("2023-12-10T12:00:00Z"),
    "deletedAt": new Date("2023-12-10T12:00:00Z"),
    "updatedAt": new Date("2023-12-10T12:00:00Z"),
    "email": "user1@gmail.com",
    "passwordHash": "abc",
    "firstName": "John",
    "lastName": "Doe",
    "birthday": new Date("1990-01-01"),
    "phoneNumber": "123-456-7890",
    lastLogin: new Date("2001-01-01")
  },
  {
    "id": 2,
    "createdAt": new Date("2023-12-10T12:01:00Z"),
    "deletedAt": new Date("2023-12-10T12:01:00Z"),
    "updatedAt": new Date("2023-12-10T12:01:00Z"),
    "email": "user2@gmail.com",
    "passwordHash": "abc",
    "firstName": "Jane",
    "lastName": "Doe",
    "birthday": new Date("1985-05-15"),
    "phoneNumber": "987-654-3210",
    lastLogin: new Date("2001-01-01")
  },
  {
    "id": 3,
    "createdAt": new Date("2023-12-10T12:02:00Z"),
    "deletedAt": new Date("2023-12-10T12:02:00Z"),
    "updatedAt": new Date("2023-12-10T12:02:00Z"),
    "email": "user3@gmail.com",
    "passwordHash": "abc",
    "firstName": "Alice",
    "lastName": "Johnson",
    "birthday": new Date("1992-08-20"),
    "phoneNumber": "555-123-4567",
    lastLogin: new Date("2001-01-01")
  },
  {
    "id": 4,
    "createdAt": new Date("2023-12-10T12:03:00Z"),
    "deletedAt": new Date("2023-12-10T12:03:00Z"),
    "updatedAt": new Date("2023-12-10T12:03:00Z"),
    "email": "user4@gmail.com",
    "passwordHash": "abc",
    "firstName": "Bob",
    "lastName": "Smith",
    "birthday": new Date("1988-12-10"),
    "phoneNumber": "789-321-6540",
    lastLogin: new Date("2001-01-01")
  },
  {
    "id": 5,
    "createdAt": new Date("2023-12-10T12:04:00Z"),
    "deletedAt": new Date("2023-12-10T12:04:00Z"),
    "updatedAt": new Date("2023-12-10T12:04:00Z"),
    "email": "user5@gmail.com",
    "passwordHash": "abc",
    "firstName": "Eva",
    "lastName": "Brown",
    "birthday": new Date("1995-04-25"),
    "phoneNumber": "456-789-0123",
    lastLogin: new Date("2001-01-01")
  },
  {
    "id": 6,
    "createdAt": new Date("2023-12-10T12:05:00Z"),
    "deletedAt": new Date("2023-12-10T12:05:00Z"),
    "updatedAt": new Date("2023-12-10T12:05:00Z"),
    "email": "user6@gmail.com",
    "passwordHash": "abc",
    "firstName": "Chris",
    "lastName": "Johnson",
    "birthday": new Date("1991-06-30"),
    "phoneNumber": "111-222-3333",
    lastLogin: new Date("2001-01-01")
  },
  {
    "id": 7,
    "createdAt": new Date("2023-12-10T12:06:00Z"),
    "deletedAt": new Date("2023-12-10T12:06:00Z"),
    "updatedAt": new Date("2023-12-10T12:06:00Z"),
    "email": "user7@gmail.com",
    "passwordHash": "abc",
    "firstName": "Diana",
    "lastName": "Miller",
    "birthday": new Date("1987-09-12"),
    "phoneNumber": "222-333-4444",
    lastLogin: new Date("2001-01-01")
  },
  {
    "id": 8,
    "createdAt": new Date("2023-12-10T12:07:00Z"),
    "deletedAt": new Date("2023-12-10T12:07:00Z"),
    "updatedAt": new Date("2023-12-10T12:07:00Z"),
    "email": "user8@gmail.com",
    "passwordHash": "abc",
    "firstName": "Frank",
    "lastName": "Williams",
    "birthday": new Date("1993-11-05"),
    "phoneNumber": "333-444-5555",
    lastLogin: new Date("2001-01-01")
  },
  {
    "id": 9,
    "createdAt": new Date("2023-12-10T12:08:00Z"),
    "deletedAt": new Date("2023-12-10T12:08:00Z"),
    "updatedAt": new Date("2023-12-10T12:08:00Z"),
    "email": "user9@gmail.com",
    "passwordHash": "abc",
    "firstName": "Grace",
    "lastName": "Taylor",
    "birthday": new Date("1994-02-18"),
    "phoneNumber": "444-555-6666",
    lastLogin: new Date("2001-01-01")
  },
  {
    "id": 10,
    "createdAt": new Date("2023-12-10T12:09:00Z"),
    "deletedAt": new Date("2023-12-10T12:09:00Z"),
    "updatedAt": new Date("2023-12-10T12:09:00Z"),
    "email": "user10@gmail.com",
    "passwordHash": "abc",
    "firstName": "Henry",
    "lastName": "Davis",
    "birthday": new Date("1989-03-22"),
    "phoneNumber": "666-777-8888",
    lastLogin: new Date("2001-01-01")
  }
]



export function testUseUsers() {
  const [users, setUsers] = useState<User[]>(mockUsers)

  function addUser(user: User) {
    setUsers(prev => [...prev, user])
  }

  return { users, addUser }
}