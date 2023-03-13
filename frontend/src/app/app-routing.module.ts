import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { LoginComponent } from './pages/login/login.component';
import { RegisterComponent } from './pages/register/register.component';
import { MenuComponent } from './pages/menu/menu.component';
import { PaymentLinksComponent } from './pages/payment-links/payment-links.component';
import { AccountComponent } from './pages/account/account.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: 'register', component: RegisterComponent },
  { path: 'Receivesol', component: MenuComponent },

  {
    path: 'Receivesol/account',
    component: AccountComponent,
  },
  {
    path: 'Receivesol/payment-links',
    component: PaymentLinksComponent,
  },

  { path: '', redirectTo: 'login', pathMatch: 'full' },
  { path: '**', redirectTo: 'login' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
