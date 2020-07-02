import { NgModule } from '@angular/core';
import { Router, RouterModule, Routes } from '@angular/router';
import { AccessLevel } from 'src/app/core/typings/UserRole';
import { RouteGuard } from 'src/app/route.guard';


const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    redirectTo: 'home'
  },
  {
    path: 'home',
    loadChildren: () => import('./home/home.module').then(mod => mod.HomeModule),
    canLoad: [RouteGuard],
    data: {
      accessLevel: AccessLevel.User,
      redirect: (router: Router) => { router.navigate(['auth', 'sign-in']); }
    }
  },
  {
    path: 'auth',
    loadChildren: () => import('./auth/auth.module').then(mod => mod.AuthModule),
  },
  {
    path: 'sandbox',
    loadChildren: () => import('./sandbox/sandbox.module').then(mod => mod.SandboxModule),
  },
  {
    path: '**',
    redirectTo: ''
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
