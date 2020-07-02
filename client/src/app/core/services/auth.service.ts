import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { BackendService } from 'src/app/core/services/backend.service';
import { IUser } from 'src/app/core/typings/IUser';
import { AccessLevel, UserRole } from 'src/app/core/typings/UserRole';

@Injectable()
export class AuthService {
  constructor(private be: BackendService) { }

  public checkSession(): Observable<IUser> {
    return this.be.get('auth/check').pipe(
      catchError(() => of(void 0))
    );
  }

  public singIn(username: string, password: string): Observable<IUser> {
    return this.be.post('auth/signin', { username, password });
  }

  public singUp(email: string, password: string): Observable<void> {
    return this.be.post('auth/signup', { email, password });
  }

  public signUpConfirm(token: string): Observable<IUser> {
    return this.be.get(`auth/confirm/${token}`);
  }

  public singOut(): Observable<void> {
    return this.be.post('auth/signout');
  }

  public authorize(accessLevel: AccessLevel, userRole: UserRole): boolean {
    // tslint:disable-next-line:no-bitwise
    return (accessLevel & userRole) > 0;
  }
}
