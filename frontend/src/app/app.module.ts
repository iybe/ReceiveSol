import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { FormsModule } from '@angular/forms';
import { ReactiveFormsModule } from '@angular/forms';

import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { LoginComponent } from './pages/login/login.component';
import { RegisterComponent } from './pages/register/register.component';
import { MenuComponent } from './pages/menu/menu.component';
import { PaymentLinksComponent } from './pages/payment-links/payment-links.component';
import { AccountComponent } from './pages/account/account.component';
import { NewAccountComponent } from './modals/new-account/new-account.component';
import { NewPaymentLinksComponent } from './modals/new-payment-link/new-payment-link.component';
import { CreatePermalinkComponent } from './pages/create-permalink/create-permalink.component';

import { NZ_I18N } from 'ng-zorro-antd/i18n';
import { en_US } from 'ng-zorro-antd/i18n';
import { registerLocaleData } from '@angular/common';
import en from '@angular/common/locales/en';

import { NzInputModule } from 'ng-zorro-antd/input';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzMenuModule } from 'ng-zorro-antd/menu';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzNotificationModule } from 'ng-zorro-antd/notification';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzPopoverModule } from 'ng-zorro-antd/popover';
import { NzDividerModule } from 'ng-zorro-antd/divider';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzBadgeModule } from 'ng-zorro-antd/badge';
import { NzDropDownModule } from 'ng-zorro-antd/dropdown';
import { NzListModule } from 'ng-zorro-antd/list';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { OptionsPermalinkComponent } from './pages/options-permalink/options-permalink.component';

registerLocaleData(en);

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    RegisterComponent,
    MenuComponent,
    PaymentLinksComponent,
    AccountComponent,
    NewAccountComponent,
    NewPaymentLinksComponent,
    CreatePermalinkComponent,
    OptionsPermalinkComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,

    FormsModule,
    ReactiveFormsModule,

    HttpClientModule,
    BrowserAnimationsModule,

    NzInputModule,
    NzButtonModule,
    NzMenuModule,
    NzIconModule,
    NzNotificationModule,
    NzTableModule,
    NzPopoverModule,
    NzDividerModule,
    NzModalModule,
    NzBadgeModule,
    NzDropDownModule,
    NzListModule,
    NzSelectModule,
  ],
  providers: [{ provide: NZ_I18N, useValue: en_US }],
  bootstrap: [AppComponent],
})
export class AppModule {}
