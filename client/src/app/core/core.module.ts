import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { APP_INITIALIZER, NgModule, Optional, SkipSelf } from '@angular/core';
import { EffectsModule } from '@ngrx/effects';
import { StoreModule } from '@ngrx/store';
import { AuthService } from 'src/app/core/services/auth.service';
import { CoreEffects } from 'src/app/core/services/core.effects.service';
import { CoreStoreService } from 'src/app/core/services/core.store.service';
import { getCoreReducer } from 'src/app/core/store/core-reducers';
import { coreFeatureKey } from 'src/app/core/store/core-state';

@NgModule({
  imports: [
    CommonModule,
    HttpClientModule,
    StoreModule.forFeature(coreFeatureKey, getCoreReducer),
    EffectsModule.forFeature([
      CoreEffects
    ])
  ],
  providers: [
    AuthService,
    CoreStoreService,
    {
      provide: APP_INITIALIZER,
      useFactory: (cs: CoreStoreService) => () => cs.initStore(),
      deps: [CoreStoreService],
      multi: true
    }
  ],
})
export class CoreModule {
  constructor(@Optional() @SkipSelf() parentModule: CoreModule) {
    if (parentModule) {
      throw new Error('CoreModule is already loaded. Import it in the AppModule only');
    }
  }
}
