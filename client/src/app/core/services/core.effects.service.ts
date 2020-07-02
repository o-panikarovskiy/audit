
import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { tap } from 'rxjs/operators';
import { WebSocketService } from 'src/app/core/services/socket.service';
import * as coreActions from 'src/app/core/store/core-actions';
import { IUser } from 'src/app/core/typings/IUser';

@Injectable()
export class CoreEffects {
  constructor(
    private readonly actions$: Actions,
    private readonly ws: WebSocketService
  ) { }

  public signIn$ = createEffect(() =>
    this.actions$.pipe(
      ofType(coreActions.initStore, coreActions.signIn),
      tap(({ user }) => {
        if (user) {
          this.onUserLogin(user);
        } else {
          this.onUserLogout();
        }
      })
    ),
    { dispatch: false }
  );

  public signOut$ = createEffect(() =>
    this.actions$.pipe(
      ofType(coreActions.signOut),
      tap(() => {
        this.onUserLogout();
      })
    ),
    { dispatch: false }
  );

  private onUserLogin(user: IUser) {
    this.ws.connect('app');
  }

  private onUserLogout() {
    this.ws.disconnect('app');
  }
}
