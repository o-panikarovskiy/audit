import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { IUser } from 'src/app/core/typings/IUser';
import { AccessLevel, UserRole } from 'src/app/core/typings/UserRole';

@Injectable()
export class AuthService {
  constructor(private http: HttpClient) { }

  public checkSession(): Observable<IUser | undefined> {
    return of({ id: '1', role: UserRole.Anonymous });
  }

  public authorize(accessLevel: AccessLevel, userRole: UserRole): boolean {
    // tslint:disable-next-line:no-bitwise
    return (accessLevel & userRole) > 0;
  }
}
