
import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { EMPTY, forkJoin } from 'rxjs';
import { catchError, map, switchMap } from 'rxjs/operators';
import { AuthService } from 'src/app/core/services/auth.service';
import * as coreActions from 'src/app/core/store/core-actions';

@Injectable()
export class CoreEffects {
  constructor(
    private readonly actions$: Actions,
    private readonly authService: AuthService,
  ) { }

  public loadStore$ = createEffect(() =>
    this.actions$.pipe(
      ofType(coreActions.loadData),
      switchMap(() =>
        forkJoin(
          [
            this.authService.checkSession(),
          ]
        ).pipe(
          map(([user]) => {
            return coreActions.initStore({ user });
          }),
          catchError(() => EMPTY)
        )
      )
    )
  );
}
