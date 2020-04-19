import { createAction, props } from '@ngrx/store';
import { IInitStoreParams } from 'src/app/core/store/core-store-params';

export const loadData = createAction(
  '[CORE] Load data'
);

export const initStore = createAction(
  '[CORE] Init store',
  props<IInitStoreParams>()
);
