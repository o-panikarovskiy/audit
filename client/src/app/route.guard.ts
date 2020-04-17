import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, CanLoad, Route, Router, RouterStateSnapshot, UrlSegment } from '@angular/router';
import { Observable } from 'rxjs';
import { filter, flatMap, map, take, tap } from 'rxjs/operators';
import { AuthService } from 'src/app/core/services/auth.service';
import { CoreStoreService } from 'src/app/core/services/core.store.service';
import { UserRole } from 'src/app/core/typings/UserRole';

@Injectable({
  providedIn: 'root'
})
export class RouteGuard implements CanActivate, CanLoad {

  constructor(
    private router: Router,
    private authService: AuthService,
    private coreStore: CoreStoreService
  ) {

  }

  canLoad(route: Route, segments: UrlSegment[]): Observable<boolean> {
    return this.checkRouteAccess(route).pipe(
      take(1),
      tap((isOk) => {
        if (!isOk) {
          route.data.redirect(this.router);
        }
      })
    );
  }


  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<boolean> {
    throw new Error('Method not implemented.');
  }

  private checkRouteAccess(route: Route | ActivatedRouteSnapshot): Observable<boolean> {
    return this.coreStore.isInited$.pipe(
      filter((isInited) => isInited),
      flatMap(() => this.coreStore.user$),
      map(user => {
        const accessLevel = route.data?.accessLevel;
        const userRole = user ? user.role : UserRole.Anonymous;
        return this.authService.authorize(accessLevel, userRole);
      })
    );
  }
}
