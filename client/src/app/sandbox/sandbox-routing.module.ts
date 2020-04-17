import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { SocketComponent } from 'src/app/sandbox/socket/socket.component';


const routes: Routes = [
  {
    path: 'ws',
    pathMatch: 'full',
    component: SocketComponent
  },
  {
    path: '**',
    redirectTo: 'ws'
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class SandBoxRoutingModule { }
