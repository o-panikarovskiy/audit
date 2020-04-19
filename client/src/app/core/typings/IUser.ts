import { UserRole } from 'src/app/core/typings/UserRole';

export interface IUser {
  readonly id: string;
  readonly role: UserRole;
}
