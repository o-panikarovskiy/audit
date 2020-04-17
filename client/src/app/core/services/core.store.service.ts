import { Injectable } from '@angular/core';
import { Actions, ofType } from '@ngrx/effects';
import { select, Store } from '@ngrx/store';
import { Observable } from 'rxjs';
import { take } from 'rxjs/operators';
import * as coreActions from 'src/app/core/store/core-actions';
import * as coreSelectors from 'src/app/core/store/core-selectors';
import { CoreState } from 'src/app/core/store/core-state';
import { IUser } from 'src/app/core/typings/IUser';


@Injectable()
export class CoreStoreService {
  public readonly user$: Observable<IUser>;
  public readonly isInited$: Observable<boolean>;

  constructor(
    private actions$: Actions,
    private readonly store: Store<CoreState>
  ) {
    this.user$ = this.store.pipe(select(coreSelectors.getUserSelector));
    this.isInited$ = this.store.pipe(select(coreSelectors.getInitSelector));
  }

  public initStore() {
    this.store.dispatch(coreActions.loadData());
    return this.waitUntilStoreInit();
  }

  private waitUntilStoreInit() {
    return this.actions$.pipe(
      ofType(coreActions.initStore),
      take(1)
    );
  }
}
