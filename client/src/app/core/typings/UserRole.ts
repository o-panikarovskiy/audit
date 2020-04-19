export enum UserRole {
  Anonymous = 1,
  User = 2,
  Admin = 4,
}

// tslint:disable:no-bitwise
export enum AccessLevel {
  Anonymous = UserRole.Anonymous,
  Public = UserRole.Anonymous | UserRole.User | UserRole.Admin,
  User = UserRole.User | UserRole.Admin,
  Admin = UserRole.Admin,
}
