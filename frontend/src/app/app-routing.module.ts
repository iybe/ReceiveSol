import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { LoginComponent } from './pages/login/login.component';
import { RegisterComponent } from './pages/register/register.component';
import { MenuComponent } from './pages/menu/menu.component';
import { PaymentLinksComponent } from './pages/payment-links/payment-links.component';
import { AccountComponent } from './pages/account/account.component';
import { CreatePermalinkComponent } from './pages/create-permalink/create-permalink.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: 'register', component: RegisterComponent },
  { path: 'receivesol', component: MenuComponent },

  { path: 'permalink/:userId', component: CreatePermalinkComponent },

  {
    path: 'receivesol/account',
    component: AccountComponent,
  },
  {
    path: 'receivesol/payment-links',
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
