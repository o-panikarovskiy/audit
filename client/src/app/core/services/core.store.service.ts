import { Injectable } from '@angular/core';
import { select, Store } from '@ngrx/store';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';
import { AuthService } from 'src/app/core/services/auth.service';
import * as coreActions from 'src/app/core/store/core-actions';
import * as coreSelectors from 'src/app/core/store/core-selectors';
import { CoreState } from 'src/app/core/store/core-state';
import { IUser } from 'src/app/core/typings/IUser';


@Injectable()
export class CoreStoreService {
  public readonly user$: Observable<IUser>;
  public readonly isInited$: Observable<boolean>;

  constructor(
    private authService: AuthService,
    private readonly store: Store<CoreState>
  ) {
    this.user$ = this.store.pipe(select(coreSelectors.getUserSelector));
    this.isInited$ = this.store.pipe(select(coreSelectors.getInitSelector));
  }

  public initStore() {
    return this.authService.checkSession().pipe(
      tap((user) => {
        this.store.dispatch(coreActions.initStore({ user }));
      })
    ).toPromise();
  }

  public signIn(username: string, password: string) {
    return this.authService.singIn(username, password).pipe(
      tap((user) => {
        this.store.dispatch(coreActions.signIn({ user }));
      })
    );
  }
}
