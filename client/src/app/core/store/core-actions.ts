import { createAction, props } from '@ngrx/store';
import { IInitStoreParams } from 'src/app/core/store/core-store-params';
import { IUser } from 'src/app/core/typings/IUser';

export const loadData = createAction(
  '[CORE] Load data'
);

export const initStore = createAction(
  '[CORE] Init store',
  props<IInitStoreParams>()
);

export const signIn = createAction(
  '[CORE] Sign In',
  props<{ user: IUser }>()
);

