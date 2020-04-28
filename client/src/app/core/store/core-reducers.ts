import { Action, createReducer, on } from '@ngrx/store';
import * as actions from 'src/app/core/store/core-actions';
import { CoreState } from 'src/app/core/store/core-state';


export const initialState: CoreState = {
  user: void 0,
  isInited: false,
};

const coreReducer = createReducer(
  initialState,

  on(actions.initStore, actions.signIn, (state, { user }) => {
    return {
      ...state,
      user,
      isInited: true
    };
  }),

);

export function getCoreReducer(state = initialState, action: Action): CoreState {
  return coreReducer(state, action);
}
