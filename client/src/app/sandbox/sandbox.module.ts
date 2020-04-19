import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SocketComponent } from './socket/socket.component';
import { SandBoxRoutingModule } from 'src/app/sandbox/sandbox-routing.module';
import { ReactiveFormsModule } from '@angular/forms';


@NgModule({
  declarations: [SocketComponent],
  imports: [
    CommonModule,
    SandBoxRoutingModule,
    ReactiveFormsModule,
  ]
})
export class SandboxModule { }
