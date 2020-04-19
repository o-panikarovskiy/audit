import { createFeatureSelector, createSelector } from '@ngrx/store';
import { coreFeatureKey, CoreState } from 'src/app/core/store/core-state';


export const getSchemeStateSelector = createFeatureSelector<CoreState>(coreFeatureKey);

export const getUserSelector = createSelector(getSchemeStateSelector, state => state.user);
export const getInitSelector = createSelector(getSchemeStateSelector, state => state.isInited);
