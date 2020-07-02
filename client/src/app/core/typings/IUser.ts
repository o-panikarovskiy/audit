import { UserRole } from 'src/app/core/typings/UserRole';

export interface IUser {
  readonly id: string;
  readonly email: string;
  readonly name: string;
  readonly created: number | Date;
  readonly role: UserRole;
}
