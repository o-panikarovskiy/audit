import { IUser } from 'src/app/core/typings/IUser';

export const coreFeatureKey = 'core';

export interface CoreState {
  readonly user?: IUser;
  readonly isInited: boolean;
}
