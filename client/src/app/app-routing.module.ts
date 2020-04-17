import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';


const routes: Routes = [
  {
    path: 'sandbox',
    loadChildren: () => import('./sandbox/sandbox.module').then(mod => mod.SandboxModule),
  },
  {
    path: '**',
    redirectTo: 'sandbox'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
