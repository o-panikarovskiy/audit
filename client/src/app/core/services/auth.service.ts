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

  public authorize(accessLevel: AccessLevel, userRole: UserRole): boolean {
    // tslint:disable-next-line:no-bitwise
    return (accessLevel & userRole) > 0;
  }
}
